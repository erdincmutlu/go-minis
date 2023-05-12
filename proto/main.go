package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"

	"github.com/erdincmutlu/go-minis/proto/protopb"
)

func main() {
	fmt.Println("Hello World")

	erdinc := &protopb.Person{
		Name: "Erdinc",
		Age:  30,
		SocialFollowers: &protopb.SocialFollowers{
			Youtube: 1,
			Twitter: 2,
		},
	}

	data, err := proto.Marshal(erdinc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	newErdinc := &protopb.Person{}
	err = proto.Unmarshal(data, newErdinc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newErdinc.GetName())
	fmt.Println(newErdinc.GetAge())
	fmt.Println(newErdinc.SocialFollowers.GetYoutube())
	fmt.Println(newErdinc.SocialFollowers.GetTwitter())
}
