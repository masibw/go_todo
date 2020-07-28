package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"
	//"net/http/httptest"
)

func TestUsersGet(t *testing.T) {
	res, err := http.Get("http://localhost:8080/api/1.0/users")
	if err != nil {
		t.Fatalf("geterror %v", err)
		return
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("%v", res.StatusCode)
	}
}

func TestUserCreate(t *testing.T) {
	str := `{"email":"example@example.com","password":"password"}`

	req, err := http.NewRequest(
		"POST",
		"http://localhost:8080/api/1.0/users",
		bytes.NewBuffer([]byte(str)),
	)

	if err != nil {
		t.Fatalf("Error Occured %v", err)
	}

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}

	res, _ := client.Do(req)
	if res.StatusCode != 200 {
		t.Fatalf("%v", res.StatusCode)
	}
	defer res.Body.Close()
}
