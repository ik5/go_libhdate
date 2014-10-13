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

/**
get the number of hebrew holiday.

diaspora if true give diaspora holidays
return the number of holiday.
*/
func (h *Hdate_Struct) Get_Holyday(diaspora C.int) C.int {
	return C.hdate_get_holyday(&h.d, diaspora)
}

/**
Return the day in the omer of the given date

param h The hdate_struct of the date to use.
return The day in the omer, starting from 1 (or 0 if not in sfirat ha omer)
*/
func (h *Hdate_Struct) Get_Omer_Day() C.int {
	return C.hdate_get_omer_day(&h.d)
}

/**
Return number of hebrew holyday type.

Holiday types: set as constants as HOLYDAY_XXXX

param holyday the holyday number
return the number of holyday type.
*/
func Get_Holyday_Type(holyday C.int) C.int {
	return C.hdate_get_holyday_type(holyday)
}

/**
size of hebrew year in days.

param hebrew_year the hebrew year.
return size of Hebrew year
*/
func Get_Hebrew_Year_Size(year C.int) C.int {
	return C.hdate_get_size_of_hebrew_year(year)
}

/**
Days since Tishrey 3744

author Amos Shapir 1984 (rev. 1985, 1992) Yaacov Zamir 2003-2005

param hebrew_year The Hebrew year
return Number of days since 3,1,3744
*/
func Days_from_3744(year C.int) C.int {
	return C.hdate_days_from_3744(year)
}

/**
Return Hebrew year type based on size and first week day of year.

param size_of_year Length of year in days
param new_year_dw First week day of year
return the number for year type (1..14)
*/
func Get_Year_Type(year_size, new_year_dw C.int) C.int {
	return C.hdate_get_year_type(year_size, new_year_dw)
}

/**
Compute Julian day from Gregorian date

author Yaacov Zamir (algorithm from Henry F. Fliegel and Thomas C. Van Flandern ,1968)

param day Day of month 1..31
param month Month 1..12
param year Year in 4 digits e.g. 2001
return the julian day number
*/
func GDate_to_JD(day, month, year C.int) C.int {
	return C.hdate_gdate_to_jd(day, month, year)
}

/**
Compute Julian day from Hebrew day, month and year

author Amos Shapir 1984 (rev. 1985, 1992) Yaacov Zamir 2003-2005

param day Day of month 1..31
param month Month 1..14 (13 - Adar 1, 14 - Adar 2)
param year Hebrew year in 4 digits e.g. 5753

return Hdate_Julian struct{
  jd_tishrey1, jd_tishrey1_next_yearm day (julian)
}
param jd_tishrey1 return the julian number of 1 Tishrey this year
param jd_tishrey1_next_year return the julian number of 1 Tishrey next year
return the julian day number
*/
func HDate_To_JD(day, month, year C.int) Hdate_Julian {
	var jd_tishrey1 C.int
	var jd_tishrey1_next_year C.int

	result := C.hdate_hdate_to_jd(day, month, year,
		&jd_tishrey1, &jd_tishrey1_next_year)

	return Hdate_Julian{JD_Tishrey1: jd_tishrey1,
		JD_Tishrey1_Next_Year: jd_tishrey1_next_year,
		Day: result,
	}
}
