package response

type TistoryBlogInfo struct {
	Tistory struct {
		Status string `json:"status"`
		Item   struct {
			Id     string `json:"id"`
			UserId string `json:"userId"`
			Blogs  []struct {
				Name                     string `json:"name"`
				Url                      string `json:"url"`
				SecondaryUrl             string `json:"secondaryUrl"`
				Nickname                 string `json:"nickname"`
				Title                    string `json:"title"`
				Description              string `json:"description"`
				Default                  string `json:"default"`
				BlogIconUrl              string `json:"blogIconUrl"`
				FaviconUrl               string `json:"faviconUrl"`
				ProfileThumbnailImageUrl string `json:"profileThumbnailImageUrl"`
				ProfileImageUrl          string `json:"profileImageUrl"`
				Role                     string `json:"role"`
				BlogId                   string `json:"blogId"`
				IsEmpty                  string `json:"isEmpty"`
				Statistics               struct {
					Post      int    `json:"post,string"`
					Comment   string `json:"comment"`
					Trackback string `json:"trackback"`
					Guestbook string `json:"guestbook"`
				} `json:"statistics"`
			} `json:"blogs"`
		} `json:"item"`
	} `json:"tistory"`
}
