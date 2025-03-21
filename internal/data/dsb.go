// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	dsb_Locale LocaleData
)

func init() {
	dsb_Locale = merge(nil, LocaleData{
		Name:      "dsb",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefghijklnorstuwyzóěłńśź`),
		Translations: map[string][]string{
			"wotpołdnja": {"pm"},
			"dopołdnja":  {"am"},
			"ponjezele":  {"monday"},
			"september":  {"september"},
			"septembra":  {"september"},
			"december":   {"december"},
			"decembra":   {"december"},
			"februara":   {"february"},
			"nowember":   {"november"},
			"nowembra":   {"november"},
			"awgusta":    {"august"},
			"februar":    {"february"},
			"januara":    {"january"},
			"njezela":    {"sunday"},
			"oktober":    {"october"},
			"oktobra":    {"october"},
			"sekunda":    {"second"},
			"stwortk":    {"thursday"},
			"wałtora":    {"tuesday"},
			"apryla":     {"april"},
			"awgust":     {"august"},
			"gozina":     {"hour"},
			"januar":     {"january"},
			"julija":     {"july"},
			"junija":     {"june"},
			"minuta":     {"minute"},
			"mjasec":     {"month"},
			"sobota":     {"saturday"},
			"srjoda":     {"wednesday"},
			"apryl":      {"april"},
			"julij":      {"july"},
			"junij":      {"june"},
			"merca":      {"march"},
			"tyzen":      {"week"},
			"leto":       {"year"},
			"maja":       {"may"},
			"merc":       {"march"},
			"mjas":       {"month"},
			"petk":       {"friday"},
			"apr":        {"april"},
			"awg":        {"august"},
			"dec":        {"december"},
			"feb":        {"february"},
			"gmt":        {"gmt"},
			"goz":        {"hour"},
			"jan":        {"january"},
			"jul":        {"july"},
			"jun":        {"june"},
			"maj":        {"may"},
			"mer":        {"march"},
			"min":        {"minute"},
			"nje":        {"sunday"},
			"now":        {"november"},
			"okt":        {"october"},
			"pet":        {"friday"},
			"pon":        {"monday"},
			"sek":        {"second"},
			"sep":        {"september"},
			"sob":        {"saturday"},
			"srj":        {"wednesday"},
			"stw":        {"thursday"},
			"tyz":        {"week"},
			"utc":        {"utc"},
			"wał":        {"tuesday"},
			"zen":        {"day"},
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
			"g":          {"hour"},
			"l":          {"year"},
			"m":          {"minute"},
			"s":          {"second"},
			"z":          {"z"},
			"|":          {""},
		},
		RelativeType: map[string]string{
			"psiducy mjasec": "in 1 month",
			"psiducy tyzen":  "in 1 week",
			"sledny mjasec":  "1 month ago",
			"sledny tyzen":   "1 week ago",
			"this minute":    "0 minute ago",
			"ten mjasec":     "0 month ago",
			"ten tyzen":      "0 week ago",
			"this hour":      "0 hour ago",
			"letosa":         "0 year ago",
			"witse":          "in 1 day",
			"zinsa":          "0 day ago",
			"znowa":          "in 1 year",
			"cora":           "1 day ago",
			"łoni":           "1 year ago",
			"now":            "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) mjasecami`), "$1 month ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) sekundami`), "$1 second ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) tyzenjami`), "$1 week ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) gozinami`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) minutami`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) mjasecom`), "$1 month ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) tyzenjom`), "$1 week ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) sekundu`), "$1 second ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) dnjami`), "$1 day ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) gozinu`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) letami`), "$1 year ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) minutu`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) mjasecow`), "in $1 month"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sekundow`), "in $1 second"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) tyzenjow`), "in $1 week"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) dnjom`), "$1 day ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) letom`), "$1 year ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) minutow`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sekundu`), "in $1 second"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) mjas`), "$1 month ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) gozinu`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) minutu`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) mjasec`), "in $1 month"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) dnj`), "$1 day ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) goz`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) min`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) sek`), "$1 second ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) tyz`), "$1 week ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) dnjow`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) gozin`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) tyzen`), "in $1 week"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) leto`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) mjas`), "in $1 month"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) d`), "$1 day ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) g`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) l`), "$1 year ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) m`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)psed (\d+[.,]?\d*) s`), "$1 second ago"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) dnj`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) goz`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) let`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) sek`), "in $1 second"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) tyz`), "in $1 week"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) zen`), "in $1 day"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) g`), "in $1 hour"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) l`), "in $1 year"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) m`), "in $1 minute"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) s`), "in $1 second"},
			{regexp.MustCompile(`(?i)za (\d+[.,]?\d*) z`), "in $1 day"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(psed \d+[.,]?\d* mjasecami|psed \d+[.,]?\d* sekundami|psed \d+[.,]?\d* tyzenjami|psed \d+[.,]?\d* gozinami|psed \d+[.,]?\d* minutami|psed \d+[.,]?\d* mjasecom|psed \d+[.,]?\d* tyzenjom|psed \d+[.,]?\d* sekundu|psed \d+[.,]?\d* dnjami|psed \d+[.,]?\d* gozinu|psed \d+[.,]?\d* letami|psed \d+[.,]?\d* minutu|za \d+[.,]?\d* mjasecow|za \d+[.,]?\d* sekundow|za \d+[.,]?\d* tyzenjow|psed \d+[.,]?\d* dnjom|psed \d+[.,]?\d* letom|za \d+[.,]?\d* minutow|za \d+[.,]?\d* sekundu|psed \d+[.,]?\d* mjas|za \d+[.,]?\d* gozinu|za \d+[.,]?\d* minutu|za \d+[.,]?\d* mjasec|psed \d+[.,]?\d* dnj|psed \d+[.,]?\d* goz|psed \d+[.,]?\d* min|psed \d+[.,]?\d* sek|psed \d+[.,]?\d* tyz|za \d+[.,]?\d* dnjow|za \d+[.,]?\d* gozin|za \d+[.,]?\d* tyzen|za \d+[.,]?\d* leto|za \d+[.,]?\d* mjas|psed \d+[.,]?\d* d|psed \d+[.,]?\d* g|psed \d+[.,]?\d* l|psed \d+[.,]?\d* m|psed \d+[.,]?\d* s|za \d+[.,]?\d* dnj|za \d+[.,]?\d* goz|za \d+[.,]?\d* let|za \d+[.,]?\d* min|za \d+[.,]?\d* sek|za \d+[.,]?\d* tyz|za \d+[.,]?\d* zen|za \d+[.,]?\d* g|za \d+[.,]?\d* l|za \d+[.,]?\d* m|za \d+[.,]?\d* s|za \d+[.,]?\d* z)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(psed \d+[.,]?\d* mjasecami|psed \d+[.,]?\d* sekundami|psed \d+[.,]?\d* tyzenjami|psed \d+[.,]?\d* gozinami|psed \d+[.,]?\d* minutami|psed \d+[.,]?\d* mjasecom|psed \d+[.,]?\d* tyzenjom|psed \d+[.,]?\d* sekundu|psed \d+[.,]?\d* dnjami|psed \d+[.,]?\d* gozinu|psed \d+[.,]?\d* letami|psed \d+[.,]?\d* minutu|za \d+[.,]?\d* mjasecow|za \d+[.,]?\d* sekundow|za \d+[.,]?\d* tyzenjow|psed \d+[.,]?\d* dnjom|psed \d+[.,]?\d* letom|za \d+[.,]?\d* minutow|za \d+[.,]?\d* sekundu|psed \d+[.,]?\d* mjas|za \d+[.,]?\d* gozinu|za \d+[.,]?\d* minutu|za \d+[.,]?\d* mjasec|psed \d+[.,]?\d* dnj|psed \d+[.,]?\d* goz|psed \d+[.,]?\d* min|psed \d+[.,]?\d* sek|psed \d+[.,]?\d* tyz|za \d+[.,]?\d* dnjow|za \d+[.,]?\d* gozin|za \d+[.,]?\d* tyzen|za \d+[.,]?\d* leto|za \d+[.,]?\d* mjas|psed \d+[.,]?\d* d|psed \d+[.,]?\d* g|psed \d+[.,]?\d* l|psed \d+[.,]?\d* m|psed \d+[.,]?\d* s|za \d+[.,]?\d* dnj|za \d+[.,]?\d* goz|za \d+[.,]?\d* let|za \d+[.,]?\d* min|za \d+[.,]?\d* sek|za \d+[.,]?\d* tyz|za \d+[.,]?\d* zen|za \d+[.,]?\d* g|za \d+[.,]?\d* l|za \d+[.,]?\d* m|za \d+[.,]?\d* s|za \d+[.,]?\d* z)$`),
		KnownWords:      []string{"psiducy mjasec", "psiducy tyzen", "sledny mjasec", "sledny tyzen", "this minute", "ten mjasec", "wotpołdnja", "dopołdnja", "ponjezele", "september", "septembra", "ten tyzen", "this hour", "december", "decembra", "februara", "nowember", "nowembra", "awgusta", "februar", "januara", "njezela", "oktober", "oktobra", "sekunda", "stwortk", "wałtora", "apryla", "awgust", "gozina", "januar", "julija", "junija", "letosa", "minuta", "mjasec", "sobota", "srjoda", "apryl", "julij", "junij", "merca", "tyzen", "witse", "zinsa", "znowa", "cora", "leto", "maja", "merc", "mjas", "petk", "łoni", "apr", "awg", "dec", "feb", "gmt", "goz", "jan", "jul", "jun", "maj", "mer", "min", "nje", "now", "okt", "pet", "pon", "sek", "sep", "sob", "srj", "stw", "tyz", "utc", "wał", "zen", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "g", "l", "m", "s", "z", "|"},
	})
}
