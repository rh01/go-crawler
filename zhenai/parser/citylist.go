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
	// contro
	limit := 10
	//-----

	for _, value := range allMatchStrings {
		result.Items = append(result.Items, "City " + string(value[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        value[1],
			ParserFunc: ParseUser,
		})
		limit--
		if limit==0 {
			break
		}
	}



	return result

}
