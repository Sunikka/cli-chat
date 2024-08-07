package routes

import (
	"net/http"

	"github.com/Sunikka/termitalk/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Username string    `json:"username"`
	Token    jwt.Token `json:"token"`
}

type testRes struct {
	Message string `json:"message"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, 200, testRes{Message: "Hello from the login handler!"})
}
