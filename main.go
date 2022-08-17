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
	db                 *gorm.DB                       = config.Connection()
	jwtService         services.JWTService            = services.NewJWTService()
	notifRepo          repositories.NotifRepo         = repositories.NewNotifRepo(db)
	authRepo           repositories.AuthRepo          = repositories.NewAuthRepo(db)
	authService        services.AuthService           = services.NewAuthService(authRepo)
	authController     controllers.AuthController     = controllers.NewAuthController(authService, jwtService)
	userRepo           repositories.UserRepo          = repositories.NewUserRepo(db)
	userService        services.UserService           = services.NewUserService(userRepo)
	userController     controllers.UserController     = controllers.NewUserController(userService, jwtService)
	categoryRepo       repositories.CategoryRepo      = repositories.NewCategoryRepo(db)
	categoryService    services.CategoryService       = services.NewCategoryService(categoryRepo)
	categoryController controllers.CategoryController = controllers.NewCategoryController(categoryService, jwtService)
	walletRepo         repositories.WalletRepo        = repositories.NewWalletRepo(db)
	walletService      services.WalletService         = services.NewWalletService(walletRepo)
	walletController   controllers.WalletController   = controllers.NewWalletController(walletService, jwtService)
	owRepo             repositories.OWRepo            = repositories.NewOWRepo(db)
	owService          services.OWService             = services.NewOWService(owRepo, notifRepo)
	owController       controllers.OWController       = controllers.NewOWController(owService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()
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
		userRoutes.PUT("/create_password", userController.CreatePassword)

	}

	catergoryRoutes := r.Group("/v1/category", middlewares.AuthorizeJWT(jwtService))
	{
		catergoryRoutes.POST("/add_category", categoryController.AddCategory)
		catergoryRoutes.GET("/get_all_category", categoryController.GetAllCategory)
	}

	walletRoutes := r.Group("/v1/wallet", middlewares.AuthorizeJWT(jwtService))
	{
		walletRoutes.POST("/create_wallet", walletController.CreateWallet)
		walletRoutes.GET("/get_all_wallet", walletController.GetAllWallet)
		walletRoutes.GET("/get_invitation_wallet", walletController.GetInvitationWallet)
	}

	owRoutes := r.Group("/v1/ow", middlewares.AuthorizeJWT(jwtService))
	{
		owRoutes.GET("/get_ow_user", owController.GetOwUser)
		owRoutes.GET("/get_for_member", owController.GetForMember)
		owRoutes.POST("/add_member", owController.AddMember)
		owRoutes.PUT("/confirm_invitation", owController.ConfirmInvitation)

	}

	r.Run()
}
