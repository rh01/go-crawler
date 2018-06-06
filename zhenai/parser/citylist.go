package parser

import (
	"justforunc/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	regx := regexp.MustCompile(cityListRe)

	allMatchStrings := regx.FindAllStringSubmatch(string(contents), -1)

	result := engine.ParseResult{}

	for _, value := range allMatchStrings {
		result.Items = append(result.Items, string(value[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        value[1],
			ParserFunc: engine.NilParser,
		})
	}

	return result

}
