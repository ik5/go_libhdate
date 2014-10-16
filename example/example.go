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

package main

import (
	"fmt"
	"github.com/ik5/go_libhdate"
	"time"
)

func basic() {
	h := hdate.Init()
	defer h.Destruct()

	// Gregorian date
	fmt.Println("Today is: ")
	fmt.Printf("%d/%d/%d\n", h.Get_GDay(), h.Get_GMonth(), h.Get_GYear())

	// print hebrew date: 0 - israely holidays, 0 - long format
	fmt.Printf("%s\n\n", h.Get_Format_Date(0, 0))

}

type City_Struct struct {
	Latitude  float64
	Longitude float64
	TimeZone  int
	Name      string
}

var Cities = [...]City_Struct{
	{Latitude: 31.78, Longitude: 35.22, TimeZone: 2 * 60, Name: "Jerusalem"},
	{Latitude: 32.07, Longitude: 34.77, TimeZone: 2 * 60, Name: "Tel Aviv-Jafa"},
	{Latitude: 32.82, Longitude: 34.99, TimeZone: 2 * 60, Name: "Hifa"},
	{Latitude: 31.96, Longitude: 34.80, TimeZone: 2 * 60, Name: "Rishon Lezion"},
	{Latitude: 31.80, Longitude: 34.64, TimeZone: 2 * 60, Name: "Ashdod"},
	{Latitude: 31.25, Longitude: 34.80, TimeZone: 2 * 60, Name: "Be'er Sheva"},
	{Latitude: 32.09, Longitude: 34.88, TimeZone: 2 * 60, Name: "Petach Tiqva"},
	{Latitude: 32.33, Longitude: 34.86, TimeZone: 2 * 60, Name: "Netanya"},
	{Latitude: 32.02, Longitude: 34.76, TimeZone: 2 * 60, Name: "Holon"},
	{Latitude: 32.09, Longitude: 34.85, TimeZone: 2 * 60, Name: "B'ene Beraq"},
	{Latitude: 32.02, Longitude: 34.75, TimeZone: 2 * 60, Name: "Bat Yam"},
	{Latitude: 32.08, Longitude: 34.80, TimeZone: 2 * 60, Name: "Ramat Gan"},
	{Latitude: 31.67, Longitude: 34.56, TimeZone: 2 * 60, Name: "Ashqelon"},
	{Latitude: 31.89, Longitude: 34.80, TimeZone: 2 * 60, Name: "Rehovot"},
	{Latitude: 32.17, Longitude: 34.84, TimeZone: 2 * 60, Name: "Herzeliyya"},
	{Latitude: 32.19, Longitude: 34.91, TimeZone: 2 * 60, Name: "Kfar Saba"},
	{Latitude: 32.45, Longitude: 34.92, TimeZone: 2 * 60, Name: "Hadera"},
	{Latitude: 32.19, Longitude: 34.88, TimeZone: 2 * 60, Name: "Ra'anana"},
	{Latitude: 31.96, Longitude: 34.90, TimeZone: 2 * 60, Name: "Lod"},
	{Latitude: 31.93, Longitude: 34.86, TimeZone: 2 * 60, Name: "Ramla"},
}

var dayLightSaving = 1 * 60

func city_info() {
	h := hdate.Init()
	defer h.Destruct()

	now := time.Now()
	//  _, zone := now.Zone()

	fmt.Printf("Current Time: %02d:%02d:%02d\n", now.Hour(), now.Minute(), now.Second())

	for _, v := range Cities {
		long := v.Longitude
		lat := v.Latitude
		name := v.Name

		fmt.Printf("City: %s\n", name)
		fmt.Printf("\tLatitude: %2.2f\n", lat)
		fmt.Printf("\tLongitude: %2.2f\n", long)

		// print Time using a degree of sunrise/sunset
		deg := 90
		fmt.Printf("\tDegree: %d\n", deg)

		sun := hdate.Get_UTC_Sun_Time_Deg(now.Day(), int(now.Month()),
			now.Year(), long, lat, float64(deg))
		//    sunrise := int64(sun.Sunrise + zone * 60)
		//    sunset  := int64(sun.Sunset + zone * 60)

		t := sun.Sunrise + v.TimeZone + dayLightSaving
		fmt.Printf("\tHours of sunrise: %02d:%02d\n", t/60, t%60)
		t = sun.Sunset + v.TimeZone + dayLightSaving
		fmt.Printf("\tHours of sunset: %02d:%02d\n", t/60, t%60)

		full_time := hdate.Get_UTC_Sun_Time_Full(now.Day(), int(now.Month()),
			now.Year(), long, lat)

		fmt.Printf("\tTemporary hour length: %.02f\n", full_time.Sun_Hour)

		t = full_time.First_Light + v.TimeZone + dayLightSaving
		fmt.Printf("\tFirst light: %02d:%02d\n", t/60, t%60)

		t = full_time.Talit + v.TimeZone + dayLightSaving
		fmt.Printf("\tTalit time: %02d:%02d\n", t/60, t%60)

		t = full_time.Sunrise + v.TimeZone + dayLightSaving
		fmt.Printf("\tFull Sunrise: %02d:%02d\n", t/60, t%60)

		t = full_time.Sunset + v.TimeZone + dayLightSaving
		fmt.Printf("\tFull Sunset: %02d:%02d\n", t/60, t%60)

		t = full_time.First_Stars + v.TimeZone + dayLightSaving
		fmt.Printf("\tFirst stars: %02d:%02d\n", t/60, t%60)

		t = full_time.Three_Stars + v.TimeZone + dayLightSaving
		fmt.Printf("\tThree stars: %02d:%02d\n", t/60, t%60)

		fmt.Println("")
	}
}

func main() {
	basic()
	city_info()
}
