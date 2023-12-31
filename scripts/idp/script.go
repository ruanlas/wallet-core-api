package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nerzal/gocloak/v13"
	"github.com/joho/godotenv"
)

var (
	tokenJwt       *gocloak.JWT
	keycloakClient *gocloak.GoCloak

	client           *gocloak.Client
	clientSecret     *gocloak.CredentialRepresentation
	clientIdentifier string
	realm            *gocloak.RealmRepresentation
	realmName        string

	fileEnvPath string

	enable bool
)

func init() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		panic("You must to pass path to env as arg. Example [ $ ./idp-config-init /path/to/.env ]")
	}
	fileEnvPath = argsWithoutProg[0]
	err := godotenv.Load(fileEnvPath)
	if err != nil {
		panic(err)
	}
	keycloakAddress := fmt.Sprintf("%s:%s", os.Getenv("IDP_HOST"), os.Getenv("IDP_PORT"))
	keycloakClient = gocloak.NewClient(keycloakAddress)
	realmName = os.Getenv("IDP_REALM")
	clientIdentifier = os.Getenv("IDP_CLIENT_IDENTIFIER")
	enable = true
}

func main() {
	ctx := context.Background()

	makeLoginAdmin(ctx)
	processRealm(ctx)
	processClient(ctx)
	updateEnvFile()
	processUserFixture(ctx)
}

func makeLoginAdmin(ctx context.Context) {
	var err error

	tokenJwt, err = keycloakClient.LoginAdmin(ctx, os.Getenv("IDP_USER_ADMIN"), os.Getenv("IDP_PASSWORD_ADMIN"), os.Getenv("IDP_MAIN_REALM"))
	if err != nil {
		log.Fatalln(err)
	}
}

func processRealm(ctx context.Context) {
	var apiErr *gocloak.APIError
	var err error

	log.Printf("Buscando o realm [ %s ]\n", realmName)
	realm, err = keycloakClient.GetRealm(ctx, tokenJwt.AccessToken, realmName)
	apiErr, _ = err.(*gocloak.APIError)
	if apiErr != nil && apiErr.Code == http.StatusNotFound {
		log.Printf("Não encontrou o realm [ %s ]\n", realmName)
		log.Printf("Criando o realm [ %s ]\n", realmName)

		_, err = keycloakClient.CreateRealm(ctx, tokenJwt.AccessToken, gocloak.RealmRepresentation{Realm: &realmName, Enabled: &enable})
		if err != nil {
			log.Fatalln(err)
		}
		realm, err = keycloakClient.GetRealm(ctx, tokenJwt.AccessToken, realmName)
		if err != nil {
			log.Fatalln(err)
		}
	} else if apiErr != nil {
		log.Fatalln(apiErr)
	} else {
		log.Printf("Encontrado o realm [ %s ]\n", realmName)
	}
}

func processClient(ctx context.Context) {
	var apiErr *gocloak.APIError
	var err error

	log.Printf("Buscando o client [ %s ] dentro do realm [ %s ]\n", clientIdentifier, realmName)
	clients, err := keycloakClient.GetClients(ctx, tokenJwt.AccessToken, *realm.Realm, gocloak.GetClientsParams{ClientID: &clientIdentifier})
	apiErr, _ = err.(*gocloak.APIError)
	if apiErr != nil {
		log.Fatalln(apiErr)
	}
	if len(clients) == 0 {
		log.Printf("Não encontrou o client [ %s ] dentro do realm [ %s ]\n", clientIdentifier, realmName)
		log.Printf("Criando o client [ %s ] dentro do realm [ %s ]\n", clientIdentifier, realmName)
		clientIdString, err := keycloakClient.CreateClient(ctx, tokenJwt.AccessToken, *realm.Realm, gocloak.Client{ClientID: &clientIdentifier, DirectAccessGrantsEnabled: &enable})
		if err != nil {
			log.Fatalln(err)
		}
		client, err = keycloakClient.GetClient(ctx, tokenJwt.AccessToken, *realm.Realm, clientIdString)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Encontrado o client [ %s ] dentro do realm [ %s ]\n", clientIdentifier, realmName)
		client = clients[0]
	}

	log.Printf("Gerando uma secret para o client [ %s ]\n", clientIdentifier)
	clientSecret, err = keycloakClient.RegenerateClientSecret(ctx, tokenJwt.AccessToken, *realm.Realm, *client.ID)
	apiErr, _ = err.(*gocloak.APIError)
	if apiErr != nil {
		log.Fatalln(apiErr)
	}
	log.Printf("Secret gerada para o client [ %s ]\n", clientIdentifier)
}

func processUserFixture(ctx context.Context) {
	var apiErr *gocloak.APIError
	var err error
	usernameTeste := "testeuser"
	password := "123456"

	log.Printf("Verificando a existência de um usuário de teste no realm [ %s ]\n", realmName)
	users, err := keycloakClient.GetUsers(ctx, tokenJwt.AccessToken, *realm.Realm, gocloak.GetUsersParams{Username: &usernameTeste})
	apiErr, _ = err.(*gocloak.APIError)
	if apiErr != nil {
		log.Fatalln(apiErr)
	}
	if len(users) == 0 {
		log.Printf("Criando o usuário de teste no realm [ %s ]\n", realmName)
		enable := true
		userId, err := keycloakClient.CreateUser(ctx, tokenJwt.AccessToken, *realm.Realm, gocloak.User{Username: &usernameTeste, Enabled: &enable})
		apiErr, _ = err.(*gocloak.APIError)
		if apiErr != nil {
			log.Fatalln(apiErr)
		}
		err = keycloakClient.SetPassword(ctx, tokenJwt.AccessToken, userId, *realm.Realm, password, false)
		apiErr, _ = err.(*gocloak.APIError)
		if apiErr != nil {
			log.Fatalln(apiErr)
		}
		generateUpdateSql(userId)
		log.Printf("Usuário de teste criado no realm [ %s ]\n", realmName)
	} else {
		log.Printf("Usuário de teste encontrado no realm [ %s ]\n", realmName)
		generateUpdateSql(*users[0].ID)
	}
	log.Println("---- Dados do usuário -------")
	log.Println("Usuário para teste: ", usernameTeste)
	log.Println("Senha do usuário: ", password)
}

func generateUpdateSql(userId string) {
	updateGainProjection := fmt.Sprintf(`UPDATE gain_projection SET user_id = "%s" WHERE user_id = "user1";`, userId)
	updateGain := fmt.Sprintf(`UPDATE gain SET user_id = "%s" WHERE user_id = "user1";`, userId)
	updateInvoiceProjection := fmt.Sprintf(`UPDATE invoice_projection SET user_id = "%s" WHERE user_id = "user1";`, userId)
	updateInvoice := fmt.Sprintf(`UPDATE invoice SET user_id = "%s" WHERE user_id = "user1";`, userId)

	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("Execute os seguintes comandos no banco de dados:")
	fmt.Println()
	fmt.Println(updateGainProjection)
	fmt.Println(updateGain)
	fmt.Println(updateInvoiceProjection)
	fmt.Println(updateInvoice)
	fmt.Println()
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	path := "/scripts/mysql"
	file := "/update_user.sql"

	f, err := os.Create(currentDir + path + file)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	_, err = f.WriteString("USE wallet_core;" + "\n")
	if err != nil {
		log.Println(err)
	}
	_, err = f.WriteString(updateGainProjection + "\n")
	if err != nil {
		log.Println(err)
	}
	_, err = f.WriteString(updateGain + "\n")
	if err != nil {
		log.Println(err)
	}
	_, err = f.WriteString(updateInvoiceProjection + "\n")
	if err != nil {
		log.Println(err)
	}
	_, err = f.WriteString(updateInvoice + "\n")
	if err != nil {
		log.Println(err)
	}
	f.Sync()
	w := bufio.NewWriter(f)
	w.Flush()

	fmt.Println()
	fmt.Println("O arquivo .sql com os scripts de atualização encontra-se em", path+file)
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println()
}

func updateEnvFile() {
	log.Println("Atualizando as variáveis de ambiente do arquivo .env")
	envs, err := godotenv.Read(fileEnvPath)
	if err != nil {
		log.Fatalln(err)
	}
	envs["IDP_CLIENT_SECRET"] = *clientSecret.Value
	godotenv.Write(envs, fileEnvPath)
	log.Println("Variáveis atualizadas")
}
