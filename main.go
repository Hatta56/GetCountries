package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struktur untuk data Negara
type Country struct {
	Name      string   `json:"name"`
	Region    string   `json:"region"`
	Timezones []string `json:"timezones"`
}

func main() {
	// Menentukan route untuk endpoint API
	http.HandleFunc("/countries", GetCountries)

	// Menjalankan server HTTP pada port 8080
	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Handler untuk endpoint /countries
func GetCountries(w http.ResponseWriter, r *http.Request) {
	// Mengambil data JSON dari URL publik
	url := "https://gist.githubusercontent.com/herysepty/ba286b815417363bfbcc472a5197edd0/raw/aed8ce8f5154208f9fe7f7b04195e05de5f81fda/coutries.json"
	response, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Mengurai data JSON
	var countries []Country
	err = json.NewDecoder(response.Body).Decode(&countries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengambil parameter yang diperlukan (name, region, timezones)
	var result []map[string]interface{}
	for _, country := range countries {
		result = append(result, map[string]interface{}{
			"name":      country.Name,
			"region":    country.Region,
			"timezones": country.Timezones,
		})
	}

	// Mengembalikan data JSON sebagai respons
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
