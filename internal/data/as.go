// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var as_Locale = merge(nil, LocaleData{
	Name:      "as",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"বহষপতিবাৰ": "thursday",
		"ফেবৰৱাৰী":  "february",
		"ছেপতেমবৰ":  "september",
		"ডিচেমবৰ":   "december",
		"জানৱাৰী":   "january",
		"মঙগলবাৰ":   "tuesday",
		"পৰবাহণ":    "am",
		"শকৰবাৰ":    "friday",
		"সোমবাৰ":    "monday",
		"নৱেমবৰ":    "november",
		"অকটোবৰ":    "october",
		"অপৰাহণ":    "pm",
		"শনিবাৰ":    "saturday",
		"ছেকেণড":    "second",
		"দেওবাৰ":    "sunday",
		"বহষপতি":    "thursday",
		"এপৰিল":     "april",
		"মিনিট":     "minute",
		"বধবাৰ":     "wednesday",
		"সপতাহ":     "week",
		"আগষট":      "august",
		"ডিসে":      "december",
		"ফেবৰ":      "february",
		"ঘণটা":      "hour",
		"জলাই":      "july",
		"মাৰচ":      "march",
		"অকটো":      "october",
		"সেপট":      "september",
		"মঙগল":      "tuesday",
		"দিন":       "day",
		"শকৰ":       "friday",
		"gmt":       "gmt",
		"জান":       "january",
		"সোম":       "monday",
		"মাহ":       "month",
		"নভে":       "november",
		"শনি":       "saturday",
		"ৰবি":       "sunday",
		"utc":       "utc",
		"বছৰ":       "year",
		"am":        "am",
		"আগ":        "august",
		"জন":        "june",
		"মে":        "may",
		"pm":        "pm",
		"বধ":        "wednesday",
		"'":         "",
		",":         "",
		";":         "",
		"@":         "",
		"[":         "",
		"]":         "",
		"|":         "",
		" ":         " ",
		"+":         "+",
		"-":         "-",
		".":         ".",
		"/":         "/",
		":":         ":",
		"z":         "z",
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
		"কাইলৈ":       "in 1 day",
		"কালি":        "1 day ago",
		"আজি":         "0 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|বহষপতিবাৰ|ছেপতেমবৰ|ফেবৰৱাৰী|জানৱাৰী|ডিচেমবৰ|মঙগলবাৰ|অকটোবৰ|অপৰাহণ|ছেকেণড|দেওবাৰ|নৱেমবৰ|পৰবাহণ|বহষপতি|শকৰবাৰ|শনিবাৰ|সোমবাৰ|এপৰিল|কাইলৈ|বধবাৰ|মিনিট|সপতাহ|অকটো|আগষট|কালি|ঘণটা|জলাই|ডিসে|ফেবৰ|মঙগল|মাৰচ|সেপট|gmt|now|utc|আজি|জান|দিন|নভে|বছৰ|মাহ|শকৰ|শনি|সোম|ৰবি|\+|\.|\[|\]|\||am|pm|আগ|জন|বধ|মে| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
