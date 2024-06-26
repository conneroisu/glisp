package domain

import "net/url"

// NewDocumentURIFromURL converts an URL into a DocumentURI
func NewDocumentURIFromURL(inURL string) (TextDocumentURI, error) {
	uri, err := url.Parse(inURL)
	if err != nil {
		return TextDocumentURI{}, err
	}
	return TextDocumentURI{URL: *uri}, nil
}
