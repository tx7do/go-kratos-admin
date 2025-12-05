package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/tx7do/go-crud/entgo/mixin"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table:     "files",
			Charset:   "utf8mb4",
			Collation: "utf8mb4_bin",
		},
		entsql.WithComments(true),
		schema.Comment("文件表"),
	}
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("provider").
			Comment("OSS供应商").
			NamedValues(
				"Unknown", "UNKNOWN",
				"MinIO", "MINIO",
				"Aliyun", "ALIYUN",
				"Qiniu", "QINIU",
				"Tencent", "TENCENT",
				"AWS", "AWS",
				"Google", "GOOGLE",
				"Azure", "AZURE",
				"Baidu", "BAIDU",
				"Huawei", "HUAWEI",
				"QCloud", "QCLOUD",
				"Local", "LOCAL",
			).
			Optional().
			Nillable(),

		field.String("bucket_name").
			Comment("存储桶名称").
			Optional().
			Nillable(),

		field.String("file_directory").
			Comment("文件目录").
			Optional().
			Nillable(),

		field.String("file_guid").
			Comment("文件Guid").
			Optional().
			Nillable(),

		field.String("save_file_name").
			Comment("保存文件名").
			Optional().
			Nillable(),

		field.String("file_name").
			Comment("文件名").
			Optional().
			Nillable(),

		field.String("extension").
			Comment("文件扩展名").
			Optional().
			Nillable(),

		field.Uint64("size").
			Comment("文件字节长度").
			Optional().
			Nillable(),

		field.String("size_format").
			Comment("文件大小格式化").
			Optional().
			Nillable(),

		field.String("link_url").
			Comment("链接地址").
			Optional().
			Nillable(),

		field.String("md5").
			Comment("md5码，防止上传重复文件").
			Optional().
			Nillable(),
	}
}

// Mixin of the File.
func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AutoIncrementId{},
		mixin.TimeAt{},
		mixin.OperatorID{},
		mixin.Remark{},
		mixin.TenantID{},
	}
}
