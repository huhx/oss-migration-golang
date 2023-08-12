package model

import "time"

type MarkdownImage struct {
	ImageName    string
	MarkdownName string
	LineNumber   int
	ImageTag     string
}

type ListResponse struct {
	ImageName    string
	ImagePath    string
	CreateTime   time.Time
	ImageSize    int64
	IsUsed       bool
	ImageTag     *string
	LineNumber   *int
	MarkdownName *string
}
type PlanResponse struct {
	ImageName string
	Path      string
	ImageIn   string
}
