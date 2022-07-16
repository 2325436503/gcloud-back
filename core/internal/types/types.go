// Code generated by goctl. DO NOT EDIT.
package types

type FileUploadChunkCompleteRequest struct {
	Key        string      `json:"key"`
	UploadId   string      `json:"upload_id"`
	CosObjects []CosObject `json:"cos_objects"`
}

type CosObject struct {
	PartNumber int    `json:"part_number"`
	Etag       string `json:"etag"`
}

type FileUploadChunkCompleteReply struct {
}

type FileUploadChunkRequest struct {
}

type FileUploadChunkReply struct {
	Etag string `json:"etag"` // MD5
}

type FileUploadPrepareRequest struct {
	Md5  string `json:"md5"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
}

type FileUploadPrepareReply struct {
	Identity string `json:"identity"`
	UploadId string `json:"upload_id"`
	Key      string `json:"key"`
}

type RefreshAuthorizationRequest struct {
}

type RefreshAuthorizationReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ShareBasicSaveRequest struct {
	RepositoryIdentity string `json:"repository_identity"`
	ParentId           int64  `json:"parent_id"`
}

type ShareBasicSaveReply struct {
	Identity string `json:"identity"`
}

type ShareBasicDetailRequest struct {
	Identity string `json:"identity"`
}

type ShareBasicDetailReply struct {
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Size               int64  `json:"size"`
	Path               string `json:"path"`
	ClickNum           int    `json:"click_num"`
}

type ShareBasicCreateRequest struct {
	UserRepositoryIdentity string `json:"user_repository_identity"`
	ExpiredTime            int    `json:"expired_time"`
}

type ShareBasicCreateReply struct {
	Identity string `json:"identity"`
}

type UserFileMoveRequest struct {
	Identity       string `json:"identity"`
	ParentIdentity string `json:"parent_identity"`
}

type UserFileMoveReply struct {
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFileDeleteReply struct {
}

type UserFolderCreateRequest struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type UserFolderCreateReply struct {
	Identity string `json:"identity"`
}

type UserFileNameUpdateRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileNameUpdateReply struct {
}

type UserFileListRequest struct {
	Id   int64 `json:"id,optional"`
	Page int   `json:"page,optional"`
	Size int   `json:"size,optional"`
}

type UserFileListReply struct {
	Count int64       `json:"count"`
	List  []*UserFile `json:"list"`
}

type UserFile struct {
	Id                 int64  `json:"id"`
	ParentId           int64  `json:"parent_id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Size               int64  `json:"size"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	UpdatedAt          string `json:"updated_at"`
}

type UserRepositorySaveRequest struct {
	RepositoryIdentity string `json:"repositoryIdentity"`
	ParentId           int64  `json:"parentId"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveReply struct {
}

type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
	Ext  string `json:"ext,optional"`
}

type FileUploadReply struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Msg          string `json:"msg"`
}

type UserDetailRequest struct {
}

type UserDetailReply struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Msg      string `json:"msg"`
	Identity string `json:"identity"`
}

type MailCodeSendRequest struct {
	Email string `json:"email"`
}

type MailCodeSendReply struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type UserRegisterReply struct {
	Msg string `json:"msg"`
}
