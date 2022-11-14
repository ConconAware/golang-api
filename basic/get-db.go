package test

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Itest interface {
	SetString(id int64, name string)
	GetString() (User, error)
}

type User struct {
	id   int64
	name string
}

func (u *User) SetString(id int64, name string) {
	u.id = id
	u.name = name
}

func (u User) GetString() (User, error) {
	return u, nil
}

func GetDataDbMongo(c *gin.Context) {
	fmt.Println("GetDataDbMongo")

	// acc.SetString(12, "5555")
	fmt.Println(User{5, "Jele"})

	fmt.Println("555555555555555")
}
