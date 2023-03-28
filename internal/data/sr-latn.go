// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var (
	sr_Latn_Locale    LocaleData
	sr_Latn_BA_Locale LocaleData
	sr_Latn_ME_Locale LocaleData
	sr_Latn_XK_Locale LocaleData
)

func init() {
	sr_Latn_Locale = merge(nil, LocaleData{
		Name:      "sr-Latn",
		DateOrder: "DMY.",
		Charset:   []rune(`bcdefgijklnorstuvzćčš`),
		Translations: map[string][]string{
			"ponedeljak": {"monday"},
			"pre podne":  {"am"},
			"septembar":  {"september"},
			"cetvrtak":   {"thursday"},
			"decembar":   {"december"},
			"novembar":   {"november"},
			"po podne":   {"pm"},
			"februar":    {"february"},
			"nedelja":    {"week", "sunday"},
			"oktobar":    {"october"},
			"avgust":     {"august"},
			"godina":     {"year"},
			"januar":     {"january"},
			"sekund":     {"second"},
			"subota":     {"saturday"},
			"utorak":     {"tuesday"},
			"april":      {"april"},
			"mesec":      {"month"},
			"minut":      {"minute"},
			"petak":      {"friday"},
			"sreda":      {"wednesday"},
			"mart":       {"march"},
			"apr":        {"april"},
			"avg":        {"august"},
			"cet":        {"thursday"},
			"dan":        {"day"},
			"dec":        {"december"},
			"feb":        {"february"},
			"gmt":        {"gmt"},
			"god":        {"year"},
			"jan":        {"january"},
			"jul":        {"july"},
			"jun":        {"june"},
			"maj":        {"may"},
			"mar":        {"march"},
			"mes":        {"month"},
			"min":        {"minute"},
			"ned":        {"week", "sunday"},
			"nov":        {"november"},
			"okt":        {"october"},
			"pet":        {"friday"},
			"pon":        {"monday"},
			"sat":        {"hour"},
			"sek":        {"second"},
			"sep":        {"september"},
			"sre":        {"wednesday"},
			"sub":        {"saturday"},
			"utc":        {"utc"},
			"uto":        {"tuesday"},
			"am":         {"am"},
			"pm":         {"pm"},
			" ":          {" "},
			"'":          {""},
			"+":          {"+"},
			",":          {""},
			"-":          {"-"},
			".":          {"."},
			"/":          {"/"},
			":":          {":"},
			";":          {""},
			"@":          {""},
			"[":          {""},
			"]":          {""},
			"c":          {"hour"},
			"d":          {"day"},
			"g":          {"year"},
			"m":          {"month"},
			"n":          {"week"},
			"s":          {"second"},
			"z":          {"z"},
			"|":          {""},
		},
		RelativeType: map[string]string{
			"sledece nedelje": "in 1 week",
			"sledeceg meseca": "in 1 month",
			"prosle nedelje":  "1 week ago",
			"proslog meseca":  "1 month ago",
			"sledece godine":  "in 1 year",
			"prosle godine":   "1 year ago",
			"ove nedelje":     "0 week ago",
			"ovog meseca":     "0 month ago",
			"ovog minuta":     "0 minute ago",
			"ove godine":      "0 year ago",
			"ovog sata":       "0 hour ago",
			"danas":           "0 day ago",
			"sutra":           "in 1 day",
			"juce":            "1 day ago",
			"sada":            "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) nedelja`), "$1 week ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) nedelje`), "$1 week ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) sekunde`), "$1 second ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) sekundi`), "$1 second ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) godina`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) godine`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) meseca`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) meseci`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) minuta`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) nedelja`), "in $1 week"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) nedelju`), "in $1 week"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sekundi`), "in $1 second"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sekundu`), "in $1 second"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) godina`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) godinu`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) meseci`), "in $1 month"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) minuta`), "in $1 minute"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) dana`), "$1 day ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) sata`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) sati`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) mesec`), "in $1 month"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) minut`), "in $1 minute"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) god`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) mes`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) min`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) ned`), "$1 week ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) sek`), "$1 second ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) dana`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sati`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) dan`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) god`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) mes`), "in $1 month"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) ned`), "in $1 week"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sat`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sek`), "in $1 second"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) c`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) d`), "$1 day ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) g`), "$1 year ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) m`), "$1 month ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) n`), "$1 week ago"},
			{regexp.MustCompile(`(?i)pre (\d+[.,]?\d*) s`), "$1 second ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) c`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) d`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) g`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) m`), "in $1 month"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) n`), "in $1 week"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) s`), "in $1 second"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(pre \d+[.,]?\d* nedelja|pre \d+[.,]?\d* nedelje|pre \d+[.,]?\d* sekunde|pre \d+[.,]?\d* sekundi|pre \d+[.,]?\d* godina|pre \d+[.,]?\d* godine|pre \d+[.,]?\d* meseca|pre \d+[.,]?\d* meseci|pre \d+[.,]?\d* minuta|za \d+[.,]?\d* nedelja|za \d+[.,]?\d* nedelju|za \d+[.,]?\d* sekundi|za \d+[.,]?\d* sekundu|za \d+[.,]?\d* godina|za \d+[.,]?\d* godinu|za \d+[.,]?\d* meseci|za \d+[.,]?\d* minuta|pre \d+[.,]?\d* dana|pre \d+[.,]?\d* sata|pre \d+[.,]?\d* sati|za \d+[.,]?\d* mesec|za \d+[.,]?\d* minut|pre \d+[.,]?\d* god|pre \d+[.,]?\d* mes|pre \d+[.,]?\d* min|pre \d+[.,]?\d* ned|pre \d+[.,]?\d* sek|za \d+[.,]?\d* dana|za \d+[.,]?\d* sati|za \d+[.,]?\d* dan|za \d+[.,]?\d* god|za \d+[.,]?\d* mes|za \d+[.,]?\d* min|za \d+[.,]?\d* ned|za \d+[.,]?\d* sat|za \d+[.,]?\d* sek|pre \d+[.,]?\d* c|pre \d+[.,]?\d* d|pre \d+[.,]?\d* g|pre \d+[.,]?\d* m|pre \d+[.,]?\d* n|pre \d+[.,]?\d* s|za \d+[.,]?\d* c|za \d+[.,]?\d* d|za \d+[.,]?\d* g|za \d+[.,]?\d* m|za \d+[.,]?\d* n|za \d+[.,]?\d* s)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(pre \d+[.,]?\d* nedelja|pre \d+[.,]?\d* nedelje|pre \d+[.,]?\d* sekunde|pre \d+[.,]?\d* sekundi|pre \d+[.,]?\d* godina|pre \d+[.,]?\d* godine|pre \d+[.,]?\d* meseca|pre \d+[.,]?\d* meseci|pre \d+[.,]?\d* minuta|za \d+[.,]?\d* nedelja|za \d+[.,]?\d* nedelju|za \d+[.,]?\d* sekundi|za \d+[.,]?\d* sekundu|za \d+[.,]?\d* godina|za \d+[.,]?\d* godinu|za \d+[.,]?\d* meseci|za \d+[.,]?\d* minuta|pre \d+[.,]?\d* dana|pre \d+[.,]?\d* sata|pre \d+[.,]?\d* sati|za \d+[.,]?\d* mesec|za \d+[.,]?\d* minut|pre \d+[.,]?\d* god|pre \d+[.,]?\d* mes|pre \d+[.,]?\d* min|pre \d+[.,]?\d* ned|pre \d+[.,]?\d* sek|za \d+[.,]?\d* dana|za \d+[.,]?\d* sati|za \d+[.,]?\d* dan|za \d+[.,]?\d* god|za \d+[.,]?\d* mes|za \d+[.,]?\d* min|za \d+[.,]?\d* ned|za \d+[.,]?\d* sat|za \d+[.,]?\d* sek|pre \d+[.,]?\d* c|pre \d+[.,]?\d* d|pre \d+[.,]?\d* g|pre \d+[.,]?\d* m|pre \d+[.,]?\d* n|pre \d+[.,]?\d* s|za \d+[.,]?\d* c|za \d+[.,]?\d* d|za \d+[.,]?\d* g|za \d+[.,]?\d* m|za \d+[.,]?\d* n|za \d+[.,]?\d* s)$`),
		KnownWords:      []string{"sledece nedelje", "sledeceg meseca", "prosle nedelje", "proslog meseca", "sledece godine", "prosle godine", "ove nedelje", "ovog meseca", "ovog minuta", "ove godine", "ponedeljak", "ovog sata", "pre podne", "septembar", "cetvrtak", "decembar", "novembar", "po podne", "februar", "nedelja", "oktobar", "avgust", "godina", "januar", "sekund", "subota", "utorak", "april", "danas", "mesec", "minut", "petak", "sreda", "sutra", "juce", "mart", "sada", "apr", "avg", "cet", "dan", "dec", "feb", "gmt", "god", "jan", "jul", "jun", "maj", "mar", "mes", "min", "ned", "nov", "okt", "pet", "pon", "sat", "sek", "sep", "sre", "sub", "utc", "uto", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "c", "d", "g", "m", "n", "s", "z", "|"},
	})

	sr_Latn_BA_Locale = merge(&sr_Latn_Locale, LocaleData{
		Name:      "sr-Latn-BA",
		DateOrder: "DMY.",
		Translations: map[string][]string{
			"prije podne": {"am"},
			"nedjelja":    {"sunday"},
			"srijeda":     {"wednesday"},
			"sept":        {"september"},
			"sr":          {"wednesday"},
			"ut":          {"tuesday"},
		},
		KnownWords: []string{"sledece nedelje", "sledeceg meseca", "prosle nedelje", "proslog meseca", "sledece godine", "prosle godine", "ove nedelje", "ovog meseca", "ovog minuta", "prije podne", "ove godine", "ponedeljak", "ovog sata", "pre podne", "septembar", "cetvrtak", "decembar", "nedjelja", "novembar", "po podne", "februar", "nedelja", "oktobar", "srijeda", "avgust", "godina", "januar", "sekund", "subota", "utorak", "april", "danas", "mesec", "minut", "petak", "sreda", "sutra", "juce", "mart", "sada", "sept", "apr", "avg", "cet", "dan", "dec", "feb", "gmt", "god", "jan", "jul", "jun", "maj", "mar", "mes", "min", "ned", "nov", "okt", "pet", "pon", "sat", "sek", "sep", "sre", "sub", "utc", "uto", "am", "pm", "sr", "ut", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "c", "d", "g", "m", "n", "s", "z", "|"},
	})

	sr_Latn_ME_Locale = merge(&sr_Latn_Locale, LocaleData{
		Name:      "sr-Latn-ME",
		DateOrder: "DMY.",
		Translations: map[string][]string{
			"prije podne": {"am"},
			"nedjelja":    {"sunday"},
			"srijeda":     {"wednesday"},
			"sept":        {"september"},
			"sr":          {"wednesday"},
			"ut":          {"tuesday"},
		},
		KnownWords: []string{"sledece nedelje", "sledeceg meseca", "prosle nedelje", "proslog meseca", "sledece godine", "prosle godine", "ove nedelje", "ovog meseca", "ovog minuta", "prije podne", "ove godine", "ponedeljak", "ovog sata", "pre podne", "septembar", "cetvrtak", "decembar", "nedjelja", "novembar", "po podne", "februar", "nedelja", "oktobar", "srijeda", "avgust", "godina", "januar", "sekund", "subota", "utorak", "april", "danas", "mesec", "minut", "petak", "sreda", "sutra", "juce", "mart", "sada", "sept", "apr", "avg", "cet", "dan", "dec", "feb", "gmt", "god", "jan", "jul", "jun", "maj", "mar", "mes", "min", "ned", "nov", "okt", "pet", "pon", "sat", "sek", "sep", "sre", "sub", "utc", "uto", "am", "pm", "sr", "ut", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "c", "d", "g", "m", "n", "s", "z", "|"},
	})

	sr_Latn_XK_Locale = merge(&sr_Latn_Locale, LocaleData{
		Name:      "sr-Latn-XK",
		DateOrder: "DMY.",
		Translations: map[string][]string{
			"sept": {"september"},
			"sr":   {"wednesday"},
			"ut":   {"tuesday"},
		},
		KnownWords: []string{"sledece nedelje", "sledeceg meseca", "prosle nedelje", "proslog meseca", "sledece godine", "prosle godine", "ove nedelje", "ovog meseca", "ovog minuta", "ove godine", "ponedeljak", "ovog sata", "pre podne", "septembar", "cetvrtak", "decembar", "novembar", "po podne", "februar", "nedelja", "oktobar", "avgust", "godina", "januar", "sekund", "subota", "utorak", "april", "danas", "mesec", "minut", "petak", "sreda", "sutra", "juce", "mart", "sada", "sept", "apr", "avg", "cet", "dan", "dec", "feb", "gmt", "god", "jan", "jul", "jun", "maj", "mar", "mes", "min", "ned", "nov", "okt", "pet", "pon", "sat", "sek", "sep", "sre", "sub", "utc", "uto", "am", "pm", "sr", "ut", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "c", "d", "g", "m", "n", "s", "z", "|"},
	})
}
