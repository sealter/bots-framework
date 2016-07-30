package gae_host

import (
	"github.com/strongo/bots-framework/core"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"net/http"
	"github.com/strongo/app"
)

type GaeBotHost struct {
}

var _ bots.BotHost = (*GaeBotHost)(nil)

func (h GaeBotHost) Logger(r *http.Request) strongo.Logger {
	return NewGaeLogger(appengine.NewContext(r))
}

func (h GaeBotHost) GetHttpClient(r *http.Request) *http.Client {
	return &http.Client{Transport: &urlfetch.Transport{Context: appengine.NewContext(r)}}
}

func (h GaeBotHost) GetBotCoreStores(platform string, appContext bots.BotAppContext, r *http.Request) bots.BotCoreStores {
	switch platform {
	case "telegram":
		logger := h.Logger(r)
		appUserStore := NewGaeAppUserStore(logger, r, appContext.AppUserEntityKind(), appContext.AppUserEntityType(), appContext.NewBotAppUserEntity)
		return bots.BotCoreStores{
			BotChatStore: NewGaeTelegramChatStore(logger, r),
			BotUserStore: NewGaeTelegramUserStore(logger, r, appUserStore),
			BotAppUserStore: appUserStore,
		}
	default:
		panic("Unknown platform: " + platform)
	}
}
