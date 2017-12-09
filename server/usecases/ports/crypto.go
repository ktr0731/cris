package ports

type CryptoPort interface {
	HashDigest(src []byte) string
	Verify(pubkey, msg, signature []byte) bool
}
