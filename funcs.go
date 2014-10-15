/*
Copyright (C) 2014  Ido Kanner <idokan at@at gmail dot.dot com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.
*/
package hdate

/*
  #cgo pkg-config: libhdate
  #include <stdlib.h>
  #include <hdate.h>
*/
import "C"
import "unsafe"

/**
Initialize the Hdate_Struct

return Zero'd Hdate_Struct
*/
func Init() *Hdate_Struct {
	var h *Hdate_Struct = new(Hdate_Struct)
	var d *C.hdate_struct
	d = C.new_hdate()
	h.d = d

	return h
}

/**
Destruct memory of Hdate_Struct, must be called last
*/
func (h *Hdate_Struct) Destruct() {
	C.delete_hdate(h.d)
}

/**
compute date structure from the Gregorian date

d Day of month 1..31
m Month 1..12
  if m or d is 0 return current date.
y Year in 4 digits e.g. 2001
*/

func (h *Hdate_Struct) Set_Gdate(d, m, y int) {
	hdate := C.hdate_set_gdate(h.d, C.int(d), C.int(m), C.int(y))
	if hdate != h.d {
		h.d = hdate
	}
}

/**
compute date structure from the Hebrew date

d Day of month 1..31
m Month 1..14 ,(13 - Adar 1, 14 - Adar 2)
	if m or d is 0 return current date.
y Year in 4 digits e.g. 5731
*/
func (h *Hdate_Struct) Set_Hdate(d, m, y int) {
	hdate := C.hdate_set_hdate(h.d, C.int(d), C.int(m), C.int(y))
	if hdate != h.d {
		h.d = hdate
	}
}

/**
compute date structure from the Julian day

jd the julian day number.
*/
func (h *Hdate_Struct) Set_jd(jd int) {
	hdate := C.hdate_set_jd(h.d, C.int(jd))
	if hdate != h.d {
		h.d = hdate
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
func (h *Hdate_Struct) Get_Format_Date(diaspora, s int) string {
	var ch *C.char
	str := ""

	ch = C.hdate_get_format_date(h.d, C.int(diaspora), C.int(s))
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
func (h *Hdate_Struct) Get_Parasha(diaspora int) int {
	return int(C.hdate_get_parasha(h.d, C.int(diaspora)))
}

/**
get the number of hebrew holiday.

diaspora if true give diaspora holidays
return the number of holiday.
*/
func (h *Hdate_Struct) Get_Holyday(diaspora int) int {
	return int(C.hdate_get_holyday(h.d, C.int(diaspora)))
}

/**
Return the day in the omer of the given date

param h The hdate_struct of the date to use.
return The day in the omer, starting from 1 (or 0 if not in sfirat ha omer)
*/
func (h *Hdate_Struct) Get_Omer_Day() int {
	return int(C.hdate_get_omer_day(h.d))
}

/**
Return number of hebrew holyday type.

Holiday types: set as constants as HOLYDAY_XXXX

param holyday the holyday number
return the number of holyday type.
*/
func Get_Holyday_Type(holyday int) int {
	return int(C.hdate_get_holyday_type(C.int(holyday)))
}

/**
size of hebrew year in days.

param hebrew_year the hebrew year.
return size of Hebrew year
*/
func Get_Hebrew_Year_Size(year int) int {
	return int(C.hdate_get_size_of_hebrew_year(C.int(year)))
}

/**
Days since Tishrey 3744

author Amos Shapir 1984 (rev. 1985, 1992) Yaacov Zamir 2003-2005

param hebrew_year The Hebrew year
return Number of days since 3,1,3744
*/
func Days_from_3744(year int) int {
	return int(C.hdate_days_from_3744(C.int(year)))
}

/**
Return Hebrew year type based on size and first week day of year.

param size_of_year Length of year in days
param new_year_dw First week day of year
return the number for year type (1..14)
*/
func Get_Year_Type(year_size, new_year_dw int) int {
	return int(C.hdate_get_year_type(C.int(year_size), C.int(new_year_dw)))
}

/**
Compute Julian day from Gregorian date

author Yaacov Zamir (algorithm from Henry F. Fliegel and Thomas C. Van Flandern ,1968)

param day Day of month 1..31
param month Month 1..12
param year Year in 4 digits e.g. 2001
return the julian day number
*/
func GDate_to_JD(day, month, year int) int {
	return int(C.hdate_gdate_to_jd(C.int(day), C.int(month), C.int(year)))
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
func HDate_To_JD(day, month, year int) Hdate_Julian {
	var jd_tishrey1 C.int
	var jd_tishrey1_next_year C.int

	result := C.hdate_hdate_to_jd(C.int(day), C.int(month), C.int(year),
		&jd_tishrey1, &jd_tishrey1_next_year)

	return Hdate_Julian{
		JD_Tishrey1:           int(jd_tishrey1),
		JD_Tishrey1_Next_Year: int(jd_tishrey1_next_year),
		Day: int(result),
	}
}

/**
days from 1 january

param day this day of month
param month this month
param year this year
return the days from 1 jan
*/
func Get_Day_of_Year(day, month, year int) int {
	return int(C.hdate_get_day_of_year(C.int(day), C.int(month), C.int(year)))
}

/**
utc sun times for altitude at a gregorian date

Returns the sunset and sunrise times in minutes from 00:00 (utc time)
if sun altitude in sunrise is deg degries.
This function only works for altitudes sun realy is.
If the sun never get to this altitude, the returned sunset and sunrise values
will be negative. This can happen in low altitude when latitude is
nearing the pols in winter times, the sun never goes very high in
the sky there.

param day this day of month
param month this month
param year this year
param longitude longitude to use in calculations
param latitude latitude to use in calculations
param deg degrees of sun's altitude (0 -  Zenith .. 90 - Horizon)

return:
param sunrise return the utc sunrise in minutes
param sunset return the utc sunset in minutes
*/
func Get_UTC_Sun_Time_Deg(day, month, year int, latitude, longitude, deg float64) Sunrise_Sunset {
	var sunrise, sunset C.int
	C.hdate_get_utc_sun_time_deg(C.int(day), C.int(month), C.int(year),
		C.double(latitude), C.double(longitude), C.double(deg), &sunrise, &sunset)

	return Sunrise_Sunset{
		Sunrise: int(sunrise),
		Sunset:  int(sunset),
	}
}

/**
utc sunrise/set time for a gregorian date

param day this day of month
param month this month
param year this year
param longitude longitude to use in calculations
degrees, negative values are east
param latitude latitude to use in calculations
degrees, negative values are south

returns
param sunrise return the utc sunrise in minutes after midnight (00:00)
param sunset return the utc sunset in minutes after midnight (00:00)
*/
func Get_UTC_Sun_Time(day, month, year int, latitude, longitude float64) Sunrise_Sunset {
	var sunrise, sunset C.int
	C.hdate_get_utc_sun_time(C.int(day), C.int(month), C.int(year),
		C.double(latitude), C.double(longitude), &sunrise, &sunset)

	return Sunrise_Sunset{
		Sunrise: int(sunrise),
		Sunset:  int(sunset),
	}
}

/**
utc sunrise/set time for a gregorian date

param day this day of month
param month this month
param year this year
param longitude longitude to use in calculations
param latitude latitude to use in calculations

returns
param sun_hour return the length of shaa zaminit in minutes
param first_light return the utc alut ha-shachar in minutes
param talit return the utc tphilin and talit in minutes
param sunrise return the utc sunrise in minutes
param midday return the utc midday in minutes
param sunset return the utc sunset in minutes
param first_stars return the utc tzeit hacochavim in minutes
param three_stars return the utc shlosha cochavim in minutes
*/
func Get_UTC_Sun_Time_Full(day, month, year int, latitude, longitude float64) Sun_Time {
	var sun_hour, first_light, talit, sunrise, midday, sunset C.int
	var first_stars, three_stars C.int

	C.hdate_get_utc_sun_time_full(C.int(day), C.int(month), C.int(year),
		C.double(latitude), C.double(longitude),
		&sun_hour, &first_light, &talit, &sunrise, &midday, &sunset,
		&first_stars, &three_stars)

	return Sun_Time{
		Sun_Hour:    int(sun_hour),
		First_Light: int(first_light),
		Talit:       int(talit),
		Sunrise:     int(sunrise),
		MidDay:      int(midday),
		Sunset:      int(sunset),
		First_Stars: int(first_stars),
		Three_Stars: int(three_stars),
	}
}

/**
get the Gregorian day of the month

return the Gregorian day of the month, 1..31.
*/
func (h *Hdate_Struct) Get_GDay() int {
	return int(C.hdate_get_gday(h.d))
}

/**
get the Gregorian month

return the Gregorian month, jan = 1.
*/
func (h *Hdate_Struct) Get_GMonth() int {
	return int(C.hdate_get_gmonth(h.d))
}

/**
get the Gregorian year

return the Gregorian year.
*/
func (h *Hdate_Struct) Get_GYear() int {
	return int(C.hdate_get_gyear(h.d))
}

/**
get the Hebrew day of the month

return the Hebrew day of the month, 1..30.
*/
func (h *Hdate_Struct) Get_HDay() int {
	return int(C.hdate_get_hday(h.d))
}

/**
get the Hebrew month

return the Hebrew month, Tishery = 1 .. Adar I =13, Adar II = 14.
*/
func (h *Hdate_Struct) Get_HMonth() int {
	return int(C.hdate_get_hmonth(h.d))
}

/**
get the Hebrew year

return the Hebrew year.
*/
func (h *Hdate_Struct) Get_HYear() int {
	return int(C.hdate_get_hyear(h.d))
}

/**
get the day of the week

return the the day of the week.
*/
func (h *Hdate_Struct) Get_Day_Of_Week() int {
	return int(C.hdate_get_day_of_the_week(h.d))
}

/**
get the size of the hebrew year

return the the size of the hebrew year.
*/
func (h *Hdate_Struct) Get_Size_of_Year() int {
	return int(C.hdate_get_size_of_year(h.d))
}

/**
get the new year day of the week

return the the new year day of the week.
*/
func (h *Hdate_Struct) Get_New_Year_Day_of_Week() int {
	return int(C.hdate_get_new_year_day_of_the_week(h.d))
}

/**
get the Julian day number

return the Julian day number.
*/
func (h *Hdate_Struct) Get_Julian() int {
	return int(C.hdate_get_julian(h.d))
}

/**
get the number of days passed since 1 tishrey

return the number of days passed since 1 tishrey.
*/
func (h *Hdate_Struct) Get_Days() int {
	return int(C.hdate_get_days(h.d))
}

/**
get the number of weeks passed since 1 tishrey

return the number of weeks passed since 1 tishrey.
*/
func (h *Hdate_Struct) Get_Weeks() int {
	return int(C.hdate_get_weeks(h.d))
}

/**
Return a static string, with the package name and version

return a string, with the package name and version
*/
func Get_Version() string {
	return C.GoString(C.hdate_get_version_string())
}

/**
name of translator

return a string with name of translator, or empty string if none.
*/
func Get_Translator() string {
	trans := C.hdate_get_translator_string()
	result := ""

	if trans != nil {
		result = C.GoString(trans)
	}

	return result
}

/**
Helper function to find Hebrew local

return true if Hebrew locale
*/
func Is_Hebrew_Locale() bool {
	locale := C.hdate_is_hebrew_locale()
	return locale == -1
}

/**
 Return string values for hdate information

 param type_of_string 	0 = integer, 1 = day of week, 2 = parshaot,
						3 = hmonth, 4 = gmonth, 5 = holiday, 6 = omer
 param index			integer		( 0 < n < 11000)
						day of week ( 0 < n <  8 )
						parshaot	( 0 , n < 62 )
						hmonth		( 0 < n < 15 )
						gmonth		( 0 < n < 13 )
						holiday		( 0 < n < 37 )
						omer		( 0 < n < 50 )
 param short_form   0 = short format
 param hebrew_form  0 = not hebrew (native/embedded)

 return  a string containing the information.
         return empty string upon failure.
*/
func Hdate_String(string_type, index, short_form, hebrew_form int) string {
	ch := C.hdate_string(C.int(string_type),
		C.int(index),
		C.int(short_form),
		C.int(hebrew_form))
	result := ""

	if ch != nil {
		result = C.GoString(ch)
		C.free(unsafe.Pointer(ch))
	}

	return result
}
