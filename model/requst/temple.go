package requst

import (
	"ifuzz/cmd"
	"log"
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
	Num           int
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

func GetRequestTemplate(opt cmd.Option) *RequetsPool {
	strs := Strings{}
	strs = append(strs, opt.Basic.Target...)

	payload := cmd.ParsePayload(opt)
	RequestTemplate.payload = cmd.ParseIterator(payload, opt)

	RequestTemplate.Head = opt.Fuzz.Headers
	RequestTemplate.Method = opt.Basic.Method
	RequestTemplate.Body = opt.Fuzz.Body
	RequestTemplate.Num = opt.Basic.FuzzNum

	pool := InitRequestPool(opt.Fuzz.Thread, opt.Basic.AttackType)

	for v := range strs.Iterator() {
		TaskWaitGroup.Add(1)
		RequestTemplate.Url = v
		RequestTemplate = HasFuzz(RequestTemplate)
		if err := pool.Pool.Invoke(RequestTemplate); err != nil {
			log.Fatal(err.Error())
		}
	}
	return pool
}
