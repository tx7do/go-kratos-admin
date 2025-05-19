package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/file"

	fileV1 "kratos-admin/api/gen/go/file/service/v1"
)

type FileRepo struct {
	data *Data
	log  *log.Helper
}

func NewFileRepo(data *Data, logger log.Logger) *FileRepo {
	l := log.NewHelper(log.With(logger, "module", "file/repo/admin-service"))
	return &FileRepo{
		data: data,
		log:  l,
	}
}

func (r *FileRepo) toProtoProvider(in *file.Provider) *fileV1.OSSProvider {
	if in == nil {
		return nil
	}

	switch *in {
	case file.ProviderMinIO:
		return trans.Ptr(fileV1.OSSProvider_MinIO)

	case file.ProviderAliyun:
		return trans.Ptr(fileV1.OSSProvider_Aliyun)

	case file.ProviderAWS:
		return trans.Ptr(fileV1.OSSProvider_AWS)

	case file.ProviderAzure:
		return trans.Ptr(fileV1.OSSProvider_Azure)

	case file.ProviderBaidu:
		return trans.Ptr(fileV1.OSSProvider_Baidu)

	case file.ProviderQiniu:
		return trans.Ptr(fileV1.OSSProvider_Qiniu)

	case file.ProviderTencent:
		return trans.Ptr(fileV1.OSSProvider_Tencent)

	case file.ProviderGoogle:
		return trans.Ptr(fileV1.OSSProvider_Google)

	case file.ProviderHuawei:
		return trans.Ptr(fileV1.OSSProvider_Huawei)

	case file.ProviderQCloud:
		return trans.Ptr(fileV1.OSSProvider_QCloud)

	case file.ProviderLocal:
		return trans.Ptr(fileV1.OSSProvider_Local)

	default:
		return nil
	}
}

func (r *FileRepo) toEntProvider(in *fileV1.OSSProvider) *file.Provider {
	if in == nil {
		return nil
	}

	switch *in {
	case fileV1.OSSProvider_MinIO:
		return trans.Ptr(file.ProviderMinIO)

	case fileV1.OSSProvider_Aliyun:
		return trans.Ptr(file.ProviderAliyun)

	case fileV1.OSSProvider_AWS:
		return trans.Ptr(file.ProviderAWS)

	case fileV1.OSSProvider_Azure:
		return trans.Ptr(file.ProviderAzure)

	case fileV1.OSSProvider_Baidu:
		return trans.Ptr(file.ProviderBaidu)

	case fileV1.OSSProvider_Qiniu:
		return trans.Ptr(file.ProviderQiniu)

	case fileV1.OSSProvider_Tencent:
		return trans.Ptr(file.ProviderTencent)

	case fileV1.OSSProvider_Google:
		return trans.Ptr(file.ProviderGoogle)

	case fileV1.OSSProvider_Huawei:
		return trans.Ptr(file.ProviderHuawei)

	case fileV1.OSSProvider_QCloud:
		return trans.Ptr(file.ProviderQCloud)

	case fileV1.OSSProvider_Local:
		return trans.Ptr(file.ProviderLocal)

	default:
		return nil
	}
}

func (r *FileRepo) convertEntToProto(in *ent.File) *fileV1.File {
	if in == nil {
		return nil
	}
	return &fileV1.File{
		Id:            trans.Ptr(in.ID),
		Provider:      r.toProtoProvider(in.Provider),
		BucketName:    in.BucketName,
		FileDirectory: in.FileDirectory,
		FileGuid:      in.FileGUID,
		SaveFileName:  in.SaveFileName,
		FileName:      in.FileName,
		Extension:     in.Extension,
		Size:          in.Size,
		SizeFormat:    in.SizeFormat,
		LinkUrl:       in.LinkURL,
		Md5:           in.Md5,
		CreateBy:      in.CreateBy,
		//UpdateBy:      in.UpdateBy,
		CreateTime: timeutil.TimeToTimeString(in.CreateTime),
		UpdateTime: timeutil.TimeToTimeString(in.UpdateTime),
		DeleteTime: timeutil.TimeToTimeString(in.DeleteTime),
	}
}

func (r *FileRepo) Count(ctx context.Context, whereCond []func(s *sql.Selector)) (int, error) {
	builder := r.data.db.Client().File.Query()
	if len(whereCond) != 0 {
		builder.Modify(whereCond...)
	}

	count, err := builder.Count(ctx)
	if err != nil {
		r.log.Errorf("query count failed: %s", err.Error())
		return 0, fileV1.ErrorInternalServerError("query count failed")
	}

	return count, nil
}

func (r *FileRepo) List(ctx context.Context, req *pagination.PagingRequest) (*fileV1.ListFileResponse, error) {
	if req == nil {
		return nil, fileV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().File.Query()

	err, whereSelectors, querySelectors := entgo.BuildQuerySelector(
		req.GetQuery(), req.GetOrQuery(),
		req.GetPage(), req.GetPageSize(), req.GetNoPaging(),
		req.GetOrderBy(), file.FieldCreateTime,
		req.GetFieldMask().GetPaths(),
	)
	if err != nil {
		r.log.Errorf("parse list param error [%s]", err.Error())
		return nil, fileV1.ErrorBadRequest("invalid query parameter")
	}

	if querySelectors != nil {
		builder.Modify(querySelectors...)
	}

	results, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, fileV1.ErrorInternalServerError("query list failed")
	}

	items := make([]*fileV1.File, 0, len(results))
	for _, res := range results {
		item := r.convertEntToProto(res)
		items = append(items, item)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &fileV1.ListFileResponse{
		Total: uint32(count),
		Items: items,
	}, err
}

func (r *FileRepo) IsExist(ctx context.Context, id uint32) (bool, error) {
	exist, err := r.data.db.Client().File.Query().
		Where(file.IDEQ(id)).
		Exist(ctx)
	if err != nil {
		r.log.Errorf("query exist failed: %s", err.Error())
		return false, fileV1.ErrorInternalServerError("query exist failed")
	}
	return exist, nil
}

func (r *FileRepo) Get(ctx context.Context, req *fileV1.GetFileRequest) (*fileV1.File, error) {
	if req == nil {
		return nil, fileV1.ErrorBadRequest("invalid parameter")
	}

	ret, err := r.data.db.Client().File.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fileV1.ErrorFileNotFound("file not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, fileV1.ErrorInternalServerError("query data failed")
	}

	return r.convertEntToProto(ret), nil
}

func (r *FileRepo) Create(ctx context.Context, req *fileV1.CreateFileRequest) error {
	if req == nil || req.Data == nil {
		return fileV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().File.Create().
		SetNillableProvider(r.toEntProvider(req.Data.Provider)).
		SetNillableBucketName(req.Data.BucketName).
		SetNillableFileDirectory(req.Data.FileDirectory).
		SetNillableFileGUID(req.Data.FileGuid).
		SetNillableSaveFileName(req.Data.SaveFileName).
		SetNillableFileName(req.Data.FileName).
		SetNillableExtension(req.Data.Extension).
		SetNillableSize(req.Data.Size).
		SetNillableSizeFormat(req.Data.SizeFormat).
		SetNillableLinkURL(req.Data.LinkUrl).
		SetNillableMd5(req.Data.Md5).
		SetNillableCreateBy(req.Data.CreateBy).
		SetNillableCreateTime(timeutil.StringTimeToTime(req.Data.CreateTime))

	if req.Data.CreateTime == nil {
		builder.SetCreateTime(time.Now())
	}

	if req.Data.Id != nil {
		builder.SetID(req.Data.GetId())
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("insert one data failed: %s", err.Error())
		return fileV1.ErrorInternalServerError("insert data failed")
	}

	return nil
}

func (r *FileRepo) Update(ctx context.Context, req *fileV1.UpdateFileRequest) error {
	if req == nil || req.Data == nil {
		return fileV1.ErrorBadRequest("invalid parameter")
	}

	// 如果不存在则创建
	if req.GetAllowMissing() {
		exist, err := r.IsExist(ctx, req.GetData().GetId())
		if err != nil {
			return err
		}
		if !exist {
			createReq := &fileV1.CreateFileRequest{Data: req.Data}
			createReq.Data.CreateBy = createReq.Data.UpdateBy
			createReq.Data.UpdateBy = nil
			return r.Create(ctx, createReq)
		}
	}

	if req.UpdateMask != nil {
		req.UpdateMask.Normalize()
		if !req.UpdateMask.IsValid(req.Data) {
			return fileV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().File.UpdateOneID(req.Data.GetId()).
		SetNillableProvider(r.toEntProvider(req.Data.Provider)).
		SetNillableBucketName(req.Data.BucketName).
		SetNillableFileDirectory(req.Data.FileDirectory).
		SetNillableFileGUID(req.Data.FileGuid).
		SetNillableSaveFileName(req.Data.SaveFileName).
		SetNillableFileName(req.Data.FileName).
		SetNillableExtension(req.Data.Extension).
		SetNillableSize(req.Data.Size).
		SetNillableSizeFormat(req.Data.SizeFormat).
		SetNillableLinkURL(req.Data.LinkUrl).
		SetNillableMd5(req.Data.Md5).
		//SetNillableUpdateBy(trans.Ptr(operator.UserId)).
		SetNillableUpdateTime(timeutil.StringTimeToTime(req.Data.UpdateTime))

	if req.Data.UpdateTime == nil {
		builder.SetUpdateTime(time.Now())
	}

	if req.UpdateMask != nil {
		nilPaths := fieldmaskutil.NilValuePaths(req.Data, req.GetUpdateMask().GetPaths())
		nilUpdater := entgoUpdate.BuildSetNullUpdater(nilPaths)
		if nilUpdater != nil {
			builder.Modify(nilUpdater)
		}
	}

	if err := builder.Exec(ctx); err != nil {
		r.log.Errorf("update one data failed: %s", err.Error())
		return fileV1.ErrorInternalServerError("update data failed")
	}

	return nil
}

func (r *FileRepo) Delete(ctx context.Context, req *fileV1.DeleteFileRequest) error {
	if req == nil {
		return fileV1.ErrorBadRequest("invalid parameter")
	}

	if err := r.data.db.Client().File.DeleteOneID(req.GetId()).Exec(ctx); err != nil {
		if ent.IsNotFound(err) {
			return fileV1.ErrorNotFound("file not found")
		}

		r.log.Errorf("delete one data failed: %s", err.Error())

		return fileV1.ErrorInternalServerError("delete failed")
	}

	return nil
}
