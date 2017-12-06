package ports

import "github.com/ktr0731/cris/usecases/dto"

type BlockchainPort interface {
	FindTxByHash(string) (*dto.Transaction, error)
}
