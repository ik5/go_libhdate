package hdate

/*
  #cgo pkg-config: libhdate
  #include <stdlib.h>
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

	// for function hdate_string: identifies string type: integer
	HDATE_STRING_INT = 0

	// for function hdate_string: identifies string type: day of week
	HDATE_STRING_DOW = 1

	// for function hdate_string: identifies string type: parasha
	HDATE_STRING_PARASHA = 2

	// for function hdate_string: identifies string type: hebrew_month
	HDATE_STRING_HMONTH = 3

	// for function hdate_string: identifies string type: gregorian_month
	HDATE_STRING_GMONTH = 4

	// for function hdate_string: identifies string type: holiday
	HDATE_STRING_HOLIDAY = 5

	// for function hdate_string: identifies string type: holiday
	HDATE_STRING_OMER = 6

	// for function hdate_string: use short form, if one exists
	HDATE_STRING_SHORT = 1

	// for function hdate_string: use long form
	HDATE_STRING_LONG = 0

	// for function hdate_string: use embedded hebrew string
	HDATE_STRING_HEBREW = 1

	// for function hdate_string: use local locale string
	HDATE_STRING_LOCAL = 0
)

// libdate Heberew date struct
type Hdate_Struct struct {
	d C.hdate_struct
}

/**
compute date structure from the Gregorian date

d Day of month 1..31
m Month 1..12
  if m or d is 0 return current date.
y Year in 4 digits e.g. 2001
*/

func (h *Hdate_Struct) Set_Gdate(d, m, y C.int) {
	hdate := C.hdate_set_gdate(&h.d, d, m, y)
	if hdate != &h.d {
		h.d = *hdate
	}
}

/**
compute date structure from the Hebrew date

d Day of month 1..31
m Month 1..14 ,(13 - Adar 1, 14 - Adar 2)
	if m or d is 0 return current date.
y Year in 4 digits e.g. 5731
*/
func (h *Hdate_Struct) Set_Hdate(d, m, y C.int) {
	hdate := C.hdate_set_hdate(&h.d, d, m, y)
	if hdate != &h.d {
		h.d = *hdate
	}
}

/**
compute date structure from the Julian day

jd the julian day number.
*/
func (h *Hdate_Struct) Set_jd(jd C.int) {
	hdate := C.hdate_set_jd(&h.d, jd)
	if hdate != &h.d {
		h.d = *hdate
	}
}

/**
Return a string, with the hebrew date.

return empty string upon failure, upon success, a string containing the
short ( e.g. "1 Tishrey" ) or long (e.g. "Tuesday 18 Tishrey 5763 Hol hamoed
Sukot" ) formated date. You must free() the pointer after use.

param diaspora if true give diaspora holydays
param short_format A short flag (true - returns a short string, false returns a long string).

warning This was originally written using a local static string,
         calling for output to be copied away.

*/
func (h *Hdate_Struct) Get_Format_Date(diaspora, s C.int) string {
	var ch *C.char
	str := ""

	ch = C.hdate_get_format_date(&h.d, diaspora, s)
	if ch != nil {
		str = C.GoString(ch)
		C.free(unsafe.Pointer(ch))
	}

	return str
}

/**
get the number of hebrew parasha.

param diaspora if true give diaspora readings
return the number of parasha 1. Bereshit etc..
  (55 through 61 are joined strings e.g. Vayakhel Pekudei)
*/
func (h *Hdate_Struct) Get_Parasha(diaspora C.int) C.int {
	return C.hdate_get_parasha(&h.d, diaspora)
}
