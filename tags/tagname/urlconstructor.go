package tagname

type URLConstructor interface {
	ConstructURL(host string, tag string) string
}

type URLConstructorFunc func(host string, tag string) string

func (fn URLConstructorFunc) ConstructURL(host string, tag string) string {
	return fn(host, tag)
}

var DefaultURLConstructor URLConstructor = URLConstructorFunc(rssURL)
