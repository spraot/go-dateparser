// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var to_Locale = merge(nil, LocaleData{
	Name:      "to",
	DateOrder: "DMY",
	Charset:   []rune(`cefghiklnorstuvzáéíóúāēīōū`),
	Translations: map[string][]string{
		"tu'apulelulu": {"thursday"},
		"hengihengi":   {"am"},
		"'epeleli":     {"april"},
		"'okatopa":     {"october"},
		"pulelulu":     {"wednesday"},
		"sepitema":     {"september"},
		"tokonaki":     {"saturday"},
		"'aokosi":      {"august"},
		"falaite":      {"friday"},
		"fepueli":      {"february"},
		"sanuali":      {"january"},
		"efiafi":       {"pm"},
		"ma'asi":       {"march"},
		"mahina":       {"month"},
		"miniti":       {"minute"},
		"monite":       {"monday"},
		"novema":       {"november"},
		"sapate":       {"sunday"},
		"sekoni":       {"second"},
		"siulai":       {"july"},
		"tisema":       {"december"},
		"tusite":       {"tuesday"},
		"'aho":         {"day"},
		"'aok":         {"august"},
		"'epe":         {"april"},
		"'oka":         {"october"},
		"houa":         {"hour"},
		"ma'a":         {"march"},
		"sune":         {"june"},
		"ta'u":         {"year"},
		"tu'a":         {"thursday"},
		"uike":         {"week"},
		"fal":          {"friday"},
		"fep":          {"february"},
		"gmt":          {"gmt"},
		"mon":          {"monday"},
		"nov":          {"november"},
		"pul":          {"wednesday"},
		"san":          {"january"},
		"sap":          {"sunday"},
		"sep":          {"september"},
		"siu":          {"july"},
		"sun":          {"june"},
		"tis":          {"december"},
		"tok":          {"saturday"},
		"tus":          {"tuesday"},
		"utc":          {"utc"},
		"am":           {"am"},
		"ea":           {"pm"},
		"hh":           {"am"},
		"me":           {"may"},
		"pm":           {"pm"},
		" ":            {" "},
		"'":            {""},
		"+":            {"+"},
		",":            {""},
		"-":            {"-"},
		".":            {"."},
		"/":            {"/"},
		":":            {":"},
		";":            {""},
		"@":            {""},
		"[":            {""},
		"]":            {""},
		"z":            {"z"},
		"|":            {""},
	},
	RelativeType: map[string]string{
		"mahina kuo'osi": "1 month ago",
		"mahina kaha'u":  "in 1 month",
		"'apongipongi":   "in 1 day",
		"ta'u kuo'osi":   "1 year ago",
		"uike kuo'osi":   "1 week ago",
		"ta'u kaha'u":    "in 1 year",
		"this minute":    "0 minute ago",
		"uike kaha'u":    "in 1 week",
		"mahina ni":      "0 month ago",
		"this hour":      "0 hour ago",
		"taimi ni":       "0 second ago",
		"'aho ni":        "0 day ago",
		"'aneafi":        "1 day ago",
		"ta'u ni":        "0 year ago",
		"uike ni":        "0 week ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)mahina 'e (\d+) kuo'osi`), "$1 month ago"},
		{regexp.MustCompile(`(?i)miniti 'e (\d+) kuo'osi`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)sekoni 'e (\d+) kuo'osi`), "$1 second ago"},
		{regexp.MustCompile(`(?i)'aho 'e (\d+) kuo'osi`), "$1 day ago"},
		{regexp.MustCompile(`(?i)'i he mahina 'e (\d+)`), "in $1 month"},
		{regexp.MustCompile(`(?i)'i he miniti 'e (\d+)`), "in $1 minute"},
		{regexp.MustCompile(`(?i)'i he sekoni 'e (\d+)`), "in $1 second"},
		{regexp.MustCompile(`(?i)houa 'e (\d+) kuo'osi`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)ta'u 'e (\d+) kuo'osi`), "$1 year ago"},
		{regexp.MustCompile(`(?i)uike 'e (\d+) kuo'osi`), "$1 week ago"},
		{regexp.MustCompile(`(?i)'i he 'aho 'e (\d+)`), "in $1 day"},
		{regexp.MustCompile(`(?i)'i he houa 'e (\d+)`), "in $1 hour"},
		{regexp.MustCompile(`(?i)'i he ta'u 'e (\d+)`), "in $1 year"},
		{regexp.MustCompile(`(?i)'i he uike 'e (\d+)`), "in $1 week"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(mahina 'e \d+ kuo'osi|miniti 'e \d+ kuo'osi|sekoni 'e \d+ kuo'osi|'aho 'e \d+ kuo'osi|'i he mahina 'e \d+|'i he miniti 'e \d+|'i he sekoni 'e \d+|houa 'e \d+ kuo'osi|ta'u 'e \d+ kuo'osi|uike 'e \d+ kuo'osi|'i he 'aho 'e \d+|'i he houa 'e \d+|'i he ta'u 'e \d+|'i he uike 'e \d+)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(mahina 'e \d+ kuo'osi|miniti 'e \d+ kuo'osi|sekoni 'e \d+ kuo'osi|'aho 'e \d+ kuo'osi|'i he mahina 'e \d+|'i he miniti 'e \d+|'i he sekoni 'e \d+|houa 'e \d+ kuo'osi|ta'u 'e \d+ kuo'osi|uike 'e \d+ kuo'osi|'i he 'aho 'e \d+|'i he houa 'e \d+|'i he ta'u 'e \d+|'i he uike 'e \d+)$`),
	KnownWords:      []string{"mahina kuo'osi", "mahina kaha'u", "'apongipongi", "ta'u kuo'osi", "tu'apulelulu", "uike kuo'osi", "ta'u kaha'u", "this minute", "uike kaha'u", "hengihengi", "mahina ni", "this hour", "'epeleli", "'okatopa", "pulelulu", "sepitema", "taimi ni", "tokonaki", "'aho ni", "'aneafi", "'aokosi", "falaite", "fepueli", "sanuali", "ta'u ni", "uike ni", "efiafi", "ma'asi", "mahina", "miniti", "monite", "novema", "sapate", "sekoni", "siulai", "tisema", "tusite", "'aho", "'aok", "'epe", "'oka", "houa", "ma'a", "sune", "ta'u", "tu'a", "uike", "fal", "fep", "gmt", "mon", "nov", "pul", "san", "sap", "sep", "siu", "sun", "tis", "tok", "tus", "utc", "am", "ea", "hh", "me", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
})
