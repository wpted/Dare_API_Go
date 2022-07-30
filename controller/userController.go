package controller

import (
	"dareAPI/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var MockUser = model.User{
	UserName: "Admin",
	Password: "Password",
}
var MockSecretKey = "BDABF2B3DF0E000B2C927DCF1E2235320AC3256AB0E97B6CFD390490B43582FA"

type AuthHandler struct{}

// Claims is a set of statement made by the creator to tell info about the subject
type Claims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// JWTOutput is the output format of the combination of a string token according to the payload and the expiration time
type JWTOutput struct {
	Token   string
	Expires time.Time
}

// AuthMiddleWare is a translation middle layer for implementing JWT tokens to endpoints
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		claims := &Claims{}

		parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(MockSecretKey), nil
		})

		if err != nil || parsedToken == nil || !parsedToken.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

// SignInHandler is a controller for getting the JWT token
func (a *AuthHandler) SignInHandler(c *gin.Context) {
	var user model.User

	// if user input matches the model.user struct, bind the input data to the user variable
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// verify whether the user input is within the database
	if user.UserName != MockUser.UserName || user.Password != MockUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authorized",
		})
		return
	}

	// expirationTime is 10 minutes after the request
	var expirationTime = time.Now().Add(10 * time.Minute)

	// create claims according to the user input
	claims := Claims{
		UserName: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			// NewNumericDate constructs a new *NumericDate from a standard library time.Time struct.
			// It will truncate the timestamp according to the precision specified in TimePrecision.
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// NewWithClaims creates a type *Token from the given hashing algorithm and the claims
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	// SignedString is Parsing and Validating the two: the given signature and the signature generated with the secret key
	tokenString, err := token.SignedString([]byte(MockSecretKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	output := JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}

	c.JSON(http.StatusOK, output)

}
