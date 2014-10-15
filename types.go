package hdate

// #include <hdate.h>
import "C"

// libdate Heberew date struct
type Hdate_Struct struct {
	d *C.hdate_struct
}

type Hdate_Julian struct {
	JD_Tishrey1           int
	JD_Tishrey1_Next_Year int
	Day                   int
}

type Sunrise_Sunset struct {
	Sunrise int
	Sunset  int
}

type Sun_Time struct {
	Sun_Hour    int
	First_Light int
	Talit       int
	Sunrise     int
	MidDay      int
	Sunset      int
	First_Stars int
	Three_Stars int
}
