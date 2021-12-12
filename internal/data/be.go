// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var be_Locale = LocaleData{
	Name:                  "be",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "каля", "у", "і", "ў"},
	January:               []string{"стд", "сту", "студзень", "студзеня"},
	February:              []string{"лют", "лютага", "люты"},
	March:                 []string{"сак", "сакавік", "сакавіка"},
	April:                 []string{"кра", "красавік", "красавіка", "крс"},
	May:                   []string{"май", "мая", "тра", "траўня"},
	June:                  []string{"чэр", "чэрвень", "чэрвеня"},
	July:                  []string{"ліп", "ліпень", "ліпеня"},
	August:                []string{"жнв", "жні", "жнівень", "жнівеня", "жніўня"},
	September:             []string{"вер", "верасень", "верасня", "врс"},
	October:               []string{"кас", "кастрычнік", "кастрычніка", "кст"},
	November:              []string{"ліс", "лістапад", "лістапада"},
	December:              []string{"сне", "снежань", "снежня", "снж"},
	Monday:                []string{"панядзелак", "пн", "пнд"},
	Tuesday:               []string{"аў", "аўт", "аўторак"},
	Wednesday:             []string{"серада", "ср"},
	Thursday:              []string{"чацвер", "чв", "чц", "чцв"},
	Friday:                []string{"пт", "пят", "пятніца"},
	Saturday:              []string{"сб", "суб", "субота"},
	Sunday:                []string{"нд", "няд", "нядзеля"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"г", "гадоў", "гады", "год", "года"},
	Month:                 []string{"мес", "месяц", "месяца", "месяцаў", "месяцы"},
	Week:                  []string{"тыд", "тыдзень", "тыдня", "тыдняў", "тыдні"},
	Day:                   []string{"д", "дзен", "дзень", "дзён", "дні"},
	Hour:                  []string{"гадз", "гадзін", "гадзіна", "гадзіну", "гадзіны"},
	Minute:                []string{"хв", "хвіл", "хвілін", "хвіліна", "хвіліну", "хвіліны"},
	Second:                []string{"с", "сек", "секунд", "секунда", "секунду", "секунды"},
	In:                    []string{"на працягу"},
	Ago:                   []string{"назад", "таму", "таму назад"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`сення`, `сёння`},
		`0 hour ago`:   {`у гэту гадзіну`},
		`0 minute ago`: {`у гэту хвіліну`},
		`0 month ago`:  {`у гэтым месяцы`},
		`0 second ago`: {`цяпер`},
		`0 week ago`:   {`на гэтым тыдні`},
		`0 year ago`:   {`у гэтым годзе`},
		`1 day ago`:    {`учора`, `ўчора`},
		`1 month ago`:  {`у мінулым месяцы`},
		`1 week ago`:   {`на мінулым тыдні`},
		`1 year ago`:   {`у мінулым годзе`},
		`2 day ago`:    {`пазаўчора`},
		`in 1 day`:     {`заўтра`},
		`in 1 month`:   {`у наступным месяцы`},
		`in 1 week`:    {`на наступным тыдні`},
		`in 1 year`:    {`у наступным годзе`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) д таму`),
			regexp.MustCompile(`(?i)(\d+) дзень таму`),
			regexp.MustCompile(`(?i)(\d+) дня таму`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) гадз таму`),
			regexp.MustCompile(`(?i)(\d+) гадзіну таму`),
			regexp.MustCompile(`(?i)(\d+) гадзіны таму`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) хв таму`),
			regexp.MustCompile(`(?i)(\d+) хвіліну таму`),
			regexp.MustCompile(`(?i)(\d+) хвіліны таму`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) мес таму`),
			regexp.MustCompile(`(?i)(\d+) месяц таму`),
			regexp.MustCompile(`(?i)(\d+) месяца таму`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) с таму`),
			regexp.MustCompile(`(?i)(\d+) секунду таму`),
			regexp.MustCompile(`(?i)(\d+) секунды таму`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) тыд таму`),
			regexp.MustCompile(`(?i)(\d+) тыдзень таму`),
			regexp.MustCompile(`(?i)(\d+) тыдня таму`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) г таму`),
			regexp.MustCompile(`(?i)(\d+) год таму`),
			regexp.MustCompile(`(?i)(\d+) года таму`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)праз (\d+) д`),
			regexp.MustCompile(`(?i)праз (\d+) дзень`),
			regexp.MustCompile(`(?i)праз (\d+) дня`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)праз (\d+) гадз`),
			regexp.MustCompile(`(?i)праз (\d+) гадзіну`),
			regexp.MustCompile(`(?i)праз (\d+) гадзіны`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)праз (\d+) хв`),
			regexp.MustCompile(`(?i)праз (\d+) хвіліну`),
			regexp.MustCompile(`(?i)праз (\d+) хвіліны`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)праз (\d+) мес`),
			regexp.MustCompile(`(?i)праз (\d+) месяц`),
			regexp.MustCompile(`(?i)праз (\d+) месяца`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)праз (\d+) с`),
			regexp.MustCompile(`(?i)праз (\d+) секунду`),
			regexp.MustCompile(`(?i)праз (\d+) секунды`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)праз (\d+) тыд`),
			regexp.MustCompile(`(?i)праз (\d+) тыдзень`),
			regexp.MustCompile(`(?i)праз (\d+) тыдня`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)праз (\d+) г`),
			regexp.MustCompile(`(?i)праз (\d+) год`),
			regexp.MustCompile(`(?i)праз (\d+) года`),
		},
	},
	Simplifications: map[string]string{
		`(\d+)\s*гадзін\s(\d+)\s*хвілін`: `$1:$2`,
		`^гадзіна`:        `1 гадзіна`,
		`^секунду`:        `1 секунду`,
		`^хвіліну`:        `1 хвіліну`,
		`гадзіну`:         `1 гадзіну`,
		`некалькі секунд`: `44 секунды`,
		`некалькі хвілін`: `2 хвіліны`,
	},
}
