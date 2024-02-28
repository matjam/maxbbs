package system

import "github.com/matjam/maxbbs/internal/config"

type BBS struct{}

func NewBBS() *BBS {
	return &BBS{}
}

func (b *BBS) SystemName() string {
	return config.Get().Server.Name
}

func (b *BBS) SysopName() string {
	return config.Get().Server.Sysop
}
