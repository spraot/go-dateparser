// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var te_Locale = merge(nil, LocaleData{
	Name:      "te",
	DateOrder: "DMY",
	Charset:   []rune(`cgtuzంఅఆఈఏకగచజటడతదధనపఫబమరలళవశషసాిుూెేైో్`),
	Translations: map[string]string{
		"శుకరవరం": "friday",
		"గురువరం": "thursday",
		"మంగళవరం": "tuesday",
		"సంవతసరం": "year",
		"సపటంబర":  "september",
		"బుధవరం":  "wednesday",
		"ఆగసటు":   "august",
		"డసంబర":   "december",
		"ఫబరవర":   "february",
		"నమషము":   "minute",
		"సమవరం":   "monday",
		"నవంబర":   "november",
		"అకటబర":   "october",
		"శనవరం":   "saturday",
		"ఆదవరం":   "sunday",
		"ఏపరల":    "april",
		"శుకర":    "friday",
		"జనవర":    "january",
		"సకను":    "second",
		"సపటం":    "september",
		"గురు":    "thursday",
		"మంగళ":    "tuesday",
		"వరము":    "week",
		"ఏపర":     "april",
		"దనం":     "day",
		"డసం":     "december",
		"ఫబర":     "february",
		"gmt":     "gmt",
		"గంట":     "hour",
		"జుల":     "july",
		"జూన":     "june",
		"మరచ":     "march",
		"నవం":     "november",
		"అకట":     "october",
		"utc":     "utc",
		"బుధ":     "wednesday",
		"am":      "am",
		"ఆగ":      "august",
		"గం":      "hour",
		"జన":      "january",
		"నమ":      "minute",
		"సమ":      "monday",
		"నల":      "month",
		"pm":      "pm",
		"శన":      "saturday",
		"సక":      "second",
		"ఆద":      "sunday",
		"సం":      "year",
		"'":       "",
		",":       "",
		";":       "",
		"@":       "",
		"[":       "",
		"]":       "",
		"|":       "",
		" ":       " ",
		"+":       "+",
		"-":       "-",
		".":       ".",
		"/":       "/",
		":":       ":",
		"ద":       "day",
		"మ":       "may",
		"న":       "minute",
		"వ":       "week",
		"z":       "z",
	},
	RelativeType: map[string]string{
		"తదుపర సంవతసరం": "in 1 year",
		"గత సంవతసరం":    "1 year ago",
		"ఈ సంవతసరం":     "0 year ago",
		"తదుపర వరం":     "in 1 week",
		"తదుపర నల":      "in 1 month",
		"పరసతుతం":       "0 second ago",
		"ఈ నమషం":        "0 minute ago",
		"గత వరం":        "1 week ago",
		"ఈ రజు":         "0 day ago",
		"ఈ గంట":         "0 hour ago",
		"ఈ వరం":         "0 week ago",
		"గత నల":         "1 month ago",
		"ఈ నల":          "0 month ago",
		"ననన":           "1 day ago",
		"రపు":           "in 1 day",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) సంవతసరం కరతం`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) సంవతసరల కరతం`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) రజుల కరతం`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) గంటల కరతం`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) నమషం కరతం`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) నమషల కరతం`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) సకనల కరతం`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) సకను కరతం`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) రజు కరతం`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) గంట కరతం`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) నలల కరతం`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) వరం కరతం`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) వరల కరతం`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) సంవతసరంల`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) సంవతసరలల`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) గం కరతం`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) నమ కరతం`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) నల కరతం`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) సక కరతం`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) సం కరతం`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) రజులల`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) గంటలల`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) నమషంల`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) నమషలల`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) సకనలల`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) సకనుల`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) రజుల`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) గంటల`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) నలలల`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) సక ల`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) వరంల`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) వరలల`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) గంల`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) నమల`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) నలల`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) సంల`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ సంవతసరం కరతం|\d+ సంవతసరల కరతం|\d+ గంటల కరతం|\d+ నమషం కరతం|\d+ నమషల కరతం|\d+ రజుల కరతం|\d+ సకనల కరతం|\d+ సకను కరతం|\d+ గంట కరతం|\d+ నలల కరతం|\d+ రజు కరతం|\d+ వరం కరతం|\d+ వరల కరతం|\d+ సంవతసరంల|\d+ సంవతసరలల|\d+ గం కరతం|\d+ నమ కరతం|\d+ నల కరతం|\d+ సం కరతం|\d+ సక కరతం|\d+ గంటలల|\d+ నమషంల|\d+ నమషలల|\d+ రజులల|\d+ సకనలల|\d+ సకనుల|\d+ గంటల|\d+ నలలల|\d+ రజుల|\d+ వరంల|\d+ వరలల|\d+ సక ల|\d+ గంల|\d+ నమల|\d+ నలల|\d+ సంల)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ సంవతసరం కరతం|\d+ సంవతసరల కరతం|\d+ గంటల కరతం|\d+ నమషం కరతం|\d+ నమషల కరతం|\d+ రజుల కరతం|\d+ సకనల కరతం|\d+ సకను కరతం|\d+ గంట కరతం|\d+ నలల కరతం|\d+ రజు కరతం|\d+ వరం కరతం|\d+ వరల కరతం|\d+ సంవతసరంల|\d+ సంవతసరలల|\d+ గం కరతం|\d+ నమ కరతం|\d+ నల కరతం|\d+ సం కరతం|\d+ సక కరతం|\d+ గంటలల|\d+ నమషంల|\d+ నమషలల|\d+ రజులల|\d+ సకనలల|\d+ సకనుల|\d+ గంటల|\d+ నలలల|\d+ రజుల|\d+ వరంల|\d+ వరలల|\d+ సక ల|\d+ గంల|\d+ నమల|\d+ నలల|\d+ సంల)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(తదుపర సంవతసరం|గత సంవతసరం|ఈ సంవతసరం|తదుపర వరం|తదుపర నల|గురువరం|పరసతుతం|మంగళవరం|శుకరవరం|సంవతసరం|ఈ నమషం|గత వరం|బుధవరం|సపటంబర|అకటబర|ఆగసటు|ఆదవరం|ఈ గంట|ఈ రజు|ఈ వరం|గత నల|డసంబర|నమషము|నవంబర|ఫబరవర|శనవరం|సమవరం|ఈ నల|ఏపరల|గురు|జనవర|మంగళ|వరము|శుకర|సకను|సపటం|gmt|utc|అకట|ఏపర|గంట|జుల|జూన|డసం|దనం|ననన|నవం|ఫబర|బుధ|మరచ|రపు|\+|\.|\[|\]|\||am|pm|ఆగ|ఆద|గం|జన|నమ|నల|శన|సం|సక|సమ| |'|,|-|/|:|;|@|z|ద|న|మ|వ)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
