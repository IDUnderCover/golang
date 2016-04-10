package main

import "fmt"

func main() {
    /*
    str := "Hello World"
    n := len(str)
    for i := 0; i < n; i++ {
        ch := str[i]
        fmt.Println(i ,ch)   
    }
    fmt.Println("Hello world")

    var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
    var mySlice []int = myArray[:5]
    fmt.Println("Elements of myArray: ")

    mySlice[1] = 111
    for _, v := range mySlice{
        v = v + 1
        fmt.Println(v, " ")    
    }
    for _, v := range myArray{
        fmt.Println(v, " ")    
    }*/
    test_make()
    mymap := make(map[string] PersonInfo, 100)
    mymap["1234"] = PersonInfo{"1","Jack"}
    fmt.Println(mymap)
    fmt.Println(is_existed(mymap, "1231"))

    test_many_args(1,2,3,4)

}

func test_many_args(arg ...int){
    for  i,v := range arg{
        fmt.Println(i,v)
        }
    }

func test_make(){
    a := make([]int, 5)
    // b := make(int, 5)
    print_list(a)
    // print_list(b)
    }

func print_list(slice []int) (res int, err error){
        for i,v := range slice{
            fmt.Println(i,v)
            }
        return 1,nil
    }

type PersonInfo struct{
    ID string
    Name string 
    }

func is_existed(mymap map[string] PersonInfo, key string) (res bool, err error){
    _, ok := mymap[key]
    return ok,nil
    }
