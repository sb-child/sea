package io

type EncryptStage interface {
	SetConnectionRoute()
	SetMessageType()
	SetMessageData()
	Package()
	Encrypt()
	Signature()
	Take()
}

type DecryptStage interface {
	Put()
	Unsignature()
	Decrypt()
	Unpackage()
	GetConnectionRoute()
	GetMessageType()
	GetMessageData()
}

type ConnectionRoute struct {
	CurrentHash  [256]byte
	SenderHash   [256]byte
	ReceiverHash [256]byte
	RelayHash    [][256]byte
}

func (cr *ConnectionRoute) Init() *ConnectionRoute {
	cr.CurrentHash = *new([256]byte)
	cr.SenderHash = *new([256]byte)
	cr.ReceiverHash = *new([256]byte)
	cr.RelayHash = *new([][256]byte)
	return cr
}
func (cr *ConnectionRoute) SetCurrentHash(v [256]byte) *ConnectionRoute {
	cr.CurrentHash = v
	return cr
}
func (cr *ConnectionRoute) SetSenderHash(v [256]byte) *ConnectionRoute {
	cr.SenderHash = v
	return cr
}
func (cr *ConnectionRoute) SetReceiverHash(v [256]byte) *ConnectionRoute {
	cr.ReceiverHash = v
	return cr
}
func (cr *ConnectionRoute) AddRelayHash(v [256]byte) *ConnectionRoute {
	cr.RelayHash = append(cr.RelayHash, v)
	return cr
}
