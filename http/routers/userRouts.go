package routers

import "practice-project/http/middleware"

func StartUserRoutes(){
	router.POST(BASE_PATH+"/register", handler.RegistrationController)
	router.GET(BASE_PATH+"/users",middleware.SessionAuthentication(), handler.GetUsersController)
	router.GET(BASE_PATH+"/users/:id",middleware.JwtTokenAuthentication(), handler.GetUserController)
}