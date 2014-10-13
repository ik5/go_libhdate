package hdate

// #include <hdate.h>
import "C"

// libdate Heberew date struct
type Hdate_Struct struct {
	d C.hdate_struct
}
