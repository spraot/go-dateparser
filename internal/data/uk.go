// Code is generated by script; DO NOT EDIT.

package data

import "regexp"

var uk_Locale = LocaleData{
	Name:                  "uk",
	DateOrder:             "DMY",
	SentenceSplitterGroup: 1,
	SkipWords:             []string{"'", ",", "-", "/", ";", "@", "[", "]", "|", "близько", "о", "об", "і"},
	January:               []string{"січ", "січень", "січня"},
	February:              []string{"лют", "лютий", "лютого"},
	March:                 []string{"бер", "березень", "березня"},
	April:                 []string{"кві", "квіт", "квітень", "квітня"},
	May:                   []string{"тра", "трав", "травень", "травня"},
	June:                  []string{"чер", "черв", "червень", "червня"},
	July:                  []string{"лип", "липень", "липня"},
	August:                []string{"сер", "серп", "серпень", "серпня"},
	September:             []string{"вер", "вересень", "вересня"},
	October:               []string{"жов", "жовт", "жовтень", "жовтня"},
	November:              []string{"лис", "лист", "листопад", "листопада"},
	December:              []string{"гру", "груд", "грудень", "грудня"},
	Monday:                []string{"пн", "пон", "понеділок"},
	Tuesday:               []string{"вт", "вів", "вівторок"},
	Wednesday:             []string{"середа", "ср"},
	Thursday:              []string{"чет", "четвер", "чт"},
	Friday:                []string{"п'ятниця", "пт"},
	Saturday:              []string{"сб", "суб", "субота"},
	Sunday:                []string{"нд", "нед", "неділя"},
	AM:                    []string{"дп"},
	PM:                    []string{"пп"},
	Year:                  []string{"р", "роки", "років", "рік"},
	Month:                 []string{"міс", "місяць", "місяці", "місяців"},
	Week:                  []string{"тиж", "тиждень", "тижні", "тижнів"},
	Day:                   []string{"д", "день", "дні", "днів"},
	Hour:                  []string{"г", "год", "годин", "година", "години", "годину"},
	Minute:                []string{"хв", "хвилин", "хвилина", "хвилини", "хвилину"},
	Second:                []string{"с", "сек", "секунд", "секунда", "секунди", "секунду"},
	In:                    []string{"протягом"},
	Ago:                   []string{"назад", "тому"},
	RelativeType: map[string][]string{
		`0 day ago`:    {`сьогодні`},
		`0 hour ago`:   {`цієї години`},
		`0 minute ago`: {`цієї хвилини`},
		`0 month ago`:  {`цього місяця`},
		`0 second ago`: {`зараз`},
		`0 week ago`:   {`цього тижня`},
		`0 year ago`:   {`цього року`},
		`1 day ago`:    {`вчора`, `учора`},
		`1 month ago`:  {`минулого місяця`},
		`1 week ago`:   {`минулого тижня`},
		`1 year ago`:   {`торік`},
		`2 day ago`:    {`позавчора`},
		`in 1 day`:     {`завтра`},
		`in 1 month`:   {`наступного місяця`},
		`in 1 week`:    {`наступного тижня`},
		`in 1 year`:    {`наступного року`},
	},
	RelativeTypeRegex: map[string][]*regexp.Regexp{
		`$1 day ago`: {
			regexp.MustCompile(`(?i)(\d+) д тому`),
			regexp.MustCompile(`(?i)(\d+) день тому`),
			regexp.MustCompile(`(?i)(\d+) дн тому`),
			regexp.MustCompile(`(?i)(\d+) дня тому`),
		},
		`$1 hour ago`: {
			regexp.MustCompile(`(?i)(\d+) год тому`),
			regexp.MustCompile(`(?i)(\d+) години тому`),
			regexp.MustCompile(`(?i)(\d+) годину тому`),
		},
		`$1 minute ago`: {
			regexp.MustCompile(`(?i)(\d+) хв тому`),
			regexp.MustCompile(`(?i)(\d+) хвилини тому`),
			regexp.MustCompile(`(?i)(\d+) хвилину тому`),
		},
		`$1 month ago`: {
			regexp.MustCompile(`(?i)(\d+) міс тому`),
			regexp.MustCompile(`(?i)(\d+) місяць тому`),
			regexp.MustCompile(`(?i)(\d+) місяця тому`),
		},
		`$1 second ago`: {
			regexp.MustCompile(`(?i)(\d+) с тому`),
			regexp.MustCompile(`(?i)(\d+) секунди тому`),
			regexp.MustCompile(`(?i)(\d+) секунду тому`),
		},
		`$1 week ago`: {
			regexp.MustCompile(`(?i)(\d+) тиж тому`),
			regexp.MustCompile(`(?i)(\d+) тиждень тому`),
			regexp.MustCompile(`(?i)(\d+) тижня тому`),
		},
		`$1 year ago`: {
			regexp.MustCompile(`(?i)(\d+) р тому`),
			regexp.MustCompile(`(?i)(\d+) року тому`),
			regexp.MustCompile(`(?i)(\d+) рік тому`),
		},
		`in $1 day`: {
			regexp.MustCompile(`(?i)за (\d+) д`),
			regexp.MustCompile(`(?i)через (\d+) день`),
			regexp.MustCompile(`(?i)через (\d+) дн`),
			regexp.MustCompile(`(?i)через (\d+) дня`),
		},
		`in $1 hour`: {
			regexp.MustCompile(`(?i)за (\d+) год`),
			regexp.MustCompile(`(?i)через (\d+) год`),
			regexp.MustCompile(`(?i)через (\d+) години`),
			regexp.MustCompile(`(?i)через (\d+) годину`),
		},
		`in $1 minute`: {
			regexp.MustCompile(`(?i)за (\d+) хв`),
			regexp.MustCompile(`(?i)через (\d+) хв`),
			regexp.MustCompile(`(?i)через (\d+) хвилини`),
			regexp.MustCompile(`(?i)через (\d+) хвилину`),
		},
		`in $1 month`: {
			regexp.MustCompile(`(?i)за (\d+) міс`),
			regexp.MustCompile(`(?i)через (\d+) міс`),
			regexp.MustCompile(`(?i)через (\d+) місяць`),
			regexp.MustCompile(`(?i)через (\d+) місяця`),
		},
		`in $1 second`: {
			regexp.MustCompile(`(?i)за (\d+) с`),
			regexp.MustCompile(`(?i)через (\d+) с`),
			regexp.MustCompile(`(?i)через (\d+) секунди`),
			regexp.MustCompile(`(?i)через (\d+) секунду`),
		},
		`in $1 week`: {
			regexp.MustCompile(`(?i)за (\d+) тиж`),
			regexp.MustCompile(`(?i)через (\d+) тиж`),
			regexp.MustCompile(`(?i)через (\d+) тиждень`),
			regexp.MustCompile(`(?i)через (\d+) тижня`),
		},
		`in $1 year`: {
			regexp.MustCompile(`(?i)за (\d+) р`),
			regexp.MustCompile(`(?i)через (\d+) р`),
			regexp.MustCompile(`(?i)через (\d+) року`),
			regexp.MustCompile(`(?i)через (\d+) рік`),
		},
	},
	Simplifications: map[string]string{
		`^година`:       `1 година`,
		`^годину`:       `1 годину`,
		`^секунду`:      `1 секунду`,
		`^хвилину`:      `1 хвилину`,
		`за чверть`:     `15 хвилин`,
		`кілька секунд`: `44 секунди`,
		`кілька хвилин`: `2 хвилини`,
		`опів на`:       `30 хвилин`,
	},
}
