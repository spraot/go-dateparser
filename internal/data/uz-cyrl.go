// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var uz_Cyrl_Locale = LocaleData{
	Name:                  "uz-Cyrl",
	DateOrder:             "DMY",
	January:               []string{"янв", "январ"},
	February:              []string{"фев", "феврал"},
	March:                 []string{"мар", "март"},
	April:                 []string{"апр", "апрел"},
	May:                   []string{"май"},
	June:                  []string{"июн"},
	July:                  []string{"июл"},
	August:                []string{"авг", "август"},
	September:             []string{"сен", "сентябр"},
	October:               []string{"окт", "октябр"},
	November:              []string{"ноя", "ноябр"},
	December:              []string{"дек", "декабр"},
	Monday:                []string{"душ", "душанба"},
	Tuesday:               []string{"сеш", "сешанба"},
	Wednesday:             []string{"чор", "чоршанба"},
	Thursday:              []string{"пай", "пайшанба"},
	Friday:                []string{"жум", "жума"},
	Saturday:              []string{"шан", "шанба"},
	Sunday:                []string{"якш", "якшанба"},
	AM:                    []string{"то"},
	PM:                    []string{"тк"},
	Year:                  []string{"йил"},
	Month:                 []string{"ой"},
	Week:                  []string{"ҳафта"},
	Day:                   []string{"кун"},
	Hour:                  []string{"соат"},
	Minute:                []string{"дақиқа"},
	Second:                []string{"сония"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`бугун`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`бу ой`},
		`0 second ago`: {`ҳозир`},
		`0 week ago`:   {`бу ҳафта`},
		`0 year ago`:   {`бу йил`},
		`1 day ago`:    {`кеча`},
		`1 month ago`:  {`ўтган ой`},
		`1 week ago`:   {`ўтган ҳафта`},
		`1 year ago`:   {`ўтган йил`},
		`in 1 day`:     {`эртага`},
		`in 1 month`:   {`кейинги ой`},
		`in 1 week`:    {`кейинги ҳафта`},
		`in 1 year`:    {`кейинги йил`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) кун олдин`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) соат олдин`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) дақиқа олдин`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ой аввал`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) сония олдин`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) ҳафта олдин`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) йил аввал`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)(\d+) кундан сўнг`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)(\d+) соатдан сўнг`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)(\d+) дақиқадан сўнг`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)(\d+) ойдан сўнг`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)(\d+) сониядан сўнг`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)(\d+) ҳафтадан сўнг`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)(\d+) йилдан сўнг`),
		},
	},
}
