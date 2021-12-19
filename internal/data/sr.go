// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sr_Locale = merge(nil, LocaleData{
	Name:      "sr",
	DateOrder: "DMY.",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|"},
	Translations: map[string]string{
		"пре подне": "am",
		"понедељак": "monday",
		"септембар": "september",
		"децембар":  "december",
		"новембар":  "november",
		"по подне":  "pm",
		"четвртак":  "thursday",
		"фебруар":   "february",
		"октобар":   "october",
		"август":    "august",
		"јануар":    "january",
		"субота":    "saturday",
		"секунд":    "second",
		"недеља":    "sunday",
		"уторак":    "tuesday",
		"година":    "year",
		"април":     "april",
		"петак":     "friday",
		"минут":     "minute",
		"месец":     "month",
		"среда":     "wednesday",
		"март":      "march",
		"апр":       "april",
		"авг":       "august",
		"дан":       "day",
		"дец":       "december",
		"феб":       "february",
		"пет":       "friday",
		"gmt":       "gmt",
		"сат":       "hour",
		"јан":       "january",
		"јул":       "july",
		"јун":       "june",
		"мар":       "march",
		"мај":       "may",
		"мин":       "minute",
		"пон":       "monday",
		"мес":       "month",
		"нов":       "november",
		"окт":       "october",
		"суб":       "saturday",
		"сек":       "second",
		"сеп":       "september",
		"нед":       "sunday",
		"чет":       "thursday",
		"уто":       "tuesday",
		"utc":       "utc",
		"сре":       "wednesday",
		"год":       "year",
		"am":        "am",
		"pm":        "pm",
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
		"д":         "day",
		"ч":         "hour",
		"м":         "month",
		"с":         "second",
		"н":         "week",
		"г":         "year",
		"z":         "z",
	},
	RelativeType: map[string]string{
		"следећег месеца": "in 1 month",
		"прошлог месеца":  "1 month ago",
		"следеће недеље":  "in 1 week",
		"следеће године":  "in 1 year",
		"прошле недеље":   "1 week ago",
		"прошле године":   "1 year ago",
		"овог минута":     "0 minute ago",
		"овог месеца":     "0 month ago",
		"ове недеље":      "0 week ago",
		"ове године":      "0 year ago",
		"овог сата":       "0 hour ago",
		"данас":           "0 day ago",
		"сутра":           "in 1 day",
		"сада":            "0 second ago",
		"јуче":            "1 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)пре (\d+) секунде`), "$1 second ago"},
		{regexp.MustCompile(`(?i)пре (\d+) секунди`), "$1 second ago"},
		{regexp.MustCompile(`(?i)пре (\d+) минута`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)пре (\d+) месеца`), "$1 month ago"},
		{regexp.MustCompile(`(?i)пре (\d+) месеци`), "$1 month ago"},
		{regexp.MustCompile(`(?i)пре (\d+) недеља`), "$1 week ago"},
		{regexp.MustCompile(`(?i)пре (\d+) недеље`), "$1 week ago"},
		{regexp.MustCompile(`(?i)пре (\d+) година`), "$1 year ago"},
		{regexp.MustCompile(`(?i)пре (\d+) године`), "$1 year ago"},
		{regexp.MustCompile(`(?i)за (\d+) секунди`), "in $1 second"},
		{regexp.MustCompile(`(?i)за (\d+) секунду`), "in $1 second"},
		{regexp.MustCompile(`(?i)за (\d+) минута`), "in $1 minute"},
		{regexp.MustCompile(`(?i)за (\d+) месеци`), "in $1 month"},
		{regexp.MustCompile(`(?i)за (\d+) недеља`), "in $1 week"},
		{regexp.MustCompile(`(?i)за (\d+) недељу`), "in $1 week"},
		{regexp.MustCompile(`(?i)за (\d+) година`), "in $1 year"},
		{regexp.MustCompile(`(?i)за (\d+) годину`), "in $1 year"},
		{regexp.MustCompile(`(?i)пре (\d+) дана`), "$1 day ago"},
		{regexp.MustCompile(`(?i)пре (\d+) сата`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)пре (\d+) сати`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)за (\d+) минут`), "in $1 minute"},
		{regexp.MustCompile(`(?i)за (\d+) месец`), "in $1 month"},
		{regexp.MustCompile(`(?i)пре (\d+) мин`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)пре (\d+) мес`), "$1 month ago"},
		{regexp.MustCompile(`(?i)пре (\d+) сек`), "$1 second ago"},
		{regexp.MustCompile(`(?i)пре (\d+) нед`), "$1 week ago"},
		{regexp.MustCompile(`(?i)пре (\d+) год`), "$1 year ago"},
		{regexp.MustCompile(`(?i)за (\d+) дана`), "in $1 day"},
		{regexp.MustCompile(`(?i)за (\d+) сати`), "in $1 hour"},
		{regexp.MustCompile(`(?i)за (\d+) дан`), "in $1 day"},
		{regexp.MustCompile(`(?i)за (\d+) сат`), "in $1 hour"},
		{regexp.MustCompile(`(?i)за (\d+) мин`), "in $1 minute"},
		{regexp.MustCompile(`(?i)за (\d+) мес`), "in $1 month"},
		{regexp.MustCompile(`(?i)за (\d+) сек`), "in $1 second"},
		{regexp.MustCompile(`(?i)за (\d+) нед`), "in $1 week"},
		{regexp.MustCompile(`(?i)за (\d+) год`), "in $1 year"},
		{regexp.MustCompile(`(?i)пре (\d+) д`), "$1 day ago"},
		{regexp.MustCompile(`(?i)пре (\d+) ч`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)пре (\d+) м`), "$1 month ago"},
		{regexp.MustCompile(`(?i)пре (\d+) с`), "$1 second ago"},
		{regexp.MustCompile(`(?i)пре (\d+) н`), "$1 week ago"},
		{regexp.MustCompile(`(?i)пре (\d+) г`), "$1 year ago"},
		{regexp.MustCompile(`(?i)за (\d+) д`), "in $1 day"},
		{regexp.MustCompile(`(?i)за (\d+) ч`), "in $1 hour"},
		{regexp.MustCompile(`(?i)за (\d+) м`), "in $1 month"},
		{regexp.MustCompile(`(?i)за (\d+) с`), "in $1 second"},
		{regexp.MustCompile(`(?i)за (\d+) н`), "in $1 week"},
		{regexp.MustCompile(`(?i)за (\d+) г`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(пре \d+ секунде|пре \d+ секунди|за \d+ секунди|за \d+ секунду|пре \d+ година|пре \d+ године|пре \d+ месеца|пре \d+ месеци|пре \d+ минута|пре \d+ недеља|пре \d+ недеље|за \d+ година|за \d+ годину|за \d+ месеци|за \d+ минута|за \d+ недеља|за \d+ недељу|за \d+ месец|за \d+ минут|пре \d+ дана|пре \d+ сата|пре \d+ сати|за \d+ дана|за \d+ сати|пре \d+ год|пре \d+ мес|пре \d+ мин|пре \d+ нед|пре \d+ сек|за \d+ год|за \d+ дан|за \d+ мес|за \d+ мин|за \d+ нед|за \d+ сат|за \d+ сек|пре \d+ г|пре \d+ д|пре \d+ м|пре \d+ н|пре \d+ с|пре \d+ ч|за \d+ г|за \d+ д|за \d+ м|за \d+ н|за \d+ с|за \d+ ч)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(пре \d+ секунде|пре \d+ секунди|за \d+ секунди|за \d+ секунду|пре \d+ година|пре \d+ године|пре \d+ месеца|пре \d+ месеци|пре \d+ минута|пре \d+ недеља|пре \d+ недеље|за \d+ година|за \d+ годину|за \d+ месеци|за \d+ минута|за \d+ недеља|за \d+ недељу|за \d+ месец|за \d+ минут|пре \d+ дана|пре \d+ сата|пре \d+ сати|за \d+ дана|за \d+ сати|пре \d+ год|пре \d+ мес|пре \d+ мин|пре \d+ нед|пре \d+ сек|за \d+ год|за \d+ дан|за \d+ мес|за \d+ мин|за \d+ нед|за \d+ сат|за \d+ сек|пре \d+ г|пре \d+ д|пре \d+ м|пре \d+ н|пре \d+ с|пре \d+ ч|за \d+ г|за \d+ д|за \d+ м|за \d+ н|за \d+ с|за \d+ ч)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(следећег месеца|прошлог месеца|следеће године|следеће недеље|прошле године|прошле недеље|овог месеца|овог минута|ове године|ове недеље|овог сата|понедељак|пре подне|септембар|децембар|новембар|по подне|четвртак|октобар|фебруар|август|година|недеља|секунд|субота|уторак|јануар|април|данас|месец|минут|петак|среда|сутра|март|сада|јуче|gmt|utc|авг|апр|год|дан|дец|мар|мај|мес|мин|нед|нов|окт|пет|пон|сат|сек|сеп|сре|суб|уто|феб|чет|јан|јул|јун|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|z|г|д|м|н|с|ч)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
