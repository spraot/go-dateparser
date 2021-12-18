// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var mfe_Locale = merge(nil, LocaleData{
	Name:      "mfe",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
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
		"vandredi":    "friday",
		"merkredi":    "wednesday",
		"fevriye":     "february",
		"zanvie":      "january",
		"segonn":      "second",
		"septam":      "september",
		"dimans":      "sunday",
		"semenn":      "week",
		"zordi":       "0 day ago",
		"avril":       "april",
		"desam":       "december",
		"demin":       "in 1 day",
		"zilye":       "july",
		"minit":       "minute",
		"lindi":       "monday",
		"novam":       "november",
		"oktob":       "october",
		"samdi":       "saturday",
		"mardi":       "tuesday",
		"zour":        "day",
		"mars":        "march",
		"zedi":        "thursday",
		"lane":        "year",
		"now":         "0 second ago",
		"yer":         "1 day ago",
		"avr":         "april",
		"out":         "august",
		"des":         "december",
		"fev":         "february",
		"van":         "friday",
		"gmt":         "gmt",
		"ler":         "hour",
		"zan":         "january",
		"zil":         "july",
		"zin":         "june",
		"lin":         "monday",
		"mwa":         "month",
		"nov":         "november",
		"okt":         "october",
		"sam":         "saturday",
		"sep":         "september",
		"dim":         "sunday",
		"mar":         "tuesday",
		"utc":         "utc",
		"mer":         "wednesday",
		"am":          "am",
		"me":          "may",
		"pm":          "pm",
		"ze":          "thursday",
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
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|merkredi|vandredi|fevriye|dimans|segonn|semenn|septam|zanvie|avril|demin|desam|lindi|mardi|minit|novam|oktob|samdi|zilye|zordi|lane|mars|zedi|zour|avr|des|dim|fev|gmt|ler|lin|mar|mer|mwa|nov|now|okt|out|sam|sep|utc|van|yer|zan|zil|zin|\+|\.|\[|\]|\||am|me|pm|ze| |'|,|-|/|:|;|@|z)((?:\z|\W|_|\d).*)$`),
})
