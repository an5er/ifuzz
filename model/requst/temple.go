package requst

import (
	"ifuzz/cmd"
)

type Template struct {
	Url           string
	RequestUrl    string
	hasFuzzUrl    bool
	RequestBody   string
	Body          string
	hasFuzzBody   bool
	Head          string
	Method        string
	RequestMethod string
	hasFuzzMethod bool
	payload       cmd.IteratorDic
}

var RequestTemplate Template

type Strings []string

func (strs Strings) Iterator() <-chan string {
	c := make(chan string)
	go func() {
		for _, v := range strs {
			c <- v
		}
		close(c)
	}()
	return c
}

func GetRequestTemplate(opt cmd.Option) {
	strs := Strings{}
	strs = append(strs, opt.Basic.Target...)
	payload := cmd.ParsePayload(opt)
	RequestTemplate.payload = cmd.ParseIterator(payload, opt)
	RequestTemplate.Head = opt.Fuzz.Headers
	RequestTemplate.Method = opt.Basic.Method
	RequestTemplate.Body = opt.Fuzz.Body

	for v := range strs.Iterator() {
		RequestTemplate.Url = v
		RequestTemplate = HasFuzz(RequestTemplate)
		if opt.Basic.AttackType == "Sniper" || opt.Basic.AttackType == "Cluster bomb" {
			Newtemplate(RequestTemplate, opt.Basic.FuzzNum)
		}
		if opt.Basic.AttackType == "Pitchfork" {
			NewZiptemplate(RequestTemplate, opt.Basic.FuzzNum)
		}
	}

}
