// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var fo_Locale = LocaleData{
	Name:                  "fo",
	DateOrder:             "DMY",
	January:               []string{"jan", "januar"},
	February:              []string{"feb", "februar"},
	March:                 []string{"mar", "mars"},
	April:                 []string{"apr", "apríl"},
	May:                   []string{"mai"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "juli"},
	August:                []string{"aug", "august"},
	September:             []string{"sep", "september"},
	October:               []string{"okt", "oktober"},
	November:              []string{"nov", "november"},
	December:              []string{"des", "desember"},
	Monday:                []string{"mán", "mánadagur"},
	Tuesday:               []string{"týs", "týsdagur"},
	Wednesday:             []string{"mik", "mikudagur"},
	Thursday:              []string{"hós", "hósdagur"},
	Friday:                []string{"frí", "fríggjadagur"},
	Saturday:              []string{"ley", "leygardagur"},
	Sunday:                []string{"sun", "sunnudagur"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"ár"},
	Month:                 []string{"mnð", "mánaður"},
	Week:                  []string{"v", "vi", "vika"},
	Day:                   []string{"d", "da", "dagur"},
	Hour:                  []string{"t", "tími"},
	Minute:                []string{"m", "min", "minuttur"},
	Second:                []string{"s", "sek", "sekund"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`í dag`},
		`0 hour ago`:   {`hendan tíman`},
		`0 minute ago`: {`hendan minuttin`},
		`0 month ago`:  {`henda mánaðin`},
		`0 second ago`: {`nú`},
		`0 week ago`:   {`hesu viku`},
		`0 year ago`:   {`í ár`},
		`1 day ago`:    {`í gjár`},
		`1 month ago`:  {`seinasta mánað`},
		`1 week ago`:   {`seinastu viku`},
		`1 year ago`:   {`í fjør`},
		`in 1 day`:     {`í morgin`},
		`in 1 month`:   {`næsta mánað`},
		`in 1 week`:    {`næstu viku`},
		`in 1 year`:    {`næsta ár`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) d síðan`),
			regexp.MustCompile(`(?i)(\d+) da síðan`),
			regexp.MustCompile(`(?i)(\d+) dagar síðan`),
			regexp.MustCompile(`(?i)(\d+) dagur síðan`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) t síðan`),
			regexp.MustCompile(`(?i)(\d+) tímar síðan`),
			regexp.MustCompile(`(?i)(\d+) tími síðan`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) m síðan`),
			regexp.MustCompile(`(?i)(\d+) min síðan`),
			regexp.MustCompile(`(?i)(\d+) minutt síðan`),
			regexp.MustCompile(`(?i)(\d+) minuttir síðan`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) mnð síðan`),
			regexp.MustCompile(`(?i)(\d+) mánað síðan`),
			regexp.MustCompile(`(?i)(\d+) mánaðir síðan`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) s síðan`),
			regexp.MustCompile(`(?i)(\d+) sek síðan`),
			regexp.MustCompile(`(?i)(\d+) sekund síðan`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) v síðan`),
			regexp.MustCompile(`(?i)(\d+) vi síðan`),
			regexp.MustCompile(`(?i)(\d+) vika síðan`),
			regexp.MustCompile(`(?i)(\d+) vikur síðan`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) ár síðan`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)um (\d+) d`),
			regexp.MustCompile(`(?i)um (\d+) da`),
			regexp.MustCompile(`(?i)um (\d+) dag`),
			regexp.MustCompile(`(?i)um (\d+) dagar`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)um (\d+) t`),
			regexp.MustCompile(`(?i)um (\d+) tíma`),
			regexp.MustCompile(`(?i)um (\d+) tímar`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)um (\d+) m`),
			regexp.MustCompile(`(?i)um (\d+) min`),
			regexp.MustCompile(`(?i)um (\d+) minutt`),
			regexp.MustCompile(`(?i)um (\d+) minuttir`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)um (\d+) mnð`),
			regexp.MustCompile(`(?i)um (\d+) mánað`),
			regexp.MustCompile(`(?i)um (\d+) mánaðir`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)um (\d+) s`),
			regexp.MustCompile(`(?i)um (\d+) sek`),
			regexp.MustCompile(`(?i)um (\d+) sekund`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)um (\d+) v`),
			regexp.MustCompile(`(?i)um (\d+) vi`),
			regexp.MustCompile(`(?i)um (\d+) viku`),
			regexp.MustCompile(`(?i)um (\d+) vikur`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)um (\d+) ár`),
		},
	},
}

var fo_DK_Locale = LocaleData{
	Name:                  "fo-DK",
	DateOrder:             "DMY",
	January:               []string{"jan", "januar"},
	February:              []string{"feb", "februar"},
	March:                 []string{"mar", "mars"},
	April:                 []string{"apr", "apríl"},
	May:                   []string{"mai"},
	June:                  []string{"jun", "juni"},
	July:                  []string{"jul", "juli"},
	August:                []string{"aug", "august"},
	September:             []string{"sep", "september"},
	October:               []string{"okt", "oktober"},
	November:              []string{"nov", "november"},
	December:              []string{"des", "desember"},
	Monday:                []string{"mán", "mánadagur"},
	Tuesday:               []string{"týs", "týsdagur"},
	Wednesday:             []string{"mik", "mikudagur"},
	Thursday:              []string{"hós", "hósdagur"},
	Friday:                []string{"frí", "fríggjadagur"},
	Saturday:              []string{"ley", "leygardagur"},
	Sunday:                []string{"sun", "sunnudagur"},
	AM:                    []string{"am"},
	PM:                    []string{"pm"},
	Year:                  []string{"ár"},
	Month:                 []string{"mnð", "mánaður"},
	Week:                  []string{"v", "vi", "vika"},
	Day:                   []string{"d", "da", "dagur"},
	Hour:                  []string{"t", "tími"},
	Minute:                []string{"m", "min", "minuttur"},
	Second:                []string{"s", "sek", "sekund"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`í dag`},
		`0 hour ago`:   {`hendan tíman`},
		`0 minute ago`: {`hendan minuttin`},
		`0 month ago`:  {`henda mánaðin`},
		`0 second ago`: {`nú`},
		`0 week ago`:   {`hesu viku`},
		`0 year ago`:   {`í ár`},
		`1 day ago`:    {`í gjár`},
		`1 month ago`:  {`seinasta mánað`},
		`1 week ago`:   {`seinastu viku`},
		`1 year ago`:   {`í fjør`},
		`in 1 day`:     {`í morgin`},
		`in 1 month`:   {`næsta mánað`},
		`in 1 week`:    {`næstu viku`},
		`in 1 year`:    {`næsta ár`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) d síðan`),
			regexp.MustCompile(`(?i)(\d+) da síðan`),
			regexp.MustCompile(`(?i)(\d+) dagar síðan`),
			regexp.MustCompile(`(?i)(\d+) dagur síðan`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) t síðan`),
			regexp.MustCompile(`(?i)(\d+) tímar síðan`),
			regexp.MustCompile(`(?i)(\d+) tími síðan`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) m síðan`),
			regexp.MustCompile(`(?i)(\d+) min síðan`),
			regexp.MustCompile(`(?i)(\d+) minutt síðan`),
			regexp.MustCompile(`(?i)(\d+) minuttir síðan`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) mnð síðan`),
			regexp.MustCompile(`(?i)(\d+) mánað síðan`),
			regexp.MustCompile(`(?i)(\d+) mánaðir síðan`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) s síðan`),
			regexp.MustCompile(`(?i)(\d+) sek síðan`),
			regexp.MustCompile(`(?i)(\d+) sekund síðan`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) v síðan`),
			regexp.MustCompile(`(?i)(\d+) vi síðan`),
			regexp.MustCompile(`(?i)(\d+) vika síðan`),
			regexp.MustCompile(`(?i)(\d+) vikur síðan`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) ár síðan`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)um (\d+) d`),
			regexp.MustCompile(`(?i)um (\d+) da`),
			regexp.MustCompile(`(?i)um (\d+) dag`),
			regexp.MustCompile(`(?i)um (\d+) dagar`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)um (\d+) t`),
			regexp.MustCompile(`(?i)um (\d+) tíma`),
			regexp.MustCompile(`(?i)um (\d+) tímar`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)um (\d+) m`),
			regexp.MustCompile(`(?i)um (\d+) min`),
			regexp.MustCompile(`(?i)um (\d+) minutt`),
			regexp.MustCompile(`(?i)um (\d+) minuttir`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)um (\d+) mnð`),
			regexp.MustCompile(`(?i)um (\d+) mánað`),
			regexp.MustCompile(`(?i)um (\d+) mánaðir`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)um (\d+) s`),
			regexp.MustCompile(`(?i)um (\d+) sek`),
			regexp.MustCompile(`(?i)um (\d+) sekund`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)um (\d+) v`),
			regexp.MustCompile(`(?i)um (\d+) vi`),
			regexp.MustCompile(`(?i)um (\d+) viku`),
			regexp.MustCompile(`(?i)um (\d+) vikur`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)um (\d+) ár`),
		},
	},
}
