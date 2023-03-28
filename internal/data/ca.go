// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var (
	ca_Locale    LocaleData
	ca_AD_Locale LocaleData
	ca_FR_Locale LocaleData
	ca_IT_Locale LocaleData
)

func init() {
	ca_Locale = merge(nil, LocaleData{
		Name:      "ca",
		DateOrder: "DMY",
		Charset:   []rune(`-bcdefghijlnorstuvxyzàçòú`),
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)una(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)un(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		},
		Translations: map[string][]string{
			"de desembre": {"december"},
			"de novembre": {"november"},
			"de setembre": {"september"},
			"d'octubre":   {"october"},
			"de febrer":   {"february"},
			"de juliol":   {"july"},
			"divendres":   {"friday"},
			"de gener":    {"january"},
			"desembre":    {"december"},
			"dimecres":    {"wednesday"},
			"dissabte":    {"saturday"},
			"diumenge":    {"sunday"},
			"novembre":    {"november"},
			"setembre":    {"september"},
			"d'abril":     {"april"},
			"d'agost":     {"august"},
			"de febr":     {"february"},
			"de juny":     {"june"},
			"de maig":     {"may"},
			"de marc":     {"march"},
			"dilluns":     {"monday"},
			"dimarts":     {"tuesday"},
			"octubre":     {"october"},
			"setmana":     {"week"},
			"de des":      {"december"},
			"de gen":      {"january"},
			"de jul":      {"july"},
			"de nov":      {"november"},
			"de set":      {"september"},
			"dijous":      {"thursday"},
			"febrer":      {"february"},
			"juliol":      {"july"},
			"abril":       {"april"},
			"agost":       {"august"},
			"d'abr":       {"april"},
			"d'oct":       {"october"},
			"gener":       {"january"},
			"minut":       {"minute"},
			"segon":       {"second"},
			"d'ag":        {"august"},
			"febr":        {"february"},
			"hora":        {"hour"},
			"juny":        {"june"},
			"maig":        {"may"},
			"marc":        {"march"},
			"setm":        {"week"},
			"abr":         {"april"},
			"any":         {"year"},
			"del":         {""},
			"des":         {"december"},
			"dia":         {"day"},
			"gen":         {"january"},
			"gmt":         {"gmt"},
			"jul":         {"july"},
			"mes":         {"month"},
			"min":         {"minute"},
			"nov":         {"november"},
			"oct":         {"october"},
			"set":         {"september"},
			"utc":         {"utc"},
			"ag":          {"august"},
			"am":          {"am"},
			"dc":          {"wednesday"},
			"de":          {""},
			"dg":          {"sunday"},
			"dj":          {"thursday"},
			"dl":          {"monday"},
			"ds":          {"saturday"},
			"dt":          {"tuesday"},
			"dv":          {"friday"},
			"en":          {"in"},
			"l'":          {""},
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
			"i":           {""},
			"s":           {"second"},
			"z":           {"z"},
			"|":           {""},
		},
		RelativeType: map[string]string{
			"la propera setmana": "in 1 week",
			"la proxima setmana": "in 1 week",
			"la setmana passada": "1 week ago",
			"la setmana que ve":  "in 1 week",
			"la setmana vinent":  "in 1 week",
			"aquesta setmana":    "0 week ago",
			"la setm passada":    "1 week ago",
			"la setm que ve":     "in 1 week",
			"el mes passat":      "1 month ago",
			"el mes que ve":      "in 1 month",
			"endema passat":      "in 3 day",
			"abans-d’ahir":       "2 day ago",
			"aquest minut":       "0 minute ago",
			"aquesta hora":       "0 hour ago",
			"aquesta setm":       "0 week ago",
			"l'any passat":       "1 year ago",
			"l'any que ve":       "in 1 year",
			"setm passada":       "1 week ago",
			"dema passat":        "in 2 day",
			"despus-ahir":        "2 day ago",
			"despus-dema":        "in 2 day",
			"passat dema":        "in 2 day",
			"setm vinent":        "in 1 week",
			"aquest mes":         "0 month ago",
			"della-ahir":         "2 day ago",
			"mes passat":         "1 month ago",
			"mes vinent":         "in 1 month",
			"enguany":            "0 year ago",
			"sendema":            "in 2 day",
			"endema":             "in 2 day",
			"ahir":               "1 day ago",
			"avui":               "0 day ago",
			"dema":               "in 1 day",
			"ara":                "0 second ago",
			"hui":                "0 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) setmanes`), "in $1 week"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) setmana`), "in $1 week"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) minuts`), "in $1 minute"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) segons`), "in $1 second"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) hores`), "in $1 hour"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) mesos`), "in $1 month"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) minut`), "in $1 minute"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) segon`), "in $1 second"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) anys`), "in $1 year"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) dies`), "in $1 day"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) hora`), "in $1 hour"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) setm`), "in $1 week"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) any`), "in $1 year"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) dia`), "in $1 day"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) mes`), "in $1 month"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) setmanes`), "$1 week ago"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) h`), "in $1 hour"},
			{regexp.MustCompile(`(?i)d'aqui a (\d+[.,]?\d*) s`), "in $1 second"},
			{regexp.MustCompile(`(?i)d‘aqui a (\d+[.,]?\d*) h`), "in $1 hour"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) setmana`), "$1 week ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) minuts`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) segons`), "$1 second ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) hores`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) mesos`), "$1 month ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) minut`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) segon`), "$1 second ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) anys`), "$1 year ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) dies`), "$1 day ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) hora`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) setm`), "$1 week ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) any`), "$1 year ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) dia`), "$1 day ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) mes`), "$1 month ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) min`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) h`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)fa (\d+[.,]?\d*) s`), "$1 second ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(d'aqui a \d+[.,]?\d* setmanes|d'aqui a \d+[.,]?\d* setmana|d'aqui a \d+[.,]?\d* minuts|d'aqui a \d+[.,]?\d* segons|d'aqui a \d+[.,]?\d* hores|d'aqui a \d+[.,]?\d* mesos|d'aqui a \d+[.,]?\d* minut|d'aqui a \d+[.,]?\d* segon|d'aqui a \d+[.,]?\d* anys|d'aqui a \d+[.,]?\d* dies|d'aqui a \d+[.,]?\d* hora|d'aqui a \d+[.,]?\d* setm|d'aqui a \d+[.,]?\d* any|d'aqui a \d+[.,]?\d* dia|d'aqui a \d+[.,]?\d* mes|d'aqui a \d+[.,]?\d* min|fa \d+[.,]?\d* setmanes|d'aqui a \d+[.,]?\d* h|d'aqui a \d+[.,]?\d* s|d‘aqui a \d+[.,]?\d* h|fa \d+[.,]?\d* setmana|fa \d+[.,]?\d* minuts|fa \d+[.,]?\d* segons|fa \d+[.,]?\d* hores|fa \d+[.,]?\d* mesos|fa \d+[.,]?\d* minut|fa \d+[.,]?\d* segon|fa \d+[.,]?\d* anys|fa \d+[.,]?\d* dies|fa \d+[.,]?\d* hora|fa \d+[.,]?\d* setm|fa \d+[.,]?\d* any|fa \d+[.,]?\d* dia|fa \d+[.,]?\d* mes|fa \d+[.,]?\d* min|fa \d+[.,]?\d* h|fa \d+[.,]?\d* s)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(d'aqui a \d+[.,]?\d* setmanes|d'aqui a \d+[.,]?\d* setmana|d'aqui a \d+[.,]?\d* minuts|d'aqui a \d+[.,]?\d* segons|d'aqui a \d+[.,]?\d* hores|d'aqui a \d+[.,]?\d* mesos|d'aqui a \d+[.,]?\d* minut|d'aqui a \d+[.,]?\d* segon|d'aqui a \d+[.,]?\d* anys|d'aqui a \d+[.,]?\d* dies|d'aqui a \d+[.,]?\d* hora|d'aqui a \d+[.,]?\d* setm|d'aqui a \d+[.,]?\d* any|d'aqui a \d+[.,]?\d* dia|d'aqui a \d+[.,]?\d* mes|d'aqui a \d+[.,]?\d* min|fa \d+[.,]?\d* setmanes|d'aqui a \d+[.,]?\d* h|d'aqui a \d+[.,]?\d* s|d‘aqui a \d+[.,]?\d* h|fa \d+[.,]?\d* setmana|fa \d+[.,]?\d* minuts|fa \d+[.,]?\d* segons|fa \d+[.,]?\d* hores|fa \d+[.,]?\d* mesos|fa \d+[.,]?\d* minut|fa \d+[.,]?\d* segon|fa \d+[.,]?\d* anys|fa \d+[.,]?\d* dies|fa \d+[.,]?\d* hora|fa \d+[.,]?\d* setm|fa \d+[.,]?\d* any|fa \d+[.,]?\d* dia|fa \d+[.,]?\d* mes|fa \d+[.,]?\d* min|fa \d+[.,]?\d* h|fa \d+[.,]?\d* s)$`),
		KnownWords:      []string{"la propera setmana", "la proxima setmana", "la setmana passada", "la setmana que ve", "la setmana vinent", "aquesta setmana", "la setm passada", "la setm que ve", "el mes passat", "el mes que ve", "endema passat", "abans-d’ahir", "aquest minut", "aquesta hora", "aquesta setm", "l'any passat", "l'any que ve", "setm passada", "de desembre", "de novembre", "de setembre", "dema passat", "despus-ahir", "despus-dema", "passat dema", "setm vinent", "aquest mes", "della-ahir", "mes passat", "mes vinent", "d'octubre", "de febrer", "de juliol", "divendres", "de gener", "desembre", "dimecres", "dissabte", "diumenge", "novembre", "setembre", "d'abril", "d'agost", "de febr", "de juny", "de maig", "de marc", "dilluns", "dimarts", "enguany", "octubre", "sendema", "setmana", "de des", "de gen", "de jul", "de nov", "de set", "dijous", "endema", "febrer", "juliol", "abril", "agost", "d'abr", "d'oct", "gener", "minut", "segon", "ahir", "avui", "d'ag", "dema", "febr", "hora", "juny", "maig", "marc", "setm", "abr", "any", "ara", "del", "des", "dia", "gen", "gmt", "hui", "jul", "mes", "min", "nov", "oct", "set", "utc", "ag", "am", "dc", "de", "dg", "dj", "dl", "ds", "dt", "dv", "en", "l'", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "h", "i", "s", "z", "|"},
	})

	ca_AD_Locale = merge(&ca_Locale, LocaleData{
		Name:      "ca-AD",
		DateOrder: "DMY",
	})

	ca_FR_Locale = merge(&ca_Locale, LocaleData{
		Name:      "ca-FR",
		DateOrder: "DMY",
	})

	ca_IT_Locale = merge(&ca_Locale, LocaleData{
		Name:      "ca-IT",
		DateOrder: "DMY",
	})
}
