// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var zh_Hant_Locale = LocaleData{
	Name:                  "zh-Hant",
	DateOrder:             "YMD",
	NoWordSpacing:         true,
	SentenceSplitterGroup: 4,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|"},
	January:               []string{"1月"},
	February:              []string{"2月"},
	March:                 []string{"3月"},
	April:                 []string{"4月"},
	May:                   []string{"5月"},
	June:                  []string{"6月"},
	July:                  []string{"7月"},
	August:                []string{"8月"},
	September:             []string{"9月"},
	October:               []string{"10月"},
	November:              []string{"11月"},
	December:              []string{"12月"},
	Monday:                []string{"星期一", "週一"},
	Tuesday:               []string{"星期二", "週二"},
	Wednesday:             []string{"星期三", "週三"},
	Thursday:              []string{"星期四", "週四"},
	Friday:                []string{"星期五", "週五"},
	Saturday:              []string{"星期六", "週六"},
	Sunday:                []string{"星期日", "週日"},
	AM:                    []string{"上午"},
	PM:                    []string{"下午"},
	Year:                  []string{"年"},
	Month:                 []string{"月"},
	Week:                  []string{"週"},
	Day:                   []string{"日"},
	Hour:                  []string{"小時"},
	Minute:                []string{"分鐘"},
	Second:                []string{"秒"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`今天`},
		`0 hour ago`:   {`這一小時`},
		`0 minute ago`: {`這一分鐘`},
		`0 month ago`:  {`本月`},
		`0 second ago`: {`現在`},
		`0 week ago`:   {`本週`},
		`0 year ago`:   {`今年`},
		`1 day ago`:    {`昨天`},
		`1 month ago`:  {`上個月`},
		`1 week ago`:   {`上週`},
		`1 year ago`:   {`去年`},
		`in 1 day`:     {`明天`},
		`in 1 month`:   {`下個月`},
		`in 1 week`:    {`下週`},
		`in 1 year`:    {`明年`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) 天前`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) 小時前`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) 分鐘前`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) 個月前`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) 秒前`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) 週前`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) 年前`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)(\d+) 天後`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)(\d+) 小時後`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)(\d+) 分鐘後`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)(\d+) 個月後`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)(\d+) 秒後`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)(\d+) 週後`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)(\d+) 年後`),
		},
	},
}

var zh_Hant_HK_Locale = LocaleData{
	Name:                  "zh-Hant-HK",
	DateOrder:             "DMY",
	NoWordSpacing:         true,
	SentenceSplitterGroup: 4,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|"},
	January:               []string{"1月"},
	February:              []string{"2月"},
	March:                 []string{"3月"},
	April:                 []string{"4月"},
	May:                   []string{"5月"},
	June:                  []string{"6月"},
	July:                  []string{"7月"},
	August:                []string{"8月"},
	September:             []string{"9月"},
	October:               []string{"10月"},
	November:              []string{"11月"},
	December:              []string{"12月"},
	Monday:                []string{"星期一", "週一"},
	Tuesday:               []string{"星期二", "週二"},
	Wednesday:             []string{"星期三", "週三"},
	Thursday:              []string{"星期四", "週四"},
	Friday:                []string{"星期五", "週五"},
	Saturday:              []string{"星期六", "週六"},
	Sunday:                []string{"星期日", "週日"},
	AM:                    []string{"上午"},
	PM:                    []string{"下午"},
	Year:                  []string{"年"},
	Month:                 []string{"月"},
	Week:                  []string{"星期", "週"},
	Day:                   []string{"日"},
	Hour:                  []string{"小時", "時"},
	Minute:                []string{"分", "分鐘"},
	Second:                []string{"秒"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`今日`},
		`0 hour ago`:   {`這個小時`},
		`0 minute ago`: {`這分鐘`},
		`0 month ago`:  {`本月`},
		`0 second ago`: {`現在`},
		`0 week ago`:   {`本星期`},
		`0 year ago`:   {`今年`},
		`1 day ago`:    {`昨日`},
		`1 month ago`:  {`上個月`, `上月`},
		`1 week ago`:   {`上星期`},
		`1 year ago`:   {`上年`},
		`in 1 day`:     {`明日`},
		`in 1 month`:   {`下個月`, `下月`},
		`in 1 week`:    {`下星期`},
		`in 1 year`:    {`下年`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) 日前`),
			regexp.MustCompile(`(?i)(\d+)日前`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) 小時前`),
			regexp.MustCompile(`(?i)(\d+)小時前`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) 分鐘前`),
			regexp.MustCompile(`(?i)(\d+)分前`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) 個月前`),
			regexp.MustCompile(`(?i)(\d+)個月前`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) 秒前`),
			regexp.MustCompile(`(?i)(\d+)秒前`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) 星期前`),
			regexp.MustCompile(`(?i)(\d+)週前`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) 年前`),
			regexp.MustCompile(`(?i)(\d+)年前`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)(\d+) 日後`),
			regexp.MustCompile(`(?i)(\d+)日後`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)(\d+) 小時後`),
			regexp.MustCompile(`(?i)(\d+)小時後`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)(\d+) 分鐘後`),
			regexp.MustCompile(`(?i)(\d+)分後`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)(\d+) 個月後`),
			regexp.MustCompile(`(?i)(\d+)個月後`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)(\d+) 秒後`),
			regexp.MustCompile(`(?i)(\d+)秒後`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)(\d+) 星期後`),
			regexp.MustCompile(`(?i)(\d+)週後`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)(\d+) 年後`),
			regexp.MustCompile(`(?i)(\d+)年後`),
		},
	},
}

var zh_Hant_MO_Locale = LocaleData{
	Name:                  "zh-Hant-MO",
	DateOrder:             "DMY",
	NoWordSpacing:         true,
	SentenceSplitterGroup: 4,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|"},
	January:               []string{"1月"},
	February:              []string{"2月"},
	March:                 []string{"3月"},
	April:                 []string{"4月"},
	May:                   []string{"5月"},
	June:                  []string{"6月"},
	July:                  []string{"7月"},
	August:                []string{"8月"},
	September:             []string{"9月"},
	October:               []string{"10月"},
	November:              []string{"11月"},
	December:              []string{"12月"},
	Monday:                []string{"星期一", "週一"},
	Tuesday:               []string{"星期二", "週二"},
	Wednesday:             []string{"星期三", "週三"},
	Thursday:              []string{"星期四", "週四"},
	Friday:                []string{"星期五", "週五"},
	Saturday:              []string{"星期六", "週六"},
	Sunday:                []string{"星期日", "週日"},
	AM:                    []string{"上午"},
	PM:                    []string{"下午"},
	Year:                  []string{"年"},
	Month:                 []string{"月"},
	Week:                  []string{"星期", "週"},
	Day:                   []string{"日"},
	Hour:                  []string{"小時", "時"},
	Minute:                []string{"分", "分鐘"},
	Second:                []string{"秒"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`今日`},
		`0 hour ago`:   {`這個小時`},
		`0 minute ago`: {`這分鐘`},
		`0 month ago`:  {`本月`},
		`0 second ago`: {`現在`},
		`0 week ago`:   {`本星期`},
		`0 year ago`:   {`今年`},
		`1 day ago`:    {`昨日`},
		`1 month ago`:  {`上個月`, `上月`},
		`1 week ago`:   {`上星期`},
		`1 year ago`:   {`上年`},
		`in 1 day`:     {`明日`},
		`in 1 month`:   {`下個月`, `下月`},
		`in 1 week`:    {`下星期`},
		`in 1 year`:    {`下年`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) 日前`),
			regexp.MustCompile(`(?i)(\d+)日前`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) 小時前`),
			regexp.MustCompile(`(?i)(\d+)小時前`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) 分鐘前`),
			regexp.MustCompile(`(?i)(\d+)分前`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) 個月前`),
			regexp.MustCompile(`(?i)(\d+)個月前`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) 秒前`),
			regexp.MustCompile(`(?i)(\d+)秒前`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) 星期前`),
			regexp.MustCompile(`(?i)(\d+)週前`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) 年前`),
			regexp.MustCompile(`(?i)(\d+)年前`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)(\d+) 日後`),
			regexp.MustCompile(`(?i)(\d+)日後`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)(\d+) 小時後`),
			regexp.MustCompile(`(?i)(\d+)小時後`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)(\d+) 分鐘後`),
			regexp.MustCompile(`(?i)(\d+)分後`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)(\d+) 個月後`),
			regexp.MustCompile(`(?i)(\d+)個月後`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)(\d+) 秒後`),
			regexp.MustCompile(`(?i)(\d+)秒後`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)(\d+) 星期後`),
			regexp.MustCompile(`(?i)(\d+)週後`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)(\d+) 年後`),
			regexp.MustCompile(`(?i)(\d+)年後`),
		},
	},
}
