// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var smn_Locale = merge(nil, LocaleData{
	Name:      "smn",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"vastuppeeivi": "friday",
		"uđđaivemaanu": "january",
		"roovvadmaanu": "october",
		"cuaŋuimaanu":  "april",
		"juovlamaanu":  "december",
		"vastuppeivi":  "friday",
		"syeinimaanu":  "july",
		"njuhcamaanu":  "march",
		"skammamaanu":  "november",
		"porgemaanu":   "august",
		"kuovamaanu":   "february",
		"vyesimaanu":   "may",
		"vuossaarga":   "monday",
		"cohcamaanu":   "september",
		"pasepeeivi":   "sunday",
		"majebaarga":   "tuesday",
		"kesimaanu":    "june",
		"vuossarga":    "monday",
		"pasepeivi":    "sunday",
		"tuorastah":    "thursday",
		"tuorastuv":    "thursday",
		"majebarga":    "tuesday",
		"lavurdah":     "saturday",
		"lavurduv":     "saturday",
		"koskokko":     "wednesday",
		"roovvad":      "october",
		"koskoho":      "wednesday",
		"cuaŋui":       "april",
		"juovla":       "december",
		"syeini":       "july",
		"njuhca":       "march",
		"minute":       "minute",
		"skamma":       "november",
		"second":       "second",
		"porge":        "august",
		"kuova":        "february",
		"vyesi":        "may",
		"month":        "month",
		"cohca":        "september",
		"hour":         "hour",
		"uđiv":         "january",
		"kesi":         "june",
		"week":         "week",
		"year":         "year",
		"day":          "day",
		"vas":          "friday",
		"gmt":          "gmt",
		"vuo":          "monday",
		"lav":          "saturday",
		"pas":          "sunday",
		"tuo":          "thursday",
		"maj":          "tuesday",
		"utc":          "utc",
		"kos":          "wednesday",
		"am":           "am",
		"ip":           "am",
		"ep":           "pm",
		"pm":           "pm",
		"'":            "",
		",":            "",
		";":            "",
		"@":            "",
		"[":            "",
		"]":            "",
		"|":            "",
		" ":            " ",
		"+":            "+",
		"-":            "-",
		".":            ".",
		"/":            "/",
		":":            ":",
		"z":            "z",
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
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(roovvadmaanu|uđđaivemaanu|vastuppeeivi|cuaŋuimaanu|juovlamaanu|njuhcamaanu|skammamaanu|syeinimaanu|this minute|vastuppeivi|cohcamaanu|kuovamaanu|last month|majebaarga|next month|pasepeeivi|porgemaanu|this month|vuossaarga|vyesimaanu|kesimaanu|last week|last year|majebarga|next week|next year|pasepeivi|this hour|this week|this year|tuorastah|tuorastuv|vuossarga|yesterday|koskokko|lavurdah|lavurduv|tomorrow|koskoho|roovvad|cuaŋui|juovla|minute|njuhca|second|skamma|syeini|cohca|kuova|month|porge|today|vyesi|hour|kesi|uđiv|week|year|day|gmt|kos|lav|maj|now|pas|tuo|utc|vas|vuo|\+|\.|\[|\]|\||am|ep|ip|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
