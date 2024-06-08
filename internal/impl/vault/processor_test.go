package vault

import (
	"context"

	api "github.com/redpanda-data/connect/v4/internal/impl/vault/openapi"
)

var (
	h api.Handler = (*handler)(nil)
)

func newServer() (*api.Server, error) {
	return api.NewServer(&handler{})
}

type handler struct {
}

// AppRoleDeleteBindSecretID implements app-role-delete-bind-secret-id operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/bind-secret-id
func (h *handler) AppRoleDeleteBindSecretID(ctx context.Context, params api.AppRoleDeleteBindSecretIDParams) error {
	panic("not implemented")
}

// AppRoleDeleteBoundCidrList implements app-role-delete-bound-cidr-list operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list
func (h *handler) AppRoleDeleteBoundCidrList(ctx context.Context, params api.AppRoleDeleteBoundCidrListParams) error {
	panic("not implemented")
}

// AppRoleDeletePeriod implements app-role-delete-period operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/period
func (h *handler) AppRoleDeletePeriod(ctx context.Context, params api.AppRoleDeletePeriodParams) error {
	panic("not implemented")
}

// AppRoleDeletePolicies implements app-role-delete-policies operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/policies
func (h *handler) AppRoleDeletePolicies(ctx context.Context, params api.AppRoleDeletePoliciesParams) error {
	panic("not implemented")
}

// AppRoleDeleteRole implements app-role-delete-role operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}
func (h *handler) AppRoleDeleteRole(ctx context.Context, params api.AppRoleDeleteRoleParams) error {
	panic("not implemented")
}

// AppRoleDeleteSecretIDBoundCidrs implements app-role-delete-secret-id-bound-cidrs operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs
func (h *handler) AppRoleDeleteSecretIDBoundCidrs(ctx context.Context, params api.AppRoleDeleteSecretIDBoundCidrsParams) error {
	panic("not implemented")
}

// AppRoleDeleteSecretIDNumUses implements app-role-delete-secret-id-num-uses operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses
func (h *handler) AppRoleDeleteSecretIDNumUses(ctx context.Context, params api.AppRoleDeleteSecretIDNumUsesParams) error {
	panic("not implemented")
}

// AppRoleDeleteSecretIDTTL implements app-role-delete-secret-id-ttl operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl
func (h *handler) AppRoleDeleteSecretIDTTL(ctx context.Context, params api.AppRoleDeleteSecretIDTTLParams) error {
	panic("not implemented")
}

// AppRoleDeleteTokenBoundCidrs implements app-role-delete-token-bound-cidrs operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs
func (h *handler) AppRoleDeleteTokenBoundCidrs(ctx context.Context, params api.AppRoleDeleteTokenBoundCidrsParams) error {
	panic("not implemented")
}

// AppRoleDeleteTokenMaxTTL implements app-role-delete-token-max-ttl operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/token-max-ttl
func (h *handler) AppRoleDeleteTokenMaxTTL(ctx context.Context, params api.AppRoleDeleteTokenMaxTTLParams) error {
	panic("not implemented")
}

// AppRoleDeleteTokenNumUses implements app-role-delete-token-num-uses operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/token-num-uses
func (h *handler) AppRoleDeleteTokenNumUses(ctx context.Context, params api.AppRoleDeleteTokenNumUsesParams) error {
	panic("not implemented")
}

// AppRoleDeleteTokenTTL implements app-role-delete-token-ttl operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/token-ttl
func (h *handler) AppRoleDeleteTokenTTL(ctx context.Context, params api.AppRoleDeleteTokenTTLParams) error {
	panic("not implemented")
}

// AppRoleDestroySecretID implements app-role-destroy-secret-id operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/secret-id/destroy
func (h *handler) AppRoleDestroySecretID(ctx context.Context, req *api.AppRoleDestroySecretIdRequest, params api.AppRoleDestroySecretIDParams) error {
	panic("not implemented")
}

// AppRoleDestroySecretIDByAccessor implements app-role-destroy-secret-id-by-accessor operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy
func (h *handler) AppRoleDestroySecretIDByAccessor(ctx context.Context, req *api.AppRoleDestroySecretIdByAccessorRequest, params api.AppRoleDestroySecretIDByAccessorParams) error {
	panic("not implemented")
}

// AppRoleDestroySecretIDByAccessor2 implements app-role-destroy-secret-id-by-accessor2 operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/destroy
func (h *handler) AppRoleDestroySecretIDByAccessor2(ctx context.Context, params api.AppRoleDestroySecretIDByAccessor2Params) error {
	panic("not implemented")
}

// AppRoleDestroySecretId2 implements app-role-destroy-secret-id2 operation.
//
// DELETE /auth/{approle_mount_path}/role/{role_name}/secret-id/destroy
func (h *handler) AppRoleDestroySecretId2(ctx context.Context, params api.AppRoleDestroySecretId2Params) error {
	panic("not implemented")
}

// AppRoleListRoles implements app-role-list-roles operation.
//
// GET /auth/{approle_mount_path}/role/
func (h *handler) AppRoleListRoles(ctx context.Context, params api.AppRoleListRolesParams) (*api.StandardListResponse, error) {
	panic("not implemented")
}

// AppRoleListSecretIds implements app-role-list-secret-ids operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/secret-id/
func (h *handler) AppRoleListSecretIds(ctx context.Context, params api.AppRoleListSecretIdsParams) (*api.StandardListResponse, error) {
	panic("not implemented")
}

// AppRoleLogin implements app-role-login operation.
//
// POST /auth/{approle_mount_path}/login
func (h *handler) AppRoleLogin(ctx context.Context, req *api.AppRoleLoginRequest, params api.AppRoleLoginParams) error {
	panic("not implemented")
}

// AppRoleLookUpSecretID implements app-role-look-up-secret-id operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/secret-id/lookup
func (h *handler) AppRoleLookUpSecretID(ctx context.Context, req *api.AppRoleLookUpSecretIdRequest, params api.AppRoleLookUpSecretIDParams) (*api.AppRoleLookUpSecretIdResponse, error) {
	panic("not implemented")
}

// AppRoleLookUpSecretIDByAccessor implements app-role-look-up-secret-id-by-accessor operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/secret-id-accessor/lookup
func (h *handler) AppRoleLookUpSecretIDByAccessor(ctx context.Context, req *api.AppRoleLookUpSecretIdByAccessorRequest, params api.AppRoleLookUpSecretIDByAccessorParams) (*api.AppRoleLookUpSecretIdByAccessorResponse, error) {
	panic("not implemented")
}

// AppRoleReadBindSecretID implements app-role-read-bind-secret-id operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/bind-secret-id
func (h *handler) AppRoleReadBindSecretID(ctx context.Context, params api.AppRoleReadBindSecretIDParams) (*api.AppRoleReadBindSecretIdResponse, error) {
	panic("not implemented")
}

// AppRoleReadBoundCidrList implements app-role-read-bound-cidr-list operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list
func (h *handler) AppRoleReadBoundCidrList(ctx context.Context, params api.AppRoleReadBoundCidrListParams) (*api.AppRoleReadBoundCidrListResponse, error) {
	panic("not implemented")
}

// AppRoleReadLocalSecretIds implements app-role-read-local-secret-ids operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/local-secret-ids
func (h *handler) AppRoleReadLocalSecretIds(ctx context.Context, params api.AppRoleReadLocalSecretIdsParams) (*api.AppRoleReadLocalSecretIdsResponse, error) {
	panic("not implemented")
}

// AppRoleReadPeriod implements app-role-read-period operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/period
func (h *handler) AppRoleReadPeriod(ctx context.Context, params api.AppRoleReadPeriodParams) (*api.AppRoleReadPeriodResponse, error) {
	panic("not implemented")
}

// AppRoleReadPolicies implements app-role-read-policies operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/policies
func (h *handler) AppRoleReadPolicies(ctx context.Context, params api.AppRoleReadPoliciesParams) (*api.AppRoleReadPoliciesResponse, error) {
	panic("not implemented")
}

// AppRoleReadRole implements app-role-read-role operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}
func (h *handler) AppRoleReadRole(ctx context.Context, params api.AppRoleReadRoleParams) (*api.AppRoleReadRoleResponse, error) {
	panic("not implemented")
}

// AppRoleReadRoleID implements app-role-read-role-id operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/role-id
func (h *handler) AppRoleReadRoleID(ctx context.Context, params api.AppRoleReadRoleIDParams) (*api.AppRoleReadRoleIdResponse, error) {
	panic("not implemented")
}

// AppRoleReadSecretIDBoundCidrs implements app-role-read-secret-id-bound-cidrs operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs
func (h *handler) AppRoleReadSecretIDBoundCidrs(ctx context.Context, params api.AppRoleReadSecretIDBoundCidrsParams) (*api.AppRoleReadSecretIdBoundCidrsResponse, error) {
	panic("not implemented")
}

// AppRoleReadSecretIDNumUses implements app-role-read-secret-id-num-uses operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses
func (h *handler) AppRoleReadSecretIDNumUses(ctx context.Context, params api.AppRoleReadSecretIDNumUsesParams) (*api.AppRoleReadSecretIdNumUsesResponse, error) {
	panic("not implemented")
}

// AppRoleReadSecretIDTTL implements app-role-read-secret-id-ttl operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl
func (h *handler) AppRoleReadSecretIDTTL(ctx context.Context, params api.AppRoleReadSecretIDTTLParams) (*api.AppRoleReadSecretIdTtlResponse, error) {
	panic("not implemented")
}

// AppRoleReadTokenBoundCidrs implements app-role-read-token-bound-cidrs operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs
func (h *handler) AppRoleReadTokenBoundCidrs(ctx context.Context, params api.AppRoleReadTokenBoundCidrsParams) (*api.AppRoleReadTokenBoundCidrsResponse, error) {
	panic("not implemented")
}

// AppRoleReadTokenMaxTTL implements app-role-read-token-max-ttl operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/token-max-ttl
func (h *handler) AppRoleReadTokenMaxTTL(ctx context.Context, params api.AppRoleReadTokenMaxTTLParams) (*api.AppRoleReadTokenMaxTtlResponse, error) {
	panic("not implemented")
}

// AppRoleReadTokenNumUses implements app-role-read-token-num-uses operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/token-num-uses
func (h *handler) AppRoleReadTokenNumUses(ctx context.Context, params api.AppRoleReadTokenNumUsesParams) (*api.AppRoleReadTokenNumUsesResponse, error) {
	panic("not implemented")
}

// AppRoleReadTokenTTL implements app-role-read-token-ttl operation.
//
// GET /auth/{approle_mount_path}/role/{role_name}/token-ttl
func (h *handler) AppRoleReadTokenTTL(ctx context.Context, params api.AppRoleReadTokenTTLParams) (*api.AppRoleReadTokenTtlResponse, error) {
	panic("not implemented")
}

// AppRoleTidySecretID implements app-role-tidy-secret-id operation.
//
// POST /auth/{approle_mount_path}/tidy/secret-id
func (h *handler) AppRoleTidySecretID(ctx context.Context, params api.AppRoleTidySecretIDParams) error {
	panic("not implemented")
}

// AppRoleWriteBindSecretID implements app-role-write-bind-secret-id operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/bind-secret-id
func (h *handler) AppRoleWriteBindSecretID(ctx context.Context, req *api.AppRoleWriteBindSecretIdRequest, params api.AppRoleWriteBindSecretIDParams) error {
	panic("not implemented")
}

// AppRoleWriteBoundCidrList implements app-role-write-bound-cidr-list operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/bound-cidr-list
func (h *handler) AppRoleWriteBoundCidrList(ctx context.Context, req *api.AppRoleWriteBoundCidrListRequest, params api.AppRoleWriteBoundCidrListParams) error {
	panic("not implemented")
}

// AppRoleWriteCustomSecretID implements app-role-write-custom-secret-id operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/custom-secret-id
func (h *handler) AppRoleWriteCustomSecretID(ctx context.Context, req *api.AppRoleWriteCustomSecretIdRequest, params api.AppRoleWriteCustomSecretIDParams) (*api.AppRoleWriteCustomSecretIdResponse, error) {
	panic("not implemented")
}

// AppRoleWritePeriod implements app-role-write-period operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/period
func (h *handler) AppRoleWritePeriod(ctx context.Context, req *api.AppRoleWritePeriodRequest, params api.AppRoleWritePeriodParams) error {
	panic("not implemented")
}

// AppRoleWritePolicies implements app-role-write-policies operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/policies
func (h *handler) AppRoleWritePolicies(ctx context.Context, req *api.AppRoleWritePoliciesRequest, params api.AppRoleWritePoliciesParams) error {
	panic("not implemented")
}

// AppRoleWriteRole implements app-role-write-role operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}
func (h *handler) AppRoleWriteRole(ctx context.Context, req *api.AppRoleWriteRoleRequest, params api.AppRoleWriteRoleParams) error {
	panic("not implemented")
}

// AppRoleWriteRoleID implements app-role-write-role-id operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/role-id
func (h *handler) AppRoleWriteRoleID(ctx context.Context, req *api.AppRoleWriteRoleIdRequest, params api.AppRoleWriteRoleIDParams) error {
	panic("not implemented")
}

// AppRoleWriteSecretID implements app-role-write-secret-id operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/secret-id
func (h *handler) AppRoleWriteSecretID(ctx context.Context, req *api.AppRoleWriteSecretIdRequest, params api.AppRoleWriteSecretIDParams) (*api.AppRoleWriteSecretIdResponse, error) {
	panic("not implemented")
}

// AppRoleWriteSecretIDBoundCidrs implements app-role-write-secret-id-bound-cidrs operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/secret-id-bound-cidrs
func (h *handler) AppRoleWriteSecretIDBoundCidrs(ctx context.Context, req *api.AppRoleWriteSecretIdBoundCidrsRequest, params api.AppRoleWriteSecretIDBoundCidrsParams) error {
	panic("not implemented")
}

// AppRoleWriteSecretIDNumUses implements app-role-write-secret-id-num-uses operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/secret-id-num-uses
func (h *handler) AppRoleWriteSecretIDNumUses(ctx context.Context, req *api.AppRoleWriteSecretIdNumUsesRequest, params api.AppRoleWriteSecretIDNumUsesParams) error {
	panic("not implemented")
}

// AppRoleWriteSecretIDTTL implements app-role-write-secret-id-ttl operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/secret-id-ttl
func (h *handler) AppRoleWriteSecretIDTTL(ctx context.Context, req *api.AppRoleWriteSecretIdTtlRequest, params api.AppRoleWriteSecretIDTTLParams) error {
	panic("not implemented")
}

// AppRoleWriteTokenBoundCidrs implements app-role-write-token-bound-cidrs operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/token-bound-cidrs
func (h *handler) AppRoleWriteTokenBoundCidrs(ctx context.Context, req *api.AppRoleWriteTokenBoundCidrsRequest, params api.AppRoleWriteTokenBoundCidrsParams) error {
	panic("not implemented")
}

// AppRoleWriteTokenMaxTTL implements app-role-write-token-max-ttl operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/token-max-ttl
func (h *handler) AppRoleWriteTokenMaxTTL(ctx context.Context, req *api.AppRoleWriteTokenMaxTtlRequest, params api.AppRoleWriteTokenMaxTTLParams) error {
	panic("not implemented")
}

// AppRoleWriteTokenNumUses implements app-role-write-token-num-uses operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/token-num-uses
func (h *handler) AppRoleWriteTokenNumUses(ctx context.Context, req *api.AppRoleWriteTokenNumUsesRequest, params api.AppRoleWriteTokenNumUsesParams) error {
	panic("not implemented")
}

// AppRoleWriteTokenTTL implements app-role-write-token-ttl operation.
//
// POST /auth/{approle_mount_path}/role/{role_name}/token-ttl
func (h *handler) AppRoleWriteTokenTTL(ctx context.Context, req *api.AppRoleWriteTokenTtlRequest, params api.AppRoleWriteTokenTTLParams) error {
	panic("not implemented")
}

// KvV2Configure implements kv-v2-configure operation.
//
// Configure backend level settings that are applied to every key in the key-value store.
//
// POST /{kv_v2_mount_path}/config
func (h *handler) KvV2Configure(ctx context.Context, req *api.KvV2ConfigureRequest, params api.KvV2ConfigureParams) error {
	panic("not implemented")
}

// KvV2Delete implements kv-v2-delete operation.
//
// DELETE /{kv_v2_mount_path}/data/{path}
func (h *handler) KvV2Delete(ctx context.Context, params api.KvV2DeleteParams) error {
	panic("not implemented")
}

// KvV2DeleteMetadataAndAllVersions implements kv-v2-delete-metadata-and-all-versions operation.
//
// DELETE /{kv_v2_mount_path}/metadata/{path}
func (h *handler) KvV2DeleteMetadataAndAllVersions(ctx context.Context, params api.KvV2DeleteMetadataAndAllVersionsParams) error {
	panic("not implemented")
}

// KvV2DeleteVersions implements kv-v2-delete-versions operation.
//
// POST /{kv_v2_mount_path}/delete/{path}
func (h *handler) KvV2DeleteVersions(ctx context.Context, req *api.KvV2DeleteVersionsRequest, params api.KvV2DeleteVersionsParams) error {
	panic("not implemented")
}

// KvV2DestroyVersions implements kv-v2-destroy-versions operation.
//
// POST /{kv_v2_mount_path}/destroy/{path}
func (h *handler) KvV2DestroyVersions(ctx context.Context, req *api.KvV2DestroyVersionsRequest, params api.KvV2DestroyVersionsParams) error {
	panic("not implemented")
}

// KvV2List implements kv-v2-list operation.
//
// GET /{kv_v2_mount_path}/metadata/{path}/
func (h *handler) KvV2List(ctx context.Context, params api.KvV2ListParams) (*api.StandardListResponse, error) {
	panic("not implemented")
}

// KvV2Read implements kv-v2-read operation.
//
// GET /{kv_v2_mount_path}/data/{path}
func (h *handler) KvV2Read(ctx context.Context, params api.KvV2ReadParams) (*api.KvV2ReadResponse, error) {
	panic("not implemented")
}

// KvV2ReadConfiguration implements kv-v2-read-configuration operation.
//
// Read the backend level settings.
//
// GET /{kv_v2_mount_path}/config
func (h *handler) KvV2ReadConfiguration(ctx context.Context, params api.KvV2ReadConfigurationParams) (*api.KvV2ReadConfigurationResponse, error) {
	panic("not implemented")
}

// KvV2ReadMetadata implements kv-v2-read-metadata operation.
//
// GET /{kv_v2_mount_path}/metadata/{path}
func (h *handler) KvV2ReadMetadata(ctx context.Context, params api.KvV2ReadMetadataParams) (*api.KvV2ReadMetadataResponse, error) {
	panic("not implemented")
}

// KvV2ReadSubkeys implements kv-v2-read-subkeys operation.
//
// GET /{kv_v2_mount_path}/subkeys/{path}
func (h *handler) KvV2ReadSubkeys(ctx context.Context, params api.KvV2ReadSubkeysParams) (*api.KvV2ReadSubkeysResponse, error) {
	panic("not implemented")
}

// KvV2UndeleteVersions implements kv-v2-undelete-versions operation.
//
// POST /{kv_v2_mount_path}/undelete/{path}
func (h *handler) KvV2UndeleteVersions(ctx context.Context, req *api.KvV2UndeleteVersionsRequest, params api.KvV2UndeleteVersionsParams) error {
	panic("not implemented")
}

// KvV2Write implements kv-v2-write operation.
//
// POST /{kv_v2_mount_path}/data/{path}
func (h *handler) KvV2Write(ctx context.Context, req *api.KvV2WriteRequest, params api.KvV2WriteParams) (*api.KvV2WriteResponse, error) {
	panic("not implemented")
}

// KvV2WriteMetadata implements kv-v2-write-metadata operation.
//
// POST /{kv_v2_mount_path}/metadata/{path}
func (h *handler) KvV2WriteMetadata(ctx context.Context, req *api.KvV2WriteMetadataRequest, params api.KvV2WriteMetadataParams) error {
	panic("not implemented")
}
