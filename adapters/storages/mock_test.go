package storages

import (
	"strings"
	"testing"

	"github.com/ktr0731/cris/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMockStorageAdapter_Upload(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		storage := NewMockStorage()
		key := "kurisu"
		expected := strings.NewReader("kurisu makise")
		_, err := storage.Upload(key, expected)
		require.NoError(t, err)
		v, ok := storage.storage.Load(key)
		require.True(t, ok)
		actual, ok := v.(*strings.Reader)
		require.True(t, ok)
		assert.Equal(t, expected, actual)
	})

	t.Run("empty content", func(t *testing.T) {
		storage := NewMockStorage()
		key := "sern"
		expected := strings.NewReader("")
		_, err := storage.Upload(key, expected)
		assert.Equal(t, usecases.ErrEmptyContent, err)
	})
}

func TestMockStorageAdapter_Download(t *testing.T) {
	storage := NewMockStorage()
	key := "kurisu"
	expected := strings.NewReader("kurisu makise")
	url, err := storage.Upload(key, expected)
	v, err := storage.Download(url)
	require.NoError(t, err)
	actual, ok := v.(*strings.Reader)
	require.True(t, ok)
	assert.Equal(t, expected, actual)
}
