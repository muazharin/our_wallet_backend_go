package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/muazharin/our_wallet_backend_go/config"
	"github.com/muazharin/our_wallet_backend_go/src/controllers"
	"github.com/muazharin/our_wallet_backend_go/src/middlewares"
	"github.com/muazharin/our_wallet_backend_go/src/repositories"
	"github.com/muazharin/our_wallet_backend_go/src/services"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.Connection()
	jwtService     services.JWTService        = services.NewJWTService()
	authRepo       repositories.AuthRepo      = repositories.NewAuthRepo(db)
	authService    services.AuthService       = services.NewAuthService(authRepo)
	authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
	userRepo       repositories.UserRepo      = repositories.NewUserRepo(db)
	userService    services.UserService       = services.NewUserService(userRepo)
	userController controllers.UserController = controllers.NewUserController(userService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSMiddleware(), middlewares.APIKey())

	authRoutes := r.Group("/v1/auth")
	{
		authRoutes.POST("/signup", authController.SignUp)
		authRoutes.POST("/signin", authController.SignIn)
		authRoutes.POST("/check_phone_number", authController.CheckPhone)
	}

	userRoutes := r.Group("/v1/user", middlewares.AuthorizeJWT(jwtService))
	{
		userRoutes.POST("/create_password", userController.CreatePassword)

	}

	r.Run()
}
