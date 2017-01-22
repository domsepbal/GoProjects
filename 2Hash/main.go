package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Get a book
	res, err := http.Get("http://www.gutenber.org/files/2701/old/moby10b.txt")
	if err == nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	//Scan the page
	scanner := bufio.NewScanner(res.Body)
	//Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	//Create a slice to hold counts
	buckets := make([]int, 200)
	//Loop over words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}

	fmt.Println(buckets[65:123])
	fmt.Println("**************")
	for i := 28; i <= 126; i++ {
		fmt.Printf("%v - %c - %v \n", i, i, buckets[i])
	}

}

func HashBucket(word string, bucket int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}

	return sum % bucket
}
