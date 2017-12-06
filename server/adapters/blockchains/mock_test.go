package blockchains

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMockBlockchainAdapter_FindTxByHash(t *testing.T) {
	adapter := NewMockBlockchain()
	hash := "C204"
	content := strings.NewReader("Cristina")
	expected, err := adapter.CreateTx(hash, content)
	require.NoError(t, err)
	actual, err := adapter.FindTxByHash(hash)
	require.NoError(t, err)

	assert.Exactly(t, expected, actual)
}
