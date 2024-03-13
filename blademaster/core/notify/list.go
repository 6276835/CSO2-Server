package notify

import (
	"net"

	. "github.com/6276835/CSO2-Server/blademaster/typestruct"
	. "github.com/6276835/CSO2-Server/kerlong"
	. "github.com/6276835/CSO2-Server/servermanager"
	. "github.com/6276835/CSO2-Server/verbose"
)

func OnNotifyList(p *PacketData, client net.Conn) {
	//检索数据报
	var pkt InNotifyListPacket
	if !p.PraseNotifyListPacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a illegal notifylist packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "try to request notifylist but not in server !")
		return
	}
	//发送数据
	rst := BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeNotify), BuildMailList())
	SendPacket(rst, uPtr.CurrentConnection)
	DebugInfo(2, "Sent a null notify list to User", uPtr.UserName)

}

func BuildMailList() []byte {
	return []byte{0x00, 0x01}
}
