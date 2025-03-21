// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	mk_Locale LocaleData
)

func init() {
	mk_Locale = merge(nil, LocaleData{
		Name:      "mk",
		DateOrder: "DMY",
		Charset:   []rune(`cgtuzабвгдеиклмнопрстуфцчј`),
		Translations: map[string][]string{
			"понеделник": {"monday"},
			"претпладне": {"am"},
			"септември":  {"september"},
			"декември":   {"december"},
			"октомври":   {"october"},
			"попладне":   {"pm"},
			"февруари":   {"february"},
			"четврток":   {"thursday"},
			"вторник":    {"tuesday"},
			"ноември":    {"november"},
			"секунда":    {"second"},
			"јануари":    {"january"},
			"август":     {"august"},
			"година":     {"year"},
			"минута":     {"minute"},
			"недела":     {"week", "sunday"},
			"претпл":     {"am"},
			"сабота":     {"saturday"},
			"април":      {"april"},
			"месец":      {"month"},
			"петок":      {"friday"},
			"среда":      {"wednesday"},
			"март":       {"march"},
			"ноем":       {"november"},
			"попл":       {"pm"},
			"септ":       {"september"},
			"јули":       {"july"},
			"јуни":       {"june"},
			"gmt":        {"gmt"},
			"utc":        {"utc"},
			"авг":        {"august"},
			"апр":        {"april"},
			"вто":        {"tuesday"},
			"год":        {"year"},
			"дек":        {"december"},
			"ден":        {"day"},
			"мар":        {"march"},
			"мај":        {"may"},
			"мес":        {"month"},
			"мин":        {"minute"},
			"нед":        {"sunday"},
			"окт":        {"october"},
			"пет":        {"friday"},
			"пон":        {"monday"},
			"саб":        {"saturday"},
			"сед":        {"week"},
			"сек":        {"second"},
			"сре":        {"wednesday"},
			"фев":        {"february"},
			"час":        {"hour"},
			"чет":        {"thursday"},
			"јан":        {"january"},
			"јул":        {"july"},
			"јун":        {"june"},
			"am":         {"am"},
			"pm":         {"pm"},
			"вт":         {"tuesday"},
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
			"минатата седмица": "1 week ago",
			"следната седмица": "in 1 week",
			"минатата година":  "1 year ago",
			"следната година":  "in 1 year",
			"минатиот месец":   "1 month ago",
			"следниот месец":   "in 1 month",
			"оваа седмица":     "0 week ago",
			"оваа година":      "0 year ago",
			"оваа минута":      "0 minute ago",
			"овој месец":       "0 month ago",
			"вчера":            "1 day ago",
			"денес":            "0 day ago",
			"часов":            "0 hour ago",
			"сега":             "0 second ago",
			"утре":             "in 1 day",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) седмица`), "$1 week ago"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) седмици`), "$1 week ago"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) секунда`), "$1 second ago"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) секунди`), "$1 second ago"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) година`), "$1 year ago"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) години`), "$1 year ago"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) месеци`), "$1 month ago"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) минута`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) минути`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) седмица`), "in $1 week"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) седмици`), "in $1 week"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) секунда`), "in $1 second"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) секунди`), "in $1 second"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) месец`), "$1 month ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) година`), "in $1 year"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) години`), "in $1 year"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) месеци`), "in $1 month"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) минута`), "in $1 minute"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) минути`), "in $1 minute"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) дена`), "$1 day ago"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) часа`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) месец`), "in $1 month"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) ден`), "$1 day ago"},
			{regexp.MustCompile(`(?i)пред (\d+[.,]?\d*) час`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) дена`), "in $1 day"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) часа`), "in $1 hour"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) ден`), "in $1 day"},
			{regexp.MustCompile(`(?i)за (\d+[.,]?\d*) час`), "in $1 hour"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(пред \d+[.,]?\d* седмица|пред \d+[.,]?\d* седмици|пред \d+[.,]?\d* секунда|пред \d+[.,]?\d* секунди|пред \d+[.,]?\d* година|пред \d+[.,]?\d* години|пред \d+[.,]?\d* месеци|пред \d+[.,]?\d* минута|пред \d+[.,]?\d* минути|за \d+[.,]?\d* седмица|за \d+[.,]?\d* седмици|за \d+[.,]?\d* секунда|за \d+[.,]?\d* секунди|пред \d+[.,]?\d* месец|за \d+[.,]?\d* година|за \d+[.,]?\d* години|за \d+[.,]?\d* месеци|за \d+[.,]?\d* минута|за \d+[.,]?\d* минути|пред \d+[.,]?\d* дена|пред \d+[.,]?\d* часа|за \d+[.,]?\d* месец|пред \d+[.,]?\d* ден|пред \d+[.,]?\d* час|за \d+[.,]?\d* дена|за \d+[.,]?\d* часа|за \d+[.,]?\d* ден|за \d+[.,]?\d* час)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(пред \d+[.,]?\d* седмица|пред \d+[.,]?\d* седмици|пред \d+[.,]?\d* секунда|пред \d+[.,]?\d* секунди|пред \d+[.,]?\d* година|пред \d+[.,]?\d* години|пред \d+[.,]?\d* месеци|пред \d+[.,]?\d* минута|пред \d+[.,]?\d* минути|за \d+[.,]?\d* седмица|за \d+[.,]?\d* седмици|за \d+[.,]?\d* секунда|за \d+[.,]?\d* секунди|пред \d+[.,]?\d* месец|за \d+[.,]?\d* година|за \d+[.,]?\d* години|за \d+[.,]?\d* месеци|за \d+[.,]?\d* минута|за \d+[.,]?\d* минути|пред \d+[.,]?\d* дена|пред \d+[.,]?\d* часа|за \d+[.,]?\d* месец|пред \d+[.,]?\d* ден|пред \d+[.,]?\d* час|за \d+[.,]?\d* дена|за \d+[.,]?\d* часа|за \d+[.,]?\d* ден|за \d+[.,]?\d* час)$`),
		KnownWords:      []string{"минатата седмица", "следната седмица", "минатата година", "следната година", "минатиот месец", "следниот месец", "оваа седмица", "оваа година", "оваа минута", "овој месец", "понеделник", "претпладне", "септември", "декември", "октомври", "попладне", "февруари", "четврток", "вторник", "ноември", "секунда", "јануари", "август", "година", "минута", "недела", "претпл", "сабота", "април", "вчера", "денес", "месец", "петок", "среда", "часов", "март", "ноем", "попл", "сега", "септ", "утре", "јули", "јуни", "gmt", "utc", "авг", "апр", "вто", "год", "дек", "ден", "мар", "мај", "мес", "мин", "нед", "окт", "пет", "пон", "саб", "сед", "сек", "сре", "фев", "час", "чет", "јан", "јул", "јун", "am", "pm", "вт", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|"},
	})
}
