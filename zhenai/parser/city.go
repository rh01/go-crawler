package parser

import (
	"justforunc/engine"
	"regexp"
)

const UserRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseUser(contents []byte) engine.ParseResult {
	regx := regexp.MustCompile(UserRe)

	allMatchStrings := regx.FindAllStringSubmatch(string(contents), -1)

	result := engine.ParseResult{}

	for _, value := range allMatchStrings {
		name := string(value[2])
		result.Items = append(result.Items, "User " +  string(value[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        value[1],
			// function bibao
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	return result

}
