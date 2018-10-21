package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var (
	nameRe = regexp.MustCompile(`<a class="name[^>]*">([^<]+)</a>`)
	// nameRe       = regexp.MustCompile(`<td><span class="label">姓名：</span>([^<]+)</td>`)
	genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	ageRe    = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
	heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
	// weightRe     = regexp.MustCompile(`<td><span class="label">体重：</span>([\d]+)KG</td>`)
	weightRe     = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
	incomeRe     = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	marriageRe   = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
	hukouRe      = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	xingzuoRe    = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
	houseRe      = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe        = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	profile.Name = extractString(contents, nameRe)
	profile.Gender = extractString(contents, genderRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	// fmt.Println("result:", result)
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
