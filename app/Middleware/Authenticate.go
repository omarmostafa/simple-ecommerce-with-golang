package Middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"simple-ecommerce/app/Errors"
	"simple-ecommerce/app/controllers"
	"simple-ecommerce/app/models"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		notAuth := []string{"register", "login","get_categories","get_products"}
		requestPath := mux.CurrentRoute(req).GetName()

		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(res, req)
				return
			}
		}

		tokenHeader := req.Header.Get("Authorization")

		if tokenHeader == "" {
			controllers.RespondForbidden(res,req,Errors.NotAuthorizedToken{"Token is missing"}.Error(),nil)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			controllers.RespondForbidden(res,req,Errors.NotAuthorizedToken{"Invalid/Malformed auth token"}.Error(),nil)
			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			controllers.RespondForbidden(res,req,Errors.NotAuthorizedToken{"Malformed authentication token"}.Error(),nil)
			return
		}

		if !token.Valid {
			controllers.RespondForbidden(res,req,Errors.NotAuthorizedToken{"Invalid Token"}.Error(),nil)
			return
		}

		next.ServeHTTP(res, req)
	});
}
