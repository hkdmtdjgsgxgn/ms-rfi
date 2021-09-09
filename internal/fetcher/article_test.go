package fetcher

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/hi20160616/exhtml"
	"github.com/hkdmtdjgsgxgn/ms-rfi/configs"
	"github.com/pkg/errors"
)

// pass test
func TestFetchArticle(t *testing.T) {
	tests := []struct {
		url string
		err error
	}{
		{"https://www.rfi.com/zhongwen/simp/world-57255390", ErrTimeOverDays},
		{"https://www.rfi.com/zhongwen/simp/uk-57264136", nil},
	}
	for _, tc := range tests {
		a := NewArticle()
		a, err := a.fetchArticle(tc.url)
		if err != nil {
			if !errors.Is(err, ErrTimeOverDays) {
				t.Error(err)
			} else {
				fmt.Println("ignore pass test: ", tc.url)
			}
		} else {
			fmt.Println("pass test: ", a.Content)
		}
	}
}

func TestFetchUpdateTime(t *testing.T) {
	tests := []struct {
		url string
		err error
	}{
		{"https://www.rfi.com/zhongwen/simp/uk-57264136", nil},
		{"https://www.rfi.com/zhongwen/simp/world-57341255", nil},
	}
	var err error
	if err := configs.Reset("../../"); err != nil {
		t.Error(err)
	}

	for _, tc := range tests {
		a := NewArticle()
		a.U, err = url.Parse(tc.url)
		if err != nil {
			t.Error(err)
		}
		// Dail
		a.raw, a.doc, err = exhtml.GetRawAndDoc(a.U, timeout)
		if err != nil {
			t.Error(err)
		}
		tt, err := a.fetchUpdateTime()
		if err != nil {
			t.Error(err)
		} else {
			ttt := tt.AsTime()
			fmt.Println(shanghai(ttt))
		}
	}

}
