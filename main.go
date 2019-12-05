package main

import (
	"fmt"

	"github.com/leonistor/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Nutzi",
		LastName:  "Porcoasa",
	}
	fmt.Println(u)
}
