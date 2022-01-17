package oauth_test

import (
	"errors"
	"net/url"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/oauth"
	"github.com/stretchr/testify/assert"
)

func TestNewGroupTokensFromJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, wantToken *oauth.GroupTokens, wantErr string) {
		t.Helper()

		token, err := oauth.NewGroupTokensFromJSON(data)
		if err != nil || wantErr != "" {
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
	t.Parallel()

	f := func(rawurl string, wantToken *oauth.GroupTokens, wantErr string) {
		t.Helper()

		u, _ := url.Parse(rawurl)

		token, err := oauth.NewGroupTokensFromURL(u)
		if err != nil || wantErr != "" {
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
	t.Parallel()

	_, err := oauth.NewGroupTokensFromURL(&url.URL{
		Fragment: "%gh&%ij",
	})
	assert.EqualError(t, err, `invalid URL escape "%gh"`)
}

func TestAuthCodeFlowGroup_URL(t *testing.T) {
	t.Parallel()

	f := func(acf *oauth.AuthCodeFlowGroup, wantURL string) {
		t.Helper()

		assert.Equal(t, wantURL, acf.URL().String())
	}

	f(oauth.NewAuthCodeFlowGroup(oauth.GroupParams{
		ClientID:    6888183,
		GroupIDs:    []int{1234},
		Scope:       oauth.ScopeGroupPhotos + oauth.ScopeGroupDocs,
		RedirectURI: "https://example.com/callback",
		V:           "5.100",
	}, "clientSecret"), "https://oauth.vk.com/authorize?client_id=6888183&display=&group_ids=1234&redirect_uri=https%3A%2F%2Fexample.com%2Fcallback&response_type=code&scope=131076&state=&v=5.100")
	f(oauth.NewAuthCodeFlowGroup(oauth.GroupParams{
		ClientID: 6888183,
		GroupIDs: []int{1234, 321},
		Scope:    oauth.ScopeGroupPhotos + oauth.ScopeGroupDocs,
	}, "clientSecret"), "https://oauth.vk.com/authorize?client_id=6888183&display=&group_ids=1234%2C321&redirect_uri=https%3A%2F%2Foauth.vk.com%2Fblank.html&response_type=code&scope=131076&state=&v=5.131")
}

func TestAuthCodeFlowGroup_Token(t *testing.T) {
	t.Parallel()

	acf := oauth.NewAuthCodeFlowGroup(oauth.GroupParams{
		ClientID: 6888183,
		GroupIDs: []int{1234},
		Scope:    oauth.ScopeGroupPhotos + oauth.ScopeGroupDocs,
	}, "clientSecretGroup")

	u, _ := url.Parse("https://oauth.vk.com/blank.html#code=2fb239386220842e7c")

	_, err := acf.Token(u)

	var e *oauth.Error
	if !errors.As(err, &e) {
		t.Error(err)
	}

	f := func(rawURL, errString string) {
		t.Helper()

		u, _ := url.Parse(rawURL)
		_, err := acf.Token(u)
		assert.EqualError(t, err, errString)
	}

	// f("https://oauth.vk.com/blank.html#;", "invalid semicolon separator in query")
	f("https://oauth.vk.com/blank.html#error=invalid_request&error_description=Invalid+display+parameter", "oauth: Invalid display parameter")
}

func TestImplicitFlowGroup(t *testing.T) {
	t.Parallel()

	f := func(p oauth.GroupParams, wantURL string) {
		t.Helper()

		assert.Equal(t, wantURL, oauth.ImplicitFlowGroup(p).String())
	}

	f(oauth.GroupParams{
		ClientID:    6888183,
		GroupIDs:    []int{1234},
		Scope:       oauth.ScopeGroupPhotos + oauth.ScopeGroupDocs,
		RedirectURI: "https://example.com/callback",
		V:           "5.100",
	}, "https://oauth.vk.com/authorize?client_id=6888183&display=&group_ids=1234&redirect_uri=https%3A%2F%2Fexample.com%2Fcallback&response_type=token&scope=131076&state=&v=5.100")
	f(oauth.GroupParams{
		ClientID: 6888183,
		GroupIDs: []int{1234},
		Scope:    oauth.ScopeGroupPhotos + oauth.ScopeGroupDocs,
	}, "https://oauth.vk.com/authorize?client_id=6888183&display=&group_ids=1234&redirect_uri=https%3A%2F%2Foauth.vk.com%2Fblank.html&response_type=token&scope=131076&state=&v=5.131")
}
