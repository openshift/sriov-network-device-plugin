module github.com/k8snetworkplumbingwg/sriov-network-device-plugin

go 1.19

require (
	github.com/Mellanox/rdmamap v1.0.0
	github.com/Mellanox/sriovnet v1.0.3
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/jaypipes/ghw v0.6.0
	github.com/jaypipes/pcidb v0.5.0
	github.com/k8snetworkplumbingwg/govdpa v0.1.3
	github.com/k8snetworkplumbingwg/network-attachment-definition-client v1.1.1-0.20201119153432-9d213757d22d
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.20.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.0
	github.com/vishvananda/netlink v1.1.1-0.20211101163509-b10eb8fe5cf6
	google.golang.org/grpc v1.47.0
	k8s.io/kubelet v0.25.15
)

require (
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/containernetworking/cni v0.7.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.8.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mitchellh/go-homedir v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nxadm/tail v1.4.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.4.1 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/vishvananda/netns v0.0.0-20200728191858-db3c7e526aae // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220502173005-c8bf987b8c21 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	howett.net/plist v0.0.0-20181124034731-591f970eefbb // indirect
	k8s.io/api v0.25.15 // indirect
	k8s.io/apimachinery v0.25.15 // indirect
	k8s.io/client-go v0.25.15 // indirect
	k8s.io/klog/v2 v2.70.1 // indirect
	k8s.io/kube-openapi v0.0.0-20220803162953-67bda5d908f1 // indirect
	k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	github.com/containernetworking/cni => github.com/containernetworking/cni v0.8.1
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
)
