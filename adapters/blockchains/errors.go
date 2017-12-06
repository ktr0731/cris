package blockchains

import "errors"

type errTxNotFound struct {
	error
}

func (e errTxNotFound) ClientError() bool {
	return true
}

var (
	ErrTxNotFound = errTxNotFound{errors.New("no such transaction")}
)
