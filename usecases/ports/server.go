package ports

type RequestPort interface {
	Listen() error
}
