package image

import "testing"

func TestSize(t *testing.T) {
	width, height, _ := Size("https://mika-resource.oss-cn-hangzhou.aliyuncs.com/blog/avatar.png")
	if width != 641 || height != 641 {
		t.Errorf("Expected 641, got %v", width)
		t.Errorf("Expected 641, got %v", height)
	}
}
