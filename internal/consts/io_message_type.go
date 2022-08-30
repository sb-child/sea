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
	IO_MESSAGE_TYPE_MINOR_PING_REQUEST IOMessageTypeMinor = iota
	IO_MESSAGE_TYPE_MINOR_PING_RESPONSE
)

// IOMessageTypeMinor -> entity operation -> auth
const (
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_AUTH_GET_NEW_SESSION IOMessageTypeMinor = iota
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_AUTH_GET_SESSION_INFO
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_AUTH_LOGIN_WITH_USERNAME_PASSWORD
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_AUTH_LOGIN_WITH_ACCESS_TOKEN
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_AUTH_LOGIN_REQUIRE_2FA
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_AUTH_LOGIN_WITH_2FA
)

// IOMessageTypeMinor -> entity operation -> storage
const (
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STORAGE_GET_INFO IOMessageTypeMinor = iota
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STORAGE_GET_BYTES
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STORAGE_CREATE
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STORAGE_DELETE
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STORAGE_UPLOAD
)

// IOMessageTypeMinor -> entity operation -> messaging
const (
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_MESSAGING_SEND IOMessageTypeMinor = iota
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_MESSAGING_GET
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_MESSAGING_EDIT
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_MESSAGING_DELETE
)

// IOMessageTypeMinor -> entity operation -> stream
const (
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STREAM_GET_NEW_SESSION IOMessageTypeMinor = iota
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STREAM_GET_SESSION_INFO
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STREAM_KEEPALIVE
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STREAM_CONNECT
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STREAM_SEND
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STREAM_RECV
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STREAM_CLOSE
)

// IOMessageTypeMinor -> base services -> direct message
const (
	IO_MESSAGE_TYPE_MINOR_BASE_SERVICES_DIRECT_MESSAGE_GET_HISTORY IOMessageTypeMinor = iota
	IO_MESSAGE_TYPE_MINOR_BASE_SERVICES_DIRECT_MESSAGE_GET_INFO
	IO_MESSAGE_TYPE_MINOR_BASE_SERVICES_DIRECT_MESSAGE_DELETE_HISTORY
	IO_MESSAGE_TYPE_MINOR_BASE_SERVICES_DIRECT_MESSAGE_CONVERT_TO_GROUP
)

// IOMessageTypeMinor -> base services -> group
const (
	IO_MESSAGE_TYPE_MINOR_BASE_SERVICES_GROUP_GET_HISTORY IOMessageTypeMinor = iota
	IO_MESSAGE_TYPE_MINOR_BASE_SERVICES_GROUP_GET_INFO
	IO_MESSAGE_TYPE_MINOR_BASE_SERVICES_GROUP_DELETE_HISTORY
	IO_MESSAGE_TYPE_MINOR_BASE_SERVICES_GROUP_ADD_MEMBER
	IO_MESSAGE_TYPE_MINOR_BASE_SERVICES_GROUP_DELETE_MEMBER
	IO_MESSAGE_TYPE_MINOR_BASE_SERVICES_GROUP_SET_MEMBER_PERMISSION
)
