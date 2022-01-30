// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var pl_Locale = merge(nil, LocaleData{
	Name:                  "pl",
	DateOrder:             "DMY",
	Charset:               []rune(`bcdegijklnorstuwyząęłńśź`),
	Abbreviations:         []string{"r"},
	SentenceSplitterGroup: 1,
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)dzis(\z|[^\pL\pM\d]|_)`), "${1}0 dnia${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)dzisiaj(\z|[^\pL\pM\d]|_)`), "${1}0 dnia${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)onegdaj(\z|[^\pL\pM\d]|_)`), "${1}2 dnia${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)przedwczoraj(\z|[^\pL\pM\d]|_)`), "${1}2 dnia temu${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)jutro(\z|[^\pL\pM\d]|_)`), "${1}za 1 dnia${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)pojutrze(\z|[^\pL\pM\d]|_)`), "${1}za 2 dnia${2}"},
	},
	Translations: map[string]string{
		"poniedzialek": "monday",
		"poniedziałek": "monday",
		"pazdziernika": "october",
		"pazdzierniku": "october",
		"listopadzie":  "november",
		"pazdziernik":  "october",
		"listopada":    "november",
		"niedziela":    "sunday",
		"niedziele":    "sunday",
		"kwiecien":     "april",
		"kwietnia":     "april",
		"kwietniu":     "april",
		"sierpien":     "august",
		"sierpnia":     "august",
		"sierpniu":     "august",
		"grudzien":     "december",
		"stycznia":     "january",
		"styczniu":     "january",
		"czerwiec":     "june",
		"miesiaca":     "month",
		"miesiace":     "month",
		"miesiecy":     "month",
		"listopad":     "november",
		"wrzesien":     "september",
		"wrzesnia":     "september",
		"wrzesniu":     "september",
		"czwartek":     "thursday",
		"tygodnie":     "week",
		"grudnia":      "december",
		"grudniu":      "december",
		"godzina":      "hour",
		"godzine":      "hour",
		"godziny":      "hour",
		"styczen":      "january",
		"czerwca":      "june",
		"czerwcu":      "june",
		"miesiac":      "month",
		"sekunda":      "second",
		"sekunde":      "second",
		"sekundy":      "second",
		"tydzien":      "week",
		"tygodni":      "week",
		"lutego":       "february",
		"piatek":       "friday",
		"godzin":       "hour",
		"lipiec":       "july",
		"marzec":       "march",
		"minuta":       "minute",
		"minute":       "minute",
		"minuty":       "minute",
		"sobota":       "saturday",
		"sobote":       "saturday",
		"sekund":       "second",
		"wtorek":       "tuesday",
		"dzien":        "day",
		"lutym":        "february",
		"lipca":        "july",
		"lipcu":        "july",
		"marca":        "march",
		"marcu":        "march",
		"minut":        "minute",
		"niedz":        "sunday",
		"sroda":        "wednesday",
		"srode":        "wednesday",
		"roku":         "",
		"temu":         "ago",
		"kwie":         "april",
		"dnia":         "day",
		"dniu":         "day",
		"luty":         "february",
		"godz":         "hour",
		"maja":         "may",
		"maju":         "may",
		"mies":         "month",
		"tydz":         "week",
		"lata":         "year",
		"kwi":          "april",
		"sie":          "august",
		"dni":          "day",
		"gru":          "december",
		"lut":          "february",
		"pia":          "friday",
		"gmt":          "gmt",
		"sty":          "january",
		"lip":          "july",
		"cze":          "june",
		"mar":          "march",
		"maj":          "may",
		"min":          "minute",
		"pon":          "monday",
		"lis":          "november",
		"paz":          "october",
		"sob":          "saturday",
		"sek":          "second",
		"wrz":          "september",
		"nie":          "sunday",
		"czw":          "thursday",
		"wto":          "tuesday",
		"utc":          "utc",
		"sro":          "wednesday",
		"lat":          "year",
		"rok":          "year",
		"r.":           "",
		"am":           "am",
		"pi":           "friday",
		"pt":           "friday",
		"za":           "in",
		"pn":           "monday",
		"mc":           "month",
		"pm":           "pm",
		"sb":           "saturday",
		"so":           "saturday",
		"nd":           "sunday",
		"cz":           "thursday",
		"wt":           "tuesday",
		"sr":           "wednesday",
		"'":            "",
		",":            "",
		";":            "",
		"@":            "",
		"[":            "",
		"]":            "",
		"i":            "",
		"o":            "",
		"w":            "",
		"|":            "",
		" ":            " ",
		"+":            "+",
		"-":            "-",
		".":            ".",
		"/":            "/",
		":":            ":",
		"g":            "hour",
		"s":            "second",
		"r":            "year",
		"z":            "z",
	},
	RelativeType: map[string]string{
		"w przyszłym miesiacu": "in 1 month",
		"w przyszłym tygodniu": "in 1 week",
		"w zeszłym miesiacu":   "1 month ago",
		"w zeszłym tygodniu":   "1 week ago",
		"w przyszłym roku":     "in 1 year",
		"w tym miesiacu":       "0 month ago",
		"w tym tygodniu":       "0 week ago",
		"w zeszłym roku":       "1 year ago",
		"ta godzina":           "0 hour ago",
		"w tym roku":           "0 year ago",
		"ta minuta":            "0 minute ago",
		"dzisiaj":              "0 day ago",
		"wczoraj":              "1 day ago",
		"teraz":                "0 second ago",
		"jutro":                "in 1 day",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) miesiaca temu`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) tygodnia temu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) godzine temu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) godziny temu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) miesiac temu`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) sekunde temu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) sekundy temu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) tydzien temu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) minute temu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) minuty temu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)za (\d+) miesiaca`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) tygodnia`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) dzien temu`), "$1 day ago"},
		{regexp.MustCompile(`(?i)za (\d+) godzine`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) godziny`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) miesiac`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) sekunde`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) sekundy`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) tydzien`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) dnia temu`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) godz temu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) mies temu`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) tydz temu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) roku temu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)za (\d+) minute`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) minuty`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) min temu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) sek temu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) tyg temu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) rok temu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)za (\d+) dzien`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) dnia`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) godz`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) mies`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) tydz`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) roku`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) g temu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) s temu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)za (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) tyg`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) rok`), "in $1 year"},
		{regexp.MustCompile(`(?i)–(\d+) mies`), "$1 month ago"},
		{regexp.MustCompile(`(?i)za (\d+) g`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) s`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ miesiaca temu|\d+ tygodnia temu|\d+ godzine temu|\d+ godziny temu|\d+ miesiac temu|\d+ sekunde temu|\d+ sekundy temu|\d+ tydzien temu|\d+ minute temu|\d+ minuty temu|za \d+ miesiaca|za \d+ tygodnia|\d+ dzien temu|za \d+ godzine|za \d+ godziny|za \d+ miesiac|za \d+ sekunde|za \d+ sekundy|za \d+ tydzien|\d+ dnia temu|\d+ godz temu|\d+ mies temu|\d+ roku temu|\d+ tydz temu|za \d+ minute|za \d+ minuty|\d+ min temu|\d+ rok temu|\d+ sek temu|\d+ tyg temu|za \d+ dzien|za \d+ dnia|za \d+ godz|za \d+ mies|za \d+ roku|za \d+ tydz|\d+ g temu|\d+ s temu|za \d+ min|za \d+ rok|za \d+ sek|za \d+ tyg|–\d+ mies|za \d+ g|za \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ miesiaca temu|\d+ tygodnia temu|\d+ godzine temu|\d+ godziny temu|\d+ miesiac temu|\d+ sekunde temu|\d+ sekundy temu|\d+ tydzien temu|\d+ minute temu|\d+ minuty temu|za \d+ miesiaca|za \d+ tygodnia|\d+ dzien temu|za \d+ godzine|za \d+ godziny|za \d+ miesiac|za \d+ sekunde|za \d+ sekundy|za \d+ tydzien|\d+ dnia temu|\d+ godz temu|\d+ mies temu|\d+ roku temu|\d+ tydz temu|za \d+ minute|za \d+ minuty|\d+ min temu|\d+ rok temu|\d+ sek temu|\d+ tyg temu|za \d+ dzien|za \d+ dnia|za \d+ godz|za \d+ mies|za \d+ roku|za \d+ tydz|\d+ g temu|\d+ s temu|za \d+ min|za \d+ rok|za \d+ sek|za \d+ tyg|–\d+ mies|za \d+ g|za \d+ s)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(w przyszłym miesiacu|w przyszłym tygodniu|w zeszłym miesiacu|w zeszłym tygodniu|w przyszłym roku|w tym miesiacu|w tym tygodniu|w zeszłym roku|pazdziernika|pazdzierniku|poniedzialek|poniedziałek|listopadzie|pazdziernik|ta godzina|w tym roku|listopada|niedziela|niedziele|ta minuta|czerwiec|czwartek|grudzien|kwiecien|kwietnia|kwietniu|listopad|miesiaca|miesiace|miesiecy|sierpien|sierpnia|sierpniu|stycznia|styczniu|tygodnie|wrzesien|wrzesnia|wrzesniu|czerwca|czerwcu|dzisiaj|godzina|godzine|godziny|grudnia|grudniu|miesiac|sekunda|sekunde|sekundy|styczen|tydzien|tygodni|wczoraj|godzin|lipiec|lutego|marzec|minuta|minute|minuty|piatek|sekund|sobota|sobote|wtorek|dzien|jutro|lipca|lipcu|lutym|marca|marcu|minut|niedz|sroda|srode|teraz|dnia|dniu|godz|kwie|lata|luty|maja|maju|mies|roku|temu|tydz|cze|czw|dni|gmt|gru|kwi|lat|lip|lis|lut|maj|mar|min|nie|paz|pia|pon|r\.|rok|sek|sie|sob|sro|sty|utc|wrz|wto|\+|\.|\[|\]|\||am|cz|mc|nd|pi|pm|pn|pt|sb|so|sr|wt|za| |'|,|-|/|:|;|@|g|i|o|r|s|w|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
