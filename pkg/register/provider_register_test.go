package register

import (
  "fmt"
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
)

func ExampleProviderRegistrar_GetNodeID() {
	fmt.Println(pr.GetNodeID())
	// Output: AA
}

func ExampleProviderRegistrar_GetAddress() {
	fmt.Println(pr.GetAddress())
	// Output: AA
}

func ExampleProviderRegistrar_GetRegionCode() {
	fmt.Println(pr.GetRegionCode())
	// Output: AA
}

func ExampleProviderRegistrar_GetNetworkInfoGateway() {
	fmt.Println(pr.GetNetworkInfoGateway())
	// Output: AA
}

func ExampleProviderRegistrar_GetNetworkInfoClient() {
	fmt.Println(pr.GetNetworkInfoClient())
	// Output: AA
}

func ExampleProviderRegistrar_GetNetworkInfoAdmin() {
	fmt.Println(pr.GetNetworkInfoAdmin())
	// Output: AA
}

func ExampleProviderRegistrar_GetRootSigningKey() {
	v, err := pr.GetRootSigningKey()
	fmt.Println(v, err)
	// Output: <nil> incorrect secp256k1 public key length: 32
}

func ExampleProviderRegistrar_GetSigningKey() {
	v, err := pr.GetSigningKey()
	fmt.Println(v, err)
	// Output: <nil> incorrect secp256k1 public key length: 32
}
