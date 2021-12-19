// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ksh_Locale = merge(nil, LocaleData{
	Name:      "ksh",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"uhr vormiddaachs": "am",
		"uhr nommendaachs": "pm",
		"dunnersdaach":     "thursday",
		"vormeddaach":      "am",
		"nommendaach":      "pm",
		"dinnsdaach":       "tuesday",
		"friidaach":        "friday",
		"mohndaach":        "monday",
		"samsdaach":        "saturday",
		"septamber":        "september",
		"sunndaach":        "sunday",
		"dezamber":         "december",
		"novamber":         "november",
		"oktohber":         "october",
		"fabrowa":          "february",
		"schtund":          "hour",
		"jannewa":          "january",
		"metwoch":          "wednesday",
		"aprell":           "april",
		"menutt":           "minute",
		"sekond":           "second",
		"oujoß":            "august",
		"daach":            "day",
		"juuli":            "july",
		"juuni":            "june",
		"mohnd":            "month",
		"maaz":             "march",
		"woch":             "week",
		"johr":             "year",
		"apr":              "april",
		"ouj":              "august",
		"dez":              "december",
		"fab":              "february",
		"gmt":              "gmt",
		"std":              "hour",
		"jan":              "january",
		"jul":              "july",
		"jun":              "june",
		"maz":              "march",
		"mai":              "may",
		"min":              "minute",
		"nov":              "november",
		"okt":              "october",
		"sek":              "second",
		"sap":              "september",
		"utc":              "utc",
		"am":               "am",
		"vm":               "am",
		"fr":               "friday",
		"mo":               "monday",
		"nm":               "pm",
		"pm":               "pm",
		"sa":               "saturday",
		"su":               "sunday",
		"du":               "thursday",
		"di":               "tuesday",
		"me":               "wednesday",
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
		"d":                "day",
		"s":                "hour",
		"m":                "month",
		"w":                "week",
		"j":                "year",
		"z":                "z",
	},
	RelativeType: map[string]string{
		"nachste mohnd": "in 1 month",
		"nachste woche": "in 1 week",
		"latzde mohnd":  "1 month ago",
		"this minute":   "0 minute ago",
		"diese mohnd":   "0 month ago",
		"this hour":     "0 hour ago",
		"diß johr":      "0 year ago",
		"laz woch":      "1 week ago",
		"laz johr":      "1 year ago",
		"nax johr":      "in 1 year",
		"di woch":       "0 week ago",
		"jestere":       "1 day ago",
		"morje":         "in 1 day",
		"huck":          "0 day ago",
		"now":           "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)vor (\d+) johre`), "$1 year ago"},
		{regexp.MustCompile(`(?i)vor (\d+) johr`), "$1 year ago"},
		{regexp.MustCompile(`(?i)en (\d+) johre`), "in $1 year"},
		{regexp.MustCompile(`(?i)en (\d+) johr`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(vor \d+ johre|en \d+ johre|vor \d+ johr|en \d+ johr)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(vor \d+ johre|en \d+ johre|vor \d+ johr|en \d+ johr)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(uhr nommendaachs|uhr vormiddaachs|nachste mohnd|nachste woche|dunnersdaach|latzde mohnd|diese mohnd|nommendaach|this minute|vormeddaach|dinnsdaach|friidaach|mohndaach|samsdaach|septamber|sunndaach|this hour|dezamber|diß johr|laz johr|laz woch|nax johr|novamber|oktohber|di woch|fabrowa|jannewa|jestere|metwoch|schtund|aprell|menutt|sekond|daach|juuli|juuni|mohnd|morje|oujoß|huck|johr|maaz|woch|apr|dez|fab|gmt|jan|jul|jun|mai|maz|min|nov|now|okt|ouj|sap|sek|std|utc|\+|\.|\[|\]|\||am|di|du|fr|me|mo|nm|pm|sa|su|vm| |'|,|-|/|:|;|@|d|j|m|s|w|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
