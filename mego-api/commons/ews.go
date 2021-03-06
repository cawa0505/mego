package commons

import (
	"github.com/mhewedy/ews"
	"github.com/mhewedy/go-conf"
)

// used for testing to be assigned by mocked value
var DefaultEWSClient ews.Client

func NewEWSClient(username, password string) ews.Client {

	if DefaultEWSClient != nil {
		return DefaultEWSClient
	}

	if dn := conf.Get("ews.ad_domain_name", ""); len(dn) > 0 {
		username = dn + "\\" + username
	}

	return ews.NewClient(
		conf.Get("ews.exchange_url"),
		username,
		password,
		&ews.Config{
			Dump:    conf.GetBool("ews.dump", false),
			NTLM:    conf.GetBool("ews.ntlm", true),
			SkipTLS: conf.GetBool("ews.skip_tls", false),
		},
	)
}
