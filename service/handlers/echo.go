package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/karlmoad/k8s-example/model"
)

func Echo(rw http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var r model.Request
	err := decoder.Decode(&r)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	var response model.Response
	response.Input = r;
	statement := fmt.Sprintf("RSP:%s", r.Statement)
	response.Response = statement

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}

