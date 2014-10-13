package hdate

/*
  #cgo pkg-config: libhdate
  //cgo LDFLAGS: -lhdate
  #include <hdate.h>
*/
import "C"
//import "unsafe"
import "fmt"

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

func (h *Hdate_Struct) Set_gdate(d, m, y C.int) {
  hdate := C.hdate_set_gdate(&h.d, d, m, y)
  if *hdate != &h.d {
    h.d = *hdate
  }

}
