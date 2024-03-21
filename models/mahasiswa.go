package models

type Mahasiswa struct {
	MahasiswaId   int64  `gorm:"primaryKey" json:"mahasiswa_id"`
	MahasiswaNama string `gorm:"type:varchar(300)" json:"mahasiswa_nama"`
	MahasiswaNim  int    `gorm:"type:text" json:"mahasiswa_nim"`
}
