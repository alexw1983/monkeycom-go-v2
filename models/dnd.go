package models

import (
	"math"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HitPoints struct {
	Current   int
	Temporary int
	Maximum   int
}

type Character struct {
	Id     primitive.ObjectID `bson:"_id"`
	UserId string
	Slug   string
	Name   string
	Bio    string

	Level      int
	Background string
	Species    string
	Class      string
	SubClass   string
	HitPoints  HitPoints
	XP         int

	Strength     Attribute
	Dexterity    Attribute
	Constitution Attribute
	Intelligence Attribute
	Wisdom       Attribute
	Charisma     Attribute

	Athletics      Skill
	Acrobatics     Skill
	SleightOfHand  Skill
	Stealth        Skill
	Arcana         Skill
	History        Skill
	Investigation  Skill
	Nature         Skill
	Religion       Skill
	AnimalHandling Skill
	Insight        Skill
	Medicine       Skill
	Perception     Skill
	Survival       Skill
	Deception      Skill
	Intimidation   Skill
	Performance    Skill
	Persuasion     Skill
}

type Attribute struct {
	Value       int
	SavingThrow bool
}

type Skill struct {
	Proficient bool
	Expertise  bool
}

func (a Attribute) GetModifier() int {
	return int(math.Round((float64(a.Value) - 10.1) / 2))
}

func (c Character) ArmourClass() int {
	return 10 + c.Dexterity.GetModifier()
}

func (c Character) GetProficiencyBonus() int {
	switch {
	case c.Level >= 29:
		return 9
	case c.Level >= 25:
		return 8
	case c.Level >= 21:
		return 7
	case c.Level >= 17:
		return 6
	case c.Level >= 13:
		return 5
	case c.Level >= 9:
		return 4
	case c.Level >= 5:
		return 3
	default:
		return 2
	}
}
