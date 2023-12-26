package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var mySigningKey = []byte("jg4893gj938ght938h")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Anu"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		err = fmt.Errorf("something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func GenerateEndpoint(w http.ResponseWriter, r *http.Request) {
	tokenString, err := GenerateJWT()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error generating token")
		return
	}
	fmt.Fprintln(w, tokenString)
}

func ValidateEndpoint(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("Authorization")
	token, err := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid token")
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Fprintf(w, "Token is valid. Claims: %s", claims)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Invalid token")
	}
}

func main() {
	router := mux.NewRouter()
	router.Use(ValidateEndpoint())
	router.HandleFunc("/generate", GenerateEndpoint).Methods("GET")
	router.HandleFunc("/validate", ValidateEndpoint).Methods("GET")
	router.HandleFunc("/login", func login(w http.ResponseWriter, r *http.Request){}).Methods("GET")
	http.ListenAndServe(":8000", router)
}
