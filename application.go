package main

import (
    "os"
    "log"
	"time"
	"strconv"
	"strings"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/secure"
    "github.com/appleboy/gin-jwt"
	"github.com/BI/api/db"
	"github.com/BI/api/models"
	"github.com/BI/api/utils"
	"github.com/BI/api/httpmodels"
	"github.com/BI/api/controllers"
)

func helloHandler(c *gin.Context) {
    claims := jwt.ExtractClaims(c)
    c.JSON(200, gin.H{
        "userID": claims["id"],
        "text":   "Hello World.",
    })
}

var identityKey = "email"

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
    secureConfig := secure.New(secure.Config{
        FrameDeny:             true,
        ContentTypeNosniff:    true,
        BrowserXssFilter:      true,
        ContentSecurityPolicy: "default-src 'self'",
        IENoOpen:              true,
        ReferrerPolicy:        "strict-origin-when-cross-origin",
        SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
    })
    router.Use(secureConfig)
	db.InitDB()
    Migrate()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Users); ok {
				return jwt.MapClaims{
					identityKey:    v.Email,
					"name":         v.Name,
					"id":           strconv.FormatUint(uint64(v.ID), 10),
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			log.Printf("claims = %+v\n", claims)
	        //logger.USER_EMAIL = claims["email"].(string)
			return &models.Users{
				Email: claims["email"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
            loginVals := &httpmodels.Login{}
            if err := c.ShouldBind(&loginVals); err != nil {
                return "", jwt.ErrMissingLoginValues
            }
            user := models.Users{}
            db.DB.Where("email = ? AND status = ?", loginVals.Email, true).First(&user)
            log.Printf("usuario = %+v\n", &user)

            //logger.InfowJson("Login", "body", loginVals.Email)
            err := utils.ComparePassword(user.Password, loginVals.Password)
            if err != nil {
                return nil, jwt.ErrFailedAuthentication
            }
            return &user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	router.POST("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"status": false, "message": "Page not found"})
	})
	auth := router.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
	    auth.POST("/Kpi/GetKpiTypeWeapon", controllers.GetKpiTypeWeapon)
	    auth.POST("/Kpi/GetKpiBrandWeapon", controllers.GetKpiBrandWeapon)
	    auth.POST("/Kpi/GetKpiQuantityWeapon", controllers.GetKpiQuantityWeapon)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	erro := router.Run(":" + port)
	log.Panic(erro)
}

func CORSMiddleware() gin.HandlerFunc {
	log.Printf("en CORSMiddleware")
	return func(c *gin.Context) {
        path := c.Request.URL.Path
        for _, param := range c.Params {
            path = strings.Replace(path, param.Value, ":"+param.Key, -1)
        }
        //logger.ENDPOINT = path
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Writer.Header().Set("Feature-Policy", "geolocation 'self'");
		if c.Request.Method == "OPTIONS" {
			log.Printf("era OPTIONS")
			c.AbortWithStatus(200)
		}
	}
}

func Migrate(){
    db.DB.AutoMigrate(&models.Users{})
    db.DB.AutoMigrate(&models.Company{})
    db.DB.AutoMigrate(&models.Brands{})
    db.DB.AutoMigrate(&models.Instruments{}).AddForeignKey("brands_id", "brands(id)", "RESTRICT", "RESTRICT")
    db.DB.AutoMigrate(&models.TypeWeapons{})
    db.DB.AutoMigrate(&models.TypePerformances{})
    db.DB.AutoMigrate(&models.BarrelWeapons{})
    db.DB.AutoMigrate(&models.BrandWeapons{})
    db.DB.AutoMigrate(&models.CartridgeBrands{})
    db.DB.AutoMigrate(&models.Reports{}).AddForeignKey("company_id", "company(id)", "RESTRICT", "RESTRICT").AddForeignKey("users_id", "users(id)", "RESTRICT", "RESTRICT")
    db.DB.AutoMigrate(&models.ReportsConclusions{}).AddForeignKey("reports_id", "reports(id)", "RESTRICT", "RESTRICT")
    db.DB.AutoMigrate(&models.Emp{}).AddForeignKey("reports_id", "reports(id)", "RESTRICT", "RESTRICT").AddForeignKey("type_weapons_id", "type_weapons(id)", "RESTRICT", "RESTRICT").AddForeignKey("brand_weapons_id", "brand_weapons(id)", "RESTRICT", "RESTRICT").AddForeignKey("barrel_weapons_id", "barrel_weapons(id)", "RESTRICT", "RESTRICT").AddForeignKey("type_performances_id", "type_performances(id)", "RESTRICT", "RESTRICT")
    db.DB.AutoMigrate(&models.EmpImage{}).AddForeignKey("emp_id", "emp(id)", "RESTRICT", "RESTRICT")
    db.DB.AutoMigrate(&models.EmpCartridge{}).AddForeignKey("emp_id", "emp(id)", "RESTRICT", "RESTRICT").AddForeignKey("cartridge_brands_id", "cartridge_brands(id)", "RESTRICT", "RESTRICT")
}
