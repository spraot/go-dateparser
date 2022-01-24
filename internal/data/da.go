// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var da_Locale = merge(nil, LocaleData{
	Name:      "da",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "cirka", "d.", "kl", "kl.", "|"},
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)en(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)et(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+)\s*hr(s?)(\z|[^\pL\pM\d]|_)`), "${1}${2} time${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+)\s*min(s?)(\z|[^\pL\pM\d]|_)`), "${1}${2} minut${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+)\s*sec(s?)(\z|[^\pL\pM\d]|_)`), "${1}${2} sekund${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)middag(\z|[^\pL\pM\d]|_)`), "${1}12:00${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)midnat(\z|[^\pL\pM\d]|_)`), "${1}00:00${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+)h(\d+)m?(\z|[^\pL\pM\d]|_)`), "${1}${2}:${3}${4}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)mindre end 1 minut siden(\z|[^\pL\pM\d]|_)`), "${1}45 seconds${2}"},
	},
	Translations: map[string]string{
		"september": "september",
		"december":  "december",
		"minutter":  "minute",
		"november":  "november",
		"sekunder":  "second",
		"februar":   "february",
		"maneder":   "month",
		"oktober":   "october",
		"torsdag":   "thursday",
		"tirsdag":   "tuesday",
		"august":    "august",
		"fredag":    "friday",
		"januar":    "january",
		"mandag":    "monday",
		"lørdag":    "saturday",
		"sekund":    "second",
		"søndag":    "sunday",
		"onsdag":    "wednesday",
		"cirka":     "",
		"siden":     "ago",
		"april":     "april",
		"timer":     "hour",
		"marts":     "march",
		"minut":     "minute",
		"maned":     "month",
		"dage":      "day",
		"time":      "hour",
		"juli":      "july",
		"juni":      "june",
		"uger":      "week",
		"kl.":       "",
		"apr":       "april",
		"aug":       "august",
		"dag":       "day",
		"dec":       "december",
		"feb":       "february",
		"fre":       "friday",
		"gmt":       "gmt",
		"jan":       "january",
		"jul":       "july",
		"jun":       "june",
		"mar":       "march",
		"maj":       "may",
		"min":       "minute",
		"man":       "monday",
		"nov":       "november",
		"okt":       "october",
		"lør":       "saturday",
		"sek":       "second",
		"sep":       "september",
		"søn":       "sunday",
		"tor":       "thursday",
		"tir":       "tuesday",
		"utc":       "utc",
		"ons":       "wednesday",
		"uge":       "week",
		"d.":        "",
		"kl":        "",
		"am":        "am",
		"md":        "month",
		"pm":        "pm",
		"ar":        "year",
		"'":         "",
		",":         "",
		";":         "",
		"@":         "",
		"[":         "",
		"]":         "",
		"|":         "",
		" ":         " ",
		"+":         "+",
		"-":         "-",
		".":         ".",
		"/":         "/",
		":":         ":",
		"t":         "hour",
		"i":         "in",
		"s":         "second",
		"z":         "z",
	},
	RelativeType: map[string]string{
		"i det kommende minut": "0 minute ago",
		"i den kommende time":  "0 hour ago",
		"sidste maned":         "1 month ago",
		"denne maned":          "0 month ago",
		"næste maned":          "in 1 month",
		"sidste uge":           "1 week ago",
		"denne uge":            "0 week ago",
		"sidste md":            "1 month ago",
		"sidste ar":            "1 year ago",
		"næste uge":            "in 1 week",
		"denne md":             "0 month ago",
		"i morgen":             "in 1 day",
		"næste md":             "in 1 month",
		"næste ar":             "in 1 year",
		"i dag":                "0 day ago",
		"i gar":                "1 day ago",
		"i ar":                 "0 year ago",
		"nu":                   "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)for (\d+) minutter siden`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)for (\d+) sekunder siden`), "$1 second ago"},
		{regexp.MustCompile(`(?i)for (\d+) maneder siden`), "$1 month ago"},
		{regexp.MustCompile(`(?i)for (\d+) sekund siden`), "$1 second ago"},
		{regexp.MustCompile(`(?i)for (\d+) timer siden`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)for (\d+) minut siden`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)for (\d+) maned siden`), "$1 month ago"},
		{regexp.MustCompile(`(?i)for (\d+) dage siden`), "$1 day ago"},
		{regexp.MustCompile(`(?i)for (\d+) time siden`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)for (\d+) uger siden`), "$1 week ago"},
		{regexp.MustCompile(`(?i)for (\d+) dag siden`), "$1 day ago"},
		{regexp.MustCompile(`(?i)for (\d+) min siden`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)for (\d+) mdr siden`), "$1 month ago"},
		{regexp.MustCompile(`(?i)for (\d+) sek siden`), "$1 second ago"},
		{regexp.MustCompile(`(?i)for (\d+) uge siden`), "$1 week ago"},
		{regexp.MustCompile(`(?i)for (\d+) minutter`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)for (\d+) md siden`), "$1 month ago"},
		{regexp.MustCompile(`(?i)for (\d+) sekunder`), "$1 second ago"},
		{regexp.MustCompile(`(?i)for (\d+) ar siden`), "$1 year ago"},
		{regexp.MustCompile(`(?i)om (\d+) minutter`), "in $1 minute"},
		{regexp.MustCompile(`(?i)om (\d+) sekunder`), "in $1 second"},
		{regexp.MustCompile(`(?i)om (\d+) maneder`), "in $1 month"},
		{regexp.MustCompile(`(?i)for (\d+) timer`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)om (\d+) sekund`), "in $1 second"},
		{regexp.MustCompile(`(?i)om (\d+) timer`), "in $1 hour"},
		{regexp.MustCompile(`(?i)om (\d+) minut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)om (\d+) maned`), "in $1 month"},
		{regexp.MustCompile(`(?i)for (\d+)\s*h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)for (\d+)\s*m`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)for (\d+)\s*s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)om (\d+) dage`), "in $1 day"},
		{regexp.MustCompile(`(?i)om (\d+) time`), "in $1 hour"},
		{regexp.MustCompile(`(?i)om (\d+) uger`), "in $1 week"},
		{regexp.MustCompile(`(?i)om (\d+) dag`), "in $1 day"},
		{regexp.MustCompile(`(?i)om (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)om (\d+) mdr`), "in $1 month"},
		{regexp.MustCompile(`(?i)om (\d+) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)om (\d+) uge`), "in $1 week"},
		{regexp.MustCompile(`(?i)om (\d+) md`), "in $1 month"},
		{regexp.MustCompile(`(?i)om (\d+) ar`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(for \d+ minutter siden|for \d+ sekunder siden|for \d+ maneder siden|for \d+ sekund siden|for \d+ maned siden|for \d+ minut siden|for \d+ timer siden|for \d+ dage siden|for \d+ time siden|for \d+ uger siden|for \d+ dag siden|for \d+ mdr siden|for \d+ min siden|for \d+ sek siden|for \d+ uge siden|for \d+ ar siden|for \d+ md siden|for \d+ minutter|for \d+ sekunder|om \d+ minutter|om \d+ sekunder|om \d+ maneder|for \d+ timer|om \d+ sekund|om \d+ maned|om \d+ minut|om \d+ timer|for \d+\s*h|for \d+\s*m|for \d+\s*s|om \d+ dage|om \d+ time|om \d+ uger|om \d+ dag|om \d+ mdr|om \d+ min|om \d+ sek|om \d+ uge|om \d+ ar|om \d+ md)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(for \d+ minutter siden|for \d+ sekunder siden|for \d+ maneder siden|for \d+ sekund siden|for \d+ maned siden|for \d+ minut siden|for \d+ timer siden|for \d+ dage siden|for \d+ time siden|for \d+ uger siden|for \d+ dag siden|for \d+ mdr siden|for \d+ min siden|for \d+ sek siden|for \d+ uge siden|for \d+ ar siden|for \d+ md siden|for \d+ minutter|for \d+ sekunder|om \d+ minutter|om \d+ sekunder|om \d+ maneder|for \d+ timer|om \d+ sekund|om \d+ maned|om \d+ minut|om \d+ timer|for \d+\s*h|for \d+\s*m|for \d+\s*s|om \d+ dage|om \d+ time|om \d+ uger|om \d+ dag|om \d+ mdr|om \d+ min|om \d+ sek|om \d+ uge|om \d+ ar|om \d+ md)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(i det kommende minut|i den kommende time|sidste maned|denne maned|næste maned|sidste uge|denne uge|næste uge|september|sidste ar|sidste md|december|denne md|i morgen|minutter|november|næste ar|næste md|sekunder|februar|maneder|oktober|tirsdag|torsdag|august|fredag|januar|lørdag|mandag|onsdag|sekund|søndag|april|cirka|i dag|i gar|maned|marts|minut|siden|timer|dage|i ar|juli|juni|kl\.|time|uger|apr|aug|d\.|dag|dec|feb|fre|gmt|jan|jul|jun|lør|maj|man|mar|min|nov|okt|ons|sek|sep|søn|tir|tor|uge|utc|\+|\.|\[|\]|\||am|ar|kl|md|nu|pm| |'|,|-|/|:|;|@|i|s|t|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var da_GL_Locale = merge(&da_Locale, LocaleData{
	Name:            "da-GL",
	DateOrder:       "DMY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(for \d+ minutter siden|for \d+ sekunder siden|for \d+ maneder siden|for \d+ sekund siden|for \d+ maned siden|for \d+ minut siden|for \d+ timer siden|for \d+ dage siden|for \d+ time siden|for \d+ uger siden|for \d+ dag siden|for \d+ mdr siden|for \d+ min siden|for \d+ sek siden|for \d+ uge siden|for \d+ ar siden|for \d+ md siden|for \d+ minutter|for \d+ sekunder|om \d+ minutter|om \d+ sekunder|om \d+ maneder|for \d+ timer|om \d+ sekund|om \d+ maned|om \d+ minut|om \d+ timer|for \d+\s*h|for \d+\s*m|for \d+\s*s|om \d+ dage|om \d+ time|om \d+ uger|om \d+ dag|om \d+ mdr|om \d+ min|om \d+ sek|om \d+ uge|om \d+ ar|om \d+ md)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(for \d+ minutter siden|for \d+ sekunder siden|for \d+ maneder siden|for \d+ sekund siden|for \d+ maned siden|for \d+ minut siden|for \d+ timer siden|for \d+ dage siden|for \d+ time siden|for \d+ uger siden|for \d+ dag siden|for \d+ mdr siden|for \d+ min siden|for \d+ sek siden|for \d+ uge siden|for \d+ ar siden|for \d+ md siden|for \d+ minutter|for \d+ sekunder|om \d+ minutter|om \d+ sekunder|om \d+ maneder|for \d+ timer|om \d+ sekund|om \d+ maned|om \d+ minut|om \d+ timer|for \d+\s*h|for \d+\s*m|for \d+\s*s|om \d+ dage|om \d+ time|om \d+ uger|om \d+ dag|om \d+ mdr|om \d+ min|om \d+ sek|om \d+ uge|om \d+ ar|om \d+ md)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(i det kommende minut|i den kommende time|sidste maned|denne maned|næste maned|sidste uge|denne uge|næste uge|september|sidste ar|sidste md|december|denne md|i morgen|minutter|november|næste ar|næste md|sekunder|februar|maneder|oktober|tirsdag|torsdag|august|fredag|januar|lørdag|mandag|onsdag|sekund|søndag|april|cirka|i dag|i gar|maned|marts|minut|siden|timer|dage|i ar|juli|juni|kl\.|time|uger|apr|aug|d\.|dag|dec|feb|fre|gmt|jan|jul|jun|lør|maj|man|mar|min|nov|okt|ons|sek|sep|søn|tir|tor|uge|utc|\+|\.|\[|\]|\||am|ar|kl|md|nu|pm| |'|,|-|/|:|;|@|i|s|t|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
