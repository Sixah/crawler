package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var (
	nameRe       = regexp.MustCompile(`<a class="name[^>]*">([^<]+)</a>`)
	genderRe     = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	ageRe        = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
	heightRe     = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
	weightRe     = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
	incomeRe     = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	marriageRe   = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	educationRe  = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
	occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
	workPlaceRe  = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
	hukouRe      = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
	xingzuoRe    = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
	houseRe      = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe        = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
	shengXiaoRe  = regexp.MustCompile(`<td><span class="label">生肖：</span><span field="">([^<]+)</span></td>`)
	bodyTypeRe   = regexp.MustCompile(`<td><span class="label">体型：</span><span field="">([^<]+)</span></td>`)
	monologRe    = regexp.MustCompile(`<p class="fs14 lh20 c5e slider-area-js[^"]*"[^>]*>([^<]+)<span class="info-mark"></span></p>`)
	idUrlRe      = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
	urlRe        = regexp.MustCompile(`<link rel="canonical" href="(http://album.zhenai.com/u/[^"]+)" />`)
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
	profile.WorkPlace = extractString(contents, workPlaceRe)
	profile.ShengXiao = extractString(contents, shengXiaoRe)
	profile.BodyType = extractString(contents, bodyTypeRe)
	profile.Monolog = extractString(contents, monologRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	url := extractString(contents, urlRe)
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	// matches := guessRe.FindAllSubmatch(contents,-1)
	// for _,m := range matches {
	// 	name := string(m[2])
	// 	url := string(m[1])
	// 	result.Requests = append(result.Requests,engine.Request{
	// 		Url: url,
	// 		ParserFunc: func(c []byte) engine.ParseResult {
	// 			return ParseProfile(c, url,name)
	// 		}
	// 	})
	// }
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
