package main

import (
	"fmt"
	"ifuzz/asset"
	"ifuzz/cmd"
	"ifuzz/model/requst"
)

func main() {
	fmt.Println(asset.FuckingCoding)

	opt := cmd.FlagParse()
	opt = cmd.ParseUrl(opt)
	requst.GetRequestTemplate(opt)

}
