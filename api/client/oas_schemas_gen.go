// Code generated by ogen, DO NOT EDIT.

package client

import (
	ht "github.com/ogen-go/ogen/http"
)

// DeleteNoteNoContent is response for DeleteNote operation.
type DeleteNoteNoContent struct{}

// Ref: #/components/schemas/draft
type DraftMultipart struct {
	// Note title.
	Title string `json:"title"`
	// Note body.
	Text string `json:"text"`
	// Optional (and not verifiable) author of note.
	Author OptString `json:"author"`
	// Do not make list of attachments.
	HideAttachments OptBool `json:"hide_attachments"`
	// File attachment.
	Attachment []ht.MultipartFile `json:"attachment"`
}

// GetTitle returns the value of Title.
func (s *DraftMultipart) GetTitle() string {
	return s.Title
}

// GetText returns the value of Text.
func (s *DraftMultipart) GetText() string {
	return s.Text
}

// GetAuthor returns the value of Author.
func (s *DraftMultipart) GetAuthor() OptString {
	return s.Author
}

// GetHideAttachments returns the value of HideAttachments.
func (s *DraftMultipart) GetHideAttachments() OptBool {
	return s.HideAttachments
}

// GetAttachment returns the value of Attachment.
func (s *DraftMultipart) GetAttachment() []ht.MultipartFile {
	return s.Attachment
}

// SetTitle sets the value of Title.
func (s *DraftMultipart) SetTitle(val string) {
	s.Title = val
}

// SetText sets the value of Text.
func (s *DraftMultipart) SetText(val string) {
	s.Text = val
}

// SetAuthor sets the value of Author.
func (s *DraftMultipart) SetAuthor(val OptString) {
	s.Author = val
}

// SetHideAttachments sets the value of HideAttachments.
func (s *DraftMultipart) SetHideAttachments(val OptBool) {
	s.HideAttachments = val
}

// SetAttachment sets the value of Attachment.
func (s *DraftMultipart) SetAttachment(val []ht.MultipartFile) {
	s.Attachment = val
}

type HeaderAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *HeaderAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *HeaderAuth) SetAPIKey(val string) {
	s.APIKey = val
}

type ID string

// Ref: #/components/schemas/note
type Note struct {
	ID ID `json:"id"`
	// Public URL.
	PublicURL string `json:"public_url"`
}

// GetID returns the value of ID.
func (s *Note) GetID() ID {
	return s.ID
}

// GetPublicURL returns the value of PublicURL.
func (s *Note) GetPublicURL() string {
	return s.PublicURL
}

// SetID sets the value of ID.
func (s *Note) SetID(val ID) {
	s.ID = val
}

// SetPublicURL sets the value of PublicURL.
func (s *Note) SetPublicURL(val string) {
	s.PublicURL = val
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type QueryAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *QueryAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *QueryAuth) SetAPIKey(val string) {
	s.APIKey = val
}

// UpdateNoteNoContent is response for UpdateNote operation.
type UpdateNoteNoContent struct{}
