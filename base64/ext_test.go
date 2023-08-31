package base64

import "testing"

func TestExt(t *testing.T) {
	ext, _ := Ext("data:image/png")
	if ext != "png" {
		t.Errorf("Expected png, got %v", ext)
	}
}
