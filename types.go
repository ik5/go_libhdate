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
