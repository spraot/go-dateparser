package timezone

import "github.com/spraot/go-dateparser/internal/regexp"

// Info is the data for parsing timezone from a string.
type Info struct {
	RegexPatterns       []string
	AlternativePatterns map[*regexp.Regexp]string
	Timezones           map[string]int
}

// These timezone list are taken fromseveral sources:
// - http://stackoverflow.com/q/1703546
// - http://en.wikipedia.org/wiki/List_of_time_zone_abbreviations
// - https://github.com/scrapinghub/dateparser/pull/4
// - http://en.wikipedia.org/wiki/List_of_UTC_time_offsets
var timezoneInfoList = []Info{
	{
		RegexPatterns: []string{
			`(.)%s$`,
		},
		AlternativePatterns: map[*regexp.Regexp]string{
			// UTC+n, UTC-n, GMT+n, GMT-n:
			regexp.MustCompile(`(?:UTC|GMT)\\?([+-])0(\d):00`): `(?:UTC|GMT)\$1$2`,
			// UTC+n:mm, UTC-n:mm, GMT+n:mm, GMT-n:mm:
			regexp.MustCompile(`(?:UTC|GMT)\\?([+-])0(\d):(\d{2})`): `(?:UTC|GMT)\$1$2:$3`,
			// UTC+nn, UTC-nn, GMT+nn, GMT-nn:
			regexp.MustCompile(`(?:UTC|GMT)\\?([+-])(\d{2}):00`): `(?:UTC|GMT)\$1$2`,
			// UTC+nnmm, UTC-nnmm, GMT+nnmm, GMT-nnmm:
			regexp.MustCompile(`(?:UTC|GMT)\\?([+-])(\d{2}):(\d{2})`): `(?:UTC|GMT)\$1$2:?$3.*`,
			// Others:
			regexp.MustCompile(`UTC`):   "",
			regexp.MustCompile(`:`):     "",
			regexp.MustCompile(`:|UTC`): "",
			regexp.MustCompile(`UTC`):   "GMT",
		},
		Timezones: map[string]int{
			`UTC-12:00`: -43200,
			`UTC-11:00`: -39600,
			`UTC-10:00`: -36000,
			`UTC-09:30`: -34200,
			`UTC-09:00`: -32400,
			`UTC-08:00`: -28800,
			`UTC-07:00`: -25200,
			`UTC-06:00`: -21600,
			`UTC-05:00`: -18000,
			`UTC-04:30`: -16200,
			`UTC-04:00`: -14400,
			`UTC-03:30`: -12600,
			`UTC-03:00`: -10800,
			`UTC-02:30`: -9000,
			`UTC-02:00`: -7200,
			`UTC-01:00`: -3600,
			`UTC-00:00`: 0,
			`UTC+00:00`: 0,
			`UTC+01:00`: 3600,
			`UTC+02:00`: 7200,
			`UTC+03:00`: 10800,
			`UTC+03:30`: 12600,
			`UTC+04:00`: 14400,
			`UTC+04:30`: 16200,
			`UTC+05:00`: 18000,
			`UTC+05:30`: 19800,
			`UTC+05:45`: 20700,
			`UTC+06:00`: 21600,
			`UTC+06:30`: 23400,
			`UTC+07:00`: 25200,
			`UTC+08:00`: 28800,
			`UTC+08:45`: 31500,
			`UTC+09:00`: 32400,
			`UTC+09:30`: 34200,
			`UTC+10:00`: 36000,
			`UTC+10:30`: 37800,
			`UTC+11:00`: 39600,
			`UTC+11:30`: 41400,
			`UTC+12:00`: 43200,
			`UTC+12:45`: 45900,
			`UTC+13:00`: 46800,
			`UTC+14:00`: 50400,
		},
	}, {
		RegexPatterns: []string{
			(`(\W|\d|_)%s($|\W)`),
		},
		Timezones: map[string]int{
			`ACDT`:  37800,
			`ACST`:  34200,
			`ACT`:   -18000,
			`ACWDT`: 35100,
			`ACWST`: 31500,
			`ADDT`:  -7200,
			`ADMT`:  9300,
			`ADT`:   -10800,
			`AEDT`:  39600,
			`AEST`:  36000,
			`AFT`:   16200,
			`AHDT`:  -32400,
			`AHST`:  -36000,
			`AKDT`:  -28800,
			`AKST`:  -32400,
			`AKTST`: 21600,
			`AKTT`:  18000,
			`ALMST`: 25200,
			`ALMT`:  21600,
			`AMST`:  18000,
			`AMT`:   14400,
			`ANAST`: 43200,
			`ANAT`:  43200,
			`ANT`:   -16200,
			`APT`:   -10800,
			`AQTST`: 21600,
			`AQTT`:  18000,
			`ARST`:  -10800,
			`ART`:   -10800,
			`ASHST`: 21600,
			`ASHT`:  18000,
			`AST`:   -14400,
			`AWDT`:  32400,
			`AWST`:  28800,
			`AWT`:   -10800,
			`AZOMT`: 0,
			`AZOST`: -3600,
			`AZOT`:  -3600,
			`AZST`:  18000,
			`AZT`:   14400,
			`BAKST`: 14400,
			`BAKT`:  10800,
			`BDST`:  7200,
			`BDT`:   28800,
			`BEAT`:  9000,
			`BEAUT`: 9900,
			`BIOT`:  21600,
			`BMT`:   1800,
			`BNT`:   28800,
			`BORT`:  28800,
			`BOST`:  -12780,
			`BOT`:   -14400,
			`BRST`:  -7200,
			`BRT`:   -10800,
			`BST`:   39600,
			`BTT`:   21600,
			`BURT`:  23400,
			`CANT`:  -3600,
			`CAPT`:  -32400,
			`CAST`:  10800,
			`CAT`:   7200,
			`CAWT`:  -32400,
			`CCT`:   23400,
			`CDDT`:  -14400,
			`CDT`:   -18000,
			`CEDT`:  7200,
			`CEMT`:  10800,
			`CEST`:  7200,
			`CET`:   3600,
			`CGST`:  -3600,
			`CGT`:   -7200,
			`CHADT`: 49500,
			`CHAST`: 45900,
			`CHDT`:  -19800,
			`CHOST`: 36000,
			`CHOT`:  28800,
			`CIST`:  -28800,
			`CKHST`: -34200,
			`CKT`:   -36000,
			`CLST`:  -10800,
			`CLT`:   -14400,
			`CMT`:   -16080,
			`COST`:  -14400,
			`COT`:   -18000,
			`CPT`:   -18000,
			`CST`:   -21600,
			`CUT`:   8400,
			`CVST`:  -3600,
			`CVT`:   -3600,
			`CWT`:   -18000,
			`CXT`:   25200,
			`ChST`:  36000,
			`DACT`:  21600,
			`DAVT`:  25200,
			`DDUT`:  36000,
			`DFT`:   3600,
			`DMT`:   -1500,
			`DUSST`: 21600,
			`DUST`:  21600,
			`EASST`: -18000,
			`EAST`:  -21600,
			`EAT`:   10800,
			`ECT`:   -18000,
			`EDDT`:  -10800,
			`EDT`:   -14400,
			`EEDT`:  10800,
			`EEST`:  10800,
			`EET`:   7200,
			`EGST`:  0,
			`EGT`:   -3600,
			`EHDT`:  -16200,
			`EMT`:   -26220,
			`EPT`:   -14400,
			`EST`:   -18000,
			`ET`:    -18000,
			`EWT`:   -14400,
			`FET`:   10800,
			`FFMT`:  -14640,
			`FJST`:  46800,
			`FJT`:   43200,
			`FKST`:  -10800,
			`FKT`:   -14400,
			`FMT`:   -4080,
			`FNST`:  -3600,
			`FNT`:   -7200,
			`FORT`:  14400,
			`FRUST`: 25200,
			`FRUT`:  18000,
			`GALT`:  -21600,
			`GAMT`:  -32400,
			`GBGT`:  -13500,
			`GEST`:  14400,
			`GET`:   14400,
			`GFT`:   -10800,
			`GHST`:  1200,
			`GILT`:  43200,
			`GIT`:   -32400,
			`GMT`:   0,
			`GST`:   14400,
			`GYT`:   -14400,
			`HAA`:   -10800,
			`HAC`:   -18000,
			`HADT`:  -32400,
			`HAE`:   -14400,
			`HAP`:   -25200,
			`HAR`:   -21600,
			`HAST`:  -36000,
			`HAT`:   -9000,
			`HAY`:   -28800,
			`HDT`:   -34200,
			`HKST`:  32400,
			`HKT`:   28800,
			`HLV`:   -16200,
			`HMT`:   18000,
			`HNA`:   -14400,
			`HNC`:   -21600,
			`HNE`:   -18000,
			`HNP`:   -28800,
			`HNR`:   -25200,
			`HNT`:   -12600,
			`HNY`:   -32400,
			`HOVST`: 28800,
			`HOVT`:  25200,
			`HST`:   -36000,
			`ICT`:   25200,
			`IDDT`:  14400,
			`IDT`:   10800,
			`IHST`:  21600,
			`IMT`:   7020,
			`IOT`:   21600,
			`IRDT`:  16200,
			`IRKST`: 32400,
			`IRKT`:  28800,
			`IRST`:  12600,
			`ISST`:  0,
			`IST`:   7200,
			`JAVT`:  26400,
			`JCST`:  32400,
			`JDT`:   36000,
			`JMT`:   8460,
			`JST`:   32400,
			`JWST`:  28800,
			`KART`:  18000,
			`KDT`:   32400,
			`KGST`:  21600,
			`KGT`:   21600,
			`KIZST`: 21600,
			`KIZT`:  18000,
			`KMT`:   5760,
			`KOST`:  39600,
			`KRAST`: 28800,
			`KRAT`:  25200,
			`KST`:   32400,
			`KUYST`: 18000,
			`KUYT`:  14400,
			`KWAT`:  -43200,
			`LHDT`:  39600,
			`LHST`:  37800,
			`LINT`:  50400,
			`LKT`:   23400,
			`LMT`:   -20160,
			// `LMT`:   -17640,
			// `LMT`:   -20580,
			// `LMT`:   -14400,
			`LRT`:   -2640,
			`LST`:   9420,
			`MADMT`: 3600,
			`MADST`: 0,
			`MADT`:  -3600,
			`MAGST`: 43200,
			`MAGT`:  39600,
			`MALST`: 26400,
			`MALT`:  27000,
			`MART`:  -34200,
			`MAWT`:  18000,
			`MDDT`:  -18000,
			`MDST`:  16260,
			`MDT`:   -21600,
			`MEST`:  7200,
			`MESZ`:  7200,
			`MET`:   3600,
			`MEZ`:   3600,
			`MHT`:   43200,
			`MIST`:  39600,
			`MIT`:   -34200,
			`MMT`:   23400,
			`MOST`:  32400,
			`MOT`:   28800,
			`MPT`:   -21600,
			`MSD`:   14400,
			`MSK`:   10800,
			`MSM`:   18000,
			`MST`:   -25200,
			`MUST`:  18000,
			`MUT`:   14400,
			`MVT`:   18000,
			`MWT`:   -21600,
			`MYT`:   28800,
			`NCST`:  43200,
			`NCT`:   39600,
			`NDDT`:  -5400,
			`NDT`:   -9000,
			`NEGT`:  -12600,
			`NEST`:  4800,
			`NET`:   1200,
			`NFT`:   41400,
			`NMT`:   40320,
			`NOVST`: 25200,
			`NOVT`:  21600,
			`NPT`:   20700,
			`NRT`:   41400,
			`NST`:   -12600,
			`NT`:    -12600,
			`NUT`:   -39600,
			`NWT`:   -36000,
			`NZDT`:  46800,
			`NZMT`:  41400,
			`NZST`:  43200,
			`OMSST`: 25200,
			`OMST`:  21600,
			`ORAST`: 18000,
			`ORAT`:  18000,
			`PDDT`:  -21600,
			`PDT`:   -25200,
			`PEST`:  -14400,
			`PET`:   -18000,
			`PETST`: 43200,
			`PETT`:  43200,
			`PGT`:   36000,
			`PHOT`:  46800,
			`PHST`:  32400,
			`PHT`:   28800,
			`PKST`:  21600,
			`PKT`:   18000,
			`PLMT`:  25620,
			`PMDT`:  -7200,
			`PMMT`:  35340,
			`PMST`:  -10800,
			`PMT`:   540,
			`PNT`:   -30600,
			`PONT`:  39600,
			`PPMT`:  -17340,
			`PPT`:   -25200,
			`PST`:   -28800,
			`PT`:    -28800,
			`PWT`:   -25200,
			`PYST`:  -10800,
			`PYT`:   -14400,
			`QMT`:   -18840,
			`QYZST`: 25200,
			`QYZT`:  21600,
			`RET`:   14400,
			`RMT`:   3000,
			`ROTT`:  -10800,
			`SAKST`: 43200,
			`SAKT`:  39600,
			`SAMT`:  14400,
			`SAST`:  7200,
			`SBT`:   39600,
			`SCT`:   14400,
			`SDMT`:  -16800,
			`SDT`:   -36000,
			`SEČ`:   3600,
			`SET`:   3600,
			`SELČ`:  7200,
			`SGT`:   28800,
			`SHEST`: 21600,
			`SHET`:  18000,
			`SJMT`:  -20160,
			`SLT`:   19800,
			`SMT`:   -13860,
			`SRET`:  39600,
			`SRT`:   -10800,
			`SST`:   -39600,
			`STAT`:  10800,
			`SVEST`: 21600,
			`SVET`:  14400,
			`SWAT`:  5400,
			`SYOT`:  10800,
			`TAHT`:  -36000,
			`TASST`: 25200,
			`TAST`:  21600,
			`TBIST`: 18000,
			`TBIT`:  10800,
			`TBMT`:  10740,
			`TFT`:   18000,
			`THA`:   25200,
			`TJT`:   18000,
			`TKT`:   -39600,
			`TLT`:   32400,
			`TMT`:   18000,
			`TOST`:  50400,
			`TOT`:   46800,
			`TRST`:  14400,
			`TRT`:   10800,
			`TSAT`:  10800,
			`TVT`:   43200,
			`ULAST`: 32400,
			`ULAT`:  28800,
			`URAST`: 18000,
			`URAT`:  18000,
			`UT`:    0,
			`UTC`:   0,
			`UYHST`: -9000,
			`UYST`:  -7200,
			`UYT`:   -10800,
			`UZST`:  21600,
			`UZT`:   18000,
			`VEČ`:   7200,
			`VET`:   -16200,
			`VLAST`: 39600,
			`VLAT`:  36000,
			`VOLST`: 14400,
			`VOLT`:  14400,
			`VOST`:  21600,
			`VUST`:  43200,
			`VUT`:   39600,
			`WARST`: -10800,
			`WART`:  -14400,
			`WAST`:  7200,
			`WAT`:   3600,
			`WDT`:   32400,
			`WEDT`:  3600,
			`WEMT`:  7200,
			`WEST`:  3600,
			`WET`:   0,
			`WFT`:   43200,
			`WGST`:  -7200,
			`WGT`:   -10800,
			`WIB`:   25200,
			`WIT`:   32400,
			`WITA`:  28800,
			`WMT`:   5040,
			`WSDT`:  50400,
			`WSST`:  46800,
			`WST`:   28800,
			`WT`:    0,
			`XJT`:   21600,
			`YAKST`: 36000,
			`YAKT`:  32400,
			`YAPT`:  36000,
			`YDDT`:  -25200,
			`YDT`:   -28800,
			`YEKST`: 21600,
			`YEKT`:  18000,
			`YERST`: 14400,
			`YERT`:  10800,
			`YPT`:   -28800,
			`YST`:   -32400,
			`YWT`:   -28800,
			`zzz`:   0,
			`Z`:     0,
			`ZEČ`:   0,
		},
	},
}
