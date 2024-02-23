package caas

import (
	"calebsideras.com/temporary/src/components/server"
	"github.com/a-h/templ"
)

func Page_() templ.Component {

	return server.ProjectTab(server.ProjectTabType{
		Title:       "CAAS",
		ProjectURL:  "https://github.com/caleb-sideras/caas",
		ReadMeURL:   "/projects/caas/readme",
		VideoURL:    "/projects/caas/videos",
		InitialBody: Readme_(),
	})
}
