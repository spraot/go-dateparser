// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var ka_Locale = merge(nil, LocaleData{
	Name:      "ka",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "|", "და", "დაახლოებით", "ზე", "ის"},
	Simplifications: []ReplacementData{
		{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)ერთ(\z|[^\pL\pM\d]|_)`), "${1}1${2}"},
	},
	Translations: map[string]string{
		"შუადღ შემდეგ": "pm",
		"დაახლოებით":   "",
		"სექტემბერი":   "september",
		"დეკემბერი":    "december",
		"თებერვალი":    "february",
		"პარასკევი":    "friday",
		"ოქტომბერი":    "october",
		"ხუთშაბათი":    "thursday",
		"სამშაბათი":    "tuesday",
		"ოთხშაბათი":    "wednesday",
		"ორშაბათი":     "monday",
		"ნოემბერი":     "november",
		"აგვისტო":      "august",
		"დღეიდან":      "in",
		"იანვარი":      "january",
		"აპრილი":       "april",
		"ივლისი":       "july",
		"ივნისი":       "june",
		"შაბათი":       "saturday",
		"საათი":        "hour",
		"მარტი":        "march",
		"მაისი":        "may",
		"კვირა":        "sunday",
		"წუთი":         "minute",
		"წამი":         "second",
		"წელი":         "year",
		"წლის":         "year",
		"წინ":          "ago",
		"აპრ":          "april",
		"აგვ":          "august",
		"დღე":          "day",
		"დეკ":          "december",
		"თებ":          "february",
		"პარ":          "friday",
		"gmt":          "gmt",
		"იან":          "january",
		"ივლ":          "july",
		"ივნ":          "june",
		"მარ":          "march",
		"მაი":          "may",
		"ორშ":          "monday",
		"თვე":          "month",
		"ნოე":          "november",
		"ოქტ":          "october",
		"შაბ":          "saturday",
		"სექ":          "september",
		"კვი":          "sunday",
		"ხუთ":          "thursday",
		"სამ":          "tuesday",
		"utc":          "utc",
		"ოთხ":          "wednesday",
		"და":           "",
		"ზე":           "",
		"ის":           "",
		"am":           "am",
		"სთ":           "hour",
		"წთ":           "minute",
		"pm":           "pm",
		"წმ":           "second",
		"კვ":           "week",
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
		"წ":            "year",
		"z":            "z",
	},
	RelativeType: map[string]string{
		"მომავალ კვირაში": "in 1 week",
		"გასულ კვირაში":   "1 week ago",
		"მომავალ თვეს":    "in 1 month",
		"მომავალ წელს":    "in 1 year",
		"ამ კვირაში":      "0 week ago",
		"გასულ თვეს":      "1 month ago",
		"გასულ წელს":      "1 year ago",
		"ამ საათში":       "0 hour ago",
		"ამ წუთში":        "0 minute ago",
		"ამ თვეში":        "0 month ago",
		"ამ წელს":         "0 year ago",
		"გუშინ":           "1 day ago",
		"დღეს":            "0 day ago",
		"ახლა":            "0 second ago",
		"ხვალ":            "in 1 day",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) საათის წინ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) კვირის წინ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) წუთის წინ`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) წამის წინ`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) წელიწადში`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) დღის წინ`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) თვის წინ`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) წლის წინ`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) კვირაში`), "in $1 week"},
		{regexp.MustCompile(`(?i)(\d+) სთ წინ`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) წთ წინ`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) წმ წინ`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) კვ წინ`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) საათში`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) დღეში`), "in $1 day"},
		{regexp.MustCompile(`(?i)(\d+) წუთში`), "in $1 minute"},
		{regexp.MustCompile(`(?i)(\d+) თვეში`), "in $1 month"},
		{regexp.MustCompile(`(?i)(\d+) წამში`), "in $1 second"},
		{regexp.MustCompile(`(?i)(\d+) წელში`), "in $1 year"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ კვირის წინ|\d+ საათის წინ|\d+ წამის წინ|\d+ წელიწადში|\d+ წუთის წინ|\d+ დღის წინ|\d+ თვის წინ|\d+ წლის წინ|\d+ კვირაში|\d+ კვ წინ|\d+ საათში|\d+ სთ წინ|\d+ წთ წინ|\d+ წმ წინ|\d+ დღეში|\d+ თვეში|\d+ წამში|\d+ წელში|\d+ წუთში)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ კვირის წინ|\d+ საათის წინ|\d+ წამის წინ|\d+ წელიწადში|\d+ წუთის წინ|\d+ დღის წინ|\d+ თვის წინ|\d+ წლის წინ|\d+ კვირაში|\d+ კვ წინ|\d+ საათში|\d+ სთ წინ|\d+ წთ წინ|\d+ წმ წინ|\d+ დღეში|\d+ თვეში|\d+ წამში|\d+ წელში|\d+ წუთში)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(მომავალ კვირაში|გასულ კვირაში|მომავალ თვეს|მომავალ წელს|შუადღ შემდეგ|ამ კვირაში|გასულ თვეს|გასულ წელს|დაახლოებით|სექტემბერი|ამ საათში|დეკემბერი|თებერვალი|ოთხშაბათი|ოქტომბერი|პარასკევი|სამშაბათი|ხუთშაბათი|ამ თვეში|ამ წუთში|ნოემბერი|ორშაბათი|აგვისტო|ამ წელს|დღეიდან|იანვარი|აპრილი|ივლისი|ივნისი|შაბათი|გუშინ|კვირა|მაისი|მარტი|საათი|ახლა|დღეს|წამი|წელი|წლის|წუთი|ხვალ|gmt|utc|აგვ|აპრ|დეკ|დღე|თებ|თვე|იან|ივლ|ივნ|კვი|მაი|მარ|ნოე|ოთხ|ორშ|ოქტ|პარ|სამ|სექ|შაბ|წინ|ხუთ|\+|\.|\[|\]|\||am|pm|და|ზე|ის|კვ|სთ|წთ|წმ| |'|,|-|/|:|;|@|z|წ)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
