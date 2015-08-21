package main

import "net/http"

type Route struct {
	Action  string
	Method  string
	Pattern string
	HandlerFunc http.HandlerFunc
}

// Use the Secure() handler wrapper to block unauthorized access to
var routes = []Route{
    Route{
        "index",
        "GET",
        "/configurations",
        Secure(IndexConfigurationsHandler),
    },
    Route{
        "create",  
        "POST",
        "/configurations",
        Secure(CreateConfigurationsHandler),
    },
    Route{
        "show",    
        "GET",    
        "/configurations/{id:[a-zA-Z0-9]+}", 
        Secure(ShowConfigurationsHandler),
    },
    Route{
        "update",  
        "PUT",
        "/configurations/{id:[a-zA-Z0-9]+}", 
        Secure(UpdateConfigurationsHandler),
    },
    Route{
        "destroy",
        "DELETE", 
        "/configurations/{id:[a-zA-Z0-9]+}",
        Secure(DestroyConfigurationsHandler),
    },
    Route{
        "register",    
        "POST",
        "/register", 
        RegistrationHandler,
    },
    Route{
        "login",
        "POST",
        "/login",
        LoginHandler,
    },
    Route{
        "logout",
        "POST",
        "/logout",
        Secure(LogoutHandler),
    },
}