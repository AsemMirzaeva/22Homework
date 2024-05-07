package main

import (
	g "git/git"
	"log"
	"os"
)

func main(){
	file, err := os.Open("text.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	username, err := g.GetUsername()
	if err != nil{
		log.Fatal(err)
	}
	_, err = file.WriteString(username + "\n")
	if err != nil{
		log.Fatal(err)
	}
	
}