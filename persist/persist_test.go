package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1505672417",
		Type: "zhenai",
		Id:   "1505672417",
		Payload: model.Profile{
			Name:       "元气老阿姨",
			Gender:     "女",
			Age:        25,
			ShengXiao:  "鸡",
			Height:     160,
			Weight:     50,
			BodyType:   "一般",
			Xingzuo:    "天秤座",
			Education:  "高中及以下",
			Occupation: "销售",
			Income:     "8001-12000元",
			Marriage:   "未婚",
			House:      "租房",
			Car:        "未购车",
			WorkPlace:  "浙江杭州",
			Hukou:      "浙江杭州",
			Monolog:    "性格一般，脾气一点就炸。但是来的快，去的也快。希望未来的另一半能多包容一些我的坏脾气。喜欢那种温柔性格的男生请绕道。大天蝎，性子直，有什么说什么，不喜欢藏着掖着。除了不会做饭，其他家务都会。可以学做饭，但是得出现一个我愿意为他做饭的人。其他好像也没什么想说的了，感情方面，都是有过故事的人，如果你对我有兴趣，那就慢慢了解吧。至于对另一半的要求目前就外在条件肯定要能带的出门，不能太胖或者矮，不接受170以下男生，我比较注重眼缘，第一眼的感觉挺重要，其他的了解以后看情况吧",
		},
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	// Save expected item
	err = save(client, "sixah", expected)
	if err != nil {
		t.Logf("save error: %v", err)
		panic(err)
	}

	// TODO: Try to start up elastic search
	// here using docker go client

	// Fetch saved item
	resp, err := client.Get().Index("sixah").Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	t.Logf("%v", actual)

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
