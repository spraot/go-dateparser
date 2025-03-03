// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	ml_Locale LocaleData
)

func init() {
	ml_Locale = merge(nil, LocaleData{
		Name:      "ml",
		DateOrder: "DMY",
		Charset:   []rune(`cgtuzംഅആഇഈഏഒഓകഗങചജഞടഡണതദധനപഫബമയരറലളഴവശഷസാിുൂെേൈൊോ്ൺൻർൽൾ‌`),
		Translations: map[string][]string{
			"വെളളിയാഴ‌ച": {"friday"},
			"തിങകളാഴ‌ച":  {"monday"},
			"ചൊവവാഴ‌ച":   {"tuesday"},
			"ഞായറാഴ‌ച":   {"sunday"},
			"വയാഴാഴ‌ച":   {"thursday"},
			"ശനിയാഴ‌ച":   {"saturday"},
			"സെപററംബർ":   {"september"},
			"ഒക‌ടോബർ":    {"october"},
			"ചൊവവാഴച":    {"tuesday"},
			"ഫെബരവരി":    {"february"},
			"ബധനാഴ‌ച":    {"wednesday"},
			"ഡിസംബർ":     {"december"},
			"മണികകർ":     {"hour"},
			"മിനിററ":     {"minute"},
			"സെകകൻഡ":     {"second"},
			"സെപററം":     {"september"},
			"ഏപരിൽ":      {"april"},
			"ഓഗസററ":      {"august"},
			"ജനവരി":      {"january"},
			"തിങകൾ":      {"monday"},
			"ദിവസം":      {"day"},
			"നവംബർ":      {"november"},
			"മാർചച":      {"march"},
			"വയാഴം":      {"thursday"},
			"വെളളി":      {"friday"},
			"ഏപരി":       {"april"},
			"ഒകടോ":       {"october"},
			"ചൊവവ":       {"tuesday"},
			"ഞായർ":       {"sunday"},
			"ഡിസം":       {"december"},
			"ഫെബര":       {"february"},
			"മാസം":       {"month"},
			"വർഷം":       {"year"},
			"gmt":        {"gmt"},
			"utc":        {"utc"},
			"ആഴച":        {"week"},
			"ജലൈ":        {"july"},
			"നവം":        {"november"},
			"ബധൻ":        {"wednesday"},
			"മാർ":        {"march"},
			"മേയ":        {"may"},
			"ശനി":        {"saturday"},
			"am":         {"am"},
			"pm":         {"pm"},
			"ഓഗ":         {"august"},
			"ജന":         {"january"},
			"ജൺ":         {"june"},
			"മാ":         {"month"},
			"മി":         {"minute"},
			"സെ":         {"second"},
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
			"ആ":          {"week"},
			"മ":          {"hour"},
			"വ":          {"year"},
		},
		RelativeType: map[string]string{
			"ഈ മണികകറിൽ": "0 hour ago",
			"ഈ മിനിററിൽ": "0 minute ago",
			"കഴിഞഞ ആഴ‌ച": "1 week ago",
			"കഴിഞഞ മാസം": "1 month ago",
			"കഴിഞഞ വർഷം": "1 year ago",
			"അടതത മാസം":  "in 1 month",
			"അടതത ആഴച":   "in 1 week",
			"അടതതവർഷം":   "in 1 year",
			"ഈ വർ‌ഷം":    "0 year ago",
			"ഈ മാസം":     "0 month ago",
			"ഇനനലെ":      "1 day ago",
			"ഇപപോൾ":      "0 second ago",
			"ഈ ആഴച":      "0 week ago",
			"നാളെ":       "in 1 day",
			"ഇനന":        "0 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) മണികകർ മമപ`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) മിനിററ മമപ`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) സെകകൻഡ മമപ`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ദിവസം മമപ`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ദിവസതതിൽ`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) മണികകറിൽ`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) മാസം മമപ`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) മിനിററിൽ`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) വർഷം മമപ`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) സെകകൻഡിൽ`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ആഴച മമപ`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) മാസതതിൽ`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) വർഷതതിൽ`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ആഴചയിൽ`), "in $1 week"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* മണികകർ മമപ|\d+[.,]?\d* മിനിററ മമപ|\d+[.,]?\d* സെകകൻഡ മമപ|\d+[.,]?\d* ദിവസം മമപ|\d+[.,]?\d* ദിവസതതിൽ|\d+[.,]?\d* മണികകറിൽ|\d+[.,]?\d* മാസം മമപ|\d+[.,]?\d* മിനിററിൽ|\d+[.,]?\d* വർഷം മമപ|\d+[.,]?\d* സെകകൻഡിൽ|\d+[.,]?\d* ആഴച മമപ|\d+[.,]?\d* മാസതതിൽ|\d+[.,]?\d* വർഷതതിൽ|\d+[.,]?\d* ആഴചയിൽ)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* മണികകർ മമപ|\d+[.,]?\d* മിനിററ മമപ|\d+[.,]?\d* സെകകൻഡ മമപ|\d+[.,]?\d* ദിവസം മമപ|\d+[.,]?\d* ദിവസതതിൽ|\d+[.,]?\d* മണികകറിൽ|\d+[.,]?\d* മാസം മമപ|\d+[.,]?\d* മിനിററിൽ|\d+[.,]?\d* വർഷം മമപ|\d+[.,]?\d* സെകകൻഡിൽ|\d+[.,]?\d* ആഴച മമപ|\d+[.,]?\d* മാസതതിൽ|\d+[.,]?\d* വർഷതതിൽ|\d+[.,]?\d* ആഴചയിൽ)$`),
		KnownWords:      []string{"ഈ മണികകറിൽ", "ഈ മിനിററിൽ", "കഴിഞഞ ആഴ‌ച", "കഴിഞഞ മാസം", "കഴിഞഞ വർഷം", "വെളളിയാഴ‌ച", "അടതത മാസം", "തിങകളാഴ‌ച", "അടതത ആഴച", "അടതതവർഷം", "ചൊവവാഴ‌ച", "ഞായറാഴ‌ച", "വയാഴാഴ‌ച", "ശനിയാഴ‌ച", "സെപററംബർ", "ഈ വർ‌ഷം", "ഒക‌ടോബർ", "ചൊവവാഴച", "ഫെബരവരി", "ബധനാഴ‌ച", "ഈ മാസം", "ഡിസംബർ", "മണികകർ", "മിനിററ", "സെകകൻഡ", "സെപററം", "ഇനനലെ", "ഇപപോൾ", "ഈ ആഴച", "ഏപരിൽ", "ഓഗസററ", "ജനവരി", "തിങകൾ", "ദിവസം", "നവംബർ", "മാർചച", "വയാഴം", "വെളളി", "ഏപരി", "ഒകടോ", "ചൊവവ", "ഞായർ", "ഡിസം", "നാളെ", "ഫെബര", "മാസം", "വർഷം", "gmt", "utc", "ആഴച", "ഇനന", "ജലൈ", "നവം", "ബധൻ", "മാർ", "മേയ", "ശനി", "am", "pm", "ഓഗ", "ജന", "ജൺ", "മാ", "മി", "സെ", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "ആ", "മ", "വ"},
	})
}
