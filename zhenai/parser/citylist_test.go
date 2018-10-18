package parser

import (
	"testing"
)

func TestParseCityList(t *testing.T) {
	// contents, err := ioutil.ReadFile("citylist_test_data.html")
	// if err != nil {
	// 	t.Log(err)
	// }
	// t.Logf("%s\n", contents)
	/*
		file, err := os.Create("citytext.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		write := bufio.NewWriter(file)
		defer write.Flush()

		fmt.Fprintf(write, "%s", contents)
	*/
	// result := ParserCityList(contents)
	// const resultSize = 470
	// expectedUrls := []string{
	// 	"http://www.zhenai.com/zhenghun/aba",
	// 	"http://www.zhenai.com/zhenghun/baicheng1",
	// 	"http://www.zhenai.com/zhenghun/cangzhou",
	// }
	// expectedCites := []string{}
	// if len(result.Requests) != resultSize {
	// 	t.Errorf("result should hava %d requests,but had %d", resultSize, len(result.Requests))
	// }
	// for i, url := range expectedUrls {
	// 	if result.Requests[i].Url != url {
	// 		t.Errorf("expected url #%d: %s,but was %s", i, url, result.Requests[i].Url)
	// 	}
	// }
	//
	// if len(result.Items) != resultSize {
	// 	t.Errorf("result should hava %d requests,but had %d", resultSize, len(result.Items))
	// }
	// for i, city := range expectedCites {
	// 	if result.Items[i].(string) != city {
	// 		t.Errorf("expected url #%d: %s,but was %s", i, city, result.Items[i].(string))
	// 	}
	// }
}
