
// Code generated by gox; DO NOT EDIT.
package gox
import (
	"github.com/caleb-sideras/gox2/src/app/projects/caas/cluster"
	"github.com/caleb-sideras/gox2/src/app/projects"
	"github.com/caleb-sideras/gox2/src/app/projects/peoplepedia"
	"github.com/caleb-sideras/gox2/src/app/projects/goxFramework"
	"github.com/caleb-sideras/gox2/src/app/projects/musicgpt"
	"github.com/caleb-sideras/gox2/src/app/blog"
	"github.com/caleb-sideras/gox2/src/app"
	"github.com/caleb-sideras/gox2/src/app/blog/temporarystandard"
	"github.com/caleb-sideras/gox2/src/app/experience"
	"github.com/caleb-sideras/gox2/src/app/projects/caas"
	"github.com/caleb-sideras/gox2/src/app/projects/caas/processor"
	"github.com/caleb-sideras/gox2/src/app/projects/tweetailyze"
	"github.com/caleb-sideras/gox2/src/app/home_"
	"github.com/caleb-sideras/gox2/src/app/projects/caas/segmentation"
)

var IndexList = map[string]IndexDefaultFunc{
	"/projects/goxFramework" : app.Index,
	"/blog" : app.Index,
	"/projects/caas/processor" : app.Index,
	"/projects/caas/cluster" : app.Index,
	"/experience" : app.Index,
	"/projects" : app.Index,
	"/projects/tweetailyze" : app.Index,
	"/projects/caas" : app.Index,
	"/projects/musicgpt" : app.Index,
	"/projects/peoplepedia" : app.Index,
	"/blog/temporarystandard" : app.Index,
	"/" : app.Index,
	"/projects/caas/segmentation" : app.Index,
}

var PageRenderList = []RenderDefault{
	{"/experience", experience.Page_},
	{"/blog", blog.Page_},
}

var RouteRenderList = []RenderDefault{
	{"/example", home.Example_},
}

var PageHandleList = []Handler{
	{"/blog/temporarystandard", temporarystandard.Page, ResReqHandler},
	{"/projects/caas/cluster", caas_cluster.Page, ResReqHandler},
	{"/projects", projects.Page, ResReqHandler},
	{"/projects/caas", caas.Page, ResReqHandler},
	{"/projects/goxFramework", gox.Page, ResReqHandler},
	{"/projects/musicgpt", musicgpt.Page, ResReqHandler},
	{"/projects/peoplepedia", peoplepedia.Page, ResReqHandler},
	{"/projects/tweetailyze", tweetailyze.Page, ResReqHandler},
	{"/", home.Page, ResReqHandler},
	{"/projects/caas/processor", caas_processor.Page, ResReqHandler},
	{"/projects/caas/segmentation", caas_segmentation.Page, ResReqHandler},
}

var RouteHandleList = []Handler{
	{"/projects/caas/readme", caas.Readme, DefaultHandler},
	{"/projects/caas/videos", caas.Videos, DefaultHandler},
	{"/projects/goxFramework/readme", gox.Readme, DefaultHandler},
	{"/projects/goxFramework/videos", gox.Videos, DefaultHandler},
	{"/projects/musicgpt/readme", musicgpt.Readme, DefaultHandler},
	{"/projects/musicgpt/videos", musicgpt.Videos, DefaultHandler},
}
