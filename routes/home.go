package routes

import (
	"net/http"

	"github.com/martini-contrib/render"

	"../db/documents"
	"../models"
	"../models/data"
)

/* Render index template */
func IndexHandler(rnd render.Render, r *http.Request) {

	/* Init posts data */
	postDocuments := []documents.PostDocument{}
	postsCollection.Find(nil).All(&postDocuments)
	posts := []models.Post{}
	for _, doc := range postDocuments {
		post := models.Post{
			doc.Id,
			doc.Title,
			doc.ContentHtml,
			doc.ContentMarkdown,
			doc.Time,
			doc.Owner,
		}
		posts = append(posts, post)
	}

	/* Init User data */
	userData, _ := getPublicCurrentUserData(r)

	/* Init IndexData */
	data := data.IndexData{posts, userData}

	/* Render html template */
	rnd.HTML(200, "index", data)
}
