package fcrmessages

import (
	"encoding/json"
	"fmt"
)

// ProtocolChangeResponse message is sent to indicate that the gateway is requesting the other entity change protocol version
type ProtocolChangeResponse struct {
	DesiredVersion int32 `json:"desired_version"`
}

// EncodeProtocolChangeResponse is used to get the FCRMessage of ProtocolChangeResponse
func EncodeProtocolChangeResponse(desiredVersion int32) (*FCRMessage, error) {
	body, err := json.Marshal(ProtocolChangeResponse{
		DesiredVersion: desiredVersion,
	})
	if err != nil {
		return nil, err
	}
	return &FCRMessage{
		MessageType:       ProtocolChangeResponseType,
		ProtocolVersion:   protocolVersion,
		ProtocolSupported: protocolSupported,
		MessageBody:       body,
	}, nil
}

// DecodeProtocolChangeResponse is used to get the fields from FCRMessage of ProtocolChangeResponse
func DecodeProtocolChangeResponse(fcrMsg *FCRMessage) (
	int32, // desired version
	error, // error
) {
	if fcrMsg.MessageType != ProtocolChangeResponseType {
		return 0, fmt.Errorf("Message type mismatch")
	}
	msg := ProtocolChangeResponse{}
	err := json.Unmarshal(fcrMsg.MessageBody, &msg)
	if err != nil {
		return 0, err
	}
	return msg.DesiredVersion, nil
}

// ProtocolMismatchResponse message is sent to indicate that there are no common protocol versions between the gateway and the requesting entity.
type ProtocolMismatchResponse struct {
}

// EncodeProtocolMismatchResponse is used to get the FCRMessage of ProtocolMismatchResponse
func EncodeProtocolMismatchResponse() (*FCRMessage, error) {
	body, err := json.Marshal(ProtocolMismatchResponse{})
	if err != nil {
		return nil, err
	}
	return &FCRMessage{
		MessageType:       ProtocolMismatchResponseType,
		ProtocolVersion:   protocolVersion,
		ProtocolSupported: protocolSupported,
		MessageBody:       body,
	}, nil
}

// DecodeProtocolMismatchResponse is used to get the fields from FCRMessage of ProtocolMismatchResponse
func DecodeProtocolMismatchResponse(fcrMsg *FCRMessage) error {
	if fcrMsg.MessageType != ProtocolMismatchResponseType {
		return fmt.Errorf("Message type mismatch")
	}
	msg := ProtocolMismatchResponse{}
	return json.Unmarshal(fcrMsg.MessageBody, &msg)
}

// InvalidMessageResponse message is sent to indicate that the message is invalid
type InvalidMessageResponse struct {
}

// EncodeInvalidMessageResponse is used to get the FCRMessage of InvalidMessageResponse
func EncodeInvalidMessageResponse() (*FCRMessage, error) {
	body, err := json.Marshal(InvalidMessageResponse{})
	if err != nil {
		return nil, err
	}
	return &FCRMessage{
		MessageType:       InvalidMessageResponseType,
		ProtocolVersion:   protocolVersion,
		ProtocolSupported: protocolSupported,
		MessageBody:       body,
	}, nil
}

// DecodeInvalidMessageResponse is used to get the fields from FCRMessage of InvalidMessageResponse
func DecodeInvalidMessageResponse(fcrMsg *FCRMessage) error {
	if fcrMsg.MessageType != InvalidMessageResponseType {
		return fmt.Errorf("Message type mismatch")
	}
	msg := InvalidMessageResponse{}
	return json.Unmarshal(fcrMsg.MessageBody, &msg)
}

// InsufficientFundsResponse message is sent to indicate the there are not sufficient funds to finish the request
type InsufficientFundsResponse struct {
	PaymentChannelID int64 `json:"payment_channel_id"`
}

// EncodeInsufficientFundsResponse is used to get the FCRMessage of InsufficientFundsResponse
func EncodeInsufficientFundsResponse(paymentChannelID int64) (*FCRMessage, error) {
	body, err := json.Marshal(InsufficientFundsResponse{
		PaymentChannelID: paymentChannelID,
	})
	if err != nil {
		return nil, err
	}
	return &FCRMessage{
		MessageType:       InsufficientFundsResponseType,
		ProtocolVersion:   protocolVersion,
		ProtocolSupported: protocolSupported,
		MessageBody:       body,
	}, nil
}

// DecodeInsufficientFundsResponse is used to get the fields from FCRMessage of InsufficientFundsResponse
func DecodeInsufficientFundsResponse(fcrMsg *FCRMessage) (
	int64, // payment channel id
	error, // error
) {
	if fcrMsg.MessageType != InsufficientFundsResponseType {
		return 0, fmt.Errorf("Message type mismatch")
	}
	msg := InsufficientFundsResponse{}
	err := json.Unmarshal(fcrMsg.MessageBody, &msg)
	if err != nil {
		return 0, err
	}
	return msg.PaymentChannelID, nil
}
