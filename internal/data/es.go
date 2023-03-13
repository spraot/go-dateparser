// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var es_Locale = merge(nil, LocaleData{
	Name:                  "es",
	DateOrder:             "DMY",
	Charset:               []rune(`bcdefghijlnorstuvxyzáéíñó`),
	SentenceSplitterGroup: 2,
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)una(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)un(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
	},
	Translations: map[string][]string{
		"septiembre": {"september"},
		"diciembre":  {"december"},
		"miercoles":  {"wednesday"},
		"noviembre":  {"november"},
		"setiembre":  {"september"},
		"segundos":   {"second"},
		"domingo":    {"sunday"},
		"febrero":    {"february"},
		"minutos":    {"minute"},
		"octubre":    {"october"},
		"segundo":    {"second"},
		"semanas":    {"week"},
		"viernes":    {"friday"},
		"agosto":     {"august"},
		"jueves":     {"thursday"},
		"martes":     {"tuesday"},
		"minuto":     {"minute"},
		"sabado":     {"saturday"},
		"semana":     {"week"},
		"a las":      {""},
		"abril":      {"april"},
		"cerca":      {""},
		"enero":      {"january"},
		"horas":      {"hour"},
		"julio":      {"july"},
		"junio":      {"june"},
		"lunes":      {"monday"},
		"marzo":      {"march"},
		"meses":      {"month"},
		"anos":       {"year"},
		"dias":       {"day"},
		"hace":       {"ago"},
		"hora":       {"hour"},
		"mayo":       {"may"},
		"sept":       {"september"},
		"abr":        {"april"},
		"ago":        {"august"},
		"ano":        {"year"},
		"del":        {""},
		"dia":        {"day"},
		"dic":        {"december"},
		"dom":        {"sunday"},
		"ene":        {"january"},
		"feb":        {"february"},
		"gmt":        {"gmt"},
		"jue":        {"thursday"},
		"jul":        {"july"},
		"jun":        {"june"},
		"lun":        {"monday"},
		"mar":        {"march", "tuesday"},
		"may":        {"may"},
		"mes":        {"month"},
		"mie":        {"wednesday"},
		"min":        {"minute"},
		"nov":        {"november"},
		"oct":        {"october"},
		"sab":        {"saturday"},
		"sem":        {"week"},
		"sep":        {"september"},
		"set":        {"september"},
		"utc":        {"utc"},
		"vie":        {"friday"},
		"am":         {"am"},
		"de":         {""},
		"do":         {"sunday"},
		"en":         {"in"},
		"ju":         {"thursday"},
		"lu":         {"monday"},
		"mi":         {"wednesday"},
		"pm":         {"pm"},
		"sa":         {"saturday"},
		"vi":         {"friday"},
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
		"a":          {"year"},
		"d":          {"day"},
		"h":          {"hour"},
		"m":          {"month"},
		"s":          {"second"},
		"y":          {""},
		"z":          {"z"},
		"|":          {""},
	},
	RelativeType: map[string]string{
		"la proxima semana": "in 1 week",
		"la semana pasada":  "1 week ago",
		"el proximo ano":    "in 1 year",
		"el proximo mes":    "in 1 month",
		"el ano pasado":     "1 year ago",
		"el mes pasado":     "1 month ago",
		"esta semana":       "0 week ago",
		"este minuto":       "0 minute ago",
		"esta hora":         "0 hour ago",
		"anteayer":          "2 day ago",
		"este ano":          "0 year ago",
		"este mes":          "0 month ago",
		"manana":            "in 1 day",
		"ahora":             "0 second ago",
		"ayer":              "1 day ago",
		"hoy":               "0 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)dentro de (\d+) segundos`), "in $1 second"},
		{regexp.MustCompile(`(?i)dentro de (\d+) minutos`), "in $1 minute"},
		{regexp.MustCompile(`(?i)dentro de (\d+) segundo`), "in $1 second"},
		{regexp.MustCompile(`(?i)dentro de (\d+) semanas`), "in $1 week"},
		{regexp.MustCompile(`(?i)dentro de (\d+) minuto`), "in $1 minute"},
		{regexp.MustCompile(`(?i)dentro de (\d+) semana`), "in $1 week"},
		{regexp.MustCompile(`(?i)dentro de (\d+) horas`), "in $1 hour"},
		{regexp.MustCompile(`(?i)dentro de (\d+) meses`), "in $1 month"},
		{regexp.MustCompile(`(?i)dentro de (\d+) anos`), "in $1 year"},
		{regexp.MustCompile(`(?i)dentro de (\d+) dias`), "in $1 day"},
		{regexp.MustCompile(`(?i)dentro de (\d+) hora`), "in $1 hour"},
		{regexp.MustCompile(`(?i)dentro de (\d+) ano`), "in $1 year"},
		{regexp.MustCompile(`(?i)dentro de (\d+) dia`), "in $1 day"},
		{regexp.MustCompile(`(?i)dentro de (\d+) mes`), "in $1 month"},
		{regexp.MustCompile(`(?i)dentro de (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)dentro de (\d+) sem`), "in $1 week"},
		{regexp.MustCompile(`(?i)hace (\d+) segundos`), "$1 second ago"},
		{regexp.MustCompile(`(?i)hace (\d+) minutos`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)hace (\d+) segundo`), "$1 second ago"},
		{regexp.MustCompile(`(?i)hace (\d+) semanas`), "$1 week ago"},
		{regexp.MustCompile(`(?i)dentro de (\d+) a`), "in $1 year"},
		{regexp.MustCompile(`(?i)dentro de (\d+) h`), "in $1 hour"},
		{regexp.MustCompile(`(?i)dentro de (\d+) m`), "in $1 month"},
		{regexp.MustCompile(`(?i)dentro de (\d+) s`), "in $1 second"},
		{regexp.MustCompile(`(?i)hace (\d+) minuto`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)hace (\d+) semana`), "$1 week ago"},
		{regexp.MustCompile(`(?i)hace (\d+) horas`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)hace (\d+) meses`), "$1 month ago"},
		{regexp.MustCompile(`(?i)hace (\d+) anos`), "$1 year ago"},
		{regexp.MustCompile(`(?i)hace (\d+) dias`), "$1 day ago"},
		{regexp.MustCompile(`(?i)hace (\d+) hora`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)hace (\d+) ano`), "$1 year ago"},
		{regexp.MustCompile(`(?i)hace (\d+) dia`), "$1 day ago"},
		{regexp.MustCompile(`(?i)hace (\d+) mes`), "$1 month ago"},
		{regexp.MustCompile(`(?i)hace (\d+) min`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)hace (\d+) sem`), "$1 week ago"},
		{regexp.MustCompile(`(?i)hace (\d+) a`), "$1 year ago"},
		{regexp.MustCompile(`(?i)hace (\d+) h`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)hace (\d+) m`), "$1 month ago"},
		{regexp.MustCompile(`(?i)hace (\d+) s`), "$1 second ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(dentro de \d+ segundos|dentro de \d+ minutos|dentro de \d+ segundo|dentro de \d+ semanas|dentro de \d+ minuto|dentro de \d+ semana|dentro de \d+ horas|dentro de \d+ meses|dentro de \d+ anos|dentro de \d+ dias|dentro de \d+ hora|dentro de \d+ ano|dentro de \d+ dia|dentro de \d+ mes|dentro de \d+ min|dentro de \d+ sem|hace \d+ segundos|hace \d+ minutos|hace \d+ segundo|hace \d+ semanas|dentro de \d+ a|dentro de \d+ h|dentro de \d+ m|dentro de \d+ s|hace \d+ minuto|hace \d+ semana|hace \d+ horas|hace \d+ meses|hace \d+ anos|hace \d+ dias|hace \d+ hora|hace \d+ ano|hace \d+ dia|hace \d+ mes|hace \d+ min|hace \d+ sem|hace \d+ a|hace \d+ h|hace \d+ m|hace \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(dentro de \d+ segundos|dentro de \d+ minutos|dentro de \d+ segundo|dentro de \d+ semanas|dentro de \d+ minuto|dentro de \d+ semana|dentro de \d+ horas|dentro de \d+ meses|dentro de \d+ anos|dentro de \d+ dias|dentro de \d+ hora|dentro de \d+ ano|dentro de \d+ dia|dentro de \d+ mes|dentro de \d+ min|dentro de \d+ sem|hace \d+ segundos|hace \d+ minutos|hace \d+ segundo|hace \d+ semanas|dentro de \d+ a|dentro de \d+ h|dentro de \d+ m|dentro de \d+ s|hace \d+ minuto|hace \d+ semana|hace \d+ horas|hace \d+ meses|hace \d+ anos|hace \d+ dias|hace \d+ hora|hace \d+ ano|hace \d+ dia|hace \d+ mes|hace \d+ min|hace \d+ sem|hace \d+ a|hace \d+ h|hace \d+ m|hace \d+ s)$`),
	KnownWords:      []string{"la proxima semana", "la semana pasada", "el proximo ano", "el proximo mes", "el ano pasado", "el mes pasado", "esta semana", "este minuto", "septiembre", "diciembre", "esta hora", "miercoles", "noviembre", "setiembre", "anteayer", "este ano", "este mes", "segundos", "domingo", "febrero", "minutos", "octubre", "segundo", "semanas", "viernes", "agosto", "jueves", "manana", "martes", "minuto", "sabado", "semana", "a las", "abril", "ahora", "cerca", "enero", "horas", "julio", "junio", "lunes", "marzo", "meses", "anos", "ayer", "dias", "hace", "hora", "mayo", "sept", "abr", "ago", "ano", "del", "dia", "dic", "dom", "ene", "feb", "gmt", "hoy", "jue", "jul", "jun", "lun", "mar", "may", "mes", "mie", "min", "nov", "oct", "sab", "sem", "sep", "set", "utc", "vie", "am", "de", "do", "en", "ju", "lu", "mi", "pm", "sa", "vi", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "d", "h", "m", "s", "y", "z", "|"},
})

var es_419_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-419",
	DateOrder: "DMY",
})

var es_AR_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-AR",
	DateOrder: "DMY",
	Translations: map[string][]string{
		"seg": {"second"},
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)dentro de (\d+) seg`), "in $1 second"},
		{regexp.MustCompile(`(?i)hace (\d+) seg`), "$1 second ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(dentro de \d+ segundos|dentro de \d+ minutos|dentro de \d+ segundo|dentro de \d+ semanas|dentro de \d+ minuto|dentro de \d+ semana|dentro de \d+ horas|dentro de \d+ meses|dentro de \d+ anos|dentro de \d+ dias|dentro de \d+ hora|dentro de \d+ ano|dentro de \d+ dia|dentro de \d+ mes|dentro de \d+ min|dentro de \d+ seg|dentro de \d+ sem|hace \d+ segundos|hace \d+ minutos|hace \d+ segundo|hace \d+ semanas|dentro de \d+ a|dentro de \d+ h|dentro de \d+ m|dentro de \d+ s|hace \d+ minuto|hace \d+ semana|hace \d+ horas|hace \d+ meses|hace \d+ anos|hace \d+ dias|hace \d+ hora|hace \d+ ano|hace \d+ dia|hace \d+ mes|hace \d+ min|hace \d+ seg|hace \d+ sem|hace \d+ a|hace \d+ h|hace \d+ m|hace \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(dentro de \d+ segundos|dentro de \d+ minutos|dentro de \d+ segundo|dentro de \d+ semanas|dentro de \d+ minuto|dentro de \d+ semana|dentro de \d+ horas|dentro de \d+ meses|dentro de \d+ anos|dentro de \d+ dias|dentro de \d+ hora|dentro de \d+ ano|dentro de \d+ dia|dentro de \d+ mes|dentro de \d+ min|dentro de \d+ seg|dentro de \d+ sem|hace \d+ segundos|hace \d+ minutos|hace \d+ segundo|hace \d+ semanas|dentro de \d+ a|dentro de \d+ h|dentro de \d+ m|dentro de \d+ s|hace \d+ minuto|hace \d+ semana|hace \d+ horas|hace \d+ meses|hace \d+ anos|hace \d+ dias|hace \d+ hora|hace \d+ ano|hace \d+ dia|hace \d+ mes|hace \d+ min|hace \d+ seg|hace \d+ sem|hace \d+ a|hace \d+ h|hace \d+ m|hace \d+ s)$`),
	KnownWords:      []string{"la proxima semana", "la semana pasada", "el proximo ano", "el proximo mes", "el ano pasado", "el mes pasado", "esta semana", "este minuto", "septiembre", "diciembre", "esta hora", "miercoles", "noviembre", "setiembre", "anteayer", "este ano", "este mes", "segundos", "domingo", "febrero", "minutos", "octubre", "segundo", "semanas", "viernes", "agosto", "jueves", "manana", "martes", "minuto", "sabado", "semana", "a las", "abril", "ahora", "cerca", "enero", "horas", "julio", "junio", "lunes", "marzo", "meses", "anos", "ayer", "dias", "hace", "hora", "mayo", "sept", "abr", "ago", "ano", "del", "dia", "dic", "dom", "ene", "feb", "gmt", "hoy", "jue", "jul", "jun", "lun", "mar", "may", "mes", "mie", "min", "nov", "oct", "sab", "seg", "sem", "sep", "set", "utc", "vie", "am", "de", "do", "en", "ju", "lu", "mi", "pm", "sa", "vi", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "d", "h", "m", "s", "y", "z", "|"},
})

var es_BO_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-BO",
	DateOrder: "DMY",
})

var es_BR_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-BR",
	DateOrder: "DMY",
})

var es_BZ_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-BZ",
	DateOrder: "DMY",
})

var es_CL_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-CL",
	DateOrder: "DMY",
})

var es_CO_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-CO",
	DateOrder: "DMY",
})

var es_CR_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-CR",
	DateOrder: "DMY",
})

var es_CU_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-CU",
	DateOrder: "DMY",
})

var es_DO_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-DO",
	DateOrder: "DMY",
})

var es_EA_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-EA",
	DateOrder: "DMY",
})

var es_EC_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-EC",
	DateOrder: "DMY",
})

var es_GQ_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-GQ",
	DateOrder: "DMY",
})

var es_GT_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-GT",
	DateOrder: "DMY",
})

var es_HN_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-HN",
	DateOrder: "DMY",
})

var es_IC_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-IC",
	DateOrder: "DMY",
})

var es_MX_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-MX",
	DateOrder: "DMY",
	RelativeType: map[string]string{
		"la semana proxima": "in 1 week",
		"el ano proximo":    "in 1 year",
		"el mes proximo":    "in 1 month",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)en (\d+) meses`), "in $1 month"},
		{regexp.MustCompile(`(?i)en (\d+) dias`), "in $1 day"},
		{regexp.MustCompile(`(?i)en (\d+) dia`), "in $1 day"},
		{regexp.MustCompile(`(?i)en (\d+) mes`), "in $1 month"},
		{regexp.MustCompile(`(?i)en (\d+) min`), "in $1 minute"},
		{regexp.MustCompile(`(?i)en (\d+) sem`), "in $1 week"},
		{regexp.MustCompile(`(?i)en (\d+) a`), "in $1 year"},
		{regexp.MustCompile(`(?i)en (\d+) h`), "in $1 hour"},
		{regexp.MustCompile(`(?i)en (\d+) m`), "in $1 month"},
		{regexp.MustCompile(`(?i)en (\d+) n`), "in $1 hour"},
		{regexp.MustCompile(`(?i)en (\d+) s`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(dentro de \d+ segundos|dentro de \d+ minutos|dentro de \d+ segundo|dentro de \d+ semanas|dentro de \d+ minuto|dentro de \d+ semana|dentro de \d+ horas|dentro de \d+ meses|dentro de \d+ anos|dentro de \d+ dias|dentro de \d+ hora|dentro de \d+ ano|dentro de \d+ dia|dentro de \d+ mes|dentro de \d+ min|dentro de \d+ sem|hace \d+ segundos|hace \d+ minutos|hace \d+ segundo|hace \d+ semanas|dentro de \d+ a|dentro de \d+ h|dentro de \d+ m|dentro de \d+ s|hace \d+ minuto|hace \d+ semana|hace \d+ horas|hace \d+ meses|hace \d+ anos|hace \d+ dias|hace \d+ hora|en \d+ meses|hace \d+ ano|hace \d+ dia|hace \d+ mes|hace \d+ min|hace \d+ sem|en \d+ dias|en \d+ dia|en \d+ mes|en \d+ min|en \d+ sem|hace \d+ a|hace \d+ h|hace \d+ m|hace \d+ s|en \d+ a|en \d+ h|en \d+ m|en \d+ n|en \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(dentro de \d+ segundos|dentro de \d+ minutos|dentro de \d+ segundo|dentro de \d+ semanas|dentro de \d+ minuto|dentro de \d+ semana|dentro de \d+ horas|dentro de \d+ meses|dentro de \d+ anos|dentro de \d+ dias|dentro de \d+ hora|dentro de \d+ ano|dentro de \d+ dia|dentro de \d+ mes|dentro de \d+ min|dentro de \d+ sem|hace \d+ segundos|hace \d+ minutos|hace \d+ segundo|hace \d+ semanas|dentro de \d+ a|dentro de \d+ h|dentro de \d+ m|dentro de \d+ s|hace \d+ minuto|hace \d+ semana|hace \d+ horas|hace \d+ meses|hace \d+ anos|hace \d+ dias|hace \d+ hora|en \d+ meses|hace \d+ ano|hace \d+ dia|hace \d+ mes|hace \d+ min|hace \d+ sem|en \d+ dias|en \d+ dia|en \d+ mes|en \d+ min|en \d+ sem|hace \d+ a|hace \d+ h|hace \d+ m|hace \d+ s|en \d+ a|en \d+ h|en \d+ m|en \d+ n|en \d+ s)$`),
	KnownWords:      []string{"la proxima semana", "la semana proxima", "la semana pasada", "el ano proximo", "el mes proximo", "el proximo ano", "el proximo mes", "el ano pasado", "el mes pasado", "esta semana", "este minuto", "septiembre", "diciembre", "esta hora", "miercoles", "noviembre", "setiembre", "anteayer", "este ano", "este mes", "segundos", "domingo", "febrero", "minutos", "octubre", "segundo", "semanas", "viernes", "agosto", "jueves", "manana", "martes", "minuto", "sabado", "semana", "a las", "abril", "ahora", "cerca", "enero", "horas", "julio", "junio", "lunes", "marzo", "meses", "anos", "ayer", "dias", "hace", "hora", "mayo", "sept", "abr", "ago", "ano", "del", "dia", "dic", "dom", "ene", "feb", "gmt", "hoy", "jue", "jul", "jun", "lun", "mar", "may", "mes", "mie", "min", "nov", "oct", "sab", "sem", "sep", "set", "utc", "vie", "am", "de", "do", "en", "ju", "lu", "mi", "pm", "sa", "vi", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "d", "h", "m", "s", "y", "z", "|"},
})

var es_NI_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-NI",
	DateOrder: "DMY",
})

var es_PA_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-PA",
	DateOrder: "MDY",
})

var es_PE_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-PE",
	DateOrder: "DMY",
})

var es_PH_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-PH",
	DateOrder: "DMY",
})

var es_PR_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-PR",
	DateOrder: "MDY",
})

var es_PY_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-PY",
	DateOrder: "DMY",
	Translations: map[string][]string{
		"seg": {"second"},
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)dentro de (\d+) seg`), "in $1 second"},
		{regexp.MustCompile(`(?i)hace (\d+) seg`), "$1 second ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(dentro de \d+ segundos|dentro de \d+ minutos|dentro de \d+ segundo|dentro de \d+ semanas|dentro de \d+ minuto|dentro de \d+ semana|dentro de \d+ horas|dentro de \d+ meses|dentro de \d+ anos|dentro de \d+ dias|dentro de \d+ hora|dentro de \d+ ano|dentro de \d+ dia|dentro de \d+ mes|dentro de \d+ min|dentro de \d+ seg|dentro de \d+ sem|hace \d+ segundos|hace \d+ minutos|hace \d+ segundo|hace \d+ semanas|dentro de \d+ a|dentro de \d+ h|dentro de \d+ m|dentro de \d+ s|hace \d+ minuto|hace \d+ semana|hace \d+ horas|hace \d+ meses|hace \d+ anos|hace \d+ dias|hace \d+ hora|hace \d+ ano|hace \d+ dia|hace \d+ mes|hace \d+ min|hace \d+ seg|hace \d+ sem|hace \d+ a|hace \d+ h|hace \d+ m|hace \d+ s)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(dentro de \d+ segundos|dentro de \d+ minutos|dentro de \d+ segundo|dentro de \d+ semanas|dentro de \d+ minuto|dentro de \d+ semana|dentro de \d+ horas|dentro de \d+ meses|dentro de \d+ anos|dentro de \d+ dias|dentro de \d+ hora|dentro de \d+ ano|dentro de \d+ dia|dentro de \d+ mes|dentro de \d+ min|dentro de \d+ seg|dentro de \d+ sem|hace \d+ segundos|hace \d+ minutos|hace \d+ segundo|hace \d+ semanas|dentro de \d+ a|dentro de \d+ h|dentro de \d+ m|dentro de \d+ s|hace \d+ minuto|hace \d+ semana|hace \d+ horas|hace \d+ meses|hace \d+ anos|hace \d+ dias|hace \d+ hora|hace \d+ ano|hace \d+ dia|hace \d+ mes|hace \d+ min|hace \d+ seg|hace \d+ sem|hace \d+ a|hace \d+ h|hace \d+ m|hace \d+ s)$`),
	KnownWords:      []string{"la proxima semana", "la semana pasada", "el proximo ano", "el proximo mes", "el ano pasado", "el mes pasado", "esta semana", "este minuto", "septiembre", "diciembre", "esta hora", "miercoles", "noviembre", "setiembre", "anteayer", "este ano", "este mes", "segundos", "domingo", "febrero", "minutos", "octubre", "segundo", "semanas", "viernes", "agosto", "jueves", "manana", "martes", "minuto", "sabado", "semana", "a las", "abril", "ahora", "cerca", "enero", "horas", "julio", "junio", "lunes", "marzo", "meses", "anos", "ayer", "dias", "hace", "hora", "mayo", "sept", "abr", "ago", "ano", "del", "dia", "dic", "dom", "ene", "feb", "gmt", "hoy", "jue", "jul", "jun", "lun", "mar", "may", "mes", "mie", "min", "nov", "oct", "sab", "seg", "sem", "sep", "set", "utc", "vie", "am", "de", "do", "en", "ju", "lu", "mi", "pm", "sa", "vi", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "d", "h", "m", "s", "y", "z", "|"},
})

var es_SV_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-SV",
	DateOrder: "DMY",
})

var es_US_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-US",
	DateOrder: "DMY",
})

var es_UY_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-UY",
	DateOrder: "DMY",
})

var es_VE_Locale = merge(&es_Locale, LocaleData{
	Name:      "es-VE",
	DateOrder: "DMY",
})
