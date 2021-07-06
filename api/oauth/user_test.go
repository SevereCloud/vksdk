package oauth_test

import (
	"net/url"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/oauth"
	"github.com/stretchr/testify/assert"
)

func TestParseJSON(t *testing.T) {
	f := func(data []byte, wantToken *oauth.UserToken, wantErr string) {
		token, err := oauth.NewUserTokenFromJSON(data)
		if err != nil {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, token, wantToken)
	}

	f([]byte("123"), nil, "json: cannot unmarshal number into Go value of type oauth.Error")
	f(
		[]byte(`{"access_token":"533bacf01e11f55b536a565b57531ac114461ae8736d6506a3", "expires_in":43200, "user_id":66748}`),
		&oauth.UserToken{
			AccessToken: "533bacf01e11f55b536a565b57531ac114461ae8736d6506a3",
			ExpiresIn:   43200,
			UserID:      66748,
		},
		"",
	)
	f(
		[]byte(`{"error":"invalid_grant","error_description":"Code is expired."}`),
		nil,
		"oauth: Code is expired.",
	)
}

func TestParseURL(t *testing.T) {
	f := func(rawurl string, wantToken *oauth.UserToken, wantErr string) {
		u, _ := url.Parse(rawurl)

		token, err := oauth.NewUserTokenFromURL(u)
		if err != nil {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, token, wantToken)
	}

	f(
		"http://REDIRECT_URI#access_token=533bacf01e11f55b536a565b57531ad114461ae8736d6506a3&expires_in=86400&user_id=8492&state=123456",
		&oauth.UserToken{
			AccessToken: "533bacf01e11f55b536a565b57531ad114461ae8736d6506a3",
			ExpiresIn:   86400,
			UserID:      8492,
			State:       "123456",
		},
		"",
	)
	f(
		"http://REDIRECT_URI#error=access_denied&error_description=The+user+or+authorization+server+denied+the+request.",
		nil,
		"oauth: The user or authorization server denied the request.",
	)
}

func TestParseURL_Error(t *testing.T) {
	_, err := oauth.NewUserTokenFromURL(&url.URL{
		Fragment: "%gh&%ij",
	})
	assert.EqualError(t, err, `invalid URL escape "%gh"`)
}
