// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var to_Locale = LocaleData{
	Name:                  "to",
	DateOrder:             "DMY",
	January:               []string{"sān", "sānuali"},
	February:              []string{"fēp", "fēpueli"},
	March:                 []string{"ma'a", "ma'asi"},
	April:                 []string{"'epe", "'epeleli"},
	May:                   []string{"mē"},
	June:                  []string{"sun", "sune"},
	July:                  []string{"siu", "siulai"},
	August:                []string{"'aok", "'aokosi"},
	September:             []string{"sep", "sepitema"},
	October:               []string{"'oka", "'okatopa"},
	November:              []string{"nōv", "nōvema"},
	December:              []string{"tīs", "tīsema"},
	Monday:                []string{"mōn", "mōnite"},
	Tuesday:               []string{"tūs", "tūsite"},
	Wednesday:             []string{"pul", "pulelulu"},
	Thursday:              []string{"tu'a", "tu'apulelulu"},
	Friday:                []string{"fal", "falaite"},
	Saturday:              []string{"tok", "tokonaki"},
	Sunday:                []string{"sāp", "sāpate"},
	AM:                    []string{"am", "hengihengi", "hh"},
	PM:                    []string{"ea", "efiafi", "pm"},
	Year:                  []string{"ta'u"},
	Month:                 []string{"māhina"},
	Week:                  []string{"uike"},
	Day:                   []string{"'aho"},
	Hour:                  []string{"houa"},
	Minute:                []string{"miniti"},
	Second:                []string{"sekoni"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`'ahó ni`},
		`0 hour ago`:   {`this hour`},
		`0 minute ago`: {`this minute`},
		`0 month ago`:  {`māhiná ni`},
		`0 second ago`: {`taimí ni`},
		`0 week ago`:   {`uiké ni`},
		`0 year ago`:   {`ta'ú ni`},
		`1 day ago`:    {`'aneafi`},
		`1 month ago`:  {`māhina kuo'osi`},
		`1 week ago`:   {`uike kuo'osi`},
		`1 year ago`:   {`ta'u kuo'osi`},
		`in 1 day`:     {`'apongipongi`},
		`in 1 month`:   {`māhina kaha'u`},
		`in 1 week`:    {`uike kaha'u`},
		`in 1 year`:    {`ta'u kaha'u`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)'aho 'e (\d+) kuo'osi`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)houa 'e (\d+) kuo'osi`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)miniti 'e (\d+) kuo'osi`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)māhina 'e (\d+) kuo'osi`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)sekoni 'e (\d+) kuo'osi`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)uike 'e (\d+) kuo'osi`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)ta'u 'e (\d+) kuo'osi`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)'i he 'aho 'e (\d+)`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)'i he houa 'e (\d+)`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)'i he miniti 'e (\d+)`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)'i he māhina 'e (\d+)`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)'i he sekoni 'e (\d+)`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)'i he uike 'e (\d+)`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)'i he ta'u 'e (\d+)`),
		},
	},
}
