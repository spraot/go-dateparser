// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var el_Locale = merge(nil, LocaleData{
	Name:      "el",
	DateOrder: "DMY",
	Charset:   []rune(`cgtuzΐάέήίαβγδεηθικλμνοπρςστυφχωόύώ`),
	Translations: map[string]string{
		"δευτερολεπτο": "second",
		"φεβρουαριος":  "february",
		"φεβρουαριου":  "february",
		"σεπτεμβριος":  "september",
		"σεπτεμβριου":  "september",
		"δεκεμβριος":   "december",
		"δεκεμβριου":   "december",
		"ιανουαριος":   "january",
		"ιανουαριου":   "january",
		"αυγουστος":    "august",
		"αυγουστου":    "august",
		"παρασκευη":    "friday",
		"νοεμβριος":    "november",
		"νοεμβριου":    "november",
		"οκτωβριος":    "october",
		"οκτωβριου":    "october",
		"απριλιος":     "april",
		"απριλιου":     "april",
		"εβδομαδα":     "week",
		"ιουλιος":      "july",
		"ιουλιου":      "july",
		"ιουνιος":      "june",
		"ιουνιου":      "june",
		"μαρτιος":      "march",
		"μαρτιου":      "march",
		"δευτερα":      "monday",
		"σαββατο":      "saturday",
		"κυριακη":      "sunday",
		"τεταρτη":      "wednesday",
		"πεμπτη":       "thursday",
		"ημερα":        "day",
		"μαιος":        "may",
		"μαιου":        "may",
		"λεπτο":        "minute",
		"μηνας":        "month",
		"τριτη":        "tuesday",
		"ιουλ":         "july",
		"ιουν":         "june",
		"δευτ":         "second",
		"ετος":         "year",
		"απρ":          "april",
		"αυγ":          "august",
		"δεκ":          "december",
		"φεβ":          "february",
		"παρ":          "friday",
		"gmt":          "gmt",
		"ωρα":          "hour",
		"ιαν":          "january",
		"μαρ":          "march",
		"μαι":          "may",
		"λεπ":          "minute",
		"δευ":          "monday",
		"μην":          "month",
		"νοε":          "november",
		"οκτ":          "october",
		"σαβ":          "saturday",
		"σεπ":          "september",
		"κυρ":          "sunday",
		"πεμ":          "thursday",
		"τρι":          "tuesday",
		"utc":          "utc",
		"τετ":          "wednesday",
		"εβδ":          "week",
		"am":           "am",
		"πμ":           "am",
		"ωρ":           "hour",
		"pm":           "pm",
		"μμ":           "pm",
		"ετ":           "year",
		"'":            "",
		",":            "",
		";":            "",
		"@":            "",
		"[":            "",
		"]":            "",
		"|":            "",
		" ":            " ",
		"+":            "+",
		"-":            "-",
		".":            ".",
		"/":            "/",
		":":            ":",
		"ω":            "hour",
		"λ":            "minute",
		"δ":            "second",
		"z":            "z",
	},
	RelativeType: map[string]string{
		"προηγουμενη εβδομαδα": "1 week ago",
		"αυτην την εβδομαδα":   "0 week ago",
		"προηγουμενος μηνας":   "1 month ago",
		"επομενη εβδομαδα":     "in 1 week",
		"επομενος μηνας":       "in 1 month",
		"αυτην την ωρα":        "0 hour ago",
		"αυτο το λεπτο":        "0 minute ago",
		"τρεχων μηνας":         "0 month ago",
		"επομενο ετος":         "in 1 year",
		"σημερα":               "0 day ago",
		"φετος":                "0 year ago",
		"περσι":                "1 year ago",
		"αυριο":                "in 1 day",
		"τωρα":                 "0 second ago",
		"χθες":                 "1 day ago",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)πριν απο (\d+) δευτερολεπτα`), "$1 second ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) δευτερολεπτο`), "$1 second ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) εβδομαδες`), "$1 week ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) εβδομαδα`), "$1 week ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) ημερες`), "$1 day ago"},
		{regexp.MustCompile(`(?i)σε (\d+) δευτερολεπτα`), "in $1 second"},
		{regexp.MustCompile(`(?i)σε (\d+) δευτερολεπτο`), "in $1 second"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) ημερα`), "$1 day ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) λεπτα`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) λεπτο`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) μηνες`), "$1 month ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) ωρες`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) μηνα`), "$1 month ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) δευτ`), "$1 second ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) ετος`), "$1 year ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) ωρα`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) λεπ`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) εβδ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) ετη`), "$1 year ago"},
		{regexp.MustCompile(`(?i)σε (\d+) εβδομαδες`), "in $1 week"},
		{regexp.MustCompile(`(?i)πριν απο (\d+) ωρ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)σε (\d+) εβδομαδα`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) ετος πριν`), "$1 year ago"},
		{regexp.MustCompile(`(?i)σε (\d+) ημερες`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) εβδ πριν`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) ετη πριν`), "$1 year ago"},
		{regexp.MustCompile(`(?i)σε (\d+) ημερα`), "in $1 day"},
		{regexp.MustCompile(`(?i)σε (\d+) λεπτα`), "in $1 minute"},
		{regexp.MustCompile(`(?i)σε (\d+) λεπτο`), "in $1 minute"},
		{regexp.MustCompile(`(?i)σε (\d+) μηνες`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) ημ πριν`), "$1 day ago"},
		{regexp.MustCompile(`(?i)σε (\d+) ωρες`), "in $1 hour"},
		{regexp.MustCompile(`(?i)σε (\d+) μηνα`), "in $1 month"},
		{regexp.MustCompile(`(?i)σε (\d+) δευτ`), "in $1 second"},
		{regexp.MustCompile(`(?i)σε (\d+) ετος`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) ω πριν`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) λ πριν`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) μ πριν`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) δ πριν`), "$1 second ago"},
		{regexp.MustCompile(`(?i)σε (\d+) ωρα`), "in $1 hour"},
		{regexp.MustCompile(`(?i)σε (\d+) λεπ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)σε (\d+) εβδ`), "in $1 week"},
		{regexp.MustCompile(`(?i)σε (\d+) ετη`), "in $1 year"},
		{regexp.MustCompile(`(?i)σε (\d+) ημ`), "in $1 day"},
		{regexp.MustCompile(`(?i)σε (\d+) ωρ`), "in $1 hour"},
		{regexp.MustCompile(`(?i)σε (\d+) ω`), "in $1 hour"},
		{regexp.MustCompile(`(?i)σε (\d+) λ`), "in $1 minute"},
		{regexp.MustCompile(`(?i)σε (\d+) μ`), "in $1 month"},
		{regexp.MustCompile(`(?i)σε (\d+) δ`), "in $1 second"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(πριν απο \d+ δευτερολεπτα|πριν απο \d+ δευτερολεπτο|πριν απο \d+ εβδομαδες|πριν απο \d+ εβδομαδα|πριν απο \d+ ημερες|σε \d+ δευτερολεπτα|σε \d+ δευτερολεπτο|πριν απο \d+ ημερα|πριν απο \d+ λεπτα|πριν απο \d+ λεπτο|πριν απο \d+ μηνες|πριν απο \d+ δευτ|πριν απο \d+ ετος|πριν απο \d+ μηνα|πριν απο \d+ ωρες|πριν απο \d+ εβδ|πριν απο \d+ ετη|πριν απο \d+ λεπ|πριν απο \d+ ωρα|σε \d+ εβδομαδες|πριν απο \d+ ωρ|σε \d+ εβδομαδα|\d+ ετος πριν|σε \d+ ημερες|\d+ εβδ πριν|\d+ ετη πριν|σε \d+ ημερα|σε \d+ λεπτα|σε \d+ λεπτο|σε \d+ μηνες|\d+ ημ πριν|σε \d+ δευτ|σε \d+ ετος|σε \d+ μηνα|σε \d+ ωρες|\d+ δ πριν|\d+ λ πριν|\d+ μ πριν|\d+ ω πριν|σε \d+ εβδ|σε \d+ ετη|σε \d+ λεπ|σε \d+ ωρα|σε \d+ ημ|σε \d+ ωρ|σε \d+ δ|σε \d+ λ|σε \d+ μ|σε \d+ ω)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(πριν απο \d+ δευτερολεπτα|πριν απο \d+ δευτερολεπτο|πριν απο \d+ εβδομαδες|πριν απο \d+ εβδομαδα|πριν απο \d+ ημερες|σε \d+ δευτερολεπτα|σε \d+ δευτερολεπτο|πριν απο \d+ ημερα|πριν απο \d+ λεπτα|πριν απο \d+ λεπτο|πριν απο \d+ μηνες|πριν απο \d+ δευτ|πριν απο \d+ ετος|πριν απο \d+ μηνα|πριν απο \d+ ωρες|πριν απο \d+ εβδ|πριν απο \d+ ετη|πριν απο \d+ λεπ|πριν απο \d+ ωρα|σε \d+ εβδομαδες|πριν απο \d+ ωρ|σε \d+ εβδομαδα|\d+ ετος πριν|σε \d+ ημερες|\d+ εβδ πριν|\d+ ετη πριν|σε \d+ ημερα|σε \d+ λεπτα|σε \d+ λεπτο|σε \d+ μηνες|\d+ ημ πριν|σε \d+ δευτ|σε \d+ ετος|σε \d+ μηνα|σε \d+ ωρες|\d+ δ πριν|\d+ λ πριν|\d+ μ πριν|\d+ ω πριν|σε \d+ εβδ|σε \d+ ετη|σε \d+ λεπ|σε \d+ ωρα|σε \d+ ημ|σε \d+ ωρ|σε \d+ δ|σε \d+ λ|σε \d+ μ|σε \d+ ω)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(προηγουμενη εβδομαδα|αυτην την εβδομαδα|προηγουμενος μηνας|επομενη εβδομαδα|επομενος μηνας|αυτην την ωρα|αυτο το λεπτο|δευτερολεπτο|επομενο ετος|τρεχων μηνας|σεπτεμβριος|σεπτεμβριου|φεβρουαριος|φεβρουαριου|δεκεμβριος|δεκεμβριου|ιανουαριος|ιανουαριου|αυγουστος|αυγουστου|νοεμβριος|νοεμβριου|οκτωβριος|οκτωβριου|παρασκευη|απριλιος|απριλιου|εβδομαδα|δευτερα|ιουλιος|ιουλιου|ιουνιος|ιουνιου|κυριακη|μαρτιος|μαρτιου|σαββατο|τεταρτη|πεμπτη|σημερα|αυριο|ημερα|λεπτο|μαιος|μαιου|μηνας|περσι|τριτη|φετος|δευτ|ετος|ιουλ|ιουν|τωρα|χθες|gmt|utc|απρ|αυγ|δεκ|δευ|εβδ|ιαν|κυρ|λεπ|μαι|μαρ|μην|νοε|οκτ|παρ|πεμ|σαβ|σεπ|τετ|τρι|φεβ|ωρα|\+|\.|\[|\]|\||am|pm|ετ|μμ|πμ|ωρ| |'|,|-|/|:|;|@|z|δ|λ|ω)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})

var el_CY_Locale = merge(&el_Locale, LocaleData{
	Name:            "el-CY",
	DateOrder:       "DMY",
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(πριν απο \d+ δευτερολεπτα|πριν απο \d+ δευτερολεπτο|πριν απο \d+ εβδομαδες|πριν απο \d+ εβδομαδα|πριν απο \d+ ημερες|σε \d+ δευτερολεπτα|σε \d+ δευτερολεπτο|πριν απο \d+ ημερα|πριν απο \d+ λεπτα|πριν απο \d+ λεπτο|πριν απο \d+ μηνες|πριν απο \d+ δευτ|πριν απο \d+ ετος|πριν απο \d+ μηνα|πριν απο \d+ ωρες|πριν απο \d+ εβδ|πριν απο \d+ ετη|πριν απο \d+ λεπ|πριν απο \d+ ωρα|σε \d+ εβδομαδες|πριν απο \d+ ωρ|σε \d+ εβδομαδα|\d+ ετος πριν|σε \d+ ημερες|\d+ εβδ πριν|\d+ ετη πριν|σε \d+ ημερα|σε \d+ λεπτα|σε \d+ λεπτο|σε \d+ μηνες|\d+ ημ πριν|σε \d+ δευτ|σε \d+ ετος|σε \d+ μηνα|σε \d+ ωρες|\d+ δ πριν|\d+ λ πριν|\d+ μ πριν|\d+ ω πριν|σε \d+ εβδ|σε \d+ ετη|σε \d+ λεπ|σε \d+ ωρα|σε \d+ ημ|σε \d+ ωρ|σε \d+ δ|σε \d+ λ|σε \d+ μ|σε \d+ ω)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(πριν απο \d+ δευτερολεπτα|πριν απο \d+ δευτερολεπτο|πριν απο \d+ εβδομαδες|πριν απο \d+ εβδομαδα|πριν απο \d+ ημερες|σε \d+ δευτερολεπτα|σε \d+ δευτερολεπτο|πριν απο \d+ ημερα|πριν απο \d+ λεπτα|πριν απο \d+ λεπτο|πριν απο \d+ μηνες|πριν απο \d+ δευτ|πριν απο \d+ ετος|πριν απο \d+ μηνα|πριν απο \d+ ωρες|πριν απο \d+ εβδ|πριν απο \d+ ετη|πριν απο \d+ λεπ|πριν απο \d+ ωρα|σε \d+ εβδομαδες|πριν απο \d+ ωρ|σε \d+ εβδομαδα|\d+ ετος πριν|σε \d+ ημερες|\d+ εβδ πριν|\d+ ετη πριν|σε \d+ ημερα|σε \d+ λεπτα|σε \d+ λεπτο|σε \d+ μηνες|\d+ ημ πριν|σε \d+ δευτ|σε \d+ ετος|σε \d+ μηνα|σε \d+ ωρες|\d+ δ πριν|\d+ λ πριν|\d+ μ πριν|\d+ ω πριν|σε \d+ εβδ|σε \d+ ετη|σε \d+ λεπ|σε \d+ ωρα|σε \d+ ημ|σε \d+ ωρ|σε \d+ δ|σε \d+ λ|σε \d+ μ|σε \d+ ω)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(προηγουμενη εβδομαδα|αυτην την εβδομαδα|προηγουμενος μηνας|επομενη εβδομαδα|επομενος μηνας|αυτην την ωρα|αυτο το λεπτο|δευτερολεπτο|επομενο ετος|τρεχων μηνας|σεπτεμβριος|σεπτεμβριου|φεβρουαριος|φεβρουαριου|δεκεμβριος|δεκεμβριου|ιανουαριος|ιανουαριου|αυγουστος|αυγουστου|νοεμβριος|νοεμβριου|οκτωβριος|οκτωβριου|παρασκευη|απριλιος|απριλιου|εβδομαδα|δευτερα|ιουλιος|ιουλιου|ιουνιος|ιουνιου|κυριακη|μαρτιος|μαρτιου|σαββατο|τεταρτη|πεμπτη|σημερα|αυριο|ημερα|λεπτο|μαιος|μαιου|μηνας|περσι|τριτη|φετος|δευτ|ετος|ιουλ|ιουν|τωρα|χθες|gmt|utc|απρ|αυγ|δεκ|δευ|εβδ|ιαν|κυρ|λεπ|μαι|μαρ|μην|νοε|οκτ|παρ|πεμ|σαβ|σεπ|τετ|τρι|φεβ|ωρα|\+|\.|\[|\]|\||am|pm|ετ|μμ|πμ|ωρ| |'|,|-|/|:|;|@|z|δ|λ|ω)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
