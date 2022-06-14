package appreq

type ArticleRequest struct {
	ArticleId string `json:"article_id"`
}

type NewArticleRequest struct {
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Author string `json:"author"`
}
