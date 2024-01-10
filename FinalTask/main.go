package main

import (
	"encoding/json"
	"net/http"
)








// Define the API handlers
func createTemplateHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a Template
	var template Template
	err := json.NewDecoder(r.Body).Decode(&template)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new DB and call CreateTemplate
	// db := &DB{}
	// err = db.CreateTemplate(template.Key, template)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// // Respond with a success message
	// fmt.Fprint(w, "Template created successfully")
}

// Similar handlers would be implemented for UpdateTemplate, DeleteTemplate, Refresh and Test

func main() {
	http.HandleFunc("/createTemplate", createTemplateHandler)
	// Similar routes would be added for UpdateTemplate, DeleteTemplate, Refresh and Test
	http.ListenAndServe(":8080", nil)
}


