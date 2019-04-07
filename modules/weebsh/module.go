package weebsh

import (
	"fmt"
	"github.com/Daniele122898/weeb.go/src"
	"github.com/Daniele122898/weeb.go/src/data"
	"github.com/moedevs/Vigne/commands"
	"github.com/moedevs/Vigne/database/interfaces"
	"github.com/moedevs/Vigne/errors"
	"github.com/moedevs/Vigne/server"
	"strings"
	"time"
)

type WeebshModule struct {
	typesCache        []string
	TypesCacheExpires time.Time
	Authenticated     bool
	Token             interfaces.StringValue
	TokenType         interfaces.StringValue
	Container         interfaces.Container
}

func (m WeebshModule) GetName() string {
	return "weeb.sh"
}

func (m *WeebshModule) Init(server *server.Server) error {
	m.typesCache = []string{}
	cmdInterface, err := server.GetModuleByName("Commands")
	if err != nil {
		return err
	}
	cmd := cmdInterface.(*commands.CommandsModule)
	cmd.RegisterCommand(&TypesCommand{m})
	cmd.RegisterCommand(&ImageCommand{module:m})
	m.Container = server.Database.GetContainer("weebsh")
	m.Token = m.Container.Value("token")
	m.TokenType = m.Container.Value("tokenType")
	return nil
}

func (m *WeebshModule) Setup() error {
	if !m.Authenticated {
		token := m.Token.Get()
		if token == "" {
			return errors.New("weebsh:token is missing", "This service is unavailable")
		}
		tokenTypeS := m.TokenType.Get()
		if tokenTypeS == "" {
			return errors.New("weebsh:tokenType is missing", "This service is unavailable")
		}
		tokenTypeS = strings.ToLower(tokenTypeS)
		var tokenType data.TokenType
		switch tokenTypeS {
		case "bearer":
			tokenType = data.BEARER
			break
		case "wolke":
			tokenType = data.WOLKE
			break
		default:
			return errors.New("weebsh:tokenType is neither bearer nor wolke", "This service is unavailable")
		}

		err := weebgo.Authenticate(token, tokenType)
		if err != nil {
			return errors.New("Couldn't authenticate.", "This service is unavailable")
		}

		m.Authenticated = true
	}
	return nil
}

func (m *WeebshModule) updateTypes()  {
	if time.Now().After(m.TypesCacheExpires) {
		fmt.Println("Updating types cache...")
		data, err := weebgo.GetTypes(false)
		if err != nil {
			return
		}
		m.typesCache = data.Types
		m.TypesCacheExpires = time.Now().Add(time.Hour)
	}
}

func (m *WeebshModule) GetTypes() ([]string, error) {
	err := m.Setup()
	if err != nil {
		return nil, err
	}
	m.updateTypes()
	if m.typesCache == nil {
		return nil, errors.New("Couldn't fetch types", "This service is not available")
	}
	return m.typesCache, nil
}

func (m *WeebshModule) GetImageByType(typeName string) (string, error) {
	err := m.Setup()
	if err != nil {
		return "", err
	}
	image, err := weebgo.GetRandomImage(typeName, nil, data.ANY, data.FALSE, false)
	if err != nil {
		return "", err
	}
	return image.Url, nil
}