// Copyright 2020-2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufapimodule

import (
	"context"
	"errors"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/bufbuild/buf/private/bufpkg/bufmanifest"
	"github.com/bufbuild/buf/private/bufpkg/bufmodule/bufmoduleref"
	"github.com/bufbuild/buf/private/gen/proto/connect/buf/alpha/registry/v1alpha1/registryv1alpha1connect"
	modulev1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/module/v1alpha1"
	registryv1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
	"github.com/bufbuild/buf/private/pkg/manifest"
	"github.com/bufbuild/buf/private/pkg/storage/storagemem"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDownload(t *testing.T) {
	t.Parallel()
	testDownload(
		t,
		"does-not-exist error",
		newMockDownloadService(
			t,
			withError(connect.NewError(connect.CodeNotFound, nil)),
		),
		"does not exist",
	)
	testDownload(
		t,
		"unexpected download service error",
		newMockDownloadService(
			t,
			withError(errors.New("internal")),
		),
		"internal",
	)
	testDownload(
		t,
		"success but response has all empty fields",
		newMockDownloadService(t),
		"expected non-nil manifest",
	)
	testDownload(
		t,
		"success with empty manifest module",
		newMockDownloadService(
			t,
			withBlobsFromMap(map[string][]byte{}),
		),
		"",
	)
	testDownload(
		t,
		"success with valid manifest module",
		newMockDownloadService(
			t,
			withBlobsFromMap(map[string][]byte{
				"test.proto": []byte(`syntax = "proto3";
message Test {}
`),
			}),
		),
		"",
	)
	testDownload(
		t,
		"manifest module with invalid lock file",
		newMockDownloadService(
			t,
			withBlobsFromMap(map[string][]byte{
				"buf.lock": []byte("invalid lock file"),
			}),
		),
		"failed to decode lock file",
	)
	testDownload(
		t,
		"no manifest",
		newMockDownloadService(
			t,
			withModule(&modulev1alpha1.Module{
				Files: []*modulev1alpha1.ModuleFile{
					{
						Path: "foo.proto",
					},
				},
			}),
		),
		"expected non-nil manifest",
	)
}

func testDownload(
	t *testing.T,
	desc string,
	mock *mockDownloadService,
	errorContains string,
) {
	t.Helper()
	t.Run(desc, func(t *testing.T) {
		t.Parallel()
		var moduleReaderOpts []ModuleReaderOption
		moduleReader := newModuleReader(mock.factory, moduleReaderOpts...)
		ctx := context.Background()
		pin, err := bufmoduleref.NewModulePin(
			"remote",
			"owner",
			"repository",
			"branch",
			"commit",
			"digest",
			time.Now(),
		)
		require.NoError(t, err)
		module, err := moduleReader.GetModule(ctx, pin)
		if errorContains != "" {
			assert.ErrorContains(t, err, errorContains)
		} else {
			assert.NotNil(t, module)
			assert.NoError(t, err)
			for _, path := range module.Manifest().Paths() {
				moduleFile, err := module.GetModuleFile(ctx, path)
				require.NoError(t, err)
				assert.Equal(t, pin.Commit(), moduleFile.Commit())
			}
		}
	})
}

type mockDownloadService struct {
	module       *modulev1alpha1.Module
	manifestBlob *modulev1alpha1.Blob
	blobs        []*modulev1alpha1.Blob
	err          error
}

type option interface {
	apply(*mockDownloadService) error
}

type filemap map[string][]byte

func (fm filemap) apply(m *mockDownloadService) error {
	bucket, err := storagemem.NewReadBucket(fm)
	if err != nil {
		return err
	}
	ctx := context.Background()
	moduleManifest, blobSet, err := manifest.NewFromBucket(ctx, bucket)
	if err != nil {
		return err
	}
	mBlob, err := moduleManifest.Blob()
	if err != nil {
		return err
	}
	m.manifestBlob, err = bufmanifest.AsProtoBlob(ctx, mBlob)
	if err != nil {
		return err
	}
	blobs := blobSet.Blobs()
	m.blobs = make([]*modulev1alpha1.Blob, 0, len(blobs))
	for _, blob := range blobs {
		protoBlob, err := bufmanifest.AsProtoBlob(ctx, blob)
		if err != nil {
			return err
		}
		m.blobs = append(m.blobs, protoBlob)
	}
	return nil
}

func withBlobsFromMap(files map[string][]byte) option {
	return filemap(files)
}

type retErr struct{ err error }

func (re retErr) apply(m *mockDownloadService) error {
	m.err = re.err
	return nil
}

func withError(err error) option {
	return retErr{err: err}
}

type retModule struct{ module *modulev1alpha1.Module }

func (rm retModule) apply(m *mockDownloadService) error {
	m.module = rm.module
	return nil
}

func withModule(module *modulev1alpha1.Module) option {
	return retModule{module: module}
}

func newMockDownloadService(
	t *testing.T,
	opts ...option,
) *mockDownloadService {
	m := &mockDownloadService{}
	for _, opt := range opts {
		if err := opt.apply(m); err != nil {
			t.Error(err)
		}
	}
	return m
}

func (m *mockDownloadService) factory(_ string) registryv1alpha1connect.DownloadServiceClient {
	return m
}

func (m *mockDownloadService) Download(
	context.Context,
	*connect.Request[registryv1alpha1.DownloadRequest],
) (*connect.Response[registryv1alpha1.DownloadResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("Download RPC should not be called, use DownloadManifestAndBlobs RPC instead"))
}

func (m *mockDownloadService) DownloadManifestAndBlobs(
	context.Context,
	*connect.Request[registryv1alpha1.DownloadManifestAndBlobsRequest],
) (*connect.Response[registryv1alpha1.DownloadManifestAndBlobsResponse], error) {
	if m.err != nil {
		return nil, m.err
	}
	return connect.NewResponse(&registryv1alpha1.DownloadManifestAndBlobsResponse{
		Manifest: m.manifestBlob,
		Blobs:    m.blobs,
	}), nil
}
