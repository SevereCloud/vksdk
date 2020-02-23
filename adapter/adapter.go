package adapter

import (
	"context"
	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/longpoll-bot"
	"github.com/SevereCloud/vksdk/object"
	"github.com/go-joe/joe"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
)

type BotAdapter struct {
	vk      *api.VK
	lp      longpoll.Longpoll
	context context.Context
	logger  *zap.Logger
}

type Config struct {
	Token  string
	Logger *zap.Logger
}

var ErrGetBotInfo = errors.New("failed to get bot info")

func NewAdapter(ctx context.Context, conf Config) (*BotAdapter, error) {
	vk := api.Init(conf.Token)

	logger := conf.Logger
	if logger == nil {
		logger = zap.NewNop()
	}

	r, err := vk.GroupsGetByID(api.Params{})
	if err != nil {
		return nil, errors.Wrap(err, ErrGetBotInfo.Error())
	}

	if len(r) < 1 {
		return nil, ErrGetBotInfo
	}

	group := r[0]
	logger.Info("got group info",
		zap.Int("group_id", group.ID),
		zap.String("name", group.Name),
	)

	lp, err := longpoll.Init(vk, group.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init longpolling")
	}

	b := &BotAdapter{
		vk:      vk,
		lp:      lp,
		context: ctx,
		logger:  logger,
	}

	return b, nil
}

func (b *BotAdapter) Send(text, chat string) error {
	peerID, err := strconv.ParseInt(chat, 10, 64)
	if err != nil {
		return b.sendDomain(text, chat)
	}

	return b.sendPeerID(text, peerID)
}

func (b *BotAdapter) sendDomain(text, domain string) error {
	b.logger.Info("Sending message to chat",
		zap.String("domain", domain),
	)

	_, err := b.vk.MessagesSend(api.Params{
		"domain":    domain,
		"message":   text,
		"random_id": rand.Intn(100),
	})

	return err
}

func (b *BotAdapter) sendPeerID(text string, peerID int64) error {
	b.logger.Info("Sending message to chat",
		zap.Int64("peer_id", peerID),
	)

	_, err := b.vk.MessagesSend(api.Params{
		"peer_id":   peerID,
		"message":   text,
		"random_id": rand.Intn(100),
	})

	return err
}

func (b *BotAdapter) RegisterAt(brain *joe.Brain) {
	go func() {
		err := b.lp.Run()
		if err != nil {
			b.logger.Error("longpoll run failed", zap.Error(err))
		}
	}()

	b.lp.MessageNew(func(object object.MessageNewObject, i int) {
		b.logger.Info("Message received",
			zap.Int("peer_id", object.Message.PeerID),
		)
		brain.Emit(b.dispatch(object))
	})
}

func (b *BotAdapter) dispatch(object object.MessageNewObject) interface{} {
	switch object.Message.Action.Type {
	case "chat_create":
		return ChatCreate{
			Channel: strconv.Itoa(object.Message.PeerID),
			Text:    object.Message.Action.Text,
			Data:    object,
		}

	case "chat_title_update":
		return ChatTitleUpdate{
			Channel: strconv.Itoa(object.Message.PeerID),
			NewText: object.Message.Action.Text,
			Data:    object,
		}

	case "chat_photo_update", "chat_photo_remove":
		return ChatPhotoUpdate{
			Channel:  strconv.Itoa(object.Message.PeerID),
			NewPhoto: object.Message.Action.Photo,
			Removed:  object.Message.Action.Type == "chat_photo_remove",
			Data:     object,
		}

	case "chat_pin_message", "chat_unpin_message":
		return ChatPinUpdate{
			Channel:   strconv.Itoa(object.Message.PeerID),
			UserID:    strconv.Itoa(object.Message.Action.MemberID),
			MessageID: strconv.Itoa(object.Message.Action.ConversationMessageID),
			Unpinned:  object.Message.Action.Type == "chat_unpin_message",
			Data:      object,
		}

	case "chat_invite_user", "chat_invite_user_by_link":
		return UserEnteredChat{
			Channel: strconv.Itoa(object.Message.PeerID),
			UserID:  strconv.Itoa(object.Message.Action.MemberID),
			ByLink:  object.Message.Action.Type == "chat_invite_user_by_link",
			Data:    object,
		}

	case "chat_kick_user":
		return UserLeavedChat{
			Channel: strconv.Itoa(object.Message.PeerID),
			UserID:  strconv.Itoa(object.Message.Action.MemberID),
			Data:    object,
		}

	default:
		return joe.ReceiveMessageEvent{
			Text:     object.Message.Text,
			ID:       strconv.Itoa(object.Message.ID),
			AuthorID: strconv.Itoa(object.Message.FromID),
			Channel:  strconv.Itoa(object.Message.PeerID),
			Data:     object,
		}
	}
}

func (b *BotAdapter) Close() error {
	b.lp.Shutdown()
	return nil
}
