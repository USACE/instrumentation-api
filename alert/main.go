package main

import (
	"fmt"
	"log"
	"time"

	"github.com/USACE/instrumentation-api/api/dbutils"
	"github.com/USACE/instrumentation-api/api/handlers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/kelseyhightower/envconfig"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Config stores configuration information stored in environment variables
type Config struct {
	DBUser              string
	DBPass              string
	DBName              string
	DBHost              string
	DBSSLMode           string
	AWSECSTriggerMocked bool   `envconfig:"INSTRUMENTATION_AWS_ECS_TRIGGER_MOCKED"`
	AWSSESMocked        bool   `envconfig:"INSTRUMENTATION_AWS_SES_MOCKED"`
	AWSSESEmailSender   string `envconfig:"INSTRUMENTATION_AWS_SES_EMAIL_SENDER"`
	AWSSESDisableSSL    bool   `envconfig:"INSTRUMENTATION_AWS_SES_DISABLE_SSL"`
	AWSSESRegion        string `envconfig:"INSTRUMENTATION_AWS_SES_REGION"`
}

func awsConfig(cfg *Config) *aws.Config {
	awsConfig := aws.NewConfig().WithRegion(cfg.AWSSESRegion)
	awsConfig.WithDisableSSL(cfg.AWSSESDisableSSL)
	return awsConfig
}

func (c *Config) dbConnStr() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", c.DBUser, c.DBPass, c.DBName, c.DBHost, c.DBSSLMode)
}

func checkAlerts(cfg Config) {
	awsCfg := awsConfig(&cfg)
	sess := session.Must(session.NewSession(awsCfg))
	sesc := ses.New(sess)

	db := dbutils.Connection(cfg.dbConnStr())
	defer func() error {
		if err := db.Close(); err != nil {
			log.Fatal(err.Error())
		}
		return nil
	}()

	if err := handlers.DoAlertChecks(db, sesc, cfg.AWSSESEmailSender, cfg.AWSSESMocked); err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("successfully completed alert checks at %s", time.Now())
}

func main() {
	var cfg Config
	if err := envconfig.Process("instrumentation", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	if cfg.AWSECSTriggerMocked {
		for {
			checkAlerts(cfg)
			time.Sleep(15 * time.Second)
		}
	} else {
		checkAlerts(cfg)
	}
}
