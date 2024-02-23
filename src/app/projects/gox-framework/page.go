package gox

import (
	"calebsideras.com/temporary/src/components/server"
	"github.com/a-h/templ"
)

func Page_() templ.Component {

	return server.ProjectTab(server.ProjectTabType{
		Title:       "GoX Framework",
		ProjectURL:  "https://goxFramework.org",
		ReadMeURL:   "/projects/gox-framework/readme",
		VideoURL:    "/projects/gox-framework/videos",
		InitialBody: Readme_(),
	})
}
