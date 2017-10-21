package handler

import "testing"

func TestGenerate(t *testing.T) {
	err := Generate()
	if err != nil {
		t.Errorf("generate error: %v", err)
	}
}