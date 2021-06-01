package games

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"
)

// ParamsVerification represents verification struct.
type ParamsVerification struct {
	ClientSecret string
}

// NewParamsVerification return *ParamsVerification.
func NewParamsVerification(clientSecret string) *ParamsVerification {
	pv := &ParamsVerification{
		ClientSecret: clientSecret,
	}

	return pv
}

// getSignParams return sort sign parameters.
func getSignParams(rawValues url.Values) string {
	singParams := make(url.Values)
	signKeys := strings.Split(rawValues["sign_keys"][0], ",")

	for _, key := range signKeys {
		values, ok := rawValues[key]
		if !ok {
			continue
		}

		for _, value := range values {
			singParams.Add(key, value)
		}
	}

	return singParams.Encode() // sorted by key.
}

// Sign return signature in base64.
func (pv *ParamsVerification) Sign(p []byte) string {
	// Generate hash code
	mac := hmac.New(sha256.New, []byte(pv.ClientSecret))
	_, _ = mac.Write(p)
	expectedMAC := mac.Sum(nil)

	// Generate base64
	base64Sign := base64.StdEncoding.EncodeToString(expectedMAC)
	base64Sign = strings.ReplaceAll(base64Sign, "+", "-")
	base64Sign = strings.ReplaceAll(base64Sign, "/", "_")
	base64Sign = strings.TrimRight(base64Sign, "=")

	return base64Sign
}

// Verify verifies the signature in URL.
func (pv *ParamsVerification) Verify(u *url.URL) (bool, error) {
	values, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return false, err
	}

	if len(values["sign"]) == 0 || len(values["sign_keys"]) == 0 {
		return false, nil
	}

	singParams := getSignParams(values)
	base64Sign := pv.Sign([]byte(singParams))

	return base64Sign == values["sign"][0], nil
}

// VerifyMiddleware func.
func (pv *ParamsVerification) VerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ok, err := pv.Verify(r.URL)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if ok {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// ParamsVerify verifies the signature in link using client secret.
func ParamsVerify(link, clientSecret string) (bool, error) {
	pv := NewParamsVerification(clientSecret)

	u, err := url.Parse(link)
	if err != nil {
		return false, err
	}

	return pv.Verify(u)
}
