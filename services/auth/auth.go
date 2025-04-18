package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexw1983/monkeycom-go-v2/config"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/google"
)

const SessionName = "monkeycom_session"

type AuthService struct{}

func NewAuthService(store sessions.Store) *AuthService {
	gothic.Store = store
	goth.UseProviders(
		google.New(
			config.Envs.GoogleClientID,
			config.Envs.GoogleClientSecret,
			buildCallbackURL("google"),
		),
		discord.New(
			config.Envs.DiscordClientID,
			config.Envs.DiscordClientSecret,
			buildCallbackURL("discord"),
			discord.ScopeIdentify,
			discord.ScopeEmail,
		),
	)

	return &AuthService{}
}

func buildCallbackURL(provider string) string {
	return fmt.Sprintf(
		"%s:%d/auth/%s/callback",
		config.Envs.PublicHost,
		config.Envs.Port,
		provider,
	)
}

func (s *AuthService) StoreUserSession(w http.ResponseWriter, r *http.Request, user goth.User) error {
	session, err := gothic.Store.Get(r, SessionName)
	if err != nil {
		return err
	}

	session.Values["user"] = user

	if err := session.Save(r, w); err != nil {
		http.Error(w, "Unable to save session", http.StatusInternalServerError)
		return err
	}

	return nil
}

func RequireAuth(handlerFunc http.HandlerFunc, auth *AuthService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session, err := auth.GetUserSession(r)
		if err != nil {
			log.Println("User is not logged in:")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		log.Printf("User is logged in: %v", session.Email)

		handlerFunc(w, r)
	}

}

func (s *AuthService) GetUserSession(r *http.Request) (goth.User, error) {
	session, err := gothic.Store.Get(r, SessionName)
	if err != nil {
		return goth.User{}, err
	}

	user, ok := session.Values["user"].(goth.User)
	if !ok {
		return goth.User{}, fmt.Errorf("user not found in session")
	}

	return user, nil
}
