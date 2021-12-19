// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var be_Locale = merge(nil, LocaleData{
	Name:      "be",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|", "каля", "у", "і"},
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)^гадзіна(\z|[^\pL\pM\d]|_)`), "${1}1 гадзіна${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)гадзіну(\z|[^\pL\pM\d]|_)`), "${1}1 гадзіну${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)^хвіліну(\z|[^\pL\pM\d]|_)`), "${1}1 хвіліну${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)^секунду(\z|[^\pL\pM\d]|_)`), "${1}1 секунду${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)некалькі секунд(\z|[^\pL\pM\d]|_)`), "${1}44 секунды${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)некалькі хвілін(\z|[^\pL\pM\d]|_)`), "${1}2 хвіліны${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+)\s*гадзін\s(\d+)\s*хвілін(\z|[^\pL\pM\d]|_)`), "${1}${2}:${3}${4}"},
	},
	Translations: map[string]string{
		"кастрычніка": "october",
		"таму назад":  "ago",
		"на працягу":  "in",
		"панядзелак":  "monday",
		"кастрычнік":  "october",
		"красавіка":   "april",
		"лістапада":   "november",
		"красавік":    "april",
		"студзень":    "january",
		"студзеня":    "january",
		"сакавіка":    "march",
		"лістапад":    "november",
		"верасень":    "september",
		"жнівень":     "august",
		"жнівеня":     "august",
		"снежань":     "december",
		"пятніца":     "friday",
		"гадзіна":     "hour",
		"гадзіну":     "hour",
		"гадзіны":     "hour",
		"чэрвень":     "june",
		"чэрвеня":     "june",
		"сакавік":     "march",
		"хвіліна":     "minute",
		"хвіліну":     "minute",
		"хвіліны":     "minute",
		"месяцау":     "month",
		"секунда":     "second",
		"секунду":     "second",
		"секунды":     "second",
		"верасня":     "september",
		"нядзеля":     "sunday",
		"ауторак":     "tuesday",
		"тыдзень":     "week",
		"жніуня":      "august",
		"снежня":      "december",
		"лютага":      "february",
		"гадзін":      "hour",
		"ліпень":      "july",
		"ліпеня":      "july",
		"трауня":      "may",
		"хвілін":      "minute",
		"месяца":      "month",
		"месяцы":      "month",
		"субота":      "saturday",
		"секунд":      "second",
		"чацвер":      "thursday",
		"серада":      "wednesday",
		"тыдняу":      "week",
		"назад":       "ago",
		"дзень":       "day",
		"месяц":       "month",
		"тыдня":       "week",
		"тыдні":       "week",
		"гадоу":       "year",
		"каля":        "",
		"таму":        "ago",
		"дзен":        "day",
		"люты":        "february",
		"гадз":        "hour",
		"хвіл":        "minute",
		"гады":        "year",
		"года":        "year",
		"кра":         "april",
		"крс":         "april",
		"жнв":         "august",
		"жні":         "august",
		"дні":         "day",
		"сне":         "december",
		"снж":         "december",
		"лют":         "february",
		"пят":         "friday",
		"gmt":         "gmt",
		"стд":         "january",
		"сту":         "january",
		"ліп":         "july",
		"чэр":         "june",
		"сак":         "march",
		"маи":         "may",
		"мая":         "may",
		"тра":         "may",
		"пнд":         "monday",
		"мес":         "month",
		"ліс":         "november",
		"кас":         "october",
		"кст":         "october",
		"суб":         "saturday",
		"сек":         "second",
		"вер":         "september",
		"врс":         "september",
		"няд":         "sunday",
		"чцв":         "thursday",
		"аут":         "tuesday",
		"utc":         "utc",
		"тыд":         "week",
		"год":         "year",
		"am":          "am",
		"пт":          "friday",
		"хв":          "minute",
		"пн":          "monday",
		"pm":          "pm",
		"сб":          "saturday",
		"нд":          "sunday",
		"чв":          "thursday",
		"чц":          "thursday",
		"ау":          "tuesday",
		"ср":          "wednesday",
		"'":           "",
		",":           "",
		";":           "",
		"@":           "",
		"[":           "",
		"]":           "",
		"|":           "",
		"у":           "",
		"і":           "",
		" ":           " ",
		"+":           "+",
		"-":           "-",
		".":           ".",
		"/":           "/",
		":":           ":",
		"д":           "day",
		"с":           "second",
		"г":           "year",
		"z":           "z",
	},
	RelativeType: map[string]string{
		"у наступным месяцы": "in 1 month",
		"на наступным тыдні": "in 1 week",
		"у наступным годзе":  "in 1 year",
		"у мінулым месяцы":   "1 month ago",
		"на мінулым тыдні":   "1 week ago",
		"у мінулым годзе":    "1 year ago",
		"у гэту гадзіну":     "0 hour ago",
		"у гэту хвіліну":     "0 minute ago",
		"у гэтым месяцы":     "0 month ago",
		"на гэтым тыдні":     "0 week ago",
		"у гэтым годзе":      "0 year ago",
		"пазаучора":          "2 day ago",
		"заутра":             "in 1 day",
		"сення":              "0 day ago",
		"цяпер":              "0 second ago",
		"учора":              "1 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) гадзіну таму`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) гадзіны таму`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) хвіліну таму`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) хвіліны таму`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) секунду таму`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) секунды таму`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) тыдзень таму`), "$1 week ago"},
		{regexp.MustCompile(`(?i)праз (\d+) гадзіну`), "in $1 hour"},
		{regexp.MustCompile(`(?i)праз (\d+) гадзіны`), "in $1 hour"},
		{regexp.MustCompile(`(?i)праз (\d+) хвіліну`), "in $1 minute"},
		{regexp.MustCompile(`(?i)праз (\d+) хвіліны`), "in $1 minute"},
		{regexp.MustCompile(`(?i)праз (\d+) секунду`), "in $1 second"},
		{regexp.MustCompile(`(?i)праз (\d+) секунды`), "in $1 second"},
		{regexp.MustCompile(`(?i)праз (\d+) тыдзень`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) месяца таму`), "$1 month ago"},
		{regexp.MustCompile(`(?i)праз (\d+) месяца`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) дзень таму`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) месяц таму`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) тыдня таму`), "$1 week ago"},
		{regexp.MustCompile(`(?i)праз (\d+) дзень`), "in $1 day"},
		{regexp.MustCompile(`(?i)праз (\d+) месяц`), "in $1 month"},
		{regexp.MustCompile(`(?i)праз (\d+) тыдня`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) гадз таму`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) года таму`), "$1 year ago"},
		{regexp.MustCompile(`(?i)праз (\d+) гадз`), "in $1 hour"},
		{regexp.MustCompile(`(?i)праз (\d+) года`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) дня таму`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) мес таму`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) тыд таму`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) год таму`), "$1 year ago"},
		{regexp.MustCompile(`(?i)праз (\d+) дня`), "in $1 day"},
		{regexp.MustCompile(`(?i)праз (\d+) мес`), "in $1 month"},
		{regexp.MustCompile(`(?i)праз (\d+) тыд`), "in $1 week"},
		{regexp.MustCompile(`(?i)праз (\d+) год`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) хв таму`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)праз (\d+) хв`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) д таму`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) с таму`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) г таму`), "$1 year ago"},
		{regexp.MustCompile(`(?i)праз (\d+) д`), "in $1 day"},
		{regexp.MustCompile(`(?i)праз (\d+) с`), "in $1 second"},
		{regexp.MustCompile(`(?i)праз (\d+) г`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ гадзіну таму|\d+ гадзіны таму|\d+ секунду таму|\d+ секунды таму|\d+ тыдзень таму|\d+ хвіліну таму|\d+ хвіліны таму|праз \d+ гадзіну|праз \d+ гадзіны|праз \d+ секунду|праз \d+ секунды|праз \d+ тыдзень|праз \d+ хвіліну|праз \d+ хвіліны|\d+ месяца таму|праз \d+ месяца|\d+ дзень таму|\d+ месяц таму|\d+ тыдня таму|праз \d+ дзень|праз \d+ месяц|праз \d+ тыдня|\d+ гадз таму|\d+ года таму|праз \d+ гадз|праз \d+ года|\d+ год таму|\d+ дня таму|\d+ мес таму|\d+ тыд таму|праз \d+ год|праз \d+ дня|праз \d+ мес|праз \d+ тыд|\d+ хв таму|праз \d+ хв|\d+ г таму|\d+ д таму|\d+ с таму|праз \d+ г|праз \d+ д|праз \d+ с)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ гадзіну таму|\d+ гадзіны таму|\d+ секунду таму|\d+ секунды таму|\d+ тыдзень таму|\d+ хвіліну таму|\d+ хвіліны таму|праз \d+ гадзіну|праз \d+ гадзіны|праз \d+ секунду|праз \d+ секунды|праз \d+ тыдзень|праз \d+ хвіліну|праз \d+ хвіліны|\d+ месяца таму|праз \d+ месяца|\d+ дзень таму|\d+ месяц таму|\d+ тыдня таму|праз \d+ дзень|праз \d+ месяц|праз \d+ тыдня|\d+ гадз таму|\d+ года таму|праз \d+ гадз|праз \d+ года|\d+ год таму|\d+ дня таму|\d+ мес таму|\d+ тыд таму|праз \d+ год|праз \d+ дня|праз \d+ мес|праз \d+ тыд|\d+ хв таму|праз \d+ хв|\d+ г таму|\d+ д таму|\d+ с таму|праз \d+ г|праз \d+ д|праз \d+ с)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(на наступным тыдні|у наступным месяцы|у наступным годзе|на мінулым тыдні|у мінулым месяцы|у мінулым годзе|на гэтым тыдні|у гэту гадзіну|у гэту хвіліну|у гэтым месяцы|у гэтым годзе|кастрычніка|кастрычнік|на працягу|панядзелак|таму назад|красавіка|лістапада|пазаучора|верасень|красавік|лістапад|сакавіка|студзень|студзеня|ауторак|верасня|гадзіна|гадзіну|гадзіны|жнівень|жнівеня|месяцау|нядзеля|пятніца|сакавік|секунда|секунду|секунды|снежань|тыдзень|хвіліна|хвіліну|хвіліны|чэрвень|чэрвеня|гадзін|жніуня|заутра|лютага|ліпень|ліпеня|месяца|месяцы|секунд|серада|снежня|субота|трауня|тыдняу|хвілін|чацвер|гадоу|дзень|месяц|назад|сення|тыдня|тыдні|учора|цяпер|гадз|гады|года|дзен|каля|люты|таму|хвіл|gmt|utc|аут|вер|врс|год|дні|жнв|жні|кас|кра|крс|кст|лют|ліп|ліс|маи|мая|мес|няд|пнд|пят|сак|сек|сне|снж|стд|сту|суб|тра|тыд|чцв|чэр|\+|\.|\[|\]|\||am|pm|ау|нд|пн|пт|сб|ср|хв|чв|чц| |'|,|-|/|:|;|@|z|г|д|с|у|і)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
