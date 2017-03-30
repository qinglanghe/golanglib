package dateutils

import (
	"time"
	"strings"
)

//2006-01-02 15:04:05
const (
	yyyy = "2006"
	MM = "01"
	dd = "02"
	HH = "15"
	mm = "04"
	ss = "05"
)

func getPattern(pattern string) string {
	s := strings.Replace(pattern,"yyyy",yyyy,1)
	s = strings.Replace(s,"MM",MM,1)
	s = strings.Replace(s,"dd",dd,1)
	s = strings.Replace(s,"HH",HH,1)
	s = strings.Replace(s,"mm",mm,1)
	s = strings.Replace(s,"ss",ss,1)
	return s
}

func Parse(pattern string,str string) time.Time  {
	ret,_ := time.Parse(getPattern(pattern),str)
	return ret
}

func Format(pattern string,t time.Time) string  {
	return t.Format(getPattern(pattern))
}
