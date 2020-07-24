package example

import (
    "fmt"
    "github.com/codecov/example-go/awesome"
)

var result string

func Setup() {

    // Comment

    result = awesome.Smile()
    fmt.Println("dev")

}

func GetResult() string {

    /*
    Comment
    */

    return result

}
