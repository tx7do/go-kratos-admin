package data

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/copierutil"
	entgo "github.com/tx7do/go-utils/entgo/query"
	entgoUpdate "github.com/tx7do/go-utils/entgo/update"
	"github.com/tx7do/go-utils/fieldmaskutil"
	"github.com/tx7do/go-utils/mapper"
	"github.com/tx7do/go-utils/timeutil"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"

	"kratos-admin/app/admin/service/internal/data/ent"
	"kratos-admin/app/admin/service/internal/data/ent/file"

	fileV1 "kratos-admin/api/gen/go/file/service/v1"
)

type FileRepo struct {
	data *Data
	log  *log.Helper

	mapper            *mapper.CopierMapper[fileV1.File, ent.File]
	providerConverter *mapper.EnumTypeConverter[fileV1.OSSProvider, file.Provider]
}

func NewFileRepo(data *Data, logger log.Logger) *FileRepo {
	repo := &FileRepo{
		log:               log.NewHelper(log.With(logger, "module", "file/repo/admin-service")),
		data:              data,
		mapper:            mapper.NewCopierMapper[fileV1.File, ent.File](),
		providerConverter: mapper.NewEnumTypeConverter[fileV1.OSSProvider, file.Provider](fileV1.OSSProvider_name, fileV1.OSSProvider_value),
	}

	repo.init()

	return repo
}

func (r *FileRepo) init() {
	r.mapper.AppendConverters(copierutil.NewTimeStringConverterPair())
	r.mapper.AppendConverters(copierutil.NewTimeTimestamppbConverterPair())

	r.mapper.AppendConverters(r.providerConverter.NewConverterPair())
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

	entities, err := builder.All(ctx)
	if err != nil {
		r.log.Errorf("query list failed: %s", err.Error())
		return nil, fileV1.ErrorInternalServerError("query list failed")
	}

	dtos := make([]*fileV1.File, 0, len(entities))
	for _, entity := range entities {
		dto := r.mapper.ToDTO(entity)
		dtos = append(dtos, dto)
	}

	count, err := r.Count(ctx, whereSelectors)
	if err != nil {
		return nil, err
	}

	return &fileV1.ListFileResponse{
		Total: uint32(count),
		Items: dtos,
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

	entity, err := r.data.db.Client().File.Get(ctx, req.GetId())
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fileV1.ErrorFileNotFound("file not found")
		}

		r.log.Errorf("query one data failed: %s", err.Error())

		return nil, fileV1.ErrorInternalServerError("query data failed")
	}

	return r.mapper.ToDTO(entity), nil
}

func (r *FileRepo) Create(ctx context.Context, req *fileV1.CreateFileRequest) error {
	if req == nil || req.Data == nil {
		return fileV1.ErrorBadRequest("invalid parameter")
	}

	builder := r.data.db.Client().File.Create().
		SetNillableProvider(r.providerConverter.ToEntity(req.Data.Provider)).
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
		SetNillableCreateTime(timeutil.TimestamppbToTime(req.Data.CreateTime))

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
			r.log.Errorf("invalid field mask [%v]", req.UpdateMask)
			return fileV1.ErrorBadRequest("invalid field mask")
		}
		fieldmaskutil.Filter(req.GetData(), req.UpdateMask.GetPaths())
	}

	builder := r.data.db.Client().File.UpdateOneID(req.Data.GetId()).
		SetNillableProvider(r.providerConverter.ToEntity(req.Data.Provider)).
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
		SetNillableUpdateTime(timeutil.TimestamppbToTime(req.Data.UpdateTime))

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
