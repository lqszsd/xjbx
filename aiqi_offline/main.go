package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {
	conn, err := NewDb()
	if err != nil {
		panic(err)
	}
	db = conn
}
func main() {
	db,err:=NewDb()
	errFix(err)
	for{
		result:=db.DB().Ping()
		fmt.Println(result)
	}

}
func errFix(err error) {
	if err != nil {
		panic(err)
	}
	println(err)
}
func httpGet() string {
	resp, err := http.Get("http://47.106.92.111:9501/home/index/index")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body)
}
