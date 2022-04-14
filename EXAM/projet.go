//Examen Golang pour DevOps
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// un type Task, qui est une struct avec deux champs: • ”Description”, de type string • ”Done”, de type bool
type Task struct {
	Description string
	Done        bool
	ID          string
}

//variable globale ”task” qui est une slice de Task.
var tasks = []Task{
	{
		Description: "Faire ses courses",
		Done:        false,
		ID:          "0",
	},
	{
		Description: "Payer ses amendes",
		Done:        false,
		ID:          "1",
	},
}

//trois HandleFunc

func listHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "LIST")
}

func doneHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "DONE")
}

func addHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "ADD")
}

// func list(rw http.ResponseWriter, _ *http.Request)

func main() {

	list := func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Toutes les tasks pas encore finis
		for _, task := range tasks {
			if task.Done == false {
				data := ([]byte(" l'ID est : " + task.ID + " Tasks: " + task.Description))
				w.Write([]byte(data))
			}
		}
	}

	done := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "done")
		//filtrer les tasks finis
		switch r.Method {
		case http.MethodGet:
			for _, task := range tasks {
				if task.Done == true {
					data := ([]byte(" l'ID est : " + task.ID + " Tasks: " + task.Description))
					w.Write([]byte(data))
				}
			}
		}
		//	+ Taches via POST

	}

	add := func(req http.ResponseWriter, r *http.Request) {
		io.WriteString(req, "add")
		if r.Method != http.MethodPost {
			req.WriteHeader(http.StatusBadRequest)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error  body: %v", err)
			http.Error(
				req,
				"No read body", http.StatusBadRequest)
			return
		}
		description := string(body)
		tasks = append(tasks, Task{description, false, "1"})
		req.WriteHeader(http.StatusOK)
	}

	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)
	//fonction ListenAndServe

	http.ListenAndServe("localhost:8080", nil)

}
