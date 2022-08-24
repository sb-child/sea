package consts

type IOMessageTypeMajor uint64
type IOMessageTypeMinor uint64

// IOMessageTypeMajor
const (

	// ping

	IO_MESSAGE_TYPE_MAJOR_PING IOMessageTypeMajor = iota

	// entity operation

	IO_MESSAGE_TYPE_MAJOR_AUTH
	IO_MESSAGE_TYPE_MAJOR_STORAGE
	IO_MESSAGE_TYPE_MAJOR_MESSAGING
	IO_MESSAGE_TYPE_MAJOR_STREAM

	// base services

	IO_MESSAGE_TYPE_MAJOR_DIRECT_MESSAGE
	IO_MESSAGE_TYPE_MAJOR_GROUP
	IO_MESSAGE_TYPE_MAJOR_CHANNEL
	IO_MESSAGE_TYPE_MAJOR_SERVER_NOTIFY

	// base entity

	IO_MESSAGE_TYPE_MAJOR_USER
	IO_MESSAGE_TYPE_MAJOR_BOT

	// admin

	IO_MESSAGE_TYPE_MAJOR_ADMIN

	// server

	IO_MESSAGE_TYPE_MAJOR_SERVER
)

// IOMessageTypeMinor -> ping
const (
	IO_MESSAGE_TYPE_MINOR_PING_DIRECT IOMessageTypeMinor = iota
)

// IOMessageTypeMinor -> entity operation -> auth
const (
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_AUTH_GET_SESSION IOMessageTypeMinor = iota
)
