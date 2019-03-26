package main

import (
	"fmt"
	"path/filepath"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
	registerapi "k8s.io/kubernetes/pkg/kubelet/apis/pluginregistration/v1"
)

func TestSriovdp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sriov Device Plugin Suite")
}

type mockStreamA struct{ grpc.ServerStream }
type mockStreamB struct{ grpc.ServerStream }

func (msa *mockStreamA) Send(resp *pluginapi.ListAndWatchResponse) error {
	return nil
}

func (msb *mockStreamB) Send(resp *pluginapi.ListAndWatchResponse) error {
	return fmt.Errorf("Error. Cannot update device states")
}

var _ = Describe("Device Plugin APIs", func() {

	var err error
	var devices []string
	var ctx context.Context
	var empty *pluginapi.Empty
	var sm *sriovManager

	BeforeEach(func() {
		ctx = context.Background()
		empty = new(pluginapi.Empty)
		sm = &sriovManager{
			devices:     make(map[string]pluginapi.Device),
			socketFile:  fmt.Sprintf("%s.sock", pluginEndpointPrefix),
			termSignal:  make(chan bool, 1),
			stopWatcher: make(chan bool),
		}

		devices = []string{"0000:00:01.0", "0000:00:01.1", "0000:00:02.0"}
		for _, dev := range devices {
			sm.devices[dev] = pluginapi.Device{ID: dev, Health: pluginapi.Healthy}
		}
	})

	It("Check GetInfo", func() {
		giRqt := new(registerapi.InfoRequest)
		response := new(registerapi.PluginInfo)

		response, err = sm.GetInfo(ctx, giRqt)

		Expect(err).NotTo(HaveOccurred())
		Expect(*response).Should(BeAssignableToTypeOf(registerapi.PluginInfo{}))
		Expect(response.Type).To(Equal(registerapi.DevicePlugin))
		Expect(response.Name).To(Equal(resourceName))
		Expect(response.Endpoint).To(Equal(filepath.Join(pluginMountPath, sm.socketFile)))
		Expect(response.SupportedVersions).To(Equal([]string{"v1beta1"}))
	})

	It("Check NotifyRegistrationStatus", func() {
		rStat := new(registerapi.RegistrationStatus)
		rStat.PluginRegistered = true
		response := new(registerapi.RegistrationStatusResponse)

		response, err = sm.NotifyRegistrationStatus(ctx, rStat)

		Expect(err).NotTo(HaveOccurred())
		Expect(*response).Should(BeAssignableToTypeOf(registerapi.RegistrationStatusResponse{}))
	})

	It("Check PreStartContainer", func() {
		psRqt := new(pluginapi.PreStartContainerRequest)
		response := new(pluginapi.PreStartContainerResponse)

		response, err = sm.PreStartContainer(ctx, psRqt)

		Expect(err).NotTo(HaveOccurred())
		Expect(*response).Should(BeAssignableToTypeOf(pluginapi.PreStartContainerResponse{}))
	})

	It("Check GetDevicePluginOptions", func() {
		response := new(pluginapi.DevicePluginOptions)

		response, err = sm.GetDevicePluginOptions(ctx, empty)

		Expect(err).NotTo(HaveOccurred())
		Expect(*response).Should(BeAssignableToTypeOf(pluginapi.DevicePluginOptions{}))
		Expect(response.PreStartRequired).To(Equal(false))
	})

	It("Check Allocate success", func() {
		request := new(pluginapi.AllocateRequest)
		response := new(pluginapi.AllocateResponse)

		request = &pluginapi.AllocateRequest{
			ContainerRequests: []*pluginapi.ContainerAllocateRequest{
				{DevicesIDs: devices},
			},
		}

		response, err = sm.Allocate(ctx, request)

		Expect(err).NotTo(HaveOccurred())

		envmap := make(map[string]string)
		envmap["SRIOV-VF-PCI-ADDR"] = "0000:00:01.0,0000:00:01.1,0000:00:02.0,"
		Expect(response.ContainerResponses[0].Envs).To(Equal(envmap))
	})

	It("Check Allocate with non-existing device", func() {
		request := new(pluginapi.AllocateRequest)
		response := new(pluginapi.AllocateResponse)

		request = &pluginapi.AllocateRequest{
			ContainerRequests: []*pluginapi.ContainerAllocateRequest{
				{DevicesIDs: []string{"0000:00:02.1"}},
			},
		}

		response, err = sm.Allocate(ctx, request)

		Expect(response).To(BeNil())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).
			To(ContainSubstring("Invalid allocation request with non-existing device 0000:00:02.1"))
	})

	It("Check Allocate with unhealthy device", func() {
		request := new(pluginapi.AllocateRequest)
		response := new(pluginapi.AllocateResponse)
		sm.devices["0000:00:02.0"] = pluginapi.Device{ID: "0000:00:02.0", Health: "Unhealthy"}
		request = &pluginapi.AllocateRequest{
			ContainerRequests: []*pluginapi.ContainerAllocateRequest{
				{DevicesIDs: devices},
			},
		}

		response, err = sm.Allocate(ctx, request)

		Expect(response).To(BeNil())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).
			To(ContainSubstring("Invalid allocation request with unhealthy device 0000:00:02.0"))
	})

	It("Check ListAndWatch success", func() {
		sm.termSignal <- true
		stream := &mockStreamA{}

		err = sm.ListAndWatch(empty, stream)

		Expect(err).NotTo(HaveOccurred())
	})

	It("Check ListAndWatch send device states fail", func() {
		sm.grpcServer = grpc.NewServer()
		sm.termSignal <- true
		stream := &mockStreamB{}

		err = sm.ListAndWatch(empty, stream)

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).
			To(ContainSubstring("Error. Cannot update device states"))
	})
})
