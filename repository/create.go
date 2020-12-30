package repository

import "fmt"

func CreateTODO(name, roll_no string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?,?)", name, roll_no)

	defer insertQ.Close()
	if err != nil {
		fmt.Println("this is the error : ", err)
		return err
	}
	return nil
}

func DeleteTODO(name string) error {
	insertQ, err := con.Query("DELETE FROM TODO where name=?", name)
	defer insertQ.Close()
	if err != nil {
		fmt.Println("this is the error : ", err)
		return err
	}
	return nil
}
