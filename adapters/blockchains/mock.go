package blockchains

import (
	"errors"
	"io"
	"io/ioutil"

	"github.com/ktr0731/cris/usecases/dto"
	"golang.org/x/sync/syncmap"
)

type MockBlockchainAdapter struct {
	store syncmap.Map
}

func (a *MockBlockchainAdapter) FindTxByHash(hash string) (*dto.Transaction, error) {
	v, ok := a.store.Load(hash)
	if !ok {
		return nil, ErrTxNotFound
	}
	tx, ok := v.(*dto.Transaction)
	if !ok {
		return nil, errors.New("type assertion failed")
	}
	return tx, nil
}

// CreateTx is used to testing only
func (a *MockBlockchainAdapter) CreateTx(hash string, content io.Reader) (*dto.Transaction, error) {
	b, err := ioutil.ReadAll(content)
	if err != nil {
		return nil, err
	}
	tx := &dto.Transaction{
		Data: b,
	}
	a.store.Store(hash, tx)
	return tx, nil
}

func NewMockBlockchain() *MockBlockchainAdapter {
	return &MockBlockchainAdapter{}
}
