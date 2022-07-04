package controller

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/Luka-Spa/GoAPI/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

type Payload struct {
	Aud         []string
	Azp         string
	Iat         int
	Iss         string
	Permissions []string
	Scope       string
	Sub         string
}

func UseAuthorisation(permissions []string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		rawToken := c.GetHeader("Authorization")
		data, err := decode(rawToken)
		if err != nil {
			log.Error(err)
			c.AbortWithStatusJSON(403, gin.H{"message": model.ErrUnauthorized.Error()})
			return
		}
		for _, s := range permissions {
			if slices.Contains(data.Permissions, s) {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(403, gin.H{"message": model.ErrUnauthorized.Error()})
	}
	return gin.HandlerFunc(fn)
}

func decode(token string) (Payload, error) {
	var data Payload
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return data, model.ErrInvalidToken
	}
	bytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return data, model.ErrInvalidToken
	}
	err = json.Unmarshal([]byte(bytes), &data)
	if err != nil {
		return data, model.ErrInvalidToken
	}
	return data, nil
}
