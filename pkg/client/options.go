/*
Copyright © 2021 Cisco and/or its affiliates. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/banzaicloud/go-cruise-control/pkg/types"
)

const (
	HTTPHeaderUserAgent   = "User-Agent"
	HTTPHeaderAccept      = "Accept"
	HTTPHeaderContentType = "Content-Type"
	MIMETypeJSON          = "application/json"
	ChartSetUTF8          = "utf-8"
)

type RequestOptions interface {
	apply(*http.Request) error
}

type requestOptionApplier func(*http.Request) error

func (c requestOptionApplier) apply(r *http.Request) error {
	return c(r)
}

func WithAuthInfo(a AuthInfo) RequestOptions {
	return requestOptionApplier(func(r *http.Request) error {
		if a != nil {
			return a.Apply(r)
		}
		return nil
	})
}

func WithServerURL(u *url.URL) RequestOptions {
	return requestOptionApplier(func(r *http.Request) error {
		if r.URL == nil {
			r.URL = &url.URL{}
		}
		baseURL := *u
		reqURL := baseURL.ResolveReference(r.URL)
		r.URL = reqURL
		return nil
	})
}

func WithEndpoint(endpoint types.APIEndpoint) RequestOptions {
	return requestOptionApplier(func(r *http.Request) error {
		if r.URL == nil {
			r.URL = &url.URL{}
		}
		r.URL.Path = path.Join(r.URL.Path, endpoint.Path())
		return nil
	})
}

func WithMethod(m string) RequestOptions {
	return requestOptionApplier(func(r *http.Request) error {
		r.Method = m
		return nil
	})
}

func WithHeader(h string, v string) RequestOptions {
	return requestOptionApplier(func(r *http.Request) error {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set(h, v)
		return nil
	})
}

func WithUserAgent(agent string) RequestOptions {
	return WithHeader(HTTPHeaderUserAgent, agent)
}

func WithAcceptJSON() RequestOptions {
	return WithHeader(HTTPHeaderAccept, MIMETypeJSON)
}

func WithContentTypeJSON() RequestOptions {
	return WithHeader(HTTPHeaderContentType, fmt.Sprintf("%s; charset=%s", MIMETypeJSON, ChartSetUTF8))
}

func WithJSONQuery() RequestOptions {
	return requestOptionApplier(func(r *http.Request) error {
		if r.URL == nil {
			r.URL = &url.URL{}
		}
		q, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			return err
		}
		q.Set("json", "true")
		r.URL.RawQuery = q.Encode()
		return nil
	})
}

func WithContext(ctx context.Context) RequestOptions {
	return requestOptionApplier(func(r *http.Request) error {
		r2 := r.WithContext(ctx)
		*r = *r2
		return nil
	})
}
