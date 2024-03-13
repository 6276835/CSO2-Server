package host

import (
	"net"

	. "github.com/6276835/CSO2-Server/blademaster/typestruct"
	. "github.com/6276835/CSO2-Server/verbose"
)

const (
	WeaponsPoint = 3
	UserKillOne  = 7
	UserDeath    = 8
	UserAssist   = 9
	GameScore    = 10 //比分类
	TeamChanging = 11
	HostUnk00    = 12 //未知
	BuyItem      = 14
	UserRevived  = 20
	OnGameEnd    = 21

	SetInventory = 101 // 不一定是101，可能其他的数据包
	ItemUsing    = 105
	SetLoadout   = 107
	SetBuyMenu   = 111

	//kill类型
	KillSelf   = 0xFF //自杀
	KillOne    = 1
	KillTeamCt = 2
	KillTeamTr = 1
)

func OnHost(p *PacketData, client net.Conn) {
	var pkt InHostPacket
	if p.PraseHostPacket(&pkt) {
		switch pkt.InHostType {
		case GameStart:
			OnHostGameStart(client)
		case OnGameEnd:
			OnHostGameEnd(p, client)
		case SetInventory:
			OnHostSetUserInventory(p, client)
		case SetLoadout:
			OnHostSetUserLoadout(p, client)
		case SetBuyMenu:
			OnHostSetUserBuyMenu(p, client)
		case TeamChanging:
			OnChangingTeam(p, client)
		case ItemUsing:
			OnHostItemUsing(p, client)
		case WeaponsPoint:
			OnHostWeaponPoint(p, client)
		case UserKillOne:
			OnHostKillPacket(p, client)
		case UserDeath:
			OnHostDeathPacket(p, client)
		case UserAssist:
			OnHostAssistPacket(p, client)
		case BuyItem:
			OnHostBuyItem(p, client)
		case UserRevived:
			OnHostRevivedPacket(p, client)
		case GameScore:
			OnHostGameScorePacket(p, client)
		//case HostUnk00:
		//fmt.Println("TeamWinPacket", p.data[:p.datalen], "from", client.RemoteAddr().String())
		default:
			DebugInfo(2, "Unknown host packet", pkt.InHostType, "from", client.RemoteAddr().String(), p.Data)
		}
	} else {
		DebugInfo(2, "Error : Recived a illegal host packet from", client.RemoteAddr().String())
	}
}
