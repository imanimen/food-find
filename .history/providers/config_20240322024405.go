package providers


func IConfig interface {
	Get(key string) string
}