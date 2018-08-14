package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
)

func testIsPrime(w http.ResponseWriter, r *http.Request) {
	testString := r.URL.Query().Get("test")
	testNumber, err := strconv.ParseInt(testString, 10, 64)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	isPrime := true
	for i := int64(2); i <= int64(math.Floor(math.Sqrt(float64(testNumber)))); i++ {
		if testNumber%i == 0 {
			isPrime = false
		}
	}

	if isPrime {
		isPrime = testNumber > 1
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	responseJSON, err := json.Marshal(isPrime)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	w.Write(responseJSON)

}

func main() {
	http.HandleFunc("/", testIsPrime)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
