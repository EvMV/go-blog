package provider

import "net/http"

type AuthorizeChecker struct {
	userProvider *UserProvider
}

func NewAuthorizeChecker(userProvider *UserProvider) *AuthorizeChecker {
	return &AuthorizeChecker{userProvider: userProvider}
}

func (c *AuthorizeChecker) Authorize(r *http.Request) error {
	_, err := c.userProvider.ProvideCurrentUser(r)

	return err
}
