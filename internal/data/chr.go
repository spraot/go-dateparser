// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	chr_Locale LocaleData
)

func init() {
	chr_Locale = merge(nil, LocaleData{
		Name:      "chr",
		DateOrder: "MDY",
		Charset:   []rune(`cgtuzᏸᏹᏼꭰꭱꭲꭴꭵꭶꭷꭸꭹꭺꭻꭼꭽꭿꮂꮄꮅꮆꮇꮈꮎꮑꮒꮓꮕꮖꮘꮙꮚꮜꮝꮞꮠꮡꮢꮣꮤꮥꮧꮨꮩꮪꮫꮯꮲꮵꮶꮷꮼꮿ`),
		Translations: map[string][]string{
			"ꭴꮥꮨᏼꮜꮧꮢꭲ": {"year"},
			"ꭲꮿꮤꮼꮝꮤꮕ":  {"minute"},
			"ꭴꮎꮩꮣꮖꮝꭼ":  {"sunday"},
			"ꭴꮎꮩꮣꮘꮥꮎ":  {"saturday"},
			"ꭴꮎꮩꮣꮙꮕꭿ":  {"monday"},
			"ꮢꮎꮩꮣꮖꮝꮧ":  {"week"},
			"ꮢꭿᏹꭲꮧꮲ":   {"pm"},
			"ꮷꮎꭹꮆꮝꮧ":   {"friday"},
			"ꭰꮒꮝꭼꮨ":    {"may"},
			"ꭴꮓꮈꮤꮕ":    {"january"},
			"ꮕꭹꮑꭲꭶ":    {"thursday"},
			"ꮤꮅꮑꭲꭶ":    {"tuesday"},
			"ꮶꭲꮑꭲꭶ":    {"wednesday"},
			"ꭵꮝꭹᏹ":     {"december"},
			"ꭻᏸꮙꮒ":     {"july"},
			"ꮕꮣꮥꮖ":     {"november"},
			"ꮡꮯꮆꮣ":     {"hour"},
			"ꮥꭽꮇᏹ":     {"june"},
			"ꮪꮅꮝꮧ":     {"september"},
			"ꮪꮒꮕꮧ":     {"october"},
			"gmt":      {"gmt"},
			"utc":      {"utc"},
			"ꭰꮕᏹ":      {"march"},
			"ꭰꮞꮲ":      {"second"},
			"ꭲꮿꮤ":      {"minute"},
			"ꭶꮆꮒ":      {"august"},
			"ꭷꭶꮅ":      {"february"},
			"ꭷꮈꭲ":      {"month"},
			"ꭷꮼꮒ":      {"april"},
			"ꮕꭹꮑ":      {"thursday"},
			"ꮖꮝꭼ":      {"sunday"},
			"ꮘꮥꮎ":      {"saturday"},
			"ꮙꮕꭿ":      {"monday"},
			"ꮜꮎꮄ":      {"am"},
			"ꮤꮅꮑ":      {"tuesday"},
			"ꮶꭲꮑ":      {"wednesday"},
			"ꮷꮎꭹ":      {"friday"},
			"am":       {"am"},
			"pm":       {"pm"},
			"ꭰꮒ":       {"may"},
			"ꭰꮕ":       {"march"},
			"ꭲꭶ":       {"day"},
			"ꭴꮓ":       {"january"},
			"ꭴꮥ":       {"year"},
			"ꭵꮝ":       {"december"},
			"ꭶꮆ":       {"august"},
			"ꭷꭶ":       {"february"},
			"ꭷꮈ":       {"month"},
			"ꭷꮼ":       {"april"},
			"ꭻᏸ":       {"july"},
			"ꮕꮣ":       {"november"},
			"ꮡꮯ":       {"hour"},
			"ꮢꮎ":       {"week"},
			"ꮥꭽ":       {"june"},
			"ꮪꮅ":       {"september"},
			"ꮪꮒ":       {"october"},
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
		},
		RelativeType: map[string]string{
			"ꭿꭰ ꭲꮿꮤꮼꮝꮤꮕ": "0 minute ago",
			"ꭿꭰ ꮷꮥꮨᏼꮢꮨ":  "0 year ago",
			"ꭷꮈꭲ ꮵꭸꮢ":    "1 month ago",
			"ꭿꭰ ꭰꮅꮅꮜ":    "0 week ago",
			"ꭿꭰ ꮡꮯꮆꮣ":    "0 hour ago",
			"ꮤꮅꮑ ꭷꮈꭲ":    "in 1 month",
			"ꮵꮫꮅᏹꮅꮢꭲ":    "1 week ago",
			"ꭱꮨ ꮵꭸꮢ":     "1 year ago",
			"ꭿꭰ ꭷꮈꭲ":     "0 month ago",
			"ꭺꭿ ꭲꭶ":      "0 day ago",
			"ꮠꮖꮄꮕꮂ":      "in 1 week",
			"ꭱꮨᏼꭲ":       "in 1 year",
			"ꮜꮎꮄꭲ":       "in 1 day",
			"ꮓꮚ":         "0 second ago",
			"ꮢꭿ":         "1 day ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭲᏻꮎꮩꮣꮖꮝꮧ ꮵꭸꮢ`), "$1 week ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭲꮿꮤꮼꮝꮤꮕ ꮵꭸꮢ`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꮢꮎꮩꮣꮖꮝꮧ ꮵꭸꮢ`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ꭲꮷꮥꮨᏼꮜꮧꮢꭲ ꮵꭸꮢ`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ꭿꮈꮝꭹ ꮷꮢꭿꮫ ꮵꭸꮢ`), "$1 day ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꮣꮣꮎꮹꮝꭼ ꮵꭸꮢ`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ꭴꮥꮨᏼꮜꮧꮢꭲ ꮵꭸꮢ`), "$1 year ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭲꮷꮥꮨᏼꮜꮧꮢꭲ`), "in $1 year"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭿꮈꮝꭹ ꮷꮢꭿꮫ`), "in $1 day"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭲᏻꮎꮩꮣꮖꮝꮧ`), "in $1 week"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭴꮥꮨᏼꮜꮧꮢꭲ`), "in $1 year"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꮧꭷꮈꭲ ꮵꭸꮢ`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ꮣꮣꮎꮹꮝꭼ ꮵꭸꮢ`), "$1 second ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭲꮿꮤ ꮵꭸꮢ`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭲꮿꮤꮼꮝꮤꮕ`), "in $1 minute"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭷꮈꭲ ꮵꭸꮢ`), "$1 month ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꮢꮎꮩꮣꮖꮝꮧ`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ꭲᏻꮯꮆꮣ ꮵꭸꮢ`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭴꮥ ꮵꭸꮢ`), "$1 year ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭷꮈ ꮵꭸꮢ`), "$1 month ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꮡꮯ ꮵꭸꮢ`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꮢꮎ ꮵꭸꮢ`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ꮡꮯꮆꮣ ꮵꭸꮢ`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭲᏻꮯꮆꮣ`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ꭰꮞꮲ ꮵꭸꮢ`), "$1 second ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꮡꮯꮆꮣ`), "in $1 hour"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꮧꭷꮈꭲ`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) ꭲꭶ ꮵꭸꮢ`), "$1 day ago"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭰꮞꮲ`), "in $1 second"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭲꮿꮤ`), "in $1 minute"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭷꮈꭲ`), "in $1 month"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭲꭶ`), "in $1 day"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭴꮥ`), "in $1 year"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꭷꮈ`), "in $1 month"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꮡꮯ`), "in $1 hour"},
			{regexp.MustCompile(`(?i)ꮎꮏ (\d+[.,]?\d*) ꮢꮎ`), "in $1 week"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(ꮎꮏ \d+[.,]?\d* ꭲᏻꮎꮩꮣꮖꮝꮧ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭲꮿꮤꮼꮝꮤꮕ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮢꮎꮩꮣꮖꮝꮧ ꮵꭸꮢ|\d+[.,]?\d* ꭲꮷꮥꮨᏼꮜꮧꮢꭲ ꮵꭸꮢ|\d+[.,]?\d* ꭿꮈꮝꭹ ꮷꮢꭿꮫ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮣꮣꮎꮹꮝꭼ ꮵꭸꮢ|\d+[.,]?\d* ꭴꮥꮨᏼꮜꮧꮢꭲ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭲꮷꮥꮨᏼꮜꮧꮢꭲ|ꮎꮏ \d+[.,]?\d* ꭿꮈꮝꭹ ꮷꮢꭿꮫ|ꮎꮏ \d+[.,]?\d* ꭲᏻꮎꮩꮣꮖꮝꮧ|ꮎꮏ \d+[.,]?\d* ꭴꮥꮨᏼꮜꮧꮢꭲ|ꮎꮏ \d+[.,]?\d* ꮧꭷꮈꭲ ꮵꭸꮢ|\d+[.,]?\d* ꮣꮣꮎꮹꮝꭼ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭲꮿꮤ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭲꮿꮤꮼꮝꮤꮕ|ꮎꮏ \d+[.,]?\d* ꭷꮈꭲ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮢꮎꮩꮣꮖꮝꮧ|\d+[.,]?\d* ꭲᏻꮯꮆꮣ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭴꮥ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭷꮈ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮡꮯ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮢꮎ ꮵꭸꮢ|\d+[.,]?\d* ꮡꮯꮆꮣ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭲᏻꮯꮆꮣ|\d+[.,]?\d* ꭰꮞꮲ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮡꮯꮆꮣ|ꮎꮏ \d+[.,]?\d* ꮧꭷꮈꭲ|\d+[.,]?\d* ꭲꭶ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭰꮞꮲ|ꮎꮏ \d+[.,]?\d* ꭲꮿꮤ|ꮎꮏ \d+[.,]?\d* ꭷꮈꭲ|ꮎꮏ \d+[.,]?\d* ꭲꭶ|ꮎꮏ \d+[.,]?\d* ꭴꮥ|ꮎꮏ \d+[.,]?\d* ꭷꮈ|ꮎꮏ \d+[.,]?\d* ꮡꮯ|ꮎꮏ \d+[.,]?\d* ꮢꮎ)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(ꮎꮏ \d+[.,]?\d* ꭲᏻꮎꮩꮣꮖꮝꮧ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭲꮿꮤꮼꮝꮤꮕ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮢꮎꮩꮣꮖꮝꮧ ꮵꭸꮢ|\d+[.,]?\d* ꭲꮷꮥꮨᏼꮜꮧꮢꭲ ꮵꭸꮢ|\d+[.,]?\d* ꭿꮈꮝꭹ ꮷꮢꭿꮫ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮣꮣꮎꮹꮝꭼ ꮵꭸꮢ|\d+[.,]?\d* ꭴꮥꮨᏼꮜꮧꮢꭲ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭲꮷꮥꮨᏼꮜꮧꮢꭲ|ꮎꮏ \d+[.,]?\d* ꭿꮈꮝꭹ ꮷꮢꭿꮫ|ꮎꮏ \d+[.,]?\d* ꭲᏻꮎꮩꮣꮖꮝꮧ|ꮎꮏ \d+[.,]?\d* ꭴꮥꮨᏼꮜꮧꮢꭲ|ꮎꮏ \d+[.,]?\d* ꮧꭷꮈꭲ ꮵꭸꮢ|\d+[.,]?\d* ꮣꮣꮎꮹꮝꭼ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭲꮿꮤ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭲꮿꮤꮼꮝꮤꮕ|ꮎꮏ \d+[.,]?\d* ꭷꮈꭲ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮢꮎꮩꮣꮖꮝꮧ|\d+[.,]?\d* ꭲᏻꮯꮆꮣ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭴꮥ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭷꮈ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮡꮯ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮢꮎ ꮵꭸꮢ|\d+[.,]?\d* ꮡꮯꮆꮣ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭲᏻꮯꮆꮣ|\d+[.,]?\d* ꭰꮞꮲ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꮡꮯꮆꮣ|ꮎꮏ \d+[.,]?\d* ꮧꭷꮈꭲ|\d+[.,]?\d* ꭲꭶ ꮵꭸꮢ|ꮎꮏ \d+[.,]?\d* ꭰꮞꮲ|ꮎꮏ \d+[.,]?\d* ꭲꮿꮤ|ꮎꮏ \d+[.,]?\d* ꭷꮈꭲ|ꮎꮏ \d+[.,]?\d* ꭲꭶ|ꮎꮏ \d+[.,]?\d* ꭴꮥ|ꮎꮏ \d+[.,]?\d* ꭷꮈ|ꮎꮏ \d+[.,]?\d* ꮡꮯ|ꮎꮏ \d+[.,]?\d* ꮢꮎ)$`),
		KnownWords:      []string{"ꭿꭰ ꭲꮿꮤꮼꮝꮤꮕ", "ꭿꭰ ꮷꮥꮨᏼꮢꮨ", "ꭴꮥꮨᏼꮜꮧꮢꭲ", "ꭲꮿꮤꮼꮝꮤꮕ", "ꭴꮎꮩꮣꮖꮝꭼ", "ꭴꮎꮩꮣꮘꮥꮎ", "ꭴꮎꮩꮣꮙꮕꭿ", "ꭷꮈꭲ ꮵꭸꮢ", "ꭿꭰ ꭰꮅꮅꮜ", "ꭿꭰ ꮡꮯꮆꮣ", "ꮢꮎꮩꮣꮖꮝꮧ", "ꮤꮅꮑ ꭷꮈꭲ", "ꮵꮫꮅᏹꮅꮢꭲ", "ꭱꮨ ꮵꭸꮢ", "ꭿꭰ ꭷꮈꭲ", "ꮢꭿᏹꭲꮧꮲ", "ꮷꮎꭹꮆꮝꮧ", "ꭰꮒꮝꭼꮨ", "ꭴꮓꮈꮤꮕ", "ꭺꭿ ꭲꭶ", "ꮕꭹꮑꭲꭶ", "ꮠꮖꮄꮕꮂ", "ꮤꮅꮑꭲꭶ", "ꮶꭲꮑꭲꭶ", "ꭱꮨᏼꭲ", "ꭵꮝꭹᏹ", "ꭻᏸꮙꮒ", "ꮕꮣꮥꮖ", "ꮜꮎꮄꭲ", "ꮡꮯꮆꮣ", "ꮥꭽꮇᏹ", "ꮪꮅꮝꮧ", "ꮪꮒꮕꮧ", "gmt", "utc", "ꭰꮕᏹ", "ꭰꮞꮲ", "ꭲꮿꮤ", "ꭶꮆꮒ", "ꭷꭶꮅ", "ꭷꮈꭲ", "ꭷꮼꮒ", "ꮕꭹꮑ", "ꮖꮝꭼ", "ꮘꮥꮎ", "ꮙꮕꭿ", "ꮜꮎꮄ", "ꮤꮅꮑ", "ꮶꭲꮑ", "ꮷꮎꭹ", "am", "pm", "ꭰꮒ", "ꭰꮕ", "ꭲꭶ", "ꭴꮓ", "ꭴꮥ", "ꭵꮝ", "ꭶꮆ", "ꭷꭶ", "ꭷꮈ", "ꭷꮼ", "ꭻᏸ", "ꮓꮚ", "ꮕꮣ", "ꮡꮯ", "ꮢꭿ", "ꮢꮎ", "ꮥꭽ", "ꮪꮅ", "ꮪꮒ", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
