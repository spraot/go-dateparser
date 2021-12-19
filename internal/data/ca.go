// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ca_Locale = merge(nil, LocaleData{
	Name:         "ca",
	DateOrder:    "DMY",
	SkipWords:    []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "de", "del", "i", "l'", "|"},
	PertainWords: []string{"de", "del"},
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)una(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)un(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
	},
	Translations: map[string]string{
		"de desembre": "december",
		"de novembre": "november",
		"de setembre": "september",
		"de febrer":   "february",
		"divendres":   "friday",
		"de juliol":   "july",
		"d'octubre":   "october",
		"desembre":    "december",
		"de gener":    "january",
		"novembre":    "november",
		"dissabte":    "saturday",
		"setembre":    "september",
		"diumenge":    "sunday",
		"dimecres":    "wednesday",
		"d'abril":     "april",
		"d'agost":     "august",
		"de febr":     "february",
		"de juny":     "june",
		"de marc":     "march",
		"de maig":     "may",
		"dilluns":     "monday",
		"octubre":     "october",
		"dimarts":     "tuesday",
		"setmana":     "week",
		"de des":      "december",
		"febrer":      "february",
		"de gen":      "january",
		"de jul":      "july",
		"juliol":      "july",
		"de nov":      "november",
		"de set":      "september",
		"dijous":      "thursday",
		"abril":       "april",
		"d'abr":       "april",
		"agost":       "august",
		"gener":       "january",
		"minut":       "minute",
		"d'oct":       "october",
		"segon":       "second",
		"d'ag":        "august",
		"febr":        "february",
		"hora":        "hour",
		"juny":        "june",
		"marc":        "march",
		"maig":        "may",
		"setm":        "week",
		"del":         "",
		"abr":         "april",
		"dia":         "day",
		"des":         "december",
		"gmt":         "gmt",
		"gen":         "january",
		"jul":         "july",
		"min":         "minute",
		"mes":         "month",
		"nov":         "november",
		"oct":         "october",
		"set":         "september",
		"utc":         "utc",
		"any":         "year",
		"de":          "",
		"l'":          "",
		"am":          "am",
		"ag":          "august",
		"dv":          "friday",
		"en":          "in",
		"dl":          "monday",
		"pm":          "pm",
		"ds":          "saturday",
		"dg":          "sunday",
		"dj":          "thursday",
		"dt":          "tuesday",
		"dc":          "wednesday",
		"'":           "",
		",":           "",
		";":           "",
		"@":           "",
		"[":           "",
		"]":           "",
		"i":           "",
		"|":           "",
		" ":           " ",
		"+":           "+",
		"-":           "-",
		".":           ".",
		"/":           "/",
		":":           ":",
		"h":           "hour",
		"s":           "second",
		"z":           "z",
	},
	RelativeType: map[string]string{
		"la setmana passada": "1 week ago",
		"la propera setmana": "in 1 week",
		"la proxima setmana": "in 1 week",
		"la setmana que ve":  "in 1 week",
		"la setmana vinent":  "in 1 week",
		"aquesta setmana":    "0 week ago",
		"la setm passada":    "1 week ago",
		"la setm que ve":     "in 1 week",
		"el mes passat":      "1 month ago",
		"el mes que ve":      "in 1 month",
		"endema passat":      "in 3 day",
		"aquesta hora":       "0 hour ago",
		"aquest minut":       "0 minute ago",
		"aquesta setm":       "0 week ago",
		"setm passada":       "1 week ago",
		"l'any passat":       "1 year ago",
		"abans-d’ahir":       "2 day ago",
		"l'any que ve":       "in 1 year",
		"despus-ahir":        "2 day ago",
		"setm vinent":        "in 1 week",
		"dema passat":        "in 2 day",
		"despus-dema":        "in 2 day",
		"passat dema":        "in 2 day",
		"aquest mes":         "0 month ago",
		"mes passat":         "1 month ago",
		"della-ahir":         "2 day ago",
		"mes vinent":         "in 1 month",
		"enguany":            "0 year ago",
		"sendema":            "in 2 day",
		"endema":             "in 2 day",
		"avui":               "0 day ago",
		"ahir":               "1 day ago",
		"dema":               "in 1 day",
		"hui":                "0 day ago",
		"ara":                "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)d'aqui a (\d+) setmanes`), "in $1 week"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) setmana`), "in $1 week"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) minuts`), "in $1 minute"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) segons`), "in $1 second"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) hores`), "in $1 hour"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) minut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) mesos`), "in $1 month"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) segon`), "in $1 second"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) dies`), "in $1 day"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) hora`), "in $1 hour"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) setm`), "in $1 week"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) anys`), "in $1 year"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) dia`), "in $1 day"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) mes`), "in $1 month"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) any`), "in $1 year"},
		{regexp.MustCompile(`(?i)fa (\d+) setmanes`), "$1 week ago"},
		{regexp.MustCompile(`(?i)fa (\d+) setmana`), "$1 week ago"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) h`), "in $1 hour"},
		{regexp.MustCompile(`(?i)d‘aqui a (\d+) h`), "in $1 hour"},
		{regexp.MustCompile(`(?i)d'aqui a (\d+) s`), "in $1 second"},
		{regexp.MustCompile(`(?i)fa (\d+) minuts`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)fa (\d+) segons`), "$1 second ago"},
		{regexp.MustCompile(`(?i)fa (\d+) hores`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)fa (\d+) minut`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)fa (\d+) mesos`), "$1 month ago"},
		{regexp.MustCompile(`(?i)fa (\d+) segon`), "$1 second ago"},
		{regexp.MustCompile(`(?i)fa (\d+) dies`), "$1 day ago"},
		{regexp.MustCompile(`(?i)fa (\d+) hora`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)fa (\d+) setm`), "$1 week ago"},
		{regexp.MustCompile(`(?i)fa (\d+) anys`), "$1 year ago"},
		{regexp.MustCompile(`(?i)fa (\d+) dia`), "$1 day ago"},
		{regexp.MustCompile(`(?i)fa (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)fa (\d+) mes`), "$1 month ago"},
		{regexp.MustCompile(`(?i)fa (\d+) any`), "$1 year ago"},
		{regexp.MustCompile(`(?i)fa (\d+) h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)fa (\d+) s`), "$1 second ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(d'aqui a \d+ setmanes|d'aqui a \d+ setmana|d'aqui a \d+ minuts|d'aqui a \d+ segons|d'aqui a \d+ hores|d'aqui a \d+ mesos|d'aqui a \d+ minut|d'aqui a \d+ segon|d'aqui a \d+ anys|d'aqui a \d+ dies|d'aqui a \d+ hora|d'aqui a \d+ setm|d'aqui a \d+ any|d'aqui a \d+ dia|d'aqui a \d+ mes|d'aqui a \d+ min|fa \d+ setmanes|d'aqui a \d+ h|d'aqui a \d+ s|d‘aqui a \d+ h|fa \d+ setmana|fa \d+ minuts|fa \d+ segons|fa \d+ hores|fa \d+ mesos|fa \d+ minut|fa \d+ segon|fa \d+ anys|fa \d+ dies|fa \d+ hora|fa \d+ setm|fa \d+ any|fa \d+ dia|fa \d+ mes|fa \d+ min|fa \d+ h|fa \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(d'aqui a \d+ setmanes|d'aqui a \d+ setmana|d'aqui a \d+ minuts|d'aqui a \d+ segons|d'aqui a \d+ hores|d'aqui a \d+ mesos|d'aqui a \d+ minut|d'aqui a \d+ segon|d'aqui a \d+ anys|d'aqui a \d+ dies|d'aqui a \d+ hora|d'aqui a \d+ setm|d'aqui a \d+ any|d'aqui a \d+ dia|d'aqui a \d+ mes|d'aqui a \d+ min|fa \d+ setmanes|d'aqui a \d+ h|d'aqui a \d+ s|d‘aqui a \d+ h|fa \d+ setmana|fa \d+ minuts|fa \d+ segons|fa \d+ hores|fa \d+ mesos|fa \d+ minut|fa \d+ segon|fa \d+ anys|fa \d+ dies|fa \d+ hora|fa \d+ setm|fa \d+ any|fa \d+ dia|fa \d+ mes|fa \d+ min|fa \d+ h|fa \d+ s)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(la propera setmana|la proxima setmana|la setmana passada|la setmana que ve|la setmana vinent|aquesta setmana|la setm passada|la setm que ve|el mes passat|el mes que ve|endema passat|abans-d’ahir|aquest minut|aquesta hora|aquesta setm|l'any passat|l'any que ve|setm passada|de desembre|de novembre|de setembre|dema passat|despus-ahir|despus-dema|passat dema|setm vinent|aquest mes|della-ahir|mes passat|mes vinent|d'octubre|de febrer|de juliol|divendres|de gener|desembre|dimecres|dissabte|diumenge|novembre|setembre|d'abril|d'agost|de febr|de juny|de maig|de marc|dilluns|dimarts|enguany|octubre|sendema|setmana|de des|de gen|de jul|de nov|de set|dijous|endema|febrer|juliol|abril|agost|d'abr|d'oct|gener|minut|segon|ahir|avui|d'ag|dema|febr|hora|juny|maig|marc|setm|abr|any|ara|del|des|dia|gen|gmt|hui|jul|mes|min|nov|oct|set|utc|\+|\.|\[|\]|\||ag|am|dc|de|dg|dj|dl|ds|dt|dv|en|l'|pm| |'|,|-|/|:|;|@|h|i|s|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ca_AD_Locale = merge(&ca_Locale, LocaleData{
	Name:      "ca-AD",
	DateOrder: "DMY",
})

var ca_FR_Locale = merge(&ca_Locale, LocaleData{
	Name:      "ca-FR",
	DateOrder: "DMY",
})

var ca_IT_Locale = merge(&ca_Locale, LocaleData{
	Name:      "ca-IT",
	DateOrder: "DMY",
})
