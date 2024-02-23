package peoplepedia

import (
	"calebsideras.com/temporary/src/components/server"
	"calebsideras.com/temporary/src/utils"
	"github.com/a-h/templ"
)

func Page_() templ.Component {

	newTempl, err := utils.MdFileToTempl("src/app/projects/_markdown/peoplepedia.md")

	if err != nil {
		panic(err)
	}

	return server.Project(server.ProjectType{
		Title:      "PeoplePedia",
		ProjectURL: "https://github.com/caleb-sideras/PeoplePedia-Frontend",
		Body:       newTempl,
	})
}
