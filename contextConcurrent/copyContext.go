package contextConcurrent

import "context"

type mapReq struct {
	TotalNum         int         `json:"totalNum"`
	StuStatisticInfo map[int]int `json:"stuStatisticInfo"`
}

func CopyContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}

	fields := []string{"x_rpcid", "x_trace_id", "X-Businessline-Id", "HEADER_BRAND", "IS_BENCHMARK", "logid", "IS_PLAYBACK", "hostname"}
	newCtx := context.Background()

	for _, field := range fields {
		if v := ctx.Value(field); v != nil {
			newCtx = context.WithValue(newCtx, field, v)
		}
	}

	return newCtx
}
