package base64

import (
	"errors"
	"strings"
)

func Ext(base64Str string) (ext string, err error) {
	if strings.HasPrefix(base64Str, "data:image/") {
		ext = strings.Split(base64Str, ";")[0][11:]
	} else {
		err = errors.New("not an image")
	}
	return
}
