package sanity

type Project struct {
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Description  string `json:"description"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type Response[T any] struct {
	Result   []T      `json:"result"`
	Ms       int      `json:"ms"`
	SyncTags []string `json:"syncTags"`
}
