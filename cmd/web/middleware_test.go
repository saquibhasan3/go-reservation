package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// Will not do anything
	default:
		t.Error(fmt.Sprintf("Expecting http handler but got %T ", v))

	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// Will not do anything
	default:
		t.Error(fmt.Sprintf("Expecting http handler but got %T ", v))

	}
}
