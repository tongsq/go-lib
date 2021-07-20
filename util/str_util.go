package util

import (
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func HasPrefix(str string, prefixList []string) bool {
	for _, p := range prefixList {
		if strings.HasPrefix(str, p) {
			return true
		}
	}
	return false
}

func Contains(str string, prefixList []string) bool {
	for _, p := range prefixList {
		if strings.Contains(str, p) {
			return true
		}
	}
	return false
}

func GbkToUtf8(s string) (string, error) {
	reader := transform.NewReader(strings.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d[:]), nil
}

func Utf8ToGbk(s string) (string, error) {
	reader := transform.NewReader(strings.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return "", e
	}
	return string(d[:]), nil
}
