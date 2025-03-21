// Code is generated by script; DO NOT EDIT.

package data

import "github.com/spraot/go-dateparser/internal/regexp"

var (
	te_Locale LocaleData
)

func init() {
	te_Locale = merge(nil, LocaleData{
		Name:      "te",
		DateOrder: "DMY",
		Charset:   []rune(`cgtuzంఅఆఈఏకగచజటడతదధనపఫబమరలళవశషసాిుూెేైో్`),
		Translations: map[string][]string{
			"గురువరం": {"thursday"},
			"మంగళవరం": {"tuesday"},
			"శుకరవరం": {"friday"},
			"సంవతసరం": {"year"},
			"బుధవరం":  {"wednesday"},
			"సపటంబర":  {"september"},
			"అకటబర":   {"october"},
			"ఆగసటు":   {"august"},
			"ఆదవరం":   {"sunday"},
			"డసంబర":   {"december"},
			"నమషము":   {"minute"},
			"నవంబర":   {"november"},
			"ఫబరవర":   {"february"},
			"శనవరం":   {"saturday"},
			"సమవరం":   {"monday"},
			"ఏపరల":    {"april"},
			"గురు":    {"thursday"},
			"జనవర":    {"january"},
			"మంగళ":    {"tuesday"},
			"వరము":    {"week"},
			"శుకర":    {"friday"},
			"సకను":    {"second"},
			"సపటం":    {"september"},
			"gmt":     {"gmt"},
			"utc":     {"utc"},
			"అకట":     {"october"},
			"ఏపర":     {"april"},
			"గంట":     {"hour"},
			"జుల":     {"july"},
			"జూన":     {"june"},
			"డసం":     {"december"},
			"దనం":     {"day"},
			"నవం":     {"november"},
			"ఫబర":     {"february"},
			"బుధ":     {"wednesday"},
			"మరచ":     {"march"},
			"am":      {"am"},
			"pm":      {"pm"},
			"ఆగ":      {"august"},
			"ఆద":      {"sunday"},
			"గం":      {"hour"},
			"జన":      {"january"},
			"నమ":      {"minute"},
			"నల":      {"month"},
			"శన":      {"saturday"},
			"సం":      {"year"},
			"సక":      {"second"},
			"సమ":      {"monday"},
			" ":       {" "},
			"'":       {""},
			"+":       {"+"},
			",":       {""},
			"-":       {"-"},
			".":       {"."},
			"/":       {"/"},
			":":       {":"},
			";":       {""},
			"@":       {""},
			"[":       {""},
			"]":       {""},
			"z":       {"z"},
			"|":       {""},
			"ద":       {"day"},
			"న":       {"month", "minute"},
			"మ":       {"may"},
			"వ":       {"week"},
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
			"ఈ గంట":         "0 hour ago",
			"ఈ రజు":         "0 day ago",
			"ఈ వరం":         "0 week ago",
			"గత నల":         "1 month ago",
			"ఈ నల":          "0 month ago",
			"ననన":           "1 day ago",
			"రపు":           "in 1 day",
		},
		RelativeTypeRegexes: []ReplacementData{
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సంవతసరం కరతం`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సంవతసరల కరతం`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) గంటల కరతం`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) నమషం కరతం`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) నమషల కరతం`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) రజుల కరతం`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సకనల కరతం`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సకను కరతం`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) గంట కరతం`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) నలల కరతం`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) రజు కరతం`), "$1 day ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) వరం కరతం`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) వరల కరతం`), "$1 week ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సంవతసరంల`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సంవతసరలల`), "in $1 year"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) గం కరతం`), "$1 hour ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) నమ కరతం`), "$1 minute ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) నల కరతం`), "$1 month ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సం కరతం`), "$1 year ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సక కరతం`), "$1 second ago"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) గంటలల`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) నమషంల`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) నమషలల`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) రజులల`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సకనలల`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సకనుల`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) గంటల`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) నలలల`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) రజుల`), "in $1 day"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) వరంల`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) వరలల`), "in $1 week"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సక ల`), "in $1 second"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) గంల`), "in $1 hour"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) నమల`), "in $1 minute"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) నలల`), "in $1 month"},
			{regexp.MustCompile(`(?i)(\d+[.,]?\d*) సంల`), "in $1 year"},
		},
		RxCombined:      regexp.MustCompile(`(?i)(\A|[^\pL\pM\d]|_)(\d+[.,]?\d* సంవతసరం కరతం|\d+[.,]?\d* సంవతసరల కరతం|\d+[.,]?\d* గంటల కరతం|\d+[.,]?\d* నమషం కరతం|\d+[.,]?\d* నమషల కరతం|\d+[.,]?\d* రజుల కరతం|\d+[.,]?\d* సకనల కరతం|\d+[.,]?\d* సకను కరతం|\d+[.,]?\d* గంట కరతం|\d+[.,]?\d* నలల కరతం|\d+[.,]?\d* రజు కరతం|\d+[.,]?\d* వరం కరతం|\d+[.,]?\d* వరల కరతం|\d+[.,]?\d* సంవతసరంల|\d+[.,]?\d* సంవతసరలల|\d+[.,]?\d* గం కరతం|\d+[.,]?\d* నమ కరతం|\d+[.,]?\d* నల కరతం|\d+[.,]?\d* సం కరతం|\d+[.,]?\d* సక కరతం|\d+[.,]?\d* గంటలల|\d+[.,]?\d* నమషంల|\d+[.,]?\d* నమషలల|\d+[.,]?\d* రజులల|\d+[.,]?\d* సకనలల|\d+[.,]?\d* సకనుల|\d+[.,]?\d* గంటల|\d+[.,]?\d* నలలల|\d+[.,]?\d* రజుల|\d+[.,]?\d* వరంల|\d+[.,]?\d* వరలల|\d+[.,]?\d* సక ల|\d+[.,]?\d* గంల|\d+[.,]?\d* నమల|\d+[.,]?\d* నలల|\d+[.,]?\d* సంల)(\z|[^\pL\pM\d]|_)`),
		RxExactCombined: regexp.MustCompile(`(?i)^(\d+[.,]?\d* సంవతసరం కరతం|\d+[.,]?\d* సంవతసరల కరతం|\d+[.,]?\d* గంటల కరతం|\d+[.,]?\d* నమషం కరతం|\d+[.,]?\d* నమషల కరతం|\d+[.,]?\d* రజుల కరతం|\d+[.,]?\d* సకనల కరతం|\d+[.,]?\d* సకను కరతం|\d+[.,]?\d* గంట కరతం|\d+[.,]?\d* నలల కరతం|\d+[.,]?\d* రజు కరతం|\d+[.,]?\d* వరం కరతం|\d+[.,]?\d* వరల కరతం|\d+[.,]?\d* సంవతసరంల|\d+[.,]?\d* సంవతసరలల|\d+[.,]?\d* గం కరతం|\d+[.,]?\d* నమ కరతం|\d+[.,]?\d* నల కరతం|\d+[.,]?\d* సం కరతం|\d+[.,]?\d* సక కరతం|\d+[.,]?\d* గంటలల|\d+[.,]?\d* నమషంల|\d+[.,]?\d* నమషలల|\d+[.,]?\d* రజులల|\d+[.,]?\d* సకనలల|\d+[.,]?\d* సకనుల|\d+[.,]?\d* గంటల|\d+[.,]?\d* నలలల|\d+[.,]?\d* రజుల|\d+[.,]?\d* వరంల|\d+[.,]?\d* వరలల|\d+[.,]?\d* సక ల|\d+[.,]?\d* గంల|\d+[.,]?\d* నమల|\d+[.,]?\d* నలల|\d+[.,]?\d* సంల)$`),
		KnownWords:      []string{"తదుపర సంవతసరం", "గత సంవతసరం", "ఈ సంవతసరం", "తదుపర వరం", "తదుపర నల", "గురువరం", "పరసతుతం", "మంగళవరం", "శుకరవరం", "సంవతసరం", "ఈ నమషం", "గత వరం", "బుధవరం", "సపటంబర", "అకటబర", "ఆగసటు", "ఆదవరం", "ఈ గంట", "ఈ రజు", "ఈ వరం", "గత నల", "డసంబర", "నమషము", "నవంబర", "ఫబరవర", "శనవరం", "సమవరం", "ఈ నల", "ఏపరల", "గురు", "జనవర", "మంగళ", "వరము", "శుకర", "సకను", "సపటం", "gmt", "utc", "అకట", "ఏపర", "గంట", "జుల", "జూన", "డసం", "దనం", "ననన", "నవం", "ఫబర", "బుధ", "మరచ", "రపు", "am", "pm", "ఆగ", "ఆద", "గం", "జన", "నమ", "నల", "శన", "సం", "సక", "సమ", " ", "'", "+", ",", "-", ".", "/", ":", ";", "@", "[", "]", "z", "|", "ద", "న", "మ", "వ"},
	})
}
