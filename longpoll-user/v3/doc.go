/*
Package wrapper implement User Long Poll API wrapper v3.

	w := wrapper.NewWrapper(&lp)

	w.OnNewMessage(func(m wrapper.NewMessage) {
		fmt.Printf("4 wrapper.NewMessage: %v\n", m)
	})

VK documentation https://dev.vk.com/ru/api/user-long-poll/getting-started
*/
package wrapper // import "github.com/SevereCloud/vksdk/v2/longpoll-user/v3"
