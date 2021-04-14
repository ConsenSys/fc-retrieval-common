package fcrtcpcomms

import (
  "errors"
  "net"
  "sync"

  log "github.com/ConsenSys/fc-retrieval-common/pkg/logging"
  "github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// Constants for identifying the correct access point
const (
	AccessFromGateway  = 0
	AccessFromProvider = 1
)

// CommunicationChannel holds the connection for sending outgoing TCP requests.
// CommsLock is used to ensure only one thread can access the tcp connection at any time.
// Conn is the net connection for sending outgoing TCP requests.
type CommunicationChannel struct {
	CommsLock sync.RWMutex
	Conn      net.Conn
}

// CommunicationPool holds the node address map and active node connections.
type CommunicationPool struct {
	// AddressMap stores mapping from node id (big int in string repr) to its node info.
	RegisteredNodeMap     *map[string]RegisteredNode
	RegisteredNodeMapLock *sync.RWMutex

	// ActiveNodes store connected active nodes for outgoing request:
	// A map from node id (big int in string repr) to a CommunicationChannel.
	ActiveNodes     map[string](*CommunicationChannel)
	ActiveNodesLock sync.RWMutex
}

// NewCommunicationPool creates a new communication commPool.
func NewCommunicationPool(registeredNodeMap *map[string]RegisteredNode, registeredNodeMapLock *sync.RWMutex) *CommunicationPool {
	return &CommunicationPool{
		RegisteredNodeMap:     registeredNodeMap,
		RegisteredNodeMapLock: registeredNodeMapLock,
		ActiveNodes:           make(map[string](*CommunicationChannel)),
		ActiveNodesLock:       sync.RWMutex{},
	}
}

// GetConnForRequestingNode returns the connection for sending request to a node with given id.
// It will reuse any active connection.
func (commPool *CommunicationPool) GetConnForRequestingNode(nodeID *nodeid.NodeID, accessFrom int) (*CommunicationChannel, error) {
	log.Info("Get active connection, nodeID: %v", nodeID.ToString())
	commPool.ActiveNodesLock.RLock()
	comm := commPool.ActiveNodes[nodeID.ToString()]
	commPool.ActiveNodesLock.RUnlock()
	if comm == nil {
		log.Info("No active connection, connect to peer")
		var address string
		commPool.RegisteredNodeMapLock.RLock()
		node, ok := (*commPool.RegisteredNodeMap)[nodeID.ToString()]
		if ok {
			switch accessFrom {
			case AccessFromGateway:
				address = node.GetNetworkInfoGateway()
			case AccessFromProvider:
				address = node.GetNetworkInfoProvider()
			}
		}
		commPool.RegisteredNodeMapLock.RUnlock()
		if !ok {
			return nil, errors.New("Node not found in register")
		}
		log.Debug("Got address: %v", address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Error("Unable to get connection: %v", err)
			return nil, err
		}
		comm = &CommunicationChannel{
			CommsLock: sync.RWMutex{},
			Conn:      conn}
		err = commPool.RegisterNodeCommunication(nodeID, comm)
		if err != nil {
			log.Error("Unable to register node communication: %v", err)
			conn.Close()
			return nil, err
		}
	}
	return comm, nil
}

// AddRegisteredNode adds a new registered node
func (commPool *CommunicationPool) AddRegisteredNode(nodeID *nodeid.NodeID, node *RegisteredNode) {
	commPool.RegisteredNodeMapLock.Lock()
	defer commPool.RegisteredNodeMapLock.Unlock()
	(*commPool.RegisteredNodeMap)[nodeID.ToString()] = *node
}

// DeregisterNodeAddress deregisters a node address
// Fail silently
func (commPool *CommunicationPool) DeregisterNodeAddress(nodeID *nodeid.NodeID) {
	commPool.RegisteredNodeMapLock.Lock()
	defer commPool.RegisteredNodeMapLock.Unlock()
	_, exist := (*commPool.RegisteredNodeMap)[nodeID.ToString()]
	if exist {
		delete(*commPool.RegisteredNodeMap, nodeID.ToString())
	}
}

// RegisterNodeCommunication registers a node communication
func (commPool *CommunicationPool) RegisterNodeCommunication(nodeID *nodeid.NodeID, comm *CommunicationChannel) error {
	commPool.ActiveNodesLock.Lock()
	defer commPool.ActiveNodesLock.Unlock()
	_, exist := commPool.ActiveNodes[nodeID.ToString()]
	if exist {
		return errors.New("Error: connection existed")
	}
	commPool.ActiveNodes[nodeID.ToString()] = comm
	return nil
}

// DeregisterNodeCommunication deregisters a node communication
// Fail silently
func (commPool *CommunicationPool) DeregisterNodeCommunication(nodeID *nodeid.NodeID) {
	commPool.ActiveNodesLock.Lock()
	defer commPool.ActiveNodesLock.Unlock()
	comm, exist := commPool.ActiveNodes[nodeID.ToString()]
	if exist {
		comm.Conn.Close()
		delete(commPool.ActiveNodes, nodeID.ToString())
	}
}
