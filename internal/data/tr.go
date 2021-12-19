// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var tr_Locale = merge(nil, LocaleData{
	Name:      "tr",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "ve", "yaklasık", "|"},
	Translations: map[string]string{
		"icerisinde": "in",
		"pazartesi":  "monday",
		"cumartesi":  "saturday",
		"yaklasık":   "",
		"persembe":   "thursday",
		"carsamba":   "wednesday",
		"agustos":    "august",
		"haziran":    "june",
		"aralık":     "december",
		"icinde":     "in",
		"temmuz":     "july",
		"dakika":     "minute",
		"saniye":     "second",
		"nisan":      "april",
		"subat":      "february",
		"sonra":      "in",
		"mayıs":      "may",
		"kasım":      "november",
		"eylul":      "september",
		"pazar":      "sunday",
		"hafta":      "week",
		"once":       "ago",
		"cuma":       "friday",
		"saat":       "hour",
		"ocak":       "january",
		"mart":       "march",
		"ekim":       "october",
		"salı":       "tuesday",
		"sene":       "year",
		"nis":        "april",
		"agu":        "august",
		"gun":        "day",
		"ara":        "december",
		"sub":        "february",
		"cum":        "friday",
		"gmt":        "gmt",
		"oca":        "january",
		"tem":        "july",
		"haz":        "june",
		"mar":        "march",
		"may":        "may",
		"pzt":        "monday",
		"kas":        "november",
		"eki":        "october",
		"cmt":        "saturday",
		"eyl":        "september",
		"paz":        "sunday",
		"per":        "thursday",
		"prs":        "thursday",
		"sal":        "tuesday",
		"utc":        "utc",
		"car":        "wednesday",
		"crs":        "wednesday",
		"yıl":        "year",
		"ve":         "",
		"am":         "am",
		"oo":         "am",
		"ni":         "april",
		"ag":         "august",
		"ar":         "december",
		"su":         "february",
		"sa":         "hour",
		"oc":         "january",
		"te":         "july",
		"ha":         "june",
		"dk":         "minute",
		"ay":         "month",
		"ka":         "november",
		"ek":         "october",
		"os":         "pm",
		"pm":         "pm",
		"sn":         "second",
		"ey":         "september",
		"hf":         "week",
		"'":          "",
		",":          "",
		";":          "",
		"@":          "",
		"[":          "",
		"]":          "",
		"|":          "",
		" ":          " ",
		"+":          "+",
		"-":          "-",
		".":          ".",
		"/":          "/",
		":":          ":",
		"z":          "z",
	},
	RelativeType: map[string]string{
		"onumuzdeki hafta": "in 1 week",
		"onumuzdeki gun":   "in 1 day",
		"onumuzdeki yıl":   "in 1 year",
		"onumuzdeki ay":    "in 1 month",
		"gelecek hafta":    "in 1 week",
		"gecen hafta":      "1 week ago",
		"gelecek yıl":      "in 1 year",
		"gelecek ay":       "in 1 month",
		"bu dakika":        "0 minute ago",
		"gecen gun":        "1 day ago",
		"gecen yıl":        "1 year ago",
		"bu hafta":         "0 week ago",
		"gecen ay":         "1 month ago",
		"bu saat":          "0 hour ago",
		"haftaya":          "in 1 week",
		"bu yıl":           "0 year ago",
		"bugun":            "0 day ago",
		"bu ay":            "0 month ago",
		"simdi":            "0 second ago",
		"yarın":            "in 1 day",
		"dun":              "1 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) dakika sonra`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) saniye sonra`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) dakika once`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) saniye once`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) hafta sonra`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) hafta once`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) saat sonra`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) saat once`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) gun sonra`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) yıl sonra`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) gun once`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) yıl once`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) sa sonra`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) dk sonra`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ay sonra`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) sn sonra`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) hf sonra`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) sa once`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) dk once`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) ay once`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) sn once`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) hf once`), "$1 week ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ dakika sonra|\d+ saniye sonra|\d+ dakika once|\d+ hafta sonra|\d+ saniye once|\d+ hafta once|\d+ saat sonra|\d+ gun sonra|\d+ saat once|\d+ yıl sonra|\d+ ay sonra|\d+ dk sonra|\d+ gun once|\d+ hf sonra|\d+ sa sonra|\d+ sn sonra|\d+ yıl once|\d+ ay once|\d+ dk once|\d+ hf once|\d+ sa once|\d+ sn once)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ dakika sonra|\d+ saniye sonra|\d+ dakika once|\d+ hafta sonra|\d+ saniye once|\d+ hafta once|\d+ saat sonra|\d+ gun sonra|\d+ saat once|\d+ yıl sonra|\d+ ay sonra|\d+ dk sonra|\d+ gun once|\d+ hf sonra|\d+ sa sonra|\d+ sn sonra|\d+ yıl once|\d+ ay once|\d+ dk once|\d+ hf once|\d+ sa once|\d+ sn once)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(onumuzdeki hafta|onumuzdeki gun|onumuzdeki yıl|gelecek hafta|onumuzdeki ay|gecen hafta|gelecek yıl|gelecek ay|icerisinde|bu dakika|cumartesi|gecen gun|gecen yıl|pazartesi|bu hafta|carsamba|gecen ay|persembe|yaklasık|agustos|bu saat|haftaya|haziran|aralık|bu yıl|dakika|icinde|saniye|temmuz|bu ay|bugun|eylul|hafta|kasım|mayıs|nisan|pazar|simdi|sonra|subat|yarın|cuma|ekim|mart|ocak|once|saat|salı|sene|agu|ara|car|cmt|crs|cum|dun|eki|eyl|gmt|gun|haz|kas|mar|may|nis|oca|paz|per|prs|pzt|sal|sub|tem|utc|yıl|\+|\.|\[|\]|\||ag|am|ar|ay|dk|ek|ey|ha|hf|ka|ni|oc|oo|os|pm|sa|sn|su|te|ve| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var tr_CY_Locale = merge(&tr_Locale, LocaleData{
	Name:      "tr-CY",
	DateOrder: "DMY",
})
