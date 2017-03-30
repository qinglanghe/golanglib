package main

import (
	"fmt"
	"time"
	"lang/common/utils/dateutils"
)

type  User struct  {
	Name string `col:"name" json:"username"`
	Age int64 `col:"age"`
	D time.Time `col:"created_time"`
}

func (u *User) ToTime(str string) time.Time {
	return dateutils.Parse("yyyy-MM-dd HH:mm:ss",str)
}

func (u *User) String()  {
	fmt.Println("user string ")
}

func test(f interface{})  {

}


func main() {
	u := User{}
	println(u.Name == nil)
}
