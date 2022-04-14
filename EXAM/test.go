// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
// 		rw.Header().Set("Content-Type", "text/html")
// 		fmt.Fprintf(rw, "HELLO FROM GO SRV [%s]", req.URL)
// 	})

// 	http.ListenAndServe(":3000", nil)
// // }
// package main

// import "fmt"

// func main() {
// 	var intSlice = []int{10, 20, 30, 40}

// 	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))

package main

import (
	"io"
	"net/http"
)

type Task struct {
	Description string
	Done        bool
}

var tasks = []Task{
	{
		Description: "Faire courses",
		Done:        false,
	},
	{
		Description: "Payer factures",
		Done:        false,
	},
}

func main() {
	list := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "HandleFunc 1")
	}

	done := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "HandleFunc 2")
	}

	add := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "HandleFunc 3")
	}
	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)
	http.ListenAndServe("localhost:8080", nil)

}
