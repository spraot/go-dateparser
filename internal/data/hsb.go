// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var (
	hsb_Locale LocaleData
)

func init() {
	hsb_Locale = merge(nil, LocaleData{
		Name:      "hsb",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefghijklnorstuwyzóčěłńřšź`),
		Translations: map[string][]string{
			"dopołdnja": {"am"},
			"popołdnju": {"pm"},
			"september": {"september"},
			"septembra": {"september"},
			"december":  {"december"},
			"decembra":  {"december"},
			"februara":  {"february"},
			"njedzela":  {"sunday"},
			"nowember":  {"november"},
			"nowembra":  {"november"},
			"pondzela":  {"monday"},
			"awgusta":   {"august"},
			"februar":   {"february"},
			"hodzina":   {"hour"},
			"januara":   {"january"},
			"oktober":   {"october"},
			"oktobra":   {"october"},
			"sekunda":   {"second"},
			"stwortk":   {"thursday"},
			"apryla":    {"april"},
			"awgust":    {"august"},
			"januar":    {"january"},
			"julija":    {"july"},
			"junija":    {"june"},
			"minuta":    {"minute"},
			"sobota":    {"saturday"},
			"srjeda":    {"wednesday"},
			"tydzen":    {"week"},
			"wutora":    {"tuesday"},
			"apryl":     {"april"},
			"julij":     {"july"},
			"junij":     {"june"},
			"merca":     {"march"},
			"mesac":     {"month"},
			"pjatk":     {"friday"},
			"dzen":      {"day"},
			"hodz":      {"hour"},
			"leto":      {"year"},
			"meja":      {"may"},
			"meje":      {"may"},
			"merc":      {"march"},
			"tydz":      {"week"},
			"apr":       {"april"},
			"awg":       {"august"},
			"dec":       {"december"},
			"feb":       {"february"},
			"gmt":       {"gmt"},
			"jan":       {"january"},
			"jul":       {"july"},
			"jun":       {"june"},
			"mej":       {"may"},
			"mer":       {"march"},
			"mes":       {"month"},
			"min":       {"minute"},
			"nje":       {"sunday"},
			"now":       {"november"},
			"okt":       {"october"},
			"pja":       {"friday"},
			"pon":       {"monday"},
			"sek":       {"second"},
			"sep":       {"september"},
			"sob":       {"saturday"},
			"srj":       {"wednesday"},
			"stw":       {"thursday"},
			"utc":       {"utc"},
			"wut":       {"tuesday"},
			"am":        {"am"},
			"pm":        {"pm"},
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
			"d":         {"day"},
			"h":         {"hour"},
			"l":         {"year"},
			"m":         {"minute"},
			"s":         {"second"},
			"z":         {"z"},
			"|":         {""},
		},
		RelativeType: map[string]string{
			"prichodny tydzen": "in 1 week",
			"prichodny mesac":  "in 1 month",
			"tuton tydzen":     "0 week ago",
			"zasły tydzen":     "1 week ago",
			"this minute":      "0 minute ago",
			"tuton mesac":      "0 month ago",
			"zasły mesac":      "1 month ago",
			"this hour":        "0 hour ago",
			"dzensa":           "0 day ago",
			"jutre":            "in 1 day",
			"kletu":            "in 1 year",
			"letsa":            "0 year ago",
			"wcera":            "1 day ago",
			"loni":             "1 year ago",
			"now":              "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) tydzenjemi`), "$1 week ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) hodzinami`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) sekundami`), "$1 second ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) tydzenjom`), "$1 week ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) mesacami`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) minutami`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) hodzinu`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) mesacom`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) sekundu`), "$1 second ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) tydzenjow`), "in $1 week"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) dnjemi`), "$1 day ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) letami`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) minutu`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sekundow`), "in $1 second"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) dnjom`), "$1 day ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) letom`), "$1 year ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) hodzinu`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) mesacow`), "in $1 month"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) minutow`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sekundu`), "in $1 second"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) hodz`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) tydz`), "$1 week ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) hodzin`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) minutu`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) tydzen`), "in $1 week"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) dnj`), "$1 day ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) mes`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) min`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) sek`), "$1 second ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) dnjow`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) mesac`), "in $1 month"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) dzen`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) hodz`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) leto`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) tydz`), "in $1 week"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) d`), "$1 day ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) h`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) l`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) m`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)pred (\d+[.,]?\d*) s`), "$1 second ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) dnj`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) let`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) mes`), "in $1 month"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sek`), "in $1 second"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) d`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) h`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) l`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) m`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) s`), "in $1 second"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(pred \d+[.,]?\d* tydzenjemi|pred \d+[.,]?\d* hodzinami|pred \d+[.,]?\d* sekundami|pred \d+[.,]?\d* tydzenjom|pred \d+[.,]?\d* mesacami|pred \d+[.,]?\d* minutami|pred \d+[.,]?\d* hodzinu|pred \d+[.,]?\d* mesacom|pred \d+[.,]?\d* sekundu|za \d+[.,]?\d* tydzenjow|pred \d+[.,]?\d* dnjemi|pred \d+[.,]?\d* letami|pred \d+[.,]?\d* minutu|za \d+[.,]?\d* sekundow|pred \d+[.,]?\d* dnjom|pred \d+[.,]?\d* letom|za \d+[.,]?\d* hodzinu|za \d+[.,]?\d* mesacow|za \d+[.,]?\d* minutow|za \d+[.,]?\d* sekundu|pred \d+[.,]?\d* hodz|pred \d+[.,]?\d* tydz|za \d+[.,]?\d* hodzin|za \d+[.,]?\d* minutu|za \d+[.,]?\d* tydzen|pred \d+[.,]?\d* dnj|pred \d+[.,]?\d* mes|pred \d+[.,]?\d* min|pred \d+[.,]?\d* sek|za \d+[.,]?\d* dnjow|za \d+[.,]?\d* mesac|za \d+[.,]?\d* dzen|za \d+[.,]?\d* hodz|za \d+[.,]?\d* leto|za \d+[.,]?\d* tydz|pred \d+[.,]?\d* d|pred \d+[.,]?\d* h|pred \d+[.,]?\d* l|pred \d+[.,]?\d* m|pred \d+[.,]?\d* s|za \d+[.,]?\d* dnj|za \d+[.,]?\d* let|za \d+[.,]?\d* mes|za \d+[.,]?\d* min|za \d+[.,]?\d* sek|za \d+[.,]?\d* d|za \d+[.,]?\d* h|za \d+[.,]?\d* l|za \d+[.,]?\d* m|za \d+[.,]?\d* s)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(pred \d+[.,]?\d* tydzenjemi|pred \d+[.,]?\d* hodzinami|pred \d+[.,]?\d* sekundami|pred \d+[.,]?\d* tydzenjom|pred \d+[.,]?\d* mesacami|pred \d+[.,]?\d* minutami|pred \d+[.,]?\d* hodzinu|pred \d+[.,]?\d* mesacom|pred \d+[.,]?\d* sekundu|za \d+[.,]?\d* tydzenjow|pred \d+[.,]?\d* dnjemi|pred \d+[.,]?\d* letami|pred \d+[.,]?\d* minutu|za \d+[.,]?\d* sekundow|pred \d+[.,]?\d* dnjom|pred \d+[.,]?\d* letom|za \d+[.,]?\d* hodzinu|za \d+[.,]?\d* mesacow|za \d+[.,]?\d* minutow|za \d+[.,]?\d* sekundu|pred \d+[.,]?\d* hodz|pred \d+[.,]?\d* tydz|za \d+[.,]?\d* hodzin|za \d+[.,]?\d* minutu|za \d+[.,]?\d* tydzen|pred \d+[.,]?\d* dnj|pred \d+[.,]?\d* mes|pred \d+[.,]?\d* min|pred \d+[.,]?\d* sek|za \d+[.,]?\d* dnjow|za \d+[.,]?\d* mesac|za \d+[.,]?\d* dzen|za \d+[.,]?\d* hodz|za \d+[.,]?\d* leto|za \d+[.,]?\d* tydz|pred \d+[.,]?\d* d|pred \d+[.,]?\d* h|pred \d+[.,]?\d* l|pred \d+[.,]?\d* m|pred \d+[.,]?\d* s|za \d+[.,]?\d* dnj|za \d+[.,]?\d* let|za \d+[.,]?\d* mes|za \d+[.,]?\d* min|za \d+[.,]?\d* sek|za \d+[.,]?\d* d|za \d+[.,]?\d* h|za \d+[.,]?\d* l|za \d+[.,]?\d* m|za \d+[.,]?\d* s)$`),
		KnownWords:      []string{"prichodny tydzen", "prichodny mesac", "tuton tydzen", "zasły tydzen", "this minute", "tuton mesac", "zasły mesac", "dopołdnja", "popołdnju", "september", "septembra", "this hour", "december", "decembra", "februara", "njedzela", "nowember", "nowembra", "pondzela", "awgusta", "februar", "hodzina", "januara", "oktober", "oktobra", "sekunda", "stwortk", "apryla", "awgust", "dzensa", "januar", "julija", "junija", "minuta", "sobota", "srjeda", "tydzen", "wutora", "apryl", "julij", "junij", "jutre", "kletu", "letsa", "merca", "mesac", "pjatk", "wcera", "dzen", "hodz", "leto", "loni", "meja", "meje", "merc", "tydz", "apr", "awg", "dec", "feb", "gmt", "jan", "jul", "jun", "mej", "mer", "mes", "min", "nje", "now", "okt", "pja", "pon", "sek", "sep", "sob", "srj", "stw", "utc", "wut", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "h", "l", "m", "s", "z", "|"},
	})
}
