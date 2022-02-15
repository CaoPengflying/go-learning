package args

import (
	"flag"
	"testing"
)

var offline = flag.Bool("offline", false, "only perform local tests")
var skip = flag.Bool("skip", false, "only perform local tests")
var env = flag.String("env", "dev", "env")

func TestEnv(t *testing.T) {
	if *env == "dev" {
		t.Log("dev")
	}else {
		t.Skip("dev skip")
	}
}

func TestArgs(t *testing.T) {
	if !*offline {
		t.Fatal("offline if false")
	}
}

func TestSkip(t *testing.T) {
	if !*skip {

		t.Fatal("skip is false")

	}else {
		t.Skip("test skip")

	}
}


