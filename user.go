package main


type User struct {
	  Id           string `json:"id"`
    Username     string `json:"username"`
    Password     string `json:"password"`
    AccessToken  string `json:"access_token"`
}

type Users struct {
	List []User `json:"users"`
}

func UserExists(username string) bool {
  for _,user := range users.List {
    if user.Username == username {
      return true
    }
  }
  return false
}