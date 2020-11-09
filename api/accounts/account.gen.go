// Code generated by "make api"; DO NOT EDIT.
package accounts

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/scopes"
)

type Account struct {
	Id           string                 `json:"id,omitempty"`
	Scope        *scopes.ScopeInfo      `json:"scope,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Description  string                 `json:"description,omitempty"`
	CreatedTime  time.Time              `json:"created_time,omitempty"`
	UpdatedTime  time.Time              `json:"updated_time,omitempty"`
	Version      uint32                 `json:"version,omitempty"`
	Type         string                 `json:"type,omitempty"`
	AuthMethodId string                 `json:"auth_method_id,omitempty"`
	Attributes   map[string]interface{} `json:"attributes,omitempty"`

	response *api.Response
}

func (n Account) ResponseBody() *bytes.Buffer {
	return n.response.Body
}

func (n Account) ResponseMap() map[string]interface{} {
	return n.response.Map
}

func (n Account) ResponseStatus() int {
	return n.response.HttpResponse().StatusCode
}

func (n Account) ResponseTraceId() string {
	return n.response.HttpResponse().Header.Get("TraceId")
}

type AccountReadResult struct {
	Item     *Account
	response *api.Response
}

func (n AccountReadResult) GetItem() interface{} {
	return n.Item
}

func (n AccountReadResult) GetResponseBody() *bytes.Buffer {
	return n.response.Body
}

func (n AccountReadResult) GetResponseMap() map[string]interface{} {
	return n.response.Map
}

type AccountCreateResult = AccountReadResult
type AccountUpdateResult = AccountReadResult

type AccountDeleteResult struct {
	response *api.Response
}

func (n AccountDeleteResult) GetResponseBody() *bytes.Buffer {
	return n.response.Body
}

func (n AccountDeleteResult) GetResponseMap() map[string]interface{} {
	return n.response.Map
}

type AccountListResult struct {
	Items    []*Account
	response *api.Response
}

func (n AccountListResult) GetItems() interface{} {
	return n.Items
}

func (n AccountListResult) GetResponseBody() *bytes.Buffer {
	return n.response.Body
}

func (n AccountListResult) GetResponseMap() map[string]interface{} {
	return n.response.Map
}

// Client is a client for this collection
type Client struct {
	client *api.Client
}

// Creates a new client for this collection. The submitted API client is cloned;
// modifications to it after generating this client will not have effect. If you
// need to make changes to the underlying API client, use ApiClient() to access
// it.
func NewClient(c *api.Client) *Client {
	return &Client{client: c.Clone()}
}

// ApiClient returns the underlying API client
func (c *Client) ApiClient() *api.Client {
	return c.client
}

func (c *Client) Create(ctx context.Context, authMethodId string, opt ...Option) (*AccountCreateResult, error) {
	if authMethodId == "" {
		return nil, fmt.Errorf("empty authMethodId value passed into Create request")
	}

	opts, apiOpts := getOpts(opt...)

	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts.postMap["auth_method_id"] = authMethodId

	req, err := c.client.NewRequest(ctx, "POST", "accounts", opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Create request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Create call: %w", err)
	}

	target := new(AccountCreateResult)
	target.Item = new(Account)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Create response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Read(ctx context.Context, accountId string, opt ...Option) (*AccountReadResult, error) {
	if accountId == "" {
		return nil, fmt.Errorf("empty accountId value passed into Read request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("accounts/%s", accountId), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Read request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(AccountReadResult)
	target.Item = new(Account)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Update(ctx context.Context, accountId string, version uint32, opt ...Option) (*AccountUpdateResult, error) {
	if accountId == "" {
		return nil, fmt.Errorf("empty accountId value passed into Update request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into Update request and automatic versioning not specified")
		}
		existingTarget, existingErr := c.Read(ctx, accountId, opt...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	req, err := c.client.NewRequest(ctx, "PATCH", fmt.Sprintf("accounts/%s", accountId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Update request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Update call: %w", err)
	}

	target := new(AccountUpdateResult)
	target.Item = new(Account)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Update response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}

func (c *Client) Delete(ctx context.Context, accountId string, opt ...Option) (*AccountDeleteResult, error) {
	if accountId == "" {
		return nil, fmt.Errorf("empty accountId value passed into Delete request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("accounts/%s", accountId), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Delete request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Delete call: %w", err)
	}

	apiErr, err := resp.Decode(nil)
	if err != nil {
		return nil, fmt.Errorf("error decoding Delete response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}

	target := &AccountDeleteResult{
		response: resp,
	}
	return target, nil
}

func (c *Client) List(ctx context.Context, authMethodId string, opt ...Option) (*AccountListResult, error) {
	if authMethodId == "" {
		return nil, fmt.Errorf("empty authMethodId value passed into List request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	opts.queryMap["auth_method_id"] = authMethodId

	req, err := c.client.NewRequest(ctx, "GET", "accounts", nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating List request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	target := new(AccountListResult)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.response = resp
	return target, nil
}
