package api

import (
	"encoding/json"
	"log"
	"net/http"
	"net/mail"

	"github.com/srazap/luckbuy/internal"
	"github.com/srazap/luckbuy/models"
	"gorm.io/gorm"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	user := models.NewUser("", "", "")

	// read request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error while parsing request body", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	if user.Email == "" && user.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	// email validation
	if _, err := mail.ParseAddress(user.Email); err != nil {
		log.Println("Email Parse Error:", err.Error())
		respondWithError(w, http.StatusBadRequest, "Invalid Eamil Address")
		return
	}

	// refferral code email validation
	if user.ReferralCode != "" {
		if _, err := mail.ParseAddress(user.ReferralCode); err != nil {
			log.Println("Email Parse Error:", err.Error())
			respondWithError(w, http.StatusBadRequest, "Invalid Refferral Eamil Address")
			return
		}
	}

	// save user details in database
	if err := internal.Register(user); err != nil {
		log.Println("Internal Error:", err.Error())
		respondWithError(w, http.StatusBadRequest, "Internal Server Error, please try again")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "User registered successfully",
	})

}
func Login(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Email    string `json:"email"`
		Passowrd string `json:"password"`
	}{}

	// read request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("error while parsing request body", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	sessionId, err := internal.Login(req.Email, req.Passowrd)
	if err != nil {
		msg := "wrong email or password"
		if err != gorm.ErrRecordNotFound && err.Error() != msg {
			msg = "internal server error"
			log.Println(err)
		}
		respondWithError(w, http.StatusInternalServerError, msg)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"session_id": sessionId,
	})
}
func Logout(w http.ResponseWriter, r *http.Request) {
	sId := r.Header.Get("session_id")
	if err := internal.Logout(sId); err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "internal server error, please try again later")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Session closed successfully",
	})
}
func MyPoints(w http.ResponseWriter, r *http.Request) {
	// get user by session id
	sId := r.Header.Get("session_id")

	email, err := internal.GetEmailBySessionId(sId)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "internal server error, please try again later")
		return
	}

	summary, points, err := internal.MyPoints(email)
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "internal server error, please try again later")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"points":  points,
		"summary": summary,
	})
}
func Leaderboard(w http.ResponseWriter, r *http.Request) {
	board, err := internal.GetLeaderboard()
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "internal server error, please try again later")
		return
	}
	respondWithJSON(w, http.StatusOK, board)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
