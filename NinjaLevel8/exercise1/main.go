package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users) //lets see what they look like.

	byteX, err := json.Marshal(users)
	if err != nil {
		fmt.Println("PANIKKKK!!111!! There is an error!!!", err)
	}

	fmt.Println(string(byteX))

}
