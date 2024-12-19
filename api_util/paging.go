package apiutil

// 分页接口规范：入参必须有 pageno count 出参必须有 pageno count total

import (
	"fmt"
	"reflect"
)

const (
	defaultPageno          = 1
	defaultCount           = 20
	DefaultMaxAllowedCount = 10000
)

// 分页参数请求，可以被其他结构体匿名组合使用
type PagingReq struct {
	Pageno int64 `json:"pageno"` // 第几页
	Count  int64 `json:"count"`  // 每页几条数据
}

// 分页参数返回，可以被其他结构体匿名组合使用
type PagingResp struct {
	Pageno int64 `json:"pageno"` // 第几页
	Count  int64 `json:"count"`  // 每页几条数据
	Total  int64 `json:"total"`  // 总计多少条数据
}

// ModifyReqPagenoAndCount 检查分页参数，如果前端没有传 pageno 和 count，或者它们的值小于等于 0，此函数会给他们赋一个大于 0 的默认值，
// 此外，此函数会校验 count 值，如果超过允许的最大值 maxAllowedCount，则会返回错误
// 所有分页查询接口都需要在 schema 层调用此函数进行分页参数校验
// req 必须为指针，并包含 Count 和 Pageno 两个字段(名字必须一样)，且必须为 int64 类型，否则会 panic。
// req 结构体样例：
//
//	type HttpReq struct {
//		Pageno int64 `json:"pageno"`
//		Count  int64 `json:"count"`
//	}
func ModifyReqPagenoAndCount(req interface{}, maxAllowedCount int64) error {
	rv := reflect.ValueOf(req)
	// req 必须为指针
	if rv.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("incorrect parameters, %v should be pointer", req))
	}

	pageno := rv.Elem().FieldByName("Pageno")
	if pageno.Int() <= 0 {
		pageno.SetInt(defaultPageno)
	}
	count := rv.Elem().FieldByName("Count")
	if count.Int() <= 0 {
		count.SetInt(defaultCount)
	}

	if count.Int() > maxAllowedCount {
		return fmt.Errorf("count:%+v is too large, allowed max count:%+v", count.Int(), maxAllowedCount)
	}

	return nil
}
