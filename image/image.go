package image

import (
	"encoding/base64"
	"fmt"
	"image"
	"io"
	"net/http"
	"path"
	"strings"
)

func GetTypeByBase64(base64Str string) string {
	var contentType string
	if strings.HasPrefix(base64Str, "data:image/") {
		contentType = strings.Split(base64Str, ";")[0][11:]
	} else {
		fmt.Println("not an image")
	}
	return contentType
}

func ImageToBase64(imageUrl string) (base64Str string, err error) {
	response, err := http.Get(imageUrl)
	if err != nil {
		return
	}
	defer response.Body.Close()
	imageData, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	encodedImage := base64.StdEncoding.EncodeToString(imageData)
	ext := path.Ext(imageUrl)[1:]
	return "data:image/" + ext + ";base64," + encodedImage, nil
}

func GetImageSize(src string) (int, int) {
	response, err := http.Get(src)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 通过image.Decode获取图片对象
	img, _, err := image.Decode(response.Body)
	if err != nil {
		panic(err)
	}

	// 获取图片宽高信息
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	return width, height
}
