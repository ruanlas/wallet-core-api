package idpauth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Nerzal/gocloak/v13"
)

type User struct {
	Id       string `json:"sub"`
	Username string `json:"preferred_username"`
}

const AUTH_HEADER = "X-Access-Token"

var allowedPaths = []string{
	"/swagger/*any",
}

func AuthenticationMiddleware(ctx *gin.Context) {
	log.Println("----Authentication------")
	if slices.Contains(allowedPaths, ctx.FullPath()) {
		return
	}
	idpAddress := fmt.Sprintf("%s:%s", os.Getenv("IDP_HOST"), os.Getenv("IDP_PORT"))
	client := gocloak.NewClient(idpAddress)

	accessToken := ctx.GetHeader(AUTH_HEADER)
	rptResult, err := client.RetrospectToken(ctx, accessToken, os.Getenv("IDP_CLIENT_IDENTIFIER"), os.Getenv("IDP_CLIENT_SECRET"), os.Getenv("IDP_REALM"))
	if err != nil {
		log.Println("Retrospect Token failed:" + err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "Error"})
	}
	if !*rptResult.Active {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "Negado"})
	}
}

func AuthorizationMiddleware(ctx *gin.Context) {
	// log.Println(ctx.Request.URL)
	// log.Println(ctx.Request.Method)
	// fmt.Println(ctx.HandlerName())

	// fmt.Println(ctx.Request.RequestURI)
	// fmt.Println(ctx.FullPath())
	// fmt.Println(ctx.Request.URL.Path)
	// ctx.Request.Header.Add("X-EXTRA", "TESTE123")
	// ctx.Header("X-EXTRA2", "123TESTE")
	log.Println("----Authorization------")
	// if slices.Contains(allowedPaths, ctx.FullPath()) {
	// 	return
	// }
	accessToken := ctx.GetHeader(AUTH_HEADER)
	user := GetUser(accessToken)
	if user != nil {
		fmt.Println(*user)
	}
}

func GetUser(accessToken string) *User {
	if accessToken == "" {
		return nil
	}
	payload := splitToken(accessToken)[1]
	userDecodedRaw, err := base64.RawStdEncoding.DecodeString(payload)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	var user User
	err = json.Unmarshal(userDecodedRaw, &user)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &user
}

func splitToken(accessToken string) []string {
	return strings.Split(accessToken, ".")
}
