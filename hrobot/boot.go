package hrobot

import (
	"context"
	"fmt"

	"github.com/floshodan/hrobot-go/hrobot/schema"
	"github.com/google/go-querystring/query"
)

type BootClient struct {
	client *Client
}

type BootLinux struct {
	ServerIP      string
	ServerIpv6Net string
	ServerNumber  int
	Dist          string
	Arch          int
	Lang          string
	Active        bool
	Password      string
	AuthorizedKey []interface{}
	HostKey       []interface{}
}

type BootRescue struct {
	ServerIP      string
	ServerIpv6Net string
	ServerNumber  int
	Os            string
	Arch          int
	Active        bool
	Password      string
	AuthorizedKey []interface{}
	HostKey       []interface{}
}

type Boot struct {
	BootRescue
	BootLinux
}

func (c *BootClient) GetBootList(ctx context.Context, servernumber string) (*BootRescue, *Response, error) {
	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("/boot/%s", servernumber), nil)

	if err != nil {
		return nil, nil, err
	}

	var body schema.BootRescue
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return RescueFromSchema(body), resp, nil
}

func (c *BootClient) ActivateRescue(ctx context.Context, servernumber string, opt *RescueOpts) (*BootRescue, *Response, error) {
	params, _ := query.Values(opt)
	fmt.Println(params)
	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("/boot/%s/rescue", servernumber), params)

	if err != nil {
		return nil, nil, err
	}

	var body schema.BootRescue
	resp, err := c.client.Do(req, &body)
	if err != nil {
		return nil, nil, err
	}

	return RescueFromSchema(body), resp, nil
}

type RescueOpts struct {
	OS             string `url:"os"` //required
	Arch           string `url:"arch"`
	Authorized_Key string `url:"authorized_key"`
}