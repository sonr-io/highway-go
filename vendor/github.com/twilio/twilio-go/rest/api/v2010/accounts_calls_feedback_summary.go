/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// Optional parameters for the method 'CreateCallFeedbackSummary'
type CreateCallFeedbackSummaryParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// Only include feedback given on or before this date. Format is `YYYY-MM-DD` and specified in UTC.
	EndDate *string `json:"EndDate,omitempty"`
	// Whether to also include Feedback resources from all subaccounts. `true` includes feedback from all subaccounts and `false`, the default, includes feedback from only the specified account.
	IncludeSubaccounts *bool `json:"IncludeSubaccounts,omitempty"`
	// Only include feedback given on or after this date. Format is `YYYY-MM-DD` and specified in UTC.
	StartDate *string `json:"StartDate,omitempty"`
	// The URL that we will request when the feedback summary is complete.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// The HTTP method (`GET` or `POST`) we use to make the request to the `StatusCallback` URL.
	StatusCallbackMethod *string `json:"StatusCallbackMethod,omitempty"`
}

func (params *CreateCallFeedbackSummaryParams) SetPathAccountSid(PathAccountSid string) *CreateCallFeedbackSummaryParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateCallFeedbackSummaryParams) SetEndDate(EndDate string) *CreateCallFeedbackSummaryParams {
	params.EndDate = &EndDate
	return params
}
func (params *CreateCallFeedbackSummaryParams) SetIncludeSubaccounts(IncludeSubaccounts bool) *CreateCallFeedbackSummaryParams {
	params.IncludeSubaccounts = &IncludeSubaccounts
	return params
}
func (params *CreateCallFeedbackSummaryParams) SetStartDate(StartDate string) *CreateCallFeedbackSummaryParams {
	params.StartDate = &StartDate
	return params
}
func (params *CreateCallFeedbackSummaryParams) SetStatusCallback(StatusCallback string) *CreateCallFeedbackSummaryParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateCallFeedbackSummaryParams) SetStatusCallbackMethod(StatusCallbackMethod string) *CreateCallFeedbackSummaryParams {
	params.StatusCallbackMethod = &StatusCallbackMethod
	return params
}

// Create a FeedbackSummary resource for a call
func (c *ApiService) CreateCallFeedbackSummary(params *CreateCallFeedbackSummaryParams) (*ApiV2010CallFeedbackSummary, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint(*params.EndDate))
	}
	if params != nil && params.IncludeSubaccounts != nil {
		data.Set("IncludeSubaccounts", fmt.Sprint(*params.IncludeSubaccounts))
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint(*params.StartDate))
	}
	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.StatusCallbackMethod != nil {
		data.Set("StatusCallbackMethod", *params.StatusCallbackMethod)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010CallFeedbackSummary{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteCallFeedbackSummary'
type DeleteCallFeedbackSummaryParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteCallFeedbackSummaryParams) SetPathAccountSid(PathAccountSid string) *DeleteCallFeedbackSummaryParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete a FeedbackSummary resource from a call
func (c *ApiService) DeleteCallFeedbackSummary(Sid string, params *DeleteCallFeedbackSummaryParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchCallFeedbackSummary'
type FetchCallFeedbackSummaryParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchCallFeedbackSummaryParams) SetPathAccountSid(PathAccountSid string) *FetchCallFeedbackSummaryParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a FeedbackSummary resource from a call
func (c *ApiService) FetchCallFeedbackSummary(Sid string, params *FetchCallFeedbackSummaryParams) (*ApiV2010CallFeedbackSummary, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Calls/FeedbackSummary/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010CallFeedbackSummary{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
