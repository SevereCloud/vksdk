package e2e

import (
	"os"
	"sync"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"
	packer "github.com/SevereCloud/vksdk/v2/packer"
	"github.com/stretchr/testify/assert"
)

func TestManyAPICalls(t *testing.T) {
	token := os.Getenv("USER_TOKEN")
	if token == "" {
		t.Skip("USER_TOKEN empty")
	}

	vk := api.NewVK(token)
	vk.Limit = api.LimitUserToken
	packer.Default(vk)
	var wg sync.WaitGroup
	num := 500
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			defer wg.Done()
			resp, err := vk.UtilsResolveScreenName(api.Params{
				"screen_name": "durov",
			})
			assert.Nil(t, err)
			assert.Equal(t, 1, resp.ObjectID)
		}(i)
	}
	wg.Wait()
}
