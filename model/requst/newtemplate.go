package requst

import (
	"fmt"
	"strings"
)

var FUZZ_KEY_WORD = "FUZ%dZ"

func HasFuzz(template Template) Template {
	if hasFuzzString(template.Url) {
		template.hasFuzzUrl = true
	}
	if hasFuzzString(template.Body) {
		template.hasFuzzBody = true
	}
	if hasFuzzString(template.Method) {
		template.hasFuzzMethod = true
	}
	return template
}

func hasFuzzString(value string) bool {
	return strings.Contains(value, "FUZ")
}

func Newtemplate(i interface{}) {
	defer TaskWaitGroup.Done()
	RequestTemplate := i.(Template)
	for i := 0; i < RequestTemplate.Num; i++ {
		for j := 0; j < len(RequestTemplate.payload.ChainIteratorDic); j++ {

			key := fmt.Sprintf(FUZZ_KEY_WORD, i)
			replace := func(v string) string {
				return strings.ReplaceAll(
					v,
					key,
					RequestTemplate.payload.ChainIteratorDic[j],
				)
			}
			if RequestTemplate.hasFuzzUrl && strings.Contains(RequestTemplate.Url, key) {
				RequestTemplate.RequestUrl = replace(RequestTemplate.Url)
			}

			if RequestTemplate.hasFuzzBody && strings.Contains(RequestTemplate.Body, key) {
				RequestTemplate.RequestBody = replace(RequestTemplate.Body)
			}
			if RequestTemplate.hasFuzzMethod && strings.Contains(RequestTemplate.Method, key) {
				RequestTemplate.Method = replace(RequestTemplate.Method)
			}
			Requests(RequestTemplate)
		}
	}
}

func NewZiptemplate(i interface{}) {
	defer TaskWaitGroup.Done()
	RequestTemplate := i.(Template)
	var key string
	for j := 0; j < len(RequestTemplate.payload.ZipIteratorDic); j++ {
		tempfuzzresult := RequestTemplate.Url
		for i := 0; i < RequestTemplate.Num; i++ {

			key = fmt.Sprintf(FUZZ_KEY_WORD, i)
			replace := func(v string) string {
				return strings.ReplaceAll(
					v,
					key,
					RequestTemplate.payload.ZipIteratorDic[j][i],
				)
			}
			if RequestTemplate.hasFuzzUrl && strings.Contains(RequestTemplate.Url, key) {
				nextkey := fmt.Sprintf(FUZZ_KEY_WORD, i+1)
				if strings.Contains(tempfuzzresult, nextkey) {
					tempfuzzresult = replace(tempfuzzresult)
				} else {
					RequestTemplate.RequestUrl = replace(tempfuzzresult)
				}
			}
			if RequestTemplate.hasFuzzBody && strings.Contains(RequestTemplate.Body, key) {

				nextkey := fmt.Sprintf(FUZZ_KEY_WORD, i+1)
				if strings.Contains(tempfuzzresult, nextkey) {
					tempfuzzresult = replace(tempfuzzresult)
				} else {
					RequestTemplate.RequestBody = replace(tempfuzzresult)
				}

			}
			if RequestTemplate.hasFuzzMethod && strings.Contains(RequestTemplate.Method, key) {
				RequestTemplate.RequestBody = replace(RequestTemplate.Method)

			}
		}
		Requests(RequestTemplate)
	}

}
