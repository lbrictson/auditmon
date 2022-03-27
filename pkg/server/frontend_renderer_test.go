package server

import "testing"

func TestMustNewRenderer(t *testing.T) {
	r := MustNewRenderer()
	if r == nil {
		t.Error("expected a new renderer but got nil")
	}
	return
}
