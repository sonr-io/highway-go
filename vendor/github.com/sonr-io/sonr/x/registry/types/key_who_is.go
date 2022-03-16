package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WhoIsKeyPrefix is the prefix to retrieve all WhoIs
	WhoIsKeyPrefix = "WhoIs/value/"
)

// WhoIsKey returns the store key to retrieve a WhoIs from the index fields
func WhoIsKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
