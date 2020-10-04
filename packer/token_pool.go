package packer

type tokenPool struct {
	tokens chan string
}

func newTokenPool(tokens ...string) *tokenPool {
	c := make(chan string, len(tokens))
	for _, t := range tokens {
		c <- t
	}

	return &tokenPool{
		tokens: c,
	}
}

func (tp *tokenPool) Get() string {
	token := <-tp.tokens
	tp.tokens <- token

	return token
}
