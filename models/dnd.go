package models

import "math"

type Character struct {
	//Id             primitive.ObjectID `bson:"_id"`
	UserId         string
	Slug           string
	Name           string
	Bio            string
	Strength       Attribute
	Dexterity      Attribute
	Constitution   Attribute
	Intelligence   Attribute
	Wisdom         Attribute
	Charisma       Attribute
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
	Value int
}

type Skill struct {
	Proficient bool
	Expertise  bool
}

func (a Attribute) GetModifier() int {
	return int(math.Round((float64(a.Value) - 10.1) / 2))
}
