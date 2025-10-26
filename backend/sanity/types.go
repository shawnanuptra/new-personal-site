package sanity

type Project struct {
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
	PublishedAt  string `json:"publishedAt"`
	Content      string `json:"markdownContent"`
}

type Response[T any] struct {
	Result   T        `json:"result"`
	Ms       int      `json:"ms"`
	SyncTags []string `json:"syncTags"`
}

type Blog struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	PublishedAt string `json:"publishedAt"`
	Series      string `json:"series,omitempty"`
	Entry       int    `json:"entry,omitempty"`
	Content     string `json:"markdownContent"`
}
