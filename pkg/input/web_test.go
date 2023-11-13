package input

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

)

type httpClientMock struct {
	mock.Mock
}

func (m *httpClientMock) Get(url string) (resp *http.Response, err error) {
	args := m.Called(url)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*http.Response), args.Error(1)
}

func TestInput_FromWeb(t *testing.T) {
	t.Parallel()

	mockClient := httpClientMock{}
	defaultHTTPClient = &mockClient

	t.Run("Success", func(t *testing.T) {
		url := "http://example.com"
		expectedContent := "test_content"

		mockClient.On("Get", url).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(expectedContent)),
		}, nil)

		actualContent, err := FromWeb(url)
		assert.NoError(t, err)
		assert.Equal(t, []byte(expectedContent), actualContent)
	})

	t.Run("NotFoundError_Fail", func(t *testing.T) {
		url := "http://example2.com"

		mockClient.On("Get", url).Return(&http.Response{
			StatusCode: http.StatusNotFound,
		}, nil)

		_, err := FromWeb(url)
		assert.Error(t, err)
	})

	t.Run("ClientError_Fail", func(t *testing.T) {
		url := "http://example3.com"

		mockClient.On("Get", url).Return(nil, http.ErrHandlerTimeout)

		_, err := FromWeb(url)
		assert.Error(t, err)
	})
}
