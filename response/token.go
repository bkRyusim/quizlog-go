package response

type Login struct {
	Token     string `json:"token"`
	UserId    string `json:"user_id"`
	IsNewUser bool   `json:"new_user"`
	Name      string `json:"name"`
}
