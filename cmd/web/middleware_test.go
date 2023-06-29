package main

import (
	"fmt"
	"net/http"
	"testing"
)

// TestNoSurf tests if the output of NoSurf is of type http.Handler
func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
			// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}

// TestSessionLoad tests if the output of SessionLoad is of type http.Handler
func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
			// do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}