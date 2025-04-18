package pageviewer

import (
	"time"

	"github.com/go-rod/rod"
)

var (
	DefaultWaitStableTimeout = time.Second * 20
)

func NewVisitOptions(opts ...VisitOption) *VisitOptions {
	// 生成配置项
	vo := &VisitOptions{
		PageOptions: &PageOptions{
			beforeRequest: nil,
			waitTimeout:   DefaultWaitStableTimeout,
		},
		browser: nil, // 初始化为nil，不要立即创建默认浏览器
	}
	for _, opt := range opts {
		opt(vo)
	}
	return vo
}

type VisitOptions struct {
	*PageOptions
	browser *Browser // 浏览器对象，只在Visit调用时有效
}

// VisitOption 访问配置项
type VisitOption func(vo *VisitOptions)

func WithWaitTimeout(timeout time.Duration) VisitOption {
	return func(vo *VisitOptions) {
		vo.waitTimeout = timeout
	}
}
func WithBrowser(browser *Browser) VisitOption {
	return func(vo *VisitOptions) {
		vo.browser = browser
	}
}

func WithBeforeRequest(f func(page *rod.Page) error) VisitOption {
	return func(vo *VisitOptions) {
		vo.beforeRequest = f
	}
}
func WithRemoveInvisibleDiv(removeInvisibleDiv bool) VisitOption {
	return func(vo *VisitOptions) {
		vo.removeInvisibleDiv = removeInvisibleDiv
	}
}

func Visit(u string, onPageLoad func(page *rod.Page) error, opts ...VisitOption) (err error) {
	// 生成配置项
	vo := NewVisitOptions(opts...)

	if vo.browser == nil {
		vo.browser = DefaultBrowser()
	}

	return vo.browser.run(u, onPageLoad, vo.PageOptions)
}
