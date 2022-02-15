package unit_test

import (
	"net/http"
	"testing"
)

const checkMark = "\u27113"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	url := "http://www.baidu.com"
	statusCode := 200

	t.Log("Given the need to test downloading content")
	{
		t.Logf("\t when checking %s for status code %d ", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("should be able to make the Get call", ballotX, err)
			}
			t.Log("should be able to make the Get call", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("should receive a %d status %v", statusCode, checkMark)
			} else {
				t.Errorf("should receive a %d status %v %v", statusCode, ballotX, resp.StatusCode)
			}

		}
	}
}
