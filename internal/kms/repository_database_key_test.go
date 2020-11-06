package kms_test

import (
	"context"
	"testing"
	"time"

	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/boundary/internal/iam"
	"github.com/hashicorp/boundary/internal/kms"
	"github.com/hashicorp/boundary/internal/oplog"
	wrapping "github.com/hashicorp/go-kms-wrapping"
	"github.com/hashicorp/go-kms-wrapping/wrappers/aead"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestRepository_CreateDatabaseKey(t *testing.T) {
	t.Parallel()
	conn, _ := db.TestSetup(t, "postgres")
	rw := db.New(conn)
	wrapper := db.TestWrapper(t)
	repo, err := kms.NewRepository(rw, rw)
	require.NoError(t, err)
	org, _ := iam.TestScopes(t, iam.TestRepo(t, conn, wrapper))
	require.NoError(t, conn.Where("1=1").Delete(kms.AllocRootKey()).Error)
	rk := kms.TestRootKey(t, conn, org.PublicId)
	_, rkvWrapper := kms.TestRootKeyVersion(t, conn, wrapper, rk.PrivateId)

	type args struct {
		scopeId    string
		key        []byte
		keyWrapper wrapping.Wrapper
		opt        []kms.Option
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantIsError error
	}{
		{
			name: "valid-org",
			args: args{
				scopeId:    org.PublicId,
				key:        []byte("test key"),
				keyWrapper: rkvWrapper,
			},
			wantErr: false,
		},
		{
			name: "nil-key",
			args: args{
				scopeId:    org.PublicId,
				keyWrapper: rkvWrapper,
			},
			wantErr:     true,
			wantIsError: errors.ErrInvalidParameter,
		},
		{
			name: "empty-key",
			args: args{
				scopeId:    org.PublicId,
				keyWrapper: rkvWrapper,
				key:        []byte(""),
			},
			wantErr:     true,
			wantIsError: errors.ErrInvalidParameter,
		},
		{
			name: "nil-wrapper",
			args: args{
				scopeId:    org.PublicId,
				key:        []byte("test key"),
				keyWrapper: nil,
			},
			wantErr:     true,
			wantIsError: errors.ErrInvalidParameter,
		},
		{
			name: "not-rkv-wrapper",
			args: args{
				scopeId:    org.PublicId,
				key:        []byte("test key"),
				keyWrapper: wrapper,
			},
			wantErr:     true,
			wantIsError: errors.ErrInvalidParameter,
		},
		{
			name: "wrapper-missing-id",
			args: args{
				scopeId: org.PublicId,
				key:     []byte("test key"),
				keyWrapper: func() wrapping.Wrapper {
					w := db.TestWrapper(t)
					_, err = w.(*aead.Wrapper).SetConfig(map[string]string{
						"key_id": "",
					})
					require.NoError(t, err)
					return w
				}(),
			},
			wantErr:     true,
			wantIsError: errors.ErrInvalidParameter,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert, require := assert.New(t), require.New(t)
			dk, dv, err := repo.CreateDatabaseKey(context.Background(), tt.args.keyWrapper, tt.args.key, tt.args.opt...)
			if tt.wantErr {
				assert.Error(err)
				assert.Nil(dk)
				if tt.wantIsError != nil {
					assert.True(errors.Is(err, tt.wantIsError))
				}
				return
			}
			require.NoError(err)
			assert.NotNil(dk.CreateTime)
			foundKey, err := repo.LookupDatabaseKey(context.Background(), dk.PrivateId)
			assert.NoError(err)
			assert.True(proto.Equal(foundKey, dk))

			// make sure there was no oplog written
			err = db.TestVerifyOplog(t, rw, dk.PrivateId, db.WithOperation(oplog.OpType_OP_TYPE_CREATE), db.WithCreateNotBefore(10*time.Second))
			assert.Error(err)
			assert.True(errors.Is(err, errors.ErrRecordNotFound))

			assert.NotNil(dv.CreateTime)
			foundKeyVersion, err := repo.LookupDatabaseKeyVersion(context.Background(), tt.args.keyWrapper, dv.PrivateId)
			assert.NoError(err)
			assert.True(proto.Equal(foundKeyVersion, dv))

			// make sure there was no oplog written
			err = db.TestVerifyOplog(t, rw, dv.PrivateId, db.WithOperation(oplog.OpType_OP_TYPE_CREATE), db.WithCreateNotBefore(10*time.Second))
			assert.Error(err)
			assert.True(errors.Is(err, errors.ErrRecordNotFound))
		})
	}
}

func TestRepository_DeleteDatabaseKey(t *testing.T) {
	t.Parallel()
	conn, _ := db.TestSetup(t, "postgres")
	rw := db.New(conn)
	wrapper := db.TestWrapper(t)
	repo, err := kms.NewRepository(rw, rw)
	require.NoError(t, err)
	org, _ := iam.TestScopes(t, iam.TestRepo(t, conn, wrapper))
	require.NoError(t, conn.Where("1=1").Delete(kms.AllocRootKey()).Error)
	rk := kms.TestRootKey(t, conn, org.PublicId)

	type args struct {
		key *kms.DatabaseKey
		opt []kms.Option
	}
	tests := []struct {
		name            string
		args            args
		wantRowsDeleted int
		wantErr         bool
		wantIsError     error
	}{
		{
			name: "valid",
			args: args{
				key: kms.TestDatabaseKey(t, conn, rk.PrivateId),
			},
			wantRowsDeleted: 1,
			wantErr:         false,
		},
		{
			name: "no-private-id",
			args: args{
				key: func() *kms.DatabaseKey {
					k := kms.AllocDatabaseKey()
					return &k
				}(),
			},
			wantRowsDeleted: 0,
			wantErr:         true,
			wantIsError:     errors.ErrInvalidParameter,
		},
		{
			name: "not-found",
			args: args{
				key: func() *kms.DatabaseKey {
					id, err := db.NewPublicId(kms.RootKeyPrefix)
					require.NoError(t, err)
					k := kms.AllocDatabaseKey()
					k.PrivateId = id
					require.NoError(t, err)
					return &k
				}(),
			},
			wantRowsDeleted: 0,
			wantErr:         true,
			wantIsError:     errors.ErrRecordNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert, require := assert.New(t), require.New(t)
			deletedRows, err := repo.DeleteDatabaseKey(context.Background(), tt.args.key.PrivateId, tt.args.opt...)
			if tt.wantErr {
				require.Error(err)
				assert.Equal(0, deletedRows)
				if tt.wantIsError != nil {
					assert.True(errors.Is(err, tt.wantIsError))
				}
				// make sure there was no oplog written
				err = db.TestVerifyOplog(t, rw, tt.args.key.PrivateId, db.WithOperation(oplog.OpType_OP_TYPE_DELETE), db.WithCreateNotBefore(10*time.Second))
				assert.Error(err)
				assert.True(errors.Is(errors.ErrRecordNotFound, err))
				return
			}
			require.NoError(err)
			assert.Equal(tt.wantRowsDeleted, deletedRows)
			foundKey, err := repo.LookupDatabaseKey(context.Background(), tt.args.key.PrivateId)
			assert.Error(err)
			assert.Nil(foundKey)
			assert.True(errors.Is(err, errors.ErrRecordNotFound))

			// make sure there was no oplog written
			err = db.TestVerifyOplog(t, rw, tt.args.key.PrivateId, db.WithOperation(oplog.OpType_OP_TYPE_DELETE), db.WithCreateNotBefore(10*time.Second))
			assert.Error(err)
			assert.True(errors.Is(errors.ErrRecordNotFound, err))
		})
	}
}

func TestRepository_ListDatabaseKeys(t *testing.T) {
	t.Parallel()
	conn, _ := db.TestSetup(t, "postgres")
	const testLimit = 10
	rw := db.New(conn)
	repo, err := kms.NewRepository(rw, rw, kms.WithLimit(testLimit))
	require.NoError(t, err)
	wrapper := db.TestWrapper(t)

	type args struct {
		opt []kms.Option
	}
	tests := []struct {
		name      string
		createCnt int
		args      args
		wantCnt   int
		wantErr   bool
	}{
		{
			name:      "no-limit",
			createCnt: repo.DefaultLimit() + 1,
			args: args{
				opt: []kms.Option{kms.WithLimit(-1)},
			},
			wantCnt: repo.DefaultLimit() + 1, // org and project both have keys, plus global scope
			wantErr: false,
		},
		{
			name:      "default-limit",
			createCnt: repo.DefaultLimit() + 1,
			args:      args{},
			wantCnt:   repo.DefaultLimit(),
			wantErr:   false,
		},
		{
			name:      "custom-limit",
			createCnt: repo.DefaultLimit() + 1,
			args: args{
				opt: []kms.Option{kms.WithLimit(3)},
			},
			wantCnt: 3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		require.NoError(t, conn.Where("1=1").Delete(kms.AllocRootKey()).Error)
		t.Run(tt.name, func(t *testing.T) {
			assert, require := assert.New(t), require.New(t)
			for i := 0; i < tt.createCnt; i++ {
				org, proj := iam.TestScopes(t, iam.TestRepo(t, conn, wrapper))
				require.NoError(conn.Where("scope_id in(?)", []interface{}{"global", org.PublicId, proj.PublicId}).Delete(kms.AllocRootKey()).Error)
				rk := kms.TestRootKey(t, conn, proj.PublicId)
				kms.TestDatabaseKey(t, conn, rk.PrivateId)
				require.NoError(err)
			}
			got, err := repo.ListDatabaseKeys(context.Background(), tt.args.opt...)
			if tt.wantErr {
				require.Error(err)
				return
			}
			require.NoError(err)
			assert.Equal(tt.wantCnt, len(got))
		})
	}
}
