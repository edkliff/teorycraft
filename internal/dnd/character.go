package dnd

type Character struct {
	 WeaponRight *Weapon
	 WeaponLeft *Weapon
	 MainAttrMod int
	 Mastery int
	 Armor int
	 Health int
}

func (defender *Character) HitBy(attacker *Character)  {
	dmg := attacker.WeaponRight.Hit(defender.Armor, 19, attacker.Mastery, attacker.MainAttrMod)
	if attacker.WeaponLeft != nil {
		dmg +=  attacker.WeaponRight.Hit(defender.Armor, 19, attacker.Mastery, 0)
	}
	defender.Health -= dmg
}

func DualWilder() Character {
	return Character{
		WeaponRight: BattleAxe(),
		WeaponLeft:  BattleAxe(),
		MainAttrMod: 3,
		Mastery:     2,
		Armor:       20,
		Health: 40,
	}
}

func Shielder() Character  {
	return Character{
		WeaponRight: BattleAxe(),
		WeaponLeft:  nil,
		MainAttrMod: 3,
		Mastery:     2,
		Armor:       21,
		Health: 40,
	}
}

func TwoHanded() Character {
	return Character{
		WeaponRight: Sekiro(),
		WeaponLeft:  nil,
		MainAttrMod: 3,
		Mastery:     2,
		Armor:       19,
		Health: 40,
	}
}