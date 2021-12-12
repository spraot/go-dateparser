// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var bs_Latn_Locale = LocaleData{
	Name:                  "bs-Latn",
	DateOrder:             "DMY.",
	January:               []string{"jan", "januar"},
	February:              []string{"feb", "februar"},
	March:                 []string{"mar", "mart"},
	April:                 []string{"apr", "april"},
	May:                   []string{"maj"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "juli"},
	August:                []string{"avg", "avgust"},
	September:             []string{"sep", "septembar"},
	October:               []string{"okt", "oktobar"},
	November:              []string{"nov", "novembar"},
	December:              []string{"dec", "decembar"},
	Monday:                []string{"pon", "ponedjeljak"},
	Tuesday:               []string{"uto", "utorak"},
	Wednesday:             []string{"sri", "srijeda"},
	Thursday:              []string{"čet", "četvrtak"},
	Friday:                []string{"pet", "petak"},
	Saturday:              []string{"sub", "subota"},
	Sunday:                []string{"ned", "nedjelja"},
	AM:                    []string{"prijepodne"},
	PM:                    []string{"popodne"},
	Year:                  []string{"g", "god", "godina"},
	Month:                 []string{"mj", "mjesec"},
	Week:                  []string{"sed", "sedmica"},
	Day:                   []string{"dan"},
	Hour:                  []string{"h", "sat"},
	Minute:                []string{"min", "minuta"},
	Second:                []string{"s", "sek", "sekunda"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`danas`},
		`0 hour ago`:   {`ovaj sat`},
		`0 minute ago`: {`ova minuta`},
		`0 month ago`:  {`ovaj mjesec`},
		`0 second ago`: {`sada`},
		`0 week ago`:   {`ove sedmice`},
		`0 year ago`:   {`ove godine`},
		`1 day ago`:    {`jučer`},
		`1 month ago`:  {`prošli mjesec`},
		`1 week ago`:   {`prošle sedmice`},
		`1 year ago`:   {`prošle godine`},
		`in 1 day`:     {`sutra`},
		`in 1 month`:   {`sljedeći mjesec`},
		`in 1 week`:    {`sljedeće sedmice`},
		`in 1 year`:    {`sljedeće godine`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)prije (\d+) d`),
			regexp.MustCompile(`(?i)prije (\d+) dan`),
			regexp.MustCompile(`(?i)prije (\d+) dana`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)prije (\d+) sat`),
			regexp.MustCompile(`(?i)prije (\d+) sati`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)prije (\d+) min`),
			regexp.MustCompile(`(?i)prije (\d+) minuta`),
			regexp.MustCompile(`(?i)prije (\d+) minutu`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)prije (\d+) mj`),
			regexp.MustCompile(`(?i)prije (\d+) mjesec`),
			regexp.MustCompile(`(?i)prije (\d+) mjeseci`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)prije (\d+) sek`),
			regexp.MustCompile(`(?i)prije (\d+) sekundi`),
			regexp.MustCompile(`(?i)prije (\d+) sekundu`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)prije (\d+) sed`),
			regexp.MustCompile(`(?i)prije (\d+) sedmica`),
			regexp.MustCompile(`(?i)prije (\d+) sedmicu`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)prije (\d+) g`),
			regexp.MustCompile(`(?i)prije (\d+) god`),
			regexp.MustCompile(`(?i)prije (\d+) godina`),
			regexp.MustCompile(`(?i)prije (\d+) godinu`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)za (\d+) d`),
			regexp.MustCompile(`(?i)za (\d+) dan`),
			regexp.MustCompile(`(?i)za (\d+) dana`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)za (\d+) sat`),
			regexp.MustCompile(`(?i)za (\d+) sati`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)za (\d+) min`),
			regexp.MustCompile(`(?i)za (\d+) minuta`),
			regexp.MustCompile(`(?i)za (\d+) minutu`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)za (\d+) mj`),
			regexp.MustCompile(`(?i)za (\d+) mjesec`),
			regexp.MustCompile(`(?i)za (\d+) mjeseci`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)za (\d+) sek`),
			regexp.MustCompile(`(?i)za (\d+) sekundi`),
			regexp.MustCompile(`(?i)za (\d+) sekundu`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)za (\d+) sed`),
			regexp.MustCompile(`(?i)za (\d+) sedmica`),
			regexp.MustCompile(`(?i)za (\d+) sedmicu`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)za (\d+) g`),
			regexp.MustCompile(`(?i)za (\d+) god`),
			regexp.MustCompile(`(?i)za (\d+) godina`),
			regexp.MustCompile(`(?i)za (\d+) godinu`),
		},
	},
}
