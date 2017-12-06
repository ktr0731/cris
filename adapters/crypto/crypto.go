package crypto

import "crypto/sha256"

type CryptoAdapter struct{}

func (c *CryptoAdapter) HashDigest(src []byte) string {
	sum := sha256.Sum256(src)
	return string(sum[:])
}

func NewCryptoAdapter() *CryptoAdapter {
	return &CryptoAdapter{}
}
