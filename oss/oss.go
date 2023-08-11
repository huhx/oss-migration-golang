package oss

import "time"

type MarkdownImage struct {
	ImageName    string
	MarkdownName string
}

type ListResponse struct {
	ImageName    string
	ImagePath    string
	CreateTime   time.Time
	ImageSize    int64
	IsUsed       bool
	MarkdownName *string
}
type PlanResponse struct {
	ImageName string
	Path      string
	ImageIn   string
}
