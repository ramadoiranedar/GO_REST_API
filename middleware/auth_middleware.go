package middleware

import (
	"net/http"

	"github.com/ramadoiranedar/go_restapi/helper"
	"github.com/ramadoiranedar/go_restapi/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(
	handler http.Handler,
) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "my_key_xxx" == request.Header.Get("X-API-KEY") {
		// OK
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// ERROR
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   nil,
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
