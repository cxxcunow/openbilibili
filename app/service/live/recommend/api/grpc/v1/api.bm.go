// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api/grpc/v1/api.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	api/grpc/v1/api.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

// ===================
// Recommend Interface
// ===================

type Recommend interface {
	// 获取n个推荐, 得到的结果是在线的房间
	// 去重，不会重复推荐
	// 如果没有足够推荐的结果则返回空的结果，调用方需要补位
	RandomRecsByUser(ctx context.Context, req *GetRandomRecReq) (resp *GetRandomRecResp, err error)

	// 清空推荐缓存，清空推荐过的集合
	ClearRecommendCache(ctx context.Context, req *ClearRecommendRequest) (resp *ClearRecommendResponse, err error)
}

var v1RecommendSvc Recommend

// @params GetRandomRecReq
// @router GET /xlive/recommend/v1/recommend/random_recs_by_user
// @response GetRandomRecResp
func recommendRandomRecsByUser(c *bm.Context) {
	p := new(GetRandomRecReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RecommendSvc.RandomRecsByUser(c, p)
	c.JSON(resp, err)
}

// @params ClearRecommendRequest
// @router GET /xlive/recommend/v1/recommend/clear_recommend_cache
// @response ClearRecommendResponse
func recommendClearRecommendCache(c *bm.Context) {
	p := new(ClearRecommendRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RecommendSvc.ClearRecommendCache(c, p)
	c.JSON(resp, err)
}

// RegisterV1RecommendService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1RecommendService(e *bm.Engine, svc Recommend, midMap map[string]bm.HandlerFunc) {
	v1RecommendSvc = svc
	e.GET("/xlive/recommend/v1/recommend/random_recs_by_user", recommendRandomRecsByUser)
	e.GET("/xlive/recommend/v1/recommend/clear_recommend_cache", recommendClearRecommendCache)
}
