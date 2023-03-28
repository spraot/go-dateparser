// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var (
	tl_Locale LocaleData
)

func init() {
	tl_Locale = merge(nil, LocaleData{
		Name:                  "tl",
		DateOrder:             "",
		Charset:               []rune(`bcdeghiklnorstuwyz`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)kahapon(\z|[^\pL\pM\d]|_)`), "${1}1 araw nakaraan${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)ngayon(\z|[^\pL\pM\d]|_)`), "${1}0 segundo nakalipas${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)isang araw(\z|[^\pL\pM\d]|_)`), "${1}2 araw${2}"},
		},
		Translations: map[string][]string{
			"miyerkules": {"wednesday"},
			"disyembre":  {"december"},
			"nakalipas":  {"ago"},
			"nobyembre":  {"november"},
			"setyembre":  {"september"},
			"biyernes":   {"friday"},
			"nakaraan":   {"ago"},
			"huwebes":    {"thursday"},
			"oktubre":    {"october"},
			"pebrero":    {"february"},
			"segundo":    {"second"},
			"agosto":     {"august"},
			"linggo":     {"week", "sunday"},
			"martes":     {"tuesday"},
			"minuto":     {"minute"},
			"sabado":     {"saturday"},
			"abril":      {"april"},
			"buwan":      {"month"},
			"enero":      {"january"},
			"ganap":      {""},
			"hulyo":      {"july"},
			"hunyo":      {"june"},
			"lunes":      {"monday"},
			"marso":      {"march"},
			"noong":      {""},
			"araw":       {"day"},
			"mayo":       {"may"},
			"noon":       {""},
			"oras":       {"hour"},
			"taon":       {"year"},
			"abr":        {"april"},
			"ago":        {"august"},
			"biy":        {"friday"},
			"dis":        {"december"},
			"ene":        {"january"},
			"gmt":        {"gmt"},
			"hul":        {"july"},
			"hun":        {"june"},
			"huw":        {"thursday"},
			"lin":        {"sunday"},
			"lun":        {"monday"},
			"mar":        {"march"},
			"may":        {"may"},
			"miy":        {"wednesday"},
			"nob":        {"november"},
			"okt":        {"october"},
			"peb":        {"february"},
			"sab":        {"saturday"},
			"set":        {"september"},
			"utc":        {"utc"},
			"am":         {"am"},
			"na":         {""},
			"pm":         {"pm"},
			"sa":         {"in"},
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
			"z":          {"z"},
			"|":          {""},
		},
		KnownWords: []string{"miyerkules", "disyembre", "nakalipas", "nobyembre", "setyembre", "biyernes", "nakaraan", "huwebes", "oktubre", "pebrero", "segundo", "agosto", "linggo", "martes", "minuto", "sabado", "abril", "buwan", "enero", "ganap", "hulyo", "hunyo", "lunes", "marso", "noong", "araw", "mayo", "noon", "oras", "taon", "abr", "ago", "biy", "dis", "ene", "gmt", "hul", "hun", "huw", "lin", "lun", "mar", "may", "miy", "nob", "okt", "peb", "sab", "set", "utc", "am", "na", "pm", "sa", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
