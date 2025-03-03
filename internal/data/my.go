// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	my_Locale LocaleData
)

func init() {
	my_Locale = merge(nil, LocaleData{
		Name:      "my",
		DateOrder: "DMY",
		Charset:   []rune(`cgtuzကခဂငစဇညတဒဓနပဖဗဘမယရလဝသဟအဤဧဩါာိီုူေဲံ့း္်ျြွှ`),
		Translations: map[string][]string{
			"ကြာသပတေး": {"thursday"},
			"ဖေဖောဝါရ": {"february"},
			"အောကတဘာ":  {"october"},
			"စကတငဘာ":   {"september"},
			"ဇနနဝါရ":   {"january"},
			"တနငဂနေ":   {"sunday"},
			"သောကြာ":   {"friday"},
			"တနငလာ":    {"monday"},
			"ဒဇငဘာ":    {"december"},
			"နဝငဘာ":    {"november"},
			"ဗဒဓဟး":    {"wednesday"},
			"စကကန":     {"second"},
			"အငဂါ":     {"tuesday"},
			"အောက":     {"october"},
			"gmt":      {"gmt"},
			"utc":      {"utc"},
			"စနေ":      {"saturday"},
			"ဇလင":      {"july"},
			"ညနေ":      {"pm"},
			"နနက":      {"am"},
			"နာရ":      {"hour"},
			"မနစ":      {"minute"},
			"ဧပြ":      {"april"},
			"ဩဂတ":      {"august"},
			"am":       {"am"},
			"pm":       {"pm"},
			"စက":       {"september"},
			"ဇန":       {"june", "january"},
			"နစ":       {"year"},
			"ပတ":       {"week"},
			"ဖေ":       {"february"},
			"မတ":       {"march"},
			"မေ":       {"may"},
			"ရက":       {"day"},
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
			"z":        {"z"},
			"|":        {""},
			"ဇ":        {"july"},
			"ဒ":        {"december"},
			"န":        {"november"},
			"လ":        {"month"},
			"ဧ":        {"april"},
			"ဩ":        {"august"},
		},
		RelativeType: map[string]string{
			"ပြးခသည သတငးပတ": "1 week ago",
			"လာမည သတငးပတ":   "in 1 week",
			"ယခ သတငးပတ":     "0 week ago",
			"ပြးခသညလ":       "1 month ago",
			"မနကဖြန":        "in 1 day",
			"လာမညနစ":        "in 1 year",
			"ယမနနစ":         "1 year ago",
			"လာမညလ":         "in 1 month",
			"ဤအချန":         "0 hour ago",
			"မနေက":          "1 day ago",
			"ယခနစ":          "0 year ago",
			"ဤမနစ":          "0 minute ago",
			"ယခလ":           "0 month ago",
			"ယနေ":           "0 day ago",
			"ယခ":            "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)ပြးခသည (\d+[.,]?\d*) စကကန`), "$1 second ago"},
			{regexp.MustCompile(`(?i)ပြးခသည (\d+[.,]?\d*) နာရ`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)ပြးခသည (\d+[.,]?\d*) မနစ`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)ပြးခသည (\d+[.,]?\d*) နစ`), "$1 year ago"},
			{regexp.MustCompile(`(?i)ပြးခသည (\d+[.,]?\d*) ပတ`), "$1 week ago"},
			{regexp.MustCompile(`(?i)ပြးခသည (\d+[.,]?\d*) ရက`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) စကကနအတငး`), "in $1 second"},
			{regexp.MustCompile(`(?i)ပြးခသည (\d+[.,]?\d*) လ`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) နာရအတငး`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) မနစအတငး`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) နစအတငး`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ပတအတငး`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ရကအတငး`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) လအတငး`), "in $1 month"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(ပြးခသည \d+[.,]?\d* စကကန|ပြးခသည \d+[.,]?\d* နာရ|ပြးခသည \d+[.,]?\d* မနစ|ပြးခသည \d+[.,]?\d* နစ|ပြးခသည \d+[.,]?\d* ပတ|ပြးခသည \d+[.,]?\d* ရက|\d+[.,]?\d* စကကနအတငး|ပြးခသည \d+[.,]?\d* လ|\d+[.,]?\d* နာရအတငး|\d+[.,]?\d* မနစအတငး|\d+[.,]?\d* နစအတငး|\d+[.,]?\d* ပတအတငး|\d+[.,]?\d* ရကအတငး|\d+[.,]?\d* လအတငး)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(ပြးခသည \d+[.,]?\d* စကကန|ပြးခသည \d+[.,]?\d* နာရ|ပြးခသည \d+[.,]?\d* မနစ|ပြးခသည \d+[.,]?\d* နစ|ပြးခသည \d+[.,]?\d* ပတ|ပြးခသည \d+[.,]?\d* ရက|\d+[.,]?\d* စကကနအတငး|ပြးခသည \d+[.,]?\d* လ|\d+[.,]?\d* နာရအတငး|\d+[.,]?\d* မနစအတငး|\d+[.,]?\d* နစအတငး|\d+[.,]?\d* ပတအတငး|\d+[.,]?\d* ရကအတငး|\d+[.,]?\d* လအတငး)$`),
		KnownWords:      []string{"ပြးခသည သတငးပတ", "လာမည သတငးပတ", "ယခ သတငးပတ", "ကြာသပတေး", "ဖေဖောဝါရ", "ပြးခသညလ", "အောကတဘာ", "စကတငဘာ", "ဇနနဝါရ", "တနငဂနေ", "မနကဖြန", "လာမညနစ", "သောကြာ", "တနငလာ", "ဒဇငဘာ", "နဝငဘာ", "ဗဒဓဟး", "ယမနနစ", "လာမညလ", "ဤအချန", "စကကန", "မနေက", "ယခနစ", "အငဂါ", "အောက", "ဤမနစ", "gmt", "utc", "စနေ", "ဇလင", "ညနေ", "နနက", "နာရ", "မနစ", "ယခလ", "ယနေ", "ဧပြ", "ဩဂတ", "am", "pm", "စက", "ဇန", "နစ", "ပတ", "ဖေ", "မတ", "မေ", "ယခ", "ရက", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "ဇ", "ဒ", "န", "လ", "ဧ", "ဩ"},
	})
}
