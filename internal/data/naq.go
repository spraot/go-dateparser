// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var naq_Locale = merge(nil, LocaleData{
	Name:      "naq",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"taraǀkhuumuǁkhab": "september",
		"aoǁkhuumuǁkhab":   "august",
		"dondertaxtsees":   "thursday",
		"satertaxtsees":    "saturday",
		"hoasoreǁkhab":     "december",
		"fraitaxtsees":     "friday",
		"denstaxtsees":     "tuesday",
		"wunstaxtsees":     "wednesday",
		"mantaxtsees":      "monday",
		"sontaxtsees":      "sunday",
		"ǃhoaǂkhaib":       "april",
		"ǃkhanǀgoab":       "february",
		"ǀkhuuǁkhab":       "march",
		"ǂnuǁnaiseb":       "october",
		"ǂkhoesaob":        "july",
		"ǃkhaitsab":        "may",
		"ǀhooǂgaeb":        "november",
		"gamaǀaeb":         "june",
		"ǁgoagas":          "am",
		"ǃkhanni":          "january",
		"wekheb":           "week",
		"tsees":            "day",
		"ǁkhab":            "month",
		"ǃuias":            "pm",
		"ǀgaub":            "second",
		"kurib":            "year",
		"iiri":             "hour",
		"haib":             "minute",
		"apr":              "april",
		"aug":              "august",
		"dec":              "december",
		"feb":              "february",
		"gmt":              "gmt",
		"jan":              "january",
		"jul":              "july",
		"jun":              "june",
		"mar":              "march",
		"may":              "may",
		"nov":              "november",
		"oct":              "october",
		"sat":              "saturday",
		"sep":              "september",
		"son":              "sunday",
		"utc":              "utc",
		"am":               "am",
		"fr":               "friday",
		"ma":               "monday",
		"pm":               "pm",
		"do":               "thursday",
		"de":               "tuesday",
		"wu":               "wednesday",
		"'":                "",
		",":                "",
		";":                "",
		"@":                "",
		"[":                "",
		"]":                "",
		"|":                "",
		" ":                " ",
		"+":                "+",
		"-":                "-",
		".":                ".",
		"/":                "/",
		":":                ":",
		"z":                "z",
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
		"neetsee":     "0 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(taraǀkhuumuǁkhab|aoǁkhuumuǁkhab|dondertaxtsees|satertaxtsees|denstaxtsees|fraitaxtsees|hoasoreǁkhab|wunstaxtsees|mantaxtsees|sontaxtsees|this minute|last month|next month|this month|ǀkhuuǁkhab|ǂnuǁnaiseb|ǃhoaǂkhaib|ǃkhanǀgoab|last week|last year|next week|next year|this hour|this week|this year|yesterday|ǀhooǂgaeb|ǂkhoesaob|ǃkhaitsab|gamaǀaeb|tomorrow|neetsee|ǁgoagas|ǃkhanni|wekheb|kurib|tsees|ǀgaub|ǁkhab|ǃuias|haib|iiri|apr|aug|dec|feb|gmt|jan|jul|jun|mar|may|nov|now|oct|sat|sep|son|utc|\+|\.|\[|\]|\||am|de|do|fr|ma|pm|wu| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
