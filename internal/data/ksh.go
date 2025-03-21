// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	ksh_Locale LocaleData
)

func init() {
	ksh_Locale = merge(nil, LocaleData{
		Name:      "ksh",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefghijklnorstuvwxzßäöü`),
		Translations: map[string][]string{
			"uhr nommendaachs": {"pm"},
			"uhr vormiddaachs": {"am"},
			"dunnersdaach":     {"thursday"},
			"nommendaach":      {"pm"},
			"vormeddaach":      {"am"},
			"dinnsdaach":       {"tuesday"},
			"friidaach":        {"friday"},
			"mohndaach":        {"monday"},
			"samsdaach":        {"saturday"},
			"septamber":        {"september"},
			"sunndaach":        {"sunday"},
			"dezamber":         {"december"},
			"novamber":         {"november"},
			"oktohber":         {"october"},
			"fabrowa":          {"february"},
			"jannewa":          {"january"},
			"metwoch":          {"wednesday"},
			"schtund":          {"hour"},
			"aprell":           {"april"},
			"menutt":           {"minute"},
			"sekond":           {"second"},
			"daach":            {"day"},
			"juuli":            {"july"},
			"juuni":            {"june"},
			"mohnd":            {"month"},
			"oujoß":            {"august"},
			"johr":             {"year"},
			"maaz":             {"march"},
			"woch":             {"week"},
			"apr":              {"april"},
			"dez":              {"december"},
			"fab":              {"february"},
			"gmt":              {"gmt"},
			"jan":              {"january"},
			"jul":              {"july"},
			"jun":              {"june"},
			"mai":              {"may"},
			"maz":              {"march"},
			"min":              {"minute"},
			"nov":              {"november"},
			"okt":              {"october"},
			"ouj":              {"august"},
			"sap":              {"september"},
			"sek":              {"second"},
			"std":              {"hour"},
			"utc":              {"utc"},
			"am":               {"am"},
			"di":               {"tuesday"},
			"du":               {"thursday"},
			"fr":               {"friday"},
			"me":               {"wednesday"},
			"mo":               {"monday"},
			"nm":               {"pm"},
			"pm":               {"pm"},
			"sa":               {"saturday"},
			"su":               {"sunday"},
			"vm":               {"am"},
			" ":                {" "},
			"'":                {""},
			"+":                {"+"},
			",":                {""},
			"-":                {"-"},
			".":                {"."},
			"/":                {"/"},
			":":                {":"},
			";":                {""},
			"@":                {""},
			"[":                {""},
			"]":                {""},
			"d":                {"day"},
			"j":                {"year"},
			"m":                {"month", "minute"},
			"s":                {"hour", "second"},
			"w":                {"week"},
			"z":                {"z"},
			"|":                {""},
		},
		RelativeType: map[string]string{
			"nachste mohnd": "in 1 month",
			"nachste woche": "in 1 week",
			"latzde mohnd":  "1 month ago",
			"diese mohnd":   "0 month ago",
			"this minute":   "0 minute ago",
			"this hour":     "0 hour ago",
			"diß johr":      "0 year ago",
			"laz johr":      "1 year ago",
			"laz woch":      "1 week ago",
			"nax johr":      "in 1 year",
			"di woch":       "0 week ago",
			"jestere":       "1 day ago",
			"morje":         "in 1 day",
			"huck":          "0 day ago",
			"now":           "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) johre`), "$1 year ago"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) johre`), "in $1 year"},
			{regexp.MustCompile(`(?i)vor (\d+[.,]?\d*) johr`), "$1 year ago"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) johr`), "in $1 year"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(vor \d+[.,]?\d* johre|en \d+[.,]?\d* johre|vor \d+[.,]?\d* johr|en \d+[.,]?\d* johr)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(vor \d+[.,]?\d* johre|en \d+[.,]?\d* johre|vor \d+[.,]?\d* johr|en \d+[.,]?\d* johr)$`),
		KnownWords:      []string{"uhr nommendaachs", "uhr vormiddaachs", "nachste mohnd", "nachste woche", "dunnersdaach", "latzde mohnd", "diese mohnd", "nommendaach", "this minute", "vormeddaach", "dinnsdaach", "friidaach", "mohndaach", "samsdaach", "septamber", "sunndaach", "this hour", "dezamber", "diß johr", "laz johr", "laz woch", "nax johr", "novamber", "oktohber", "di woch", "fabrowa", "jannewa", "jestere", "metwoch", "schtund", "aprell", "menutt", "sekond", "daach", "juuli", "juuni", "mohnd", "morje", "oujoß", "huck", "johr", "maaz", "woch", "apr", "dez", "fab", "gmt", "jan", "jul", "jun", "mai", "maz", "min", "nov", "now", "okt", "ouj", "sap", "sek", "std", "utc", "am", "di", "du", "fr", "me", "mo", "nm", "pm", "sa", "su", "vm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "j", "m", "s", "w", "z", "|"},
	})
}
