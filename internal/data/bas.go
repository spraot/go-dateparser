// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var bas_Locale = merge(nil, LocaleData{
	Name:      "bas",
	DateOrder: "DMY",
	Charset:   []rune(`bcdeghijklnorstuwxyzàèêìòôùûńŋɓɔɛ̀̂`),
	Translations: map[string]string{
		"ŋgwa njaŋgumba": "monday",
		"libuy li nyee":  "december",
		"i bikɛgla":      "am",
		"i ɓugajɔp":      "pm",
		"hiŋgeŋget":      "second",
		"ŋgwa mbɔk":      "thursday",
		"ŋgwa kɔɔ":       "friday",
		"ŋgwa jon":       "saturday",
		"ŋgwa nɔy":       "sunday",
		"ŋgwa ŋge":       "wednesday",
		"hilondɛ":        "june",
		"mayɛsep":        "november",
		"ŋgwa um":        "tuesday",
		"kɔndɔŋ":         "january",
		"matumb":         "march",
		"matop":          "april",
		"hikaŋ":          "august",
		"macɛl":          "february",
		"njeba":          "july",
		"mpuyɛ":          "may",
		"bioom":          "october",
		"dipɔs":          "september",
		"sɔndɛ":          "week",
		"ŋgɛŋ":           "hour",
		"ŋget":           "minute",
		"ŋwii":           "year",
		"mto":            "april",
		"hik":            "august",
		"kɛl":            "day",
		"liɓ":            "december",
		"mac":            "february",
		"kɔɔ":            "friday",
		"gmt":            "gmt",
		"kɔn":            "january",
		"nje":            "july",
		"hil":            "june",
		"mat":            "march",
		"mpu":            "may",
		"nja":            "monday",
		"soŋ":            "month",
		"may":            "november",
		"bio":            "october",
		"jon":            "saturday",
		"dip":            "september",
		"nɔy":            "sunday",
		"mbɔ":            "thursday",
		"uum":            "tuesday",
		"utc":            "utc",
		"ŋge":            "wednesday",
		"am":             "am",
		"pm":             "pm",
		"'":              "",
		",":              "",
		";":              "",
		"@":              "",
		"[":              "",
		"]":              "",
		"|":              "",
		" ":              " ",
		"+":              "+",
		"-":              "-",
		".":              ".",
		"/":              "/",
		":":              ":",
		"z":              "z",
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
		"yaani":       "1 day ago",
		"yani":        "in 1 day",
		"lɛn":         "0 day ago",
		"now":         "0 second ago",
	},
	RxKnownWords: regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(ŋgwa njaŋgumba|libuy li nyee|this minute|last month|next month|this month|hiŋgeŋget|i bikɛgla|i ɓugajɔp|last week|last year|next week|next year|this hour|this week|this year|ŋgwa mbɔk|ŋgwa jon|ŋgwa kɔɔ|ŋgwa nɔy|ŋgwa ŋge|hilondɛ|mayɛsep|ŋgwa um|kɔndɔŋ|matumb|bioom|dipɔs|hikaŋ|macɛl|matop|mpuyɛ|njeba|sɔndɛ|yaani|yani|ŋget|ŋgɛŋ|ŋwii|bio|dip|gmt|hik|hil|jon|kɔn|kɔɔ|kɛl|liɓ|lɛn|mac|mat|may|mbɔ|mpu|mto|nja|nje|now|nɔy|soŋ|utc|uum|ŋge|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
