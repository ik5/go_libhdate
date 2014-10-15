hdate
=====
The following Go package is a binding for the C library named libhdate in Go.

Example
-------
I have provided an example file (At the moment, it does not calculate time
properly).

To install the package:
   go get github.com/ik5/go_libhdate

Simple usage:
   package main

   import (
           "fmt"
           "github.com/ik5/go_libhdate"
   )

   func main() {
        h := hdate.Init()
        defer h.Destruct()

        // Gregorian date
        fmt.Println("Today is:")
        fmt.Printf("%d, %d, %d\n", h.Get_GDay(), h.Get_GMonth(), h.Get_GYear())

        // print hebrew date: 0 - israely holidays, 0 - long format
        fmt.Printf("%s\n\n", h.Get_Format_Date(0, 0))

   }

License
-------
The following package is provided as LGPL v3. libhdate is GPLi v3.

    Copyright (C) 2014  Ido Kanner <idokan at@at gmail dot.dot com>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Lesser General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Lesser General Public License for more details.


