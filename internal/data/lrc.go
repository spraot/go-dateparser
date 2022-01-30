// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var lrc_Locale = merge(nil, LocaleData{
	Name:      "lrc",
	DateOrder: "YMD",
	Charset:   []rune(`cdefghiklnorstuwxyzآأئاتثجدرزسشصفقلمنويٙپڤکگھیە`),
	Translations: map[string]string{
		"يوکتوڤر": "october",
		"سيپتامر": "september",
		"ديسامر":  "december",
		"فيڤریە":  "february",
		"جانڤیە":  "january",
		"نوڤامر":  "november",
		"اڤریل":   "april",
		"اگوست":   "august",
		"ديیقە":   "minute",
		"ثانیە":   "second",
		"ھافتە":   "week",
		"ساات":    "hour",
		"جولا":    "july",
		"جوان":    "june",
		"مارس":    "march",
		"روز":     "day",
		"fri":     "friday",
		"gmt":     "gmt",
		"ميی":     "may",
		"mon":     "monday",
		"sat":     "saturday",
		"sun":     "sunday",
		"thu":     "thursday",
		"tue":     "tuesday",
		"utc":     "utc",
		"wed":     "wednesday",
		"سال":     "year",
		"am":      "am",
		"ما":      "month",
		"pm":      "pm",
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
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"دیروز":       "1 day ago",
		"امرو":        "0 day ago",
		"شوصو":        "in 1 day",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|سيپتامر|يوکتوڤر|جانڤیە|ديسامر|فيڤریە|نوڤامر|اڤریل|اگوست|ثانیە|ديیقە|دیروز|ھافتە|امرو|جوان|جولا|ساات|شوصو|مارس|fri|gmt|mon|now|sat|sun|thu|tue|utc|wed|روز|سال|ميی|\+|\.|\[|\]|\||am|pm|ما| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var lrc_IQ_Locale = merge(&lrc_Locale, LocaleData{
	Name:         "lrc-IQ",
	DateOrder:    "YMD",
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|سيپتامر|يوکتوڤر|جانڤیە|ديسامر|فيڤریە|نوڤامر|اڤریل|اگوست|ثانیە|ديیقە|دیروز|ھافتە|امرو|جوان|جولا|ساات|شوصو|مارس|fri|gmt|mon|now|sat|sun|thu|tue|utc|wed|روز|سال|ميی|\+|\.|\[|\]|\||am|pm|ما| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
