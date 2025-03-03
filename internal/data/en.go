// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	en_Locale     LocaleData
	en_001_Locale LocaleData
	en_150_Locale LocaleData
	en_AG_Locale  LocaleData
	en_AI_Locale  LocaleData
	en_AS_Locale  LocaleData
	en_AT_Locale  LocaleData
	en_AU_Locale  LocaleData
	en_BB_Locale  LocaleData
	en_BE_Locale  LocaleData
	en_BI_Locale  LocaleData
	en_BM_Locale  LocaleData
	en_BS_Locale  LocaleData
	en_BW_Locale  LocaleData
	en_BZ_Locale  LocaleData
	en_CA_Locale  LocaleData
	en_CC_Locale  LocaleData
	en_CH_Locale  LocaleData
	en_CK_Locale  LocaleData
	en_CM_Locale  LocaleData
	en_CX_Locale  LocaleData
	en_CY_Locale  LocaleData
	en_DE_Locale  LocaleData
	en_DG_Locale  LocaleData
	en_DK_Locale  LocaleData
	en_DM_Locale  LocaleData
	en_ER_Locale  LocaleData
	en_FI_Locale  LocaleData
	en_FJ_Locale  LocaleData
	en_FK_Locale  LocaleData
	en_FM_Locale  LocaleData
	en_GB_Locale  LocaleData
	en_GD_Locale  LocaleData
	en_GG_Locale  LocaleData
	en_GH_Locale  LocaleData
	en_GI_Locale  LocaleData
	en_GM_Locale  LocaleData
	en_GU_Locale  LocaleData
	en_GY_Locale  LocaleData
	en_HK_Locale  LocaleData
	en_IE_Locale  LocaleData
	en_IL_Locale  LocaleData
	en_IM_Locale  LocaleData
	en_IN_Locale  LocaleData
	en_IO_Locale  LocaleData
	en_JE_Locale  LocaleData
	en_JM_Locale  LocaleData
	en_KE_Locale  LocaleData
	en_KI_Locale  LocaleData
	en_KN_Locale  LocaleData
	en_KY_Locale  LocaleData
	en_LC_Locale  LocaleData
	en_LR_Locale  LocaleData
	en_LS_Locale  LocaleData
	en_MG_Locale  LocaleData
	en_MH_Locale  LocaleData
	en_MO_Locale  LocaleData
	en_MP_Locale  LocaleData
	en_MS_Locale  LocaleData
	en_MT_Locale  LocaleData
	en_MU_Locale  LocaleData
	en_MW_Locale  LocaleData
	en_MY_Locale  LocaleData
	en_NA_Locale  LocaleData
	en_NF_Locale  LocaleData
	en_NG_Locale  LocaleData
	en_NL_Locale  LocaleData
	en_NR_Locale  LocaleData
	en_NU_Locale  LocaleData
	en_NZ_Locale  LocaleData
	en_PG_Locale  LocaleData
	en_PH_Locale  LocaleData
	en_PK_Locale  LocaleData
	en_PN_Locale  LocaleData
	en_PR_Locale  LocaleData
	en_PW_Locale  LocaleData
	en_RW_Locale  LocaleData
	en_SB_Locale  LocaleData
	en_SC_Locale  LocaleData
	en_SD_Locale  LocaleData
	en_SE_Locale  LocaleData
	en_SG_Locale  LocaleData
	en_SH_Locale  LocaleData
	en_SI_Locale  LocaleData
	en_SL_Locale  LocaleData
	en_SS_Locale  LocaleData
	en_SX_Locale  LocaleData
	en_SZ_Locale  LocaleData
	en_TC_Locale  LocaleData
	en_TK_Locale  LocaleData
	en_TO_Locale  LocaleData
	en_TT_Locale  LocaleData
	en_TV_Locale  LocaleData
	en_TZ_Locale  LocaleData
	en_UG_Locale  LocaleData
	en_UM_Locale  LocaleData
	en_VC_Locale  LocaleData
	en_VG_Locale  LocaleData
	en_VI_Locale  LocaleData
	en_VU_Locale  LocaleData
	en_WS_Locale  LocaleData
	en_ZA_Locale  LocaleData
	en_ZM_Locale  LocaleData
	en_ZW_Locale  LocaleData
)

func init() {
	en_Locale = merge(nil, LocaleData{
		Name:                  "en",
		DateOrder:             "MDY",
		Charset:               []rune(`bcdefghijklnorstuvwxyz`),
		SentenceSplitterGroup: 1,
		Simplifications: []ReplacementData{
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])an(\z|[^\pL\pM\d_])`), "${1}1${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])a(\z|[^\pL\pM\d_])`), "${1}1${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])(?:12\s+)?noon(\z|[^\pL\pM\d_])`), "${1}12:00${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])(?:12\s+)?midnight(\z|[^\pL\pM\d_])`), "${1}00:00${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])(\d+[.,]?\d*)h(\d+[.,]?\d*)(\z|[^\pL\pM\d_])`), "${1}${2}:${3}${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])(from\s+)now(\z|[^\pL\pM\d_])`), "${1}${2}in${3}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])less than 1 minute ago(\z|[^\pL\pM\d_])`), "${1}45 second ago${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])(\d+[.,]?\d*) (decade|year|month|week|day|hour|minute|second)s? later(\z|[^\pL\pM\d_])`), "${1}in ${2} ${3}${4}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])one(\z|[^\pL\pM\d_])`), "${1}1${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])two(\z|[^\pL\pM\d_])`), "${1}2${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])three(\z|[^\pL\pM\d_])`), "${1}3${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])four(\z|[^\pL\pM\d_])`), "${1}4${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])five(\z|[^\pL\pM\d_])`), "${1}5${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])six(\z|[^\pL\pM\d_])`), "${1}6${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])seven(\z|[^\pL\pM\d_])`), "${1}7${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])eight(\z|[^\pL\pM\d_])`), "${1}8${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])nine(\z|[^\pL\pM\d_])`), "${1}9${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])ten(\z|[^\pL\pM\d_])`), "${1}10${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])eleven(\z|[^\pL\pM\d_])`), "${1}11${2}"},
			{regexp.MustCompile(`(?i)(\A|[^\pL\pM\d_])twelve(\z|[^\pL\pM\d_])`), "${1}12${2}"},
		},
		Translations: map[string][]string{
			"september": {"september"},
			"wednesday": {"wednesday"},
			"december":  {"december"},
			"february":  {"february"},
			"from now":  {"in"},
			"november":  {"november"},
			"saturday":  {"saturday"},
			"thursday":  {"thursday"},
			"decades":   {"decade"},
			"january":   {"january"},
			"minutes":   {"minute"},
			"october":   {"october"},
			"seconds":   {"second"},
			"tuesday":   {"tuesday"},
			"august":    {"august"},
			"before":    {"ago"},
			"decade":    {"decade"},
			"friday":    {"friday"},
			"minute":    {"minute"},
			"monday":    {"monday"},
			"months":    {"month"},
			"second":    {"second"},
			"sunday":    {"sunday"},
			"about":     {""},
			"after":     {"in"},
			"april":     {"april"},
			"hours":     {"hour"},
			"march":     {"march"},
			"month":     {"month"},
			"weeks":     {"week"},
			"years":     {"year"},
			"days":      {"day"},
			"hour":      {"hour"},
			"july":      {"july"},
			"june":      {"june"},
			"just":      {""},
			"mins":      {"minute"},
			"secs":      {"second"},
			"sept":      {"september"},
			"tues":      {"tuesday"},
			"week":      {"week"},
			"year":      {"year"},
			"ago":       {"ago"},
			"and":       {""},
			"apr":       {"april"},
			"aug":       {"august"},
			"day":       {"day"},
			"dec":       {"december"},
			"feb":       {"february"},
			"fri":       {"friday"},
			"gmt":       {"gmt"},
			"hrs":       {"hour"},
			"jan":       {"january"},
			"jul":       {"july"},
			"jun":       {"june"},
			"mar":       {"march"},
			"may":       {"may"},
			"min":       {"minute"},
			"mon":       {"monday"},
			"nov":       {"november"},
			"oct":       {"october"},
			"sat":       {"saturday"},
			"sec":       {"second"},
			"sep":       {"september"},
			"sun":       {"sunday"},
			"the":       {""},
			"thu":       {"thursday"},
			"tue":       {"tuesday"},
			"utc":       {"utc"},
			"wed":       {"wednesday"},
			"ad":        {""},
			"am":        {"am"},
			"at":        {""},
			"by":        {""},
			"hr":        {"hour"},
			"in":        {"in"},
			"mo":        {"month"},
			"nd":        {""},
			"of":        {""},
			"on":        {""},
			"pm":        {"pm"},
			"rd":        {""},
			"st":        {""},
			"th":        {""},
			"wk":        {"week"},
			"yr":        {"year"},
			" ":         {" "},
			"'":         {""},
			"+":         {"+"},
			",":         {""},
			"-":         {"-"},
			".":         {"."},
			"/":         {"/"},
			":":         {":"},
			";":         {""},
			"@":         {""},
			"[":         {""},
			"]":         {""},
			"d":         {"day"},
			"h":         {"hour"},
			"m":         {"minute"},
			"s":         {"second"},
			"y":         {"year"},
			"z":         {"z"},
			"|":         {""},
		},
		RelativeType: map[string]string{
			"day before yesterday": "2 day ago",
			"day after tomorrow":   "in 2 day",
			"last decade":          "1 decade ago",
			"next decade":          "in 1 decade",
			"this decade":          "1 decade ago",
			"this minute":          "0 minute ago",
			"last month":           "1 month ago",
			"next month":           "in 1 month",
			"this month":           "0 month ago",
			"last week":            "1 week ago",
			"last year":            "1 year ago",
			"next week":            "in 1 week",
			"next year":            "in 1 year",
			"this hour":            "0 hour ago",
			"this week":            "0 week ago",
			"this year":            "0 year ago",
			"till date":            "0 day ago",
			"yesterday":            "1 day ago",
			"tomorrow":             "in 1 day",
			"last mo":              "1 month ago",
			"last wk":              "1 week ago",
			"last yr":              "1 year ago",
			"next mo":              "in 1 month",
			"next wk":              "in 1 week",
			"next yr":              "in 1 year",
			"this mo":              "0 month ago",
			"this wk":              "0 week ago",
			"this yr":              "0 year ago",
			"today":                "0 day ago",
			"now":                  "0 second ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) decades? ago`), "${1} decade ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minutes ago`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) seconds ago`), "$1 second ago"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) decades?`), "in ${1} decade"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) minute ago`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) months ago`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) second ago`), "$1 second ago"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) minutes`), "in $1 minute"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) seconds`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) hours ago`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) month ago`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) weeks ago`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) years ago`), "$1 year ago"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) minute`), "in $1 minute"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) months`), "in $1 month"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) second`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) days ago`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) hour ago`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) week ago`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) year ago`), "$1 year ago"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) hours`), "in $1 hour"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) month`), "in $1 month"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) weeks`), "in $1 week"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) years`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) day ago`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) min ago`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) sec ago`), "$1 second ago"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) days`), "in $1 day"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) hour`), "in $1 hour"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) week`), "in $1 week"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) year`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) hr ago`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) mo ago`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) wk ago`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) yr ago`), "$1 year ago"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) day`), "in $1 day"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) min`), "in $1 minute"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) sec`), "in $1 second"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) hr`), "in $1 hour"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) mo`), "in $1 month"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) wk`), "in $1 week"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) yr`), "in $1 year"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* decades? ago|\d+[.,]?\d* minutes ago|\d+[.,]?\d* seconds ago|in \d+[.,]?\d* decades?|\d+[.,]?\d* minute ago|\d+[.,]?\d* months ago|\d+[.,]?\d* second ago|in \d+[.,]?\d* minutes|in \d+[.,]?\d* seconds|\d+[.,]?\d* hours ago|\d+[.,]?\d* month ago|\d+[.,]?\d* weeks ago|\d+[.,]?\d* years ago|in \d+[.,]?\d* minute|in \d+[.,]?\d* months|in \d+[.,]?\d* second|\d+[.,]?\d* days ago|\d+[.,]?\d* hour ago|\d+[.,]?\d* week ago|\d+[.,]?\d* year ago|in \d+[.,]?\d* hours|in \d+[.,]?\d* month|in \d+[.,]?\d* weeks|in \d+[.,]?\d* years|\d+[.,]?\d* day ago|\d+[.,]?\d* min ago|\d+[.,]?\d* sec ago|in \d+[.,]?\d* days|in \d+[.,]?\d* hour|in \d+[.,]?\d* week|in \d+[.,]?\d* year|\d+[.,]?\d* hr ago|\d+[.,]?\d* mo ago|\d+[.,]?\d* wk ago|\d+[.,]?\d* yr ago|in \d+[.,]?\d* day|in \d+[.,]?\d* min|in \d+[.,]?\d* sec|in \d+[.,]?\d* hr|in \d+[.,]?\d* mo|in \d+[.,]?\d* wk|in \d+[.,]?\d* yr)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* decades? ago|\d+[.,]?\d* minutes ago|\d+[.,]?\d* seconds ago|in \d+[.,]?\d* decades?|\d+[.,]?\d* minute ago|\d+[.,]?\d* months ago|\d+[.,]?\d* second ago|in \d+[.,]?\d* minutes|in \d+[.,]?\d* seconds|\d+[.,]?\d* hours ago|\d+[.,]?\d* month ago|\d+[.,]?\d* weeks ago|\d+[.,]?\d* years ago|in \d+[.,]?\d* minute|in \d+[.,]?\d* months|in \d+[.,]?\d* second|\d+[.,]?\d* days ago|\d+[.,]?\d* hour ago|\d+[.,]?\d* week ago|\d+[.,]?\d* year ago|in \d+[.,]?\d* hours|in \d+[.,]?\d* month|in \d+[.,]?\d* weeks|in \d+[.,]?\d* years|\d+[.,]?\d* day ago|\d+[.,]?\d* min ago|\d+[.,]?\d* sec ago|in \d+[.,]?\d* days|in \d+[.,]?\d* hour|in \d+[.,]?\d* week|in \d+[.,]?\d* year|\d+[.,]?\d* hr ago|\d+[.,]?\d* mo ago|\d+[.,]?\d* wk ago|\d+[.,]?\d* yr ago|in \d+[.,]?\d* day|in \d+[.,]?\d* min|in \d+[.,]?\d* sec|in \d+[.,]?\d* hr|in \d+[.,]?\d* mo|in \d+[.,]?\d* wk|in \d+[.,]?\d* yr)$`),
		KnownWords:      []string{"day before yesterday", "day after tomorrow", "last decade", "next decade", "this decade", "this minute", "last month", "next month", "this month", "last week", "last year", "next week", "next year", "september", "this hour", "this week", "this year", "till date", "wednesday", "yesterday", "december", "february", "from now", "november", "saturday", "thursday", "tomorrow", "decades", "january", "last mo", "last wk", "last yr", "minutes", "next mo", "next wk", "next yr", "october", "seconds", "this mo", "this wk", "this yr", "tuesday", "august", "before", "decade", "friday", "minute", "monday", "months", "second", "sunday", "about", "after", "april", "hours", "march", "month", "today", "weeks", "years", "days", "hour", "july", "june", "just", "mins", "secs", "sept", "tues", "week", "year", "ago", "and", "apr", "aug", "day", "dec", "feb", "fri", "gmt", "hrs", "jan", "jul", "jun", "mar", "may", "min", "mon", "nov", "now", "oct", "sat", "sec", "sep", "sun", "the", "thu", "tue", "utc", "wed", "ad", "am", "at", "by", "hr", "in", "mo", "nd", "of", "on", "pm", "rd", "st", "th", "wk", "yr", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "h", "m", "s", "y", "z", "|"},
	})

	en_001_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-001",
		DateOrder: "DMY",
	})

	en_150_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-150",
		DateOrder: "DMY",
	})

	en_AG_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-AG",
		DateOrder: "DMY",
	})

	en_AI_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-AI",
		DateOrder: "DMY",
	})

	en_AS_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-AS",
		DateOrder: "MDY",
	})

	en_AT_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-AT",
		DateOrder: "DMY",
	})

	en_AU_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-AU",
		DateOrder: "DMY",
	})

	en_BB_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-BB",
		DateOrder: "DMY",
	})

	en_BE_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-BE",
		DateOrder: "DMY",
	})

	en_BI_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-BI",
		DateOrder: "MDY",
	})

	en_BM_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-BM",
		DateOrder: "DMY",
	})

	en_BS_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-BS",
		DateOrder: "DMY",
	})

	en_BW_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-BW",
		DateOrder: "DMY",
	})

	en_BZ_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-BZ",
		DateOrder: "DMY",
	})

	en_CA_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-CA",
		DateOrder: "YMD",
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) mins ago`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) secs ago`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) hrs ago`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) mos ago`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) wks ago`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) yrs ago`), "$1 year ago"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) mins`), "in $1 minute"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) secs`), "in $1 second"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) hrs`), "in $1 hour"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) mos`), "in $1 month"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) wks`), "in $1 week"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) yrs`), "in $1 year"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* decades? ago|\d+[.,]?\d* minutes ago|\d+[.,]?\d* seconds ago|in \d+[.,]?\d* decades?|\d+[.,]?\d* minute ago|\d+[.,]?\d* months ago|\d+[.,]?\d* second ago|in \d+[.,]?\d* minutes|in \d+[.,]?\d* seconds|\d+[.,]?\d* hours ago|\d+[.,]?\d* month ago|\d+[.,]?\d* weeks ago|\d+[.,]?\d* years ago|in \d+[.,]?\d* minute|in \d+[.,]?\d* months|in \d+[.,]?\d* second|\d+[.,]?\d* days ago|\d+[.,]?\d* hour ago|\d+[.,]?\d* mins ago|\d+[.,]?\d* secs ago|\d+[.,]?\d* week ago|\d+[.,]?\d* year ago|in \d+[.,]?\d* hours|in \d+[.,]?\d* month|in \d+[.,]?\d* weeks|in \d+[.,]?\d* years|\d+[.,]?\d* day ago|\d+[.,]?\d* hrs ago|\d+[.,]?\d* min ago|\d+[.,]?\d* mos ago|\d+[.,]?\d* sec ago|\d+[.,]?\d* wks ago|\d+[.,]?\d* yrs ago|in \d+[.,]?\d* days|in \d+[.,]?\d* hour|in \d+[.,]?\d* mins|in \d+[.,]?\d* secs|in \d+[.,]?\d* week|in \d+[.,]?\d* year|\d+[.,]?\d* hr ago|\d+[.,]?\d* mo ago|\d+[.,]?\d* wk ago|\d+[.,]?\d* yr ago|in \d+[.,]?\d* day|in \d+[.,]?\d* hrs|in \d+[.,]?\d* min|in \d+[.,]?\d* mos|in \d+[.,]?\d* sec|in \d+[.,]?\d* wks|in \d+[.,]?\d* yrs|in \d+[.,]?\d* hr|in \d+[.,]?\d* mo|in \d+[.,]?\d* wk|in \d+[.,]?\d* yr)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* decades? ago|\d+[.,]?\d* minutes ago|\d+[.,]?\d* seconds ago|in \d+[.,]?\d* decades?|\d+[.,]?\d* minute ago|\d+[.,]?\d* months ago|\d+[.,]?\d* second ago|in \d+[.,]?\d* minutes|in \d+[.,]?\d* seconds|\d+[.,]?\d* hours ago|\d+[.,]?\d* month ago|\d+[.,]?\d* weeks ago|\d+[.,]?\d* years ago|in \d+[.,]?\d* minute|in \d+[.,]?\d* months|in \d+[.,]?\d* second|\d+[.,]?\d* days ago|\d+[.,]?\d* hour ago|\d+[.,]?\d* mins ago|\d+[.,]?\d* secs ago|\d+[.,]?\d* week ago|\d+[.,]?\d* year ago|in \d+[.,]?\d* hours|in \d+[.,]?\d* month|in \d+[.,]?\d* weeks|in \d+[.,]?\d* years|\d+[.,]?\d* day ago|\d+[.,]?\d* hrs ago|\d+[.,]?\d* min ago|\d+[.,]?\d* mos ago|\d+[.,]?\d* sec ago|\d+[.,]?\d* wks ago|\d+[.,]?\d* yrs ago|in \d+[.,]?\d* days|in \d+[.,]?\d* hour|in \d+[.,]?\d* mins|in \d+[.,]?\d* secs|in \d+[.,]?\d* week|in \d+[.,]?\d* year|\d+[.,]?\d* hr ago|\d+[.,]?\d* mo ago|\d+[.,]?\d* wk ago|\d+[.,]?\d* yr ago|in \d+[.,]?\d* day|in \d+[.,]?\d* hrs|in \d+[.,]?\d* min|in \d+[.,]?\d* mos|in \d+[.,]?\d* sec|in \d+[.,]?\d* wks|in \d+[.,]?\d* yrs|in \d+[.,]?\d* hr|in \d+[.,]?\d* mo|in \d+[.,]?\d* wk|in \d+[.,]?\d* yr)$`),
	})

	en_CC_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-CC",
		DateOrder: "DMY",
	})

	en_CH_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-CH",
		DateOrder: "DMY",
	})

	en_CK_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-CK",
		DateOrder: "DMY",
	})

	en_CM_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-CM",
		DateOrder: "DMY",
	})

	en_CX_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-CX",
		DateOrder: "DMY",
	})

	en_CY_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-CY",
		DateOrder: "DMY",
	})

	en_DE_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-DE",
		DateOrder: "DMY",
	})

	en_DG_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-DG",
		DateOrder: "DMY",
	})

	en_DK_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-DK",
		DateOrder: "DMY",
	})

	en_DM_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-DM",
		DateOrder: "DMY",
	})

	en_ER_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-ER",
		DateOrder: "DMY",
	})

	en_FI_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-FI",
		DateOrder: "DMY",
	})

	en_FJ_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-FJ",
		DateOrder: "DMY",
	})

	en_FK_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-FK",
		DateOrder: "DMY",
	})

	en_FM_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-FM",
		DateOrder: "DMY",
	})

	en_GB_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-GB",
		DateOrder: "DMY",
	})

	en_GD_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-GD",
		DateOrder: "DMY",
	})

	en_GG_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-GG",
		DateOrder: "DMY",
	})

	en_GH_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-GH",
		DateOrder: "DMY",
	})

	en_GI_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-GI",
		DateOrder: "DMY",
	})

	en_GM_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-GM",
		DateOrder: "DMY",
	})

	en_GU_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-GU",
		DateOrder: "MDY",
	})

	en_GY_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-GY",
		DateOrder: "DMY",
	})

	en_HK_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-HK",
		DateOrder: "DMY",
	})

	en_IE_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-IE",
		DateOrder: "DMY",
	})

	en_IL_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-IL",
		DateOrder: "DMY",
	})

	en_IM_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-IM",
		DateOrder: "DMY",
	})

	en_IN_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-IN",
		DateOrder: "DMY",
	})

	en_IO_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-IO",
		DateOrder: "DMY",
	})

	en_JE_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-JE",
		DateOrder: "DMY",
	})

	en_JM_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-JM",
		DateOrder: "DMY",
	})

	en_KE_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-KE",
		DateOrder: "DMY",
	})

	en_KI_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-KI",
		DateOrder: "DMY",
	})

	en_KN_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-KN",
		DateOrder: "DMY",
	})

	en_KY_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-KY",
		DateOrder: "DMY",
	})

	en_LC_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-LC",
		DateOrder: "DMY",
	})

	en_LR_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-LR",
		DateOrder: "DMY",
	})

	en_LS_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-LS",
		DateOrder: "DMY",
	})

	en_MG_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-MG",
		DateOrder: "DMY",
	})

	en_MH_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-MH",
		DateOrder: "MDY",
	})

	en_MO_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-MO",
		DateOrder: "DMY",
	})

	en_MP_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-MP",
		DateOrder: "MDY",
	})

	en_MS_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-MS",
		DateOrder: "DMY",
	})

	en_MT_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-MT",
		DateOrder: "DMY",
	})

	en_MU_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-MU",
		DateOrder: "DMY",
	})

	en_MW_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-MW",
		DateOrder: "DMY",
	})

	en_MY_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-MY",
		DateOrder: "DMY",
	})

	en_NA_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-NA",
		DateOrder: "DMY",
	})

	en_NF_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-NF",
		DateOrder: "DMY",
	})

	en_NG_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-NG",
		DateOrder: "DMY",
	})

	en_NL_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-NL",
		DateOrder: "DMY",
	})

	en_NR_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-NR",
		DateOrder: "DMY",
	})

	en_NU_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-NU",
		DateOrder: "DMY",
	})

	en_NZ_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-NZ",
		DateOrder: "DMY",
	})

	en_PG_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-PG",
		DateOrder: "DMY",
	})

	en_PH_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-PH",
		DateOrder: "DMY",
	})

	en_PK_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-PK",
		DateOrder: "DMY",
	})

	en_PN_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-PN",
		DateOrder: "DMY",
	})

	en_PR_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-PR",
		DateOrder: "MDY",
	})

	en_PW_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-PW",
		DateOrder: "DMY",
	})

	en_RW_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-RW",
		DateOrder: "DMY",
	})

	en_SB_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-SB",
		DateOrder: "DMY",
	})

	en_SC_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-SC",
		DateOrder: "DMY",
	})

	en_SD_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-SD",
		DateOrder: "DMY",
	})

	en_SE_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-SE",
		DateOrder: "YMD",
	})

	en_SG_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-SG",
		DateOrder: "DMY",
		Translations: map[string][]string{
			"mth": {"month"},
		},
		RelativeType: map[string]string{
			"last mth": "1 month ago",
			"next mth": "in 1 month",
			"this mth": "0 month ago",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) mth ago`), "$1 month ago"},
			{regexp.MustCompile(`(?i)in (\d+[.,]?\d*) mth`), "in $1 month"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* decades? ago|\d+[.,]?\d* minutes ago|\d+[.,]?\d* seconds ago|in \d+[.,]?\d* decades?|\d+[.,]?\d* minute ago|\d+[.,]?\d* months ago|\d+[.,]?\d* second ago|in \d+[.,]?\d* minutes|in \d+[.,]?\d* seconds|\d+[.,]?\d* hours ago|\d+[.,]?\d* month ago|\d+[.,]?\d* weeks ago|\d+[.,]?\d* years ago|in \d+[.,]?\d* minute|in \d+[.,]?\d* months|in \d+[.,]?\d* second|\d+[.,]?\d* days ago|\d+[.,]?\d* hour ago|\d+[.,]?\d* week ago|\d+[.,]?\d* year ago|in \d+[.,]?\d* hours|in \d+[.,]?\d* month|in \d+[.,]?\d* weeks|in \d+[.,]?\d* years|\d+[.,]?\d* day ago|\d+[.,]?\d* min ago|\d+[.,]?\d* mth ago|\d+[.,]?\d* sec ago|in \d+[.,]?\d* days|in \d+[.,]?\d* hour|in \d+[.,]?\d* week|in \d+[.,]?\d* year|\d+[.,]?\d* hr ago|\d+[.,]?\d* mo ago|\d+[.,]?\d* wk ago|\d+[.,]?\d* yr ago|in \d+[.,]?\d* day|in \d+[.,]?\d* min|in \d+[.,]?\d* mth|in \d+[.,]?\d* sec|in \d+[.,]?\d* hr|in \d+[.,]?\d* mo|in \d+[.,]?\d* wk|in \d+[.,]?\d* yr)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* decades? ago|\d+[.,]?\d* minutes ago|\d+[.,]?\d* seconds ago|in \d+[.,]?\d* decades?|\d+[.,]?\d* minute ago|\d+[.,]?\d* months ago|\d+[.,]?\d* second ago|in \d+[.,]?\d* minutes|in \d+[.,]?\d* seconds|\d+[.,]?\d* hours ago|\d+[.,]?\d* month ago|\d+[.,]?\d* weeks ago|\d+[.,]?\d* years ago|in \d+[.,]?\d* minute|in \d+[.,]?\d* months|in \d+[.,]?\d* second|\d+[.,]?\d* days ago|\d+[.,]?\d* hour ago|\d+[.,]?\d* week ago|\d+[.,]?\d* year ago|in \d+[.,]?\d* hours|in \d+[.,]?\d* month|in \d+[.,]?\d* weeks|in \d+[.,]?\d* years|\d+[.,]?\d* day ago|\d+[.,]?\d* min ago|\d+[.,]?\d* mth ago|\d+[.,]?\d* sec ago|in \d+[.,]?\d* days|in \d+[.,]?\d* hour|in \d+[.,]?\d* week|in \d+[.,]?\d* year|\d+[.,]?\d* hr ago|\d+[.,]?\d* mo ago|\d+[.,]?\d* wk ago|\d+[.,]?\d* yr ago|in \d+[.,]?\d* day|in \d+[.,]?\d* min|in \d+[.,]?\d* mth|in \d+[.,]?\d* sec|in \d+[.,]?\d* hr|in \d+[.,]?\d* mo|in \d+[.,]?\d* wk|in \d+[.,]?\d* yr)$`),
		KnownWords:      []string{"day before yesterday", "day after tomorrow", "last decade", "next decade", "this decade", "this minute", "last month", "next month", "this month", "last week", "last year", "next week", "next year", "september", "this hour", "this week", "this year", "till date", "wednesday", "yesterday", "december", "february", "from now", "last mth", "next mth", "november", "saturday", "this mth", "thursday", "tomorrow", "decades", "january", "last mo", "last wk", "last yr", "minutes", "next mo", "next wk", "next yr", "october", "seconds", "this mo", "this wk", "this yr", "tuesday", "august", "before", "decade", "friday", "minute", "monday", "months", "second", "sunday", "about", "after", "april", "hours", "march", "month", "today", "weeks", "years", "days", "hour", "july", "june", "just", "mins", "secs", "sept", "tues", "week", "year", "ago", "and", "apr", "aug", "day", "dec", "feb", "fri", "gmt", "hrs", "jan", "jul", "jun", "mar", "may", "min", "mon", "mth", "nov", "now", "oct", "sat", "sec", "sep", "sun", "the", "thu", "tue", "utc", "wed", "ad", "am", "at", "by", "hr", "in", "mo", "nd", "of", "on", "pm", "rd", "st", "th", "wk", "yr", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "d", "h", "m", "s", "y", "z", "|"},
	})

	en_SH_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-SH",
		DateOrder: "DMY",
	})

	en_SI_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-SI",
		DateOrder: "DMY",
	})

	en_SL_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-SL",
		DateOrder: "DMY",
	})

	en_SS_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-SS",
		DateOrder: "DMY",
	})

	en_SX_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-SX",
		DateOrder: "DMY",
	})

	en_SZ_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-SZ",
		DateOrder: "DMY",
	})

	en_TC_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-TC",
		DateOrder: "DMY",
	})

	en_TK_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-TK",
		DateOrder: "DMY",
	})

	en_TO_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-TO",
		DateOrder: "DMY",
	})

	en_TT_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-TT",
		DateOrder: "DMY",
	})

	en_TV_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-TV",
		DateOrder: "DMY",
	})

	en_TZ_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-TZ",
		DateOrder: "DMY",
	})

	en_UG_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-UG",
		DateOrder: "DMY",
	})

	en_UM_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-UM",
		DateOrder: "MDY",
	})

	en_VC_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-VC",
		DateOrder: "DMY",
	})

	en_VG_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-VG",
		DateOrder: "DMY",
	})

	en_VI_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-VI",
		DateOrder: "MDY",
	})

	en_VU_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-VU",
		DateOrder: "DMY",
	})

	en_WS_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-WS",
		DateOrder: "DMY",
	})

	en_ZA_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-ZA",
		DateOrder: "YMD",
	})

	en_ZM_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-ZM",
		DateOrder: "DMY",
	})

	en_ZW_Locale = merge(&en_Locale, LocaleData{
		Name:      "en-ZW",
		DateOrder: "DMY",
	})
}
