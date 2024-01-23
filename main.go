// package main

// import "fmt"

// type AirtimeAdapter struct {
// 	Name  string
// 	Email string
// }

// type PrepaidPackageRequest struct {
// 	PhoneNumber string
// 	Amount      float64
// }

// type PrepaidPackageResponse struct {
// 	ConfirmationCode string
// 	Status           string
// }




// package main

// import (
// 	"errors"
// 	"fmt"
// 	"log"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// // Configuration struct for holding application configurations
// type Configurations struct {
// 	Auth struct {
// 		MerchantSecret string
// 	}
// }

// // CustomMerchantJWTClaims struct for custom JWT claims
// type CustomMerchantJWTClaims struct {
// 	jwt.StandardClaims
// }

// // Routes that don't require authentication
// var routesThatDontNeedAuth = []string{
// 	"/api/v1/public",
// }

// // Middleware to build merchant authentication
// func BuildMerchantAuthMiddleware(config Configurations) echo.MiddlewareFunc {
// 	signingKey := []byte(config.Auth.MerchantSecret)

// 	conf := middleware.JWTConfig{
// 		Claims: CustomMerchantJWTClaims{},
// 		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
// 			keyFunc := func(t *jwt.Token) (interface{}, error) {
// 				if t.Method.Alg() != "HS256" {
// 					return nil, fmt.Errorf("unexpected JWT signing method=%v", t.Header["alg"])
// 				}
// 				return signingKey, nil
// 			}

// 			token, err := jwt.ParseWithClaims(auth, &CustomMerchantJWTClaims{}, keyFunc)
// 			if err != nil {
// 				return nil, err
// 			}
// 			if !token.Valid {
// 				return nil, errors.New("invalid token")
// 			}
// 			return token, nil
// 		},
// 		Skipper: func(c echo.Context) bool {
// 			currentRoute := c.Request().URL.Path

// 			for _, route := range routesThatDontNeedAuth {
// 				if currentRoute == route {
// 					return true
// 				}
// 			}

// 			// Additional skipper logic can be added here for specific routes
// 			// For example, admins routes can use a different middleware with a different secret
// 			// if strings.HasPrefix(currentRoute, "/api/v1/admins") {
// 			// 	return true
// 			// }

// 			return false
// 		},
// 	}

// 	return middleware.JWTWithConfig(conf)
// }

// // Route handler for protected route
// func ProtectedRouteHandler(c echo.Context) error {
// 	user := c.Get("user")
// 	token := user.(*jwt.Token)
// 	claims := token.Claims.(*CustomMerchantJWTClaims)

// 	return c.String(200, fmt.Sprintf("Protected route accessed by user: %s", claims.Subject))
// }

// func main() {
// 	// Create an instance of the Echo router
// 	e := echo.New()

// 	// Configuration
// 	config := Configurations{
// 		Auth: struct {
// 			MerchantSecret string
// 		}{
// 			MerchantSecret: "your_merchant_secret_key",
// 		},
// 	}

// 	// Define your routes
// 	merchantRoutes := e.Group("/api/v1/merchant")
// 	merchantRoutes.Use(BuildMerchantAuthMiddleware(config))
// 	merchantRoutes.GET("/protected", ProtectedRouteHandler)

// 	// Start the server
// 	log.Fatal(e.Start(":8080"))
// }

package main

import "fmt"

func Check() {
	Func1 := func() {
		fmt.Println("call one")

		// Define and initialize Func2 inline
		Func2 := func () {
			fmt.Println("func two")
		}
		Func2() // Call Func2 to print "func two"
	}

	Func1() // Call Func1 to print "call one" and "func two"
}

func main() {
	Check()
}

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// 	echojwt "github.com/labstack/echo-jwt/v4"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// // jwtCustomClaims are custom claims extending default ones.
// // See https://github.com/golang-jwt/jwt for more examples
// type jwtCustomClaims struct {
// 	Name  string `json:"name"`
// 	Admin bool   `json:"admin"`
// 	jwt.RegisteredClaims
// }

// type LoginUser struct{
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// func accessible(c echo.Context) error {
// 	return c.String(http.StatusOK, "Accessible")
// }

// func login(c echo.Context) error {
// 	user:=new(LoginUser)

// 	if err:=c.Bind(user); err!=nil{
// 		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
// 	}

// 	if user.Username != "jon" || user.Password != "shhh!" {
// 		return echo.ErrUnauthorized
// 	}

// 	claims := &jwtCustomClaims{
// 		"Jon Snow",
// 		true,
// 		jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	fmt.Println("unsigned token: ",token)

// 	t, err := token.SignedString([]byte("secret"))
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("signed token: ",t)

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"token": t,
// 	})
// }

// func restricted(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	fmt.Println("restricted/n")
// 	claims := user.Claims.(*jwtCustomClaims)
// 	name := claims.Name
// 	return c.String(http.StatusOK, "Welcome "+name+"!")
// }

// func jwtMiddleware() echo.MiddlewareFunc {
// 	config := echojwt.Config{
// 		NewClaimsFunc: func(c echo.Context) jwt.Claims {
// 			return new(jwtCustomClaims)
// 		},
// 		SigningKey: []byte("secret"),
// 	}

// 	return echojwt.WithConfig(config)
// }

// func main() {
// 	e := echo.New()
// 	e.Use(middleware.Recover())

// 	e.POST("/login", login)
// 	e.GET("/", accessible)

// 	r := e.Group("/restricted")
// 	r.Use(jwtMiddleware())
// 	r.GET("/", restricted)

// 	e.Logger.Fatal(e.Start(":1323"))
// }

// package main

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"net/http"

// 	echojwt "github.com/labstack/echo-jwt/v4"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// 	"github.com/lestrrat-go/jwx/jwk"
// )

// func getKey(token *jwt.Token) (interface{}, error) {

// 	keySet, err := jwk.Fetch(context.Background(), "https://www.googleapis.com/oauth2/v3/certs")
// 	if err != nil {
// 		return nil, err
// 	}

// 	keyID, ok := token.Header["kid"].(string)
// 	if !ok {
// 		return nil, errors.New("expecting JWT header to have a key ID in the kid field")
// 	}

// 	key, found := keySet.LookupKeyID(keyID)

// 	if !found {
// 		return nil, fmt.Errorf("unable to find key %q", keyID)
// 	}

// 	var pubkey interface{}
// 	if err := key.Raw(&pubkey); err != nil {
// 		return nil, fmt.Errorf("unable to get the public key. Error: %s", err.Error())
// 	}

// 	return pubkey, nil
// }

// func accessible(c echo.Context) error {
// 	return c.String(http.StatusOK, "Accessible")
// }

// func restricted(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)
// 	claims := user.Claims.(jwt.MapClaims)
// 	name := claims["name"].(string)
// 	return c.String(http.StatusOK, "Welcome "+name+"!")
// }

// func main() {
// 	e := echo.New()

// 	// Middleware
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	// Unauthenticated route
// 	e.GET("/", accessible)

// 	// Restricted group
// 	r := e.Group("/restricted")
// 	{
// 		config := echojwt.Config{
// 			KeyFunc: getKey,
// 		}
// 		r.Use(echojwt.WithConfig(config))
// 		r.GET("", restricted)
// 	}

//		e.Logger.Fatal(e.Start(":1323"))
//	}

// package main

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// func customSkipper(c echo.Context) bool {
// 	return c.Request().URL.Path == "/skip"
// }

// func main() {
// 	e := echo.New()

// 	// Middleware
// 	// Use the custom skipper with the Logger middleware
// 	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
// 		Skipper: customSkipper,
// 	}))

// 	// Routes
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})

// 	e.GET("/skip", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "This request is skipped!")
// 	})

// 	// Start the server
// 	e.Start(":1323")
// }

// package main

// import (
// 	"net/http"
// 	"strconv"
// 	"sync"
// 	"time"

// 	"github.com/labstack/echo/v4"
// )

// type (
// 	Stats struct {
// 		Uptime       time.Time      `json:"uptime"`
// 		RequestCount uint64         `json:"requestCount"`
// 		Statuses     map[string]int `json:"statuses"`
// 		mutex        sync.RWMutex
// 	}

// )

// func NewStats() *Stats {
// 	return &Stats{
// 		Uptime:   time.Now(),
// 		Statuses: map[string]int{},
// 	}
// }

// // Process is the middleware function.
// func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		if err := next(c); err != nil {
// 			c.Error(err)
// 		}
// 		s.mutex.Lock()
// 		defer s.mutex.Unlock()
// 		s.RequestCount++
// 		status := strconv.Itoa(c.Response().Status)
// 		s.Statuses[status]++
// 		return nil
// 	}
// }

// // Handle is the endpoint to get stats.
// func (s *Stats) Handle(c echo.Context) error {
// 	s.mutex.RLock()
// 	defer s.mutex.RUnlock()
// 	return c.JSON(http.StatusOK, s)
// }

// // ServerHeader middleware adds a `Server` header to the response.
// func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
// 		return next(c)
// 	}
// }

// func main() {
// 	e := echo.New()

// 	// Debug mode
// 	e.Debug = true

// 	//-------------------
// 	// Custom middleware
// 	//-------------------
// 	// Stats
// 	s := NewStats()
// 	e.Use(s.Process)
// 	e.GET("/stats", s.Handle) // Endpoint to get stats

// 	// Server header
// 	e.Use(ServerHeader)

// 	// Handler
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})

// 	// Start server
// 	e.Logger.Fatal(e.Start(":1323"))
// }

// package main

// import (
// 	"net/http"

// 	"github.com/go-playground/validator"
// 	"github.com/labstack/echo/v4"
// )

// type (
//   User struct {
//     Name  string `json:"name" validate:"required"`
//     Email string `json:"email" validate:"required,email"`
//   }

//   CustomValidator struct {
//     validator *validator.Validate
//   }
// )

// func (cv *CustomValidator) Validate(i interface{}) error {
//   if err := cv.validator.Struct(i); err != nil {
//        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//   }
//   return nil
// }

// func BeforeResponse(c echo.Context) error {
// 	c.Response().Before(func() {
// 	  println("before response")
// 	})
// 	c.Response().After(func() {
// 	  println("after response")
// 	})
// 	return c.JSON(http.StatusOK,"succcessful")
//   }

// func main() {
//   e := echo.New()
//   e.GET("/check",BeforeResponse)
// //   e.Validator = &CustomValidator{validator: validator.New()}
// //   e.POST("/users", func(c echo.Context) (err error) {
// //     u := new(User)
// //     if err = c.Bind(u); err != nil {
// //       return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// //     }
// //     if err = c.Validate(u); err != nil {
// //       return err
// //     }
// //     return c.JSON(http.StatusOK, u)
// //   })
//   e.Logger.Fatal(e.Start(":1323"))
// }

// package main

// import (
// 	"html/template"
// 	"io"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// // TemplateRenderer is a custom html/template renderer for Echo framework
// type TemplateRenderer struct {
// 	templates *template.Template
// }

// // Render renders a template document
// func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

// 	// Add global methods if data is a map
// 	if viewContext, isMap := data.(map[string]interface{}); isMap {
// 		viewContext["reverse"] = c.Echo().Reverse
// 	}

// 	return t.templates.ExecuteTemplate(w, name, data)
// }

func main() {
  e := echo.New()
  renderer := &TemplateRenderer{
      templates: template.Must(template.ParseGlob("*.html")),
  }
  e.Renderer = renderer

  // Named route "foobar"
  e.GET("/something", func(c echo.Context) error {
      return c.Render(http.StatusOK, "template.html", map[string]interface{}{
          "name": "Dolly!",
      })
  }).Name = "foobar"

  e.Logger.Fatal(e.Start(":8000"))
}

package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

type Handler struct {
	db map[string]*User
}

func (h *Handler) CreateUser(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    validator := validator.New()
    if err := validator.Struct(u); err != nil {
		fmt.Println("error")
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    h.db[u.Email] = u

    return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()
	h := &Handler{
		db: make(map[string]*User),
	}
	e.POST("/users", h.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
