package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/configor"
	"github.com/naveensrinivasan/instacart-delivery-availability/pkg/instacart"
)

type configuration struct {
	Email    string
	Password string
	Stores   []string
}

func main() {
	config := configuration{}
	e := configor.Load(&config, "settings.yaml")
	if e != nil {
		log.Fatal(e)
	}

	i, err := instacart.NewInstacart(config.Email, config.Password)
	if err != nil {
		log.Fatal(err)
	}
	for _, store := range config.Stores {
		result, err := i.DeliveryAvailable(store)
		if err != nil {
			log.Fatal(err)
		}
		if result {
			fmt.Println(fmt.Sprintf("There are delivery slots available %s", store))
		}
	}
}
