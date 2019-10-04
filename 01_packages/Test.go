package main

import ("fmt"; "math/rand"; "time")

func main(){
	var bytes int
	bytes = rand.Intn(100)
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(bytes)
	println("rrrrrrrrkhbjhvrrrrr")
}

