package service

var WaterKey = waterKeyService{}

type waterKeyService struct{}

const (
	WATER_KEY_STATUS_OK              = 0
	WATER_KEY_STATUS_WAIT_FOR_RESULT = 1
	WATER_KEY_STATUS_BANNED          = 2
	WATER_KEY_STATUS_NOT_FOUND       = 3
)

// GetSelfKeyID returns the self key ID
func (s *waterKeyService) GetSelfKeyID() (string, error) {
	return "", nil
}

// GetSelfKeyID returns all the keys ID stored
func (s *waterKeyService) GetKeyIDList() []string {
	return make([]string, 0)
}

func (s *waterKeyService) GetKey(id string) (string, error) {
	return "", nil
}

func (s *waterKeyService) GetKeyStatus(id string) int {
	return WATER_KEY_STATUS_OK
}

func (s *waterKeyService) SetKeyStatus(id string) error {
	return nil
}

func (s *waterKeyService) DeleteKey(id string) error {
	return nil
}
