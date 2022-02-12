package resume

import "errors"

// Resume is a container for a JSON resume blob
type Resume struct {
	slug      string
	body      string
	isVisible bool
}

func NewResume(slug, body string, isVisible bool) (*Resume, error) {
	if slug == "" {
		return nil, errors.New("empty slug")
	}
	if body == "" {
		return nil, errors.New("empty body")
	}

	return &Resume{
		slug:      slug,
		body:      body,
		isVisible: isVisible,
	}, nil
}

func (r Resume) Slug() string {
	return r.slug
}

func (r Resume) Body() string {
	return r.body
}

func (r Resume) IsVisble() bool {
	return r.isVisible
}
