package temporaryframework

import (
	"calebsideras.com/temporary/src/components/server"
	"github.com/a-h/templ"
)

func Page_() templ.Component {

	return server.ProjectTab(server.ProjectTabType{
		Title:       "Temporary Framework",
		ProjectURL:  "http://temporary-framework.org/",
		ReadMeURL:   "/projects/temporary-framework/readme",
		VideoURL:    "/projects/temporary-framework/videos",
		InitialBody: Readme_(),
	})
}
