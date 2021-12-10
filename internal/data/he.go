// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var he_Locale = LocaleData{
	Name:                  "he",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "ב-", "בסביבות", "בערך", "בקירוב", "בשעה", "ה-", "，"},
	January:               []string{"בינואר", "ינו", "ינואר", "ינו׳", "לינואר"},
	February:              []string{"בפברואר", "לפברואר", "פבר", "פברואר", "פבר׳"},
	March:                 []string{"במארס", "במרס", "במרץ", "למארס", "למרס", "למרץ", "מארס", "מרס", "מרץ"},
	April:                 []string{"אפר", "אפריל", "אפר׳", "באפריל", "לאפריל"},
	May:                   []string{"במאי", "למאי", "מאי"},
	June:                  []string{"ביוני", "יונ", "יוני", "ליוני"},
	July:                  []string{"ביולי", "יול", "יולי", "ליולי"},
	August:                []string{"אוג", "אוגוסט", "אוג׳", "באוגוסט", "לאוגוסט"},
	September:             []string{"בספטמבר", "לספטמבר", "ספט", "ספטמבר", "ספט׳"},
	October:               []string{"אוק", "אוקטובר", "אוק׳", "באוקטובר", "לאוקטובר"},
	November:              []string{"בנובמבר", "לנובמבר", "נוב", "נובמבר", "נוב׳"},
	December:              []string{"בדצמבר", "דצמ", "דצמבר", "דצמ׳", "לדצמבר"},
	Monday:                []string{"יום ב", "יום ב׳", "יום שני", "שני"},
	Tuesday:               []string{"יום ג", "יום ג׳", "יום שלישי", "שלישי"},
	Wednesday:             []string{"יום ד", "יום ד׳", "יום רביעי", "רביעי"},
	Thursday:              []string{"חמישי", "יום ה", "יום ה׳", "יום חמישי"},
	Friday:                []string{"יום ו", "יום ו׳", "יום שישי", "שישי"},
	Saturday:              []string{"יום שבת", "שבת"},
	Sunday:                []string{"יום א", "יום א׳", "יום ראשון", "ראשון"},
	AM:                    []string{"לפנה״צ"},
	PM:                    []string{"אחה״צ"},
	Year:                  []string{"בשנה", "שנה", "שנים", "שנ׳"},
	Month:                 []string{"בחודש", "חודש", "חודשים", "חו׳"},
	Week:                  []string{"שבוע", "שבועות", "שב׳"},
	Day:                   []string{"ביום", "יום", "ימים"},
	Hour:                  []string{"שעה", "שעות", "שע׳"},
	Minute:                []string{"דקה", "דקות", "דק׳"},
	Second:                []string{"שניה", "שניות", "שנייה", "שנ׳"},
	In:                    []string{"בעוד", "עוד"},
	Ago:                   []string{"לפני"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`היום`},
		`0 hour ago`:   {`בשעה זו`},
		`0 minute ago`: {`בדקה זו`},
		`0 month ago`:  {`החודש`},
		`0 second ago`: {`עכשיו`},
		`0 week ago`:   {`השבוע`},
		`0 year ago`:   {`השנה`},
		`1 day ago`:    {`אתמול`},
		`1 month ago`:  {`החודש שעבר`},
		`1 week ago`:   {`השבוע שעבר`},
		`1 year ago`:   {`השנה שעברה`},
		`in 1 day`:     {`מחר`},
		`in 1 month`:   {`החודש הבא`},
		`in 1 week`:    {`השבוע הבא`},
		`in 1 year`:    {`השנה הבאה`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`\1 day ago`: {
			regexp.MustCompile(`(?i)לפני (\d+) ימים`),
			regexp.MustCompile(`(?i)לפני (\d+) ימ׳`),
			regexp.MustCompile(`(?i)לפני יום (\d+)`),
		},
		`\1 hour ago`: {
			regexp.MustCompile(`(?i)לפני (\d+) שעות`),
			regexp.MustCompile(`(?i)לפני (\d+) שע׳`),
		},
		`\1 minute ago`: {
			regexp.MustCompile(`(?i)לפני (\d+) דקות`),
			regexp.MustCompile(`(?i)לפני (\d+) דק׳`),
		},
		`\1 month ago`: {
			regexp.MustCompile(`(?i)לפני (\d+) חודשים`),
			regexp.MustCompile(`(?i)לפני (\d+) חו׳`),
		},
		`\1 second ago`: {
			regexp.MustCompile(`(?i)לפני (\d+) שניות`),
			regexp.MustCompile(`(?i)לפני (\d+) שנ׳`),
		},
		`\1 week ago`: {
			regexp.MustCompile(`(?i)לפני (\d+) שבועות`),
			regexp.MustCompile(`(?i)לפני (\d+) שב׳`),
		},
		`\1 year ago`: {
			regexp.MustCompile(`(?i)לפני (\d+) שנים`),
		},
		`in \1 day`: {
			regexp.MustCompile(`(?i)בעוד (\d+) ימים`),
			regexp.MustCompile(`(?i)בעוד (\d+) ימ׳`),
			regexp.MustCompile(`(?i)בעוד יום (\d+)`),
		},
		`in \1 hour`: {
			regexp.MustCompile(`(?i)בעוד (\d+) שעות`),
			regexp.MustCompile(`(?i)בעוד (\d+) שע׳`),
		},
		`in \1 minute`: {
			regexp.MustCompile(`(?i)בעוד (\d+) דקות`),
			regexp.MustCompile(`(?i)בעוד (\d+) דק׳`),
		},
		`in \1 month`: {
			regexp.MustCompile(`(?i)בעוד (\d+) חודשים`),
			regexp.MustCompile(`(?i)בעוד (\d+) חו׳`),
		},
		`in \1 second`: {
			regexp.MustCompile(`(?i)בעוד (\d+) שניות`),
			regexp.MustCompile(`(?i)בעוד (\d+) שנ׳`),
		},
		`in \1 week`: {
			regexp.MustCompile(`(?i)בעוד (\d+) שבועות`),
			regexp.MustCompile(`(?i)בעוד (\d+) שב׳`),
		},
		`in \1 year`: {
			regexp.MustCompile(`(?i)בעוד (\d+) שנים`),
		},
	},
	Simplifications: map[string]string{
		`אחה"צ`:          `pm`,
		`אחר חצות`:       `am`,
		`אחרי ה?צהריי?ם`: `pm`,
		`בבוקר`:          `am`,
		`בערב`:           `pm`,
		`בצהריי?ם`:       `pm`,
		`ו(\w+)`:         `\1`,
		`וחודש`:          `1 חודש`,
		`ויום`:           `1 יום`,
		`ושבוע`:          `1 שבוע`,
		`ושנה`:           `1 שנה`,
		`חודשיים`:        `2 חודשים`,
		`חצות`:           `12 am`,
		`יומיי?ם`:        `2 ימים`,
		`לפנות בוקר`:     `am`,
		`לפנות ערב`:      `pm`,
		`מחר`:            `בעוד 1 יום`,
		`שבועיי?ם`:       `2 שבועות`,
		`שלשום`:          `2 ימים`,
		`שנתיי?ם`:        `2 שנים`,
		`שעתיי?ם`:        `2 שעות`,
	},
}
