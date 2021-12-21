// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var id_Locale = merge(nil, LocaleData{
	Name:      "id",
	DateOrder: "DMY",
	SkipWords: []string{"'", ",", "-", ".", "/", ";", "@", "[", "]", "pukul", "tanggal", "yang", "|"},
	Translations: map[string]string{
		"september": "september",
		"seminggu":  "1 week",
		"desember":  "december",
		"februari":  "february",
		"november":  "november",
		"tanggal":   "",
		"sebulan":   "1 month",
		"setahun":   "1 year",
		"agustus":   "august",
		"januari":   "january",
		"oktober":   "october",
		"sehari":    "1 day",
		"selasa":    "tuesday",
		"minggu":    "week",
		"pukul":     "",
		"april":     "april",
		"jumat":     "friday",
		"dalam":     "in",
		"maret":     "march",
		"menit":     "minute",
		"senin":     "monday",
		"bulan":     "month",
		"sabtu":     "saturday",
		"detik":     "second",
		"kamis":     "thursday",
		"tahun":     "year",
		"yang":      "",
		"lalu":      "ago",
		"hari":      "day",
		"juli":      "july",
		"juni":      "june",
		"sept":      "september",
		"ahad":      "sunday",
		"rabu":      "wednesday",
		"apr":       "april",
		"agt":       "august",
		"agu":       "august",
		"des":       "december",
		"feb":       "february",
		"jum":       "friday",
		"gmt":       "gmt",
		"jam":       "hour",
		"jan":       "january",
		"jul":       "july",
		"jun":       "june",
		"mar":       "march",
		"mei":       "may",
		"mnt":       "minute",
		"sen":       "monday",
		"bln":       "month",
		"nov":       "november",
		"okt":       "october",
		"sab":       "saturday",
		"dtk":       "second",
		"sep":       "september",
		"min":       "sunday",
		"kam":       "thursday",
		"sel":       "tuesday",
		"utc":       "utc",
		"rab":       "wednesday",
		"mgg":       "week",
		"thn":       "year",
		"am":        "am",
		"pm":        "pm",
		"'":         "",
		",":         "",
		";":         "",
		"@":         "",
		"[":         "",
		"]":         "",
		"|":         "",
		" ":         " ",
		"+":         "+",
		"-":         "-",
		".":         ".",
		"/":         "/",
		":":         ":",
		"h":         "day",
		"j":         "hour",
		"m":         "minute",
		"d":         "second",
		"z":         "z",
	},
	RelativeType: map[string]string{
		"bulan berikutnya": "in 1 month",
		"kemarin lusa":     "2 day ago",
		"minggu depan":     "in 1 week",
		"minggu lalu":      "1 week ago",
		"tahun depan":      "in 1 year",
		"minggu ini":       "0 week ago",
		"bulan lalu":       "1 month ago",
		"tahun lalu":       "1 year ago",
		"menit ini":        "0 minute ago",
		"bulan ini":        "0 month ago",
		"baru saja":        "0 second ago",
		"tahun ini":        "0 year ago",
		"hari ini":         "0 day ago",
		"sekarang":         "0 second ago",
		"jam ini":          "0 hour ago",
		"kemarin":          "1 day ago",
		"besok":            "in 1 day",
	},
	RelativeTypeRegexes: []ReplacementData{
		{regexp.MustCompile(`(?i)(\d+) minggu yang lalu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) menit yang lalu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) bulan yang lalu`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) detik yang lalu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) tahun yang lalu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)(\d+) hari yang lalu`), "$1 day ago"},
		{regexp.MustCompile(`(?i)(\d+) jam yang lalu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)dalam (\d+) minggu`), "in $1 week"},
		{regexp.MustCompile(`(?i)dalam (\d+) menit`), "in $1 minute"},
		{regexp.MustCompile(`(?i)dalam (\d+) bulan`), "in $1 month"},
		{regexp.MustCompile(`(?i)dalam (\d+) detik`), "in $1 second"},
		{regexp.MustCompile(`(?i)dalam (\d+) tahun`), "in $1 year"},
		{regexp.MustCompile(`(?i)dalam (\d+) hari`), "in $1 day"},
		{regexp.MustCompile(`(?i)dalam (\d+) jam`), "in $1 hour"},
		{regexp.MustCompile(`(?i)(\d+) jam lalu`), "$1 hour ago"},
		{regexp.MustCompile(`(?i)(\d+) mnt lalu`), "$1 minute ago"},
		{regexp.MustCompile(`(?i)(\d+) bln lalu`), "$1 month ago"},
		{regexp.MustCompile(`(?i)(\d+) dtk lalu`), "$1 second ago"},
		{regexp.MustCompile(`(?i)(\d+) mgg lalu`), "$1 week ago"},
		{regexp.MustCompile(`(?i)(\d+) thn lalu`), "$1 year ago"},
		{regexp.MustCompile(`(?i)dalam (\d+) h`), "in $1 day"},
		{regexp.MustCompile(`(?i)dlm (\d+) mnt`), "in $1 minute"},
		{regexp.MustCompile(`(?i)dlm (\d+) bln`), "in $1 month"},
		{regexp.MustCompile(`(?i)dlm (\d+) dtk`), "in $1 second"},
		{regexp.MustCompile(`(?i)dlm (\d+) mgg`), "in $1 week"},
		{regexp.MustCompile(`(?i)dlm (\d+) thn`), "in $1 year"},
		{regexp.MustCompile(`(?i)(\d+) h lalu`), "$1 day ago"},
	},
	RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+ minggu yang lalu|\d+ bulan yang lalu|\d+ detik yang lalu|\d+ menit yang lalu|\d+ tahun yang lalu|\d+ hari yang lalu|\d+ jam yang lalu|dalam \d+ minggu|dalam \d+ bulan|dalam \d+ detik|dalam \d+ menit|dalam \d+ tahun|dalam \d+ hari|dalam \d+ jam|\d+ bln lalu|\d+ dtk lalu|\d+ jam lalu|\d+ mgg lalu|\d+ mnt lalu|\d+ thn lalu|dalam \d+ h|dlm \d+ bln|dlm \d+ dtk|dlm \d+ mgg|dlm \d+ mnt|dlm \d+ thn|\d+ h lalu)(\z|[^\pL\pM\d]|_)`),
	RxExactCombined: regexp.MustCompile(`(?i)^(\d+ minggu yang lalu|\d+ bulan yang lalu|\d+ detik yang lalu|\d+ menit yang lalu|\d+ tahun yang lalu|\d+ hari yang lalu|\d+ jam yang lalu|dalam \d+ minggu|dalam \d+ bulan|dalam \d+ detik|dalam \d+ menit|dalam \d+ tahun|dalam \d+ hari|dalam \d+ jam|\d+ bln lalu|\d+ dtk lalu|\d+ jam lalu|\d+ mgg lalu|\d+ mnt lalu|\d+ thn lalu|dalam \d+ h|dlm \d+ bln|dlm \d+ dtk|dlm \d+ mgg|dlm \d+ mnt|dlm \d+ thn|\d+ h lalu)$`),
	RxKnownWords:    regexp.MustCompile(`(?i)^(.*?(?:\A|[^\pL\pM\d]|_|\d))(bulan berikutnya|kemarin lusa|minggu depan|minggu lalu|tahun depan|bulan lalu|minggu ini|tahun lalu|baru saja|bulan ini|menit ini|september|tahun ini|desember|februari|hari ini|november|sekarang|seminggu|agustus|jam ini|januari|kemarin|oktober|sebulan|setahun|tanggal|minggu|sehari|selasa|april|besok|bulan|dalam|detik|jumat|kamis|maret|menit|pukul|sabtu|senin|tahun|ahad|hari|juli|juni|lalu|rabu|sept|yang|agt|agu|apr|bln|des|dtk|feb|gmt|jam|jan|jul|jum|jun|kam|mar|mei|mgg|min|mnt|nov|okt|rab|sab|sel|sen|sep|thn|utc|\+|\.|\[|\]|\||am|pm| |'|,|-|/|:|;|@|d|h|j|m|z)((?:\z|[^\pL\pM\d]|_|\d).*)$`),
})
