package models

import "subway/tool"

const (
	HeroTypeStrength    int8 = 1
	HeroTypeAgility     int8 = 2
	HeroTypeIntelligent int8 = 3
)

type HeroInfo struct {
	Uid    string
	HeroId int32
	Name   string
}

func CreateAHero() *HeroInfo {
	return &HeroInfo{
		Uid:    tool.UniqueId(),
		HeroId: 1002,
		Name:   "",
	}
}
