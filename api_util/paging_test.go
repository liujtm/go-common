package apiutil

import (
	"testing"
)

type HttpReq struct {
	PagingReq
}

func TestModifyReqPagenoAndCount(t *testing.T) {
	tests := []struct {
		name            string
		req             *HttpReq
		maxAllowedCount int64
		expectedPageno  int64
		expectedCount   int64
		expectError     bool
	}{
		{
			name:            "Valid request with pageno and count",
			req:             &HttpReq{PagingReq: PagingReq{Pageno: 2, Count: 10}},
			maxAllowedCount: 100,
			expectedPageno:  2,
			expectedCount:   10,
			expectError:     false,
		},
		{
			name:            "Default pageno and count",
			req:             &HttpReq{PagingReq: PagingReq{Pageno: 0, Count: 0}},
			maxAllowedCount: 100,
			expectedPageno:  defaultPageno,
			expectedCount:   defaultCount,
			expectError:     false,
		},
		{
			name:            "Count exceeds max allowed count",
			req:             &HttpReq{PagingReq: PagingReq{Pageno: 1, Count: 200}},
			maxAllowedCount: 100,
			expectedPageno:  1,
			expectedCount:   100,
			expectError:     true,
		},
		{
			name:            "Negative pageno and count",
			req:             &HttpReq{PagingReq: PagingReq{Pageno: -1, Count: -1}},
			maxAllowedCount: 100,
			expectedPageno:  defaultPageno,
			expectedCount:   defaultCount,
			expectError:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ModifyReqPagenoAndCount(tt.req, tt.maxAllowedCount)
			if err != nil {
				if !tt.expectError {
					t.Errorf("ModifyReqPagenoAndCount() error = %v, expectError %v", err, tt.expectError)
				}
				return
			}
			if tt.req.Pageno != tt.expectedPageno {
				t.Errorf("expected pageno = %v, got %v", tt.expectedPageno, tt.req.Pageno)
			}
			if tt.req.Count != tt.expectedCount {
				t.Errorf("expected count = %v, got %v", tt.expectedCount, tt.req.Count)
			}
		})
	}
}
