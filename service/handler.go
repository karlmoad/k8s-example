package service

import "net/http"

type HttpGetHandler interface {
	Get(rw http.ResponseWriter, req *http.Request)
}

type HttpPutHandler interface {
	Put(rw http.ResponseWriter, req *http.Request)
}