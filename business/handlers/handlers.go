package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func GetTestHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	d := struct {
		Data string `json:"data"`
	}{
		Data: "test-data",
	}

	js, err := json.Marshal(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func PostTestHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	d := struct {
		Data string `json:"data"`
	}{
		Data: "post-test-data",
	}

	js, err := json.Marshal(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func PutTestHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	d := struct {
		Data string `json:"data"`
	}{
		Data: "put-test-data",
	}

	js, err := json.Marshal(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func DeleteTestHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	d := struct {
		Data string `json:"data"`
	}{
		Data: "delete-test-data",
	}

	js, err := json.Marshal(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
