package mates

type SearchMateDTO struct {
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Birthday string `json:"birthday"`
	Age      int    `json:"age"`
}

type AddMateDTO struct {
	UserName  string `json:"username"`
	GroupId   int32  `json:"group_id"`
	RealName  string `json:"real_name"`
	Tags      string `json:"tags"`
	Birthday  string `json:"birthday"`
	Hobby     string `json:"hobby"`
	Nickname  string `json:"nickname"`
	Images    string `json:"images"`
	Age       int32  `json:"age"`
	Province  string `json:"province"`
	Sign      string `json:"sign"`
	VideoUrl  string `json:"videourl"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
}
