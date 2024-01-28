package alert

import "time"

type Embed struct {
	// Title - title of embed
	Title string `json:"title,omitempty"`
	// Type - type of embed (always "rich" for webhook embeds)
	Type EmbedType `json:"type,omitempty"`
	// Description - description of embed
	Description string `json:"description,omitempty"`
	// URL - url of embed
	URL string `json:"url,omitempty"`
	// Timestamp - timestamp of embed content
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Color - color code of the embed
	Color int `json:"color,omitempty"`
	// Footer - footer information
	Footer *EmbedFooter `json:"footer,omitempty"`
	// Image - image information
	Image *EmbedImage `json:"image,omitempty"`
	// Thumbnail - thumbnail information
	Thumbnail *EmbedThumbnail `json:"thumbnail,omitempty"`
	// Video - video information
	Video *EmbedVideo `json:"video,omitempty"`
	// Provider - provider information
	Provider *EmbedProvider `json:"provider,omitempty"`
	// Author - author information
	Author *EmbedAuthor `json:"author,omitempty"`
	// Fields - fields information
	Fields []*EmbedField `json:"fields,omitempty"`
}
type EmbedAuthor struct {
	// Name - name of author
	Name string `json:"name,omitempty"`
	// URL - url of author
	URL string `json:"url,omitempty"`
	// IconURL - url of author icon (only supports http(s) and attachments)
	IconURL string `json:"icon_url,omitempty"`
	// ProxyIconURL - a proxied url of author icon
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}
type EmbedField struct {
	// Name - name of the field (required)
	Name string `json:"name"`
	// Value - value of the field (required)
	Value string `json:"value"`
	// Inline - whether or not this field should display inline
	Inline bool `json:"inline,omitempty"`
}
type EmbedFooter struct {
	// Text - footer text (required)
	Text string `json:"text"`
	// IconURL - url of footer icon (only supports http(s) and attachments)
	IconURL string `json:"icon_url,omitempty"`
	// ProxyIconURL - a proxied url of footer icon
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	// URl source url of image (only supports http(s) and attachments)
	URL string `json:"url,omitempty"`
	// ProxyURL - a proxied url of the image
	ProxyURL string `json:"proxy_url,omitempty"`
	// Height - height of image
	Height int `json:"height,omitempty"`
	// Width - width of image
	Width int `json:"width,omitempty"`
}
type EmbedProvider struct {
	// Name - name of provider
	Name string `json:"name,omitempty"`
	// URL - url of provider
	URL string `json:"url,omitempty"`
}

type EmbedThumbnail struct {
	// URL - source url of thumbnail (only supports http(s) and attachments)
	URL string `json:"url,omitempty"`
	// ProxyURL - a proxied url of the thumbnail
	ProxyURL string `json:"proxy_url,omitempty"`
	// Height - height of thumbnail
	Height int `json:"height,omitempty"`
	// Wifth - width of thumbnail
	Width int `json:"width,omitempty"`
}
type EmbedType string

const (
	// EmbedTypeRich - generic embed rendered from embed attributes
	EmbedTypeRich EmbedType = "rich"
	// EmbedTypeImage - image embed
	EmbedTypeImage EmbedType = "image"
	// EmbedTypeVideo - video embed
	EmbedTypeVideo EmbedType = "video"
	// EmbedTypeGifv - animated gif image embed rendered as a video embed
	EmbedTypeGifv EmbedType = "gifv"
	// EmbedTypeArticle - article embed
	EmbedTypeArticle EmbedType = "article"
	// EmbedTypeLink - link embed
	EmbedTypeLink EmbedType = "link"
)

type EmbedVideo struct {
	// URL - source url of video
	URL string `json:"url,omitempty"`
	// Height - height of video
	Height int `json:"height,omitempty"`
	// Width - width of video
	Width int `json:"width,omitempty"`
}
