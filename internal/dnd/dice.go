package dnd

import "math/rand"

type Dice struct {
	count int
	number int
}

func (d Dice) Roll(mods ...int) (natural int, full int) {
	for i := 0; i < d.count; i ++ {
		natural += rand.Intn(d.number) + 1
	}
	full = natural
	for _, mod := range mods {
		full += mod
	}
	return
}

func DiceD4() Dice {
	return Dice{1, 4}
}

func DiceD6() Dice {
	return Dice{1, 6}
}

func Dice2D6() Dice {
	return Dice{2, 6}
}

func DiceD8() Dice {
	return Dice{1, 8}
}

func Dice2D8() Dice {
	return Dice{2, 8}
}

func DiceD10() Dice {
	return Dice{1, 10}
}

func Dice2D10() Dice {
	return Dice{2, 10}
}

func DiceD12() Dice {
	return Dice{1, 12}
}

func DiceD20() Dice {
	return Dice{1, 10}
}

