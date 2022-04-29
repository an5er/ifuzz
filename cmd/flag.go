package cmd

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type Option struct {
	Basic BasicOption      `description:"basic type"`
	Fuzz  FuzzMondelOption `description:"Fuzz type"`
}

type BasicOption struct {
	Target     []string `short:"t" description:"your target"`
	Output     string   `short:"o" description:"Output method:false==cmd,ture==file" default:"fales"`
	Method     string   `short:"m" description:"Request method" default:"GET"`
	FuzzNum    int      `short:"n" description:"Fuzz Num"`
	AttackType string   `short:"a" description:"AttackType" default:"Sniper"`
}
type FuzzMondelOption struct {
	Iterator string   `short:"i" description:"Iterator mondel" default:"chain"`
	Rspcode  string   `short:"c" description:"Set the Respone code"`
	Payloads []string `short:"p" description:"Set the payloads"`
	Headers  string   `long:"header" description:"Set the Headers" default:"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36"`
	Body     string   `long:"body" description:"Set the Body"`
	Thread   int      `long:"thread" description:"The number of concurrent threads"`
}

func FlagParse() Option {
	var opt Option
	p := flags.NewParser(&opt, flags.Default)
	_, _ = p.ParseArgs(os.Args[1:])
	return opt
}
