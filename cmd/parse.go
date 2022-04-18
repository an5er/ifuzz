package cmd

import (
	"ifuzz/model/iterable"
	"ifuzz/model/payloads"
	"ifuzz/utils"
	"strings"
)

type IteratorDic struct {
	ChainIteratorDic   []string
	ProductIteratorDic [][]string
	ZipIteratorDic     [][]string
}

var FilePreix = "./asset/wordlist/"

func ParseUrl(opt Option) Option {
	for i := 0; i < len(opt.Basic.Target); i++ {
		if utils.FileIsExist(opt.Basic.Target[i]) {
			opt.Basic.Target = utils.SliceAdd(opt.Basic.Target, utils.Read(opt.Basic.Target[i]))
			opt.Basic.Target = utils.SliceDelete(opt.Basic.Target, i)
			i--
		}
	}
	return opt
}

func ParsePayload(opt Option) [][]string {
	Resultpayload := make([][]string, 0)
	for i := 0; i < len(opt.Fuzz.Payloads); i++ {
		payload := strings.Split(opt.Fuzz.Payloads[i], ",")
		switch payload[0] {
		case "list":
			Resultpayload = append(Resultpayload, payloads.LoadList(payload[1]))
		case "file":
			filename := FilePreix + payload[1]
			Resultpayload = append(Resultpayload, payloads.LoadFile(filename))
		case "range":
			Resultpayload = append(Resultpayload, payloads.LoadRange(payload[1]))
		case "stdin":
			Resultpayload = append(Resultpayload, payloads.LoadStdin())
		}
	}
	return Resultpayload
}

func ParseCode(opt Option) []string {
	codes := make([]string, 0)
	codes = strings.Split(opt.Fuzz.Rspcode, ",")
	return codes
}

func ParseIterator(payload [][]string, opt Option) IteratorDic {
	var iteratordic IteratorDic
	switch opt.Fuzz.Iterator {
	case "chain":
		iteratordic.ChainIteratorDic = iterable.Chain(payload)
	case "product":
		iteratordic.ProductIteratorDic = iterable.Product(payload)
	case "zip":
		iteratordic.ZipIteratorDic = iterable.Zip(payload)
	}
	return iteratordic
}
