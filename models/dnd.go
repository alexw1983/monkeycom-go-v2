package models

import (
	"errors"
	"log"
	"math"

	"slices"

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

	Proficiencies []Skill
	Expertise     []Skill
}

type Attribute struct {
	Value       int
	SavingThrow bool
}

type SkillProficiency struct {
	Skill      Skill
	Proficient bool
	Expertise  bool
}

type SkillRoll struct {
	DiceRoll         int
	Ability          string
	AbilityModifier  int
	Proficient       bool
	ProficiencyBonus int
	Expertise        bool
}

func (roll SkillRoll) Total() int {
	if roll.Expertise {
		return roll.DiceRoll + roll.AbilityModifier + (2 * roll.ProficiencyBonus)
	}
	if roll.Proficient {
		return roll.DiceRoll + roll.AbilityModifier + roll.ProficiencyBonus
	}

	return roll.DiceRoll + roll.AbilityModifier
}

func (roll SkillRoll) IsCriticalSuccess() bool {
	return roll.DiceRoll == 20
}

func (roll SkillRoll) IsBotch() bool {
	return roll.DiceRoll == 1
}

func (c Character) IsProficient(skill Skill) bool {
	for _, s := range c.Proficiencies {
		if s == skill {
			return true
		}
	}
	return false
}

func (c Character) GetAttributeBonus(skill Skill) int {
	var total = c.GetSkillModifier(skill)
	if c.IsProficient(skill) {
		total += c.GetProficiencyBonus()
	}

	if c.HasExpertise(skill) {
		total += c.GetProficiencyBonus()
	}

	return total
}

func (c Character) GetSkillModifier(skill Skill) int {
	switch skill {
	case Athletics:
		return c.Strength.GetModifier()
	case Acrobatics, SleightOfHand, Stealth:
		return c.Dexterity.GetModifier()
	case Arcana, History, Investigation, Nature, Religion:
		return c.Intelligence.GetModifier()
	case AnimalHandling, Insight, Medicine, Perception, Survival:
		return c.Wisdom.GetModifier()
	case Deception, Intimidation, Performance, Persuasion:
		return c.Charisma.GetModifier()
	default:
		log.Printf("Unknown skill: %s", skill)
		return 0
	}
}

func (c Character) HasExpertise(skill Skill) bool {
	return slices.Contains(c.Expertise, skill)
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

type Skill int

const (
	Athletics Skill = iota
	Acrobatics
	SleightOfHand
	Stealth
	Arcana
	History
	Investigation
	Nature
	Religion
	AnimalHandling
	Insight
	Medicine
	Perception
	Survival
	Deception
	Intimidation
	Performance
	Persuasion
)

func (s Skill) String() string {
	return [...]string{
		"Athletics",
		"Acrobatics",
		"Sleight Of Hand",
		"Stealth",
		"Arcana",
		"History",
		"Investigation",
		"Nature",
		"Religion",
		"Animal Handling",
		"Insight",
		"Medicine",
		"Perception",
		"Survival",
		"Deception",
		"Intimidation",
		"Performance",
		"Persuasion"}[s]
}

func ParseSkill(s string) (Skill, error) {
	switch s {
	case "Athletics":
		return Athletics, nil
	case "Acrobatics":
		return Acrobatics, nil
	case "SleightOfHand":
		return SleightOfHand, nil
	case "Stealth":
		return Stealth, nil
	case "Arcana":
		return Arcana, nil
	case "History":
		return History, nil
	case "Investigation":
		return Investigation, nil
	case "Nature":
		return Nature, nil
	case "Religion":
		return Religion, nil
	case "AnimalHandling":
		return AnimalHandling, nil
	case "Insight":
		return Insight, nil
	case "Medicine":
		return Medicine, nil
	case "Perception":
		return Perception, nil
	case "Survival":
		return Survival, nil
	case "Deception":
		return Deception, nil
	case "Intimidation":
		return Intimidation, nil
	case "Performance":
		return Performance, nil
	case "Persuasion":
		return Persuasion, nil
	}

	return -1, errors.New("Unknown skill: " + s)
}

func (s Skill) EnumIndex() int {
	return int(s)
}
