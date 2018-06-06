package main

import (
	"justforunc/zhenai/parser"
	"justforunc/engine"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}	)

}




