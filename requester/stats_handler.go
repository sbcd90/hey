package requester

import (
	"context"
	"google.golang.org/grpc/stats"
	"time"
)

func NewClientStatusHandler() stats.Handler {
	return &ClientStatsHandlerImpl{}
}

type ClientStatsHandlerImpl struct {
	updateTxnName bool
	ConnStart time.Time
	ConnEnd time.Time
	ResStart time.Duration
}

func (h *ClientStatsHandlerImpl) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	return ctx
}

func (h *ClientStatsHandlerImpl) HandleRPC(ctx context.Context, s stats.RPCStats)  {
	switch v := s.(type) {
	case *stats.Begin:
		h.ConnStart = v.BeginTime
	case *stats.End:
		h.ConnEnd = v.EndTime
		h.ResStart = now()
	}
}

func (h *ClientStatsHandlerImpl) TagConn(ctx context.Context, s *stats.ConnTagInfo) context.Context {
	return ctx
}

func (h *ClientStatsHandlerImpl) HandleConn(ctx context.Context, s stats.ConnStats)  {

}