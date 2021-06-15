package register

import (
  "fmt"

  "github.com/ConsenSys/fc-retrieval-common/pkg/request"
)

const (
	PrivKey = "015ed053eab6fdf18c03954373ff7f89089992017d56beb8b05305b19800d6afe0"
)

var gr = NewGatewayRegister(
    "AA",
    "AA",
    PrivKey,
    PrivKey,
    "AA",
    "AA",
    "AA",
    "AA",
    "AA",
    request.NewHttpCommunicator())

func ExampleGatewayRegistrar_GetNodeID() {
	fmt.Println(gr.GetNodeID())
	// Output: AA
}

func ExampleGatewayRegistrar_GetAddress() {
	fmt.Println(gr.GetAddress())
	// Output: AA
}

func ExampleGatewayRegistrar_GetRegionCode() {
	fmt.Println(gr.GetRegionCode())
	// Output: AA
}

func ExampleGatewayRegistrar_GetNetworkInfoGateway() {
	fmt.Println(gr.GetNetworkInfoGateway())
	// Output: AA
}

func ExampleGatewayRegistrar_GetNetworkInfoClient() {
	fmt.Println(gr.GetNetworkInfoClient())
	// Output: AA
}

func ExampleGatewayRegistrar_GetNetworkInfoAdmin() {
	fmt.Println(gr.GetNetworkInfoAdmin())
	// Output: AA
}

func ExampleGatewayRegistrar_GetRootSigningKey() {
	v, err := gr.GetRootSigningKey()
	fmt.Println(v, err)
	// Output: <nil> incorrect secp256k1 public key length: 32
}

func ExampleGatewayRegistrar_GetSigningKey() {
	v, err := gr.GetSigningKey()
	fmt.Println(v, err)
	// Output: <nil> incorrect secp256k1 public key length: 32
}

func ExampleGatewayRegistrar_RegisterGateway() {
	err := gr.RegisterGateway("")
	fmt.Println(err)
	// Output:
	// Post "/registers/gateway": unsupported protocol scheme ""
}
