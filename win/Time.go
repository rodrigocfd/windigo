package win

import (
	"time"
)

type _TimeT struct{}

// Functions to convert between Go's time.Time and Win32's time structs.
var Time _TimeT

// Decomposes time.Duration into SYSTEMTIME fields.
func (_TimeT) DurationToSystemtime(duration time.Duration, st *SYSTEMTIME) {
	*st = SYSTEMTIME{}
	st.WHour = uint16(duration / time.Hour)
	st.WMinute = uint16((duration -
		time.Duration(st.WHour)*time.Hour) / time.Minute)
	st.WSecond = uint16((duration -
		time.Duration(st.WHour)*time.Hour -
		time.Duration(st.WMinute)*time.Minute) / time.Second)
	st.WMilliseconds = uint16((duration -
		time.Duration(st.WHour)*time.Hour -
		time.Duration(st.WMinute)*time.Minute -
		time.Duration(st.WSecond)*time.Second) / time.Millisecond)
}

// Creates a Go's time.Time from a current timezone FILETIME, millisecond precision.
func (_TimeT) FromFiletime(ftLocalTime *FILETIME) time.Time {
	st := SYSTEMTIME{}
	FileTimeToSystemTime(ftLocalTime, &st)
	return Time.FromSystemtime(&st)
}

// Creates a Go's time.Time from a current timezone SYSTEMTIME, millisecond precision.
func (_TimeT) FromSystemtime(stLocalTime *SYSTEMTIME) time.Time {
	return time.Date(int(stLocalTime.WYear),
		time.Month(stLocalTime.WMonth), int(stLocalTime.WDay),
		int(stLocalTime.WHour), int(stLocalTime.WMinute), int(stLocalTime.WSecond),
		int(stLocalTime.WMilliseconds)*1_000_000,
		time.Local)
}

// Converts Go's time.Time to a current timezone FILETIME, millisecond precision.
func (_TimeT) ToFiletime(t time.Time, ftLocalTime *FILETIME) {
	st := SYSTEMTIME{}
	Time.ToSystemtime(t, &st)
	SystemTimeToFileTime(&st, ftLocalTime)
}

// Converts Go's time.Time to a current timezone SYSTEMTIME, millisecond precision.
func (_TimeT) ToSystemtime(t time.Time, stLocalTime *SYSTEMTIME) {
	// https://support.microsoft.com/en-ca/help/167296/how-to-convert-a-unix-time-t-to-a-win32-filetime-or-systemtime
	epoch := t.UnixNano()/100 + 116_444_736_000_000_000

	ft := FILETIME{}
	ft.DwLowDateTime = uint32(epoch & 0xffff_ffff)
	ft.DwHighDateTime = uint32(epoch >> 32)

	stUtc := SYSTEMTIME{}
	FileTimeToSystemTime(&ft, &stUtc)
	SystemTimeToTzSpecificLocalTime(nil, &stUtc, stLocalTime)
}
