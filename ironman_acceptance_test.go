package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/ironman-project/ironman/acceptance"
)

func FeatureContext(s *godog.Suite) {
	acceptance.VarsContext(s)
	acceptance.InstallContext(s)
	acceptance.LinkContext(s)
	acceptance.UpdateContext(s)
	acceptance.UnlinkContext(s)
	acceptance.UninstallContext(s)
}
