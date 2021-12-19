// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var dav_Locale = merge(nil, LocaleData{
	Name:      "dav",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"mori ghwa ikumi na imweri": "november",
		"mori ghwa ikumi na iwi":    "december",
		"mori ghwa karandadu":       "june",
		"mori ghwa wunyanya":        "august",
		"mori ghwa mfungade":        "july",
		"mori ghwa imbiri":          "january",
		"mori ghwa kadadu":          "march",
		"mori ghwa kasanu":          "may",
		"kuramuka jimweri":          "monday",
		"mori ghwa ikenda":          "september",
		"kuramuka kasanu":           "friday",
		"mori ghwa ikumi":           "october",
		"kuramuka kadadu":           "wednesday",
		"mori ghwa kana":            "april",
		"mori ghwa kawi":            "february",
		"ituku ja jumwa":            "sunday",
		"kuramuka kana":             "thursday",
		"kuramuka kawi":             "tuesday",
		"kifula nguwo":              "saturday",
		"luma lwa k":                "am",
		"luma lwa p":                "pm",
		"sekunde":                   "second",
		"dakika":                    "minute",
		"ituku":                     "day",
		"mwaka":                     "year",
		"mori":                      "month",
		"juma":                      "week",
		"kan":                       "april",
		"wun":                       "august",
		"iwi":                       "december",
		"kaw":                       "february",
		"gmt":                       "gmt",
		"saa":                       "hour",
		"imb":                       "january",
		"mfu":                       "july",
		"kar":                       "june",
		"kad":                       "march",
		"kas":                       "may",
		"jim":                       "monday",
		"imw":                       "november",
		"iku":                       "october",
		"ngu":                       "saturday",
		"ike":                       "september",
		"jum":                       "sunday",
		"utc":                       "utc",
		"am":                        "am",
		"pm":                        "pm",
		"'":                         "",
		",":                         "",
		";":                         "",
		"@":                         "",
		"[":                         "",
		"]":                         "",
		"|":                         "",
		" ":                         " ",
		"+":                         "+",
		"-":                         "-",
		".":                         ".",
		"/":                         "/",
		":":                         ":",
		"z":                         "z",
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
		"idime":       "0 day ago",
		"kesho":       "in 1 day",
		"iguo":        "1 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(mori ghwa ikumi na imweri|mori ghwa ikumi na iwi|mori ghwa karandadu|mori ghwa mfungade|mori ghwa wunyanya|kuramuka jimweri|mori ghwa ikenda|mori ghwa imbiri|mori ghwa kadadu|mori ghwa kasanu|kuramuka kadadu|kuramuka kasanu|mori ghwa ikumi|ituku ja jumwa|mori ghwa kana|mori ghwa kawi|kuramuka kana|kuramuka kawi|kifula nguwo|this minute|last month|luma lwa k|luma lwa p|next month|this month|last week|last year|next week|next year|this hour|this week|this year|sekunde|dakika|idime|ituku|kesho|mwaka|iguo|juma|mori|gmt|ike|iku|imb|imw|iwi|jim|jum|kad|kan|kar|kas|kaw|mfu|ngu|now|saa|utc|wun|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
