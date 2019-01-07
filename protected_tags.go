package gitlab

import (
	"fmt"
	"net/url"
)

// ProtectedTagsService handles communication with the protected tag methods of the GitLab API.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/protected_tags.html
type ProtectedTagsService struct {
	client *Client
}

// TagAccessDescription reperesents the access decription for a protected tag.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/protected_tags.html#list-protected-tags
type TagAccessDescription struct {
	AccessLevel            AccessLevelValue `json:"access_level"`
	AccessLevelDescription string           `json:"access_level_description"`
}

// ProtectedTag represents a protected tag.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/protected_tags.html#list-protected-tags
type ProtectedTag struct {
	Name               string                  `json:"name"`
	CreateAccessLevels []*TagAccessDescription `json:"create_access_levels"`
}

// ListProtectedTagsOptions represents available ListProtectedTags() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/protected_tags.html#list-protected-tags
type ListProtectedTagsOptions ListOptions

// ListProtectedTags returns a list of protected tags from a project.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/protected_tags.html#list-protected-tags
func (s *ProtectedTagsService) ListProtectedTags(pid interface{}, opt *ListProtectedTagsOptions, options ...OptionFunc) ([]*ProtectedTag, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/protected_tags", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var p []*ProtectedTag
	resp, err := s.client.Do(req, &p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// GetProtectedTag returns a single protected tag or wildcard protected tag.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/protected_tags.html#get-a-single-protected-tag-or-wildcard-protected-tag
func (s *ProtectedTagsService) GetProtectedTag(pid interface{}, tag string, options ...OptionFunc) (*ProtectedTag, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/protected_tags/%s", url.QueryEscape(project), tag)

	req, err := s.client.NewRequest("GET", u, nil, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(ProtectedTag)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// ProtectRepositoryTagsOptions represents available ProtectRepositoryTags() options.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/protected_tags.html#protect-repository-tags
type ProtectRepositoryTagsOptions struct {
	Name              string            `url:"name,omitempty" json:"name,omitempty"`
	CreateAccessLevel *AccessLevelValue `url:"create_access_level,omitempty" json:"create_access_level,omitempty"`
}

// ProtectRepositoryTags protects a single repository tag or several project repository tags using a wildcard protected tag.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/protected_tags.html#protect-repository-tags
func (s *ProtectedTagsService) ProtectRepositoryTags(pid interface{}, opt *ProtectRepositoryTagsOptions, options ...OptionFunc) (*ProtectedTag, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/protected_tags", url.QueryEscape(project))

	req, err := s.client.NewRequest("POST", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	p := new(ProtectedTag)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// UnprotectRepositoryTags unprotects the given protected tag or wildcard protected tag.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/protected_tags.html#unprotect-repository-tags
func (s *ProtectedTagsService) UnprotectRepositoryTags(pid interface{}, tag string, options ...OptionFunc) (*Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("projects/%s/protected_tags/%s", url.QueryEscape(project), tag)

	req, err := s.client.NewRequest("DELETE", u, nil, options)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
