package io

import "sea/internal/consts"

type EncryptStage interface {
	SetConnectionRoute(ConnectionRoute)
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
	GetConnectionRoute() ConnectionRoute
	GetMessageType()
	GetMessageData()
}

type ConnectionRoute struct {
	CurrentID  [consts.SERVER_ID_BYTES]byte
	SenderID   [consts.SERVER_ID_BYTES]byte
	ReceiverID [consts.SERVER_ID_BYTES]byte
	RelayID    [][consts.SERVER_ID_BYTES]byte
}

func (cr *ConnectionRoute) Init() *ConnectionRoute {
	cr.CurrentID = *new([consts.SERVER_ID_BYTES]byte)
	cr.SenderID = *new([consts.SERVER_ID_BYTES]byte)
	cr.ReceiverID = *new([consts.SERVER_ID_BYTES]byte)
	cr.RelayID = *new([][consts.SERVER_ID_BYTES]byte)
	return cr
}
func (cr *ConnectionRoute) SetCurrentID(v [consts.SERVER_ID_BYTES]byte) *ConnectionRoute {
	cr.CurrentID = v
	return cr
}
func (cr *ConnectionRoute) SetSenderID(v [consts.SERVER_ID_BYTES]byte) *ConnectionRoute {
	cr.SenderID = v
	return cr
}
func (cr *ConnectionRoute) SetReceiverID(v [consts.SERVER_ID_BYTES]byte) *ConnectionRoute {
	cr.ReceiverID = v
	return cr
}
func (cr *ConnectionRoute) AddRelayID(v [consts.SERVER_ID_BYTES]byte) *ConnectionRoute {
	cr.RelayID = append(cr.RelayID, v)
	return cr
}
func (cr *ConnectionRoute) IsValid() bool {
	check1 := (cr.CurrentID == cr.SenderID) || (cr.CurrentID == cr.ReceiverID)
	if check1 {
		return true
	}
	check2 := false
	for _, v := range cr.RelayID {
		if cr.CurrentID == v {
			check2 = true
		}
	}
	return check2
}

type MessageType struct {
	MajorType consts.IOMessageTypeMajor
	MinorType consts.IOMessageTypeMinor
}
