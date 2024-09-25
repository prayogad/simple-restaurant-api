package middleware

import (
	"context"
	"net/http"
	"simple-restaurant-web/helper"
	"simple-restaurant-web/model/web"
	"simple-restaurant-web/service"
	"strings"
)

type AuthMiddleware struct {
	Handler         http.Handler
	CustomerService service.CustomerService
}

func NewAuthMiddleware(handler http.Handler, customerService service.CustomerService) *AuthMiddleware {
	return &AuthMiddleware{
		Handler:         handler,
		CustomerService: customerService,
	}
}

// User Must Login First
func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	header := request.Header.Get("X-API-KEY")
	roles := request.Header.Get("Authorization")
	token, err := request.Cookie("auth")
	noAuthNeeded := []string{"/customer/register", "/customer/login"}
	// adminOnly := []string{"/food"}

	if header != "RAHASIA" {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Success: false,
			Message: "No header",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return
	}

	if err != nil {
		for _, path := range noAuthNeeded {
			if request.URL.Path == path {
				middleware.Handler.ServeHTTP(writer, request)
				return
			}
		}

		if roles == "admin" {
			if strings.HasPrefix(request.URL.Path, "/food") {
				middleware.Handler.ServeHTTP(writer, request)
				return
			}
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Success: false,
			Message: "Unauthorize",
		}

		helper.WriteToResponseBody(writer, webResponse)
	} else {
		// ok
		id, username := middleware.CustomerService.ValidateToken(token.Value)
		if username == "" && id == 0 {
			for _, path := range noAuthNeeded {
				if request.URL.Path == path {
					middleware.Handler.ServeHTTP(writer, request)
					return
				}
			}

			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Success: false,
				Message: "Unauthorize",
			}

			helper.WriteToResponseBody(writer, webResponse)
			return
		}
		ctx := context.WithValue(request.Context(), "idCustomer", id)
		ctx = context.WithValue(ctx, "usernameCustomer", username)
		middleware.Handler.ServeHTTP(writer, request.WithContext(ctx))
	}
}
