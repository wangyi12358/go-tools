package image

import (
	"image"
	"net/http"
)

func Size(src string) (width int, height int, err error) {
	response, err := http.Get(src)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// 通过image.Decode获取图片对象
	img, _, err := image.Decode(response.Body)
	if err != nil {
		return
	}

	// 获取图片宽高信息
	bounds := img.Bounds()
	width, height = bounds.Dx(), bounds.Dy()

	return
}
