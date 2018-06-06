package parser

import (
	"regexp"
	"justforunc/engine"
	"fmt"
	"strconv"
	"justforunc/model"
)

var nameRe = regexp.MustCompile(`<a class="name fs24">([^<]+)</a>`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([^<]+)</span></td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<]+)</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occuRe = regexp.MustCompile(`<td><span class="label">公司：</span><span field="">([^<]+)</span></td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">工作地区：</span><span field="">([^<]+)</span></td>`)
var xingzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)


func ParseProfile(contens []byte, name string) engine.ParseResult {

	profile := model.Profile{}


	profile.Name = name
	profile.Gender = extractString(contens, genderRe)

	age, err := strconv.Atoi(extractString(contens, ageRe))
	if err != nil {
		fmt.Errorf("illegal age")
	}
	profile.Age = age

	height, err := strconv.Atoi(extractString(contens, heightRe))
	if err != nil {
		fmt.Errorf("illegal height")
	}
	profile.Height = height

	weight, err := strconv.Atoi(extractString(contens, weightRe))
	if err != nil {
		fmt.Errorf("illegal weight")
	}
	profile.Weight = weight

	profile.Income = extractString(contens, incomeRe)
	profile.Marriage = extractString(contens, marriageRe)
	profile.Education = extractString(contens, educationRe)
	profile.Occupation = extractString(contens, occuRe)
	profile.Hukou = extractString(contens, hukouRe)
	profile.Xingzuo =extractString(contens, xingzuoRe)
	profile.House = extractString(contens, houseRe)
	profile.Car = extractString(contens, carRe)

	result := engine.ParseResult{}
	result.Items = append(result.Items, profile)

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
