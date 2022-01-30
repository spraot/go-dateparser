// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var shi_Latn_Locale = merge(nil, LocaleData{
	Name:      "shi-Latn",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefghijklnorstuwxyzɣḍṛṣ`),
	Translations: map[string]string{
		"dujanbir": "december",
		"nuwanbir": "november",
		"tadggwat": "pm",
		"cutanbir": "september",
		"asimwas":  "friday",
		"tasragt":  "hour",
		"tusdidt":  "minute",
		"asidyas":  "saturday",
		"imalass":  "week",
		"asggwas":  "year",
		"tifawt":   "am",
		"innayr":   "january",
		"yulyuz":   "july",
		"tasint":   "second",
		"asamas":   "sunday",
		"asinas":   "tuesday",
		"ibrir":    "april",
		"brayr":    "february",
		"yunyu":    "june",
		"mayyu":    "may",
		"aynas":    "monday",
		"ayyur":    "month",
		"ktubr":    "october",
		"akwas":    "thursday",
		"akras":    "wednesday",
		"ɣuct":     "august",
		"asim":     "friday",
		"mars":     "march",
		"asid":     "saturday",
		"ibr":      "april",
		"ɣuc":      "august",
		"ass":      "day",
		"duj":      "december",
		"bra":      "february",
		"gmt":      "gmt",
		"inn":      "january",
		"yul":      "july",
		"yun":      "june",
		"mar":      "march",
		"may":      "may",
		"ayn":      "monday",
		"nuw":      "november",
		"ktu":      "october",
		"cut":      "september",
		"asa":      "sunday",
		"akw":      "thursday",
		"asi":      "tuesday",
		"utc":      "utc",
		"akr":      "wednesday",
		"am":       "am",
		"pm":       "pm",
		"'":        "",
		",":        "",
		";":        "",
		"@":        "",
		"[":        "",
		"]":        "",
		"|":        "",
		" ":        " ",
		"+":        "+",
		"-":        "-",
		".":        ".",
		"/":        "/",
		":":        ":",
		"z":        "z",
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
		"idlli":       "1 day ago",
		"askka":       "in 1 day",
		"assa":        "0 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(this minute|last month|next month|this month|last week|last year|next week|next year|this hour|this week|this year|cutanbir|dujanbir|nuwanbir|tadggwat|asggwas|asidyas|asimwas|imalass|tasragt|tusdidt|asamas|asinas|innayr|tasint|tifawt|yulyuz|akras|akwas|askka|aynas|ayyur|brayr|ibrir|idlli|ktubr|mayyu|yunyu|asid|asim|assa|mars|ɣuct|akr|akw|asa|asi|ass|ayn|bra|cut|duj|gmt|ibr|inn|ktu|mar|may|now|nuw|utc|yul|yun|ɣuc|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
