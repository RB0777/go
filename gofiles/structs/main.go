package main

import(
	
	

	"encoding/json"
	
	
	"net/http"
	"github.com/L04DB4L4NC3R/getgoing/section4/todo-api/views"
	
)

	func main(){
		mux := http.NewServeMux()
		mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
			if r.Method == http.MethodGet {
				data := structs.Response{
					Code: http.StatusOK,
					Body: "pong",
				}

				json.NewEncoder(w).Encode(data)
			}
		})
		http.ListenAndServe(":4000", mux)


	}

