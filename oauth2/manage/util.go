package manage

import (
	"net/url"
	"strings"

	"github.com/wingfeng/idx/oauth2/errors"
)

type (
	// ValidateURIHandler validates that redirectURI is contained in baseURI
	ValidateURIHandler func(baseURI, redirectURI string) error
)

// DefaultValidateURI validates that redirectURI is contained in baseURI
func DefaultValidateURI(baseURI string, redirectURI string) error {
	urls := strings.Split(baseURI, ",")
	for _, s := range urls {
		base, err := url.Parse(s)
		if err != nil {
			return err
		}

		redirect, err := url.Parse(redirectURI)
		if err != nil {
			return err
		}
		if strings.HasSuffix(redirect.Host, base.Host) {
			return nil
		}
	}
	return errors.ErrInvalidRedirectURI
}
