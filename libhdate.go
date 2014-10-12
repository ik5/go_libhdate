package hdate

/*
  #cgo pkg-config: hdate
  #cgo LDFLAGS: -lhdate
  #include <hdate.h>
*/
import "C"
import "unsafe"

const (
	// use diaspora dates and holydays flag
	DIASPORA = -1
	// use israel dates and holydays flag
	ISRAEL_FLAG = 0

	// use short strings flag
	SHORT_FLAG = -1
	// use long strings flag
	LONG_FLAG = 0

	/*
	  for hdate_get_int_string_ and hdate_get_int_wstring

	  How large should the buffer be? Hebrew year 10,999 would
	  be י'תתקצ"ט, eight characters, each two bytes, plus an
	  end-of-string delimiter, equals 17. This could effectively
	  yield a range extending to Hebrew year 11,899, י"א תתצ"ט,
	  due to the extra ק needed for the '900' century. However,
	  for readability, I would want a an extra space at that
	  point between the millenium and the century...

	*/
	HEBREW_NUMBER_BUFFER_SIZE  = 17
	HEBREW_WNUMBER_BUFFER_SIZE = 9
)

// libdate Heberew date struct
type Struct struct {
	/** The number of day in the hebrew month (1..31). */
	hd_day C.int
	/** The number of the hebrew month 1..14 (1 - tishre, 13 - adar 1, 14 - adar 2). */
	hd_mon C.int
	/** The number of the hebrew year. */
	hd_year C.int
	/** The number of the day in the month. (1..31) */
	gd_day C.int
	/** The number of the month 1..12 (1 - jan). */
	gd_mon C.int
	/** The number of the year. */
	gd_year C.int
	/** The day of the week 1..7 (1 - sunday). */
	hd_dw C.int
	/** The length of the year in days. */
	hd_size_of_year C.int
	/** The week day of Hebrew new year. */
	hd_new_year_dw C.int
	/** The number type of year. */
	hd_year_type C.int
	/** The Julian day number */
	hd_jd C.int
	/** The number of days passed since 1 tishrey */
	hd_days C.int
	/** The number of weeks passed since 1 tishrey */
	hd_weeks C.int
}

/**
compute date structure from the Gregorian date

d Day of month 1..31
m Month 1..12
  if m or d is 0 return current date.
y Year in 4 digits e.g. 2001
*/

func (h *Struct) Set_gdate(d, m, y C.int) {
	return C.hdate_set_gdate(h, d, m, y)
}
