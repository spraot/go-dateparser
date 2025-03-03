package formatted

import (
	"testing"
	"time"

	"github.com/spraot/go-dateparser/date"
	"github.com/spraot/go-dateparser/internal/setting"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// Helper function
	dateIsParsed := func(dt date.Date) { assert.NotZero(t, dt.Time) }
	dateIsNotParsed := func(dt date.Date) { assert.Zero(t, dt.Time) }
	dateAsExpected := func(dt date.Date, expected string, expectedFormat ...string) {
		dateIsParsed(dt)

		format := "2006-01-02"
		if len(expectedFormat) > 0 {
			format = expectedFormat[0]
		}

		str := dt.Time.Format(format)
		assert.Equal(t, expected, str)
	}

	// Prepare configuration
	cfg := setting.Configuration{
		CurrentTime: time.Date(2015, 2, 4, 0, 0, 0, 0, time.UTC),
	}

	// No matching format, shouldn't be parsed
	dt := Parse(&cfg, "yesterday", "2006-01-02")
	dateIsNotParsed(dt)

	// Matching format, should be parsed
	dt = Parse(&cfg, "25-03-14", "02-01-06")
	dateAsExpected(dt, "2014-03-25")

	// Should use current year for dates without year
	dt = Parse(&cfg, "09.16", "01.02")
	dateAsExpected(dt, "2015-09-16")

	// Should use day from config for dates without day
	str, format := "August 2014", "January 2006"
	cfg = setting.Configuration{
		CurrentTime: time.Date(2014, 8, 12, 0, 0, 0, 0, time.UTC),
	}

	cfg.PreferredDayOfMonth = setting.First
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-08-01")

	cfg.PreferredDayOfMonth = setting.Last
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-08-31")

	cfg.PreferredDayOfMonth = setting.Current
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-08-12")

	// Should use month from config for dates without month
	str, format = "2014", "2006"
	cfg = setting.Configuration{
		CurrentTime: time.Date(2014, 7, 1, 0, 0, 0, 0, time.UTC),
	}

	cfg.PreferredMonthOfYear = setting.FirstMonth
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-01-01")

	cfg.PreferredMonthOfYear = setting.LastMonth
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-12-01")

	cfg.PreferredMonthOfYear = setting.CurrentMonth
	dt = Parse(&cfg, str, format)
	dateAsExpected(dt, "2014-07-01")
}
