// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var bs_Latn_Locale = merge(nil, LocaleData{
	Name:      "bs-Latn",
	DateOrder: "DMY.",
	Charset:   []rune(`bcdefghijklnorstuvzćčš`),
	Translations: map[string][]string{
		"ponedjeljak": {"monday"},
		"prijepodne":  {"am"},
		"septembar":   {"september"},
		"cetvrtak":    {"thursday"},
		"decembar":    {"december"},
		"nedjelja":    {"sunday"},
		"novembar":    {"november"},
		"februar":     {"february"},
		"oktobar":     {"october"},
		"popodne":     {"pm"},
		"sedmica":     {"week"},
		"sekunda":     {"second"},
		"srijeda":     {"wednesday"},
		"avgust":      {"august"},
		"godina":      {"year"},
		"januar":      {"january"},
		"minuta":      {"minute"},
		"mjesec":      {"month"},
		"subota":      {"saturday"},
		"utorak":      {"tuesday"},
		"april":       {"april"},
		"petak":       {"friday"},
		"juli":        {"july"},
		"juni":        {"june"},
		"mart":        {"march"},
		"apr":         {"april"},
		"avg":         {"august"},
		"cet":         {"thursday"},
		"dan":         {"day"},
		"dec":         {"december"},
		"feb":         {"february"},
		"gmt":         {"gmt"},
		"god":         {"year"},
		"jan":         {"january"},
		"jul":         {"july"},
		"jun":         {"june"},
		"maj":         {"may"},
		"mar":         {"march"},
		"min":         {"minute"},
		"ned":         {"sunday"},
		"nov":         {"november"},
		"okt":         {"october"},
		"pet":         {"friday"},
		"pon":         {"monday"},
		"sat":         {"hour"},
		"sed":         {"week"},
		"sek":         {"second"},
		"sep":         {"september"},
		"sri":         {"wednesday"},
		"sub":         {"saturday"},
		"utc":         {"utc"},
		"uto":         {"tuesday"},
		"am":          {"am"},
		"mj":          {"month"},
		"pm":          {"pm"},
		" ":           {" "},
		"'":           {""},
		"+":           {"+"},
		",":           {""},
		"-":           {"-"},
		".":           {"."},
		"/":           {"/"},
		":":           {":"},
		";":           {""},
		"@":           {""},
		"[":           {""},
		"]":           {""},
		"g":           {"year"},
		"h":           {"hour"},
		"s":           {"second"},
		"z":           {"z"},
		"|":           {""},
	},
	RelativeType: map[string]string{
		"sljedece sedmice": "in 1 week",
		"sljedece godine":  "in 1 year",
		"sljedeci mjesec":  "in 1 month",
		"prosle sedmice":   "1 week ago",
		"prosle godine":    "1 year ago",
		"prosli mjesec":    "1 month ago",
		"ovaj mjesec":      "0 month ago",
		"ove sedmice":      "0 week ago",
		"ova minuta":       "0 minute ago",
		"ove godine":       "0 year ago",
		"ovaj sat":         "0 hour ago",
		"danas":            "0 day ago",
		"jucer":            "1 day ago",
		"sutra":            "in 1 day",
		"sada":             "0 second ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)prije (\d+) mjeseci`), "$1 month ago"},
		{regexp.MustCompile(`(?i)prije (\d+) sedmica`), "$1 week ago"},
		{regexp.MustCompile(`(?i)prije (\d+) sedmicu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)prije (\d+) sekundi`), "$1 second ago"},
		{regexp.MustCompile(`(?i)prije (\d+) sekundu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)prije (\d+) godina`), "$1 year ago"},
		{regexp.MustCompile(`(?i)prije (\d+) godinu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)prije (\d+) minuta`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)prije (\d+) minutu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)prije (\d+) mjesec`), "$1 month ago"},
		{regexp.MustCompile(`(?i)prije (\d+) dana`), "$1 day ago"},
		{regexp.MustCompile(`(?i)prije (\d+) sati`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)za (\d+) mjeseci`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) sedmica`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) sedmicu`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) sekundi`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) sekundu`), "in $1 second"},
		{regexp.MustCompile(`(?i)prije (\d+) dan`), "$1 day ago"},
		{regexp.MustCompile(`(?i)prije (\d+) god`), "$1 year ago"},
		{regexp.MustCompile(`(?i)prije (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)prije (\d+) sat`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)prije (\d+) sed`), "$1 week ago"},
		{regexp.MustCompile(`(?i)prije (\d+) sek`), "$1 second ago"},
		{regexp.MustCompile(`(?i)za (\d+) godina`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) godinu`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) minuta`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) minutu`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) mjesec`), "in $1 month"},
		{regexp.MustCompile(`(?i)prije (\d+) mj`), "$1 month ago"},
		{regexp.MustCompile(`(?i)prije (\d+) d`), "$1 day ago"},
		{regexp.MustCompile(`(?i)prije (\d+) g`), "$1 year ago"},
		{regexp.MustCompile(`(?i)za (\d+) dana`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) sati`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) dan`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) god`), "in $1 year"},
		{regexp.MustCompile(`(?i)za (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)za (\d+) sat`), "in $1 hour"},
		{regexp.MustCompile(`(?i)za (\d+) sed`), "in $1 week"},
		{regexp.MustCompile(`(?i)za (\d+) sek`), "in $1 second"},
		{regexp.MustCompile(`(?i)za (\d+) mj`), "in $1 month"},
		{regexp.MustCompile(`(?i)za (\d+) d`), "in $1 day"},
		{regexp.MustCompile(`(?i)za (\d+) g`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(prije \d+ mjeseci|prije \d+ sedmica|prije \d+ sedmicu|prije \d+ sekundi|prije \d+ sekundu|prije \d+ godina|prije \d+ godinu|prije \d+ minuta|prije \d+ minutu|prije \d+ mjesec|prije \d+ dana|prije \d+ sati|za \d+ mjeseci|za \d+ sedmica|za \d+ sedmicu|za \d+ sekundi|za \d+ sekundu|prije \d+ dan|prije \d+ god|prije \d+ min|prije \d+ sat|prije \d+ sed|prije \d+ sek|za \d+ godina|za \d+ godinu|za \d+ minuta|za \d+ minutu|za \d+ mjesec|prije \d+ mj|prije \d+ d|prije \d+ g|za \d+ dana|za \d+ sati|za \d+ dan|za \d+ god|za \d+ min|za \d+ sat|za \d+ sed|za \d+ sek|za \d+ mj|za \d+ d|za \d+ g)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(prije \d+ mjeseci|prije \d+ sedmica|prije \d+ sedmicu|prije \d+ sekundi|prije \d+ sekundu|prije \d+ godina|prije \d+ godinu|prije \d+ minuta|prije \d+ minutu|prije \d+ mjesec|prije \d+ dana|prije \d+ sati|za \d+ mjeseci|za \d+ sedmica|za \d+ sedmicu|za \d+ sekundi|za \d+ sekundu|prije \d+ dan|prije \d+ god|prije \d+ min|prije \d+ sat|prije \d+ sed|prije \d+ sek|za \d+ godina|za \d+ godinu|za \d+ minuta|za \d+ minutu|za \d+ mjesec|prije \d+ mj|prije \d+ d|prije \d+ g|za \d+ dana|za \d+ sati|za \d+ dan|za \d+ god|za \d+ min|za \d+ sat|za \d+ sed|za \d+ sek|za \d+ mj|za \d+ d|za \d+ g)$`),
	KnownWords:      []string{"sljedece sedmice", "sljedece godine", "sljedeci mjesec", "prosle sedmice", "prosle godine", "prosli mjesec", "ovaj mjesec", "ove sedmice", "ponedjeljak", "ova minuta", "ove godine", "prijepodne", "septembar", "cetvrtak", "decembar", "nedjelja", "novembar", "ovaj sat", "februar", "oktobar", "popodne", "sedmica", "sekunda", "srijeda", "avgust", "godina", "januar", "minuta", "mjesec", "subota", "utorak", "april", "danas", "jucer", "petak", "sutra", "juli", "juni", "mart", "sada", "apr", "avg", "cet", "dan", "dec", "feb", "gmt", "god", "jan", "jul", "jun", "maj", "mar", "min", "ned", "nov", "okt", "pet", "pon", "sat", "sed", "sek", "sep", "sri", "sub", "utc", "uto", "am", "mj", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "g", "h", "s", "z", "|"},
})
