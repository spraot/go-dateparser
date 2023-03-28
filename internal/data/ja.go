// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var (
	ja_Locale LocaleData
)

func init() {
	ja_Locale = merge(nil, LocaleData{
		Name:                  "ja",
		DateOrder:             "YMD",
		NoWordSpacing:         true,
		Charset:               []rune(`cgtuz々かてでのらカヶ一七三九二五今以先八六内分前十午去四土在年後日明昨時曜月木水火現秒約翌週金間`),
		SentenceSplitterGroup: 4,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年(?:\s+)?(\d+[.,]?\d*)月(?:\s+)?(\d+[.,]?\d*)日`), "${1}-${2}-${3}"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)月(?:\s+)?(\d+[.,]?\d*)日`), "${1}-${2}"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)時(?:\s+)?(\d+[.,]?\d*)分(?:\s+)?(\d+[.,]?\d*)秒`), "${1}:${2}:${3}"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)時(?:\s+)?(\d+[.,]?\d*)分`), "${1}:${2}"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)時$`), "${1}:00"},
			{regexp.MustCompile(`(?i)正午`), "12:00"},
		},
		Translations: map[string][]string{
			"(土)": {"saturday"},
			"(日)": {"sunday"},
			"(月)": {"monday"},
			"(木)": {"thursday"},
			"(水)": {"wednesday"},
			"(火)": {"tuesday"},
			"(金)": {"friday"},
			"10月": {"october"},
			"11月": {"november"},
			"12月": {"december"},
			"gmt": {"gmt"},
			"utc": {"utc"},
			"今から": {"in"},
			"十一月": {"november"},
			"十二月": {"december"},
			"土曜日": {"saturday"},
			"日曜日": {"sunday"},
			"月曜日": {"monday"},
			"木曜日": {"thursday"},
			"水曜日": {"wednesday"},
			"火曜日": {"tuesday"},
			"金曜日": {"friday"},
			"1月":  {"january"},
			"2月":  {"february"},
			"3月":  {"march"},
			"4月":  {"april"},
			"5月":  {"may"},
			"6月":  {"june"},
			"7月":  {"july"},
			"8月":  {"august"},
			"9月":  {"september"},
			"am":  {"am"},
			"pm":  {"pm"},
			"か月":  {"month"},
			"カ月":  {"month"},
			"ヶ月":  {"month"},
			"一月":  {"january"},
			"七月":  {"july"},
			"三月":  {"march"},
			"九月":  {"september"},
			"二月":  {"february"},
			"五月":  {"may"},
			"八月":  {"august"},
			"六月":  {"june"},
			"分間":  {"minute"},
			"十月":  {"october"},
			"午前":  {"am"},
			"午後":  {"pm"},
			"四月":  {"april"},
			"日間":  {"day"},
			"時間":  {"hour"},
			"秒間":  {"second"},
			"週間":  {"week"},
			" ":   {" "},
			"'":   {""},
			"+":   {"+"},
			",":   {""},
			"-":   {"-"},
			".":   {"."},
			"/":   {"/"},
			":":   {":"},
			";":   {""},
			"@":   {""},
			"[":   {""},
			"]":   {""},
			"z":   {"z"},
			"|":   {""},
			"て":   {"in"},
			"の":   {""},
			"分":   {"minute"},
			"前":   {"ago"},
			"土":   {"saturday"},
			"年":   {"year"},
			"日":   {"day", "sunday"},
			"時":   {"hour"},
			"月":   {"month", "monday"},
			"木":   {"thursday"},
			"水":   {"wednesday"},
			"火":   {"tuesday"},
			"秒":   {"second"},
			"約":   {""},
			"週":   {"week"},
			"金":   {"friday"},
		},
		RelativeType: map[string]string{
			"1 時間以内": "0 hour ago",
			"1 分以内":  "0 minute ago",
			"一昨日":    "2 day ago",
			"先々週":    "2 week ago",
			"明後日":    "in 2 day",
			"今年":     "0 year ago",
			"今日":     "0 day ago",
			"今月":     "0 month ago",
			"今週":     "0 week ago",
			"先月":     "1 month ago",
			"先週":     "1 week ago",
			"去年":     "1 year ago",
			"明日":     "in 1 day",
			"昨年":     "1 year ago",
			"昨日":     "1 day ago",
			"現在":     "0 second ago",
			"翌年":     "in 1 year",
			"翌月":     "in 1 month",
			"翌週":     "in 1 week",
			"今":      "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) か月前`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) か月後`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 時間前`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 時間後`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 週間前`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 週間後`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 分前`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 分後`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 年前`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 年後`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 日前`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 日後`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 秒前`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 秒後`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)か月前`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)か月後`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)時間前`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)時間後`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)週間前`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)週間後`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)分前`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)分後`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年前`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年後`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)日前`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)日後`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)秒前`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)秒後`), "in $1 second"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\d+[.,]?\d* か月前|\d+[.,]?\d* か月後|\d+[.,]?\d* 時間前|\d+[.,]?\d* 時間後|\d+[.,]?\d* 週間前|\d+[.,]?\d* 週間後|\d+[.,]?\d* 分前|\d+[.,]?\d* 分後|\d+[.,]?\d* 年前|\d+[.,]?\d* 年後|\d+[.,]?\d* 日前|\d+[.,]?\d* 日後|\d+[.,]?\d* 秒前|\d+[.,]?\d* 秒後|\d+[.,]?\d*か月前|\d+[.,]?\d*か月後|\d+[.,]?\d*時間前|\d+[.,]?\d*時間後|\d+[.,]?\d*週間前|\d+[.,]?\d*週間後|\d+[.,]?\d*分前|\d+[.,]?\d*分後|\d+[.,]?\d*年前|\d+[.,]?\d*年後|\d+[.,]?\d*日前|\d+[.,]?\d*日後|\d+[.,]?\d*秒前|\d+[.,]?\d*秒後)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* か月前|\d+[.,]?\d* か月後|\d+[.,]?\d* 時間前|\d+[.,]?\d* 時間後|\d+[.,]?\d* 週間前|\d+[.,]?\d* 週間後|\d+[.,]?\d* 分前|\d+[.,]?\d* 分後|\d+[.,]?\d* 年前|\d+[.,]?\d* 年後|\d+[.,]?\d* 日前|\d+[.,]?\d* 日後|\d+[.,]?\d* 秒前|\d+[.,]?\d* 秒後|\d+[.,]?\d*か月前|\d+[.,]?\d*か月後|\d+[.,]?\d*時間前|\d+[.,]?\d*時間後|\d+[.,]?\d*週間前|\d+[.,]?\d*週間後|\d+[.,]?\d*分前|\d+[.,]?\d*分後|\d+[.,]?\d*年前|\d+[.,]?\d*年後|\d+[.,]?\d*日前|\d+[.,]?\d*日後|\d+[.,]?\d*秒前|\d+[.,]?\d*秒後)$`),
		KnownWords:      []string{"1 時間以内", "1 分以内", "(土)", "(日)", "(月)", "(木)", "(水)", "(火)", "(金)", "10月", "11月", "12月", "gmt", "utc", "一昨日", "今から", "先々週", "十一月", "十二月", "土曜日", "日曜日", "明後日", "月曜日", "木曜日", "水曜日", "火曜日", "金曜日", "1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "am", "pm", "か月", "カ月", "ヶ月", "一月", "七月", "三月", "九月", "二月", "五月", "今年", "今日", "今月", "今週", "先月", "先週", "八月", "六月", "分間", "十月", "午前", "午後", "去年", "四月", "日間", "明日", "昨年", "昨日", "時間", "現在", "秒間", "翌年", "翌月", "翌週", "週間", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "て", "の", "今", "分", "前", "土", "年", "日", "時", "月", "木", "水", "火", "秒", "約", "週", "金"},
	})
}
