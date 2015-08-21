package main


import (
  "strings"
  "time"
  "strconv"
)

var configurations = Configurations{[]Configuration{}}
var users = Users{[]User{}}

func ListConfigurations() Configurations {
  return configurations
}

func PaginateConfigurations(size int, from int) Configurations{
    subset := Configurations{[]Configuration{}}

    if from >= len(configurations.List) {
      from = 0 // default to 0 to prevent out of bounds err
    }

    if size > len(configurations.List) {
      size = len(configurations.List) - from
    }

    for i,j := from, 0; j < size; i,j = i+1,j+1{
      subset.List = append(subset.List, configurations.List[i])
    }
    return subset
}

func CreateConfiguration(c Configuration) string {
  if len(c.Id) > 0 {
    if ConfigurationExists(c.Id, configurations) {
      UpdateConfiguration(c.Id, c)
    } else {
      configurations.List = append(configurations.List, c)
    }
  } else {
    c.Id = strconv.FormatInt(time.Now().UnixNano(), 36)
    configurations.List = append(configurations.List, c)
  }
  return c.Id
}

func ShowConfiguration(id string) *Configuration {
  for _,config := range configurations.List {
    if config.Id == id {
      return &config
    }
  }
	return nil
}

func UpdateConfiguration(id string, c Configuration)  *Configuration{

	for i,config := range configurations.List {
    if config.Id == id {

      if len(strings.Trim(c.Id, " ")) > 0 {
        configurations.List[i].Id = c.Id
      }

      if len(strings.Trim(c.Name, " ")) > 0 {
        configurations.List[i].Name = c.Name
      }

      if len(strings.Trim(c.Hostname, " ")) > 0 {
        configurations.List[i].Hostname = c.Hostname
      }

      if c.Port > 0 {
        configurations.List[i].Port = c.Port
      }
      return &config
    }
  }
  return &c
}

func DestroyConfiguration(id string) bool {
  for i,config := range configurations.List {
    if config.Id == id {
      configurations.List = append(configurations.List[:i], configurations.List[i+1:]...)
      return true
    }
  }
  return false
}

func CreateUser(username string, password string) string {
  var id string
  if UserExists(username) {
    id = "" // username is taken
  } else {
    id = strconv.FormatInt(time.Now().UnixNano(), 36)
    password_hash := Hash([]byte(password),10)
    u := User {
      Id: id,
      Username: username,
      Password: string(password_hash),
    }
    users.List = append(users.List, u)
  }
  return id
}

func ListUsers() Users {
  return users
}

