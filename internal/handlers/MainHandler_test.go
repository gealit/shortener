package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"shortener/internal/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_mainPage(t *testing.T) {
	type want struct {
		code        int
		contentType string
		body        string
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code:        200,
				contentType: "text/plain",
				body:        "http://localhost:8080",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utils.UrlsMap = make(map[string]string)
			requestBody := strings.NewReader("https://example.ru/")
			request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			mainPage(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, res.StatusCode, 201)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}
			parsed, err := url.Parse(string(resBody))
			if err != nil {
				panic(err)
			}

			base := fmt.Sprintf("%s://%s", parsed.Scheme, parsed.Host)
			path := parsed.Path

			assert.Equal(t, base, tt.want.body)
			require.NoError(t, err)
			assert.Equal(t, res.Header.Get("Content-Type"), tt.want.contentType)

			// проверяем Ответ на запрос Get
			newRequest := httptest.NewRequest(http.MethodGet, path, nil)
			newW := httptest.NewRecorder()
			mainPage(newW, newRequest)
			newRes := newW.Result()

			for key, values := range newRes.Header {
				fmt.Printf("  %s: %v\n", key, values)
			}

			assert.Equal(t, newRes.StatusCode, 307)
			assert.Equal(t, newRes.Header.Get("Content-Type"), tt.want.contentType)
			assert.Equal(t, newRes.Header.Get("Location"), "https://example.ru/")
		})
	}
}
