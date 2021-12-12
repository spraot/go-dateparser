// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var mr_Locale = LocaleData{
	Name:                  "mr",
	DateOrder:             "DMY",
	January:               []string{"जाने", "जानेवारी"},
	February:              []string{"फेब्रु", "फेब्रुवारी"},
	March:                 []string{"मार्च"},
	April:                 []string{"एप्रि", "एप्रिल"},
	May:                   []string{"मे"},
	June:                  []string{"जून"},
	July:                  []string{"जुलै"},
	August:                []string{"ऑग", "ऑगस्ट"},
	September:             []string{"सप्टें", "सप्टेंबर"},
	October:               []string{"ऑक्टो", "ऑक्टोबर"},
	November:              []string{"नोव्हें", "नोव्हेंबर"},
	December:              []string{"डिसें", "डिसेंबर"},
	Monday:                []string{"सोम", "सोमवार"},
	Tuesday:               []string{"मंगळ", "मंगळवार"},
	Wednesday:             []string{"बुध", "बुधवार"},
	Thursday:              []string{"गुरु", "गुरुवार"},
	Friday:                []string{"शुक्र", "शुक्रवार"},
	Saturday:              []string{"शनि", "शनिवार"},
	Sunday:                []string{"रवि", "रविवार"},
	AM:                    []string{"मपू"},
	PM:                    []string{"मउ"},
	Year:                  []string{"वर्ष"},
	Month:                 []string{"महिना"},
	Week:                  []string{"आठवडा"},
	Day:                   []string{"दिवस"},
	Hour:                  []string{"तास"},
	Minute:                []string{"मि", "मिनिट"},
	Second:                []string{"से", "सेकंद"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`आज`},
		`0 hour ago`:   {`तासात`},
		`0 minute ago`: {`या मिनिटात`},
		`0 month ago`:  {`हा महिना`},
		`0 second ago`: {`आत्ता`},
		`0 week ago`:   {`हा आठवडा`},
		`0 year ago`:   {`हे वर्ष`},
		`1 day ago`:    {`काल`},
		`1 month ago`:  {`मागील महिना`},
		`1 week ago`:   {`मागील आठवडा`},
		`1 year ago`:   {`मागील वर्ष`},
		`in 1 day`:     {`उद्या`},
		`in 1 month`:   {`पुढील महिना`},
		`in 1 week`:    {`पुढील आठवडा`},
		`in 1 year`:    {`पुढील वर्ष`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) दिवसांपूर्वी`),
			regexp.MustCompile(`(?i)(\d+) दिवसापूर्वी`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) तासांपूर्वी`),
			regexp.MustCompile(`(?i)(\d+) तासापूर्वी`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) मिनि पूर्वी`),
			regexp.MustCompile(`(?i)(\d+) मिनिटांपूर्वी`),
			regexp.MustCompile(`(?i)(\d+) मिनिटापूर्वी`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) महिन्यांपूर्वी`),
			regexp.MustCompile(`(?i)(\d+) महिन्यापूर्वी`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) से पूर्वी`),
			regexp.MustCompile(`(?i)(\d+) सेकंदांपूर्वी`),
			regexp.MustCompile(`(?i)(\d+) सेकंदापूर्वी`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) आठवड्यांपूर्वी`),
			regexp.MustCompile(`(?i)(\d+) आठवड्यापूर्वी`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) वर्षांपूर्वी`),
			regexp.MustCompile(`(?i)(\d+) वर्षापूर्वी`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)(\d+) दिवसांमध्ये`),
			regexp.MustCompile(`(?i)(\d+) दिवसामध्ये`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)(\d+) तासांमध्ये`),
			regexp.MustCompile(`(?i)(\d+) तासामध्ये`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)(\d+) मिनि मध्ये`),
			regexp.MustCompile(`(?i)(\d+) मिनिटांमध्ये`),
			regexp.MustCompile(`(?i)(\d+) मिनिटामध्ये`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)(\d+) महिन्यांमध्ये`),
			regexp.MustCompile(`(?i)(\d+) महिन्यामध्ये`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)(\d+) से मध्ये`),
			regexp.MustCompile(`(?i)(\d+) सेकंदांमध्ये`),
			regexp.MustCompile(`(?i)(\d+) सेकंदामध्ये`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)(\d+) आठवड्यांमध्ये`),
			regexp.MustCompile(`(?i)(\d+) आठवड्यामध्ये`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)(\d+) वर्षांमध्ये`),
			regexp.MustCompile(`(?i)(\d+) वर्षामध्ये`),
		},
	},
}
