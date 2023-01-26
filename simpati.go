package simpati

import "database/sql"

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
