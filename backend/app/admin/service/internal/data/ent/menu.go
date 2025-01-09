// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	servicev1 "kratos-admin/api/gen/go/system/service/v1"
	"kratos-admin/app/admin/service/internal/data/ent/menu"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Menu is the model entity for the Menu schema.
type Menu struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID int32 `json:"id,omitempty"`
	// 状态
	Status *menu.Status `json:"status,omitempty"`
	// 创建时间
	CreateTime *time.Time `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// 删除时间
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// 创建者ID
	CreateBy *uint32 `json:"create_by,omitempty"`
	// 更新者ID
	UpdateBy *uint32 `json:"update_by,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty"`
	// 上一层菜单ID
	ParentID *int32 `json:"parent_id,omitempty"`
	// 菜单类型 FOLDER: 目录 MENU: 菜单 BUTTON: 按钮
	Type *menu.Type `json:"type,omitempty"`
	// 路径,当其类型为'按钮'的时候对应的数据操作名,例如:/user.service.v1.UserService/Login
	Path *string `json:"path,omitempty"`
	// 重定向地址
	Redirect *string `json:"redirect,omitempty"`
	// 路由别名
	Alias *string `json:"alias,omitempty"`
	// 路由命名，然后我们可以使用 name 而不是 path 来传递 to 属性给 <router-link>。
	Name *string `json:"name,omitempty"`
	// 前端页面组件
	Component *string `json:"component,omitempty"`
	// 前端页面组件
	Meta *servicev1.RouteMeta `json:"meta,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuQuery when eager-loading is set.
	Edges        MenuEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MenuEdges holds the relations/edges for other nodes in the graph.
type MenuEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Menu `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Menu `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) ParentOrErr() (*Menu, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: menu.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) ChildrenOrErr() ([]*Menu, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menu) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case menu.FieldMeta:
			values[i] = new([]byte)
		case menu.FieldID, menu.FieldCreateBy, menu.FieldUpdateBy, menu.FieldParentID:
			values[i] = new(sql.NullInt64)
		case menu.FieldStatus, menu.FieldRemark, menu.FieldType, menu.FieldPath, menu.FieldRedirect, menu.FieldAlias, menu.FieldName, menu.FieldComponent:
			values[i] = new(sql.NullString)
		case menu.FieldCreateTime, menu.FieldUpdateTime, menu.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menu fields.
func (m *Menu) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case menu.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int32(value.Int64)
		case menu.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				m.Status = new(menu.Status)
				*m.Status = menu.Status(value.String)
			}
		case menu.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				m.CreateTime = new(time.Time)
				*m.CreateTime = value.Time
			}
		case menu.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				m.UpdateTime = new(time.Time)
				*m.UpdateTime = value.Time
			}
		case menu.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				m.DeleteTime = new(time.Time)
				*m.DeleteTime = value.Time
			}
		case menu.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_by", values[i])
			} else if value.Valid {
				m.CreateBy = new(uint32)
				*m.CreateBy = uint32(value.Int64)
			}
		case menu.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_by", values[i])
			} else if value.Valid {
				m.UpdateBy = new(uint32)
				*m.UpdateBy = uint32(value.Int64)
			}
		case menu.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				m.Remark = new(string)
				*m.Remark = value.String
			}
		case menu.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				m.ParentID = new(int32)
				*m.ParentID = int32(value.Int64)
			}
		case menu.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				m.Type = new(menu.Type)
				*m.Type = menu.Type(value.String)
			}
		case menu.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				m.Path = new(string)
				*m.Path = value.String
			}
		case menu.FieldRedirect:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field redirect", values[i])
			} else if value.Valid {
				m.Redirect = new(string)
				*m.Redirect = value.String
			}
		case menu.FieldAlias:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alias", values[i])
			} else if value.Valid {
				m.Alias = new(string)
				*m.Alias = value.String
			}
		case menu.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = new(string)
				*m.Name = value.String
			}
		case menu.FieldComponent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field component", values[i])
			} else if value.Valid {
				m.Component = new(string)
				*m.Component = value.String
			}
		case menu.FieldMeta:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field meta", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.Meta); err != nil {
					return fmt.Errorf("unmarshal field meta: %w", err)
				}
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Menu.
// This includes values selected through modifiers, order, etc.
func (m *Menu) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the Menu entity.
func (m *Menu) QueryParent() *MenuQuery {
	return NewMenuClient(m.config).QueryParent(m)
}

// QueryChildren queries the "children" edge of the Menu entity.
func (m *Menu) QueryChildren() *MenuQuery {
	return NewMenuClient(m.config).QueryChildren(m)
}

// Update returns a builder for updating this Menu.
// Note that you need to call Menu.Unwrap() before calling this method if this Menu
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menu) Update() *MenuUpdateOne {
	return NewMenuClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Menu entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Menu) Unwrap() *Menu {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menu is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menu) String() string {
	var builder strings.Builder
	builder.WriteString("Menu(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	if v := m.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := m.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := m.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := m.CreateBy; v != nil {
		builder.WriteString("create_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.UpdateBy; v != nil {
		builder.WriteString("update_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.Remark; v != nil {
		builder.WriteString("remark=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := m.ParentID; v != nil {
		builder.WriteString("parent_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.Type; v != nil {
		builder.WriteString("type=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := m.Path; v != nil {
		builder.WriteString("path=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := m.Redirect; v != nil {
		builder.WriteString("redirect=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := m.Alias; v != nil {
		builder.WriteString("alias=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := m.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := m.Component; v != nil {
		builder.WriteString("component=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("meta=")
	builder.WriteString(fmt.Sprintf("%v", m.Meta))
	builder.WriteByte(')')
	return builder.String()
}

// Menus is a parsable slice of Menu.
type Menus []*Menu
