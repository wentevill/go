package handler

import (
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	Foo string
}

type Response struct {
	Bar string
}

// Handler function as a server handler
// @tags handler
// @Summary handler summary
// @Description handler description
// @Produce json
// @Param request body string true "request"
// @Success 200 {object} Response
// @Router / [post]
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld: received a request")
	// your logic
	fmt.Fprintf(w, "Hello! this is faas\n")
}
