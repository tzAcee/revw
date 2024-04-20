package HttpServer

import "revw/HttpServer/routes"

type HandlerCollection map[string]map[string]Handler

func CreateRoutes() HandlerCollection {
	handlers := HandlerCollection{
		"/alive": {
			"GET": NewGetHandler(routes.Alive),
		},
	}

	return handlers
}
