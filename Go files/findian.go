package main

import (
    "fmt"
    "strings"
)

func main()	{
    var str string
    fmt.Println("Please enter a string:\n")
    fmt.Scanln(&str)
    if strings.Contains(str,"a") && (str[0]=='i' || str[0]=='I') && (str[len(str)-1]=='n' || str[len(str)-1]=='N'){
        fmt.Println("Found!")
	}else{
        fmt.Println("Not Found!")
    }
}