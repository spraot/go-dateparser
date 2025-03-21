// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	sv_Locale    LocaleData
	sv_AX_Locale LocaleData
	sv_FI_Locale LocaleData
)

func init() {
	sv_Locale = merge(nil, LocaleData{
		Name:                  "sv",
		DateOrder:             "YMD",
		Charset:               []rune(`bcdefghijklnorstuvzäåö`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])en(\z|[^\pL\pM\d_])`), "${1}1${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])tva(\z|[^\pL\pM\d_])`), "${1}2${2}"},
		},
		Translations: map[string][]string{
			"eftermiddag": {"pm"},
			"formiddag":   {"am"},
			"september":   {"september"},
			"december":    {"december"},
			"februari":    {"february"},
			"november":    {"november"},
			"sekunder":    {"second"},
			"augusti":     {"august"},
			"fran nu":     {"in"},
			"januari":     {"january"},
			"manaden":     {"month"},
			"manader":     {"month"},
			"minuter":     {"minute"},
			"oktober":     {"october"},
			"torsdag":     {"thursday"},
			"fredag":      {"friday"},
			"lordag":      {"saturday"},
			"mandag":      {"monday"},
			"onsdag":      {"wednesday"},
			"sekund":      {"second"},
			"sondag":      {"sunday"},
			"timmar":      {"hour"},
			"tisdag":      {"tuesday"},
			"veckor":      {"week"},
			"april":       {"april"},
			"dagar":       {"day"},
			"manad":       {"month"},
			"minut":       {"minute"},
			"sedan":       {"ago"},
			"timme":       {"hour"},
			"torsd":       {"thursday"},
			"vecka":       {"week"},
			"aret":        {"year"},
			"fred":        {"friday"},
			"juli":        {"july"},
			"juni":        {"june"},
			"lord":        {"saturday"},
			"mand":        {"monday"},
			"mars":        {"march"},
			"onsd":        {"wednesday"},
			"sept":        {"september"},
			"sond":        {"sunday"},
			"tisd":        {"tuesday"},
			"tors":        {"thursday"},
			"apr":         {"april"},
			"aug":         {"august"},
			"dag":         {"day"},
			"dec":         {"december"},
			"den":         {""},
			"feb":         {"february"},
			"fre":         {"friday"},
			"gmt":         {"gmt"},
			"jan":         {"january"},
			"jul":         {"july"},
			"jun":         {"june"},
			"lor":         {"saturday"},
			"maj":         {"may"},
			"man":         {"month", "monday"},
			"min":         {"minute"},
			"nov":         {"november"},
			"okt":         {"october"},
			"ons":         {"wednesday"},
			"sek":         {"second"},
			"sep":         {"september"},
			"son":         {"sunday"},
			"tim":         {"hour"},
			"tis":         {"tuesday"},
			"utc":         {"utc"},
			"am":          {"am"},
			"ar":          {"year"},
			"em":          {"pm"},
			"fm":          {"am"},
			"om":          {"in"},
			"pa":          {""},
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
			"h":           {"hour"},
			"m":           {"month", "minute"},
			"s":           {"second"},
			"t":           {"hour"},
			"v":           {"week"},
			"z":           {"z"},
			"|":           {""},
		},
		RelativeType: map[string]string{
			"forra manaden": "1 month ago",
			"forra veckan":  "1 week ago",
			"denna manad":   "0 month ago",
			"denna minut":   "0 minute ago",
			"denna timme":   "0 hour ago",
			"denna vecka":   "0 week ago",
			"nasta manad":   "in 1 month",
			"nasta vecka":   "in 1 week",
			"forra aret":    "1 year ago",
			"denna man":     "0 month ago",
			"forra man":     "1 month ago",
			"nasta man":     "in 1 month",
			"i morgon":      "in 1 day",
			"nasta ar":      "in 1 year",
			"denna v":       "0 week ago",
			"forra v":       "1 week ago",
			"forrgar":       "2 day ago",
			"imorgon":       "in 1 day",
			"nasta v":       "in 1 week",
			"i fjol":        "1 year ago",
			"i dag":         "0 day ago",
			"i gar":         "1 day ago",
			"i ar":          "0 year ago",
			"idag":          "0 day ago",
			"igar":          "1 day ago",
			"nu":            "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) sekunder sedan`), "$1 second ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) manader sedan`), "$1 month ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) minuter sedan`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) sekund sedan`), "$1 second ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) timmar sedan`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) veckor sedan`), "$1 week ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) dagar sedan`), "$1 day ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) manad sedan`), "$1 month ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) minut sedan`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) timme sedan`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) vecka sedan`), "$1 week ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) dag sedan`), "$1 day ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) man sedan`), "$1 month ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) min sedan`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) sek sedan`), "$1 second ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) tim sedan`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) ar sedan`), "$1 year ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) d sedan`), "$1 day ago"},
			{regexp.MustCompile(`(?i)for (\d+[.,]?\d*) v sedan`), "$1 week ago"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) sekunder`), "in $1 second"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) manader`), "in $1 month"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) minuter`), "in $1 minute"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) sekund`), "in $1 second"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) timmar`), "in $1 hour"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) veckor`), "in $1 week"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) dagar`), "in $1 day"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) manad`), "in $1 month"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) minut`), "in $1 minute"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) timme`), "in $1 hour"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) vecka`), "in $1 week"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) dag`), "in $1 day"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) man`), "in $1 month"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) sek`), "in $1 second"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) tim`), "in $1 hour"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) ar`), "in $1 year"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) d`), "in $1 day"},
			{regexp.MustCompile(`(?i)om (\d+[.,]?\d*) v`), "in $1 week"},
			{regexp.MustCompile(`(?i)−(\d+[.,]?\d*) man`), "$1 month ago"},
			{regexp.MustCompile(`(?i)−(\d+[.,]?\d*) min`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)−(\d+[.,]?\d*) ar`), "$1 year ago"},
			{regexp.MustCompile(`(?i)−(\d+[.,]?\d*) d`), "$1 day ago"},
			{regexp.MustCompile(`(?i)−(\d+[.,]?\d*) h`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)−(\d+[.,]?\d*) s`), "$1 second ago"},
			{regexp.MustCompile(`(?i)−(\d+[.,]?\d*) v`), "$1 week ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(for \d+[.,]?\d* sekunder sedan|for \d+[.,]?\d* manader sedan|for \d+[.,]?\d* minuter sedan|for \d+[.,]?\d* sekund sedan|for \d+[.,]?\d* timmar sedan|for \d+[.,]?\d* veckor sedan|for \d+[.,]?\d* dagar sedan|for \d+[.,]?\d* manad sedan|for \d+[.,]?\d* minut sedan|for \d+[.,]?\d* timme sedan|for \d+[.,]?\d* vecka sedan|for \d+[.,]?\d* dag sedan|for \d+[.,]?\d* man sedan|for \d+[.,]?\d* min sedan|for \d+[.,]?\d* sek sedan|for \d+[.,]?\d* tim sedan|for \d+[.,]?\d* ar sedan|for \d+[.,]?\d* d sedan|for \d+[.,]?\d* v sedan|om \d+[.,]?\d* sekunder|om \d+[.,]?\d* manader|om \d+[.,]?\d* minuter|om \d+[.,]?\d* sekund|om \d+[.,]?\d* timmar|om \d+[.,]?\d* veckor|om \d+[.,]?\d* dagar|om \d+[.,]?\d* manad|om \d+[.,]?\d* minut|om \d+[.,]?\d* timme|om \d+[.,]?\d* vecka|om \d+[.,]?\d* dag|om \d+[.,]?\d* man|om \d+[.,]?\d* min|om \d+[.,]?\d* sek|om \d+[.,]?\d* tim|om \d+[.,]?\d* ar|om \d+[.,]?\d* d|om \d+[.,]?\d* v|−\d+[.,]?\d* man|−\d+[.,]?\d* min|−\d+[.,]?\d* ar|−\d+[.,]?\d* d|−\d+[.,]?\d* h|−\d+[.,]?\d* s|−\d+[.,]?\d* v)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(for \d+[.,]?\d* sekunder sedan|for \d+[.,]?\d* manader sedan|for \d+[.,]?\d* minuter sedan|for \d+[.,]?\d* sekund sedan|for \d+[.,]?\d* timmar sedan|for \d+[.,]?\d* veckor sedan|for \d+[.,]?\d* dagar sedan|for \d+[.,]?\d* manad sedan|for \d+[.,]?\d* minut sedan|for \d+[.,]?\d* timme sedan|for \d+[.,]?\d* vecka sedan|for \d+[.,]?\d* dag sedan|for \d+[.,]?\d* man sedan|for \d+[.,]?\d* min sedan|for \d+[.,]?\d* sek sedan|for \d+[.,]?\d* tim sedan|for \d+[.,]?\d* ar sedan|for \d+[.,]?\d* d sedan|for \d+[.,]?\d* v sedan|om \d+[.,]?\d* sekunder|om \d+[.,]?\d* manader|om \d+[.,]?\d* minuter|om \d+[.,]?\d* sekund|om \d+[.,]?\d* timmar|om \d+[.,]?\d* veckor|om \d+[.,]?\d* dagar|om \d+[.,]?\d* manad|om \d+[.,]?\d* minut|om \d+[.,]?\d* timme|om \d+[.,]?\d* vecka|om \d+[.,]?\d* dag|om \d+[.,]?\d* man|om \d+[.,]?\d* min|om \d+[.,]?\d* sek|om \d+[.,]?\d* tim|om \d+[.,]?\d* ar|om \d+[.,]?\d* d|om \d+[.,]?\d* v|−\d+[.,]?\d* man|−\d+[.,]?\d* min|−\d+[.,]?\d* ar|−\d+[.,]?\d* d|−\d+[.,]?\d* h|−\d+[.,]?\d* s|−\d+[.,]?\d* v)$`),
		KnownWords:      []string{"forra manaden", "forra veckan", "denna manad", "denna minut", "denna timme", "denna vecka", "eftermiddag", "nasta manad", "nasta vecka", "forra aret", "denna man", "formiddag", "forra man", "nasta man", "september", "december", "februari", "i morgon", "nasta ar", "november", "sekunder", "augusti", "denna v", "forra v", "forrgar", "fran nu", "imorgon", "januari", "manaden", "manader", "minuter", "nasta v", "oktober", "torsdag", "fredag", "i fjol", "lordag", "mandag", "onsdag", "sekund", "sondag", "timmar", "tisdag", "veckor", "april", "dagar", "i dag", "i gar", "manad", "minut", "sedan", "timme", "torsd", "vecka", "aret", "fred", "i ar", "idag", "igar", "juli", "juni", "lord", "mand", "mars", "onsd", "sept", "sond", "tisd", "tors", "apr", "aug", "dag", "dec", "den", "feb", "fre", "gmt", "jan", "jul", "jun", "lor", "maj", "man", "min", "nov", "okt", "ons", "sek", "sep", "son", "tim", "tis", "utc", "am", "ar", "em", "fm", "nu", "om", "pa", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "h", "m", "s", "t", "v", "z", "|"},
	})

	sv_AX_Locale = merge(&sv_Locale, LocaleData{
		Name:      "sv-AX",
		DateOrder: "YMD",
	})

	sv_FI_Locale = merge(&sv_Locale, LocaleData{
		Name:      "sv-FI",
		DateOrder: "DMY",
	})
}
