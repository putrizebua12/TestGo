package model

import (
	"database/sql"
	"entities"
)

type UserModel struct {
	Db *sql.DB
}

func (userModel UserModel) FindAll() (user []entities.User, err error) {
	rows, err := userModel.Db.Query("select * from user")
	if err != nil {
		return nil, err
	} else {
		var users []entities.User
		for rows.Next() {
			var id int
			var nama string
			var tanggalLahir string
			var pendidikanTerakhir string
			var pekerjaan string

			err2 := rows.Scan(&id, &nama, &tanggalLahir, &pendidikanTerakhir, &pekerjaan)
			if err2 != nil {
				return nil, err2
			} else {
				user := entities.User{
					Id:                 id,
					Nama:               nama,
					TanggalLahir:       tanggalLahir,
					PendidikanTerakhir: pendidikanTerakhir,
					Pekerjaan:          pekerjaan,
				}
				users = append(users, user)
			}
		}
		return users, nil
	}
}


func (userModel UserModel) Search(keyword string) (user []entities.User, err error) {
	rows, err := userModel.Db.Query("select * from user where nama like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		var users []entities.User
		for rows.Next() {
			var id int
			var nama string
			var tanggalLahir string
			var pendidikanTerakhir string
			var pekerjaan string

			err2 := rows.Scan(&id, &nama, &tanggalLahir, &pendidikanTerakhir, &pekerjaan)
			if err2 != nil {
				return nil, err2
			} else {
				user := entities.User{
					Id:                 id,
					Nama:               nama,
					TanggalLahir:       tanggalLahir,
					PendidikanTerakhir: pendidikanTerakhir,
					Pekerjaan:          pekerjaan,
				}
				users = append(users, user)
			}
		}
		return users, nil
	}
}

func (userModel UserModel) Create(user *entities.User) (err error) {
	result, err := userModel.Db.Exec("insert into user(nama, tanggalLahir, pendidikanTerakhir, pekerjaan)
	values(?,?,?,?)", user.Nama, user.TanggalLahir, user.PendidikanTerakhir, user.Pekerjaan)
	if err != nil {
		return nil, err
	} else {
		user.Id, _ = result.LastInsertId()
		return nil
		
	}
}

func (userModel UserModel) Update(user *entities.User) (int, error) {
	result, err := userModel.Db.Exec("update user set nama = ?, tanggalLahir = ?, pendidikanTerakhir = ?, pekerjaan = ? where id = ?",
	 user.Nama, user.TanggalLahir, user.PendidikanTerakhir, user.Pekerjaan)
	if err != nil {
		return 0, err
	} else {
		
		return result.RowAffected()
		
	}
}

func (userModel UserModel) Delete(id int) (int, error) {
	result, err := userModel.Db.Exec("delete from user where id = ?", id)
	if err != nil {
		return 0, err
	} else {
		
		return result.RowAffected()
		
	}
}