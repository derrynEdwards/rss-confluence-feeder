package helpers

type Feed struct {
	Rule struct {
		Name         string `json:"name"`
		MatchesToday string `json:"matchesToday"`
		MatchesTotal string `json:"matchesTotal"`
	} `json:"rule"`
	Items []struct {
		CrawlTimeMsec string   `json:"crawlTimeMsec"`
		TimestampUsec string   `json:"timestampUsec"`
		ID            string   `json:"id"`
		Categories    []string `json:"categories"`
		Title         string   `json:"title"`
		Published     int      `json:"published"`
		Updated       int      `json:"updated"`
		Canonical     []struct {
			Href string `json:"href"`
		} `json:"canonical"`
		Alternate []struct {
			Href string `json:"href"`
			Type string `json:"type"`
		} `json:"alternate"`
		Summary struct {
			Direction string `json:"direction"`
			Content   string `json:"content"`
		} `json:"summary"`
		Author      string        `json:"author"`
		LikingUsers []interface{} `json:"likingUsers"`
		Comments    []interface{} `json:"comments"`
		CommentsNum int           `json:"commentsNum"`
		Annotations []interface{} `json:"annotations"`
		Origin      struct {
			StreamID string `json:"streamId"`
			Title    string `json:"title"`
			HTMLURL  string `json:"htmlUrl"`
		} `json:"origin"`
	} `json:"items"`
}

type BlogResponse struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	SpaceID   string `json:"spaceId"`
	AuthorID  string `json:"authorId"`
	CreatedAt string `json:"createdAt"`
	Version   struct {
		CreatedAt string `json:"createdAt"`
		Message   string `json:"message"`
		Number    int    `json:"number"`
		MinorEdit bool   `json:"minorEdit"`
		AuthorID  string `json:"authorId"`
	} `json:"version"`
	Body struct {
		Storage struct {
			Representation string `json:"representation"`
			Value          string `json:"value"`
		} `json:"storage"`
		AtlasDocFormat struct {
			Representation string `json:"representation"`
			Value          string `json:"value"`
		} `json:"atlas_doc_format"`
		View struct {
			Representation string `json:"representation"`
			Value          string `json:"value"`
		} `json:"view"`
	} `json:"body"`
	Links struct {
		Webui  string `json:"webui"`
		Editui string `json:"editui"`
		Tinyui string `json:"tinyui"`
	} `json:"_links"`
}

type BlogRequest struct {
	SpaceID string   `json:"spaceId"`
	Status  string   `json:"status"`
	Title   string   `json:"title"`
	Body    BlogBody `json:"body"`
}

type BlogBody struct {
	Representation string `json:"representation"`
	Value          string `json:"value"`
}
