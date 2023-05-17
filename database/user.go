package database

import "log"

type User struct {
	Name string `json:"name"`
}

func (u *User) Save() (int, error) {
	stmt, err := DB.Prepare("INSERT INTO public.user (name) VALUES ($1) RETURNING id")

	res, err := stmt.Exec(u.Name)

	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return int(id), nil
}
