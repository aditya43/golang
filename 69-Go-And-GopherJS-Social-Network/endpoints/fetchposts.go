package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/common/authenticate"

	"github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/common"
)

func FetchPostsEndpoint(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		gfSession, err := authenticate.SessionStore.Get(r, "gopherface-session")
		if err != nil {
			log.Print(err)
			return
		}
		uuid := gfSession.Values["uuid"].(string)

		posts, err := env.DB.FetchPosts(uuid)

		if err != nil {
			log.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(posts)

	})
}
