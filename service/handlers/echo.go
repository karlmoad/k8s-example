package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/karlmoad/k8s-example/model"
	"os"
)

func Echo(rw http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var r model.Request
	err := decoder.Decode(&r)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	secret := os.Getenv("SECRET")

	if secret == "" {
		secret = "NO MORE SECRETS:"
	}

	var response model.Response
	response.Input = r;
	response.Response = fmt.Sprintf("%s:%s", secret, r.Statement)

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}

