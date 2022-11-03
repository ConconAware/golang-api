package test

import (
	"fmt"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

func MainTest() {

	// multiple array
	fmt.Println("main test")
	ex1 := MainUniqueAnimals()
	fmt.Printf("The example1 is %s", ex1)

	// fibonacci method
	fmt.Println()
	ex2 := MainFibonacci(30)
	fmt.Printf("The example2 is %d", ex2)

	// interface
	fmt.Println()
	fmt.Println()
	fmt.Println("The example3")
	Main3()

	//51 and 52 goroutine
	fmt.Println()
	fmt.Println("The example5")
	Main51()

	fmt.Println()
	Main52()

	//http set time out
	fmt.Println()
	fmt.Println("The example4")
	Main4()
}

type Test struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func MainTestGetApi(c *gin.Context) {
	test := []Test{
		{Id: 1, Name: "Test1"},
		{Id: 2, Name: "Test2"},
	}

	fmt.Println(test)
	c.IndentedJSON(http.StatusOK, test)

}

func MainTestGetApiID(c *gin.Context) {
	intVar, _ := strconv.ParseInt(c.Param("id"), 0, 64)
	test := []Test{
		{Id: 1, Name: "Test1"},
		{Id: 2, Name: "Test2"},
	}
	var result []Test
	for _, t := range test {
		if intVar == t.Id {
			result = append(result, t)
		}
	}

	fmt.Println(test)
	c.IndentedJSON(http.StatusOK, result)

}

func MainTestPostApi(c *gin.Context) {
	var body Test

	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&body); err != nil {
		return
	}

	test := []Test{
		{Id: 1, Name: "Test1"},
		{Id: 2, Name: "Test2"},
	}
	fmt.Println(body)
	test = append(test, body)

	fmt.Println(test)
	c.IndentedJSON(http.StatusOK, test)

}
