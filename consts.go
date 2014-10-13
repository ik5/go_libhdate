package hdate

const (
	// use diaspora dates and holydays flag
	DIASPORA = -1
	// use israel dates and holydays flag
	ISRAEL_FLAG = 0

	// use short strings flag
	SHORT_FLAG = -1
	// use long strings flag
	LONG_FLAG = 0

	/*
	  for hdate_get_int_string_ and hdate_get_int_wstring

	  How large should the buffer be? Hebrew year 10,999 would
	  be י'תתקצ"ט, eight characters, each two bytes, plus an
	  end-of-string delimiter, equals 17. This could effectively
	  yield a range extending to Hebrew year 11,899, י"א תתצ"ט,
	  due to the extra ק needed for the '900' century. However,
	  for readability, I would want a an extra space at that
	  point between the millenium and the century...

	*/
	HEBREW_NUMBER_BUFFER_SIZE  = 17
	HEBREW_WNUMBER_BUFFER_SIZE = 9

	// for function hdate_string: identifies string type: integer
	STRING_INT = 0

	// for function hdate_string: identifies string type: day of week
	STRING_DOW = 1

	// for function hdate_string: identifies string type: parasha
	STRING_PARASHA = 2

	// for function hdate_string: identifies string type: hebrew_month
	STRING_HMONTH = 3

	// for function hdate_string: identifies string type: gregorian_month
	STRING_GMONTH = 4

	// for function hdate_string: identifies string type: holiday
	STRING_HOLIDAY = 5

	// for function hdate_string: identifies string type: holiday
	STRING_OMER = 6

	// for function hdate_string: use short form, if one exists
	STRING_SHORT = 1

	// for function hdate_string: use long form
	STRING_LONG = 0

	// for function hdate_string: use embedded hebrew string
	STRING_HEBREW = 1

	// for function hdate_string: use local locale string
	STRING_LOCAL = 0

	HOLIDAY_NORMAL_DAY         = 0
	HOLIDAY_YOM_TOV            = 1
	HOLIDAY_EREV_YOM_KIPPUR    = 2
	HOLIDAY_HOL_HAMOED         = 3
	HOLIDAY_HANUKA             = 4
	HOLIDAY_PURIM              = 4
	HOLIDAY_TZOMOT             = 5
	HOLIDAY_INDEPENDACE_DAY    = 6
	HOLIDAY_JERUSALEM_DAY      = 6
	HOLIDAY_LAG_BAOMER         = 7
	HOLIDAY_TU_BEAV            = 7
	HOLIDAY_TU_BESVAT          = 7
	HOLIDAY_IDF_MEMORIAL       = 8
	HOLIDAY_HOLOCAUST_MEMORIAL = 8
	HOLIDAY_NATIONAL_DAY       = 9
)
