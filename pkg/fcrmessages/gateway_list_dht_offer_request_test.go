package fcrmessages

import (
	"testing"

	"github.com/cbergoon/merkletree"
	"github.com/stretchr/testify/assert"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrmerkletree"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// TestEncodeGatewayListDHTOfferRequest success test
func TestEncodeGatewayListDHTOfferRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockCidMin, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCidMax, _ := cid.NewContentIDFromBytes([]byte{2})
	mockBlockHash := "bafy2bzaceabhr6taytcdntpr4poz43dhf3l5jml6z43e5sdv4gdlgsujsg2ze"
	mockTransactionReceipt := "bafy2bzacecz3sy5ar4rg73ri5cq3tndmwn4vnxgzt4bw4cpig7q25af6y5cnc"

	mockCids := []cid.ContentID{*mockCidMin}
	list := make([]merkletree.Content, len(mockCids))
	for i := 0; i < len(mockCids); i++ {
		list[i] = (mockCids)[i]
	}
	merkleTree, _ := fcrmerkletree.CreateMerkleTree(list)
	mockMerkleRoot := merkleTree.GetMerkleRoot()
	mockMerkleProof, _ := merkleTree.GenerateMerkleProof(mockCidMin)

	validMsg := &FCRMessage{
		messageType:       200,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","cid_min":"0000000000000000000000000000000000000000000000000000000000000001","cid_max":"0000000000000000000000000000000000000000000000000000000000000002","block_hash":"bafy2bzaceabhr6taytcdntpr4poz43dhf3l5jml6z43e5sdv4gdlgsujsg2ze","transaction_receipt":"bafy2bzacecz3sy5ar4rg73ri5cq3tndmwn4vnxgzt4bw4cpig7q25af6y5cnc","merkle_root":"c3c3a46684c07d12a9c238787df3049a6f258e7af203e5ddb66a8bd66637e108","merkle_proof":"AAAAMFsiQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBRT0iXQAAAANbMV0="}`),
		signature:         "",
	}

	msg, err := EncodeGatewayListDHTOfferRequest(mockNodeID, mockCidMin, mockCidMax, mockBlockHash, mockTransactionReceipt, mockMerkleRoot, mockMerkleProof)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayListDHTOfferRequest success test
func TestDecodeGatewayListDHTOfferRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockCidMin, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCidMax, _ := cid.NewContentIDFromBytes([]byte{2})
	mockBlockHash := "bafy2bzaceabhr6taytcdntpr4poz43dhf3l5jml6z43e5sdv4gdlgsujsg2ze"
	mockTransactionReceipt := "bafy2bzacecz3sy5ar4rg73ri5cq3tndmwn4vnxgzt4bw4cpig7q25af6y5cnc"

	mockCids := []cid.ContentID{*mockCidMin}
	list := make([]merkletree.Content, len(mockCids))
	for i := 0; i < len(mockCids); i++ {
		list[i] = (mockCids)[i]
	}
	merkleTree, _ := fcrmerkletree.CreateMerkleTree(list)
	mockMerkleRoot := merkleTree.GetMerkleRoot()
	mockMerkleProof, _ := merkleTree.GenerateMerkleProof(mockCidMin)

	validMsg := &FCRMessage{
		messageType:       200,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","cid_min":"0000000000000000000000000000000000000000000000000000000000000001","cid_max":"0000000000000000000000000000000000000000000000000000000000000002","block_hash":"bafy2bzaceabhr6taytcdntpr4poz43dhf3l5jml6z43e5sdv4gdlgsujsg2ze","transaction_receipt":"bafy2bzacecz3sy5ar4rg73ri5cq3tndmwn4vnxgzt4bw4cpig7q25af6y5cnc","merkle_root":"c3c3a46684c07d12a9c238787df3049a6f258e7af203e5ddb66a8bd66637e108","merkle_proof":"AAAAMFsiQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBRT0iXQAAAANbMV0="}`),
		signature:         "",
	}

	nodeID, cidMin, cidMax, blockHash, transactionReceipt, merkleRoot, merkleProof, err := DecodeGatewayListDHTOfferRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockNodeID)
	assert.Equal(t, cidMin, mockCidMin)
	assert.Equal(t, cidMax, mockCidMax)
	assert.Equal(t, blockHash, mockBlockHash)
	assert.Equal(t, transactionReceipt, mockTransactionReceipt)
	assert.Equal(t, merkleRoot, mockMerkleRoot)
	assert.Equal(t, merkleProof, mockMerkleProof)
}
