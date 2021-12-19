// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ta_Locale = merge(nil, LocaleData{
	Name:      "ta",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"பிபரவரி": "february",
		"செபடமபர": "september",
		"முறபகல":  "am",
		"டிசமபர":  "december",
		"நிமிடம":  "minute",
		"அகடோபர":  "october",
		"பிறபகல":  "pm",
		"விநாடி":  "second",
		"ஞாயிறு":  "sunday",
		"வியாழன":  "thursday",
		"செவவாய":  "tuesday",
		"வெளளி":   "friday",
		"ஜனவரி":   "january",
		"திஙகள":   "monday",
		"நவமபர":   "november",
		"ஏபரல":    "april",
		"ஆகஸட":    "august",
		"ஜூலை":    "july",
		"மாரச":    "march",
		"நிமி":    "minute",
		"மாதம":    "month",
		"விநா":    "second",
		"ஞாயி":    "sunday",
		"வியா":    "thursday",
		"புதன":    "wednesday",
		"வாரம":    "week",
		"ஆணடு":    "year",
		"நாள":     "day",
		"டிச":     "december",
		"பிப":     "february",
		"வெள":     "friday",
		"gmt":     "gmt",
		"மணி":     "hour",
		"ஜூன":     "june",
		"மார":     "march",
		"திங":     "monday",
		"மாத":     "month",
		"சனி":     "saturday",
		"செப":     "september",
		"செவ":     "tuesday",
		"utc":     "utc",
		"புத":     "wednesday",
		"am":      "am",
		"ஏப":      "april",
		"ஆக":      "august",
		"நா":      "day",
		"ஜன":      "january",
		"மே":      "may",
		"மா":      "month",
		"நவ":      "november",
		"அக":      "october",
		"pm":      "pm",
		"வி":      "second",
		"வா":      "week",
		"'":       "",
		",":       "",
		";":       "",
		"@":       "",
		"[":       "",
		"]":       "",
		"|":       "",
		" ":       " ",
		"+":       "+",
		"-":       "-",
		".":       ".",
		"/":       "/",
		":":       ":",
		"ம":       "hour",
		"ஆ":       "year",
		"z":       "z",
	},
	RelativeType: map[string]string{
		"இநத ஒரு மணிநேரததில": "0 hour ago",
		"இநத ஒரு நிமிடததில":  "0 minute ago",
		"அடுதத மாதம":         "in 1 month",
		"அடுதத வாரம":         "in 1 week",
		"அடுதத ஆணடு":         "in 1 year",
		"கடநத மாதம":          "1 month ago",
		"கடநத வாரம":          "1 week ago",
		"கடநத ஆணடு":          "1 year ago",
		"இநத மாதம":           "0 month ago",
		"இநத வாரம":           "0 week ago",
		"இநத ஆணடு":           "0 year ago",
		"இபபோது":             "0 second ago",
		"நேறறு":              "1 day ago",
		"இனறு":               "0 day ago",
		"நாளை":               "in 1 day",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) நிமிடஙகளுககு முன`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) விநாடிகளுககு முன`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) நிமிடததிறகு முன`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) வாரததிறகு முனபு`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) மாதஙகளுககு முன`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) வாரஙகளுககு முன`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ஆணடுகளுககு முன`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) நாடகளுககு முன`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) மாதததுககு முன`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) விநாடிககு முன`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) நாளுககு முன`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) மணிநேரம முன`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) ஆணடிறகு முன`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) மணிநேரததில`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) நிமிடஙகளில`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) விநாடிகளில`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) நிமிடததில`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) விநாடியில`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) நிமி முன`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) விநா முன`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) மாதஙகளில`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) வாரஙகளில`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ஆணடுகளில`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) மணி முன`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) மாத முன`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) வார முன`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) நாடகளில`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) மாதததில`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) வாரததில`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) நா முன`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) நி முன`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) மா முன`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) வி முன`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) வா முன`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ம முன`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) ஆ முன`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) நாளில`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) ஆணடில`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) நிமி`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) விநா`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) மணி`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) மாத`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) வார`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) நா`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) நி`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) மா`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) வி`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) வா`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ம`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) ஆ`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ நிமிடஙகளுககு முன|\d+ விநாடிகளுககு முன|\d+ நிமிடததிறகு முன|\d+ வாரததிறகு முனபு|\d+ ஆணடுகளுககு முன|\d+ மாதஙகளுககு முன|\d+ வாரஙகளுககு முன|\d+ நாடகளுககு முன|\d+ மாதததுககு முன|\d+ விநாடிககு முன|\d+ ஆணடிறகு முன|\d+ நாளுககு முன|\d+ மணிநேரம முன|\d+ நிமிடஙகளில|\d+ மணிநேரததில|\d+ விநாடிகளில|\d+ நிமிடததில|\d+ விநாடியில|\d+ ஆணடுகளில|\d+ நிமி முன|\d+ மாதஙகளில|\d+ வாரஙகளில|\d+ விநா முன|\d+ நாடகளில|\d+ மணி முன|\d+ மாத முன|\d+ மாதததில|\d+ வார முன|\d+ வாரததில|\d+ நா முன|\d+ நி முன|\d+ மா முன|\d+ வா முன|\d+ வி முன|\d+ ஆ முன|\d+ ஆணடில|\d+ நாளில|\d+ ம முன|\d+ நிமி|\d+ விநா|\d+ மணி|\d+ மாத|\d+ வார|\d+ நா|\d+ நி|\d+ மா|\d+ வா|\d+ வி|\d+ ஆ|\d+ ம)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ நிமிடஙகளுககு முன|\d+ விநாடிகளுககு முன|\d+ நிமிடததிறகு முன|\d+ வாரததிறகு முனபு|\d+ ஆணடுகளுககு முன|\d+ மாதஙகளுககு முன|\d+ வாரஙகளுககு முன|\d+ நாடகளுககு முன|\d+ மாதததுககு முன|\d+ விநாடிககு முன|\d+ ஆணடிறகு முன|\d+ நாளுககு முன|\d+ மணிநேரம முன|\d+ நிமிடஙகளில|\d+ மணிநேரததில|\d+ விநாடிகளில|\d+ நிமிடததில|\d+ விநாடியில|\d+ ஆணடுகளில|\d+ நிமி முன|\d+ மாதஙகளில|\d+ வாரஙகளில|\d+ விநா முன|\d+ நாடகளில|\d+ மணி முன|\d+ மாத முன|\d+ மாதததில|\d+ வார முன|\d+ வாரததில|\d+ நா முன|\d+ நி முன|\d+ மா முன|\d+ வா முன|\d+ வி முன|\d+ ஆ முன|\d+ ஆணடில|\d+ நாளில|\d+ ம முன|\d+ நிமி|\d+ விநா|\d+ மணி|\d+ மாத|\d+ வார|\d+ நா|\d+ நி|\d+ மா|\d+ வா|\d+ வி|\d+ ஆ|\d+ ம)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(இநத ஒரு மணிநேரததில|இநத ஒரு நிமிடததில|அடுதத ஆணடு|அடுதத மாதம|அடுதத வாரம|கடநத ஆணடு|கடநத மாதம|கடநத வாரம|இநத ஆணடு|இநத மாதம|இநத வாரம|செபடமபர|பிபரவரி|அகடோபர|இபபோது|செவவாய|ஞாயிறு|டிசமபர|நிமிடம|பிறபகல|முறபகல|விநாடி|வியாழன|ஜனவரி|திஙகள|நவமபர|நேறறு|வெளளி|ஆகஸட|ஆணடு|இனறு|ஏபரல|ஜூலை|ஞாயி|நாளை|நிமி|புதன|மாதம|மாரச|வாரம|விநா|வியா|gmt|utc|சனி|செப|செவ|ஜூன|டிச|திங|நாள|பிப|புத|மணி|மாத|மார|வெள|\+|\.|\[|\]|\||am|pm|அக|ஆக|ஏப|ஜன|நவ|நா|மா|மே|வா|வி| |'|,|-|/|:|;|@|z|ஆ|ம)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var ta_LK_Locale = merge(&ta_Locale, LocaleData{
	Name:      "ta-LK",
	DateOrder: "DMY",
})

var ta_MY_Locale = merge(&ta_Locale, LocaleData{
	Name:      "ta-MY",
	DateOrder: "DMY",
})

var ta_SG_Locale = merge(&ta_Locale, LocaleData{
	Name:      "ta-SG",
	DateOrder: "DMY",
})
