package simpati

import (
	"database/sql"
	"fmt"

	"github.com/aiteung/musik"
)

func GetUsernamefromPhonenumber(phone_number string, db *sql.DB) (username string) {
	username = GetUsernamefromPhonenumberInTable(phone_number, "simak_mst_mahasiswa", db)
	fmt.Println(username)
	if username == "" {
		username = GetUsernamefromPhonenumberInTable(phone_number, "simak_mst_dosen", db)

	}
	return username
}

func GetUsernamefromPhonenumberInTable(phone_number string, tabel string, db *sql.DB) (username string) {
	err := db.QueryRow("select Login from "+tabel+" where Handphone = ?", phone_number).Scan(&username)
	if err != nil {
		fmt.Printf("GetUsernamefromPhonenumberInTable %v: %v\n", tabel, err)
	}
	return username
}

func GetHashPasswordfromUsername(username string, db *sql.DB) (hashpassword string) {
	err := db.QueryRow("select user_password from simak_besan_users where user_name = ?", username).Scan(&hashpassword)
	if err != nil {
		fmt.Printf("GetHashPasswordfromUsername: %v\n", err)
	}
	return hashpassword
}

func UpdatePasswordfromUsername(username string, db *sql.DB) (newPassword string) {
	newPassword = musik.RandomString(10)
	var temp interface{}
	var err error

	err = db.QueryRow("update simak_mst_mahasiswa set Password = MD5(MD5(?)) where Login = ?", newPassword, username).Scan(&temp)
	if err != nil {
		fmt.Printf("UpdatePasswordfromUsername: %v\n", err)
	}
	err = db.QueryRow("update simak_mst_dosen set Password = MD5(MD5(?)) where Login = ?", newPassword, username).Scan(&temp)
	if err != nil {
		fmt.Printf("UpdatePasswordfromUsername: %v\n", err)
	}
	err = db.QueryRow("update simak_besan_users  set user_password = MD5(MD5(?)) where user_name = ?", newPassword, username).Scan(&temp)
	if err != nil {
		fmt.Printf("UpdatePasswordfromUsername: %v\n", err)
	}
	err = db.QueryRow("update besan_users  set user_password = MD5(MD5(?)) where user_name = ?", newPassword, username).Scan(&temp)
	if err != nil {
		fmt.Printf("UpdatePasswordfromUsername: %v\n", err)
	}
	return newPassword
}

func GetUserIdfromUsername(username string, db *sql.DB) (userid string) {
	err := db.QueryRow("select user_id from simak_besan_users where user_name = ?", username).Scan(&userid)
	if err != nil {
		fmt.Printf("GetHashPasswordfromUsername: %v\n", err)
	}
	return userid
}

func GetJadwalKuliah(kodedosen string, db *sql.DB) (jadwal []JadwalKuliah) {
	rows, _ := db.Query("select JadwalID, Nama, NamaKelas, HariID, JamMulai, JamSelesai, RuangID, Kehadiran from simak_trn_jadwal where DosenID = ? and TahunID = ?", kodedosen, getTahunID(db))
	jadwal = []JadwalKuliah{}
	for rows.Next() {
		var r JadwalKuliah
		rows.Scan(&r.JadwalID, &r.Nama, &r.NamaKelas, &r.HariID, &r.JamMulai, &r.JamSelesai, &r.RuangID, &r.Kehadiran)
		jadwal = append(jadwal, r)
	}
	return jadwal
}

func getTahunID(db *sql.DB) (TahunID TahunID) {
	db.QueryRow("SELECT TahunID FROM simak_mst_tahun where NA = 'N' group by TahunID order by TahunID DESC limit 1").Scan(&TahunID)
	return TahunID

}
