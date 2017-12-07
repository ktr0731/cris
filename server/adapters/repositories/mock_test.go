package repositories

import (
	"testing"

	"github.com/ktr0731/cris/server/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMockFileRepositoryAdapter_Store(t *testing.T) {
	repository := NewMockFileRepository(nil, nil)
	expected := entities.NewFile("url")
	_, err := repository.Store(expected)
	require.NoError(t, err)
	v, ok := repository.storage.Load(expected.ID)
	require.True(t, ok)
	actual, ok := v.(*entities.File)
	require.True(t, ok)
	assert.Equal(t, expected, actual)
}

func TestMockFileRepositoryAdapter_Find(t *testing.T) {
	repository := NewMockFileRepository(nil, nil)
	expected := entities.NewFile("url")
	_, err := repository.Store(expected)
	require.NoError(t, err)
	actual, err := repository.Find(expected.ID)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestMockFileRepositoryAdapter_Remove(t *testing.T) {
	repository := NewMockFileRepository(nil, nil)
	expected := entities.NewFile("url")
	_, err := repository.Store(expected)
	require.NoError(t, err)
	_, err = repository.Remove(expected.ID)
	require.NoError(t, err)
}
