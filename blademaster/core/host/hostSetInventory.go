package host

import (
	"net"

	. "github.com/6276835/CSO2-Server/blademaster/core/inventory"
	. "github.com/6276835/CSO2-Server/blademaster/typestruct"
	. "github.com/6276835/CSO2-Server/configure"
	. "github.com/6276835/CSO2-Server/kerlong"
	. "github.com/6276835/CSO2-Server/servermanager"
	. "github.com/6276835/CSO2-Server/verbose"
)

//OnHostSetUserInventory 用户发来请求将自己的装备信息发给指定user
func OnHostSetUserInventory(p *PacketData, client net.Conn) {
	//检查数据包
	var pkt InHostSetInventoryPacket
	if !p.PraseSetUserInventoryPacket(&pkt) {
		DebugInfo(2, "Error : Cannot prase a send UserInventory packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : A user request to send UserInventory but not in server!")
		return
	}
	dest := GetUserFromID(pkt.UserID)
	if dest == nil ||
		dest.Userid <= 0 {
		DebugInfo(2, "Error : A user request to send UserInventory but dest user is null!")
		return
	}
	//找到玩家的房间
	rm := GetRoomFromID(uPtr.GetUserChannelServerID(),
		uPtr.GetUserChannelID(),
		uPtr.GetUserRoomID())
	if rm == nil ||
		rm.Id <= 0 {
		DebugInfo(2, "Error : User", uPtr.UserName, "try to send UserInventory but in a null room !")
		return
	}
	//是不是房主
	if rm.HostUserID != uPtr.Userid {
		DebugInfo(2, "Error : User", uPtr.UserName, "try to send UserInventory but isn't host !")
		return
	}
	//发送用户的装备给目标user
	rst := BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeHost), BuildSetUserInventory(dest, dest.Userid))
	SendPacket(rst, uPtr.CurrentConnection)
	DebugInfo(2, "Send User", dest.UserName, "Inventory to host", uPtr.UserName)
	rst = BytesCombine(BuildHeader(uPtr.CurrentSequence, PacketTypeHost), BuildSetUserLoadout(dest))
	SendPacket(rst, uPtr.CurrentConnection)
	DebugInfo(2, "Send User", dest.UserName, "Loadout to host", uPtr.UserName)
}

//BuildSetUserInventory 建立要发给主机的玩家装备信息，按理来说应该是所有玩家的装备，待定，L-Leite是发的主机的装备加普通用户ID
func BuildSetUserInventory(u *User, destid uint32) []byte {
	if Conf.UnlockAllWeapons != 0 {
		buf := make([]byte, 10+6*len(FullInventoryItem))
		offset := 0
		WriteUint8(&buf, SetInventory, &offset)
		WriteUint32(&buf, destid, &offset)
		WriteUint8(&buf, 0, &offset)
		WriteUint16(&buf, uint16(len(FullInventoryItem)), &offset)
		for _, v := range FullInventoryItem {
			WriteUint32(&buf, v.Id, &offset)
			WriteUint16(&buf, v.Count, &offset)
		}
		return buf[:offset]
	}
	buf := make([]byte, 10+6*u.Inventory.NumOfItem)
	offset := 0
	WriteUint8(&buf, SetInventory, &offset)
	WriteUint32(&buf, destid, &offset)
	WriteUint8(&buf, 0, &offset)
	WriteUint16(&buf, u.Inventory.NumOfItem, &offset)
	for _, v := range u.Inventory.Items {
		WriteUint32(&buf, v.Id, &offset)
		WriteUint16(&buf, v.Count, &offset)
	}
	return buf[:offset]
}
