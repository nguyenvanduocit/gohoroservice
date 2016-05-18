package main

import (
	"github.com/nguyenvanduocit/gohoro"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"flag"
	"fmt"
)


type Response struct {
	Success bool `json:"success"`
	Content string `json:"content"`
}

func DailyHoro(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sign := vars["sign"]
	result, err := gohoro.GetHoroscope(sign)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Success:false, Content:"Can not get data."})
	}else if(result == ""){
		errorMessage := Response{Success:false, Content:"Sign not found"}
		json.NewEncoder(w).Encode(errorMessage)

	}else{
		json.NewEncoder(w).Encode(Response{Success:true, Content:result})
	}
}

func main() {
	var ip, port string
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip")
	flag.StringVar(&port, "port", "8080", "Port")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", ip, port)
	fmt.Println("Server is listen on ", address);
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/{sign}/daily", DailyHoro)

	log.Fatal(http.ListenAndServe(address, router))
}
