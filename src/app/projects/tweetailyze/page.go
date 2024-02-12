package tweetailyze

import (
	"net/http"

	"calebsideras.com/temporary/src/components/server"
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

func Page(w http.ResponseWriter, r *http.Request) templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/projects/_markdown/tweetailyze.md")

	if err != nil {
		panic(err)
	}

	return server.Project(server.ProjectType{
		Title:      "Tweetailyze",
		ProjectURL: "https://github.com/caleb-sideras/Tweetailyze",
		Body:       newTempl,
	})
}
