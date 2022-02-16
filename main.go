package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func fib(num int64) int64 {
	var tempA, itr, tempB int64 = 0, 0, 1
	// Iterate until desired position in sequence.
	for ; itr < num; itr++ {
		// Use temporary variable to swap values.
		temp := tempA
		tempA = tempB
		tempB = temp + tempA
	}
	return tempA
}

func fibonacci(fib_num int64) ([]int64, int) {
	series := []int64{}
	var err int
	var ret int64
	if fib_num < 0 {
		err = -1
		return series, err
	}
	if fib_num > 0 {
		err = 0
		var itr int64 = 0
		for ; itr < fib_num; itr++ {
			ret = fib(int64(itr))
			//fmt.Printf("Fibonacci %v = %v\n", itr, ret)
			series = append(series, ret)
		}
	}
	return series, err
}

func FibonacciService(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles("fbnumber.html")
		t.Execute(w, nil)
		return
	case http.MethodPost:
		//fmt.Println("username:", r.FormValue("number"))
		r.ParseForm()
		intVar, _ := strconv.Atoi(r.FormValue("number"))
		fbSeries, err := fibonacci(int64(intVar))
		if err == 0 {
			w.Header().Set("Content-Type", "application/json")
			jsonOut, err := json.Marshal(fbSeries)
			if err != nil {
				fmt.Println("error:", err)
			}
			//fmt.Println(fbSeries)
			//fmt.Println(jsonOut)
			w.Write(jsonOut)
			return
		}
		if err == -1 {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("INVALID input - Enter a value > 0"))
		}

		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

/*
func ListFibonacciNumber(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/fbseries" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:

	case http.MethodPost:
		fmt.Println(r.Form) // print information on server side.
		fmt.Println("++ path", r.URL.Path)
		fmt.Fprintf(w, "++ Hello astaxie!") // write data to response
		fmt.Println("Number:", r.Form["number"])
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Header().Set("Content-Type", "application/json")
}
*/
func main() {
	http.HandleFunc("/fbseries", FibonacciService)
	//http.HandleFunc("/fbseries", ListFibonacciNumber)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
