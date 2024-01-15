package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// определяем структуру теста
type want struct {
	code     int
	response string
}

type body struct {
	link string
}

type testStruct struct {
	name   string
	want   want
	url    string
	method string
	body   body
}

var dataTestGet = []testStruct{
	{
		name: "positive test get #1",
		want: want{
			code:     307,
			response: ``,
		},
		url:    "/45ujzdf2",
		method: http.MethodGet,
		body:   body{},
	},
	{
		name: "negative test get #1",
		want: want{
			code:     400,
			response: `Код не указан`,
		},
		url:    "/",
		method: http.MethodGet,
		body:   body{},
	},
	{
		name: "negative test get #2",
		want: want{
			code:     404,
			response: `Код не найден`,
		},
		url:    "/awdawdsdd",
		method: http.MethodGet,
		body:   body{},
	},
}

var dataTestPost = []testStruct{
	{
		name: "positive test post #1",
		want: want{
			code:     201,
			response: `jdrgjdfg`,
		},
		url:    "/",
		method: http.MethodPost,
		body: body{
			link: "https://www.google.com",
		},
	},
	{
		name: "negative test post #1",
		want: want{
			code:     400,
			response: `URL не указан`,
		},
		url:    "/",
		method: http.MethodPost,
		body: body{
			link: "",
		},
	},
	{
		name: "negative test post #2",
		want: want{
			code:     500,
			response: `Ошибка преобразования`,
		},
		url:    "/",
		method: http.MethodPost,
		body: body{
			link: "daawdawda",
		},
	},
}

func TestGetHandler(t *testing.T) {
	testHandler(t, dataTestGet)
}

func TestPostHandler(t *testing.T) {
	testHandler(t, dataTestPost)
}

func testHandler(t *testing.T, data []testStruct) {
	for _, tt := range data {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			data := url.Values{}
			if tt.method == http.MethodPost {
				data.Set("link", tt.body.link)
			}

			r := NewRouter()
			ts := httptest.NewServer(r)
			defer ts.Close()

			statusCode, body := testRequest(t, ts, tt.method, tt.url, strings.NewReader(data.Encode()))
			assert.Equal(t, tt.want.code, statusCode)
			assert.Equal(t, tt.want.response, body)
		})
	}
}

func testRequest(t *testing.T, ts *httptest.Server, method string, path string, body io.Reader) (int, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	defer resp.Body.Close()

	return resp.StatusCode, string(respBody)
}
