package presenters

import (
	"bytes"
	"testing"

	"github.com/ktr0731/cris/config"
	"github.com/ktr0731/cris/domain/entities"
	"github.com/ktr0731/cris/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPPresenter_UploadFile(t *testing.T) {
	presenter := NewHTTPPresenter(&log.Logger{}, &config.Config{})

	expectedToken := entities.FileID("token")
	res, err := presenter.UploadFile(expectedToken)
	require.NoError(t, err)
	assert.Equal(t, res.Token, string(expectedToken))
}

func TestHTTPPresenter_DownloadFile(t *testing.T) {
	presenter := NewHTTPPresenter(&log.Logger{}, &config.Config{})

	expected := []byte("foo")
	r := bytes.NewReader(expected)
	res, err := presenter.DownloadFile(r)
	require.NoError(t, err)
	assert.Equal(t, res.Content, expected)
}
