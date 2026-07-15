package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	_ "github.com/sijms/go-ora/v2"
)

type Celebrity struct {
	Id           int
	FullName     string
	Nationality  string
	ReqPhotoPath string
}

var db *sql.DB

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
		from Celebrities`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cels := []Celebrity{}

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

func GetCelebrityByIdFromDB(
	ctx context.Context,
	db *sql.DB,
	id int32,
) (*Celebrity, error) {
	var cel Celebrity

	row, err := db.QueryContext(
		ctx,
		`Select id,
				fullName,
				nationality,
				reqPhotoPath
		from Celebrities
		where id = :1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	row.Next()
	err = row.Scan(
		&cel.Id,
		&cel.FullName,
		&cel.Nationality,
		&cel.ReqPhotoPath,
	)

	if err != nil {
		return nil, err
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return &cel, nil
}

func AddCelebrityToDB(
	ctx context.Context,
	db *sql.DB,
	cel *Celebrity,
) (*Celebrity, error) {
	_, err := db.ExecContext(
		ctx,
		`Insert into Celebrities(fullName, nationality, reqPhotoPath)
		values (:1, :2, :3)`,
		cel.FullName,
		cel.Nationality,
		cel.ReqPhotoPath,
	)

	if err != nil {
		return nil, err
	}

	var cel2 Celebrity

	var row *sql.Rows

	row, err = db.QueryContext(
		ctx,
		`Select id,
				fullName,
				nationality,
				reqPhotoPath
		from Celebrities
		where fullName = :1 and
		nationality = :2 and
		reqPhotoPath = :3`,
		cel.FullName,
		cel.Nationality,
		cel.ReqPhotoPath,
	)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	row.Next()
	err = row.Scan(
		&cel2.Id,
		&cel2.FullName,
		&cel2.Nationality,
		&cel2.ReqPhotoPath,
	)

	if err != nil {
		return nil, err
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return &cel2, nil
}

func UpdateCelebrityInDB(
	ctx context.Context,
	db *sql.DB,
	id int32,
	cel *Celebrity,
) (*Celebrity, error) {
	_, err := db.ExecContext(
		ctx,
		`Update Celebrities set fullName = :1,
								nationality = :2,
								reqPhotoPath = :3
		where id = :4`,
		cel.FullName,
		cel.Nationality,
		cel.ReqPhotoPath,
		id,
	)

	if err != nil {
		return nil, err
	}

	var cel2 Celebrity

	var row *sql.Rows

	row, err = db.QueryContext(
		ctx,
		`Select id,
				fullName,
				nationality,
				reqPhotoPath
		from Celebrities
		where id = :1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	row.Next()
	err = row.Scan(
		&cel2.Id,
		&cel2.FullName,
		&cel2.Nationality,
		&cel2.ReqPhotoPath,
	)

	if err != nil {
		return nil, err
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return &cel2, nil
}

func DeleteCelebrityFromDB(
	ctx context.Context,
	db *sql.DB,
	id int32,
) (*Celebrity, error) {
	cel, err := GetCelebrityByIdFromDB(ctx, db, id)

	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(
		ctx,
		`Delete from Celebrities
		where id = :1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	return cel, nil
}

var schemaStr = `
	schema {
		query: Query
		mutation: Mutation
	}

	type Celebrity {
		id: Int!
		fullName: String!
		nationality: String!
		reqPhotoPath: String!
	}

	type Query {
		celebrities: [Celebrity!]!
		celebrity(id: Int!): Celebrity
	}

	type Mutation {
		addCelebrity(
			fullName: String!
			nationality: String!
			reqPhotoPath: String!
		): Celebrity!

		updateCelebrity(
			id: Int!
			fullName: String!
			nationality: String!
			reqPhotoPath: String!
		): Celebrity!

		deleteCelebrity(id: Int!): Celebrity!
	}
`

type Resolver struct{}

type CelebrityResolver struct {
	c Celebrity
}

func (r *CelebrityResolver) ID() int32 {
	return int32(r.c.Id)
}

func (r *CelebrityResolver) FullName() string {
	return r.c.FullName
}

func (r *CelebrityResolver) Nationality() string {
	return r.c.Nationality
}

func (r *CelebrityResolver) ReqPhotoPath() string {
	return r.c.ReqPhotoPath
}

func (r *Resolver) Celebrities(ctx context.Context) ([]*CelebrityResolver, error) {
	list, err := GetAllCelebritiesFromDB(ctx, db)
	if err != nil {
		return nil, err
	}

	res := make([]*CelebrityResolver, 0, len(list))

	for _, cel := range list {
		res = append(res, &CelebrityResolver{c: cel})
	}

	return res, nil
}

type CelebrityArgs struct {
	ID int32
}

func (r *Resolver) Celebrity(
	ctx context.Context,
	args CelebrityArgs,
) (*CelebrityResolver, error) {
	cel, err := GetCelebrityByIdFromDB(ctx, db, args.ID)
	if err != nil {
		return nil, err
	}

	return &CelebrityResolver{c: *cel}, nil
}

type AddCelebrityArgs struct {
	FullName     string
	Nationality  string
	ReqPhotoPath string
}

func (r *Resolver) AddCelebrity(
	ctx context.Context,
	args AddCelebrityArgs,
) (*CelebrityResolver, error) {
	cel := &Celebrity{
		FullName:     args.FullName,
		Nationality:  args.Nationality,
		ReqPhotoPath: args.ReqPhotoPath,
	}

	added, err := AddCelebrityToDB(ctx, db, cel)
	if err != nil {
		return nil, err
	}

	return &CelebrityResolver{c: *added}, nil
}

type UpdateCelebrityArgs struct {
	ID           int32
	FullName     string
	Nationality  string
	ReqPhotoPath string
}

func (r *Resolver) UpdateCelebrity(
	ctx context.Context,
	args UpdateCelebrityArgs,
) (*CelebrityResolver, error) {
	cel := &Celebrity{
		FullName:     args.FullName,
		Nationality:  args.Nationality,
		ReqPhotoPath: args.ReqPhotoPath,
	}

	updated, err := UpdateCelebrityInDB(ctx, db, args.ID, cel)
	if err != nil {
		return nil, err
	}

	return &CelebrityResolver{c: *updated}, nil
}

type DeleteCelebrityArgs struct {
	ID int32
}

func (r *Resolver) DeleteCelebrity(
	ctx context.Context,
	args DeleteCelebrityArgs,
) (*CelebrityResolver, error) {
	deleted, err := DeleteCelebrityFromDB(ctx, db, args.ID)
	if err != nil {
		return nil, err
	}

	return &CelebrityResolver{c: *deleted}, nil
}

func main() {
	const connString string = "oracle://cel_admin:222@localhost:1521/CEL_PDB"

	var err error

	db, err = sql.Open("oracle", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema := graphql.MustParseSchema(schemaStr, &Resolver{})

	http.Handle("/graphql", &relay.Handler{
		Schema: schema,
	})

	log.Println("Server running on http://localhost:3000/graphql")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
