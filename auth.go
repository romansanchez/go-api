package main

import (
  "net/http"
  "fmt"
  "strings"
  "encoding/json"
  "golang.org/x/crypto/bcrypt"
  "crypto/rand"
  "encoding/hex"
)

// lookup access token
func ValidToken(token string) bool {
	for _, user := range users.List {
		if user.AccessToken == token {
			return true
		}
	}
	return false
}

// wrapper to check if token is passed in requests
func Secure(fn http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		auth_value := req.Header.Get("Authorization")
		auth_value = strings.TrimPrefix(auth_value, "Bearer ")
		if len(auth_value) == 0 || !ValidToken(auth_value) {
			rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
			rw.WriteHeader(http.StatusUnauthorized)
			response := map[string]string{"status": http.StatusText(http.StatusUnauthorized)}
			json.NewEncoder(rw).Encode(response)
			return
		}
		fn(rw, req)
		fmt.Println("request successful")
	}
}

// handles login verification of username and password
func Authenticate(username string, password string) bool {
	for _,user := range users.List {
		pw := []byte(password)
		if user.Username == username && (bcrypt.CompareHashAndPassword([]byte(user.Password), pw) == nil) {
			return true
		}
	}
	return false
}

// if login is successful, token is set for user and returned
func ActivateToken(username string) string {
	for i,user := range users.List {
		if user.Username == username {
			users.List[i].AccessToken = Token()
			return users.List[i].AccessToken
		}
	}
	return ""
}

// clear token on logout
func DeactivateToken(username string) {
	for i,user := range users.List {
		if user.Username == username {
			users.List[i].AccessToken = ""
			return
		}
	}
}

// hash user's password before storing
func Hash(password []byte, cost int) []byte {
  hash, err := bcrypt.GenerateFromPassword(password, cost)
  if err != nil {
    panic(err)
  }
  return hash
}

// random hex string for access token
func Token() string {
  bytes := make([]byte, 10)
  _, err := rand.Read(bytes)
  if err != nil {
    panic(err)
  }
  return hex.EncodeToString(bytes)
}