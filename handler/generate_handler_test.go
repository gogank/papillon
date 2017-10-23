package handler

import "testing"

func TestGenerate(t *testing.T) {
	err := Generate("./config.toml")
	if err != nil {
		t.Errorf("generate error: %v", err)
	}
}
