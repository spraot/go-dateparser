// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var fur_Locale = merge(nil, LocaleData{
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
		{regexp.MustCompile(`(?i)(\d+) setemanis indaur`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) setemane indaur`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) zornadis indaur`), "$1 day ago"},
		{regexp.MustCompile(`(?i)ca di (\d+) setemanis`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) seconts indaur`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) zornade indaur`), "$1 day ago"},
		{regexp.MustCompile(`(?i)ca di (\d+) setemane`), "in $1 week"},
		{regexp.MustCompile(`(?i)ca di (\d+) zornadis`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) minuts indaur`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) secont indaur`), "$1 second ago"},
		{regexp.MustCompile(`(?i)ca di (\d+) seconts`), "in $1 second"},
		{regexp.MustCompile(`(?i)ca di (\d+) zornade`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) minut indaur`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)ca di (\d+) minuts`), "in $1 minute"},
		{regexp.MustCompile(`(?i)ca di (\d+) secont`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) agns indaur`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) oris indaur`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)ca di (\d+) minut`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) mes indaur`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) ore indaur`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)ca di (\d+) agns`), "in $1 year"},
		{regexp.MustCompile(`(?i)ca di (\d+) oris`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) an indaur`), "$1 year ago"},
		{regexp.MustCompile(`(?i)ca di (\d+) mes`), "in $1 month"},
		{regexp.MustCompile(`(?i)ca di (\d+) ore`), "in $1 hour"},
		{regexp.MustCompile(`(?i)ca di (\d+) an`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ setemanis indaur|\d+ setemane indaur|\d+ zornadis indaur|ca di \d+ setemanis|\d+ seconts indaur|\d+ zornade indaur|ca di \d+ setemane|ca di \d+ zornadis|\d+ minuts indaur|\d+ secont indaur|ca di \d+ seconts|ca di \d+ zornade|\d+ minut indaur|ca di \d+ minuts|ca di \d+ secont|\d+ agns indaur|\d+ oris indaur|ca di \d+ minut|\d+ mes indaur|\d+ ore indaur|ca di \d+ agns|ca di \d+ oris|\d+ an indaur|ca di \d+ mes|ca di \d+ ore|ca di \d+ an)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ setemanis indaur|\d+ setemane indaur|\d+ zornadis indaur|ca di \d+ setemanis|\d+ seconts indaur|\d+ zornade indaur|ca di \d+ setemane|ca di \d+ zornadis|\d+ minuts indaur|\d+ secont indaur|ca di \d+ seconts|ca di \d+ zornade|\d+ minut indaur|ca di \d+ minuts|ca di \d+ secont|\d+ agns indaur|\d+ oris indaur|ca di \d+ minut|\d+ mes indaur|\d+ ore indaur|ca di \d+ agns|ca di \d+ oris|\d+ an indaur|ca di \d+ mes|ca di \d+ ore|ca di \d+ an)$`),
	KnownWords:      []string{"this minute", "last month", "next month", "this month", "last week", "last year", "next week", "next year", "this hour", "this week", "this year", "dicembar", "novembar", "setemane", "setembar", "domenie", "martars", "miercus", "fevrar", "otubar", "sabide", "secont", "vinars", "avost", "avril", "doman", "joibe", "lunis", "minut", "zenar", "jugn", "marc", "avo", "avr", "dic", "dom", "fev", "gmt", "joi", "jug", "lui", "lun", "mai", "mar", "mes", "mie", "nov", "now", "ore", "otu", "sab", "set", "utc", "vin", "vue", "zen", "am", "an", "di", "ir", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "a", "p", "z", "|"},
})
