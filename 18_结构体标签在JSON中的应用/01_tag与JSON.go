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

	// 编码过程：结构体 -> JSON
	// 注意：json.Marshal()函数中参数的属性不能是私有的！
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		println("json.Marshal err:", err)
		return
	}
	//fmt.Printf("jsonStr = %s\n", jsonStr)
	fmt.Println("jsonStr = ", string(jsonStr))

	// 解码过程：JSON -> 结构体
	// jsonStr =  {"title":"飞驰人生","year":2019,"price":52.6,"actors":["沈腾","尹正","张本煜"]}
	movie1 := Movie{}
	err1 := json.Unmarshal(jsonStr, &movie1)
	if err1 != nil {
		println("json.Unmarshal err:", err1)
		return
	}
	fmt.Println("movie = ", movie1)
}
