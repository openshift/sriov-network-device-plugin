/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package master

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	certificatesapiv1beta1 "k8s.io/api/certificates/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilnet "k8s.io/apimachinery/pkg/util/net"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/version"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/options"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	etcdtesting "k8s.io/apiserver/pkg/storage/etcd/testing"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
	restclient "k8s.io/client-go/rest"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
	"k8s.io/kubernetes/pkg/api/testapi"
	api "k8s.io/kubernetes/pkg/apis/core"
	kubeletclient "k8s.io/kubernetes/pkg/kubelet/client"
	"k8s.io/kubernetes/pkg/master/reconcilers"
	certificatesrest "k8s.io/kubernetes/pkg/registry/certificates/rest"
	corerest "k8s.io/kubernetes/pkg/registry/core/rest"
	"k8s.io/kubernetes/pkg/registry/registrytest"
	kubeversion "k8s.io/kubernetes/pkg/version"

	"github.com/stretchr/testify/assert"
)

// setUp is a convience function for setting up for (most) tests.
func setUp(t *testing.T) (*etcdtesting.EtcdTestServer, Config, *assert.Assertions) {
	server, storageConfig := etcdtesting.NewUnsecuredEtcd3TestClientServer(t)

	config := &Config{
		GenericConfig: genericapiserver.NewConfig(legacyscheme.Codecs),
		ExtraConfig: ExtraConfig{
			APIResourceConfigSource: DefaultAPIResourceConfigSource(),
			APIServerServicePort:    443,
			MasterCount:             1,
			EndpointReconcilerType:  reconcilers.MasterCountReconcilerType,
		},
	}

	resourceEncoding := serverstorage.NewDefaultResourceEncodingConfig(legacyscheme.Scheme)
	storageFactory := serverstorage.NewDefaultStorageFactory(*storageConfig, testapi.StorageMediaType(), legacyscheme.Codecs, resourceEncoding, DefaultAPIResourceConfigSource(), nil)

	etcdOptions := options.NewEtcdOptions(storageConfig)
	// unit tests don't need watch cache and it leaks lots of goroutines with etcd testing functions during unit tests
	etcdOptions.EnableWatchCache = false
	err := etcdOptions.ApplyWithStorageFactoryTo(storageFactory, config.GenericConfig)
	if err != nil {
		t.Fatal(err)
	}

	kubeVersion := kubeversion.Get()
	config.GenericConfig.Version = &kubeVersion
	config.ExtraConfig.StorageFactory = storageFactory
	config.GenericConfig.LoopbackClientConfig = &restclient.Config{APIPath: "/api", ContentConfig: restclient.ContentConfig{NegotiatedSerializer: legacyscheme.Codecs}}
	config.GenericConfig.PublicAddress = net.ParseIP("192.168.10.4")
	config.GenericConfig.LegacyAPIGroupPrefixes = sets.NewString("/api")
	config.GenericConfig.LoopbackClientConfig = &restclient.Config{APIPath: "/api", ContentConfig: restclient.ContentConfig{NegotiatedSerializer: legacyscheme.Codecs}}
	config.ExtraConfig.KubeletClientConfig = kubeletclient.KubeletClientConfig{Port: 10250}
	config.ExtraConfig.ProxyTransport = utilnet.SetTransportDefaults(&http.Transport{
		DialContext:     func(ctx context.Context, network, addr string) (net.Conn, error) { return nil, nil },
		TLSClientConfig: &tls.Config{},
	})

	// set fake SecureServingInfo because the listener port is needed for the kubernetes service
	config.GenericConfig.SecureServing = &genericapiserver.SecureServingInfo{Listener: fakeLocalhost443Listener{}}

	clientset, err := kubernetes.NewForConfig(config.GenericConfig.LoopbackClientConfig)
	if err != nil {
		t.Fatalf("unable to create client set due to %v", err)
	}
	config.ExtraConfig.VersionedInformers = informers.NewSharedInformerFactory(clientset, config.GenericConfig.LoopbackClientConfig.Timeout)

	return server, *config, assert.New(t)
}

type fakeLocalhost443Listener struct{}

func (fakeLocalhost443Listener) Accept() (net.Conn, error) {
	return nil, nil
}

func (fakeLocalhost443Listener) Close() error {
	return nil
}

func (fakeLocalhost443Listener) Addr() net.Addr {
	return &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 443,
	}
}

// TestLegacyRestStorageStrategies ensures that all Storage objects which are using the generic registry Store have
// their various strategies properly wired up. This surfaced as a bug where strategies defined Export functions, but
// they were never used outside of unit tests because the export strategies were not assigned inside the Store.
func TestLegacyRestStorageStrategies(t *testing.T) {
	_, etcdserver, masterCfg, _ := newMaster(t)
	defer etcdserver.Terminate(t)

	storageProvider := corerest.LegacyRESTStorageProvider{
		StorageFactory:       masterCfg.ExtraConfig.StorageFactory,
		ProxyTransport:       masterCfg.ExtraConfig.ProxyTransport,
		KubeletClientConfig:  masterCfg.ExtraConfig.KubeletClientConfig,
		EventTTL:             masterCfg.ExtraConfig.EventTTL,
		ServiceIPRange:       masterCfg.ExtraConfig.ServiceIPRange,
		ServiceNodePortRange: masterCfg.ExtraConfig.ServiceNodePortRange,
		LoopbackClientConfig: masterCfg.GenericConfig.LoopbackClientConfig,
	}

	_, apiGroupInfo, err := storageProvider.NewLegacyRESTStorage(masterCfg.GenericConfig.RESTOptionsGetter)
	if err != nil {
		t.Errorf("failed to create legacy REST storage: %v", err)
	}

	// Any new stores with export logic will need to be added here:
	exceptions := registrytest.StrategyExceptions{
		// Only these stores should have an export strategy defined:
		HasExportStrategy: []string{
			"secrets",
			"limitRanges",
			"nodes",
			"podTemplates",
		},
	}

	strategyErrors := registrytest.ValidateStorageStrategies(apiGroupInfo.VersionedResourcesStorageMap["v1"], exceptions)
	for _, err := range strategyErrors {
		t.Error(err)
	}
}

func TestCertificatesRestStorageStrategies(t *testing.T) {
	_, etcdserver, masterCfg, _ := newMaster(t)
	defer etcdserver.Terminate(t)

	certStorageProvider := certificatesrest.RESTStorageProvider{}
	apiGroupInfo, _ := certStorageProvider.NewRESTStorage(masterCfg.ExtraConfig.APIResourceConfigSource, masterCfg.GenericConfig.RESTOptionsGetter)

	exceptions := registrytest.StrategyExceptions{
		HasExportStrategy: []string{
			"certificatesigningrequests",
		},
	}

	strategyErrors := registrytest.ValidateStorageStrategies(
		apiGroupInfo.VersionedResourcesStorageMap[certificatesapiv1beta1.SchemeGroupVersion.Version], exceptions)
	for _, err := range strategyErrors {
		t.Error(err)
	}
}

func newMaster(t *testing.T) (*Master, *etcdtesting.EtcdTestServer, Config, *assert.Assertions) {
	etcdserver, config, assert := setUp(t)

	master, err := config.Complete().New(genericapiserver.NewEmptyDelegate())
	if err != nil {
		t.Fatalf("Error in bringing up the master: %v", err)
	}

	return master, etcdserver, config, assert
}

// TestVersion tests /version
func TestVersion(t *testing.T) {
	s, etcdserver, _, _ := newMaster(t)
	defer etcdserver.Terminate(t)

	req, _ := http.NewRequest("GET", "/version", nil)
	resp := httptest.NewRecorder()
	s.GenericAPIServer.Handler.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Fatalf("expected http 200, got: %d", resp.Code)
	}

	var info version.Info
	err := json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(kubeversion.Get(), info) {
		t.Errorf("Expected %#v, Got %#v", kubeversion.Get(), info)
	}
}

type fakeEndpointReconciler struct{}

func (*fakeEndpointReconciler) ReconcileEndpoints(serviceName string, ip net.IP, endpointPorts []api.EndpointPort, reconcilePorts bool) error {
	return nil
}

func makeNodeList(nodes []string, nodeResources apiv1.NodeResources) *apiv1.NodeList {
	list := apiv1.NodeList{
		Items: make([]apiv1.Node, len(nodes)),
	}
	for i := range nodes {
		list.Items[i].Name = nodes[i]
		list.Items[i].Status.Capacity = nodeResources.Capacity
	}
	return &list
}

// TestGetNodeAddresses verifies that proper results are returned
// when requesting node addresses.
func TestGetNodeAddresses(t *testing.T) {
	assert := assert.New(t)

	fakeNodeClient := fake.NewSimpleClientset(makeNodeList([]string{"node1", "node2"}, apiv1.NodeResources{})).CoreV1().Nodes()
	addressProvider := nodeAddressProvider{fakeNodeClient}

	// Fail case (no addresses associated with nodes)
	nodes, _ := fakeNodeClient.List(metav1.ListOptions{})
	addrs, err := addressProvider.externalAddresses()

	assert.Error(err, "addresses should have caused an error as there are no addresses.")
	assert.Equal([]string(nil), addrs)

	// Pass case with External type IP
	nodes, _ = fakeNodeClient.List(metav1.ListOptions{})
	for index := range nodes.Items {
		nodes.Items[index].Status.Addresses = []apiv1.NodeAddress{{Type: apiv1.NodeExternalIP, Address: "127.0.0.1"}}
		fakeNodeClient.Update(&nodes.Items[index])
	}
	addrs, err = addressProvider.externalAddresses()
	assert.NoError(err, "addresses should not have returned an error.")
	assert.Equal([]string{"127.0.0.1", "127.0.0.1"}, addrs)
}

func TestGetNodeAddressesWithOnlySomeExternalIP(t *testing.T) {
	assert := assert.New(t)

	fakeNodeClient := fake.NewSimpleClientset(makeNodeList([]string{"node1", "node2", "node3"}, apiv1.NodeResources{})).CoreV1().Nodes()
	addressProvider := nodeAddressProvider{fakeNodeClient}

	// Pass case with 1 External type IP (index == 1) and nodes (indexes 0 & 2) have no External IP.
	nodes, _ := fakeNodeClient.List(metav1.ListOptions{})
	nodes.Items[1].Status.Addresses = []apiv1.NodeAddress{{Type: apiv1.NodeExternalIP, Address: "127.0.0.1"}}
	fakeNodeClient.Update(&nodes.Items[1])

	addrs, err := addressProvider.externalAddresses()
	assert.NoError(err, "addresses should not have returned an error.")
	assert.Equal([]string{"127.0.0.1"}, addrs)
}

func decodeResponse(resp *http.Response, obj interface{}) error {
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, obj); err != nil {
		return err
	}
	return nil
}

// Because we need to be backwards compatible with release 1.1, at endpoints
// that exist in release 1.1, the responses should have empty APIVersion.
func TestAPIVersionOfDiscoveryEndpoints(t *testing.T) {
	master, etcdserver, _, assert := newMaster(t)
	defer etcdserver.Terminate(t)

	server := httptest.NewServer(master.GenericAPIServer.Handler.GoRestfulContainer.ServeMux)

	// /api exists in release-1.1
	resp, err := http.Get(server.URL + "/api")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	apiVersions := metav1.APIVersions{}
	assert.NoError(decodeResponse(resp, &apiVersions))
	assert.Equal(apiVersions.APIVersion, "")

	// /api/v1 exists in release-1.1
	resp, err = http.Get(server.URL + "/api/v1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	resourceList := metav1.APIResourceList{}
	assert.NoError(decodeResponse(resp, &resourceList))
	assert.Equal(resourceList.APIVersion, "")

	// /apis exists in release-1.1
	resp, err = http.Get(server.URL + "/apis")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	groupList := metav1.APIGroupList{}
	assert.NoError(decodeResponse(resp, &groupList))
	assert.Equal(groupList.APIVersion, "")

	// /apis/extensions exists in release-1.1
	resp, err = http.Get(server.URL + "/apis/extensions")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	group := metav1.APIGroup{}
	assert.NoError(decodeResponse(resp, &group))
	assert.Equal(group.APIVersion, "")

	// /apis/extensions/v1beta1 exists in release-1.1
	resp, err = http.Get(server.URL + "/apis/extensions/v1beta1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	resourceList = metav1.APIResourceList{}
	assert.NoError(decodeResponse(resp, &resourceList))
	assert.Equal(resourceList.APIVersion, "")

	// /apis/autoscaling doesn't exist in release-1.1, so the APIVersion field
	// should be non-empty in the results returned by the server.
	resp, err = http.Get(server.URL + "/apis/autoscaling")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	group = metav1.APIGroup{}
	assert.NoError(decodeResponse(resp, &group))
	assert.Equal(group.APIVersion, "v1")

	// apis/autoscaling/v1 doesn't exist in release-1.1, so the APIVersion field
	// should be non-empty in the results returned by the server.

	resp, err = http.Get(server.URL + "/apis/autoscaling/v1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	resourceList = metav1.APIResourceList{}
	assert.NoError(decodeResponse(resp, &resourceList))
	assert.Equal(resourceList.APIVersion, "v1")

}

func TestNoAlphaVersionsEnabledByDefault(t *testing.T) {
	config := DefaultAPIResourceConfigSource()
	for gv, enable := range config.GroupVersionConfigs {
		if enable && strings.Contains(gv.Version, "alpha") {
			t.Errorf("Alpha API version %s enabled by default", gv.String())
		}
	}
}
