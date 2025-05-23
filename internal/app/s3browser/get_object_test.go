package s3browser_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mosqueiro/s3browser/internal/app/s3browser"
	"github.com/mosqueiro/s3browser/internal/app/s3browser/mocks"
	"github.com/gorilla/mux"
	"github.com/matryer/is"
	"github.com/minio/minio-go/v7"
)

func TestHandleGetObject(t *testing.T) {
	t.Parallel()

	cases := []struct {
		it                   string
		getObjectFunc        func(context.Context, string, string, minio.GetObjectOptions) (*minio.Object, error)
		bucketName           string
		objectName           string
		expectedStatusCode   int
		expectedBodyContains string
	}{
		{
			it: "returns error if there is an S3 error",
			getObjectFunc: func(context.Context, string, string, minio.GetObjectOptions) (*minio.Object, error) {
				return nil, errS3
			},
			bucketName:           "BUCKET-NAME",
			objectName:           "OBJECT-NAME",
			expectedStatusCode:   http.StatusInternalServerError,
			expectedBodyContains: http.StatusText(http.StatusInternalServerError),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.it, func(t *testing.T) {
			t.Parallel()
			is := is.New(t)

			s3 := &mocks.S3Mock{
				GetObjectFunc: tc.getObjectFunc,
			}

			r := mux.NewRouter()
			r.Handle("/buckets/{bucketName}/objects/{objectName}", s3browser.HandleGetObject(s3, true)).Methods(http.MethodGet)

			ts := httptest.NewServer(r)
			defer ts.Close()

			resp, err := http.Get(fmt.Sprintf("%s/buckets/%s/objects/%s", ts.URL, tc.bucketName, tc.objectName))
			is.NoErr(err)
			defer func() {
				err = resp.Body.Close()
				is.NoErr(err)
			}()
			body, err := io.ReadAll(resp.Body)
			is.NoErr(err)

			is.Equal(tc.expectedStatusCode, resp.StatusCode)                 // status code
			is.True(strings.Contains(string(body), tc.expectedBodyContains)) // body
		})
	}
}
