module github.com/k8snetworkplumbingwg/sriov-network-device-plugin

go 1.17

require (
	github.com/Mellanox/rdmamap v1.0.0
	github.com/Mellanox/sriovnet v1.0.3
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/jaypipes/ghw v0.6.0
	github.com/jaypipes/pcidb v0.5.0
	github.com/k8snetworkplumbingwg/network-attachment-definition-client v1.1.1-0.20201119153432-9d213757d22d
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.9.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/vishvananda/netlink v1.1.1-0.20211101163509-b10eb8fe5cf6
	google.golang.org/grpc v1.28.1
	k8s.io/kubelet v0.18.1
)

require (
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)

replace (
	github.com/containernetworking/cni => github.com/containernetworking/cni v0.8.1
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
)
