package image

import (
	"encoding/base64"
	"io"
	"net/http"
	"path"
)

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
