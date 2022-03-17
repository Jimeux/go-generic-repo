package data

import "xorm.io/xorm"

func Init(db *xorm.Engine) error {
	_, err := db.Exec(
		`CREATE TABLE job (
    id INTEGER PRIMARY KEY AUTO_INCREMENT, 
    name varchar(200) not null
)`)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`CREATE TABLE person (
    id INTEGER PRIMARY KEY AUTO_INCREMENT, 
    name varchar(200) not null
)`)
	if err != nil {
		return err
	}
	return nil
}
