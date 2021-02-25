package entities

import "fmt"

type User struct {
	Id                 int    `json:"id"`
	Nama               string `json:"nama"`
	TanggalLahir       string `json:"tanggalLahir"`
	PendidikanTerakhir string `json:"pendidikanTerakhir"`
	Pekerjaan          string `json:"pekerjaan"`
}

func (user User) ToString() string {
	return fmt.Sprintf("id: %d\nnama: %s\nTanggalLahir: %s\nPendidikanTerakhir: %s\nPekerjaan: %s",
		user.Id, user.Nama, user.TanggalLahir, user.PendidikanTerakhir, user.Pekerjaan)
}
