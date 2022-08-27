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
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STORAGE_GET_FILE_INFO IOMessageTypeMinor = iota
	IO_MESSAGE_TYPE_MINOR_ENTITY_OPERATION_STORAGE_UPLOAD_FILE
)
