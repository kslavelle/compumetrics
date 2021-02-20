package main

import (
	"github.com/kslavelle/compumetrics/pkg/identity"
)

func main() {
	identityRouter := identity.CreateIdentityProvider()
	identityRouter.Run("0.0.0.0:8888")
}
