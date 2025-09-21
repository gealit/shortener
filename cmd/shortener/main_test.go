package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setUp() {
	UrlsMap = make(map[string]string)

	for k := range UrlsMap {
		delete(UrlsMap, k)
	}
}

func MockRandSeq(n int) string {
	return "test123"
}

func TestMainPage_POST_WithMockedRand(t *testing.T) {
	setUp()

	// Сохраняем оригинальную функцию
	// originalRandSeq := RandSeq
	// // Подменяем на mock
	// RandSeq = func(n int) string {
	// 	return "test12345" // Фиксированное значение для тестов
	// }
	// Восстанавливаем после теста
	// defer func() { RandSeq = originalRandSeq }()

	requestBody := "https://example.com"
	req, err := http.NewRequest("POST", "/", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MainPage)

	handler.ServeHTTP(rr, req)

	// Проверяем что вернулся предсказуемый URL
	expectedURL := "http://localhost:8080/test123"
	if rr.Body.String() != expectedURL {
		t.Errorf("handler returned unexpected URL: got %v want %v",
			rr.Body.String(), expectedURL)
	}

	// Проверяем что URL добавлен в map
	if UrlsMap["/test123"] != requestBody {
		t.Errorf("URL not stored correctly in map")
	}
}
