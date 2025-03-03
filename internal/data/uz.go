// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	uz_Locale LocaleData
)

func init() {
	uz_Locale = merge(nil, LocaleData{
		Name:      "uz",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefghijklnorstuvyz‘`),
		Translations: map[string][]string{
			"chorshanba": {"wednesday"},
			"payshanba":  {"thursday"},
			"yakshanba":  {"sunday"},
			"dushanba":   {"monday"},
			"seshanba":   {"tuesday"},
			"sentabr":    {"september"},
			"avgust":     {"august"},
			"daqiqa":     {"minute"},
			"dekabr":     {"december"},
			"fevral":     {"february"},
			"noyabr":     {"november"},
			"oktabr":     {"october"},
			"shanba":     {"saturday"},
			"soniya":     {"second"},
			"yanvar":     {"january"},
			"aprel":      {"april"},
			"hafta":      {"week"},
			"chor":       {"wednesday"},
			"dush":       {"monday"},
			"iyul":       {"july"},
			"iyun":       {"june"},
			"juma":       {"friday"},
			"mart":       {"march"},
			"sesh":       {"tuesday"},
			"shan":       {"saturday"},
			"soat":       {"hour"},
			"apr":        {"april"},
			"avg":        {"august"},
			"daq":        {"minute"},
			"dek":        {"december"},
			"fev":        {"february"},
			"gmt":        {"gmt"},
			"iyl":        {"july"},
			"iyn":        {"june"},
			"jum":        {"friday"},
			"kun":        {"day"},
			"mar":        {"march"},
			"may":        {"may"},
			"noy":        {"november"},
			"okt":        {"october"},
			"pay":        {"thursday"},
			"sen":        {"september"},
			"son":        {"second"},
			"utc":        {"utc"},
			"yak":        {"sunday"},
			"yan":        {"january"},
			"yil":        {"year"},
			"am":         {"am"},
			"oy":         {"month"},
			"pm":         {"pm"},
			"tk":         {"pm"},
			"to":         {"am"},
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
		RelativeType: map[string]string{
			"keyingi hafta": "in 1 week",
			"o‘tgan hafta":  "1 week ago",
			"shu daqiqada":  "0 minute ago",
			"keyingi yil":   "in 1 year",
			"keyingi oy":    "in 1 month",
			"o'tgan yil":    "1 year ago",
			"o‘tgan yil":    "1 year ago",
			"shu soatda":    "0 hour ago",
			"o‘tgan oy":     "1 month ago",
			"shu hafta":     "0 week ago",
			"shu yil":       "0 year ago",
			"bu yil":        "0 year ago",
			"ertaga":        "in 1 day",
			"shu oy":        "0 month ago",
			"bugun":         "0 day ago",
			"hozir":         "0 second ago",
			"kecha":         "1 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) daqiqadan keyin`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) soniyadan keyin`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) haftadan keyin`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) soatdan keyin`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) daqiqa oldin`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) kundan keyin`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) soniya oldin`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) yildan keyin`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) hafta oldin`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) oydan keyin`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) soat oldin`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) kun oldin`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) yil oldin`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) oy oldin`), "$1 month ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* daqiqadan keyin|\d+[.,]?\d* soniyadan keyin|\d+[.,]?\d* haftadan keyin|\d+[.,]?\d* soatdan keyin|\d+[.,]?\d* daqiqa oldin|\d+[.,]?\d* kundan keyin|\d+[.,]?\d* soniya oldin|\d+[.,]?\d* yildan keyin|\d+[.,]?\d* hafta oldin|\d+[.,]?\d* oydan keyin|\d+[.,]?\d* soat oldin|\d+[.,]?\d* kun oldin|\d+[.,]?\d* yil oldin|\d+[.,]?\d* oy oldin)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* daqiqadan keyin|\d+[.,]?\d* soniyadan keyin|\d+[.,]?\d* haftadan keyin|\d+[.,]?\d* soatdan keyin|\d+[.,]?\d* daqiqa oldin|\d+[.,]?\d* kundan keyin|\d+[.,]?\d* soniya oldin|\d+[.,]?\d* yildan keyin|\d+[.,]?\d* hafta oldin|\d+[.,]?\d* oydan keyin|\d+[.,]?\d* soat oldin|\d+[.,]?\d* kun oldin|\d+[.,]?\d* yil oldin|\d+[.,]?\d* oy oldin)$`),
		KnownWords:      []string{"keyingi hafta", "o‘tgan hafta", "shu daqiqada", "keyingi yil", "chorshanba", "keyingi oy", "o'tgan yil", "o‘tgan yil", "shu soatda", "o‘tgan oy", "payshanba", "shu hafta", "yakshanba", "dushanba", "seshanba", "sentabr", "shu yil", "avgust", "bu yil", "daqiqa", "dekabr", "ertaga", "fevral", "noyabr", "oktabr", "shanba", "shu oy", "soniya", "yanvar", "aprel", "bugun", "hafta", "hozir", "kecha", "chor", "dush", "iyul", "iyun", "juma", "mart", "sesh", "shan", "soat", "apr", "avg", "daq", "dek", "fev", "gmt", "iyl", "iyn", "jum", "kun", "mar", "may", "noy", "okt", "pay", "sen", "son", "utc", "yak", "yan", "yil", "am", "oy", "pm", "tk", "to", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
