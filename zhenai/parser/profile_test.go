package parser

import (
	"testing"
	"justforunc/fetcher"
)

func TestParseProfile(t *testing.T) {
	contens, err := fetcher.Fetch("http://album.zhenai.com/u/107045901")
	if err != nil {
		panic(err)
	}

	parseResult := ParseProfile(contens)

	// verify parseList

	const resultSize int = 16

	//expectedItems := []string{
	//	"拼了命的开心",
	//	"女", "28",
	//	"0", "0",
	//	"8000", "-",
	//	"20000元", "离异",
	//	"大学本科", "私营企业",
	//	"湖南长沙", "双子座",
	//	"和家人同住", "已购车",
	//}



	if len(parseResult.Items) != 1 {
		t.Errorf("result should have %d"+
			"request; but had %d", resultSize, len(parseResult.Items))
	}



}
