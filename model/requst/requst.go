package requst

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
)

func Requests(template Template) {
	request := resty.New().R()
	request.Method = template.Method
	request.URL = template.RequestUrl
	request.SetBody(template.RequestBody)
	request.SetHeader("User-Agent", template.Head)
	resp, err := request.Send()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status Code:", resp.StatusCode(), "url", request.URL, "body", template.RequestBody)
}
