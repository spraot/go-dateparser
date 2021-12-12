// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var sah_Locale = LocaleData{
	Name:                  "sah",
	DateOrder:             "YMD",
	January:               []string{"тохс", "тохсунньу"},
	February:              []string{"олун", "олунньу"},
	March:                 []string{"клн", "кулун тутар"},
	April:                 []string{"мсу", "муус устар"},
	May:                   []string{"ыам", "ыам ыйа", "ыам ыйын"},
	June:                  []string{"бэс", "бэс ыйа", "бэс ыйын"},
	July:                  []string{"от ыйа", "от ыйын", "отй"},
	August:                []string{"атр", "атырдьых ыйа", "атырдьых ыйын"},
	September:             []string{"балаҕан ыйа", "балаҕан ыйын", "блҕ"},
	October:               []string{"алт", "алтынньы"},
	November:              []string{"сэт", "сэтинньи"},
	December:              []string{"ахс", "ахсынньы"},
	Monday:                []string{"бн", "бэнидиэнньик"},
	Tuesday:               []string{"оп", "оптуорунньук"},
	Wednesday:             []string{"сэ", "сэрэдэ"},
	Thursday:              []string{"чп", "чэппиэр"},
	Friday:                []string{"бэ", "бээтиҥсэ"},
	Saturday:              []string{"сб", "субуота"},
	Sunday:                []string{"баскыһыанньа", "бс"},
	AM:                    []string{"эи"},
	PM:                    []string{"эк"},
	Year:                  []string{"сыл"},
	Month:                 []string{"ый"},
	Week:                  []string{"нэдиэлэ"},
	Day:                   []string{"күн"},
	Hour:                  []string{"чаас"},
	Minute:                []string{"мүнүүтэ"},
	Second:                []string{"сөкүүндэ"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`бүгүн`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`бу ый`},
		`0 second ago`: {`билигин`},
		`0 week ago`:   {`бу нэдиэлэ`},
		`0 year ago`:   {`быйыл`},
		`1 day ago`:    {`бэҕэһээ`},
		`1 month ago`:  {`ааспыт ый`},
		`1 week ago`:   {`ааспыт нэдиэлэ`},
		`1 year ago`:   {`былырыын`},
		`in 1 day`:     {`сарсын`},
		`in 1 month`:   {`аныгыскы ый`},
		`in 1 week`:    {`кэлэр нэдиэлэ`},
		`in 1 year`:    {`эһиил`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) күн ынараа өттүгэр`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) чаас ынараа өттүгэр`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) мүнүүтэ ынараа өттүгэр`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ый ынараа өттүгэр`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) сөк анараа өттүгэр`),
			regexp.MustCompile(`(?i)(\d+) сөкүүндэ ынараа өттүгэр`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) нэдиэлэ анараа өттүгэр`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) сыл ынараа өттүгэр`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)(\d+) күнүнэн`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)(\d+) чааһынан`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)(\d+) мүнүүтэннэн`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)(\d+) ыйынан`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)(\d+) сөкүүндэннэн`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)(\d+) нэдиэлэннэн`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)(\d+) сылынан`),
		},
	},
}
