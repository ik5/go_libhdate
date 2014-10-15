package hdate

// #include <hdate.h>
import "C"

// libdate Heberew date struct
type Hdate_Struct struct {
	d C.hdate_struct
}

type Hdate_Julian struct {
	JD_Tishrey1           C.int
	JD_Tishrey1_Next_Year C.int
	Day                   C.int
}

type Sunrise_Sunset struct {
	Sunrise C.int
	Sunset  C.int
}

type Sun_Time struct {
	Sun_Hour    C.int
	First_Light C.int
	Talit       C.int
	Sunrise     C.int
	MidDay      C.int
	Sunset      C.int
	First_Stars C.int
	Three_Stars C.int
}
