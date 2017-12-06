package crypto

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCryptoAdapter_HashDigest(t *testing.T) {
	content := []byte("future gadget")
	sum := sha256.Sum256(content)
	expected := string(sum[:])

	assert.Equal(t, expected, NewCryptoAdapter().HashDigest(content))
}
