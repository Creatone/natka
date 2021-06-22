package instagram

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/revel/revel"
)

type MediaList struct {
	Data []struct {
		Id string `json:"id"`
	} `json:"data"`
}

type Media struct {
	MediaURL  string `json:"media_url"`
	Permalink string `json:"permalink"`
}

var (
	userID      string
	accessToken string
)

const (
	url = "https://graph.instagram.com"
)

func GetPosts() ([]Media, error) {
	var medias []Media
	var isPresent bool
	userID, isPresent = revel.Config.String("insta.user.id")
	if !isPresent {
		// TODO: Log
		return nil, fmt.Errorf("couldn't obtain instagram posts")
	}

	accessToken, isPresent = revel.Config.String("insta.access.token")
	if !isPresent {
		// TODO: Log
		return nil, fmt.Errorf("couldn't obtain instagram posts")
	}

	res, err := http.Get(fmt.Sprintf("%s/%s/media?access_token=%s", url, userID, accessToken))
	if err != nil {
		return nil, err
	}

	mediaList := MediaList{}
	err = json.NewDecoder(res.Body).Decode(&mediaList)
	if err != nil {
		return nil, err
	}

	for _, m := range mediaList.Data {
		res, err := http.Get(fmt.Sprintf("%s/%s/?fields=media_url,permalink&access_token=%s", url, m.Id, accessToken))
		if err != nil {
			return nil, err
		}

		media := Media{}
		err = json.NewDecoder(res.Body).Decode(&media)
		if err != nil {
			return nil, err
		}

		medias = append(medias, media)
	}

	return medias, nil
}
