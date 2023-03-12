package simpati

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/musik"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

func SendListMessageJadwalKuliah(waclient *whatsmeow.Client, Info *types.MessageInfo, kodedosen string, jadwalngajar []JadwalKuliah) {
	var lmsg atmessage.ListMessage
	lmsg.Title = fmt.Sprintf("Info Mengajar %v", kodedosen)
	lmsg.Description = showInfoPertemuan(jadwalngajar)
	lmsg.FooterText = showInfoJadwal(jadwalngajar)
	lmsg.ButtonText = "Info Jadwal dan Set Nama Group"
	var listrow []atmessage.WaListRow
	for _, mkjd := range jadwalngajar {
		var jdwl atmessage.WaListRow
		jdwl.Title = mkjd.Nama
		jdwl.Description = musik.KelasNumberToAbjad(mkjd.NamaKelas) + " | " + musik.NumberStringToHari(mkjd.HariID) + " | " + mkjd.JamMulai + "-" + mkjd.JamSelesai + " | " + mkjd.RuangID
		jdwl.RowId = "respon|changegroupname|" + mkjd.JadwalID + "-" + musik.KelasNumberToAbjad(mkjd.NamaKelas) + "-" + mkjd.Nama
		listrow = append(listrow, jdwl)
	}
	var sec atmessage.WaListSection
	sec.Title = "Pilih Mata Kuliah"
	sec.Rows = listrow
	var secs []atmessage.WaListSection
	secs = append(secs, sec)
	lmsg.Sections = secs

	atmessage.SendListMessage(lmsg, Info.Chat, waclient)
}

func showInfoPertemuan(jadwalngajar []JadwalKuliah) (infongajar string) {
	infongajar = "*Jumlah Pertemuan :*\n"
	for _, mk := range jadwalngajar {
		tmpajar := mk.Kehadiran + " | " + musik.KelasNumberToAbjad(mk.NamaKelas) + " : " + mk.Nama + "\n"
		infongajar = infongajar + tmpajar

	}
	return infongajar
}

func showInfoJadwal(jadwalngajar []JadwalKuliah) (infojadwal string) {
	for _, mk := range jadwalngajar {
		tmpajar := mk.Nama + "\n" + mk.JadwalID + " | " + musik.KelasNumberToAbjad(mk.NamaKelas) + " | " + mk.RuangID + " | " + musik.NumberStringToHari(mk.HariID) + " | " + mk.JamMulai + "-" + mk.JamSelesai + "\n\n"
		infojadwal = infojadwal + tmpajar
	}
	return infojadwal
}

func GetKodeDosen(phone_number string, db *sql.DB) (KodeDosen string) {
	err := db.QueryRow("select Login from simak_mst_dosen where Handphone = ?", phone_number).Scan(&KodeDosen)
	if err != nil {
		fmt.Printf("getKodeDosen: %v\n", err)
	}
	return KodeDosen
}

func IsMultiKey(Info *types.MessageInfo, Message *waProto.Message, db *sql.DB) bool {
	m := musik.NormalizeString(Message.GetConversation())
	if (strings.Contains(m, "teung") && Info.Chat.Server == "g.us") || (Info.Chat.Server == "s.whatsapp.net") {
		complete, match := musik.IsMatch(m, "jadwal", "kuliah", "pertemuan", "jumlah", "ngajar")
		fmt.Println(complete)
		if match >= 2 && GetKodeDosen(Info.Sender.User, db) != "" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
