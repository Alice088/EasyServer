package builder

type ServerDirector struct {
	Builder IServerBuilder
}

func (sd *ServerDirector) Construct() {
	sd.Builder.MakeRouter()
	sd.Builder.MakeAPI()
	sd.Builder.MakeServer()
}
