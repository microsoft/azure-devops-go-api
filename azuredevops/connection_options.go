package azuredevops

import "net/http"

type ConnectionOptionFunc func(*Connection)

// WithConnectionHTTPClient sets the HTTP client that will be used by the connection.
func WithConnectionHTTPClient(client *http.Client) ConnectionOptionFunc {
	return func(c *Connection) {
		c.httpClient = client
	}
}

// WithConnectionPersonalAccessToken specifies the PAT to use for authentication with Azure DevOps.
func WithConnectionPersonalAccessToken(personalAccessToken string) ConnectionOptionFunc {
	return func(c *Connection) {
		authorizationString := CreateBasicAuthHeaderValue("", personalAccessToken)
		c.AuthorizationString = authorizationString
	}
}
