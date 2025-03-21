// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	it_Locale    LocaleData
	it_CH_Locale LocaleData
	it_SM_Locale LocaleData
	it_VA_Locale LocaleData
)

func init() {
	it_Locale = merge(nil, LocaleData{
		Name:                  "it",
		DateOrder:             "DMY",
		Charset:               []rune(`bcdefghilnorstuvzì`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])(\d+[.,]?\d*)\s+ora(\z|[^\pL\pM\d_])`), "${1}${2} ore${3}"},
		},
		Translations: map[string][]string{
			"mercoledi": {"wednesday"},
			"settembre": {"september"},
			"settimana": {"week"},
			"settimane": {"week"},
			"dicembre":  {"december"},
			"domenica":  {"sunday"},
			"febbraio":  {"february"},
			"novembre":  {"november"},
			"gennaio":   {"january"},
			"giovedi":   {"thursday"},
			"martedi":   {"tuesday"},
			"ottobre":   {"october"},
			"secondi":   {"second"},
			"secondo":   {"second"},
			"venerdi":   {"friday"},
			"agosto":    {"august"},
			"aprile":    {"april"},
			"giorni":    {"day"},
			"giorno":    {"day"},
			"giugno":    {"june"},
			"luglio":    {"july"},
			"lunedi":    {"monday"},
			"maggio":    {"may"},
			"minuti":    {"minute"},
			"minuto":    {"minute"},
			"sabato":    {"saturday"},
			"circa":     {""},
			"marzo":     {"march"},
			"anni":      {"year"},
			"anno":      {"year"},
			"mese":      {"month"},
			"mesi":      {"month"},
			"sett":      {"week"},
			"ago":       {"august"},
			"apr":       {"april"},
			"dic":       {"december"},
			"dom":       {"sunday"},
			"feb":       {"february"},
			"gen":       {"january"},
			"gio":       {"thursday"},
			"giu":       {"june"},
			"gmt":       {"gmt"},
			"lug":       {"july"},
			"lun":       {"monday"},
			"mag":       {"may"},
			"mar":       {"march", "tuesday"},
			"mer":       {"wednesday"},
			"min":       {"minute"},
			"nov":       {"november"},
			"ora":       {"hour"},
			"ore":       {"hour"},
			"ott":       {"october"},
			"sab":       {"saturday"},
			"sec":       {"second"},
			"set":       {"september"},
			"utc":       {"utc"},
			"ven":       {"friday"},
			"am":        {"am"},
			"fa":        {"ago"},
			"in":        {"in"},
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
			"e":         {""},
			"g":         {"day"},
			"h":         {"hour"},
			"m":         {"minute"},
			"s":         {"second"},
			"z":         {"z"},
			"|":         {""},
		},
		RelativeType: map[string]string{
			"settimana prossima": "in 1 week",
			"questa settimana":   "0 week ago",
			"settimana scorsa":   "1 week ago",
			"anno prossimo":      "in 1 year",
			"mese prossimo":      "in 1 month",
			"questo minuto":      "0 minute ago",
			"anno scorso":        "1 year ago",
			"mese scorso":        "1 month ago",
			"questo mese":        "0 month ago",
			"altro ieri":         "2 day ago",
			"quest'anno":         "0 year ago",
			"quest'ora":          "0 hour ago",
			"domani":             "in 1 day",
			"ieri":               "1 day ago",
			"oggi":               "0 day ago",
			"ora":                "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) settimana`), "in $1 week"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) settimane`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) settimana fa`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) settimane fa`), "$1 week ago"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) secondi`), "in $1 second"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) secondo`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) secondi fa`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) secondo fa`), "$1 second ago"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) giorni`), "in $1 day"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) giorno`), "in $1 day"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) minuti`), "in $1 minute"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) minuto`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) giorni fa`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) giorno fa`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minuti fa`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minuto fa`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) anni`), "in $1 year"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) anno`), "in $1 year"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) mese`), "in $1 month"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) mesi`), "in $1 month"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) sett`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) anni fa`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) anno fa`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) mese fa`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) mesi fa`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sett fa`), "$1 week ago"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) ora`), "in $1 hour"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) ore`), "in $1 hour"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) sec`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) min fa`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ora fa`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ore fa`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sec fa`), "$1 second ago"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) gg`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) gg fa`), "$1 day ago"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) g`), "in $1 day"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) h`), "in $1 hour"},
			{regexp.MustCompile(`(?i)tra (\d+[.,]?\d*) s`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) g fa`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) h fa`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) s fa`), "$1 second ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(tra \d+[.,]?\d* settimana|tra \d+[.,]?\d* settimane|\d+[.,]?\d* settimana fa|\d+[.,]?\d* settimane fa|tra \d+[.,]?\d* secondi|tra \d+[.,]?\d* secondo|\d+[.,]?\d* secondi fa|\d+[.,]?\d* secondo fa|tra \d+[.,]?\d* giorni|tra \d+[.,]?\d* giorno|tra \d+[.,]?\d* minuti|tra \d+[.,]?\d* minuto|\d+[.,]?\d* giorni fa|\d+[.,]?\d* giorno fa|\d+[.,]?\d* minuti fa|\d+[.,]?\d* minuto fa|tra \d+[.,]?\d* anni|tra \d+[.,]?\d* anno|tra \d+[.,]?\d* mese|tra \d+[.,]?\d* mesi|tra \d+[.,]?\d* sett|\d+[.,]?\d* anni fa|\d+[.,]?\d* anno fa|\d+[.,]?\d* mese fa|\d+[.,]?\d* mesi fa|\d+[.,]?\d* sett fa|tra \d+[.,]?\d* min|tra \d+[.,]?\d* ora|tra \d+[.,]?\d* ore|tra \d+[.,]?\d* sec|\d+[.,]?\d* min fa|\d+[.,]?\d* ora fa|\d+[.,]?\d* ore fa|\d+[.,]?\d* sec fa|tra \d+[.,]?\d* gg|\d+[.,]?\d* gg fa|tra \d+[.,]?\d* g|tra \d+[.,]?\d* h|tra \d+[.,]?\d* s|\d+[.,]?\d* g fa|\d+[.,]?\d* h fa|\d+[.,]?\d* s fa)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(tra \d+[.,]?\d* settimana|tra \d+[.,]?\d* settimane|\d+[.,]?\d* settimana fa|\d+[.,]?\d* settimane fa|tra \d+[.,]?\d* secondi|tra \d+[.,]?\d* secondo|\d+[.,]?\d* secondi fa|\d+[.,]?\d* secondo fa|tra \d+[.,]?\d* giorni|tra \d+[.,]?\d* giorno|tra \d+[.,]?\d* minuti|tra \d+[.,]?\d* minuto|\d+[.,]?\d* giorni fa|\d+[.,]?\d* giorno fa|\d+[.,]?\d* minuti fa|\d+[.,]?\d* minuto fa|tra \d+[.,]?\d* anni|tra \d+[.,]?\d* anno|tra \d+[.,]?\d* mese|tra \d+[.,]?\d* mesi|tra \d+[.,]?\d* sett|\d+[.,]?\d* anni fa|\d+[.,]?\d* anno fa|\d+[.,]?\d* mese fa|\d+[.,]?\d* mesi fa|\d+[.,]?\d* sett fa|tra \d+[.,]?\d* min|tra \d+[.,]?\d* ora|tra \d+[.,]?\d* ore|tra \d+[.,]?\d* sec|\d+[.,]?\d* min fa|\d+[.,]?\d* ora fa|\d+[.,]?\d* ore fa|\d+[.,]?\d* sec fa|tra \d+[.,]?\d* gg|\d+[.,]?\d* gg fa|tra \d+[.,]?\d* g|tra \d+[.,]?\d* h|tra \d+[.,]?\d* s|\d+[.,]?\d* g fa|\d+[.,]?\d* h fa|\d+[.,]?\d* s fa)$`),
		KnownWords:      []string{"settimana prossima", "questa settimana", "settimana scorsa", "anno prossimo", "mese prossimo", "questo minuto", "anno scorso", "mese scorso", "questo mese", "altro ieri", "quest'anno", "mercoledi", "quest'ora", "settembre", "settimana", "settimane", "dicembre", "domenica", "febbraio", "novembre", "gennaio", "giovedi", "martedi", "ottobre", "secondi", "secondo", "venerdi", "agosto", "aprile", "domani", "giorni", "giorno", "giugno", "luglio", "lunedi", "maggio", "minuti", "minuto", "sabato", "circa", "marzo", "anni", "anno", "ieri", "mese", "mesi", "oggi", "sett", "ago", "apr", "dic", "dom", "feb", "gen", "gio", "giu", "gmt", "lug", "lun", "mag", "mar", "mer", "min", "nov", "ora", "ore", "ott", "sab", "sec", "set", "utc", "ven", "am", "fa", "in", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "e", "g", "h", "m", "s", "z", "|"},
	})

	it_CH_Locale = merge(&it_Locale, LocaleData{
		Name:      "it-CH",
		DateOrder: "DMY",
	})

	it_SM_Locale = merge(&it_Locale, LocaleData{
		Name:      "it-SM",
		DateOrder: "DMY",
	})

	it_VA_Locale = merge(&it_Locale, LocaleData{
		Name:      "it-VA",
		DateOrder: "DMY",
	})
}
