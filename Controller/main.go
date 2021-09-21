package main

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Food struct{
	Name  string
	Price float64
}

var foods []Food 

func main(){
	
	foods = []Food{
		Food{ "Osh",   		 15000 },
		Food{ "Norin", 		 20000 },
		Food{ "Qozon Kabob", 35000 },
	}

	http.HandleFunc("/", FoodController)

	http.ListenAndServe(":8001", nil)

}



func FoodController(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	
	encoder := json.NewEncoder(w)

	if r.Method == "GET"{

		encoder.Encode(foods)
	
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusCreated)

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}
		
		newFood := Food{}

		json.Unmarshal(body, &newFood)
		
		foods = append(foods, newFood)
		
	}

	
	

}