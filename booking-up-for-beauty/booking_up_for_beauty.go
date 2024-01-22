package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	parsedTime, _ := time.Parse(layout, date)

	return parsedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

	apointmentTime, _ := time.Parse(layout, date)
	currentTime := time.Now()

	return apointmentTime.Before(currentTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	appointmentTime, _ := time.Parse(layout, date)

	afternoonStart := time.Date(appointmentTime.Year(), appointmentTime.Month(), appointmentTime.Day(), 12, 0, 0, 0, appointmentTime.Location())
	afternoonEnd := time.Date(appointmentTime.Year(), appointmentTime.Month(), appointmentTime.Day(), 18, 0, 0, 0, appointmentTime.Location())

	return appointmentTime.After(afternoonStart) && appointmentTime.Before(afternoonEnd)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedTime, _ := time.Parse("1/2/2006 15:04:05", date)

	formattedDate := parsedTime.Format("Monday, January 2, 2006")
	formattedTime := parsedTime.Format("15:04")

	// dayOfWeek := parsedTime.Weekday().String()

	description := fmt.Sprintf("You have an appointment on %s, at %s.", formattedDate, formattedTime)
	return description
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	anniversary := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversary
}
