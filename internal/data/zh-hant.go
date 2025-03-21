// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	zh_Hant_Locale    LocaleData
	zh_Hant_HK_Locale LocaleData
	zh_Hant_MO_Locale LocaleData
)

func init() {
	zh_Hant_Locale = merge(nil, LocaleData{
		Name:                  "zh-Hant",
		DateOrder:             "YMD",
		NoWordSpacing:         true,
		Charset:               []rune(`cgtuz一三上下二五今個六分午去四在天小年日明星昨時月期本現秒這週鐘`),
		SentenceSplitterGroup: 4,
		Translations: map[string][]string{
			"10月": {"october"},
			"11月": {"november"},
			"12月": {"december"},
			"gmt": {"gmt"},
			"utc": {"utc"},
			"星期一": {"monday"},
			"星期三": {"wednesday"},
			"星期二": {"tuesday"},
			"星期五": {"friday"},
			"星期六": {"saturday"},
			"星期四": {"thursday"},
			"星期日": {"sunday"},
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
			"上午":  {"am"},
			"下午":  {"pm"},
			"分鐘":  {"minute"},
			"小時":  {"hour"},
			"週一":  {"monday"},
			"週三":  {"wednesday"},
			"週二":  {"tuesday"},
			"週五":  {"friday"},
			"週六":  {"saturday"},
			"週四":  {"thursday"},
			"週日":  {"sunday"},
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
			"年":   {"year"},
			"日":   {"day"},
			"月":   {"month"},
			"秒":   {"second"},
			"週":   {"week"},
		},
		RelativeType: map[string]string{
			"這一分鐘": "0 minute ago",
			"這一小時": "0 hour ago",
			"上個月":  "1 month ago",
			"下個月":  "in 1 month",
			"上週":   "1 week ago",
			"下週":   "in 1 week",
			"今天":   "0 day ago",
			"今年":   "0 year ago",
			"去年":   "1 year ago",
			"明天":   "in 1 day",
			"明年":   "in 1 year",
			"昨天":   "1 day ago",
			"本月":   "0 month ago",
			"本週":   "0 week ago",
			"現在":   "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 個月前`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 個月後`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 分鐘前`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 分鐘後`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 小時前`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 小時後`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 天前`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 天後`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 年前`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 年後`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 秒前`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 秒後`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 週前`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 週後`), "in $1 week"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\d+[.,]?\d* 個月前|\d+[.,]?\d* 個月後|\d+[.,]?\d* 分鐘前|\d+[.,]?\d* 分鐘後|\d+[.,]?\d* 小時前|\d+[.,]?\d* 小時後|\d+[.,]?\d* 天前|\d+[.,]?\d* 天後|\d+[.,]?\d* 年前|\d+[.,]?\d* 年後|\d+[.,]?\d* 秒前|\d+[.,]?\d* 秒後|\d+[.,]?\d* 週前|\d+[.,]?\d* 週後)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* 個月前|\d+[.,]?\d* 個月後|\d+[.,]?\d* 分鐘前|\d+[.,]?\d* 分鐘後|\d+[.,]?\d* 小時前|\d+[.,]?\d* 小時後|\d+[.,]?\d* 天前|\d+[.,]?\d* 天後|\d+[.,]?\d* 年前|\d+[.,]?\d* 年後|\d+[.,]?\d* 秒前|\d+[.,]?\d* 秒後|\d+[.,]?\d* 週前|\d+[.,]?\d* 週後)$`),
		KnownWords:      []string{"這一分鐘", "這一小時", "10月", "11月", "12月", "gmt", "utc", "上個月", "下個月", "星期一", "星期三", "星期二", "星期五", "星期六", "星期四", "星期日", "1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "am", "pm", "上午", "上週", "下午", "下週", "今天", "今年", "分鐘", "去年", "小時", "明天", "明年", "昨天", "本月", "本週", "現在", "週一", "週三", "週二", "週五", "週六", "週四", "週日", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "年", "日", "月", "秒", "週"},
	})

	zh_Hant_HK_Locale = merge(&zh_Hant_Locale, LocaleData{
		Name:          "zh-Hant-HK",
		DateOrder:     "DMY",
		NoWordSpacing: true,
		Translations: map[string][]string{
			"星期": {"week"},
			"分":  {"minute"},
			"時":  {"hour"},
		},
		RelativeType: map[string]string{
			"這個小時": "0 hour ago",
			"上星期":  "1 week ago",
			"下星期":  "in 1 week",
			"本星期":  "0 week ago",
			"這分鐘":  "0 minute ago",
			"上年":   "1 year ago",
			"上月":   "1 month ago",
			"下年":   "in 1 year",
			"下月":   "in 1 month",
			"今日":   "0 day ago",
			"明日":   "in 1 day",
			"昨日":   "1 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 星期前`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 星期後`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 日前`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 日後`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)個月前`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)個月後`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)小時前`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)小時後`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)分前`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)分後`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年前`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年後`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)日前`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)日後`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)秒前`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)秒後`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)週前`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)週後`), "in $1 week"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\d+[.,]?\d* 個月前|\d+[.,]?\d* 個月後|\d+[.,]?\d* 分鐘前|\d+[.,]?\d* 分鐘後|\d+[.,]?\d* 小時前|\d+[.,]?\d* 小時後|\d+[.,]?\d* 星期前|\d+[.,]?\d* 星期後|\d+[.,]?\d* 天前|\d+[.,]?\d* 天後|\d+[.,]?\d* 年前|\d+[.,]?\d* 年後|\d+[.,]?\d* 日前|\d+[.,]?\d* 日後|\d+[.,]?\d* 秒前|\d+[.,]?\d* 秒後|\d+[.,]?\d* 週前|\d+[.,]?\d* 週後|\d+[.,]?\d*個月前|\d+[.,]?\d*個月後|\d+[.,]?\d*小時前|\d+[.,]?\d*小時後|\d+[.,]?\d*分前|\d+[.,]?\d*分後|\d+[.,]?\d*年前|\d+[.,]?\d*年後|\d+[.,]?\d*日前|\d+[.,]?\d*日後|\d+[.,]?\d*秒前|\d+[.,]?\d*秒後|\d+[.,]?\d*週前|\d+[.,]?\d*週後)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* 個月前|\d+[.,]?\d* 個月後|\d+[.,]?\d* 分鐘前|\d+[.,]?\d* 分鐘後|\d+[.,]?\d* 小時前|\d+[.,]?\d* 小時後|\d+[.,]?\d* 星期前|\d+[.,]?\d* 星期後|\d+[.,]?\d* 天前|\d+[.,]?\d* 天後|\d+[.,]?\d* 年前|\d+[.,]?\d* 年後|\d+[.,]?\d* 日前|\d+[.,]?\d* 日後|\d+[.,]?\d* 秒前|\d+[.,]?\d* 秒後|\d+[.,]?\d* 週前|\d+[.,]?\d* 週後|\d+[.,]?\d*個月前|\d+[.,]?\d*個月後|\d+[.,]?\d*小時前|\d+[.,]?\d*小時後|\d+[.,]?\d*分前|\d+[.,]?\d*分後|\d+[.,]?\d*年前|\d+[.,]?\d*年後|\d+[.,]?\d*日前|\d+[.,]?\d*日後|\d+[.,]?\d*秒前|\d+[.,]?\d*秒後|\d+[.,]?\d*週前|\d+[.,]?\d*週後)$`),
		KnownWords:      []string{"這一分鐘", "這一小時", "這個小時", "10月", "11月", "12月", "gmt", "utc", "上個月", "上星期", "下個月", "下星期", "星期一", "星期三", "星期二", "星期五", "星期六", "星期四", "星期日", "本星期", "這分鐘", "1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "am", "pm", "上午", "上年", "上月", "上週", "下午", "下年", "下月", "下週", "今天", "今年", "今日", "分鐘", "去年", "小時", "明天", "明年", "明日", "星期", "昨天", "昨日", "本月", "本週", "現在", "週一", "週三", "週二", "週五", "週六", "週四", "週日", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "分", "年", "日", "時", "月", "秒", "週"},
	})

	zh_Hant_MO_Locale = merge(&zh_Hant_Locale, LocaleData{
		Name:          "zh-Hant-MO",
		DateOrder:     "DMY",
		NoWordSpacing: true,
		Translations: map[string][]string{
			"星期": {"week"},
			"分":  {"minute"},
			"時":  {"hour"},
		},
		RelativeType: map[string]string{
			"這個小時": "0 hour ago",
			"上星期":  "1 week ago",
			"下星期":  "in 1 week",
			"本星期":  "0 week ago",
			"這分鐘":  "0 minute ago",
			"上年":   "1 year ago",
			"上月":   "1 month ago",
			"下年":   "in 1 year",
			"下月":   "in 1 month",
			"今日":   "0 day ago",
			"明日":   "in 1 day",
			"昨日":   "1 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 星期前`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 星期後`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 日前`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) 日後`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)個月前`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)個月後`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)小時前`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)小時後`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)分前`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)分後`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年前`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)年後`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)日前`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)日後`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)秒前`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)秒後`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)週前`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*)週後`), "in $1 week"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\d+[.,]?\d* 個月前|\d+[.,]?\d* 個月後|\d+[.,]?\d* 分鐘前|\d+[.,]?\d* 分鐘後|\d+[.,]?\d* 小時前|\d+[.,]?\d* 小時後|\d+[.,]?\d* 星期前|\d+[.,]?\d* 星期後|\d+[.,]?\d* 天前|\d+[.,]?\d* 天後|\d+[.,]?\d* 年前|\d+[.,]?\d* 年後|\d+[.,]?\d* 日前|\d+[.,]?\d* 日後|\d+[.,]?\d* 秒前|\d+[.,]?\d* 秒後|\d+[.,]?\d* 週前|\d+[.,]?\d* 週後|\d+[.,]?\d*個月前|\d+[.,]?\d*個月後|\d+[.,]?\d*小時前|\d+[.,]?\d*小時後|\d+[.,]?\d*分前|\d+[.,]?\d*分後|\d+[.,]?\d*年前|\d+[.,]?\d*年後|\d+[.,]?\d*日前|\d+[.,]?\d*日後|\d+[.,]?\d*秒前|\d+[.,]?\d*秒後|\d+[.,]?\d*週前|\d+[.,]?\d*週後)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* 個月前|\d+[.,]?\d* 個月後|\d+[.,]?\d* 分鐘前|\d+[.,]?\d* 分鐘後|\d+[.,]?\d* 小時前|\d+[.,]?\d* 小時後|\d+[.,]?\d* 星期前|\d+[.,]?\d* 星期後|\d+[.,]?\d* 天前|\d+[.,]?\d* 天後|\d+[.,]?\d* 年前|\d+[.,]?\d* 年後|\d+[.,]?\d* 日前|\d+[.,]?\d* 日後|\d+[.,]?\d* 秒前|\d+[.,]?\d* 秒後|\d+[.,]?\d* 週前|\d+[.,]?\d* 週後|\d+[.,]?\d*個月前|\d+[.,]?\d*個月後|\d+[.,]?\d*小時前|\d+[.,]?\d*小時後|\d+[.,]?\d*分前|\d+[.,]?\d*分後|\d+[.,]?\d*年前|\d+[.,]?\d*年後|\d+[.,]?\d*日前|\d+[.,]?\d*日後|\d+[.,]?\d*秒前|\d+[.,]?\d*秒後|\d+[.,]?\d*週前|\d+[.,]?\d*週後)$`),
		KnownWords:      []string{"這一分鐘", "這一小時", "這個小時", "10月", "11月", "12月", "gmt", "utc", "上個月", "上星期", "下個月", "下星期", "星期一", "星期三", "星期二", "星期五", "星期六", "星期四", "星期日", "本星期", "這分鐘", "1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "am", "pm", "上午", "上年", "上月", "上週", "下午", "下年", "下月", "下週", "今天", "今年", "今日", "分鐘", "去年", "小時", "明天", "明年", "明日", "星期", "昨天", "昨日", "本月", "本週", "現在", "週一", "週三", "週二", "週五", "週六", "週四", "週日", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "分", "年", "日", "時", "月", "秒", "週"},
	})
}
