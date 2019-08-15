package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/step_definitions"
)

func InitiateBrowser(s *godog.Suite) {
	s.Step(`^visit dweb$`, step_definitions.OpenDWEB)
	s.Step(`^visit mweb$`, step_definitions.OpenMWEB)
	s.Step(`^visit android$`, step_definitions.OpenAndroid)
	s.Step(`^visit ios$`, step_definitions.OpenIOS)
	s.Step(`^client sends a ([^\"]*) request to "([^\"]*)"$`, step_definitions.RequestAPI)
	s.Step(`^response status should be "([^\"]*)"$`, step_definitions.ResponseAPI)

	GodogMainSupport(s)
}
