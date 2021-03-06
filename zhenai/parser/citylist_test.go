package parser

import (
	"testing"
	"justforunc/fetcher"
)

func TestParseCityList(t *testing.T) {
	contens, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	parseResult := ParseCityList(contens)

	// verify parseList

	const resultSize int = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	expectedItems := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	if len(parseResult.Requests) != resultSize {
		t.Errorf("result should have %d"+
			"request; but had %d", resultSize, len(parseResult.Requests))
	}

	if len(parseResult.Items) != resultSize {
		t.Errorf("result should have %d"+
			"request; but had %d", resultSize, len(parseResult.Items))
	}

	for i, url := range expectedUrls {
		if parseResult.Requests[i].Url != url {
			t.Errorf("Expected url is %s, actual url is %s",
				url, parseResult.Requests[i].Url)
		}
	}

	for i, item := range expectedItems {
		if parseResult.Items[i] != item {
			t.Errorf("Expected url is %s, actual url is %s",
				item, parseResult.Items[i])
		}
	}

}
