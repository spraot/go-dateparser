// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	fur_Locale LocaleData
)

func init() {
	fur_Locale = merge(nil, LocaleData{
		Name:      "fur",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefghijklnorstuvwxyzâçêìîû`),
		Translations: map[string][]string{
			"dicembar": {"december"},
			"novembar": {"november"},
			"setemane": {"week"},
			"setembar": {"september"},
			"domenie":  {"sunday"},
			"martars":  {"tuesday"},
			"miercus":  {"wednesday"},
			"fevrar":   {"february"},
			"otubar":   {"october"},
			"sabide":   {"saturday"},
			"secont":   {"second"},
			"vinars":   {"friday"},
			"avost":    {"august"},
			"avril":    {"april"},
			"joibe":    {"thursday"},
			"lunis":    {"monday"},
			"minut":    {"minute"},
			"zenar":    {"january"},
			"jugn":     {"june"},
			"marc":     {"march"},
			"avo":      {"august"},
			"avr":      {"april"},
			"dic":      {"december"},
			"dom":      {"sunday"},
			"fev":      {"february"},
			"gmt":      {"gmt"},
			"joi":      {"thursday"},
			"jug":      {"june"},
			"lui":      {"july"},
			"lun":      {"monday"},
			"mai":      {"may"},
			"mar":      {"march", "tuesday"},
			"mes":      {"month"},
			"mie":      {"wednesday"},
			"nov":      {"november"},
			"ore":      {"hour"},
			"otu":      {"october"},
			"sab":      {"saturday"},
			"set":      {"september"},
			"utc":      {"utc"},
			"vin":      {"friday"},
			"zen":      {"january"},
			"am":       {"am"},
			"an":       {"year"},
			"di":       {"day"},
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
			"a":        {"am"},
			"p":        {"pm"},
			"z":        {"z"},
			"|":        {""},
		},
		RelativeType: map[string]string{
			"this minute": "0 minute ago",
			"last month":  "1 month ago",
			"next month":  "in 1 month",
			"this month":  "0 month ago",
			"last week":   "1 week ago",
			"last year":   "1 year ago",
			"next week":   "in 1 week",
			"next year":   "in 1 year",
			"this hour":   "0 hour ago",
			"this week":   "0 week ago",
			"this year":   "0 year ago",
			"doman":       "in 1 day",
			"now":         "0 second ago",
			"vue":         "0 day ago",
			"ir":          "1 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) setemanis indaur`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) setemane indaur`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) zornadis indaur`), "$1 day ago"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) setemanis`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) seconts indaur`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) zornade indaur`), "$1 day ago"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) setemane`), "in $1 week"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) zornadis`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minuts indaur`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) secont indaur`), "$1 second ago"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) seconts`), "in $1 second"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) zornade`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minut indaur`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) minuts`), "in $1 minute"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) secont`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) agns indaur`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) oris indaur`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) minut`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) mes indaur`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ore indaur`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) agns`), "in $1 year"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) oris`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) an indaur`), "$1 year ago"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) mes`), "in $1 month"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) ore`), "in $1 hour"},
			{regexp.MustCompile(`(?i)ca di (\d+[.,]?\d*) an`), "in $1 year"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* setemanis indaur|\d+[.,]?\d* setemane indaur|\d+[.,]?\d* zornadis indaur|ca di \d+[.,]?\d* setemanis|\d+[.,]?\d* seconts indaur|\d+[.,]?\d* zornade indaur|ca di \d+[.,]?\d* setemane|ca di \d+[.,]?\d* zornadis|\d+[.,]?\d* minuts indaur|\d+[.,]?\d* secont indaur|ca di \d+[.,]?\d* seconts|ca di \d+[.,]?\d* zornade|\d+[.,]?\d* minut indaur|ca di \d+[.,]?\d* minuts|ca di \d+[.,]?\d* secont|\d+[.,]?\d* agns indaur|\d+[.,]?\d* oris indaur|ca di \d+[.,]?\d* minut|\d+[.,]?\d* mes indaur|\d+[.,]?\d* ore indaur|ca di \d+[.,]?\d* agns|ca di \d+[.,]?\d* oris|\d+[.,]?\d* an indaur|ca di \d+[.,]?\d* mes|ca di \d+[.,]?\d* ore|ca di \d+[.,]?\d* an)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* setemanis indaur|\d+[.,]?\d* setemane indaur|\d+[.,]?\d* zornadis indaur|ca di \d+[.,]?\d* setemanis|\d+[.,]?\d* seconts indaur|\d+[.,]?\d* zornade indaur|ca di \d+[.,]?\d* setemane|ca di \d+[.,]?\d* zornadis|\d+[.,]?\d* minuts indaur|\d+[.,]?\d* secont indaur|ca di \d+[.,]?\d* seconts|ca di \d+[.,]?\d* zornade|\d+[.,]?\d* minut indaur|ca di \d+[.,]?\d* minuts|ca di \d+[.,]?\d* secont|\d+[.,]?\d* agns indaur|\d+[.,]?\d* oris indaur|ca di \d+[.,]?\d* minut|\d+[.,]?\d* mes indaur|\d+[.,]?\d* ore indaur|ca di \d+[.,]?\d* agns|ca di \d+[.,]?\d* oris|\d+[.,]?\d* an indaur|ca di \d+[.,]?\d* mes|ca di \d+[.,]?\d* ore|ca di \d+[.,]?\d* an)$`),
		KnownWords:      []string{"this minute", "last month", "next month", "this month", "last week", "last year", "next week", "next year", "this hour", "this week", "this year", "dicembar", "novembar", "setemane", "setembar", "domenie", "martars", "miercus", "fevrar", "otubar", "sabide", "secont", "vinars", "avost", "avril", "doman", "joibe", "lunis", "minut", "zenar", "jugn", "marc", "avo", "avr", "dic", "dom", "fev", "gmt", "joi", "jug", "lui", "lun", "mai", "mar", "mes", "mie", "nov", "now", "ore", "otu", "sab", "set", "utc", "vin", "vue", "zen", "am", "an", "di", "ir", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "p", "z", "|"},
	})
}
