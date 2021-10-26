package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

var actual = "test"
var enc_data = "lwkl"
var shift = "18"

func decrRouter() (*http.Request, *gin.Engine) {
	path := "/decrypt"

	//	Init router
	r := gin.Default()
	r.POST(
		path,
		EncryptHandler,
	)

	req, _ := http.NewRequest(
		"POST",
		path,
		nil,
	)
	return req, r
}

//	Setup encrypt router
func encrypRouter() (*http.Request, *gin.Engine) {
	path := "/encrypt"

	//	Init router
	r := gin.Default()
	r.POST(
		path,
		EncryptHandler,
	)

	req, _ := http.NewRequest(
		"POST",
		path,
		nil,
	)
	return req, r
}

func TestDecrypt(t *testing.T) {
	got := strings.Map(func(r rune) rune {
		return caesar(r, -18)
	}, enc_data)

	if got != actual {
		t.Error("Expected: test\nGot:", got)
	}
}

func TestEncrypt(t *testing.T) {
	got := strings.Map(func(r rune) rune {
		return caesar(r, 18)
	}, actual)

	if got != enc_data {
		t.Error("Expected: lwkl\nGot:", got)
	}
}

func TestDecryptHandlerRequestValid(t *testing.T) {
	req, r := decrRouter()

	//	Add params
	url_queries := req.URL.Query()
	url_queries.Add("data", enc_data)
	url_queries.Add("shift", "-"+shift)
	req.URL.RawQuery = url_queries.Encode()

	h, _ := strconv.ParseInt(shift, 10, 64)
	println(int(h))
	//	Init recorder
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Error("Expected: 200\nGot:", w.Code)
	}

	got := w.Body.String()
	if !strings.Contains(got, actual) {
		t.Errorf("Expected: %s\nGot: %s\n", actual, got)
	}
}

func TestDecryptHandlerDataAbsent(t *testing.T) {
	req, r := decrRouter()

	//	Add params
	url_queries := req.URL.Query()
	url_queries.Add("data", actual)

	//	Init recorder
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error("Expected: 200\nGot:", w.Code)
	}
}

func TestDecryptHandlerShiftAbsent(t *testing.T) {
	req, r := decrRouter()

	//	Add params
	url_queries := req.URL.Query()
	url_queries.Add("shift", shift)

	//	Init recorder
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error("Expected: 200\nGot:", w.Code)
	}
}

//	Expect status code 200
func TestEncryptHandlerRequestValid(t *testing.T) {
	req, r := encrypRouter()

	//	Add params
	url_queries := req.URL.Query()
	url_queries.Add("data", actual)
	url_queries.Add("shift", shift)
	req.URL.RawQuery = url_queries.Encode()

	//	Init recorder
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Error("Expected: 200\nGot:", w.Code)
	}

	got := w.Body.String()
	if !strings.Contains(got, enc_data) {
		t.Errorf("Expected: %s\nGot: %s\n", enc_data, got)
	}

}

func TestEncryptHandlerDataAbsent(t *testing.T) {
	req, r := encrypRouter()

	//	Add params
	url_queries := req.URL.Query()
	url_queries.Add("data", actual)

	//	Init recorder
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error("Expected: 200\nGot:", w.Code)
	}
}

func TestEncryptHandlerShiftAbsent(t *testing.T) {
	req, r := encrypRouter()

	//	Add params
	url_queries := req.URL.Query()
	url_queries.Add("shift", shift)

	//	Init recorder
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error("Expected: 200\nGot:", w.Code)
	}
}
