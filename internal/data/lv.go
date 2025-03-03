// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	lv_Locale LocaleData
)

func init() {
	lv_Locale = merge(nil, LocaleData{
		Name:      "lv",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefghijklnorstuvzāēīļšū`),
		Translations: map[string][]string{
			"priekspusdiena": {"am"},
			"ceturtdiena":    {"thursday"},
			"pecpusdiena":    {"pm"},
			"piektdiena":     {"friday"},
			"septembris":     {"september"},
			"decembris":      {"december"},
			"februaris":      {"february"},
			"novembris":      {"november"},
			"pirmdiena":      {"monday"},
			"sestdiena":      {"saturday"},
			"svetdiena":      {"sunday"},
			"tresdiena":      {"wednesday"},
			"janvaris":       {"january"},
			"oktobris":       {"october"},
			"otrdiena":       {"tuesday"},
			"sekundes":       {"second"},
			"aprilis":        {"april"},
			"augusts":        {"august"},
			"ceturtd":        {"thursday"},
			"menesis":        {"month"},
			"minutes":        {"minute"},
			"pecpusd":        {"pm"},
			"prieksp":        {"am"},
			"stundas":        {"hour"},
			"julijs":         {"july"},
			"junijs":         {"june"},
			"nedela":         {"week"},
			"piektd":         {"friday"},
			"diena":          {"day"},
			"maijs":          {"may"},
			"marts":          {"march"},
			"pirmd":          {"monday"},
			"sestd":          {"saturday"},
			"svetd":          {"sunday"},
			"tresd":          {"wednesday"},
			"febr":           {"february"},
			"gads":           {"year"},
			"janv":           {"january"},
			"otrd":           {"tuesday"},
			"pecp":           {"pm"},
			"sept":           {"september"},
			"apr":            {"april"},
			"aug":            {"august"},
			"dec":            {"december"},
			"gmt":            {"gmt"},
			"jul":            {"july"},
			"jun":            {"june"},
			"men":            {"month"},
			"min":            {"minute"},
			"ned":            {"week"},
			"nov":            {"november"},
			"okt":            {"october"},
			"sek":            {"second"},
			"utc":            {"utc"},
			"am":             {"am"},
			"pm":             {"pm"},
			"st":             {"hour"},
			" ":              {" "},
			"'":              {""},
			"+":              {"+"},
			",":              {""},
			"-":              {"-"},
			".":              {"."},
			"/":              {"/"},
			":":              {":"},
			";":              {""},
			"@":              {""},
			"[":              {""},
			"]":              {""},
			"d":              {"day"},
			"g":              {"year"},
			"h":              {"hour"},
			"s":              {"second"},
			"z":              {"z"},
			"|":              {""},
		},
		RelativeType: map[string]string{
			"pagajusaja menesi": "1 month ago",
			"pagajusaja nedela": "1 week ago",
			"nakamaja menesi":   "in 1 month",
			"nakamaja nedela":   "in 1 week",
			"pagajusaja gada":   "1 year ago",
			"nakamaja gada":     "in 1 year",
			"saja menesi":       "0 month ago",
			"saja minute":       "0 minute ago",
			"saja nedela":       "0 week ago",
			"saja stunda":       "0 hour ago",
			"saja gada":         "0 year ago",
			"sodien":            "0 day ago",
			"tagad":             "0 second ago",
			"vakar":             "1 day ago",
			"rit":               "in 1 day",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) menesiem`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) sekundem`), "$1 second ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) sekundes`), "$1 second ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) minutem`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) minutes`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) nedelam`), "$1 week ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) nedelas`), "$1 week ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) stundam`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) stundas`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) menesiem`), "in $1 month"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) sekundem`), "in $1 second"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) sekundes`), "in $1 second"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) dienam`), "$1 day ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) dienas`), "$1 day ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) gadiem`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) menesa`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) minutem`), "in $1 minute"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) minutes`), "in $1 minute"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) nedelam`), "in $1 week"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) nedelas`), "in $1 week"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) stundam`), "in $1 hour"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) stundas`), "in $1 hour"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) dienam`), "in $1 day"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) dienas`), "in $1 day"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) gadiem`), "in $1 year"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) menesa`), "in $1 month"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) gada`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) men`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) min`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) ned`), "$1 week ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) sek`), "$1 second ago"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) gada`), "in $1 year"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) st`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) men`), "in $1 month"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) ned`), "in $1 week"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) sek`), "in $1 second"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) d`), "$1 day ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) g`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) h`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pirms (\d+[.,]?\d*) s`), "$1 second ago"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) st`), "in $1 hour"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) d`), "in $1 day"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) g`), "in $1 year"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) h`), "in $1 hour"},
			{regexp.MustCompile(`(?i)pec (\d+[.,]?\d*) s`), "in $1 second"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(pirms \d+[.,]?\d* menesiem|pirms \d+[.,]?\d* sekundem|pirms \d+[.,]?\d* sekundes|pirms \d+[.,]?\d* minutem|pirms \d+[.,]?\d* minutes|pirms \d+[.,]?\d* nedelam|pirms \d+[.,]?\d* nedelas|pirms \d+[.,]?\d* stundam|pirms \d+[.,]?\d* stundas|pec \d+[.,]?\d* menesiem|pec \d+[.,]?\d* sekundem|pec \d+[.,]?\d* sekundes|pirms \d+[.,]?\d* dienam|pirms \d+[.,]?\d* dienas|pirms \d+[.,]?\d* gadiem|pirms \d+[.,]?\d* menesa|pec \d+[.,]?\d* minutem|pec \d+[.,]?\d* minutes|pec \d+[.,]?\d* nedelam|pec \d+[.,]?\d* nedelas|pec \d+[.,]?\d* stundam|pec \d+[.,]?\d* stundas|pec \d+[.,]?\d* dienam|pec \d+[.,]?\d* dienas|pec \d+[.,]?\d* gadiem|pec \d+[.,]?\d* menesa|pirms \d+[.,]?\d* gada|pirms \d+[.,]?\d* men|pirms \d+[.,]?\d* min|pirms \d+[.,]?\d* ned|pirms \d+[.,]?\d* sek|pec \d+[.,]?\d* gada|pirms \d+[.,]?\d* st|pec \d+[.,]?\d* men|pec \d+[.,]?\d* min|pec \d+[.,]?\d* ned|pec \d+[.,]?\d* sek|pirms \d+[.,]?\d* d|pirms \d+[.,]?\d* g|pirms \d+[.,]?\d* h|pirms \d+[.,]?\d* s|pec \d+[.,]?\d* st|pec \d+[.,]?\d* d|pec \d+[.,]?\d* g|pec \d+[.,]?\d* h|pec \d+[.,]?\d* s)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(pirms \d+[.,]?\d* menesiem|pirms \d+[.,]?\d* sekundem|pirms \d+[.,]?\d* sekundes|pirms \d+[.,]?\d* minutem|pirms \d+[.,]?\d* minutes|pirms \d+[.,]?\d* nedelam|pirms \d+[.,]?\d* nedelas|pirms \d+[.,]?\d* stundam|pirms \d+[.,]?\d* stundas|pec \d+[.,]?\d* menesiem|pec \d+[.,]?\d* sekundem|pec \d+[.,]?\d* sekundes|pirms \d+[.,]?\d* dienam|pirms \d+[.,]?\d* dienas|pirms \d+[.,]?\d* gadiem|pirms \d+[.,]?\d* menesa|pec \d+[.,]?\d* minutem|pec \d+[.,]?\d* minutes|pec \d+[.,]?\d* nedelam|pec \d+[.,]?\d* nedelas|pec \d+[.,]?\d* stundam|pec \d+[.,]?\d* stundas|pec \d+[.,]?\d* dienam|pec \d+[.,]?\d* dienas|pec \d+[.,]?\d* gadiem|pec \d+[.,]?\d* menesa|pirms \d+[.,]?\d* gada|pirms \d+[.,]?\d* men|pirms \d+[.,]?\d* min|pirms \d+[.,]?\d* ned|pirms \d+[.,]?\d* sek|pec \d+[.,]?\d* gada|pirms \d+[.,]?\d* st|pec \d+[.,]?\d* men|pec \d+[.,]?\d* min|pec \d+[.,]?\d* ned|pec \d+[.,]?\d* sek|pirms \d+[.,]?\d* d|pirms \d+[.,]?\d* g|pirms \d+[.,]?\d* h|pirms \d+[.,]?\d* s|pec \d+[.,]?\d* st|pec \d+[.,]?\d* d|pec \d+[.,]?\d* g|pec \d+[.,]?\d* h|pec \d+[.,]?\d* s)$`),
		KnownWords:      []string{"pagajusaja menesi", "pagajusaja nedela", "nakamaja menesi", "nakamaja nedela", "pagajusaja gada", "priekspusdiena", "nakamaja gada", "ceturtdiena", "pecpusdiena", "saja menesi", "saja minute", "saja nedela", "saja stunda", "piektdiena", "septembris", "decembris", "februaris", "novembris", "pirmdiena", "saja gada", "sestdiena", "svetdiena", "tresdiena", "janvaris", "oktobris", "otrdiena", "sekundes", "aprilis", "augusts", "ceturtd", "menesis", "minutes", "pecpusd", "prieksp", "stundas", "julijs", "junijs", "nedela", "piektd", "sodien", "diena", "maijs", "marts", "pirmd", "sestd", "svetd", "tagad", "tresd", "vakar", "febr", "gads", "janv", "otrd", "pecp", "sept", "apr", "aug", "dec", "gmt", "jul", "jun", "men", "min", "ned", "nov", "okt", "rit", "sek", "utc", "am", "pm", "st", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "g", "h", "s", "z", "|"},
	})
}
