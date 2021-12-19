// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var yue_Locale = merge(nil, LocaleData{
	Name:          "yue",
	DateOrder:     "YMD",
	NoWordSpacing: true,
	SkipWords:     []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"12月": "december",
		"星期五": "friday",
		"gmt": "gmt",
		"星期一": "monday",
		"11月": "november",
		"10月": "october",
		"星期六": "saturday",
		"星期日": "sunday",
		"星期四": "thursday",
		"星期二": "tuesday",
		"utc": "utc",
		"星期三": "wednesday",
		"am":  "am",
		"上午":  "am",
		"4月":  "april",
		"8月":  "august",
		"2月":  "february",
		"週五":  "friday",
		"小時":  "hour",
		"1月":  "january",
		"7月":  "july",
		"6月":  "june",
		"3月":  "march",
		"5月":  "may",
		"分鐘":  "minute",
		"週一":  "monday",
		"pm":  "pm",
		"下午":  "pm",
		"週六":  "saturday",
		"9月":  "september",
		"週日":  "sunday",
		"週四":  "thursday",
		"週二":  "tuesday",
		"週三":  "wednesday",
		"'":   "",
		",":   "",
		";":   "",
		"@":   "",
		"[":   "",
		"]":   "",
		"|":   "",
		" ":   " ",
		"+":   "+",
		"-":   "-",
		".":   ".",
		"/":   "/",
		":":   ":",
		"日":   "day",
		"月":   "month",
		"秒":   "second",
		"週":   "week",
		"年":   "year",
		"z":   "z",
	},
	RelativeType: map[string]string{
		"呢個小時": "0 hour ago",
		"今個星期": "0 week ago",
		"呢分鐘":  "0 minute ago",
		"今個月":  "0 month ago",
		"上個月":  "1 month ago",
		"上星期":  "1 week ago",
		"下個月":  "in 1 month",
		"下星期":  "in 1 week",
		"今日":   "0 day ago",
		"宜家":   "0 second ago",
		"今年":   "0 year ago",
		"尋日":   "1 day ago",
		"舊年":   "1 year ago",
		"聽日":   "in 1 day",
		"下年":   "in 1 year",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) 個星期前`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) 個星期後`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) 小時前`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) 分鐘前`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) 個月前`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) 小時後`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) 分鐘後`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) 個月後`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) 日前`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) 秒前`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) 年前`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) 日後`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) 秒後`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) 年後`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\d+ 個星期前|\d+ 個星期後|\d+ 個月前|\d+ 個月後|\d+ 分鐘前|\d+ 分鐘後|\d+ 小時前|\d+ 小時後|\d+ 年前|\d+ 年後|\d+ 日前|\d+ 日後|\d+ 秒前|\d+ 秒後)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ 個星期前|\d+ 個星期後|\d+ 個月前|\d+ 個月後|\d+ 分鐘前|\d+ 分鐘後|\d+ 小時前|\d+ 小時後|\d+ 年前|\d+ 年後|\d+ 日前|\d+ 日後|\d+ 秒前|\d+ 秒後)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?)(今個星期|呢個小時|10月|11月|12月|gmt|utc|上個月|上星期|下個月|下星期|今個月|呢分鐘|星期一|星期三|星期二|星期五|星期六|星期四|星期日|1月|2月|3月|4月|5月|6月|7月|8月|9月|\+|\.|\[|\]|\||am|pm|上午|下午|下年|今年|今日|分鐘|宜家|尋日|小時|聽日|舊年|週一|週三|週二|週五|週六|週四|週日| |'|,|-|/|:|;|@|z|年|日|月|秒|週)(.*)$`),
})
