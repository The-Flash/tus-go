package uid

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

func Uid() string {
	id := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, id)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(id)
}
