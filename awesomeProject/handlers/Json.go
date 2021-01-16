package handlers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func toJson(value interface{}) []byte {
	bytes, err := json.Marshal(value)

	if err != nil {
		log.Printf("error while marshal json: %s", err)
		return []byte{}
	}

	return bytes
}

func WriteJsonToResponse(rw http.ResponseWriter, value interface{}) error {
	bytes, err := json.Marshal(value)

	if err != nil {
		return errors.Wrap(err, "error while marshal json")
	}

	rw.WriteHeader(200)
	_, err = rw.Write(bytes)
	if err != nil {
		return errors.Wrap(err, "error write response")
	}
	return nil
}
