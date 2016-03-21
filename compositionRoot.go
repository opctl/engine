package main

import (
  "github.com/dev-op-spec/engine/core"
  "github.com/dev-op-spec/engine/rest"
)

type compositionRoot interface {
  RestApi() rest.Api
}

func newCompositionRoot(
) (compositionRoot compositionRoot, err error) {

  coreApi, err := core.New()
  if (nil != err) {
    return
  }

  compositionRoot = &_compositionRoot{
    restApi:rest.New(coreApi),
  }

  return

}

type _compositionRoot struct {
  restApi rest.Api
}

func (_compositionRoot _compositionRoot) RestApi() rest.Api {
  return _compositionRoot.restApi
}