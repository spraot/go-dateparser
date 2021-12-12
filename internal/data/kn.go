// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var kn_Locale = LocaleData{
	Name:                  "kn",
	DateOrder:             "DMY",
	January:               []string{"ಜನ", "ಜನವರಿ"},
	February:              []string{"ಫೆಬ್ರ", "ಫೆಬ್ರವರಿ"},
	March:                 []string{"ಮಾರ್ಚ್"},
	April:                 []string{"ಏಪ್ರಿ", "ಏಪ್ರಿಲ್"},
	May:                   []string{"ಮೇ"},
	June:                  []string{"ಜೂನ್"},
	July:                  []string{"ಜುಲೈ"},
	August:                []string{"ಆಗ", "ಆಗಸ್ಟ್"},
	September:             []string{"ಸೆಪ್ಟೆಂ", "ಸೆಪ್ಟೆಂಬರ್"},
	October:               []string{"ಅಕ್ಟೋ", "ಅಕ್ಟೋಬರ್"},
	November:              []string{"ನವೆಂ", "ನವೆಂಬರ್"},
	December:              []string{"ಡಿಸೆಂ", "ಡಿಸೆಂಬರ್"},
	Monday:                []string{"ಸೋಮ", "ಸೋಮವಾರ"},
	Tuesday:               []string{"ಮಂಗಳ", "ಮಂಗಳವಾರ"},
	Wednesday:             []string{"ಬುಧ", "ಬುಧವಾರ"},
	Thursday:              []string{"ಗುರು", "ಗುರುವಾರ"},
	Friday:                []string{"ಶುಕ್ರ", "ಶುಕ್ರವಾರ"},
	Saturday:              []string{"ಶನಿ", "ಶನಿವಾರ"},
	Sunday:                []string{"ಭಾನು", "ಭಾನುವಾರ"},
	AM:                    []string{"ಪೂರ್ವಾಹ್ನ"},
	PM:                    []string{"ಅಪರಾಹ್ನ"},
	Year:                  []string{"ವರ್ಷ"},
	Month:                 []string{"ತಿಂಗಳು"},
	Week:                  []string{"ವಾರ"},
	Day:                   []string{"ದಿನ"},
	Hour:                  []string{"ಗಂಟೆ"},
	Minute:                []string{"ನಿಮಿಷ"},
	Second:                []string{"ಸೆಕೆಂಡ್"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`ಇಂದು`},
		`0 hour ago`:   {`ಈ ಗಂಟೆ`},
		`0 minute ago`: {`ಈ ನಿಮಿಷ`},
		`0 month ago`:  {`ಈ ತಿಂಗಳು`},
		`0 second ago`: {`ಈಗ`},
		`0 week ago`:   {`ಈ ವಾರ`},
		`0 year ago`:   {`ಈ ವರ್ಷ`},
		`1 day ago`:    {`ನಿನ್ನೆ`},
		`1 month ago`:  {`ಕಳೆದ ತಿಂಗಳು`},
		`1 week ago`:   {`ಕಳೆದ ವಾರ`},
		`1 year ago`:   {`ಕಳೆದ ವರ್ಷ`, `ಹಿಂದಿನ ವರ್ಷ`},
		`in 1 day`:     {`ನಾಳೆ`},
		`in 1 month`:   {`ಮುಂದಿನ ತಿಂಗಳು`},
		`in 1 week`:    {`ಮುಂದಿನ ವಾರ`},
		`in 1 year`:    {`ಮುಂದಿನ ವರ್ಷ`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) ದಿನಗಳ ಹಿಂದೆ`),
			regexp.MustCompile(`(?i)(\d+) ದಿನದ ಹಿಂದೆ`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) ಗಂಟೆ ಹಿಂದೆ`),
			regexp.MustCompile(`(?i)(\d+) ಗಂಟೆಗಳ ಹಿಂದೆ`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) ನಿಮಿಷಗಳ ಹಿಂದೆ`),
			regexp.MustCompile(`(?i)(\d+) ನಿಮಿಷದ ಹಿಂದೆ`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) ತಿಂಗಳ ಹಿಂದೆ`),
			regexp.MustCompile(`(?i)(\d+) ತಿಂಗಳು ಹಿಂದೆ`),
			regexp.MustCompile(`(?i)(\d+) ತಿಂಗಳುಗಳ ಹಿಂದೆ`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) ಸೆಕೆಂಡುಗಳ ಹಿಂದೆ`),
			regexp.MustCompile(`(?i)(\d+) ಸೆಕೆಂಡ್ ಹಿಂದೆ`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) ವಾರಗಳ ಹಿಂದೆ`),
			regexp.MustCompile(`(?i)(\d+) ವಾರದ ಹಿಂದೆ`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) ವರ್ಷಗಳ ಹಿಂದೆ`),
			regexp.MustCompile(`(?i)(\d+) ವರ್ಷದ ಹಿಂದೆ`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)(\d+) ದಿನಗಳಲ್ಲಿ`),
			regexp.MustCompile(`(?i)(\d+) ದಿನದಲ್ಲಿ`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)(\d+) ಗಂಟೆಗಳಲ್ಲಿ`),
			regexp.MustCompile(`(?i)(\d+) ಗಂಟೆಯಲ್ಲಿ`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)(\d+) ನಿಮಿಷಗಳಲ್ಲಿ`),
			regexp.MustCompile(`(?i)(\d+) ನಿಮಿಷದಲ್ಲಿ`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)(\d+) ತಿಂಗಳಲ್ಲಿ`),
			regexp.MustCompile(`(?i)(\d+) ತಿಂಗಳುಗಳಲ್ಲಿ`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)(\d+) ಸೆಕೆಂಡ್‌ಗಳಲ್ಲಿ`),
			regexp.MustCompile(`(?i)(\d+) ಸೆಕೆಂಡ್‌ನಲ್ಲಿ`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)(\d+) ವಾರಗಳಲ್ಲಿ`),
			regexp.MustCompile(`(?i)(\d+) ವಾರದಲ್ಲಿ`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)(\d+) ವರ್ಷಗಳಲ್ಲಿ`),
			regexp.MustCompile(`(?i)(\d+) ವರ್ಷದಲ್ಲಿ`),
		},
	},
}
