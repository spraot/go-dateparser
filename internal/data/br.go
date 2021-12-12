// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var br_Locale = LocaleData{
	Name:                  "br",
	DateOrder:             "YMD",
	January:               []string{"gen", "genver"},
	February:              []string{"c'hwe", "c'hwevrer"},
	March:                 []string{"meur", "meurzh"},
	April:                 []string{"ebr", "ebrel"},
	May:                   []string{"mae"},
	June:                  []string{"mezh", "mezheven"},
	July:                  []string{"goue", "gouere"},
	August:                []string{"eost"},
	September:             []string{"gwen", "gwengolo"},
	October:               []string{"here"},
	November:              []string{"du"},
	December:              []string{"ker", "kerzu", "kzu"},
	Monday:                []string{"lun"},
	Tuesday:               []string{"meu", "meurzh"},
	Wednesday:             []string{"mer", "merc'her"},
	Thursday:              []string{"yaou"},
	Friday:                []string{"gwe", "gwener"},
	Saturday:              []string{"sad", "sadorn"},
	Sunday:                []string{"sul"},
	AM:                    []string{"am"},
	PM:                    []string{"gm"},
	Year:                  []string{"bl", "bloaz"},
	Month:                 []string{"miz"},
	Week:                  []string{"sizhun"},
	Day:                   []string{"d", "deiz"},
	Hour:                  []string{"e", "eur"},
	Minute:                []string{"min", "munut"},
	Second:                []string{"eilenn", "s"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`hiziv`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`ar miz-mañ`},
		`0 second ago`: {`brem`, `bremañ`},
		`0 week ago`:   {`ar sizhun-mañ`},
		`0 year ago`:   {`hevlene`},
		`1 day ago`:    {`dec'h`},
		`1 month ago`:  {`ar miz diaraok`},
		`1 week ago`:   {`ar sizhun diaraok`},
		`1 year ago`:   {`warlene`},
		`in 1 day`:     {`warc'hoazh`},
		`in 1 month`:   {`ar miz a zeu`},
		`in 1 week`:    {`ar sizhun a zeu`},
		`in 1 year`:    {`ar bl a zeu`, `ar bloaz a zeu`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) d zo`),
			regexp.MustCompile(`(?i)(\d+) deiz zo`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) e zo`),
			regexp.MustCompile(`(?i)(\d+) eur zo`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) min zo`),
			regexp.MustCompile(`(?i)(\d+) munut zo`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) miz zo`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) eilenn zo`),
			regexp.MustCompile(`(?i)(\d+) s zo`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) sizhun zo`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) bl zo`),
			regexp.MustCompile(`(?i)(\d+) bloaz zo`),
			regexp.MustCompile(`(?i)(\d+) vloaz zo`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)a-benn (\d+) d`),
			regexp.MustCompile(`(?i)a-benn (\d+) deiz`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)a-benn (\d+) e`),
			regexp.MustCompile(`(?i)a-benn (\d+) eur`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)a-benn (\d+) min`),
			regexp.MustCompile(`(?i)a-benn (\d+) munut`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)a-benn (\d+) miz`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)a-benn (\d+) eilenn`),
			regexp.MustCompile(`(?i)a-benn (\d+) s`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)a-benn (\d+) sizhun`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)a-benn (\d+) bl`),
			regexp.MustCompile(`(?i)a-benn (\d+) bloaz`),
			regexp.MustCompile(`(?i)a-benn (\d+) vloaz`),
		},
	},
}
