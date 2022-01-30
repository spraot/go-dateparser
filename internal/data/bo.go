// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var bo_Locale = merge(nil, LocaleData{
	Name:      "bo",
	DateOrder: "YMD",
	Charset:   []rune(`ceghiklnorstuwxyz་།༠༡༢༣༤༥༦༧༨༩ཁགངཅཆཉདནཔཕབམཚཞཟའརལསིེོུྐྒྔྤྱྲླྷ`),
	Translations: map[string]string{
		"ཟ་བ་བཅ་གཉས་པ་": "december",
		"ཟ་བ་བཅ་གཅག་པ་": "november",
		"ཟ་བ་བཅ་གཉས་པ":  "december",
		"ཟ་བ་བཅ་གཅག་པ":  "november",
		"གཟའ་མག་དམར་":   "tuesday",
		"ཟ་བ་བརད་པ་":    "august",
		"ཟ་བ་གཉས་པ་":    "february",
		"གཟའ་པ་སངས་":    "friday",
		"ཟ་བ་བདན་པ་":    "july",
		"ཟ་བ་གསམ་པ་":    "march",
		"ཟ་བ་བཞ་པ་":     "april",
		"ཟ་བ་བརད་པ":     "august",
		"ཟ་བ་གཉས་པ":     "february",
		"ཟ་བ་དང་པ་":     "january",
		"ཟ་བ་བདན་པ":     "july",
		"ཟ་བ་གསམ་པ":     "march",
		"ཟ་བ་བཅ་པ་":     "october",
		"གཟའ་སན་པ་":     "saturday",
		"ཟ་བ་དག་པ་":     "september",
		"གཟའ་ཕར་བ་":     "thursday",
		"གཟའ་ལག་པ་":     "wednesday",
		"ཟ་བ་བཞ་པ":      "april",
		"ཟ་བ་དང་པ":      "january",
		"ཟ་བ་ལ་པ་":      "may",
		"གཟའ་ཟ་བ་":      "monday",
		"ཟ་བ་བཅ་པ":      "october",
		"ཟ་བ་དག་པ":      "september",
		"གཟའ་ཉ་མ་":      "sunday",
		"ཟ་བ་ལ་པ":       "may",
		"མག་དམར་":       "tuesday",
		"པ་སངས་":        "friday",
		"ཆ་ཚད་":         "hour",
		"སར་མ།":         "minute",
		"སན་པ་":         "saturday",
		"སར་ཆ།":         "second",
		"ཕར་བ་":         "thursday",
		"ལག་པ་":         "wednesday",
		"ས་ད་":          "am",
		"ཟ་༡༢":          "december",
		"ཟ་བ་":          "month",
		"ཟ་༡༡":          "november",
		"ཟ་༡༠":          "october",
		"ཕ་ད་":          "pm",
		"ཉ་མ་":          "sunday",
		"week":          "week",
		"ཟ་༤":           "april",
		"ཟ་༨":           "august",
		"ཉན།":           "day",
		"ཟ་༢":           "february",
		"gmt":           "gmt",
		"ཟ་༡":           "january",
		"ཟ་༧":           "july",
		"ཟ་༦":           "june",
		"ཟ་༣":           "march",
		"ཟ་༥":           "may",
		"ཟ་༩":           "september",
		"utc":           "utc",
		"am":            "am",
		"pm":            "pm",
		"ལ།":            "year",
		"'":             "",
		",":             "",
		";":             "",
		"@":             "",
		"[":             "",
		"]":             "",
		"|":             "",
		" ":             " ",
		"+":             "+",
		"-":             "-",
		".":             ".",
		"/":             "/",
		":":             ":",
		"z":             "z",
	},
	RelativeType: map[string]string{
		"this minute": "0 minute ago",
		"this month":  "0 month ago",
		"last month":  "1 month ago",
		"next month":  "in 1 month",
		"this hour":   "0 hour ago",
		"this week":   "0 week ago",
		"this year":   "0 year ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"སང་ཉན་":      "in 1 day",
		"ད་རང་":       "0 day ago",
		"ཁས་ས་":       "1 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(ཟ་བ་བཅ་གཅག་པ་|ཟ་བ་བཅ་གཉས་པ་|ཟ་བ་བཅ་གཅག་པ|ཟ་བ་བཅ་གཉས་པ|this minute|གཟའ་མག་དམར་|last month|next month|this month|གཟའ་པ་སངས་|ཟ་བ་གཉས་པ་|ཟ་བ་གསམ་པ་|ཟ་བ་བདན་པ་|ཟ་བ་བརད་པ་|last week|last year|next week|next year|this hour|this week|this year|གཟའ་ཕར་བ་|གཟའ་ལག་པ་|གཟའ་སན་པ་|ཟ་བ་གཉས་པ|ཟ་བ་གསམ་པ|ཟ་བ་དག་པ་|ཟ་བ་དང་པ་|ཟ་བ་བཅ་པ་|ཟ་བ་བདན་པ|ཟ་བ་བཞ་པ་|ཟ་བ་བརད་པ|གཟའ་ཉ་མ་|གཟའ་ཟ་བ་|ཟ་བ་དག་པ|ཟ་བ་དང་པ|ཟ་བ་བཅ་པ|ཟ་བ་བཞ་པ|ཟ་བ་ལ་པ་|མག་དམར་|ཟ་བ་ལ་པ|པ་སངས་|སང་ཉན་|ཁས་ས་|ཆ་ཚད་|ད་རང་|ཕར་བ་|ལག་པ་|སན་པ་|སར་ཆ།|སར་མ།|week|ཉ་མ་|ཕ་ད་|ཟ་༡༠|ཟ་༡༡|ཟ་༡༢|ཟ་བ་|ས་ད་|gmt|now|utc|ཉན།|ཟ་༡|ཟ་༢|ཟ་༣|ཟ་༤|ཟ་༥|ཟ་༦|ཟ་༧|ཟ་༨|ཟ་༩|\+|\.|\[|\]|\||am|pm|ལ།| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var bo_IN_Locale = merge(&bo_Locale, LocaleData{
	Name:         "bo-IN",
	DateOrder:    "YMD",
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(ཟ་བ་བཅ་གཅག་པ་|ཟ་བ་བཅ་གཉས་པ་|ཟ་བ་བཅ་གཅག་པ|ཟ་བ་བཅ་གཉས་པ|this minute|གཟའ་མག་དམར་|last month|next month|this month|གཟའ་པ་སངས་|ཟ་བ་གཉས་པ་|ཟ་བ་གསམ་པ་|ཟ་བ་བདན་པ་|ཟ་བ་བརད་པ་|last week|last year|next week|next year|this hour|this week|this year|གཟའ་ཕར་བ་|གཟའ་ལག་པ་|གཟའ་སན་པ་|ཟ་བ་གཉས་པ|ཟ་བ་གསམ་པ|ཟ་བ་དག་པ་|ཟ་བ་དང་པ་|ཟ་བ་བཅ་པ་|ཟ་བ་བདན་པ|ཟ་བ་བཞ་པ་|ཟ་བ་བརད་པ|གཟའ་ཉ་མ་|གཟའ་ཟ་བ་|ཟ་བ་དག་པ|ཟ་བ་དང་པ|ཟ་བ་བཅ་པ|ཟ་བ་བཞ་པ|ཟ་བ་ལ་པ་|མག་དམར་|ཟ་བ་ལ་པ|པ་སངས་|སང་ཉན་|ཁས་ས་|ཆ་ཚད་|ད་རང་|ཕར་བ་|ལག་པ་|སན་པ་|སར་ཆ།|སར་མ།|week|ཉ་མ་|ཕ་ད་|ཟ་༡༠|ཟ་༡༡|ཟ་༡༢|ཟ་བ་|ས་ད་|gmt|now|utc|ཉན།|ཟ་༡|ཟ་༢|ཟ་༣|ཟ་༤|ཟ་༥|ཟ་༦|ཟ་༧|ཟ་༨|ཟ་༩|\+|\.|\[|\]|\||am|pm|ལ།| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
