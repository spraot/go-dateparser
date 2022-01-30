// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var fa_Locale = merge(nil, LocaleData{
	Name:                  "fa",
	DateOrder:             "YMD",
	Charset:               []rune(`cgtuzآئابتثجدذرزسشظعفقلمنهوئپچژکگی‌`),
	SentenceSplitterGroup: 6,
	Translations: map[string]string{
		"قبل‌ازظهر": "am",
		"چهار شنبه": "wednesday",
		"بعدازظهر":  "pm",
		"روز شنبه":  "saturday",
		"پنج شنبه":  "thursday",
		"چهارشنبه":  "wednesday",
		"دو شنبه":   "saturday",
		"سپتامبر":   "september",
		"پنجشنبه":   "thursday",
		"سه‌شنبه":   "tuesday",
		"دسامبر":    "december",
		"فبروری":    "february",
		"ژانویه":    "january",
		"دوشنبه":    "monday",
		"نوامبر":    "november",
		"اکتوبر":    "october",
		"سپتمبر":    "september",
		"یکشنبه":    "sunday",
		"سهشنبه":    "tuesday",
		"اوریل":     "april",
		"اپریل":     "april",
		"دسمبر":     "december",
		"فوریه":     "february",
		"جنوری":     "january",
		"جولای":     "july",
		"ژويیه":     "july",
		"دقیقه":     "minute",
		"نومبر":     "november",
		"اکتبر":     "october",
		"ثانیه":     "second",
		"اگست":      "august",
		"جمعه":      "friday",
		"ساعت":      "hour",
		"ژوين":      "june",
		"مارس":      "march",
		"مارچ":      "march",
		"شنبه":      "saturday",
		"هفته":      "week",
		"پیش":       "ago",
		"اوت":       "august",
		"روز":       "day",
		"gmt":       "gmt",
		"جون":       "june",
		"ماه":       "month",
		"دوم":       "second",
		"utc":       "utc",
		"سال":       "year",
		"am":        "am",
		"قظ":        "am",
		"در":        "in",
		"مه":        "may",
		"می":        "may",
		"pm":        "pm",
		"بظ":        "pm",
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
		"د":         "saturday",
		"z":         "z",
	},
	RelativeType: map[string]string{
		"همین دقیقه": "0 minute ago",
		"هفته گذشته": "1 week ago",
		"هفته اینده": "in 1 week",
		"همین ساعت":  "0 hour ago",
		"ماه گذشته":  "1 month ago",
		"سال گذشته":  "1 year ago",
		"ماه اینده":  "in 1 month",
		"سال اینده":  "in 1 year",
		"این هفته":   "0 week ago",
		"این ماه":    "0 month ago",
		"ماه پیش":    "1 month ago",
		"امروز":      "0 day ago",
		"اکنون":      "0 second ago",
		"امسال":      "0 year ago",
		"دیروز":      "1 day ago",
		"فردا":       "in 1 day",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) دقیقه پیش`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) ثانیه پیش`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) دقیقه بعد`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) ثانیه بعد`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) ساعت پیش`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) هفته پیش`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ساعت بعد`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) هفته بعد`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) روز پیش`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) ماه پیش`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) سال پیش`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) روز بعد`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) ماه بعد`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) سال بعد`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ ثانیه بعد|\d+ ثانیه پیش|\d+ دقیقه بعد|\d+ دقیقه پیش|\d+ ساعت بعد|\d+ ساعت پیش|\d+ هفته بعد|\d+ هفته پیش|\d+ روز بعد|\d+ روز پیش|\d+ سال بعد|\d+ سال پیش|\d+ ماه بعد|\d+ ماه پیش)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ ثانیه بعد|\d+ ثانیه پیش|\d+ دقیقه بعد|\d+ دقیقه پیش|\d+ ساعت بعد|\d+ ساعت پیش|\d+ هفته بعد|\d+ هفته پیش|\d+ روز بعد|\d+ روز پیش|\d+ سال بعد|\d+ سال پیش|\d+ ماه بعد|\d+ ماه پیش)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(هفته اینده|هفته گذشته|همین دقیقه|سال اینده|سال گذشته|قبل‌ازظهر|ماه اینده|ماه گذشته|همین ساعت|چهار شنبه|این هفته|بعدازظهر|روز شنبه|پنج شنبه|چهارشنبه|این ماه|دو شنبه|سه‌شنبه|سپتامبر|ماه پیش|پنجشنبه|اکتوبر|دسامبر|دوشنبه|سهشنبه|سپتمبر|فبروری|نوامبر|ژانویه|یکشنبه|امروز|امسال|اوریل|اپریل|اکتبر|اکنون|ثانیه|جنوری|جولای|دسمبر|دقیقه|دیروز|فوریه|نومبر|ژويیه|اگست|جمعه|ساعت|شنبه|فردا|مارس|مارچ|هفته|ژوين|gmt|utc|اوت|جون|دوم|روز|سال|ماه|پیش|\+|\.|\[|\]|\||am|pm|بظ|در|قظ|مه|می| |'|,|-|/|:|;|@|z|د)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var fa_AF_Locale = merge(&fa_Locale, LocaleData{
	Name:      "fa-AF",
	DateOrder: "YMD",
	Translations: map[string]string{
		"دسم": "december",
		"جنو": "january",
		"جول": "july",
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ ثانیه بعد|\d+ ثانیه پیش|\d+ دقیقه بعد|\d+ دقیقه پیش|\d+ ساعت بعد|\d+ ساعت پیش|\d+ هفته بعد|\d+ هفته پیش|\d+ روز بعد|\d+ روز پیش|\d+ سال بعد|\d+ سال پیش|\d+ ماه بعد|\d+ ماه پیش)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ ثانیه بعد|\d+ ثانیه پیش|\d+ دقیقه بعد|\d+ دقیقه پیش|\d+ ساعت بعد|\d+ ساعت پیش|\d+ هفته بعد|\d+ هفته پیش|\d+ روز بعد|\d+ روز پیش|\d+ سال بعد|\d+ سال پیش|\d+ ماه بعد|\d+ ماه پیش)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(هفته اینده|هفته گذشته|همین دقیقه|سال اینده|سال گذشته|قبل‌ازظهر|ماه اینده|ماه گذشته|همین ساعت|چهار شنبه|این هفته|بعدازظهر|روز شنبه|پنج شنبه|چهارشنبه|این ماه|دو شنبه|سه‌شنبه|سپتامبر|ماه پیش|پنجشنبه|اکتوبر|دسامبر|دوشنبه|سهشنبه|سپتمبر|فبروری|نوامبر|ژانویه|یکشنبه|امروز|امسال|اوریل|اپریل|اکتبر|اکنون|ثانیه|جنوری|جولای|دسمبر|دقیقه|دیروز|فوریه|نومبر|ژويیه|اگست|جمعه|ساعت|شنبه|فردا|مارس|مارچ|هفته|ژوين|gmt|utc|اوت|جنو|جول|جون|دسم|دوم|روز|سال|ماه|پیش|\+|\.|\[|\]|\||am|pm|بظ|در|قظ|مه|می| |'|,|-|/|:|;|@|z|د)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
