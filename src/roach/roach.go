package roach

// lch listen channel to recieve request
// rch response channel to return result
type RoachServer struct {
	lch chan string
	rch chan string
}

type Request struct {
	Cmd    string
	Params map[string]string
}

type Response struct {
}

func (r *Response) Write(v interface{}) {

}

func NewRoachServer() *RoachServer {
}

func (r *RoachServer) Listen() {
	for {
		service := <-r.lch
		r.serveRoach(service)
	}
}

func (r *RoachServer) serveRoach(service string) {

}

func (r *RoachServer) serve() {

}
