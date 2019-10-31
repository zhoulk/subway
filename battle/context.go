package battle

type BattleContext struct {
	MilliSeconds int32
	SelfHeros    []*Hero
	OtherHeros   []*Hero
	Items        []*BattleItem
}

func NewBattleContext() *BattleContext {
	return &BattleContext{Items: []*BattleItem{}}
}

func (b *BattleContext) GetOtherHeros(group int) []*Hero {
	if group == 1 {
		return b.OtherHeros
	} else if group == 2 {
		return b.SelfHeros
	}
	return nil
}

func (b *BattleContext) GetSelfHeros(group int) []*Hero {
	if group == 1 {
		return b.SelfHeros
	} else if group == 2 {
		return b.OtherHeros
	}
	return nil
}

func (b *BattleContext) AddBattleItem(item *BattleItem) {
	b.Items = append(b.Items, item)
}
