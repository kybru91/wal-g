package yckms

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializeDeserializeKeyHeader(t *testing.T) {
	buffer := new(bytes.Buffer)

	encryptedKey := []byte("awesomekey")
	serializedKey := serializeEncryptedKey(encryptedKey)
	buffer.Write(serializedKey)

	deserializedKey, err := readEncryptedKey(buffer)
	assert.NoErrorf(t, err, "YcKms envelope key deserialization error: %v", err)

	assert.Equal(t, len(encryptedKey), len(deserializedKey), "YcKms deserialized envelope key len is not equal to the original one")

	for i := range encryptedKey {
		assert.Equal(t, encryptedKey[i], deserializedKey[i], "YcKms deserialized envelope key is not equal to the original one in position: %d", i)
	}
}
