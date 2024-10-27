package values

import (
	"fmt"
	"os"
	"time"
)

const InvalidHost = "The host is not a subdomain of the domain"

var domain = os.Getenv("DOMAIN")

type CNAME struct {
	Host      string    `json:"host"`
	Alias     string    `json:"alias"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewCNAME(host, alias string) CNAME {
	return CNAME{
		Host:      host,
		Alias:     alias,
		CreatedAt: time.Now(),
	}
}

func (data CNAME) validate() error {
	// check if the host is a submomain of the domain
	if fmt.Sprintf("%v.%v", data.Host, domain) != domain {
		return fmt.Errorf(InvalidHost)
	}

	return nil
}
