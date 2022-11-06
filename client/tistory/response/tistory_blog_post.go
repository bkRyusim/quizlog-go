package response

type Post struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	PostUrl    string `json:"postUrl"`
	Visibility string `json:"visibility"`
	CategoryId string `json:"categoryId"`
	Comments   string `json:"comments"`
	Trackbacks string `json:"trackbacks"`
	Date       string `json:"date"`
}

type TistoryPostList struct {
	Tistory struct {
		Status string `json:"status"`
		Item   struct {
			Url          string `json:"url"`
			SecondaryUrl string `json:"secondaryUrl"`
			Page         string `json:"page"`
			Count        string `json:"count"`
			TotalCount   string `json:"totalCount"`
			Posts        []Post `json:"posts"`
		} `json:"item"`
	} `json:"tistory"`
}
