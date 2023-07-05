package system

import "github.com/matjam/maxbbs/internal/config"

type BBS struct{}

func NewBBS() *BBS {
	return &BBS{}
}

func (b *BBS) SysName() string {
	return config.Get().SysName
}

func (b *BBS) SysopName() string {
	return config.Get().SysopName
}
