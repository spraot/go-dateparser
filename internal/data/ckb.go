// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ckb_Locale = merge(nil, LocaleData{
	Name:      "ckb",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"کانوونی دووەم": "january",
		"کانونی یەکەم":  "december",
		"تشرینی دووەم":  "november",
		"تشرینی یەکەم":  "october",
		"پێنجشەممە":     "thursday",
		"چوارشەممە":     "wednesday",
		"حوزەیران":      "june",
		"دووشەممە":      "monday",
		"یەکشەممە":      "sunday",
		"يەیلوول":       "september",
		"سێشەممە":       "tuesday",
		"تەمووز":        "july",
		"minute":        "minute",
		"second":        "second",
		"نیسان":         "april",
		"شوبات":         "february",
		"ھەینی":         "friday",
		"يازار":         "march",
		"يایار":         "may",
		"month":         "month",
		"شەممە":         "saturday",
		"hour":          "hour",
		"week":          "week",
		"year":          "year",
		"ياب":           "august",
		"day":           "day",
		"gmt":           "gmt",
		"utc":           "utc",
		"am":            "am",
		"بن":            "am",
		"pm":            "pm",
		"دن":            "pm",
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
		"yesterday":   "1 day ago",
		"last week":   "1 week ago",
		"last year":   "1 year ago",
		"next week":   "in 1 week",
		"next year":   "in 1 year",
		"tomorrow":    "in 1 day",
		"today":       "0 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(کانوونی دووەم|تشرینی دووەم|تشرینی یەکەم|کانونی یەکەم|this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|yesterday|پێنجشەممە|چوارشەممە|tomorrow|حوزەیران|دووشەممە|یەکشەممە|سێشەممە|يەیلوول|minute|second|تەمووز|month|today|شوبات|شەممە|نیسان|يازار|يایار|ھەینی|hour|week|year|day|gmt|now|utc|ياب|\+|\.|\[|\]|\||am|pm|بن|دن| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ckb_IR_Locale = merge(&ckb_Locale, LocaleData{
	Name:         "ckb-IR",
	DateOrder:    "YMD",
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(کانوونی دووەم|تشرینی دووەم|تشرینی یەکەم|کانونی یەکەم|this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|yesterday|پێنجشەممە|چوارشەممە|tomorrow|حوزەیران|دووشەممە|یەکشەممە|سێشەممە|يەیلوول|minute|second|تەمووز|month|today|شوبات|شەممە|نیسان|يازار|يایار|ھەینی|hour|week|year|day|gmt|now|utc|ياب|\+|\.|\[|\]|\||am|pm|بن|دن| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
