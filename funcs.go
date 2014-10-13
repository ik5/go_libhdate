package hdate

/*
  #cgo pkg-config: libhdate
  #include <stdlib.h>
  #include <hdate.h>
*/
import "C"
import "unsafe"

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


func 
