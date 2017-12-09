package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/ed25519"
)

type CryptoAdapter struct{}

func (c *CryptoAdapter) HashDigest(src []byte) string {
	sum := sha256.Sum256(src)
	return string(sum[:])
}

func (c *CryptoAdapter) Verify(pubkey, msg, signature []byte) bool {
	return ed25519.Verify(ed25519.PublicKey(pubkey), msg, signature)
}

func NewCryptoAdapter() *CryptoAdapter {
	return &CryptoAdapter{}
}
