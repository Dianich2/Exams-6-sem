package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	oracle "github.com/godoes/gorm-oracle"
	"github.com/gorilla/mux"
	go_ora "github.com/sijms/go-ora/v2"
	"gorm.io/gorm"
)

type Celebrity struct {
	Id           int    `gorm:"primaryKey;column:ID" json:"id"`
	FullName     string `gorm:"column:FULLNAME" json:"fullName"`
	Nationality  string `gorm:"column:NATIONALITY" json:"nationality"`
	ReqPhotoPath string `gorm:"column:REQPHOTOPATH" json:"reqPhotoPath"`
}

var db *gorm.DB

func GetAllCelebrities(
	w http.ResponseWriter,
	r *http.Request,
) {
	var cels []Celebrity
	if err := db.Find(&cels).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cels)
}

func GetCelebrityById(
	w http.ResponseWriter,
	r *http.Request,
) {
	var cel Celebrity
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if err := db.First(&cel, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Celebrity not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cel)
}

func AddCelebrity(
	w http.ResponseWriter,
	r *http.Request,
) {
	var newCel Celebrity
	if err := json.NewDecoder(r.Body).Decode(&newCel); err != nil {
		http.Error(w, "Error in body", http.StatusBadRequest)
		return
	}

	if err := db.Create(&newCel).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCel)
}

func UpdateCelebrity(
	w http.ResponseWriter,
	r *http.Request,
) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var newCel Celebrity
	if err := json.NewDecoder(r.Body).Decode(&newCel); err != nil {
		http.Error(w, "Error in body", http.StatusBadRequest)
		return
	}

	var cel Celebrity
	if err := db.First(&cel, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Celebrity not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cel.FullName = newCel.FullName
	cel.Nationality = newCel.Nationality
	cel.ReqPhotoPath = newCel.ReqPhotoPath
	cel.Id = int(id)

	db.Save(&cel)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cel)
}

func DeleteCelebrity(
	w http.ResponseWriter,
	r *http.Request,
) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var cel Celebrity
	if err := db.First(&cel, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Celebrity not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db.Delete(&cel)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cel)
}

func main() {
	r := mux.NewRouter()
	dsn := go_ora.BuildUrl("172.24.16.1", 1521, "CEL_PDB", "cel_admin", "222", nil)
	var err error
	db, err = gorm.Open(oracle.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Celebrity{})

	r.HandleFunc("/Celebrities/All", GetAllCelebrities).Methods(http.MethodGet)
	r.HandleFunc("/Celebrities/{id}", GetCelebrityById).Methods(http.MethodGet)
	r.HandleFunc("/Celebrities", AddCelebrity).Methods(http.MethodPost)
	r.HandleFunc("/Celebrities/{id}", UpdateCelebrity).Methods(http.MethodPut)
	r.HandleFunc("/Celebrities/{id}", DeleteCelebrity).Methods(http.MethodDelete)

	log.Println("Server running on", 3000)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
