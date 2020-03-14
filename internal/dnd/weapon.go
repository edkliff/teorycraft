package dnd

type Weapon struct {
	Dice Dice
	Mod int
}

func BattleAxe() *Weapon {
	return &Weapon{DiceD8(), 0}
}

func Sekiro() *Weapon {
	return &Weapon{DiceD12(), 0}
}

func (w Weapon) Hit(targetArmor, minCrit, mastery, str int, otherMods ...int) int {
	crit := false
	attackMods := append(otherMods, w.Mod, mastery, str)
	attackNatural, attackFull := DiceD20().Roll(attackMods...)
	if attackNatural == 1 {
		return 0
	}
	if attackNatural == 20 {
		crit = true
	} else if attackFull < targetArmor {
		return 0
	}
	if attackNatural >= minCrit {
		crit = true
	}
	dmgMods := append(otherMods, w.Mod, str)
	_, dmg := w.Dice.Roll(dmgMods...)
	if crit {
		_, d := w.Dice.Roll()
		dmg += d
	}
	return dmg
}