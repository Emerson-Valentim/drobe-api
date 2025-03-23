package uuid

import (
	"github.com/google/uuid"
)

// UUID wraps google/uuid to provide custom JSON marshaling
type UUID struct {
	uuid.UUID
}

// Parse creates a new UUID from string
func Parse(s string) (UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return UUID{}, err
	}
	return UUID{id}, nil
}

// New creates a new random UUID
func New() UUID {
	return UUID{uuid.New()}
}

func (u UUID) String() string {
	return u.UUID.String()
}

func (u UUID) Bytes() [16]byte {
	var bytes [16]byte
	copy(bytes[:], u.UUID[:])
	return bytes
}
