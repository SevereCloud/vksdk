package oauth_test

import (
	"errors"
	"net/url"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/oauth"
	"github.com/stretchr/testify/assert"
)

func TestParseJSON(t *testing.T) {
	t.Parallel()

	f := func(data []byte, wantToken *oauth.UserToken, wantErr string) {
		token, err := oauth.NewUserTokenFromJSON(data)
		if err != nil || wantErr != "" {
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
	t.Parallel()

	f := func(rawurl string, wantToken *oauth.UserToken, wantErr string) {
		t.Helper()

		u, _ := url.Parse(rawurl)

		token, err := oauth.NewUserTokenFromURL(u)
		if err != nil || wantErr != "" {
			assert.EqualError(t, err, wantErr)
		}

		assert.Equal(t, token, wantToken)
	}

	f(
		"http://REDIRECT_URI#access_token=533bacf01e11f55b536a565b57531ad114461ae8736d6506a3&expires_in=86400&user_id=8492&state=6888183",
		&oauth.UserToken{
			AccessToken: "533bacf01e11f55b536a565b57531ad114461ae8736d6506a3",
			ExpiresIn:   86400,
			UserID:      8492,
			State:       "6888183",
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
	t.Parallel()

	_, err := oauth.NewUserTokenFromURL(&url.URL{
		Fragment: "%gh&%ij",
	})
	assert.EqualError(t, err, `invalid URL escape "%gh"`)
}

func TestAuthCodeFlowUser_URL(t *testing.T) {
	t.Parallel()

	f := func(acf *oauth.AuthCodeFlowUser, wantURL string) {
		t.Helper()

		assert.Equal(t, wantURL, acf.URL().String())
	}

	f(oauth.NewAuthCodeFlowUser(oauth.UserParams{
		ClientID:    6888183,
		Scope:       oauth.ScopeUserPhotos + oauth.ScopeUserDocs,
		RedirectURI: "https://example.com/callback",
		V:           "5.100",
	}, "clientSecret"), "https://oauth.vk.com/authorize?client_id=6888183&display=&redirect_uri=https%3A%2F%2Fexample.com%2Fcallback&response_type=code&scope=131076&state=&v=5.100")
	f(oauth.NewAuthCodeFlowUser(oauth.UserParams{
		ClientID: 6888183,
		Scope:    oauth.ScopeUserPhotos + oauth.ScopeUserDocs,
	}, "clientSecret"), "https://oauth.vk.com/authorize?client_id=6888183&display=&redirect_uri=https%3A%2F%2Foauth.vk.com%2Fblank.html&response_type=code&scope=131076&state=&v=5.131")
}

func TestAuthCodeFlowUser_Token(t *testing.T) {
	t.Parallel()

	acf := oauth.NewAuthCodeFlowUser(oauth.UserParams{
		ClientID: 6888183,
		Scope:    oauth.ScopeUserPhotos + oauth.ScopeUserDocs,
	}, "clientSecretUser")

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

func TestImplicitFlowUser(t *testing.T) {
	t.Parallel()

	f := func(p oauth.UserParams, wantURL string) {
		t.Helper()

		assert.Equal(t, wantURL, oauth.ImplicitFlowUser(p).String())
	}

	f(oauth.UserParams{
		ClientID:    6888183,
		Scope:       oauth.ScopeUserPhotos + oauth.ScopeUserDocs,
		RedirectURI: "https://example.com/callback",
		V:           "5.100",
	}, "https://oauth.vk.com/authorize?client_id=6888183&display=&redirect_uri=https%3A%2F%2Fexample.com%2Fcallback&response_type=token&scope=131076&state=&v=5.100")
	f(oauth.UserParams{
		ClientID: 6888183,
		Scope:    oauth.ScopeUserPhotos + oauth.ScopeUserDocs,
	}, "https://oauth.vk.com/authorize?client_id=6888183&display=&redirect_uri=https%3A%2F%2Foauth.vk.com%2Fblank.html&response_type=token&scope=131076&state=&v=5.131")
}

func TestDirectAuth(t *testing.T) {
	t.Parallel()

	params := oauth.DirectAuthParams{
		ClientSecret:       "test",
		ClientID:           6888183,
		Username:           "username",
		Password:           "",
		Scope:              0,
		V:                  "5.131",
		TwoFactorSupported: true,
		ForceSMS:           true,
		Code:               "code",
		CaptchaSID:         "captcha_sid",
		CaptchaKey:         "captcha_key",
		TestRedirectURI:    true,
	}

	_, err := oauth.DirectAuth(params)

	var e *oauth.Error
	if !errors.As(err, &e) {
		t.Error(err)
	}
}
