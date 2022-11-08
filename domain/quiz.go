package domain

type Quiz struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	PostUrl  string `json:"post_url"`
	CreateAt string `json:"create_at"`
}
