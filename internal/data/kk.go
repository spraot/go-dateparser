// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var (
	kk_Locale LocaleData
)

func init() {
	kk_Locale = merge(nil, LocaleData{
		Name:      "kk",
		DateOrder: "DMY",
		Charset:   []rune(`cgtuzабгдежзийклмнопрстушыіғқңүұәө`),
		Translations: map[string][]string{
			"желтоқсан": {"december"},
			"беисенбі":  {"thursday"},
			"дүисенбі":  {"monday"},
			"жексенбі":  {"sunday"},
			"сеисенбі":  {"tuesday"},
			"сәрсенбі":  {"wednesday"},
			"қыркүиек":  {"september"},
			"маусым":    {"june"},
			"наурыз":    {"march"},
			"секунд":    {"second"},
			"қараша":    {"november"},
			"қаңтар":    {"january"},
			"ақпан":     {"february"},
			"мамыр":     {"may"},
			"минут":     {"minute"},
			"сағат":     {"hour"},
			"сенбі":     {"saturday"},
			"сәуір":     {"april"},
			"тамыз":     {"august"},
			"шілде":     {"july"},
			"қазан":     {"october"},
			"апта":      {"week"},
			"жұма":      {"friday"},
			"gmt":       {"gmt"},
			"utc":       {"utc"},
			"ақп":       {"february"},
			"жел":       {"december"},
			"жыл":       {"year"},
			"күн":       {"day"},
			"мам":       {"may"},
			"мау":       {"june"},
			"мин":       {"minute"},
			"нау":       {"march"},
			"сағ":       {"hour"},
			"сәу":       {"april"},
			"там":       {"august"},
			"шіл":       {"july"},
			"қаз":       {"october"},
			"қар":       {"november"},
			"қаң":       {"january"},
			"қыр":       {"september"},
			"am":        {"am"},
			"pm":        {"pm"},
			"аи":        {"month"},
			"ап":        {"week"},
			"бс":        {"thursday"},
			"дс":        {"monday"},
			"жм":        {"friday"},
			"жс":        {"sunday"},
			"сб":        {"saturday"},
			"ср":        {"wednesday"},
			"сс":        {"tuesday"},
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
			"z":         {"z"},
			"|":         {""},
			"ж":         {"year"},
			"с":         {"second"},
		},
		RelativeType: map[string]string{
			"былтырғы жыл": "1 year ago",
			"келесі апта":  "in 1 week",
			"биылғы жыл":   "0 year ago",
			"келесі жыл":   "in 1 year",
			"өткен апта":   "1 week ago",
			"келесі аи":    "in 1 month",
			"осы минут":    "0 minute ago",
			"осы сағат":    "0 hour ago",
			"осы апта":     "0 week ago",
			"өткен аи":     "1 month ago",
			"осы аи":       "0 month ago",
			"бүгін":        "0 day ago",
			"ертең":        "in 1 day",
			"қазір":        "0 second ago",
			"кеше":         "1 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) секундтан кеиін`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) минуттан кеиін`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сағаттан кеиін`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) аптадан кеиін`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) жылдан кеиін`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) күннен кеиін`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) секунд бұрын`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) аидан кеиін`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) минут бұрын`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сағат бұрын`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) апта бұрын`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) жыл бұрын`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) күн бұрын`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) мин бұрын`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) мин кеиін`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сағ бұрын`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сағ кеиін`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сек бұрын`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) сек кеиін`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) аи бұрын`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ап бұрын`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ап кеиін`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ж бұрын`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ж кеиін`), "in $1 year"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* секундтан кеиін|\d+[.,]?\d* минуттан кеиін|\d+[.,]?\d* сағаттан кеиін|\d+[.,]?\d* аптадан кеиін|\d+[.,]?\d* жылдан кеиін|\d+[.,]?\d* күннен кеиін|\d+[.,]?\d* секунд бұрын|\d+[.,]?\d* аидан кеиін|\d+[.,]?\d* минут бұрын|\d+[.,]?\d* сағат бұрын|\d+[.,]?\d* апта бұрын|\d+[.,]?\d* жыл бұрын|\d+[.,]?\d* күн бұрын|\d+[.,]?\d* мин бұрын|\d+[.,]?\d* мин кеиін|\d+[.,]?\d* сағ бұрын|\d+[.,]?\d* сағ кеиін|\d+[.,]?\d* сек бұрын|\d+[.,]?\d* сек кеиін|\d+[.,]?\d* аи бұрын|\d+[.,]?\d* ап бұрын|\d+[.,]?\d* ап кеиін|\d+[.,]?\d* ж бұрын|\d+[.,]?\d* ж кеиін)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* секундтан кеиін|\d+[.,]?\d* минуттан кеиін|\d+[.,]?\d* сағаттан кеиін|\d+[.,]?\d* аптадан кеиін|\d+[.,]?\d* жылдан кеиін|\d+[.,]?\d* күннен кеиін|\d+[.,]?\d* секунд бұрын|\d+[.,]?\d* аидан кеиін|\d+[.,]?\d* минут бұрын|\d+[.,]?\d* сағат бұрын|\d+[.,]?\d* апта бұрын|\d+[.,]?\d* жыл бұрын|\d+[.,]?\d* күн бұрын|\d+[.,]?\d* мин бұрын|\d+[.,]?\d* мин кеиін|\d+[.,]?\d* сағ бұрын|\d+[.,]?\d* сағ кеиін|\d+[.,]?\d* сек бұрын|\d+[.,]?\d* сек кеиін|\d+[.,]?\d* аи бұрын|\d+[.,]?\d* ап бұрын|\d+[.,]?\d* ап кеиін|\d+[.,]?\d* ж бұрын|\d+[.,]?\d* ж кеиін)$`),
		KnownWords:      []string{"былтырғы жыл", "келесі апта", "биылғы жыл", "келесі жыл", "өткен апта", "желтоқсан", "келесі аи", "осы минут", "осы сағат", "беисенбі", "дүисенбі", "жексенбі", "осы апта", "сеисенбі", "сәрсенбі", "қыркүиек", "өткен аи", "маусым", "наурыз", "осы аи", "секунд", "қараша", "қаңтар", "ақпан", "бүгін", "ертең", "мамыр", "минут", "сағат", "сенбі", "сәуір", "тамыз", "шілде", "қазан", "қазір", "апта", "жұма", "кеше", "gmt", "utc", "ақп", "жел", "жыл", "күн", "мам", "мау", "мин", "нау", "сағ", "сәу", "там", "шіл", "қаз", "қар", "қаң", "қыр", "am", "pm", "аи", "ап", "бс", "дс", "жм", "жс", "сб", "ср", "сс", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "ж", "с"},
	})
}
