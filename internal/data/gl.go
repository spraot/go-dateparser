// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	gl_Locale LocaleData
)

func init() {
	gl_Locale = merge(nil, LocaleData{
		Name:      "gl",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefghilnorstuvxzáéíñó`),
		Translations: map[string][]string{
			"decembro": {"december"},
			"febreiro": {"february"},
			"mercores": {"wednesday"},
			"novembro": {"november"},
			"setembro": {"september"},
			"domingo":  {"sunday"},
			"outubro":  {"october"},
			"segundo":  {"second"},
			"xaneiro":  {"january"},
			"agosto":   {"august"},
			"martes":   {"tuesday"},
			"minuto":   {"minute"},
			"sabado":   {"saturday"},
			"semana":   {"week"},
			"venres":   {"friday"},
			"abril":    {"april"},
			"marzo":    {"march"},
			"xoves":    {"thursday"},
			"xullo":    {"july"},
			"hora":     {"hour"},
			"luns":     {"monday"},
			"maio":     {"may"},
			"xuno":     {"june"},
			"abr":      {"april"},
			"ago":      {"august"},
			"ano":      {"year"},
			"dec":      {"december"},
			"dia":      {"day"},
			"dom":      {"sunday"},
			"feb":      {"february"},
			"gmt":      {"gmt"},
			"mar":      {"march", "tuesday"},
			"mer":      {"wednesday"},
			"mes":      {"month"},
			"min":      {"minute"},
			"nov":      {"november"},
			"out":      {"october"},
			"sab":      {"saturday"},
			"sem":      {"week"},
			"set":      {"september"},
			"utc":      {"utc"},
			"ven":      {"friday"},
			"xan":      {"january"},
			"xov":      {"thursday"},
			"xul":      {"july"},
			"am":       {"am"},
			"pm":       {"pm"},
			" ":        {" "},
			"'":        {""},
			"+":        {"+"},
			",":        {""},
			"-":        {"-"},
			".":        {"."},
			"/":        {"/"},
			":":        {":"},
			";":        {""},
			"@":        {""},
			"[":        {""},
			"]":        {""},
			"a":        {"year"},
			"d":        {"day"},
			"h":        {"hour"},
			"m":        {"month"},
			"s":        {"second"},
			"z":        {"z"},
			"|":        {""},
		},
		RelativeType: map[string]string{
			"a proxima semana": "in 1 week",
			"a semana pasada":  "1 week ago",
			"o proximo ano":    "in 1 year",
			"o proximo mes":    "in 1 month",
			"neste minuto":     "0 minute ago",
			"o ano pasado":     "1 year ago",
			"o mes pasado":     "1 month ago",
			"seguinte ano":     "in 1 year",
			"sem seguinte":     "in 1 week",
			"esta semana":      "0 week ago",
			"ano pasado":       "1 year ago",
			"m seguinte":       "in 1 month",
			"nesta hora":       "0 hour ago",
			"sem pasada":       "1 week ago",
			"esta sem":         "0 week ago",
			"este ano":         "0 year ago",
			"este mes":         "0 month ago",
			"m pasado":         "1 month ago",
			"este m":           "0 month ago",
			"agora":            "0 second ago",
			"hoxe":             "0 day ago",
			"mana":             "in 1 day",
			"onte":             "1 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) segundos`), "$1 second ago"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) segundos`), "in $1 second"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) minutos`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) segundo`), "$1 second ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) semanas`), "$1 week ago"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) minutos`), "in $1 minute"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) segundo`), "in $1 second"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) semanas`), "in $1 week"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) minuto`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) semana`), "$1 week ago"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) minuto`), "in $1 minute"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) semana`), "in $1 week"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) horas`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) meses`), "$1 month ago"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) horas`), "in $1 hour"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) meses`), "in $1 month"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) anos`), "$1 year ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) dias`), "$1 day ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) hora`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) anos`), "in $1 year"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) dias`), "in $1 day"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) hora`), "in $1 hour"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) ano`), "$1 year ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) dia`), "$1 day ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) mes`), "$1 month ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) min`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) sem`), "$1 week ago"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) ano`), "in $1 year"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) dia`), "in $1 day"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) mes`), "in $1 month"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) sem`), "in $1 week"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) a`), "$1 year ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) d`), "$1 day ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) h`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) m`), "$1 month ago"},
			{regexp.MustCompile(`(?i)hai (\d+[.,]?\d*) s`), "$1 second ago"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) a`), "in $1 year"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) d`), "in $1 day"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) h`), "in $1 hour"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) m`), "in $1 month"},
			{regexp.MustCompile(`(?i)en (\d+[.,]?\d*) s`), "in $1 second"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(hai \d+[.,]?\d* segundos|en \d+[.,]?\d* segundos|hai \d+[.,]?\d* minutos|hai \d+[.,]?\d* segundo|hai \d+[.,]?\d* semanas|en \d+[.,]?\d* minutos|en \d+[.,]?\d* segundo|en \d+[.,]?\d* semanas|hai \d+[.,]?\d* minuto|hai \d+[.,]?\d* semana|en \d+[.,]?\d* minuto|en \d+[.,]?\d* semana|hai \d+[.,]?\d* horas|hai \d+[.,]?\d* meses|en \d+[.,]?\d* horas|en \d+[.,]?\d* meses|hai \d+[.,]?\d* anos|hai \d+[.,]?\d* dias|hai \d+[.,]?\d* hora|en \d+[.,]?\d* anos|en \d+[.,]?\d* dias|en \d+[.,]?\d* hora|hai \d+[.,]?\d* ano|hai \d+[.,]?\d* dia|hai \d+[.,]?\d* mes|hai \d+[.,]?\d* min|hai \d+[.,]?\d* sem|en \d+[.,]?\d* ano|en \d+[.,]?\d* dia|en \d+[.,]?\d* mes|en \d+[.,]?\d* min|en \d+[.,]?\d* sem|hai \d+[.,]?\d* a|hai \d+[.,]?\d* d|hai \d+[.,]?\d* h|hai \d+[.,]?\d* m|hai \d+[.,]?\d* s|en \d+[.,]?\d* a|en \d+[.,]?\d* d|en \d+[.,]?\d* h|en \d+[.,]?\d* m|en \d+[.,]?\d* s)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(hai \d+[.,]?\d* segundos|en \d+[.,]?\d* segundos|hai \d+[.,]?\d* minutos|hai \d+[.,]?\d* segundo|hai \d+[.,]?\d* semanas|en \d+[.,]?\d* minutos|en \d+[.,]?\d* segundo|en \d+[.,]?\d* semanas|hai \d+[.,]?\d* minuto|hai \d+[.,]?\d* semana|en \d+[.,]?\d* minuto|en \d+[.,]?\d* semana|hai \d+[.,]?\d* horas|hai \d+[.,]?\d* meses|en \d+[.,]?\d* horas|en \d+[.,]?\d* meses|hai \d+[.,]?\d* anos|hai \d+[.,]?\d* dias|hai \d+[.,]?\d* hora|en \d+[.,]?\d* anos|en \d+[.,]?\d* dias|en \d+[.,]?\d* hora|hai \d+[.,]?\d* ano|hai \d+[.,]?\d* dia|hai \d+[.,]?\d* mes|hai \d+[.,]?\d* min|hai \d+[.,]?\d* sem|en \d+[.,]?\d* ano|en \d+[.,]?\d* dia|en \d+[.,]?\d* mes|en \d+[.,]?\d* min|en \d+[.,]?\d* sem|hai \d+[.,]?\d* a|hai \d+[.,]?\d* d|hai \d+[.,]?\d* h|hai \d+[.,]?\d* m|hai \d+[.,]?\d* s|en \d+[.,]?\d* a|en \d+[.,]?\d* d|en \d+[.,]?\d* h|en \d+[.,]?\d* m|en \d+[.,]?\d* s)$`),
		KnownWords:      []string{"a proxima semana", "a semana pasada", "o proximo ano", "o proximo mes", "neste minuto", "o ano pasado", "o mes pasado", "seguinte ano", "sem seguinte", "esta semana", "ano pasado", "m seguinte", "nesta hora", "sem pasada", "decembro", "esta sem", "este ano", "este mes", "febreiro", "m pasado", "mercores", "novembro", "setembro", "domingo", "outubro", "segundo", "xaneiro", "agosto", "este m", "martes", "minuto", "sabado", "semana", "venres", "abril", "agora", "marzo", "xoves", "xullo", "hora", "hoxe", "luns", "maio", "mana", "onte", "xuno", "abr", "ago", "ano", "dec", "dia", "dom", "feb", "gmt", "mar", "mer", "mes", "min", "nov", "out", "sab", "sem", "set", "utc", "ven", "xan", "xov", "xul", "am", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "d", "h", "m", "s", "z", "|"},
	})
}
