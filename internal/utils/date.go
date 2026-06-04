package utils

import "time"

func Now() time.Time {
	return time.Now()
}

func FormatByDate(t time.Time, format string) string {
	return t.Format(format)
}

func FormatByStr(dateStr, fromFormat, toFormat string) (string, error) {
	t, err := time.Parse(fromFormat, dateStr)
	if err != nil {
		return "", err
	}
	return t.Format(toFormat), nil
}
