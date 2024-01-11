package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetHandler(t *testing.T) {
	// определяем структуру теста
	type want struct {
		code     int
		response string
	}

	type body struct {
		link string
	}

	// создаём массив тестов: имя и желаемый результат
	tests := []struct {
		name   string
		want   want
		url    string
		method string
		body   body
	}{
		// определяем все тесты
		{
			name: "positive test #1",
			want: want{
				code:     307,
				response: ``,
			},
			url:    "/45ujzdf2",
			method: http.MethodGet,
			body:   body{},
		},
		{
			name: "positive test #2",
			want: want{
				code:     201,
				response: `jdrgjdfg`,
			},
			url:    "/",
			method: http.MethodPost,
			body: body{
				link: "http:/esefs/api/request",
			},
		},
		{
			name: "negative test #1",
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
			name: "negative test #2",
			want: want{
				code:     400,
				response: `Код не указан`,
			},
			url:    "/",
			method: http.MethodGet,
			body:   body{},
		},
		{
			name: "negative test #3",
			want: want{
				code:     404,
				response: `Код не найден`,
			},
			url:    "/awdawdsdd",
			method: http.MethodGet,
			body:   body{},
		},
		{
			name: "negative test #4",
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
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			data := url.Values{}
			if tt.method == http.MethodPost {
				// используем тип url.Values из пакета net/url
				// устанавливаем данные
				data.Set("link", tt.body.link)
			}

			request := httptest.NewRequest(tt.method, tt.url, strings.NewReader(data.Encode()))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// создаём новый Recorder
			w := httptest.NewRecorder()
			// определяем хендлер
			h := http.HandlerFunc(GetHandler)
			// запускаем сервер
			h.ServeHTTP(w, request)
			res := w.Result()

			// проверяем код ответа
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}

			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(resBody) != tt.want.response {
				t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
			}
		})
	}
}
