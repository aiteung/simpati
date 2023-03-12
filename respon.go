package simpati

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/aiteung/atmessage"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

func ListHandler(waclient *whatsmeow.Client, Info *types.MessageInfo, Message *waProto.Message) {
	rowID := strings.Split(*Message.ListResponseMessage.SingleSelectReply.SelectedRowId, "|")

	if rowID[0] == "respon" {
		switch rowID[1] {
		case "changegroupname":
			changegroupname(waclient, Info, rowID[2])
		case "getinformations":
			fmt.Println("mendapta info")
		}
	} else {
		fmt.Println(rowID)

	}
}

func List(Info *types.MessageInfo, Message *waProto.Message) bool {
	if Message.ListResponseMessage != nil {
		return true
	} else {
		return false
	}
}

func changegroupname(waclient *whatsmeow.Client, Info *types.MessageInfo, groupname string) {
	fmt.Println(Info.Chat)
	fmt.Println(groupname)
	if len(groupname) >= 25 {
		groupname = groupname[0:24]
	}
	if Info.Chat.Server == "g.us" {
		err := waclient.SetGroupName(Info.Chat, groupname)
		if err != nil {
			fmt.Printf("changegroupname: %v\n", err)
		} else {
			atmessage.SendMessage("Nama Group Berhasil diubah ke "+groupname+" silakan dilanjut dengan share Live Location ya kak..", Info.Chat, waclient)
		}
	} else {
		atmessage.SendMessage("Mohon maaf set Nama Group hanya bisa dilakukan di percakapan Group Whatsapp Saja. Silahkan kakak pindah ke Group WhatsApp dahulu", Info.Chat, waclient)
	}
}

func RunModuleResponse(waclient *whatsmeow.Client, Info *types.MessageInfo, Message *waProto.Message, db *sql.DB) {
	kodedosen := GetKodeDosen(Info.Sender.User, db)
	if kodedosen != "" {
		fmt.Println("kodedosen : ")
		fmt.Println(kodedosen)
		jadwalkuliah := GetJadwalKuliah(kodedosen, db)
		fmt.Println(jadwalkuliah)
		SendListMessageJadwalKuliah(waclient, Info, kodedosen, jadwalkuliah)
	}
}
