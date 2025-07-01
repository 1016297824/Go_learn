/*
tag在JSON中的应用
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  float32  `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{
		Title:  "飞驰人生",
		Year:   2019,
		Price:  52.6,
		Actors: []string{"沈腾", "尹正", "张本煜"},
	}
	fmt.Println("movie = ", movie)

	// 注意：json.Marshal()函数中参数的属性不能是私有的！
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		println("json.Marshal err:", err)
		return
	}

	//fmt.Printf("jsonStr = %s\n", jsonStr)
	fmt.Println("jsonStr = ", string(jsonStr))
}
