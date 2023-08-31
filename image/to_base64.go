package image

import (
	"encoding/base64"
	"io"
	"net/http"
	"path"
)

func ToBase64(src string) (base64Str string, err error) {
	response, err := http.Get(src)
	if err != nil {
		return
	}
	defer response.Body.Close()
	imageData, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	encodedImage := base64.StdEncoding.EncodeToString(imageData)
	ext := path.Ext(src)[1:]
	return "data:image/" + ext + ";base64," + encodedImage, nil
}
