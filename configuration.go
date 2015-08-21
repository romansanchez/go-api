package main


type Configuration struct {
	Id        string `json:"id"`
    Name      string `json:"name"`
    Hostname  string `json:"hostname"`
    Port      int    `json:"port"`
    Username  string `json:"username"`
}

type Configurations struct {
	List []Configuration `json:"configurations"`
}

func ConfigurationExists(id string, c Configurations) bool {
  for _,el := range c.List {
    if el.Id == id {
      return true
    }
  }
  return false
}