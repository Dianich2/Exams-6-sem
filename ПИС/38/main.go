package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "38/docs"

	"github.com/gorilla/mux"
	_ "github.com/sijms/go-ora/v2"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Celebrity struct {
	Id           int    `json:"id"`
	FullName     string `json:"fullName"`
	Nationality  string `json:"nationality"`
	ReqPhotoPath string `json:"reqPhotoPath"`
}

var db *sql.DB

// Функции для работы с БД
func GetAllCelebritiesFromDB(ctx context.Context, db *sql.DB) ([]Celebrity, error) {
	rows, err := db.QueryContext(
		ctx,
		"Select id, fullName, nationality, NVL(reqPhotoPath, ' ') from Celebrities",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := []Celebrity{}
	for rows.Next() {
		var c Celebrity
		err = rows.Scan(
			&c.Id,
			&c.FullName,
			&c.Nationality,
			&c.ReqPhotoPath,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func AddCelebrityToDB(ctx context.Context, db *sql.DB, c *Celebrity) (*Celebrity, error) {
	_, err := db.ExecContext(
		ctx,
		`Insert into Celebrities (fullName, nationality, reqPhotoPath)
		 values(:1, :2, :3)`,
		c.FullName,
		c.Nationality,
		c.ReqPhotoPath,
	)

	if err != nil {
		return nil, err
	}

	var id int
	err = db.QueryRowContext(ctx,
		`SELECT MAX(id) FROM Celebrities`,
	).Scan(&id)

	c.Id = id
	return c, nil
}

func GetCelebrityByIdFromDB(ctx context.Context, db *sql.DB, id int) (*Celebrity, error) {
	var c Celebrity
	err := db.QueryRowContext(
		ctx,
		`Select id, fullName, nationality, NVL(reqPhotoPath, ' ') 
		from Celebrities WHERE id= :1`, id,
	).
		Scan(
			&c.Id,
			&c.FullName,
			&c.Nationality,
			&c.ReqPhotoPath,
		)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func UpdateCelebrityInDB(ctx context.Context, db *sql.DB, id int, c *Celebrity) (int64, error) {
	res, err := db.ExecContext(
		ctx,
		`Update Celebrities set fullname= :1, nationality= :2,
		 reqPhotoPath= :3 where id= :4`,
		c.FullName,
		c.Nationality,
		c.ReqPhotoPath,
		id,
	)

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteCelebrityFromDB(ctx context.Context, db *sql.DB, id int) (int64, error) {
	res, err := db.ExecContext(
		ctx,
		`Delete from Celebrities where id= :1`,
		id,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Обработчики запросов

// GetAllCelebrities godoc
// @Summary Получить всех celebrities
// @Description Возвращает всех celebrities
// @Tags Celebrities
// @Produce json
// @Success 200 {array} Celebrity
// @Failure 500 {string} string "Internal server error"
// @Router /Celebrities/All [get]
func GetAllCelebrities(w http.ResponseWriter, r *http.Request) {
	cels, err := GetAllCelebritiesFromDB(r.Context(), db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cels)
}

// GetCelebrityById godoc
// @Summary Получить celebrity по id
// @Description Возвращает celebrity по id
// @Tags Celebrities
// @Produce json
// @Param id path int true "Celebrity ID"
// @Success 200 {object} Celebrity
// @Failure 404 {string} string "Celebrity not found"
// @Router /Celebrities/{id} [get]
func GetCelebrityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	cel, err := GetCelebrityByIdFromDB(r.Context(), db, id)
	if err != nil {
		http.Error(w, "Celebrity not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cel)
}

// AddCelebrity godoc
// @Summary Добавить celebrity
// @Description Новый celebrity
// @Tags Celebrities
// @Accept json
// @Produce json
// @Param celebrity body Celebrity true "Celebrity data"
// @Success 201 {object} Celebrity
// @Failure 500 {string} string "Server error"
// @Router /Celebrities [post]
func AddCelebrity(w http.ResponseWriter, r *http.Request) {
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

// UpdateCelebrity godoc
// @Summary Обновить celebrity
// @Description Обновляет celebrity
// @Tags Celebrities
// @Accept json
// @Produce json
// @Param id path int true "Celebrity ID"
// @Param celebrity body Celebrity true "Celebrity data"
// @Success 200 {object} Celebrity
// @Failure 404 {string} string "Celebrity not found"
// @Failure 500 {string} string "Server error"
// @Router /Celebrities/{id} [put]
func UpdateCelebrity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var newCel Celebrity
	if err := json.NewDecoder(r.Body).Decode(&newCel); err != nil {
		http.Error(w, "Error in body", http.StatusInternalServerError)
		return
	}

	count, err := UpdateCelebrityInDB(r.Context(), db, id, &newCel)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if count == 0 {
		http.Error(w, "Celebrity not found", http.StatusNotFound)
		return
	}

	newCel.Id = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCel)
}

// DeleteCelebrity godoc
// @Summary Удалить celebrity
// @Description Удаляет celebrity
// @Tags Celebrities
// @Produce json
// @Param id path int true "Celebrity ID"
// @Success 200 {object} Celebrity
// @Failure 404 {string} string "Celebrity not found"
// @Failure 500 {string} string "Server error"
// @Router /Celebrities/{id} [delete]
func DeleteCelebrity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	cel, err := GetCelebrityByIdFromDB(r.Context(), db, id)
	if cel == nil {
		http.Error(w, "Celebrity not found", http.StatusNotFound)
		return
	}

	_, err = DeleteCelebrityFromDB(r.Context(), db, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cel)
}

// @title Celebrities REST API
// @version 1.0
// @description Celebrities через Oracle
// @host localhost:3000
// @BasePath /
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

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server running on", 3000)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

// import (
//         "math/big"
// )

// func Height(n, m *big.Int) *big.Int {
//   zero := big.NewInt(0)
//   one := big.NewInt(1)
//   if n.Cmp(zero) == 0 || m.Cmp(zero) == 0{
//     return zero
//   }
//   if n.Cmp(one) == 0{
//     return m
//   }
//   if m.Cmp(one) == 0{
//     return one
//   }
//   if n.Cmp(m) == 1{
//     n = new(big.Int).Set(m)
//   }
//   rCount := int(n.Int64()) + 2
//   cCount := int(m.Int64()) + 2
//   f := make([][]*big.Int, rCount)
//   for i := range f {
//     f[i] = make([]*big.Int, cCount)
//   }
//   for i := 0; i <= int(n.Int64()); i++{
//     f[i][1] = one
//   }
//   for i := 0; i <= int(m.Int64()); i++{
//     f[1][i] = big.NewInt(int64(i))
//   }
//   for i := 2; i <= int(n.Int64()); i++{
//     for j := 2; j <= int(m.Int64()); j++{
//       f[i][j] = new(big.Int).Add(f[i-1][j-1], f[i][j-1])
//       f[i][j].Add(f[i][j], big.NewInt(1))
//     }
//   }
//   return f[int(n.Int64())][int(m.Int64())]
// }

// package kata

// import (
//         "math/big"
// )

// func Height(n, m *big.Int) *big.Int {
//   zero := big.NewInt(0)
//   one := big.NewInt(1)
//   if n.Cmp(zero) == 0 || m.Cmp(zero) == 0{
//     return zero
//   }
//   if n.Cmp(one) == 0{
//     return m
//   }
//   if m.Cmp(one) == 0{
//     return one
//   }
//   if n.Cmp(m) == 1{
//     n = new(big.Int).Set(m)
//   }
//   f := make([]*big.Int, int(n.Int64()) + 2)
//   for i := 0; i <= int(n.Int64()); i++{
//     f[i] = zero
//   }
//   for i := 1; i <= int(m.Int64()); i++{
//     for j := int(n.Int64()); j >= 1; j--{
//       f[j] = new(big.Int).Add(f[j], f[j-1])
//       f[j].Add(f[j], big.NewInt(1))
//     }
//   }
//   return f[int(n.Int64())]
// }
