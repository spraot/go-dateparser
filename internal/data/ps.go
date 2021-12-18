// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ps_Locale = merge(nil, LocaleData{
	Name:      "ps",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
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
		"چهارشنبه":    "wednesday",
		"پنجشنبه":     "thursday",
		"سه‌شنبه":     "tuesday",
		"فبروري":      "february",
		"minute":      "minute",
		"دوشنبه":      "monday",
		"اکتوبر":      "october",
		"second":      "second",
		"سپتمبر":      "september",
		"یکشنبه":      "sunday",
		"today":       "0 day ago",
		"اپریل":       "april",
		"دسمبر":       "december",
		"جنوري":       "january",
		"جولای":       "july",
		"month":       "month",
		"نومبر":       "november",
		"اګست":        "august",
		"جمعه":        "friday",
		"hour":        "hour",
		"مارچ":        "march",
		"شنبه":        "saturday",
		"week":        "week",
		"year":        "year",
		"now":         "0 second ago",
		"day":         "day",
		"gmt":         "gmt",
		"جون":         "june",
		"utc":         "utc",
		"am":          "am",
		"غم":          "am",
		"مۍ":          "may",
		"pm":          "pm",
		"غو":          "pm",
		"'":           "",
		",":           "",
		";":           "",
		"@":           "",
		"[":           "",
		"]":           "",
		"|":           "",
		" ":           " ",
		"+":           "+",
		"-":           "-",
		".":           ".",
		"/":           "/",
		":":           ":",
		"z":           "z",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|yesterday|tomorrow|چهارشنبه|سه‌شنبه|پنجشنبه|minute|second|اکتوبر|دوشنبه|سپتمبر|فبروري|یکشنبه|month|today|اپریل|جنوري|جولای|دسمبر|نومبر|hour|week|year|اګست|جمعه|شنبه|مارچ|day|gmt|now|utc|جون|\+|\.|\[|\]|\||am|pm|غم|غو|مۍ| |'|,|-|/|:|;|@|z)((?:\z|\W|_|\d).*)$`),
})
