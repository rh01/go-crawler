package parser

import (
	"testing"
	"justforunc/fetcher"
)

func TestParseUser(t *testing.T) {
	contens, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/aomen")
	if err != nil {
		panic(err)
	}

	parseResult := ParseUser(contens)

	// verify parseList

	const resultSize int = 16
	expectedUrls := []string{
		"http://album.zhenai.com/u/1122926649",
		"http://album.zhenai.com/u/107045901",
		"http://album.zhenai.com/u/109053717",
	}

	expectedItems := []string{
		"User 放风筝的人", "User 拼了命的开心", "User Gave",
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
