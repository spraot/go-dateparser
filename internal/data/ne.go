// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ne_Locale = merge(nil, LocaleData{
	Name:      "ne",
	DateOrder: "YMD",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"यही मिनटमा": "0 minute ago",
		"अरको महिना": "in 1 month",
		"यो घडीमा":   "0 hour ago",
		"यो महिना":   "0 month ago",
		"गत महिना":   "1 month ago",
		"आउन हपता":   "in 1 week",
		"अरको वरष":   "in 1 year",
		"यो हपता":    "0 week ago",
		"गत हपता":    "1 week ago",
		"बिहिबार":    "thursday",
		"मङगलबार":    "tuesday",
		"यो वरष":     "0 year ago",
		"गत वरष":     "1 year ago",
		"परवाहन":     "am",
		"डिसमबर":     "december",
		"फबरअरी":     "february",
		"शकरबार":     "friday",
		"सोमबार":     "monday",
		"नोभमबर":     "november",
		"अकटोबर":     "october",
		"अपराहन":     "pm",
		"शनिबार":     "saturday",
		"सपटमबर":     "september",
		"आइतबार":     "sunday",
		"अपरिल":      "april",
		"जनवरी":      "january",
		"महिना":      "month",
		"बधबार":      "wednesday",
		"हिजो":       "1 day ago",
		"अगसट":       "august",
		"घणटा":       "hour",
		"भोलि":       "in 1 day",
		"जलाई":       "july",
		"मारच":       "march",
		"मिनट":       "minute",
		"सकनड":       "second",
		"बिहि":       "thursday",
		"मङगल":       "tuesday",
		"हपता":       "week",
		"बार":        "day",
		"शकर":        "friday",
		"gmt":        "gmt",
		"सोम":        "monday",
		"शनि":        "saturday",
		"आइत":        "sunday",
		"utc":        "utc",
		"बरष":        "year",
		"वरष":        "year",
		"आज":         "0 day ago",
		"अब":         "0 second ago",
		"am":         "am",
		"जन":         "june",
		"मई":         "may",
		"pm":         "pm",
		"बध":         "wednesday",
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
		"म":          "may",
		"z":          "z",
	},
	TranslationRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) महिना पहिल`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) घणटा पहिल`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) मिनट पहिल`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) सकणड पहिल`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) हपता पहिल`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) दिन पहिल`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) वरष अघि`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) महिनामा`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) घणटामा`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) मिनटमा`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) सकणडमा`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) हपतामा`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) दिनमा`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) वरषमा`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|\W|_)(\d+ महिना पहिल|\d+ घणटा पहिल|\d+ मिनट पहिल|\d+ सकणड पहिल|\d+ हपता पहिल|\d+ दिन पहिल|\d+ महिनामा|\d+ वरष अघि|\d+ घणटामा|\d+ मिनटमा|\d+ सकणडमा|\d+ हपतामा|\d+ दिनमा|\d+ वरषमा)(\z|\W|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ महिना पहिल|\d+ घणटा पहिल|\d+ मिनट पहिल|\d+ सकणड पहिल|\d+ हपता पहिल|\d+ दिन पहिल|\d+ महिनामा|\d+ वरष अघि|\d+ घणटामा|\d+ मिनटमा|\d+ सकणडमा|\d+ हपतामा|\d+ दिनमा|\d+ वरषमा)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|\W|_|\d))(अरको महिना|यही मिनटमा|अरको वरष|आउन हपता|गत महिना|यो घडीमा|यो महिना|गत हपता|बिहिबार|मङगलबार|यो हपता|अकटोबर|अपराहन|आइतबार|गत वरष|डिसमबर|नोभमबर|परवाहन|फबरअरी|यो वरष|शकरबार|शनिबार|सपटमबर|सोमबार|अपरिल|जनवरी|बधबार|महिना|अगसट|घणटा|जलाई|बिहि|भोलि|मङगल|मारच|मिनट|सकनड|हपता|हिजो|gmt|utc|आइत|बरष|बार|वरष|शकर|शनि|सोम|\+|\.|\[|\]|\||am|pm|अब|आज|जन|बध|मई| |'|,|-|/|:|;|@|z|म)((?:\z|\W|_|\d).*)$`),
})

var ne_IN_Locale = merge(&ne_Locale, LocaleData{
	Name:      "ne-IN",
	DateOrder: "YMD",
})
