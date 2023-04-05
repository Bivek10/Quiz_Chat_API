package socket1

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewWebsocketServer),
)
