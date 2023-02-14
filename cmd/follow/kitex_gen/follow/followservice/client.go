// Code generated by Kitex v0.4.4. DO NOT EDIT.

package followservice

import (
	follow "DY_BAT/cmd/follow/kitex_gen/follow"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	FollowAction(ctx context.Context, req *follow.DouyinFollowActionRequest, callOptions ...callopt.Option) (r *follow.DouyinFollowActionResponse, err error)
	FollowList(ctx context.Context, req *follow.DouyinFollowListRequest, callOptions ...callopt.Option) (r *follow.DouyinFollowListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kFollowServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kFollowServiceClient struct {
	*kClient
}

func (p *kFollowServiceClient) FollowAction(ctx context.Context, req *follow.DouyinFollowActionRequest, callOptions ...callopt.Option) (r *follow.DouyinFollowActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowAction(ctx, req)
}

func (p *kFollowServiceClient) FollowList(ctx context.Context, req *follow.DouyinFollowListRequest, callOptions ...callopt.Option) (r *follow.DouyinFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FollowList(ctx, req)
}
