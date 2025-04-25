package utils

import (
	"errors"
	"regexp"
)

var alarmTimePattern = regexp.MustCompile(`^(?:[01]\d|2[0-3]):(?:00|10|20|30|40|50)$`)

func ValidateAlarmTimeFormat(alarmTime string) error {
	if !alarmTimePattern.MatchString(alarmTime) {
		return errors.New("invalid alarm time format, must be 'HH:MM' with 10-minute intervals")
	}
	return nil
}
