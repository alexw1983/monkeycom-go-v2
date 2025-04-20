package handlers

import (
	"net/http"

	"github.com/alexw1983/monkeycom-go-v2/models"
	"github.com/alexw1983/monkeycom-go-v2/services/dice"
	"github.com/alexw1983/monkeycom-go-v2/views/dnd"
	"github.com/alexw1983/monkeycom-go-v2/views/dnd/fragments"
	"github.com/gorilla/mux"
)

func (h *Handler) Character(w http.ResponseWriter, r *http.Request) {
	u, _ := h.auth.GetUserSession(r)

	vars := mux.Vars(r)
	slug := vars["slug"]
	character := h.db.GetCharacter(u.Email, slug)
	dnd.Character(u, character).Render(r.Context(), w)
}

func (h *Handler) Characters(w http.ResponseWriter, r *http.Request) {
	u, _ := h.auth.GetUserSession(r)
	characters := h.db.GetCharacters(u.Email)

	dnd.Characters(u, characters).Render(r.Context(), w)
}

func (h *Handler) RollSkill(w http.ResponseWriter, r *http.Request) {
	u, _ := h.auth.GetUserSession(r)

	vars := mux.Vars(r)
	slug := vars["slug"]
	s := vars["skill"]

	skill, err := models.ParseSkill(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	character := h.db.GetCharacter(u.Email, slug)

	var diceRoll = dice.RollDX(20)
	var modifider = character.GetAttributeBonus(skill)

	var result = models.SkillRoll{
		DiceRoll:         diceRoll,
		Ability:          skill.String(),
		AbilityModifier:  modifider,
		Proficient:       character.IsProficient(skill),
		ProficiencyBonus: character.GetProficiencyBonus(),
		Expertise:        character.HasExpertise(skill),
	}

	fragments.Roll(result).Render(r.Context(), w)
}
