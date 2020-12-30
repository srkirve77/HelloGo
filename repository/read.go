package repository

import "../ds"

func ReadAll() ([]ds.Postrequest, error) {
	read, err := con.Query("select * from TODO")
	if err != nil {
		return nil, err
	}
	all := []ds.Postrequest{}
	for read.Next() {
		data := ds.Postrequest{}
		read.Scan(&data.Name, &data.Roll_no)
		all = append(all, data)
	}
	return all, nil
}

func ReadByName(name string) ([]ds.Postrequest, error) {
	read, err := con.Query("select * from TODO where name = ?", name)
	if err != nil {
		return nil, err
	}
	all := []ds.Postrequest{}
	for read.Next() {
		data := ds.Postrequest{}
		read.Scan(&data.Name, &data.Roll_no)
		all = append(all, data)
	}
	return all, nil
}
