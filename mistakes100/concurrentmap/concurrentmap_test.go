// @Description concurrentmap
// @Author caopengfei 2022/4/15 10:34

package concurrentmap

import (
	"context"
	"sync"
	"testing"
)

const ResMetaDataKey = "__res_metadata"

func TestFun(t *testing.T) {
	ctx := context.Background()
	metaInfo := make(map[string]string)
	ctx = context.WithValue(ctx, ResMetaDataKey, metaInfo)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ctx2 context.Context) {
			defer wg.Done()

			meta := ctx.Value(ResMetaDataKey)
			callMetaData := make(map[string]string)
			callMetaData["time"] = "10:40"
			callMetaData["date"] = "2022-04-15"
			if meta != nil && len(callMetaData) > 0 {
				resMeta := meta.(map[string]string)
				for k, v := range callMetaData {
					resMeta[k] = v
				}
			}
		}(ctx)
	}

	wg.Wait()

}
