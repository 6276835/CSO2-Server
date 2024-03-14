package encode

import (
	"bytes"
	"io"

	"github.com/djimenez/iconv-go"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// GbkToUtf8 转换GBK编码到UTF-8编码
func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = io.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// Utf8ToGbk 转换UTF-8编码到GBK编码
func Utf8ToGbk(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewEncoder())
	b, err = io.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// Utf8ToLocal 转换UTF-8编码到本地编码
func Utf8ToLocal(str string) (b string, err error) {
	buf, err := iconv.ConvertString(str, "UTF-8", "GBK")
	return string(buf), err
}

// LocalToUtf8 转换本地编码到UTF-8编码
func LocalToUtf8(str string) (b string, err error) {
	buf, err := iconv.ConvertString(str, "GBK", "UTF-8")
	return string(buf), err
}
