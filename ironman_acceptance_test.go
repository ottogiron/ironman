package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/ironman-project/ironman/acceptance"
)

func FeatureContext(s *godog.Suite) {
	acceptance.InstallContext(s)
	acceptance.LinkContext(s)
}
