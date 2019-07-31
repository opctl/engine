package pkgs

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./fakeHandler.go --fake-name FakeHandler ./ Handler

import (
	"net/http"

	"github.com/opctl/opctl/sdks/go/node/api/handler/pkgs/ref"
	"github.com/opctl/opctl/sdks/go/node/core"
	"github.com/opctl/opctl/sdks/go/util/urlpath"
)

// Handler deprecated; use data
type Handler interface {
	Handle(
		httpResp http.ResponseWriter,
		httpReq *http.Request,
	)
}

// NewHandler returns an initialized Handler instance
func NewHandler(
	core core.Core,
) Handler {
	return _handler{
		refHandler: ref.NewHandler(core),
	}
}

type _handler struct {
	refHandler ref.Handler
}

func (hdlr _handler) Handle(
	httpResp http.ResponseWriter,
	httpReq *http.Request,
) {
	pathSegment, err := urlpath.NextSegment(httpReq.URL)
	if nil != err {
		http.Error(httpResp, err.Error(), http.StatusBadRequest)
		return
	}

	switch pathSegment {
	case "":
		http.NotFoundHandler().ServeHTTP(httpResp, httpReq)
		return
	default:
		hdlr.refHandler.Handle(
			pathSegment,
			httpResp,
			httpReq,
		)
	}
}
