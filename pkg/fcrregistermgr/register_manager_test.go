package fcrregistermgr

import (
  "fmt"
  "reflect"
  "sync"
  "testing"
  "time"

  "github.com/stretchr/testify/assert"

  "github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
  "github.com/ConsenSys/fc-retrieval-common/pkg/register"
)
const fakeRegisterAPIURL = "fakeRegisterAPIURL"
var rm = NewFCRRegisterMgr(fakeRegisterAPIURL, false, false, 1*time.Second)

func TestFCRRegisterMgr_GetGateway(t *testing.T) {

	nodeID00, _ := nodeid.NewNodeIDFromHexString("00")
	nodeID01, _ := nodeid.NewNodeIDFromHexString("01")
	nodeID02, _ := nodeid.NewNodeIDFromHexString("02")
	nodeID5A, _ := nodeid.NewNodeIDFromHexString("5A")
	nodeIDFFFF, _ := nodeid.NewNodeIDFromHexString("FFFF")

	gr := map[string]register.GatewayRegister{}
	gr[nodeID00.ToString()] = register.GatewayRegister{NodeID: nodeID00.ToString()}
	gr[nodeID01.ToString()] = register.GatewayRegister{NodeID: nodeID01.ToString()}
	gr[nodeID02.ToString()] = register.GatewayRegister{NodeID: nodeID02.ToString()}
	gr[nodeID5A.ToString()] = register.GatewayRegister{NodeID: nodeID5A.ToString()}
	gr[nodeIDFFFF.ToString()] = register.GatewayRegister{NodeID: nodeIDFFFF.ToString()}

	type fields struct {
		start                     bool
		gatewayDiscv              bool
		registeredGatewaysMap     map[string]register.GatewayRegister
		registeredGatewaysMapLock sync.RWMutex
	}
	type args struct {
		id *nodeid.NodeID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *register.GatewayRegister
	}{
		{name: "getGateway not started returns nil",
			fields: fields{
				start:                     false,
				gatewayDiscv:              false,
				registeredGatewaysMap:     gr,
				registeredGatewaysMapLock: sync.RWMutex{},
			},
			args: struct{ id *nodeid.NodeID }{id: nodeID00},
			want: nil,
		},
		{name: "getGateway - gatewayDiscv not started returns nil",
			fields: fields{
				start:                     true,
				gatewayDiscv:              false,
				registeredGatewaysMap:     gr,
				registeredGatewaysMapLock: sync.RWMutex{},
			},
			args: struct{ id *nodeid.NodeID }{id: nodeID00},
			want: nil,
		},
		{name: "getGateway - started return nodeID00",
			fields: fields{
				start:                     true,
				gatewayDiscv:              true,
				registeredGatewaysMap:     gr,
				registeredGatewaysMapLock: sync.RWMutex{},
			},
			args: struct{ id *nodeid.NodeID }{id: nodeID00},
			want: &register.GatewayRegister{NodeID: nodeID00.ToString()},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mgr := &FCRRegisterMgr{
				start:                     tt.fields.start,
				gatewayDiscv:              tt.fields.gatewayDiscv,
				registeredGatewaysMap:     tt.fields.registeredGatewaysMap,
				registeredGatewaysMapLock: tt.fields.registeredGatewaysMapLock,
			}
			if got := mgr.GetGateway(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGateway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFCRRegisterMgr_GetAllGateways(t *testing.T) {
	nodeID00, _ := nodeid.NewNodeIDFromHexString("00")
	nodeID01, _ := nodeid.NewNodeIDFromHexString("01")
	nodeID02, _ := nodeid.NewNodeIDFromHexString("02")
	nodeID5A, _ := nodeid.NewNodeIDFromHexString("5A")
	nodeIDFFFF, _ := nodeid.NewNodeIDFromHexString("FFFF")

	gr := map[string]register.GatewayRegister{}
	gr[nodeID00.ToString()] = register.GatewayRegister{NodeID: nodeID00.ToString()}
	gr[nodeID01.ToString()] = register.GatewayRegister{NodeID: nodeID01.ToString()}
	gr[nodeID02.ToString()] = register.GatewayRegister{NodeID: nodeID02.ToString()}
	gr[nodeID5A.ToString()] = register.GatewayRegister{NodeID: nodeID5A.ToString()}
	gr[nodeIDFFFF.ToString()] = register.GatewayRegister{NodeID: nodeIDFFFF.ToString()}

	type fields struct {
		start                     bool
		gatewayDiscv              bool
		registeredGatewaysMap     map[string]register.GatewayRegister
		registeredGatewaysMapLock sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields

		want []register.GatewayRegister
	}{
		{name: "getGateway not started returns nil",
			fields: fields{
				start:                     false,
				gatewayDiscv:              false,
				registeredGatewaysMap:     gr,
				registeredGatewaysMapLock: sync.RWMutex{},
			},

			want: nil,
		},
		{name: "getGateway - gatewayDiscv not started returns nil",
			fields: fields{
				start:                     true,
				gatewayDiscv:              false,
				registeredGatewaysMap:     gr,
				registeredGatewaysMapLock: sync.RWMutex{},
			},

			want: nil,
		},
		{name: "getGateway - started return nodeID00",
			fields: fields{
				start:                     true,
				gatewayDiscv:              true,
				registeredGatewaysMap:     gr,
				registeredGatewaysMapLock: sync.RWMutex{},
			},

			want: []register.GatewayRegister{

				{NodeID: nodeID02.ToString()},
				{NodeID: nodeID5A.ToString()},
				{NodeID: nodeIDFFFF.ToString()},
				{NodeID: nodeID00.ToString()},
				{NodeID: nodeID01.ToString()},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mgr := &FCRRegisterMgr{
				start:                     tt.fields.start,
				gatewayDiscv:              tt.fields.gatewayDiscv,
				registeredGatewaysMap:     tt.fields.registeredGatewaysMap,
				registeredGatewaysMapLock: tt.fields.registeredGatewaysMapLock,
			}
			got := mgr.GetAllGateways()
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}

func ExampleFCRRegisterMgr_GetRegisteredGateways() {
  v, err := rm.GetRegisteredGateways()
  fmt.Println(v, err)
  // Output:
  // [] Get "fakeRegisterAPIURL/registers/gateway": unsupported protocol scheme ""
}

func ExampleFCRRegisterMgr_GetGatewayByID_nil() {
  v, err := rm.GetGatewayByID(nil)
  fmt.Println(v, err)
  // Output:
  // {         <nil>} gateway ID is not provided in GetGatewayByID
}

func ExampleFCRRegisterMgr_GetGatewayByID_new() {
  nd := &nodeid.NodeID{}
  v, err := rm.GetGatewayByID(nd)
  fmt.Println(v, err)
  // Output:
  // {         <nil>} Get "fakeRegisterAPIURL/registers/gateway/00": unsupported protocol scheme ""
}

func ExampleFCRRegisterMgr_GetRegisteredProviders() {
  v, err := rm.GetRegisteredProviders()
  fmt.Println(v, err)
  // Output:
  // [] Get "fakeRegisterAPIURL/registers/provider": unsupported protocol scheme ""
}

func ExampleFCRRegisterMgr_GetProviderByID_nil() {
  v, err := rm.GetProviderByID(nil)
  fmt.Println(v, err)
  // Output:
  // {        <nil>} provider ID is not provided in GetProviderByID

}

func ExampleFCRRegisterMgr_GetProviderByID_new() {
  nd := &nodeid.NodeID{}
  v, err := rm.GetProviderByID(nd)
  fmt.Println(v, err)
  // Output:
  // {        <nil>} Get "fakeRegisterAPIURL/registers/provider/00": unsupported protocol scheme ""

}
