package routes

import (
	"backend/auth"
	"backend/db"
	"backend/profile"
	"backend/questions"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bhai Home route kaam kar raha hai"))
	})
	router.HandleFunc("/register-user", auth.RegisterHandler)
	router.HandleFunc("/login-user", auth.LoginHandler)
	router.HandleFunc("/logout", auth.AuthMiddleware(auth.LogoutHandler))
	router.HandleFunc("/post-question", auth.AuthMiddleware(db.QuestionFormHandler))
	router.HandleFunc("/profile", auth.AuthMiddleware(profile.AccountHandler))
	router.HandleFunc("/edit-profile", auth.AuthMiddleware(auth.EditProfileHandler))
	router.HandleFunc("/request-password-reset", auth.AuthMiddleware(auth.PasswordResetRequest))
	router.HandleFunc("/edit-password", auth.AuthMiddleware(auth.ResetPassword))
	router.HandleFunc("/submit-question", auth.AuthMiddleware(questions.SubmitHandler))
}
