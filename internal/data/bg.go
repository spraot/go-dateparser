// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	bg_Locale LocaleData
)

func init() {
	bg_Locale = merge(nil, LocaleData{
		Name:                  "bg",
		DateOrder:             "DMY 'Г'.",
		Charset:               []rune(`cgtuzабвгдезийклмнопрстуфхцчщъюя`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])един(\z|[^\pL\pM\d_])`), "${1}1${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])два(\z|[^\pL\pM\d_])`), "${1}2${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])три(\z|[^\pL\pM\d_])`), "${1}3${2}"},
		},
		Translations: map[string][]string{
			"понеделник": {"monday"},
			"септември":  {"september"},
			"четвъртък":  {"thursday"},
			"декември":   {"december"},
			"октомври":   {"october"},
			"февруари":   {"february"},
			"вторник":    {"tuesday"},
			"ноември":    {"november"},
			"седмица":    {"week"},
			"седмици":    {"week"},
			"секунда":    {"second"},
			"секунди":    {"second"},
			"август":     {"august"},
			"година":     {"year"},
			"години":     {"year"},
			"месеци":     {"month"},
			"минута":     {"minute"},
			"минути":     {"minute"},
			"неделя":     {"sunday"},
			"събота":     {"saturday"},
			"януари":     {"january"},
			"април":      {"april"},
			"върху":      {""},
			"месец":      {"month"},
			"около":      {""},
			"петък":      {"friday"},
			"подир":      {"in"},
			"после":      {"in"},
			"преди":      {"ago"},
			"септм":      {"september"},
			"сряда":      {"wednesday"},
			"дена":       {"day"},
			"март":       {"march"},
			"проб":       {"am"},
			"седм":       {"week"},
			"септ":       {"september"},
			"след":       {"in"},
			"слоб":       {"pm"},
			"часа":       {"hour"},
			"gmt":        {"gmt"},
			"utc":        {"utc"},
			"авг":        {"august"},
			"апр":        {"april"},
			"вто":        {"tuesday"},
			"год":        {"year"},
			"дек":        {"december"},
			"ден":        {"day"},
			"дни":        {"day"},
			"маи":        {"may"},
			"мес":        {"month"},
			"мин":        {"minute"},
			"ное":        {"november"},
			"окт":        {"october"},
			"под":        {""},
			"пон":        {"monday"},
			"сед":        {"week"},
			"сек":        {"second"},
			"сеп":        {"september"},
			"сря":        {"wednesday"},
			"фев":        {"february"},
			"час":        {"hour"},
			"юли":        {"july"},
			"юни":        {"june"},
			"яну":        {"january"},
			"am":         {"am"},
			"pm":         {"pm"},
			"ап":         {"april"},
			"вт":         {"tuesday"},
			"до":         {""},
			"на":         {""},
			"нд":         {"sunday"},
			"от":         {""},
			"пн":         {"monday"},
			"пт":         {"friday"},
			"сб":         {"saturday"},
			"ср":         {"wednesday"},
			"фв":         {"february"},
			"чт":         {"thursday"},
			"юл":         {"july"},
			"юн":         {"june"},
			"ян":         {"january"},
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
			"г":          {"year"},
			"д":          {"day"},
			"м":          {"month"},
			"с":          {"week", "second"},
			"ч":          {"hour"},
		},
		RelativeType: map[string]string{
			"предходната седмица": "1 week ago",
			"след 1 десетилетие":  "in 10 year",
			"следващата седмица":  "in 1 week",
			"преди десетилетие":   "10 year ago",
			"следващата година":   "in 1 year",
			"миналата седмица":    "1 week ago",
			"миналата година":     "1 year ago",
			"предходен месец":     "1 month ago",
			"в тази минута":       "0 minute ago",
			"преди седмица":       "1 week ago",
			"следващ месец":       "in 1 month",
			"преди година":        "1 year ago",
			"тази седмица":        "0 week ago",
			"тази година":         "0 year ago",
			"в този час":          "0 hour ago",
			"следв седм":          "in 1 week",
			"този месец":          "0 month ago",
			"вдругиден":           "in 2 day",
			"преди ден":           "1 day ago",
			"преди час":           "1 hour ago",
			"следв мес":           "in 1 month",
			"тази седм":           "0 week ago",
			"мин седм":            "1 week ago",
			"онзи ден":            "2 day ago",
			"след ден":            "in 1 day",
			"след час":            "in 1 hour",
			"този мес":            "0 month ago",
			"мин мес":             "1 month ago",
			"сл седм":             "in 1 week",
			"следв г":             "in 1 year",
			"вчера":               "1 day ago",
			"мин г":               "1 year ago",
			"мин м":               "1 month ago",
			"снощи":               "1 day ago",
			"днес":                "0 day ago",
			"сега":                "0 second ago",
			"сл г":                "in 1 year",
			"сл м":                "in 1 month",
			"утре":                "in 1 day",
			"т г":                 "0 year ago",
			"т м":                 "0 month ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) седмица`), "$1 week ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) седмици`), "$1 week ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) секунда`), "$1 second ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) секунди`), "$1 second ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) година`), "$1 year ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) години`), "$1 year ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) месеца`), "$1 month ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) минута`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) минути`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) седмица`), "in $1 week"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) седмици`), "in $1 week"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) секунда`), "in $1 second"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) секунди`), "in $1 second"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) месец`), "$1 month ago"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) година`), "in $1 year"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) години`), "in $1 year"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) месеца`), "in $1 month"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) минута`), "in $1 minute"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) минути`), "in $1 minute"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) седм`), "$1 week ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) часа`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) месец`), "in $1 month"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) ден`), "$1 day ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) дни`), "$1 day ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) мин`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) сек`), "$1 second ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) час`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) седм`), "in $1 week"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) часа`), "in $1 hour"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) ден`), "in $1 day"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) дни`), "in $1 day"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) мин`), "in $1 minute"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) сек`), "in $1 second"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) час`), "in $1 hour"},
			{regexp.MustCompile(`(?i)пр (\d+[.,]?\d*) седм`), "$1 week ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) г`), "$1 year ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) м`), "$1 month ago"},
			{regexp.MustCompile(`(?i)преди (\d+[.,]?\d*) ч`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)сл (\d+[.,]?\d*) седм`), "in $1 week"},
			{regexp.MustCompile(`(?i)пр (\d+[.,]?\d*) мин`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)пр (\d+[.,]?\d*) сек`), "$1 second ago"},
			{regexp.MustCompile(`(?i)сл (\d+[.,]?\d*) мин`), "in $1 minute"},
			{regexp.MustCompile(`(?i)сл (\d+[.,]?\d*) сек`), "in $1 second"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) г`), "in $1 year"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) м`), "in $1 month"},
			{regexp.MustCompile(`(?i)след (\d+[.,]?\d*) ч`), "in $1 hour"},
			{regexp.MustCompile(`(?i)пр (\d+[.,]?\d*) г`), "$1 year ago"},
			{regexp.MustCompile(`(?i)пр (\d+[.,]?\d*) д`), "$1 day ago"},
			{regexp.MustCompile(`(?i)пр (\d+[.,]?\d*) м`), "$1 month ago"},
			{regexp.MustCompile(`(?i)пр (\d+[.,]?\d*) ч`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)сл (\d+[.,]?\d*) г`), "in $1 year"},
			{regexp.MustCompile(`(?i)сл (\d+[.,]?\d*) д`), "in $1 day"},
			{regexp.MustCompile(`(?i)сл (\d+[.,]?\d*) м`), "in $1 month"},
			{regexp.MustCompile(`(?i)сл (\d+[.,]?\d*) ч`), "in $1 hour"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(преди \d+[.,]?\d* седмица|преди \d+[.,]?\d* седмици|преди \d+[.,]?\d* секунда|преди \d+[.,]?\d* секунди|преди \d+[.,]?\d* година|преди \d+[.,]?\d* години|преди \d+[.,]?\d* месеца|преди \d+[.,]?\d* минута|преди \d+[.,]?\d* минути|след \d+[.,]?\d* седмица|след \d+[.,]?\d* седмици|след \d+[.,]?\d* секунда|след \d+[.,]?\d* секунди|преди \d+[.,]?\d* месец|след \d+[.,]?\d* година|след \d+[.,]?\d* години|след \d+[.,]?\d* месеца|след \d+[.,]?\d* минута|след \d+[.,]?\d* минути|преди \d+[.,]?\d* седм|преди \d+[.,]?\d* часа|след \d+[.,]?\d* месец|преди \d+[.,]?\d* ден|преди \d+[.,]?\d* дни|преди \d+[.,]?\d* мин|преди \d+[.,]?\d* сек|преди \d+[.,]?\d* час|след \d+[.,]?\d* седм|след \d+[.,]?\d* часа|след \d+[.,]?\d* ден|след \d+[.,]?\d* дни|след \d+[.,]?\d* мин|след \d+[.,]?\d* сек|след \d+[.,]?\d* час|пр \d+[.,]?\d* седм|преди \d+[.,]?\d* г|преди \d+[.,]?\d* м|преди \d+[.,]?\d* ч|сл \d+[.,]?\d* седм|пр \d+[.,]?\d* мин|пр \d+[.,]?\d* сек|сл \d+[.,]?\d* мин|сл \d+[.,]?\d* сек|след \d+[.,]?\d* г|след \d+[.,]?\d* м|след \d+[.,]?\d* ч|пр \d+[.,]?\d* г|пр \d+[.,]?\d* д|пр \d+[.,]?\d* м|пр \d+[.,]?\d* ч|сл \d+[.,]?\d* г|сл \d+[.,]?\d* д|сл \d+[.,]?\d* м|сл \d+[.,]?\d* ч)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(преди \d+[.,]?\d* седмица|преди \d+[.,]?\d* седмици|преди \d+[.,]?\d* секунда|преди \d+[.,]?\d* секунди|преди \d+[.,]?\d* година|преди \d+[.,]?\d* години|преди \d+[.,]?\d* месеца|преди \d+[.,]?\d* минута|преди \d+[.,]?\d* минути|след \d+[.,]?\d* седмица|след \d+[.,]?\d* седмици|след \d+[.,]?\d* секунда|след \d+[.,]?\d* секунди|преди \d+[.,]?\d* месец|след \d+[.,]?\d* година|след \d+[.,]?\d* години|след \d+[.,]?\d* месеца|след \d+[.,]?\d* минута|след \d+[.,]?\d* минути|преди \d+[.,]?\d* седм|преди \d+[.,]?\d* часа|след \d+[.,]?\d* месец|преди \d+[.,]?\d* ден|преди \d+[.,]?\d* дни|преди \d+[.,]?\d* мин|преди \d+[.,]?\d* сек|преди \d+[.,]?\d* час|след \d+[.,]?\d* седм|след \d+[.,]?\d* часа|след \d+[.,]?\d* ден|след \d+[.,]?\d* дни|след \d+[.,]?\d* мин|след \d+[.,]?\d* сек|след \d+[.,]?\d* час|пр \d+[.,]?\d* седм|преди \d+[.,]?\d* г|преди \d+[.,]?\d* м|преди \d+[.,]?\d* ч|сл \d+[.,]?\d* седм|пр \d+[.,]?\d* мин|пр \d+[.,]?\d* сек|сл \d+[.,]?\d* мин|сл \d+[.,]?\d* сек|след \d+[.,]?\d* г|след \d+[.,]?\d* м|след \d+[.,]?\d* ч|пр \d+[.,]?\d* г|пр \d+[.,]?\d* д|пр \d+[.,]?\d* м|пр \d+[.,]?\d* ч|сл \d+[.,]?\d* г|сл \d+[.,]?\d* д|сл \d+[.,]?\d* м|сл \d+[.,]?\d* ч)$`),
		KnownWords:      []string{"предходната седмица", "след 1 десетилетие", "следващата седмица", "преди десетилетие", "следващата година", "миналата седмица", "миналата година", "предходен месец", "в тази минута", "преди седмица", "следващ месец", "преди година", "тази седмица", "тази година", "в този час", "понеделник", "следв седм", "този месец", "вдругиден", "преди ден", "преди час", "септември", "следв мес", "тази седм", "четвъртък", "декември", "мин седм", "октомври", "онзи ден", "след ден", "след час", "този мес", "февруари", "вторник", "мин мес", "ноември", "седмица", "седмици", "секунда", "секунди", "сл седм", "следв г", "август", "година", "години", "месеци", "минута", "минути", "неделя", "събота", "януари", "април", "вчера", "върху", "месец", "мин г", "мин м", "около", "петък", "подир", "после", "преди", "септм", "снощи", "сряда", "дена", "днес", "март", "проб", "сега", "седм", "септ", "сл г", "сл м", "след", "слоб", "утре", "часа", "gmt", "utc", "авг", "апр", "вто", "год", "дек", "ден", "дни", "маи", "мес", "мин", "ное", "окт", "под", "пон", "сед", "сек", "сеп", "сря", "т г", "т м", "фев", "час", "юли", "юни", "яну", "am", "pm", "ап", "вт", "до", "на", "нд", "от", "пн", "пт", "сб", "ср", "фв", "чт", "юл", "юн", "ян", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "г", "д", "м", "с", "ч"},
	})
}
