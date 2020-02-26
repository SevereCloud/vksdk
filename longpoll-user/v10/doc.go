/*
Package wrapper implement User Long Poll API wrapper v10.

	w := wrapper.NewWrapper(&lp)

	w.OnNewMessage(func(m wrapper.NewMessage) {
		fmt.Printf("4 wrapper.NewMessage: %v\n", m)
	})

Documentation https://vk.com/dev/using_longpoll
*/
package wrapper // import "github.com/SevereCloud/vksdk/longpoll-user/v10"
