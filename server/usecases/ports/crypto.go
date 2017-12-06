package ports

type CryptoPort interface {
	HashDigest(src []byte) string
}
