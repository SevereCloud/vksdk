package oauth_test

import (
	"net/url"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/oauth"
	"github.com/stretchr/testify/assert"
)

func TestNewGroupTokensFromJSON(t *testing.T) {
	f := func(data []byte, wantToken *oauth.GroupTokens, wantErr string) {
		token, err := oauth.NewGroupTokensFromJSON(data)
		if err != nil {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, token, wantToken)
	}

	f([]byte("123"), nil, "json: cannot unmarshal number into Go value of type oauth.Error")
	f(
		[]byte(`{"groups":[{"access_token":"533bacf01e11f55b536a565b57531ac114461ae8736d6506a3","group_id":66748}],"expires_in":43200}`),
		&oauth.GroupTokens{
			Groups: []oauth.GroupToken{
				{
					GroupID:     66748,
					AccessToken: "533bacf01e11f55b536a565b57531ac114461ae8736d6506a3",
				},
			},
			ExpiresIn: 43200,
		},
		"",
	)
	f(
		[]byte(`{"error":"invalid_grant","error_description":"Code is expired."}`),
		nil,
		"oauth: Code is expired.",
	)
}

func TestNewGroupTokensFromURL(t *testing.T) {
	f := func(rawurl string, wantToken *oauth.GroupTokens, wantErr string) {
		u, _ := url.Parse(rawurl)

		token, err := oauth.NewGroupTokensFromURL(u)
		if err != nil {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, token, wantToken)
	}

	f(
		"http://REDIRECT_URI#&test&access_token_8492=533bacf01e11f55b536a565b57531ad114461ae8736d6506a3&expires_in=86400",
		&oauth.GroupTokens{
			Groups: []oauth.GroupToken{
				{
					GroupID:     8492,
					AccessToken: "533bacf01e11f55b536a565b57531ad114461ae8736d6506a3",
				},
			},
			ExpiresIn: 86400,
		},
		"",
	)
	f(
		"http://REDIRECT_URI#error=access_denied&error_description=The+user+or+authorization+server+denied+the+request.",
		nil,
		"oauth: The user or authorization server denied the request.",
	)
}

func TestNewGroupTokensFromURL_Error(t *testing.T) {
	_, err := oauth.NewGroupTokensFromURL(&url.URL{
		Fragment: "%gh&%ij",
	})
	assert.EqualError(t, err, `invalid URL escape "%gh"`)
}
