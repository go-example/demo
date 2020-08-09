package main

import "time"

func main()  {
	Layout := "2006-01-02 15:04:05"
	t := time.Now().Format(Layout)
	println(t)

	loc, _ := time.LoadLocation("Local")
	tm,_ := time.ParseInLocation(Layout,"2019-10-10 00:00:00",loc)
	t = tm.Format("2006-01-02 15:04:05")
	println(t)


	tm,_ = time.Parse(Layout,"2019-10-10 00:00:01")
	t = tm.Format("2006-01-02 15:04:05")
	println(t)

	println(time.Now().Day())
	println(time.Now().Month())
	println(time.Now().YearDay())
	y,m,d := time.Now().Date()
	println(y,m,d)
}
