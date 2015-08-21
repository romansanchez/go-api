package main

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "strconv"
)

func IndexConfigurationsHandler(rw http.ResponseWriter, req *http.Request) {
    queryString := req.URL.Query()
    if len(queryString["size"]) > 0 && len(queryString["from"]) > 0 {
    	size,_ := strconv.Atoi(queryString["size"][0])
    	from,_ := strconv.Atoi(queryString["from"][0])
    	if size >= 0 && from >= 0 {
    		response := PaginateConfigurations(size,from)

    		// sort
    		if len(queryString["sort"]) > 0 {
    			SortConfigs(queryString["sort"], response)
    		}
			rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
			rw.WriteHeader(http.StatusOK)
		    if err := json.NewEncoder(rw).Encode(response); err != nil {
		    	panic(err)
		    }
    	} else {
			rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
			rw.WriteHeader(http.StatusBadRequest)
			response := map[string]string{"status":http.StatusText(http.StatusBadRequest)}
		    if err := json.NewEncoder(rw).Encode(response); err != nil {
		    	panic(err)
		    }
    	}
    } else if len(queryString["size"]) > 0 {
    	size,_ := strconv.Atoi(queryString["size"][0])
    	from := 0
    	if size >= 0 {
    		response := PaginateConfigurations(size,from)

    		// sort
    		if len(queryString["sort"]) > 0 {
    			SortConfigs(queryString["sort"], response)
    		}
			rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
			rw.WriteHeader(http.StatusOK)
		    if err := json.NewEncoder(rw).Encode(response); err != nil {
		    	panic(err)
		    }
    	} else {
			rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
			rw.WriteHeader(http.StatusBadRequest)
			response := map[string]string{"status":http.StatusText(http.StatusBadRequest)}
		    if err := json.NewEncoder(rw).Encode(response); err != nil {
		    	panic(err)
		    }
    	}

    } else {
		response := ListConfigurations()
		// sort
		if len(queryString["sort"]) > 0 {
			SortConfigs(queryString["sort"], response)
		}
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.WriteHeader(http.StatusOK)
	    if err := json.NewEncoder(rw).Encode(response); err != nil {
	    	panic(err)
	    }
    }
}

func CreateConfigurationsHandler(rw http.ResponseWriter, req *http.Request) {
    config := Configuration {
    	Id: req.FormValue("id"),
    	Name: req.FormValue("name"),
    	Hostname: req.FormValue("hostname"),
    	Username: req.FormValue("username"),
    }
    if len(req.FormValue("port")) > 0 {
    	config.Port,_ = strconv.Atoi(req.FormValue("port"))
    }

	id := CreateConfiguration(config)

	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
	rw.WriteHeader(http.StatusCreated)
	response := map[string]string{"id": id, "status": http.StatusText(http.StatusCreated)}
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}

func ShowConfigurationsHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	response := ShowConfiguration(vars["id"])
	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if response != nil {
		rw.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(rw).Encode(response); err != nil {
			panic(err)
		}
	} else {
		rw.WriteHeader(http.StatusNotFound)
		response := map[string]string{"id":vars["id"], "status": http.StatusText(http.StatusNotFound)}
		if err := json.NewEncoder(rw).Encode(response); err != nil {
			panic(err)
		}
	}

}
func UpdateConfigurationsHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	config := Configuration {
		Id: req.FormValue("id"),
		Name: req.FormValue("name"),
		Hostname: req.FormValue("hostname"),
		Username: req.FormValue("username"),
	}
	if len(req.FormValue("port")) > 0 {
		config.Port,_ = strconv.Atoi(req.FormValue("port"))
	}
	UpdateConfiguration(vars["id"], config)
	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	response := map[string]string{"id": vars["id"], "status": http.StatusText(http.StatusOK)}
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}

func DestroyConfigurationsHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
    if DestroyConfiguration(vars["id"]) {
    	rw.WriteHeader(http.StatusOK)
    	response := map[string]string{"id": vars["id"], "status": http.StatusText(http.StatusOK)}
    	if err := json.NewEncoder(rw).Encode(response); err != nil {
    		panic(err)
    	}
	} else {
		rw.WriteHeader(http.StatusNotFound)
		response := map[string]string{"id": vars["id"], "status": http.StatusText(http.StatusNotFound)}
		if err := json.NewEncoder(rw).Encode(response); err != nil {
			panic(err)
		}
	}
}

func IndexUsersHandler(rw http.ResponseWriter, req *http.Request) {
	response := ListUsers()
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(rw).Encode(response); err != nil {
    	panic(err)
    }
}

func RegistrationHandler(rw http.ResponseWriter, req *http.Request) {
    id := CreateUser(req.FormValue("username"), req.FormValue("password"))
    rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
    if len(id) == 0 {
    	rw.WriteHeader(http.StatusConflict)
    	response := map[string]string{"status":http.StatusText(http.StatusConflict)}
	    if err := json.NewEncoder(rw).Encode(response); err != nil {
		  panic(err)
	    }
    } else {
	    rw.WriteHeader(http.StatusCreated)
	    response := map[string]string{"id": id, "status": http.StatusText(http.StatusCreated)}
	    if err := json.NewEncoder(rw).Encode(response); err != nil {
		  panic(err)
	    }
    }
}

func LoginHandler(rw http.ResponseWriter, req *http.Request) {
	if Authenticate(req.FormValue("username"), req.FormValue("password")) {
		token := ActivateToken(req.FormValue("username"))
		rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
		rw.WriteHeader(http.StatusOK)
		response := map[string]string{"access_token": token}
		if err := json.NewEncoder(rw).Encode(response); err != nil {
			panic(err)
		}
	} else {
		rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
		rw.WriteHeader(http.StatusUnauthorized)
		response := map[string]string{"status": http.StatusText(http.StatusUnauthorized)}
		if err := json.NewEncoder(rw).Encode(response); err != nil {
			panic(err)
		}
	}
}

func LogoutHandler(rw http.ResponseWriter, req *http.Request) {
	DeactivateToken(req.FormValue("username"))
	rw.Header().Set("Content-Type", "application/json;charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	response := map[string]string{"status": http.StatusText(http.StatusOK)}
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}
