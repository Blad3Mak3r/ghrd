package structs

//GithubAsset
type GithubAsset struct {
	Url                string `json:"url"`
	BrowserDownloadUrl string `json:"browser_download_url"`

	Id            int    `json:"id"`
	NodeId        string `json:"node_id"`
	Name          string `json:"name"`
	Label         string `json:"label"`
	State         string `json:"state"`
	ContenType    string `json:"content_type"`
	Size          int    `json:"size"`
	DownloadCount int    `json:"download_Count"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

//GithubRelease
type GithubRelease struct {
	Url           string `json:"url"`
	HtmlUrl       string `json:"html_url"`
	AssetsUrl     string `json:"assets_url"`
	UploadUrl     string `json:"upload_url"`
	TarballUrl    string `json:"tarball_url"`
	ZipballUrl    string `json:"zipball_url"`
	DiscussionUrl string `json:"discussion_url"`

	Id              int    `json:"id"`
	NodeId          string `json:"node_id"`
	TagName         string `json:"tag_name"`
	TargetCommitish string `json:"target_commitish"`
	Name            string `json:"name"`
	Body            string `json:"body"`
	Draft           bool   `json:"draft"`
	PreRelease      bool   `json:"prerelease"`

	CreateAt    string `json:"created_at"`
	PublishedAt string `json:"published_at"`

	Assets []GithubAsset `json:"assets"`
}
