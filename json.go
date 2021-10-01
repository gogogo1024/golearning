package main

import (
	"encoding/json"
	"fmt"
)


type Book struct {
	Title string `json:"title"`
	Authors []string `json:"authors"`
	Publisher string `json:"publisher"`
	IsPublished bool `json:"is_published"`
	Price float64 `json:"price"`
}

func main() {
	goBook := Book{
		Title:       "Go语言编程",
		Authors:     []string{"XuShiWei","XuDaoli"},
		Publisher:   "ituring.com.cn",
		IsPublished: true,
		Price:       9.9,
	}
	b,err := json.Marshal(goBook)
	if err != nil {
		fmt.Println("error: json.Marshal")
	}
	fmt.Printf("%s\n",b)
	var book Book
	a :=[]byte(`{"Title":"Go语言编程","Sales":100000}`)
	error :=json.Unmarshal(a,&book)
	if error != nil {
		fmt.Println("error: json.Unmarshal")
	}
	fmt.Printf("%s\n",book)

	var  r interface{}
	err1 :=json.Unmarshal(b,&r)

	if err1 != nil {
		fmt.Println("error: json.Unmarshal")
	}
	mapResult,ok:=r.(map[string] interface{})
	if ok {
		for k,v := range mapResult {
			switch v2:=v.(type) {
			case string:
				fmt.Println(k,"is string",v2)
			case int:
				fmt.Println(k,"is int",v2)
			case bool:
				fmt.Println(k,"is bool",v2)
			case []interface{}:
				fmt.Println(k,"is an array:")
				for i,iv := range v2 {
					fmt.Println(i,iv)
				}
			case float64:
				fmt.Println(k,"is float64",v2)
			default:
				fmt.Println(k,"is another type not handle yet")
			}
		}
		
	}
	fmt.Printf("%s\n",r)



}