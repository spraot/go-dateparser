// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ti_Locale = merge(nil, LocaleData{
	Name:      "ti",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"ንጉሆ ሰዓተ": "am",
		"ድሕር ሰዓት": "pm",
		"minute":  "minute",
		"second":  "second",
		"month":   "month",
		"መስከረም":   "september",
		"ሚያዝያ":    "april",
		"ታሕሳስ":    "december",
		"ለካቲት":    "february",
		"hour":    "hour",
		"መጋቢት":    "march",
		"ግንቦት":    "may",
		"ጥቅምቲ":    "october",
		"ሰንበት":    "sunday",
		"week":    "week",
		"year":    "year",
		"ነሓሰ":     "august",
		"day":     "day",
		"ዓርቢ":     "friday",
		"gmt":     "gmt",
		"ሓምለ":     "july",
		"ሰኑይ":     "monday",
		"ሕዳር":     "november",
		"ቀዳም":     "saturday",
		"ሓሙስ":     "thursday",
		"ኃሙስ":     "thursday",
		"ሠሉስ":     "tuesday",
		"ሰሉስ":     "tuesday",
		"utc":     "utc",
		"ረቡዕ":     "wednesday",
		"am":      "am",
		"ሚያ":      "april",
		"ነሓ":      "august",
		"ታሕ":      "december",
		"ለካ":      "february",
		"ዓር":      "friday",
		"ጥሪ":      "january",
		"ሓም":      "july",
		"ሰነ":      "june",
		"መጋ":      "march",
		"ግን":      "may",
		"ሰኑ":      "monday",
		"ሕዳ":      "november",
		"ጥቅ":      "october",
		"pm":      "pm",
		"ቀዳ":      "saturday",
		"መስ":      "september",
		"ሰን":      "sunday",
		"ሓሙ":      "thursday",
		"ሰሉ":      "tuesday",
		"ረቡ":      "wednesday",
		"'":       "",
		",":       "",
		";":       "",
		"@":       "",
		"[":       "",
		"]":       "",
		"|":       "",
		" ":       " ",
		"+":       "+",
		"-":       "-",
		".":       ".",
		"/":       "/",
		":":       ":",
		"z":       "z",
	},
	RelativeType: map[string]string{
		"this minute": "0 minute ago",
		"this month":  "0 month ago",
		"last month":  "1 month ago",
		"next month":  "in 1 month",
		"this hour":   "0 hour ago",
		"this week":   "0 week ago",
		"this year":   "0 year ago",
		"yesterday":   "1 day ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"tomorrow":    "in 1 day",
		"today":       "0 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|yesterday|tomorrow|ንጉሆ ሰዓተ|ድሕር ሰዓት|minute|second|month|today|መስከረም|hour|week|year|ለካቲት|መጋቢት|ሚያዝያ|ሰንበት|ታሕሳስ|ግንቦት|ጥቅምቲ|day|gmt|now|utc|ሓሙስ|ሓምለ|ሕዳር|ሠሉስ|ረቡዕ|ሰሉስ|ሰኑይ|ቀዳም|ኃሙስ|ነሓሰ|ዓርቢ|\+|\.|\[|\]|\||am|pm|ለካ|ሓሙ|ሓም|ሕዳ|መስ|መጋ|ሚያ|ረቡ|ሰሉ|ሰነ|ሰኑ|ሰን|ቀዳ|ታሕ|ነሓ|ዓር|ግን|ጥሪ|ጥቅ| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ti_ER_Locale = merge(&ti_Locale, LocaleData{
	Name:      "ti-ER",
	DateOrder: "DMY",
})
