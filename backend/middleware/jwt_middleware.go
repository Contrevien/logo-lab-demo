package middleware

import (
	"context"
	"log"
	"logo-lab-demo/domain"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct{}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

// EnsureValidToken is a middleware that will check the validity of our JWT.
func JwtAuthMiddleware(auth_domain string, audience string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]

			issuerURL, err := url.Parse("https://" + auth_domain + "/")
			if err != nil {
				log.Fatalf("Failed to parse the issuer url: %v", err)
			}

			provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

			jwtValidator, err := validator.New(
				provider.KeyFunc,
				validator.RS256,
				issuerURL.String(),
				[]string{audience},
				validator.WithCustomClaims(
					func() validator.CustomClaims {
						return &CustomClaims{}
					},
				),
				validator.WithAllowedClockSkew(time.Minute),
			)
			if err != nil {
				log.Fatalf("Failed to set up the jwt validator")
			}

			claims, err := jwtValidator.ValidateToken(c, authToken)

			if err != nil {
				c.Set("x-user-id", claims.(*validator.ValidatedClaims).RegisteredClaims.Subject)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
