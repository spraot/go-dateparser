// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	az_Locale LocaleData
)

func init() {
	az_Locale = merge(nil, LocaleData{
		Name:      "az",
		DateOrder: "DMY",
		Charset:   []rune(`bcdefghiklnorstuvxyzçüışə`),
		Translations: map[string][]string{
			"cərsənbə axsamı": {"tuesday"},
			"bazar ertəsi":    {"monday"},
			"cumə axsamı":     {"thursday"},
			"cərsənbə":        {"wednesday"},
			"sentyabr":        {"september"},
			"oktyabr":         {"october"},
			"avqust":          {"august"},
			"dekabr":          {"december"},
			"dəqiqə":          {"minute"},
			"fevral":          {"february"},
			"noyabr":          {"november"},
			"saniyə":          {"second"},
			"yanvar":          {"january"},
			"aprel":           {"april"},
			"bazar":           {"sunday"},
			"həftə":           {"week"},
			"sənbə":           {"saturday"},
			"cumə":            {"friday"},
			"iyul":            {"july"},
			"iyun":            {"june"},
			"mart":            {"march"},
			"saat":            {"hour"},
			"apr":             {"april"},
			"avq":             {"august"},
			"dek":             {"december"},
			"dəq":             {"minute"},
			"fev":             {"february"},
			"gmt":             {"gmt"},
			"gun":             {"day"},
			"iyl":             {"july"},
			"iyn":             {"june"},
			"mar":             {"march"},
			"may":             {"may"},
			"noy":             {"november"},
			"okt":             {"october"},
			"san":             {"second"},
			"sen":             {"september"},
			"utc":             {"utc"},
			"yan":             {"january"},
			"am":              {"am"},
			"ay":              {"month"},
			"be":              {"monday"},
			"ca":              {"tuesday", "thursday"},
			"il":              {"year"},
			"pm":              {"pm"},
			" ":               {" "},
			"'":               {""},
			"+":               {"+"},
			",":               {""},
			"-":               {"-"},
			".":               {"."},
			"/":               {"/"},
			":":               {":"},
			";":               {""},
			"@":               {""},
			"[":               {""},
			"]":               {""},
			"b":               {"sunday"},
			"c":               {"friday", "wednesday"},
			"s":               {"saturday"},
			"z":               {"z"},
			"|":               {""},
		},
		RelativeType: map[string]string{
			"gələn həftə": "in 1 week",
			"kecən həftə": "1 week ago",
			"bu dəqiqə":   "0 minute ago",
			"bu həftə":    "0 week ago",
			"gələn ay":    "in 1 month",
			"gələn il":    "in 1 year",
			"kecən ay":    "1 month ago",
			"kecən il":    "1 year ago",
			"bu saat":     "0 hour ago",
			"bu gun":      "0 day ago",
			"bu ay":       "0 month ago",
			"bu il":       "0 year ago",
			"dunən":       "1 day ago",
			"sabah":       "in 1 day",
			"indi":        "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dəqiqə ərzində`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) saniyə ərzində`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) həftə ərzində`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) saat ərzində`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dəqiqə oncə`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) gun ərzində`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) saniyə oncə`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ay ərzində`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) həftə oncə`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) il ərzində`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) saat oncə`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) gun oncə`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ay oncə`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) il oncə`), "$1 year ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* dəqiqə ərzində|\d+[.,]?\d* saniyə ərzində|\d+[.,]?\d* həftə ərzində|\d+[.,]?\d* saat ərzində|\d+[.,]?\d* dəqiqə oncə|\d+[.,]?\d* gun ərzində|\d+[.,]?\d* saniyə oncə|\d+[.,]?\d* ay ərzində|\d+[.,]?\d* həftə oncə|\d+[.,]?\d* il ərzində|\d+[.,]?\d* saat oncə|\d+[.,]?\d* gun oncə|\d+[.,]?\d* ay oncə|\d+[.,]?\d* il oncə)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* dəqiqə ərzində|\d+[.,]?\d* saniyə ərzində|\d+[.,]?\d* həftə ərzində|\d+[.,]?\d* saat ərzində|\d+[.,]?\d* dəqiqə oncə|\d+[.,]?\d* gun ərzində|\d+[.,]?\d* saniyə oncə|\d+[.,]?\d* ay ərzində|\d+[.,]?\d* həftə oncə|\d+[.,]?\d* il ərzində|\d+[.,]?\d* saat oncə|\d+[.,]?\d* gun oncə|\d+[.,]?\d* ay oncə|\d+[.,]?\d* il oncə)$`),
		KnownWords:      []string{"cərsənbə axsamı", "bazar ertəsi", "cumə axsamı", "gələn həftə", "kecən həftə", "bu dəqiqə", "bu həftə", "cərsənbə", "gələn ay", "gələn il", "kecən ay", "kecən il", "sentyabr", "bu saat", "oktyabr", "avqust", "bu gun", "dekabr", "dəqiqə", "fevral", "noyabr", "saniyə", "yanvar", "aprel", "bazar", "bu ay", "bu il", "dunən", "həftə", "sabah", "sənbə", "cumə", "indi", "iyul", "iyun", "mart", "saat", "apr", "avq", "dek", "dəq", "fev", "gmt", "gun", "iyl", "iyn", "mar", "may", "noy", "okt", "san", "sen", "utc", "yan", "am", "ay", "be", "ca", "il", "pm", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "b", "c", "s", "z", "|"},
	})
}
