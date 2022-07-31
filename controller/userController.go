package controller

import (
	"dareAPI/configs"
	"dareAPI/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type AuthHandler struct {
	configs.Admin
}

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
	if user.UserName != a.GetAdminName() || user.Password != a.GetAdminPwd() {
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
	tokenString, err := token.SignedString([]byte(a.GetSecretKey()))

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
