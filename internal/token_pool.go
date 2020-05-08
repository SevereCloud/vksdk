package internal

// TokenPool is a simple round-robin based token pool.
type TokenPool struct {
	tokens chan string
}

// NewTokenPool returns new token pool.
func NewTokenPool(tokens ...string) TokenPool {
	c := make(chan string, len(tokens))
	for _, t := range tokens {
		c <- t
	}

	return TokenPool{
		tokens: c,
	}
}

// Get returns access token from pool.
func (tp TokenPool) Get() string {
	token := <-tp.tokens
	tp.tokens <- token

	return token
}
