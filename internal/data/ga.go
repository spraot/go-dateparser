// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ga_Locale = merge(nil, LocaleData{
	Name:      "ga",
	DateOrder: "DMY",
	Charset:   []rune(`bcdefghilnorstuzáéíóú`),
	Translations: map[string]string{
		"deireadh fomhair": "october",
		"mean fomhair":     "september",
		"de domhnaigh":     "sunday",
		"de sathairn":      "saturday",
		"de ceadaoin":      "wednesday",
		"de haoine":        "friday",
		"meitheamh":        "june",
		"bealtaine":        "may",
		"deardaoin":        "thursday",
		"seachtain":        "week",
		"de luain":         "monday",
		"de mairt":         "tuesday",
		"aibrean":          "april",
		"nollaig":          "december",
		"feabhra":          "february",
		"noimead":          "minute",
		"samhain":          "november",
		"soicind":          "second",
		"lunasa":           "august",
		"eanair":           "january",
		"bliain":           "year",
		"feabh":            "february",
		"aoine":            "friday",
		"meith":            "june",
		"marta":            "march",
		"dfomh":            "october",
		"mfomh":            "september",
		"mairt":            "tuesday",
		"noll":             "december",
		"uair":             "hour",
		"iuil":             "july",
		"beal":             "may",
		"noim":             "minute",
		"luan":             "monday",
		"samh":             "november",
		"sath":             "saturday",
		"soic":             "second",
		"domh":             "sunday",
		"dear":             "thursday",
		"cead":             "wednesday",
		"scht":             "week",
		"aib":              "april",
		"lun":              "august",
		"gmt":              "gmt",
		"ean":              "january",
		"utc":              "utc",
		"am":               "am",
		"la":               "day",
		"mi":               "month",
		"pm":               "pm",
		"bl":               "year",
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
		"u":                "hour",
		"n":                "minute",
		"s":                "second",
		"z":                "z",
	},
	RelativeType: map[string]string{
		"an tseachtain seo chugainn": "in 1 week",
		"an tseachtain seo caite":    "1 week ago",
		"an bhliain seo chugainn":    "in 1 year",
		"an tscht seo chugainn":      "in 1 week",
		"an mhi seo chugainn":        "in 1 month",
		"an bhl seo chugainn":        "in 1 year",
		"an tscht seo caite":         "1 week ago",
		"an tseachtain seo":          "0 week ago",
		"an mhi seo caite":           "1 month ago",
		"an noimead seo":             "0 minute ago",
		"an bhliain seo":             "0 year ago",
		"an tscht seo":               "0 week ago",
		"an uair seo":                "0 hour ago",
		"an mhi seo":                 "0 month ago",
		"an bhl seo":                 "0 year ago",
		"anuraidh":                   "1 year ago",
		"amarach":                    "in 1 day",
		"inniu":                      "0 day ago",
		"anois":                      "0 second ago",
		"inne":                       "1 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)i gceann (\d+) uair an chloig`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) uair an chloig o shin`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)i gceann (\d+) seachtain`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) seachtain o shin`), "$1 week ago"},
		{regexp.MustCompile(`(?i)i gceann (\d+) noimead`), "in $1 minute"},
		{regexp.MustCompile(`(?i)i gceann (\d+) soicind`), "in $1 second"},
		{regexp.MustCompile(`(?i)i gceann (\d+) bhliain`), "in $1 year"},
		{regexp.MustCompile(`(?i)i gceann (\d+) bliain`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) noimead o shin`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) soicind o shin`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) bhliain o shin`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) bliain o shin`), "$1 year ago"},
		{regexp.MustCompile(`(?i)i gceann (\d+) uair`), "in $1 hour"},
		{regexp.MustCompile(`(?i)i gceann (\d+) noim`), "in $1 minute"},
		{regexp.MustCompile(`(?i)i gceann (\d+) soic`), "in $1 second"},
		{regexp.MustCompile(`(?i)i gceann (\d+) scht`), "in $1 week"},
		{regexp.MustCompile(`(?i)i gceann (\d+) mhi`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) uair o shin`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) noim o shin`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) soic o shin`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) scht o shin`), "$1 week ago"},
		{regexp.MustCompile(`(?i)i gceann (\d+) la`), "in $1 day"},
		{regexp.MustCompile(`(?i)i gceann (\d+) mi`), "in $1 month"},
		{regexp.MustCompile(`(?i)i gceann (\d+) bl`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) mhi o shin`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) bhl o shin`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) la o shin`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) mi o shin`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) bl o shin`), "$1 year ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(i gceann \d+ uair an chloig|\d+ uair an chloig o shin|i gceann \d+ seachtain|\d+ seachtain o shin|i gceann \d+ bhliain|i gceann \d+ noimead|i gceann \d+ soicind|i gceann \d+ bliain|\d+ bhliain o shin|\d+ noimead o shin|\d+ soicind o shin|\d+ bliain o shin|i gceann \d+ noim|i gceann \d+ scht|i gceann \d+ soic|i gceann \d+ uair|i gceann \d+ mhi|\d+ noim o shin|\d+ scht o shin|\d+ soic o shin|\d+ uair o shin|i gceann \d+ bl|i gceann \d+ la|i gceann \d+ mi|\d+ bhl o shin|\d+ mhi o shin|\d+ bl o shin|\d+ la o shin|\d+ mi o shin)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(i gceann \d+ uair an chloig|\d+ uair an chloig o shin|i gceann \d+ seachtain|\d+ seachtain o shin|i gceann \d+ bhliain|i gceann \d+ noimead|i gceann \d+ soicind|i gceann \d+ bliain|\d+ bhliain o shin|\d+ noimead o shin|\d+ soicind o shin|\d+ bliain o shin|i gceann \d+ noim|i gceann \d+ scht|i gceann \d+ soic|i gceann \d+ uair|i gceann \d+ mhi|\d+ noim o shin|\d+ scht o shin|\d+ soic o shin|\d+ uair o shin|i gceann \d+ bl|i gceann \d+ la|i gceann \d+ mi|\d+ bhl o shin|\d+ mhi o shin|\d+ bl o shin|\d+ la o shin|\d+ mi o shin)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(an tseachtain seo chugainn|an bhliain seo chugainn|an tseachtain seo caite|an tscht seo chugainn|an bhl seo chugainn|an mhi seo chugainn|an tscht seo caite|an tseachtain seo|an mhi seo caite|deireadh fomhair|an bhliain seo|an noimead seo|an tscht seo|de domhnaigh|mean fomhair|an uair seo|de ceadaoin|de sathairn|an bhl seo|an mhi seo|bealtaine|de haoine|deardaoin|meitheamh|seachtain|anuraidh|de luain|de mairt|aibrean|amarach|feabhra|noimead|nollaig|samhain|soicind|bliain|eanair|lunasa|anois|aoine|dfomh|feabh|inniu|mairt|marta|meith|mfomh|beal|cead|dear|domh|inne|iuil|luan|noim|noll|samh|sath|scht|soic|uair|aib|ean|gmt|lun|utc|\+|\.|\[|\]|\||am|bl|la|mi|pm| |'|,|-|/|:|;|@|n|s|u|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
