package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println("What is your name?");
	// reader := bufio.NewReader(os.Stdin)
	// name, err := reader.ReadString('\n')
	// if err == nil {
	// 	fmt.Println("Hello", name)
	// }else {
	// 	log.Fatal(err);
	// }

	//Variable practice
	// var vName string = "Conner";
	// var v1, v2 = 1.2, 3.4;
	// var v3 = "hello"
	// v4 := 2.4;

	//Datatypes 
	//int, float64, bool, string, rune
	
	// fmt.Println(reflect.TypeOf(25))
	// fmt.Println(reflect.TypeOf(3.14))
	// fmt.Println(reflect.TypeOf(true))
	// fmt.Println(reflect.TypeOf("hello"))
	// fmt.Println(reflect.TypeOf(2))

	// cV1 := 1.5
	// cV2 := int(cV1)

	// cV3 := "50000000"
	// cV4, err := strconv.Atoi(cV3)
	// fmt.Println(cV4, err, reflect.TypeOf(cV4))

	//Int to float test

	// test := 50
	// if test2, err := strconv.ParseFloat(test, 64); err == nil
	// {
	// 	fmt.Println(test2)
	// }

	//Conditionals and Logical Operators

	// iAge := 8

	// if (iAge >= 1) && (iAge <= 18)
	// {
	// 	fmt.Println("Important Birthday")
	// }else if (iAge == 21) || (iAge == 50)
	// {
	// 	fmt.Println("Important Birthday")
	// }else if iAge >= 65
	// {
	// 	fmt.Println("Important Birthday")
	// }else {
	// 	fmt.Println("Not Important Birthday")
	// }

	s1 := "A Word"
	replacer := strings.NewReplacer("A", "Another")
	s2 := replacer.Replace(s1)
	fmt.Println(s2)

	fmt.Println(len(s2))
	fmt.Println("Contains this:", strings.Contains(s2, "Another"))
	fmt.Println("This index of item", strings.Index(s2, "o"))
	fmt.Println("Replace item", strings.Replace(s2, "o", "0", -1))
	sV3 := "\nSome Words\n" 
	sV3 = strings.TrimSpace(sV3)
	
	fmt.Println("Split : ", strings.Split("a-b-c-d", "-"))
	fmt.Println("Lower: ", strings.toLower(s2))
}