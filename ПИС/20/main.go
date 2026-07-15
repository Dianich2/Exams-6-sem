// нужно установить пакет драйвер для нужной для вас СУБД
// у меня мой любимый Oracle (go get github.com/sijms/go-ora/v2)
// ну и еще для удобства написания сервера поставим go get github.com/gorilla/mux

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/sijms/go-ora/v2"
)

type Celebrity struct {
	Id           int    `json:"id"`
	FullName     string `json:"fullName"`
	Nationality  string `json:"nationality"`
	ReqPhotoPath string `json:"reqPhotoPath"`
}

var db *sql.DB

// // Функции для работы с БД
func GetAllCelebritiesFromDB(
	ctx context.Context,
	db *sql.DB,
) ([]Celebrity, error) {
	rows, err := db.QueryContext(
		ctx,
		`Select id,
		        fullName,
				nationality,
				reqPhotoPath
				from Celebrities
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cels := make([]Celebrity, 0)
	for rows.Next() {
		var cel Celebrity
		err = rows.Scan(
			&cel.Id,
			&cel.FullName,
			&cel.Nationality,
			&cel.ReqPhotoPath,
		)

		if err != nil {
			return nil, err
		}

		cels = append(cels, cel)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cels, nil
}

func AddCelebrityToDB(
	ctx context.Context,
	db *sql.DB,
	cel *Celebrity,
) (*Celebrity, error) {
	var id int
	_, err := db.ExecContext(
		ctx,
		`insert into Celebrities (
		    fullName,
			nationality,
			reqPhotoPath
		)
		values(:1, :2, :3)
		returning id into :4`,
		cel.FullName,
		cel.Nationality,
		cel.ReqPhotoPath,
		&id,
	)

	if err != nil {
		return nil, err
	}

	cel.Id = int(id)
	return cel, nil
}

func GetCelebrityByIdFromDB(
	ctx context.Context,
	db *sql.DB,
	id int,
) (*Celebrity, error) {
	var cel Celebrity
	row := db.QueryRowContext(
		ctx,
		`select id,
		        fullName,
				nationality,
				reqPhotoPath
				from Celebrities
				where id = :1
		`,
		id,
	)

	err := row.Scan(
		&cel.Id,
		&cel.FullName,
		&cel.Nationality,
		&cel.ReqPhotoPath,
	)

	if err != nil {
		return nil, err
	}

	return &cel, nil
}

func UpdateCelebrityByIdInDB(
	ctx context.Context,
	db *sql.DB,
	id int,
	newCel *Celebrity,
) (*Celebrity, error) {
	res, err := db.ExecContext(
		ctx,
		`update Celebrities set fullName = :1,
		                        nationality = :2,
								reqPhotoPath = :3
								where id = :4
		`,
		newCel.FullName,
		newCel.Nationality,
		newCel.ReqPhotoPath,
		id,
	)

	if err != nil {
		return nil, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("cel not found")
	}

	newCel.Id = id
	return newCel, nil
}

func DeleteCelebrityByIdFromDB(
	ctx context.Context,
	db *sql.DB,
	id int,
) (*Celebrity, error) {
	cel, err := GetCelebrityByIdFromDB(ctx, db, id)
	if err != nil {
		return nil, err
	}

	res, err := db.ExecContext(
		ctx,
		`delete from Celebrities
			    where id = :1 
		`,
		id,
	)

	if err != nil {
		return nil, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, errors.New("cel not found")
	}

	return cel, nil
}

// // Обработчики запросов
func GetAllCelebrities(
	w http.ResponseWriter,
	r *http.Request,
) {
	cels, err := GetAllCelebritiesFromDB(r.Context(), db)
	if err != nil {
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
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	cel, err := GetCelebrityByIdFromDB(r.Context(), db, id)

	if err != nil {
		http.Error(w, "Celebrity not found", http.StatusNotFound)
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
		http.Error(w, "Error in body", http.StatusInternalServerError)
		return
	}

	cel, err := AddCelebrityToDB(r.Context(), db, &newCel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cel)
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
		http.Error(w, "Error in body", http.StatusInternalServerError)
		return
	}

	cel, err := UpdateCelebrityByIdInDB(r.Context(), db, id, &newCel)

	if err != nil {
		if err.Error() == "cel not found" {
			http.Error(w, "Celebrity not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	cel, err := DeleteCelebrityByIdFromDB(r.Context(), db, id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			http.Error(w, "Celebrity not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cel)
}

func main() {
	r := mux.NewRouter()
	const conn string = "oracle://cel_admin:222@localhost:1521/CEL_PDB"
	var err error
	db, err = sql.Open("oracle", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
