package entities

import "net/http"

type Route struct {
	method          string
	path            string
	handlerFunction http.Handle
}
