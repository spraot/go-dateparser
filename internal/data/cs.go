// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var cs_Locale = LocaleData{
	Name:                  "cs",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "přibližně", "v", "|", "，"},
	January:               []string{"led", "leden", "ledna"},
	February:              []string{"úno", "únor", "února"},
	March:                 []string{"bře", "březen", "března"},
	April:                 []string{"dub", "duben", "dubna"},
	May:                   []string{"kvě", "květen", "května"},
	June:                  []string{"čer", "červen", "června", "čvn"},
	July:                  []string{"července", "červenec", "črc", "čvc"},
	August:                []string{"srp", "srpen", "srpna"},
	September:             []string{"zář", "září"},
	October:               []string{"říj", "říjen", "října"},
	November:              []string{"lis", "listopad", "listopadu"},
	December:              []string{"pro", "prosince", "prosinec"},
	Monday:                []string{"po", "pon", "pondělí"},
	Tuesday:               []string{"út", "úte", "úterý"},
	Wednesday:             []string{"st", "stř", "středa", "středu"},
	Thursday:              []string{"čt", "čtv", "čtvrtek"},
	Friday:                []string{"pá", "pát", "pátek"},
	Saturday:              []string{"so", "sob", "sobota", "sobotu"},
	Sunday:                []string{"ne", "ned", "neděle", "neděli"},
	AM:                    []string{"dop"},
	PM:                    []string{"odp"},
	Year:                  []string{"r", "rok", "roky", "roků"},
	Month:                 []string{"měs", "měsíc", "měsíce", "měsíců"},
	Week:                  []string{"týd", "týden", "týdny", "týdnů"},
	Day:                   []string{"den", "dny", "dnů"},
	Hour:                  []string{"h", "hodin", "hodina", "hodinami", "hodinu", "hodiny"},
	Minute:                []string{"min", "minut", "minuta", "minuty"},
	Second:                []string{"s", "sekund", "sekunda", "sekundy", "vteřin", "vteřina", "vteřiny"},
	In:                    []string{"v", "ve"},
	Ago:                   []string{"před"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`dnes`},
		`0 hour ago`:   {`tuto hodinu`},
		`0 minute ago`: {`tuto minutu`},
		`0 month ago`:  {`tento měsíc`},
		`0 second ago`: {`nyní`},
		`0 week ago`:   {`tento týd`, `tento týden`},
		`0 year ago`:   {`tento rok`},
		`1 day ago`:    {`včera`},
		`1 month ago`:  {`minulý měsíc`},
		`1 week ago`:   {`minulý týd`, `minulý týden`},
		`1 year ago`:   {`minulý rok`},
		`2 day ago`:    {`předevčírem`},
		`in 1 day`:     {`zítra`},
		`in 1 month`:   {`příští měsíc`},
		`in 1 week`:    {`příští týd`, `příští týden`},
		`in 1 year`:    {`příští rok`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)před (\d+) dnem`),
			regexp.MustCompile(`(?i)před (\d+) dny`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)před (\d+) h`),
			regexp.MustCompile(`(?i)před (\d+) hodinami`),
			regexp.MustCompile(`(?i)před (\d+) hodinou`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)před (\d+) min`),
			regexp.MustCompile(`(?i)před (\d+) minutami`),
			regexp.MustCompile(`(?i)před (\d+) minutou`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)před (\d+) měs`),
			regexp.MustCompile(`(?i)před (\d+) měsícem`),
			regexp.MustCompile(`(?i)před (\d+) měsíci`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)před (\d+) s`),
			regexp.MustCompile(`(?i)před (\d+) sekundami`),
			regexp.MustCompile(`(?i)před (\d+) sekundou`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)před (\d+) týd`),
			regexp.MustCompile(`(?i)před (\d+) týdnem`),
			regexp.MustCompile(`(?i)před (\d+) týdny`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)před (\d+) l`),
			regexp.MustCompile(`(?i)před (\d+) lety`),
			regexp.MustCompile(`(?i)před (\d+) r`),
			regexp.MustCompile(`(?i)před (\d+) rokem`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)za (\d+) den`),
			regexp.MustCompile(`(?i)za (\d+) dní`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)za (\d+) h`),
			regexp.MustCompile(`(?i)za (\d+) hodin`),
			regexp.MustCompile(`(?i)za (\d+) hodinu`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)za (\d+) min`),
			regexp.MustCompile(`(?i)za (\d+) minut`),
			regexp.MustCompile(`(?i)za (\d+) minutu`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)za (\d+) měs`),
			regexp.MustCompile(`(?i)za (\d+) měsíc`),
			regexp.MustCompile(`(?i)za (\d+) měsíců`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)za (\d+) s`),
			regexp.MustCompile(`(?i)za (\d+) sekund`),
			regexp.MustCompile(`(?i)za (\d+) sekundu`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)za (\d+) týd`),
			regexp.MustCompile(`(?i)za (\d+) týden`),
			regexp.MustCompile(`(?i)za (\d+) týdnů`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)za (\d+) l`),
			regexp.MustCompile(`(?i)za (\d+) let`),
			regexp.MustCompile(`(?i)za (\d+) r`),
			regexp.MustCompile(`(?i)za (\d+) rok`),
		},
	},
}
