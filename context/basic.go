package contexti

type empty struct{}
type Context = RequestContext[empty, empty]
