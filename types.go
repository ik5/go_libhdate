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
