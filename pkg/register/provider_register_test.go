package register

import (
  "fmt"

  "github.com/ConsenSys/fc-retrieval-common/pkg/request"
)

var pr = NewProviderRegister(
    "AA",
    "AA",
    PrivKey,
    PrivKey,
    "AA",
    "AA",
    "AA",
    "AA",
    request.NewHttpCommunicator())

func ExampleProviderRegisterOperations_GetNodeID() {
	fmt.Println(pr.GetNodeID())
	// Output: AA
}

func ExampleProviderRegisterOperations_GetAddress() {
	fmt.Println(pr.GetAddress())
	// Output: AA
}

func ExampleProviderRegisterOperations_GetRegionCode() {
	fmt.Println(pr.GetRegionCode())
	// Output: AA
}

func ExampleProviderRegisterOperations_GetNetworkInfoGateway() {
	fmt.Println(pr.GetNetworkInfoGateway())
	// Output: AA
}

func ExampleProviderRegisterOperations_GetNetworkInfoClient() {
	fmt.Println(pr.GetNetworkInfoClient())
	// Output: AA
}

func ExampleProviderRegisterOperations_GetNetworkInfoAdmin() {
	fmt.Println(pr.GetNetworkInfoAdmin())
	// Output: AA
}

func ExampleProviderRegisterOperations_GetRootSigningKey() {
	v, err := pr.GetRootSigningKey()
	fmt.Println(v, err)
	// Output: <nil> incorrect secp256k1 public key length: 32
}

func ExampleProviderRegisterOperations_GetSigningKey() {
	v, err := pr.GetSigningKey()
	fmt.Println(v, err)
	// Output: <nil> incorrect secp256k1 public key length: 32
}
