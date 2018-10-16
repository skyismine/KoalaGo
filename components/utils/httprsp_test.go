package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpRspJson(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		HttpRspJson(w, 404, "ERROR", "0001", "path no exist", nil)
	}))
	defer ts.Close()
	api := ts.URL
	fmt.Println("url:", api)
	resp, err := http.Get(api)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes[:]))
}