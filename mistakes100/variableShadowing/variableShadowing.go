// @Description variableShadowing
// @Author caopengfei 2022/2/7 14:53

package variableShadowing

import (
	"log"
	"net/http"
)

const tracing = true

func case1() {
	// 在新的代码块钟使用 := 方式赋值，会创建一个影子变量，原值不会修改
	var client *http.Client
	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			return
		}
		log.Println(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return
		}
		log.Println(client)
	}

}

//两种方式解决

func case2() {
	//case 2
	var client *http.Client
	var err error

	if tracing {
		client, err = createClientWithTracing()
		if err != nil {
			return
		}
		log.Println(client)
	} else {
		client, err = createDefaultClient()
		if err != nil {
			return
		}
		log.Println(client)
	}
}

func case3() {
	//case 2
	var client *http.Client

	if tracing {
		c, err := createClientWithTracing()
		if err != nil {
			return
		}
		client = c
		log.Println(client)
	} else {
		// do something
	}
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}
