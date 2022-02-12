package resume

import "errors"

// Visit is used to keep track of visits to a resume
type Visit struct {
	uuid      string
	slug      string
	ipAddress string
	userAgent string
}

func NewVisit(uuid, slug, ipAddress, userAgent string) (*Visit, error) {
	if uuid == "" {
		return nil, errors.New("empty visit uuid")
	}
	if slug == "" {
		return nil, errors.New("empty slug")
	}
	if ipAddress == "" {
		return nil, errors.New("empty ip address")
	}
	if userAgent == "" {
		return nil, errors.New("empty user agent")
	}

	return &Visit{
		uuid:      uuid,
		slug:      slug,
		ipAddress: ipAddress,
		userAgent: userAgent,
	}, nil
}

func (v Visit) UUID() string {
	return v.uuid
}

func (v Visit) Slug() string {
	return v.slug
}

func (v Visit) IPAddress() string {
	return v.ipAddress
}

func (v Visit) UserAgent() string {
	return v.userAgent
}
