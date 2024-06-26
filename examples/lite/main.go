package main

import (
	"github.com/conneroisu/glisp"
	"github.com/conneroisu/glisp/domain"
)

// newR returns a new rename handler
func newR() glisp.HandlerFunc {
	return func(w glisp.ResponseWriter, r *domain.Request) {
	}
}

// main is the entry point for the server
func main() {
	server := glisp.DefaultMux

	AddRoutes(server)
}

func AddRoutes(server *glisp.ServeMux) {
	server.Handle(domain.MethodTextDocumentRename, newR())
}
