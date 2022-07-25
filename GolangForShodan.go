package main

import (
	"demo/shodan/shodan"
	"flag"
)

var model string
var ip string
var domain string
var token string

func CMDparams() {
	flag.StringVar(&model, "m", "", "查询模式,根据ip查询信息(-m a -i 1.1.1.1)和扫描子域名(-m b -d google.com)")
	flag.StringVar(&ip, "i", "", "ip地址")
	flag.StringVar(&domain, "d", "", "域名")
	flag.StringVar(&token, "t", "", "您的shodam token")
	flag.Parse()
	if model == "a" {
		var t shodan.QueryInfo
		t.OutputQueryByIP(t.QueryByIP(token, ip))
	} else if model == "b" {
		var t shodan.QueryInfo
		temp := t.QuerySubdomainsByDomain(token, domain)
		t.OutputSubdomains(temp)
	}
}

func main() {
	CMDparams()
}
