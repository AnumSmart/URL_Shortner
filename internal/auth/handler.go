package auth

import (
	"fmt"
	"net/http"
	"server/configs"
	"server/pkg/jwt"
	"server/pkg/req"
	"server/pkg/resp"
)

type AuthHandlerDep struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDep) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("/auth/login", handler.Login())
	router.HandleFunc("/auth/register", handler.Register())
}

func (a *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//нужно прочитать Body
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)

		email, err := a.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(a.Config.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := LoginResponce{
			Token: token,
		}
		resp.ResponceJson(w, 201, data)
	}
}

func (a *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Register Page Here!\n")
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}

		email, err := a.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(a.Config.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := RegisterResponce{
			Token: token,
		}
		resp.ResponceJson(w, 201, data)
	}
}
