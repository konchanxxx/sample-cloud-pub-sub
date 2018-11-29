package environment

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Environments struct {
	GCP GCP `envconfig:"GCP"`
}

type GCP struct {
	ProjectID string `split_words:"true"`
}

func NewEnvironments() *Environments {
	e := &Environments{}
	err := envconfig.Process("", e)
	if err != nil {
		log.Fatal(err)
	}

	return e
}
