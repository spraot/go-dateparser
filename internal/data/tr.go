// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	tr_Locale    LocaleData
	tr_CY_Locale LocaleData
)

func init() {
	tr_Locale = merge(nil, LocaleData{
		Name:                  "tr",
		DateOrder:             "DMY",
		Charset:               []rune(`bcdefghiklnorstuvyzçöüğış`),
		SentenceSplitterGroup: 1,
		Translations: map[string][]string{
			"icerisinde": {"in"},
			"cumartesi":  {"saturday"},
			"pazartesi":  {"monday"},
			"carsamba":   {"wednesday"},
			"persembe":   {"thursday"},
			"yaklasık":   {""},
			"agustos":    {"august"},
			"haziran":    {"june"},
			"aralık":     {"december"},
			"dakika":     {"minute"},
			"icinde":     {"in"},
			"saniye":     {"second"},
			"temmuz":     {"july"},
			"eylul":      {"september"},
			"hafta":      {"week"},
			"kasım":      {"november"},
			"mayıs":      {"may"},
			"nisan":      {"april"},
			"pazar":      {"sunday"},
			"sonra":      {"in"},
			"subat":      {"february"},
			"cuma":       {"friday"},
			"ekim":       {"october"},
			"mart":       {"march"},
			"ocak":       {"january"},
			"once":       {"ago"},
			"saat":       {"hour"},
			"salı":       {"tuesday"},
			"sene":       {"year"},
			"agu":        {"august"},
			"ara":        {"december"},
			"car":        {"wednesday"},
			"cmt":        {"saturday"},
			"crs":        {"wednesday"},
			"cum":        {"friday"},
			"eki":        {"october"},
			"eyl":        {"september"},
			"gmt":        {"gmt"},
			"gun":        {"day"},
			"haz":        {"june"},
			"kas":        {"november"},
			"mar":        {"march"},
			"may":        {"may"},
			"nis":        {"april"},
			"oca":        {"january"},
			"paz":        {"sunday"},
			"per":        {"thursday"},
			"prs":        {"thursday"},
			"pzt":        {"monday"},
			"sal":        {"tuesday"},
			"sub":        {"february"},
			"tem":        {"july"},
			"utc":        {"utc"},
			"yıl":        {"year"},
			"ag":         {"august"},
			"am":         {"am"},
			"ar":         {"december"},
			"ay":         {"month"},
			"dk":         {"minute"},
			"ek":         {"october"},
			"ey":         {"september"},
			"ha":         {"june"},
			"hf":         {"week"},
			"ka":         {"november"},
			"ni":         {"april"},
			"oc":         {"january"},
			"oo":         {"am"},
			"os":         {"pm"},
			"pm":         {"pm"},
			"sa":         {"hour"},
			"sn":         {"second"},
			"su":         {"february"},
			"te":         {"july"},
			"ve":         {""},
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
			"onumuzdeki hafta": "in 1 week",
			"onumuzdeki gun":   "in 1 day",
			"onumuzdeki yıl":   "in 1 year",
			"gelecek hafta":    "in 1 week",
			"onumuzdeki ay":    "in 1 month",
			"gecen hafta":      "1 week ago",
			"gelecek yıl":      "in 1 year",
			"gelecek ay":       "in 1 month",
			"bu dakika":        "0 minute ago",
			"gecen gun":        "1 day ago",
			"gecen yıl":        "1 year ago",
			"bu hafta":         "0 week ago",
			"gecen ay":         "1 month ago",
			"bu saat":          "0 hour ago",
			"haftaya":          "in 1 week",
			"bu yıl":           "0 year ago",
			"bu ay":            "0 month ago",
			"bugun":            "0 day ago",
			"simdi":            "0 second ago",
			"yarın":            "in 1 day",
			"dun":              "1 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dakika sonra`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) saniye sonra`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dakika once`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) hafta sonra`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) saniye once`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) hafta once`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) saat sonra`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) gun sonra`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) saat once`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) yıl sonra`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ay sonra`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dk sonra`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) gun once`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) hf sonra`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sa sonra`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sn sonra`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) yıl once`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ay once`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) dk once`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) hf once`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sa once`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sn once`), "$1 second ago"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* dakika sonra|\d+[.,]?\d* saniye sonra|\d+[.,]?\d* dakika once|\d+[.,]?\d* hafta sonra|\d+[.,]?\d* saniye once|\d+[.,]?\d* hafta once|\d+[.,]?\d* saat sonra|\d+[.,]?\d* gun sonra|\d+[.,]?\d* saat once|\d+[.,]?\d* yıl sonra|\d+[.,]?\d* ay sonra|\d+[.,]?\d* dk sonra|\d+[.,]?\d* gun once|\d+[.,]?\d* hf sonra|\d+[.,]?\d* sa sonra|\d+[.,]?\d* sn sonra|\d+[.,]?\d* yıl once|\d+[.,]?\d* ay once|\d+[.,]?\d* dk once|\d+[.,]?\d* hf once|\d+[.,]?\d* sa once|\d+[.,]?\d* sn once)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* dakika sonra|\d+[.,]?\d* saniye sonra|\d+[.,]?\d* dakika once|\d+[.,]?\d* hafta sonra|\d+[.,]?\d* saniye once|\d+[.,]?\d* hafta once|\d+[.,]?\d* saat sonra|\d+[.,]?\d* gun sonra|\d+[.,]?\d* saat once|\d+[.,]?\d* yıl sonra|\d+[.,]?\d* ay sonra|\d+[.,]?\d* dk sonra|\d+[.,]?\d* gun once|\d+[.,]?\d* hf sonra|\d+[.,]?\d* sa sonra|\d+[.,]?\d* sn sonra|\d+[.,]?\d* yıl once|\d+[.,]?\d* ay once|\d+[.,]?\d* dk once|\d+[.,]?\d* hf once|\d+[.,]?\d* sa once|\d+[.,]?\d* sn once)$`),
		KnownWords:      []string{"onumuzdeki hafta", "onumuzdeki gun", "onumuzdeki yıl", "gelecek hafta", "onumuzdeki ay", "gecen hafta", "gelecek yıl", "gelecek ay", "icerisinde", "bu dakika", "cumartesi", "gecen gun", "gecen yıl", "pazartesi", "bu hafta", "carsamba", "gecen ay", "persembe", "yaklasık", "agustos", "bu saat", "haftaya", "haziran", "aralık", "bu yıl", "dakika", "icinde", "saniye", "temmuz", "bu ay", "bugun", "eylul", "hafta", "kasım", "mayıs", "nisan", "pazar", "simdi", "sonra", "subat", "yarın", "cuma", "ekim", "mart", "ocak", "once", "saat", "salı", "sene", "agu", "ara", "car", "cmt", "crs", "cum", "dun", "eki", "eyl", "gmt", "gun", "haz", "kas", "mar", "may", "nis", "oca", "paz", "per", "prs", "pzt", "sal", "sub", "tem", "utc", "yıl", "ag", "am", "ar", "ay", "dk", "ek", "ey", "ha", "hf", "ka", "ni", "oc", "oo", "os", "pm", "sa", "sn", "su", "te", "ve", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})

	tr_CY_Locale = merge(&tr_Locale, LocaleData{
		Name:      "tr-CY",
		DateOrder: "DMY",
	})
}
