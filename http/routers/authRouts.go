package routers

import "practice-project/http/middleware"

func StartAuthRoutes() {
	router.POST(BASE_PATH+"/loginWithSession", middleware.LoginChecker(),handler.LoginSessionController)
	router.POST(BASE_PATH+"/logoutWithSession", middleware.SessionAuthentication(), handler.LogoutSessionController)
	router.POST(BASE_PATH+"/loginWithJwt",middleware.LoginChecker(), handler.LoginJwtController)
	router.POST(BASE_PATH+"/logoutWithJwt", middleware.JwtTokenAuthentication(), handler.LogoutJwtController)
}
