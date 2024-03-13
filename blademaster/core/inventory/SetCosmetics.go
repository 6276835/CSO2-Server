package inventory

import (
	"log"
	"net"

	. "github.com/6276835/CSO2-Server/blademaster/typestruct"
	. "github.com/6276835/CSO2-Server/servermanager"
	. "github.com/6276835/CSO2-Server/verbose"
)

func OnFavoriteSetCosmetics(p *PacketData, client net.Conn) {
	//检索数据包
	var pkt InFavoriteSetCosmetics
	if !p.PraseFavoriteSetCosmeticsPacket(&pkt) {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent a error SetCosmetics packet !")
		return
	}
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "try to SetCosmetics but not in server !")
		return
	}
	//设置武器
	setCosmetic(pkt.Slot, pkt.ItemId, uPtr)
	DebugInfo(1, "Setting User", uPtr.UserName, "new Cosmetic", ItemList[pkt.ItemId].Name, "id", pkt.ItemId, "to slot", pkt.Slot)
	//找到对应房间玩家
	rm := GetRoomFromID(uPtr.GetUserChannelServerID(),
		uPtr.GetUserChannelID(),
		uPtr.CurrentRoomId)
	if rm == nil ||
		rm.Id <= 0 {
		return
	}
	u := rm.RoomGetUser(uPtr.Userid)
	if u == nil {
		return
	}
	setCosmetic(pkt.Slot, pkt.ItemId, u)
}
func setCosmetic(slot uint8, itemId uint32, u *User) {
	switch slot {
	case 0:
		u.Inventory.CTModel = itemId
	case 1:
		u.Inventory.TModel = itemId
	case 2:
		u.Inventory.HeadItem = itemId
	case 3:
		u.Inventory.GloveItem = itemId
	case 4:
		u.Inventory.BackItem = itemId
	case 5:
		u.Inventory.StepsItem = itemId
	case 6:
		u.Inventory.CardItem = itemId
	case 7:
		u.Inventory.SprayItem = itemId
	default:
		log.Println("Error : User", u.UserName, "try to setCosmetic invalid slot", slot)
		return
	}
}
