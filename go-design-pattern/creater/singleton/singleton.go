package singleton

type Singleton struct {
}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetSingleton() *Singleton {
	return singleton
}
