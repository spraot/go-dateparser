// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var am_Locale = merge(nil, LocaleData{
	Name:      "am",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"ሴፕቴምበር": "september",
		"ዲሴምበር":  "december",
		"ፌብሩወሪ":  "february",
		"ጃንዩወሪ":  "january",
		"ኖቬምበር":  "november",
		"ኦክቶበር":  "october",
		"ኤፕሪል":   "april",
		"ኦገስት":   "august",
		"ከሰዓት":   "pm",
		"ሰከንድ":   "second",
		"ማክሰኞ":   "tuesday",
		"ሳምንት":   "week",
		"ጥዋት":    "am",
		"ኤፕሪ":    "april",
		"ኦገስ":    "august",
		"ዲሴም":    "december",
		"ፌብሩ":    "february",
		"ዓርብ":    "friday",
		"gmt":    "gmt",
		"ሰዓት":    "hour",
		"ጃንዩ":    "january",
		"ጁላይ":    "july",
		"ማርች":    "march",
		"ደቂቃ":    "minute",
		"ኖቬም":    "november",
		"ኦክቶ":    "october",
		"ቅዳሜ":    "saturday",
		"ሴፕቴ":    "september",
		"እሑድ":    "sunday",
		"ሐሙስ":    "thursday",
		"ማክሰ":    "tuesday",
		"utc":    "utc",
		"ረቡዕ":    "wednesday",
		"ዓመት":    "year",
		"am":     "am",
		"ቀን":     "day",
		"ጁን":     "june",
		"ሜይ":     "may",
		"ሰኞ":     "monday",
		"ወር":     "month",
		"pm":     "pm",
		"'":      "",
		",":      "",
		";":      "",
		"@":      "",
		"[":      "",
		"]":      "",
		"|":      "",
		" ":      " ",
		"+":      "+",
		"-":      "-",
		".":      ".",
		"/":      "/",
		":":      ":",
		"z":      "z",
	},
	RelativeType: map[string]string{
		"የሚቀጥለው ሳምንት": "in 1 week",
		"የሚቀጥለው ዓመት":  "in 1 year",
		"ባለፈው ሳምንት":   "1 week ago",
		"ያለፈው ሳምንት":   "1 week ago",
		"የሚቀጥለው ወር":   "in 1 month",
		"በዚህ ሣምንት":    "0 week ago",
		"በዚህ ሳምንት":    "0 week ago",
		"ያለፈው ዓመት":    "1 year ago",
		"በዚህ ዓመት":     "0 year ago",
		"ያለፈው ወር":     "1 month ago",
		"ይህ ሰዓት":      "0 hour ago",
		"ይህ ደቂቃ":      "0 minute ago",
		"በዚህ ወር":      "0 month ago",
		"ትላንትና":       "1 day ago",
		"ትናንት":        "1 day ago",
		"አሁን":         "0 second ago",
		"ዛሬ":          "0 day ago",
		"ነገ":          "in 1 day",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)ከ(\d+) ደቂቃዎች በፊት`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ሰከንዶች በፊት`), "$1 second ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ሳምንታት በፊት`), "$1 week ago"},
		{regexp.MustCompile(`(?i)በ(\d+) ደቂቃዎች ውስጥ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)በ(\d+) ሰከንዶች ውስጥ`), "in $1 second"},
		{regexp.MustCompile(`(?i)በ(\d+) ሳምንታት ውስጥ`), "in $1 week"},
		{regexp.MustCompile(`(?i)ከ(\d+) ሰዓቶች በፊት`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ሰከንድ በፊት`), "$1 second ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ሳምንት በፊት`), "$1 week ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ዓመታት በፊት`), "$1 year ago"},
		{regexp.MustCompile(`(?i)በ(\d+) ሰዓቶች ውስጥ`), "in $1 hour"},
		{regexp.MustCompile(`(?i)በ(\d+) ሰከንድ ውስጥ`), "in $1 second"},
		{regexp.MustCompile(`(?i)በ(\d+) ሳምንት ውስጥ`), "in $1 week"},
		{regexp.MustCompile(`(?i)በ(\d+) ዓመታት ውስጥ`), "in $1 year"},
		{regexp.MustCompile(`(?i)ከ (\d+) ቀን በፊት`), "$1 day ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ቀናት በፊት`), "$1 day ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ቀኖች በፊት`), "$1 day ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ሰዓት በፊት`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ደቂቃ በፊት`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ወራት በፊት`), "$1 month ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ዓመት በፊት`), "$1 year ago"},
		{regexp.MustCompile(`(?i)በ(\d+) ቀናት ውስጥ`), "in $1 day"},
		{regexp.MustCompile(`(?i)በ(\d+) ቀኖች ውስጥ`), "in $1 day"},
		{regexp.MustCompile(`(?i)በ(\d+) ሰዓት ውስጥ`), "in $1 hour"},
		{regexp.MustCompile(`(?i)በ(\d+) ደቂቃ ውስጥ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)በ(\d+) ወራት ውስጥ`), "in $1 month"},
		{regexp.MustCompile(`(?i)ከ(\d+) ቀን በፊት`), "$1 day ago"},
		{regexp.MustCompile(`(?i)ከ(\d+) ወር በፊት`), "$1 month ago"},
		{regexp.MustCompile(`(?i)በ(\d+) ቀን ውስጥ`), "in $1 day"},
		{regexp.MustCompile(`(?i)በ(\d+) ወር ውስጥ`), "in $1 month"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(በ\d+ ሰከንዶች ውስጥ|በ\d+ ሳምንታት ውስጥ|በ\d+ ደቂቃዎች ውስጥ|ከ\d+ ሰከንዶች በፊት|ከ\d+ ሳምንታት በፊት|ከ\d+ ደቂቃዎች በፊት|በ\d+ ሰከንድ ውስጥ|በ\d+ ሰዓቶች ውስጥ|በ\d+ ሳምንት ውስጥ|በ\d+ ዓመታት ውስጥ|ከ\d+ ሰከንድ በፊት|ከ\d+ ሰዓቶች በፊት|ከ\d+ ሳምንት በፊት|ከ\d+ ዓመታት በፊት|በ\d+ ሰዓት ውስጥ|በ\d+ ቀናት ውስጥ|በ\d+ ቀኖች ውስጥ|በ\d+ ወራት ውስጥ|በ\d+ ደቂቃ ውስጥ|ከ \d+ ቀን በፊት|ከ\d+ ሰዓት በፊት|ከ\d+ ቀናት በፊት|ከ\d+ ቀኖች በፊት|ከ\d+ ወራት በፊት|ከ\d+ ዓመት በፊት|ከ\d+ ደቂቃ በፊት|በ\d+ ቀን ውስጥ|በ\d+ ወር ውስጥ|ከ\d+ ቀን በፊት|ከ\d+ ወር በፊት)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(በ\d+ ሰከንዶች ውስጥ|በ\d+ ሳምንታት ውስጥ|በ\d+ ደቂቃዎች ውስጥ|ከ\d+ ሰከንዶች በፊት|ከ\d+ ሳምንታት በፊት|ከ\d+ ደቂቃዎች በፊት|በ\d+ ሰከንድ ውስጥ|በ\d+ ሰዓቶች ውስጥ|በ\d+ ሳምንት ውስጥ|በ\d+ ዓመታት ውስጥ|ከ\d+ ሰከንድ በፊት|ከ\d+ ሰዓቶች በፊት|ከ\d+ ሳምንት በፊት|ከ\d+ ዓመታት በፊት|በ\d+ ሰዓት ውስጥ|በ\d+ ቀናት ውስጥ|በ\d+ ቀኖች ውስጥ|በ\d+ ወራት ውስጥ|በ\d+ ደቂቃ ውስጥ|ከ \d+ ቀን በፊት|ከ\d+ ሰዓት በፊት|ከ\d+ ቀናት በፊት|ከ\d+ ቀኖች በፊት|ከ\d+ ወራት በፊት|ከ\d+ ዓመት በፊት|ከ\d+ ደቂቃ በፊት|በ\d+ ቀን ውስጥ|በ\d+ ወር ውስጥ|ከ\d+ ቀን በፊት|ከ\d+ ወር በፊት)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(የሚቀጥለው ሳምንት|የሚቀጥለው ዓመት|ባለፈው ሳምንት|የሚቀጥለው ወር|ያለፈው ሳምንት|በዚህ ሣምንት|በዚህ ሳምንት|ያለፈው ዓመት|በዚህ ዓመት|ያለፈው ወር|ሴፕቴምበር|በዚህ ወር|ይህ ሰዓት|ይህ ደቂቃ|ትላንትና|ኖቬምበር|ኦክቶበር|ዲሴምበር|ጃንዩወሪ|ፌብሩወሪ|ማክሰኞ|ሰከንድ|ሳምንት|ትናንት|ኤፕሪል|ኦገስት|ከሰዓት|gmt|utc|ሐሙስ|ማርች|ማክሰ|ረቡዕ|ሰዓት|ሴፕቴ|ቅዳሜ|ኖቬም|አሁን|ኤፕሪ|እሑድ|ኦክቶ|ኦገስ|ዓመት|ዓርብ|ደቂቃ|ዲሴም|ጁላይ|ጃንዩ|ጥዋት|ፌብሩ|\+|\.|\[|\]|\||am|pm|ሜይ|ሰኞ|ቀን|ነገ|ወር|ዛሬ|ጁን| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
