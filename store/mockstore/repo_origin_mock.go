// generated by gen-mocks; DO NOT EDIT

package mockstore

import (
	"crypto/rsa"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/context"
	"sourcegraph.com/sourcegraph/go-sourcegraph/sourcegraph"
	"src.sourcegraph.com/sourcegraph/store"
)

type RepoOrigin struct {
	BrandName_ func() string
	Host_      func() string
}

func (s *RepoOrigin) BrandName() string { return s.BrandName_() }

func (s *RepoOrigin) Host() string { return s.Host_() }

var _ store.RepoOrigin = (*RepoOrigin)(nil)

type RepoOriginWithPushHooks struct {
	IsPushHookEnabled_ func(ctx context.Context, repo string) (bool, error)
	EnablePushHook_    func(ctx context.Context, repo string) error
}

func (s *RepoOriginWithPushHooks) IsPushHookEnabled(ctx context.Context, repo string) (bool, error) {
	return s.IsPushHookEnabled_(ctx, repo)
}

func (s *RepoOriginWithPushHooks) EnablePushHook(ctx context.Context, repo string) error {
	return s.EnablePushHook_(ctx, repo)
}

var _ store.RepoOriginWithPushHooks = (*RepoOriginWithPushHooks)(nil)

type RepoOriginWithCommitStatuses struct {
	IsCommitStatusCapable_ func(ctx context.Context, repo string) (bool, error)
	PublishCommitStatus_   func(ctx context.Context, repo string, status *sourcegraph.RepoStatus) error
}

func (s *RepoOriginWithCommitStatuses) IsCommitStatusCapable(ctx context.Context, repo string) (bool, error) {
	return s.IsCommitStatusCapable_(ctx, repo)
}

func (s *RepoOriginWithCommitStatuses) PublishCommitStatus(ctx context.Context, repo string, status *sourcegraph.RepoStatus) error {
	return s.PublishCommitStatus_(ctx, repo, status)
}

var _ store.RepoOriginWithCommitStatuses = (*RepoOriginWithCommitStatuses)(nil)

type RepoOriginWithAuthorizedSSHKeys struct {
	IsSSHKeyAuthorized_ func(ctx context.Context, repo string, key ssh.PublicKey) (bool, error)
	AuthorizeSSHKey_    func(ctx context.Context, repo string, key ssh.PublicKey) error
	DeleteSSHKey_       func(ctx context.Context, repo string, key ssh.PublicKey) error
}

func (s *RepoOriginWithAuthorizedSSHKeys) IsSSHKeyAuthorized(ctx context.Context, repo string, key ssh.PublicKey) (bool, error) {
	return s.IsSSHKeyAuthorized_(ctx, repo, key)
}

func (s *RepoOriginWithAuthorizedSSHKeys) AuthorizeSSHKey(ctx context.Context, repo string, key ssh.PublicKey) error {
	return s.AuthorizeSSHKey_(ctx, repo, key)
}

func (s *RepoOriginWithAuthorizedSSHKeys) DeleteSSHKey(ctx context.Context, repo string, key ssh.PublicKey) error {
	return s.DeleteSSHKey_(ctx, repo, key)
}

var _ store.RepoOriginWithAuthorizedSSHKeys = (*RepoOriginWithAuthorizedSSHKeys)(nil)

type MirroredRepoSSHKeys struct {
	Create_ func(ctx context.Context, repo string, privKey *rsa.PrivateKey) error
	GetPEM_ func(ctx context.Context, repo string) ([]byte, error)
	Delete_ func(ctx context.Context, repo string) error
}

func (s *MirroredRepoSSHKeys) Create(ctx context.Context, repo string, privKey *rsa.PrivateKey) error {
	return s.Create_(ctx, repo, privKey)
}

func (s *MirroredRepoSSHKeys) GetPEM(ctx context.Context, repo string) ([]byte, error) {
	return s.GetPEM_(ctx, repo)
}

func (s *MirroredRepoSSHKeys) Delete(ctx context.Context, repo string) error {
	return s.Delete_(ctx, repo)
}

var _ store.MirroredRepoSSHKeys = (*MirroredRepoSSHKeys)(nil)
