package image

import "testing"

func TestToBase64(t *testing.T) {
	_, err := ToBase64("https://mika-resource.oss-cn-hangzhou.aliyuncs.com/blog/avatar.jpg")
	if err != nil {
		t.Errorf("Expect error to be nil, got %v", err)
		return
	}
}
