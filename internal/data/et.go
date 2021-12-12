// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var et_Locale = LocaleData{
	Name:                  "et",
	DateOrder:             "DMY",
	January:               []string{"jaan", "jaanuar"},
	February:              []string{"veebr", "veebruar"},
	March:                 []string{"märts"},
	April:                 []string{"apr", "aprill"},
	May:                   []string{"mai"},
	June:                  []string{"juuni"},
	July:                  []string{"juuli"},
	August:                []string{"aug", "august"},
	September:             []string{"sept", "september"},
	October:               []string{"okt", "oktoober"},
	November:              []string{"nov", "november"},
	December:              []string{"dets", "detsember"},
	Monday:                []string{"e", "esmaspäev"},
	Tuesday:               []string{"t", "teisipäev"},
	Wednesday:             []string{"k", "kolmapäev"},
	Thursday:              []string{"n", "neljapäev"},
	Friday:                []string{"r", "reede"},
	Saturday:              []string{"l", "laupäev"},
	Sunday:                []string{"p", "pühapäev"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"a", "aasta"},
	Month:                 []string{"k", "kuu"},
	Week:                  []string{"näd", "nädal"},
	Day:                   []string{"p", "päev"},
	Hour:                  []string{"t", "tund"},
	Minute:                []string{"min", "minut"},
	Second:                []string{"s", "sek", "sekund"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`täna`},
		`0 hour ago`:   {`praegusel tunnil`},
		`0 minute ago`: {`praegusel minutil`},
		`0 month ago`:  {`käesolev kuu`},
		`0 second ago`: {`nüüd`},
		`0 week ago`:   {`käesolev nädal`},
		`0 year ago`:   {`käesolev aasta`},
		`1 day ago`:    {`eile`},
		`1 month ago`:  {`eelmine kuu`},
		`1 week ago`:   {`eelmine nädal`},
		`1 year ago`:   {`eelmine aasta`},
		`in 1 day`:     {`homme`},
		`in 1 month`:   {`järgmine kuu`},
		`in 1 week`:    {`järgmine nädal`},
		`in 1 year`:    {`järgmine aasta`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) p eest`),
			regexp.MustCompile(`(?i)(\d+) päeva eest`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) t eest`),
			regexp.MustCompile(`(?i)(\d+) tunni eest`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) min eest`),
			regexp.MustCompile(`(?i)(\d+) minuti eest`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) k eest`),
			regexp.MustCompile(`(?i)(\d+) kuu eest`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) s eest`),
			regexp.MustCompile(`(?i)(\d+) sek eest`),
			regexp.MustCompile(`(?i)(\d+) sekundi eest`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) näd eest`),
			regexp.MustCompile(`(?i)(\d+) nädala eest`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) a eest`),
			regexp.MustCompile(`(?i)(\d+) aasta eest`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)(\d+) p pärast`),
			regexp.MustCompile(`(?i)(\d+) päeva pärast`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)(\d+) t pärast`),
			regexp.MustCompile(`(?i)(\d+) tunni pärast`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)(\d+) min pärast`),
			regexp.MustCompile(`(?i)(\d+) minuti pärast`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)(\d+) k pärast`),
			regexp.MustCompile(`(?i)(\d+) kuu pärast`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)(\d+) s pärast`),
			regexp.MustCompile(`(?i)(\d+) sek pärast`),
			regexp.MustCompile(`(?i)(\d+) sekundi pärast`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)(\d+) näd pärast`),
			regexp.MustCompile(`(?i)(\d+) nädala pärast`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)(\d+) a pärast`),
			regexp.MustCompile(`(?i)(\d+) aasta pärast`),
		},
	},
}
