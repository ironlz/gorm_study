package main

import (
	"gorm_study/db"
)

func main() {
	_, err := db.GetDB()
	if err != nil {
		panic(err)
	}

}
