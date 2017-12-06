package usecases

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/ktr0731/cris/adapters/blockchains"
	"github.com/ktr0731/cris/adapters/repositories"
	"github.com/ktr0731/cris/domain/entities"
	"github.com/ktr0731/cris/usecases/ports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type dummyOutputAdapter struct {
	ports.ServerOutputPort
	t *testing.T
}

func (a *dummyOutputAdapter) UploadFile(token entities.FileID) (*ports.UploadFileResponse, error) {
	return &ports.UploadFileResponse{Token: string(token)}, nil
}

func (a *dummyOutputAdapter) DownloadFile(content io.Reader) (*ports.DownloadFileResponse, error) {
	b, err := ioutil.ReadAll(content)
	require.NoError(a.t, err)
	return &ports.DownloadFileResponse{Content: b}, nil
}

type dummyStorage struct{}

func (s *dummyStorage) Upload(name string, content io.Reader) (string, error) {
	return "echelon", nil
}

func (s *dummyStorage) Download(url string) (io.ReadCloser, error) {
	return ioutil.NopCloser(strings.NewReader("kurisu")), nil
}

type dummyCryptoAdapter struct{}

func (a *dummyCryptoAdapter) HashDigest(src []byte) string {
	return string(src)
}

func Test_container_uploadFile(t *testing.T) {
	c := newUsecaseContainer(nil, nil)
	output := &dummyOutputAdapter{t: t}
	storage := &dummyStorage{}
	repo := repositories.NewMockFileRepository(nil, nil)

	t.Run("normal", func(t *testing.T) {
		input := &ports.UploadFileParams{}
		res, err := c.uploadFile(input, output, storage, repo)
		require.NoError(t, err)
		expected, err := repo.Find(entities.FileID(res.Token))
		require.NoError(t, err)
		assert.Equal(t, expected.URL, "echelon")
	})
}

func Test_container_downloadFile(t *testing.T) {
	c := newUsecaseContainer(nil, nil)
	output := &dummyOutputAdapter{t: t}
	storage := &dummyStorage{}
	bc := blockchains.NewMockBlockchain()
	repo := repositories.NewMockFileRepository(nil, nil)

	t.Run("no such transaction", func(t *testing.T) {
		input := &ports.DownloadFileParams{}
		_, err := c.downloadFile(input, nil, nil, bc, nil, nil)
		assert.Equal(t, blockchains.ErrTxNotFound, err)
	})

	t.Run("normal", func(t *testing.T) {
		input := &ports.DownloadFileParams{Token: "hououin", TxHash: "makise"}
		crypto := &dummyCryptoAdapter{}

		// setup
		_, err := repo.Store(&entities.File{ID: entities.FileID("hououin")})
		require.NoError(t, err)
		_, err = bc.CreateTx("makise", strings.NewReader("kurisu"))
		require.NoError(t, err)

		// check
		_, err = c.downloadFile(input, output, storage, bc, crypto, repo)
		require.NoError(t, err)
	})
}
