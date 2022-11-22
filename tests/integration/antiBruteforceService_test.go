package tests

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/Tyrqvir/anti-brute-force/internal/config"
	"github.com/Tyrqvir/anti-brute-force/proto/api"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type antiBruteForceServiceTest struct {
	client       api.AntiBruteForceServiceClient
	redisClient  *redis.Client
	lastResponse bool
	login        string
	password     string
	ip           string
}

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "../../configs/config.toml", "Path to configuration file")
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	flag.Parse()

	service := &antiBruteForceServiceTest{}

	cfg, err := NewConfig(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	_ = service.connectToServer(cfg.GRPC.Address)

	service.redisClient = redisClient(cfg.Redis.DSN)

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		service.redisClient.FlushAll(context.Background())
		return ctx, nil
	})

	ctx.Step(`^I set ip "([^"]*)"$`, service.iSetIP)
	ctx.Step(`^I set login "([^"]*)"$`, service.iSetLogin)
	ctx.Step(`^I set password "([^"]*)"$`, service.iSetPassword)
	ctx.Step(`^I send (\d+) requests to authorisation$`, service.iSendRequestsToAuthorisation)
	ctx.Step(`^the requests are not blocked$`, service.theRequestsAreNotBlocked)
	ctx.Step(`^the requests are blocked$`, service.theRequestsAreBlocked)
	ctx.Step(`^I reset bucket for login "([^"]*)" and password "([^"]*)"$`, service.iResetBucketForLoginAndPassword)
	ctx.Step(`^I add network "([^"]*)" to black list$`, service.iAddNetworkToBlackList)
	ctx.Step(`^I add network "([^"]*)" to white list$`, service.iAddNetworkToWhiteList)
	ctx.Step(`^I see network "([^"]*)" in black list$`, service.iSeeNetworkInBlackList)
	ctx.Step(`^I see network "([^"]*)" in white list$`, service.iSeeNetworkInWhiteList)
	ctx.Step(`^I don\'t see network "([^"]*)" in white list$`, service.iDontSeeNetworkInWhiteList)
	ctx.Step(`^I don\'t see network "([^"]*)" in black list$`, service.iDontSeeNetworkInBlackList)
	ctx.Step(`^I remove network "([^"]*)" from white list$`, service.iRemoveNetworkFromWhiteList)
	ctx.Step(`^I remove network "([^"]*)" from black list$`, service.iRemoveNetworkFromBlackList)
}

func (service *antiBruteForceServiceTest) iDontSeeNetworkInWhiteList(network string) error {
	result, err := service.client.ExistInWhiteList(context.Background(), &api.ListCases{SubNetwork: network})
	if err != nil {
		return err
	}
	if result.GetOk() == 1 {
		return errors.New("should be 0")
	}
	return nil
}

func (service *antiBruteForceServiceTest) iRemoveNetworkFromWhiteList(network string) error {
	_, err := service.client.RemoveFromWhiteList(context.Background(), &api.ListCases{SubNetwork: network})
	if err != nil {
		return err
	}

	return nil
}

func (service *antiBruteForceServiceTest) iDontSeeNetworkInBlackList(network string) error {
	result, err := service.client.ExistInBlackList(context.Background(), &api.ListCases{SubNetwork: network})
	if err != nil {
		return err
	}
	if result.GetOk() == 1 {
		return errors.New("should be 0")
	}
	return nil
}

func (service *antiBruteForceServiceTest) iRemoveNetworkFromBlackList(network string) error {
	_, err := service.client.RemoveFromBlackList(context.Background(), &api.ListCases{SubNetwork: network})
	if err != nil {
		return err
	}

	return nil
}

func (service *antiBruteForceServiceTest) iSeeNetworkInBlackList(network string) error {
	result, err := service.client.ExistInBlackList(context.Background(), &api.ListCases{
		SubNetwork: network,
	})
	if err != nil {
		return err
	}

	if result.GetOk() != 1 {
		return errors.New("should be 1")
	}

	return nil
}

func (service *antiBruteForceServiceTest) iAddNetworkToBlackList(network string) error {
	_, err := service.client.AddToBlackList(context.Background(), &api.ListCases{
		SubNetwork: network,
	})
	if err != nil {
		return err
	}
	return nil
}

func (service *antiBruteForceServiceTest) iSeeNetworkInWhiteList(network string) error {
	result, err := service.client.ExistInWhiteList(context.Background(), &api.ListCases{
		SubNetwork: network,
	})
	if err != nil {
		return err
	}

	if result.GetOk() != 1 {
		return errors.New("should be 1")
	}

	return nil
}

func (service *antiBruteForceServiceTest) iAddNetworkToWhiteList(network string) error {
	_, err := service.client.AddToWhiteList(context.Background(), &api.ListCases{
		SubNetwork: network,
	})
	if err != nil {
		return err
	}
	return nil
}

func (service *antiBruteForceServiceTest) iSetIP(value string) error {
	service.ip = value
	return nil
}

func (service *antiBruteForceServiceTest) iSetLogin(value string) error {
	service.login = value
	return nil
}

func (service *antiBruteForceServiceTest) iSetPassword(value string) error {
	service.password = value
	return nil
}

func (service *antiBruteForceServiceTest) iResetBucketForLoginAndPassword(login, password string) error {
	_, err := service.client.ResetBucket(context.Background(), &api.ResetBucketRequest{
		Login: login,
		Ip:    password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (service *antiBruteForceServiceTest) connectToServer(address string) error {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	service.client = api.NewAntiBruteForceServiceClient(conn)

	return nil
}

func (service *antiBruteForceServiceTest) iSendRequestsToAuthorisation(count int) error {
	for i := 0; i <= count; i++ {
		authorisation, err := service.client.Authorisation(context.Background(), &api.AuthorisationRequest{
			Login:    service.login,
			Password: service.password,
			Ip:       service.ip,
		})
		if err != nil {
			return fmt.Errorf("can't authorise: %w", err)
		}

		service.lastResponse = authorisation.GetOk()
	}

	return nil
}

func (service *antiBruteForceServiceTest) theRequestsAreNotBlocked() error {
	if !service.lastResponse {
		return fmt.Errorf("should be true")
	}
	return nil
}

func (service *antiBruteForceServiceTest) theRequestsAreBlocked() error {
	if service.lastResponse {
		return fmt.Errorf("should be false")
	}
	return nil
}

func TestMain(m *testing.M) {
	flag.Parse()
	format := "progress"
	for _, arg := range os.Args[1:] {
		if arg == "-test.v=true" { // go test transforms -v option
			format = "pretty"
			break
		}
	}

	opts := godog.Options{
		Format:    format,
		Output:    colors.Colored(os.Stdout),
		Paths:     []string{"features"},
		Randomize: 0,
	}

	status := godog.TestSuite{
		Name:                "integration-tests",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

func NewConfig(configFile string) (*config.Config, error) {
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("%v: %w", config.ErrFailedReadConfigFile, err)
	}

	cfg := new(config.Config)

	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %w", err)
	}

	return cfg, nil
}

func redisClient(dsn string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: dsn,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("error connect to redis: %v", err)
	}

	return rdb
}
