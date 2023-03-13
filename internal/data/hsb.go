// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var hsb_Locale = merge(nil, LocaleData{
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
		{regexp.MustCompile(`(?i)pred (\d+) tydzenjemi`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pred (\d+) hodzinami`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) sekundami`), "$1 second ago"},
		{regexp.MustCompile(`(?i)pred (\d+) tydzenjom`), "$1 week ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mesacami`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) minutami`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) hodzinu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mesacom`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) sekundu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)za (\d+) tydzenjow`), "in $1 week"},
		{regexp.MustCompile(`(?i)pred (\d+) dnjemi`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pred (\d+) letami`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pred (\d+) minutu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)za (\d+) sekundow`), "in $1 second"},
		{regexp.MustCompile(`(?i)pred (\d+) dnjom`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pred (\d+) letom`), "$1 year ago"},
		{regexp.MustCompile(`(?i)za (\d+) hodzinu`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) mesacow`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) minutow`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) sekundu`), "in $1 second"},
		{regexp.MustCompile(`(?i)pred (\d+) hodz`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) tydz`), "$1 week ago"},
		{regexp.MustCompile(`(?i)za (\d+) hodzin`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) minutu`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) tydzen`), "in $1 week"},
		{regexp.MustCompile(`(?i)pred (\d+) dnj`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pred (\d+) mes`), "$1 month ago"},
		{regexp.MustCompile(`(?i)pred (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) sek`), "$1 second ago"},
		{regexp.MustCompile(`(?i)za (\d+) dnjow`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) mesac`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) dzen`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) hodz`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) leto`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) tydz`), "in $1 week"},
		{regexp.MustCompile(`(?i)pred (\d+) d`), "$1 day ago"},
		{regexp.MustCompile(`(?i)pred (\d+) h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)pred (\d+) l`), "$1 year ago"},
		{regexp.MustCompile(`(?i)pred (\d+) m`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)pred (\d+) s`), "$1 second ago"},
		{regexp.MustCompile(`(?i)za (\d+) dnj`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) let`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) mes`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) d`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) h`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) l`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) m`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) s`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(pred \d+ tydzenjemi|pred \d+ hodzinami|pred \d+ sekundami|pred \d+ tydzenjom|pred \d+ mesacami|pred \d+ minutami|pred \d+ hodzinu|pred \d+ mesacom|pred \d+ sekundu|za \d+ tydzenjow|pred \d+ dnjemi|pred \d+ letami|pred \d+ minutu|za \d+ sekundow|pred \d+ dnjom|pred \d+ letom|za \d+ hodzinu|za \d+ mesacow|za \d+ minutow|za \d+ sekundu|pred \d+ hodz|pred \d+ tydz|za \d+ hodzin|za \d+ minutu|za \d+ tydzen|pred \d+ dnj|pred \d+ mes|pred \d+ min|pred \d+ sek|za \d+ dnjow|za \d+ mesac|za \d+ dzen|za \d+ hodz|za \d+ leto|za \d+ tydz|pred \d+ d|pred \d+ h|pred \d+ l|pred \d+ m|pred \d+ s|za \d+ dnj|za \d+ let|za \d+ mes|za \d+ min|za \d+ sek|za \d+ d|za \d+ h|za \d+ l|za \d+ m|za \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(pred \d+ tydzenjemi|pred \d+ hodzinami|pred \d+ sekundami|pred \d+ tydzenjom|pred \d+ mesacami|pred \d+ minutami|pred \d+ hodzinu|pred \d+ mesacom|pred \d+ sekundu|za \d+ tydzenjow|pred \d+ dnjemi|pred \d+ letami|pred \d+ minutu|za \d+ sekundow|pred \d+ dnjom|pred \d+ letom|za \d+ hodzinu|za \d+ mesacow|za \d+ minutow|za \d+ sekundu|pred \d+ hodz|pred \d+ tydz|za \d+ hodzin|za \d+ minutu|za \d+ tydzen|pred \d+ dnj|pred \d+ mes|pred \d+ min|pred \d+ sek|za \d+ dnjow|za \d+ mesac|za \d+ dzen|za \d+ hodz|za \d+ leto|za \d+ tydz|pred \d+ d|pred \d+ h|pred \d+ l|pred \d+ m|pred \d+ s|za \d+ dnj|za \d+ let|za \d+ mes|za \d+ min|za \d+ sek|za \d+ d|za \d+ h|za \d+ l|za \d+ m|za \d+ s)$`),
	KnownWords:      []string{"prichodny tydzen", "prichodny mesac", "tuton tydzen", "zasły tydzen", "this minute", "tuton mesac", "zasły mesac", "dopołdnja", "popołdnju", "september", "septembra", "this hour", "december", "decembra", "februara", "njedzela", "nowember", "nowembra", "pondzela", "awgusta", "februar", "hodzina", "januara", "oktober", "oktobra", "sekunda", "stwortk", "apryla", "awgust", "dzensa", "januar", "julija", "junija", "minuta", "sobota", "srjeda", "tydzen", "wutora", "apryl", "julij", "junij", "jutre", "kletu", "letsa", "merca", "mesac", "pjatk", "wcera", "dzen", "hodz", "leto", "loni", "meja", "meje", "merc", "tydz", "apr", "awg", "dec", "feb", "gmt", "jan", "jul", "jun", "mej", "mer", "mes", "min", "nje", "now", "okt", "pja", "pon", "sek", "sep", "sob", "srj", "stw", "utc", "wut", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "h", "l", "m", "s", "z", "|"},
})
