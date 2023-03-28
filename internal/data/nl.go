// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var (
	nl_Locale    LocaleData
	nl_AW_Locale LocaleData
	nl_BE_Locale LocaleData
	nl_BQ_Locale LocaleData
	nl_CW_Locale LocaleData
	nl_SR_Locale LocaleData
	nl_SX_Locale LocaleData
)

func init() {
	nl_Locale = merge(nil, LocaleData{
		Name:                  "nl",
		DateOrder:             "DMY",
		Charset:               []rune(`bcdefgijklnorstuvwz`),
		SentenceSplitterGroup: 1,
		Translations: map[string][]string{
			"donderdag": {"thursday"},
			"september": {"september"},
			"augustus":  {"august"},
			"december":  {"december"},
			"februari":  {"february"},
			"november":  {"november"},
			"seconden":  {"second"},
			"woensdag":  {"wednesday"},
			"zaterdag":  {"saturday"},
			"dinsdag":   {"tuesday"},
			"geleden":   {"ago"},
			"januari":   {"january"},
			"maandag":   {"monday"},
			"maanden":   {"month"},
			"minuten":   {"minute"},
			"oktober":   {"october"},
			"seconde":   {"second"},
			"vrijdag":   {"friday"},
			"minuut":    {"minute"},
			"zondag":    {"sunday"},
			"april":     {"april"},
			"dagen":     {"day"},
			"maand":     {"month"},
			"maart":     {"march"},
			"weken":     {"week"},
			"jaar":      {"year"},
			"juli":      {"july"},
			"juni":      {"june"},
			"week":      {"week"},
			"apr":       {"april"},
			"aug":       {"august"},
			"dag":       {"day"},
			"dec":       {"december"},
			"feb":       {"february"},
			"gmt":       {"gmt"},
			"jan":       {"january"},
			"jul":       {"july"},
			"jun":       {"june"},
			"mei":       {"may"},
			"min":       {"minute"},
			"mnd":       {"month"},
			"mrt":       {"march"},
			"nov":       {"november"},
			"okt":       {"october"},
			"sec":       {"second"},
			"sep":       {"september"},
			"utc":       {"utc"},
			"uur":       {"hour"},
			"am":        {"am"},
			"di":        {"tuesday"},
			"do":        {"thursday"},
			"in":        {"in"},
			"jr":        {"year"},
			"ma":        {"monday"},
			"om":        {""},
			"pm":        {"pm"},
			"vr":        {"friday"},
			"wk":        {"week"},
			"wo":        {"wednesday"},
			"za":        {"saturday"},
			"zo":        {"sunday"},
			" ":         {" "},
			"'":         {""},
			"+":         {"+"},
			",":         {""},
			"-":         {"-"},
			".":         {"."},
			"/":         {"/"},
			":":         {":"},
			";":         {""},
			"@":         {""},
			"[":         {""},
			"]":         {""},
			"s":         {"second"},
			"z":         {"z"},
			"|":         {""},
		},
		RelativeType: map[string]string{
			"binnen een minuut": "0 minute ago",
			"binnen een uur":    "0 hour ago",
			"volgende maand":    "in 1 month",
			"volgende week":     "in 1 week",
			"volgend jaar":      "in 1 year",
			"vorige maand":      "1 month ago",
			"eergisteren":       "2 day ago",
			"vorige jaar":       "1 year ago",
			"vorige week":       "1 week ago",
			"deze maand":        "0 month ago",
			"overmorgen":        "in 2 day",
			"vorig jaar":        "1 year ago",
			"deze week":         "0 week ago",
			"dit jaar":          "0 year ago",
			"gisteren":          "1 day ago",
			"vandaag":           "0 day ago",
			"morgen":            "in 1 day",
			"nu":                "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) seconden geleden`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) maanden geleden`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minuten geleden`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) seconde geleden`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minuut geleden`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dagen geleden`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) maand geleden`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) weken geleden`), "$1 week ago"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) seconden`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) jaar geleden`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) week geleden`), "$1 week ago"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) maanden`), "in $1 month"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) minuten`), "in $1 minute"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) seconde`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dag geleden`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dgn geleden`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) min geleden`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sec geleden`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) uur geleden`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) minuut`), "in $1 minute"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) dagen`), "in $1 day"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) maand`), "in $1 month"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) weken`), "in $1 week"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) jaar`), "in $1 year"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) week`), "in $1 week"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) dag`), "in $1 day"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) dgn`), "in $1 day"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) sec`), "in $1 second"},
			{regexp.MustCompile(`(?i)over (\d+[.,]?\d*) uur`), "in $1 hour"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* seconden geleden|\d+[.,]?\d* maanden geleden|\d+[.,]?\d* minuten geleden|\d+[.,]?\d* seconde geleden|\d+[.,]?\d* minuut geleden|\d+[.,]?\d* dagen geleden|\d+[.,]?\d* maand geleden|\d+[.,]?\d* weken geleden|over \d+[.,]?\d* seconden|\d+[.,]?\d* jaar geleden|\d+[.,]?\d* week geleden|over \d+[.,]?\d* maanden|over \d+[.,]?\d* minuten|over \d+[.,]?\d* seconde|\d+[.,]?\d* dag geleden|\d+[.,]?\d* dgn geleden|\d+[.,]?\d* min geleden|\d+[.,]?\d* sec geleden|\d+[.,]?\d* uur geleden|over \d+[.,]?\d* minuut|over \d+[.,]?\d* dagen|over \d+[.,]?\d* maand|over \d+[.,]?\d* weken|over \d+[.,]?\d* jaar|over \d+[.,]?\d* week|over \d+[.,]?\d* dag|over \d+[.,]?\d* dgn|over \d+[.,]?\d* min|over \d+[.,]?\d* sec|over \d+[.,]?\d* uur)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* seconden geleden|\d+[.,]?\d* maanden geleden|\d+[.,]?\d* minuten geleden|\d+[.,]?\d* seconde geleden|\d+[.,]?\d* minuut geleden|\d+[.,]?\d* dagen geleden|\d+[.,]?\d* maand geleden|\d+[.,]?\d* weken geleden|over \d+[.,]?\d* seconden|\d+[.,]?\d* jaar geleden|\d+[.,]?\d* week geleden|over \d+[.,]?\d* maanden|over \d+[.,]?\d* minuten|over \d+[.,]?\d* seconde|\d+[.,]?\d* dag geleden|\d+[.,]?\d* dgn geleden|\d+[.,]?\d* min geleden|\d+[.,]?\d* sec geleden|\d+[.,]?\d* uur geleden|over \d+[.,]?\d* minuut|over \d+[.,]?\d* dagen|over \d+[.,]?\d* maand|over \d+[.,]?\d* weken|over \d+[.,]?\d* jaar|over \d+[.,]?\d* week|over \d+[.,]?\d* dag|over \d+[.,]?\d* dgn|over \d+[.,]?\d* min|over \d+[.,]?\d* sec|over \d+[.,]?\d* uur)$`),
		KnownWords:      []string{"binnen een minuut", "binnen een uur", "volgende maand", "volgende week", "volgend jaar", "vorige maand", "eergisteren", "vorige jaar", "vorige week", "deze maand", "overmorgen", "vorig jaar", "deze week", "donderdag", "september", "augustus", "december", "dit jaar", "februari", "gisteren", "november", "seconden", "woensdag", "zaterdag", "dinsdag", "geleden", "januari", "maandag", "maanden", "minuten", "oktober", "seconde", "vandaag", "vrijdag", "minuut", "morgen", "zondag", "april", "dagen", "maand", "maart", "weken", "jaar", "juli", "juni", "week", "apr", "aug", "dag", "dec", "feb", "gmt", "jan", "jul", "jun", "mei", "min", "mnd", "mrt", "nov", "okt", "sec", "sep", "utc", "uur", "am", "di", "do", "in", "jr", "ma", "nu", "om", "pm", "vr", "wk", "wo", "za", "zo", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "s", "z", "|"},
	})

	nl_AW_Locale = merge(&nl_Locale, LocaleData{
		Name:      "nl-AW",
		DateOrder: "DMY",
	})

	nl_BE_Locale = merge(&nl_Locale, LocaleData{
		Name:      "nl-BE",
		DateOrder: "DMY",
	})

	nl_BQ_Locale = merge(&nl_Locale, LocaleData{
		Name:      "nl-BQ",
		DateOrder: "DMY",
	})

	nl_CW_Locale = merge(&nl_Locale, LocaleData{
		Name:      "nl-CW",
		DateOrder: "DMY",
	})

	nl_SR_Locale = merge(&nl_Locale, LocaleData{
		Name:      "nl-SR",
		DateOrder: "DMY",
	})

	nl_SX_Locale = merge(&nl_Locale, LocaleData{
		Name:      "nl-SX",
		DateOrder: "DMY",
	})
}
