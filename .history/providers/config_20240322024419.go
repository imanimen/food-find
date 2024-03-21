package providers


type IConfig interface {
	Get(key string) string
}

type Con