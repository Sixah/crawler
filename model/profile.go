package model

import "encoding/json"

type Profile struct {
	// Id         string // 用户id
	Name       string // 姓名
	Gender     string // 性别
	Age        int    // 年龄
	ShengXiao  string // 生肖
	Height     int    // 身高
	Weight     int    // 体重
	BodyType   string // 体型
	Income     string // 收入
	Marriage   string // 婚姻状况
	Education  string // 职业
	Occupation string // 学历
	WorkPlace  string // 工作地
	Hukou      string // 户口/籍贯
	Xingzuo    string // 星座
	House      string // 购房条件
	Car        string // 是否购车
	Monolog    string // 内心独白
	// Url        string // 链接
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}
