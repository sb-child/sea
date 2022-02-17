package consts

const (
	// water key
	WATER_KEY_STATUS_OK = iota
	WATER_KEY_STATUS_WAIT_FOR_RESULT
	WATER_KEY_STATUS_BANNED
	WATER_KEY_STATUS_NOT_FOUND
	WATER_KEY_CHECK_OK
	WATER_KEY_CHECK_TEST_FAILED
	WATER_KEY_CHECK_WRONG_SIZE
	WATER_KEY_CHECK_TYPE_ERROR
	WATER_KEY_CHECK_EXPIRED
	// water join
	JOIN_RETURN_CODE_SUCCESS            // success
	JOIN_RETURN_CODE_DECRYPTION_FAILED  // failed to decrypt
	JOIN_RETURN_CODE_SESSION_NOT_FOUND  // session not found
	JOIN_RETURN_CODE_SESSION_ERROR      // can't create session, needs a retry
	JOIN_RETURN_CODE_BAD_KEY            // invalid key, expired, a private key, banned key or empty string
	JOIN_RETURN_CODE_BAD_RANDOM_STRING  // random string is not 32 characters long
	JOIN_RETURN_CODE_KEY_ALREADY_EXISTS // this key already exists
	JOIN_RETURN_CODE_SERVER_ERROR       // server isn't ready
	JOIN_RETURN_CODE_BANNED             // key is banned
	// water stream
	STREAM_TYPE_MESSAGE // sync messages
	STREAM_TYPE_FILE    // sync files
	STREAM_TYPE_SOCKET  // communicates in socket
)
