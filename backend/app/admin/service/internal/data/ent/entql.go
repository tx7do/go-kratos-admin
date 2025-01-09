// Code generated by ent, DO NOT EDIT.

package ent

import (
	"kratos-admin/app/admin/service/internal/data/ent/dict"
	"kratos-admin/app/admin/service/internal/data/ent/menu"
	"kratos-admin/app/admin/service/internal/data/ent/organization"
	"kratos-admin/app/admin/service/internal/data/ent/position"
	"kratos-admin/app/admin/service/internal/data/ent/predicate"
	"kratos-admin/app/admin/service/internal/data/ent/role"
	"kratos-admin/app/admin/service/internal/data/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 6)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   dict.Table,
			Columns: dict.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: dict.FieldID,
			},
		},
		Type: "Dict",
		Fields: map[string]*sqlgraph.FieldSpec{
			dict.FieldCreateTime:    {Type: field.TypeTime, Column: dict.FieldCreateTime},
			dict.FieldUpdateTime:    {Type: field.TypeTime, Column: dict.FieldUpdateTime},
			dict.FieldDeleteTime:    {Type: field.TypeTime, Column: dict.FieldDeleteTime},
			dict.FieldStatus:        {Type: field.TypeEnum, Column: dict.FieldStatus},
			dict.FieldCreateBy:      {Type: field.TypeUint32, Column: dict.FieldCreateBy},
			dict.FieldUpdateBy:      {Type: field.TypeUint32, Column: dict.FieldUpdateBy},
			dict.FieldRemark:        {Type: field.TypeString, Column: dict.FieldRemark},
			dict.FieldKey:           {Type: field.TypeString, Column: dict.FieldKey},
			dict.FieldCategory:      {Type: field.TypeString, Column: dict.FieldCategory},
			dict.FieldCategoryDesc:  {Type: field.TypeString, Column: dict.FieldCategoryDesc},
			dict.FieldValue:         {Type: field.TypeString, Column: dict.FieldValue},
			dict.FieldValueDesc:     {Type: field.TypeString, Column: dict.FieldValueDesc},
			dict.FieldValueDataType: {Type: field.TypeString, Column: dict.FieldValueDataType},
			dict.FieldSortID:        {Type: field.TypeInt32, Column: dict.FieldSortID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   menu.Table,
			Columns: menu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: menu.FieldID,
			},
		},
		Type: "Menu",
		Fields: map[string]*sqlgraph.FieldSpec{
			menu.FieldStatus:     {Type: field.TypeEnum, Column: menu.FieldStatus},
			menu.FieldCreateTime: {Type: field.TypeTime, Column: menu.FieldCreateTime},
			menu.FieldUpdateTime: {Type: field.TypeTime, Column: menu.FieldUpdateTime},
			menu.FieldDeleteTime: {Type: field.TypeTime, Column: menu.FieldDeleteTime},
			menu.FieldCreateBy:   {Type: field.TypeUint32, Column: menu.FieldCreateBy},
			menu.FieldUpdateBy:   {Type: field.TypeUint32, Column: menu.FieldUpdateBy},
			menu.FieldRemark:     {Type: field.TypeString, Column: menu.FieldRemark},
			menu.FieldParentID:   {Type: field.TypeInt32, Column: menu.FieldParentID},
			menu.FieldType:       {Type: field.TypeEnum, Column: menu.FieldType},
			menu.FieldPath:       {Type: field.TypeString, Column: menu.FieldPath},
			menu.FieldRedirect:   {Type: field.TypeString, Column: menu.FieldRedirect},
			menu.FieldAlias:      {Type: field.TypeString, Column: menu.FieldAlias},
			menu.FieldName:       {Type: field.TypeString, Column: menu.FieldName},
			menu.FieldComponent:  {Type: field.TypeString, Column: menu.FieldComponent},
			menu.FieldMeta:       {Type: field.TypeJSON, Column: menu.FieldMeta},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   organization.Table,
			Columns: organization.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: organization.FieldID,
			},
		},
		Type: "Organization",
		Fields: map[string]*sqlgraph.FieldSpec{
			organization.FieldCreateTime: {Type: field.TypeTime, Column: organization.FieldCreateTime},
			organization.FieldUpdateTime: {Type: field.TypeTime, Column: organization.FieldUpdateTime},
			organization.FieldDeleteTime: {Type: field.TypeTime, Column: organization.FieldDeleteTime},
			organization.FieldStatus:     {Type: field.TypeEnum, Column: organization.FieldStatus},
			organization.FieldCreateBy:   {Type: field.TypeUint32, Column: organization.FieldCreateBy},
			organization.FieldUpdateBy:   {Type: field.TypeUint32, Column: organization.FieldUpdateBy},
			organization.FieldRemark:     {Type: field.TypeString, Column: organization.FieldRemark},
			organization.FieldName:       {Type: field.TypeString, Column: organization.FieldName},
			organization.FieldParentID:   {Type: field.TypeUint32, Column: organization.FieldParentID},
			organization.FieldSortID:     {Type: field.TypeInt32, Column: organization.FieldSortID},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   position.Table,
			Columns: position.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: position.FieldID,
			},
		},
		Type: "Position",
		Fields: map[string]*sqlgraph.FieldSpec{
			position.FieldCreateTime: {Type: field.TypeTime, Column: position.FieldCreateTime},
			position.FieldUpdateTime: {Type: field.TypeTime, Column: position.FieldUpdateTime},
			position.FieldDeleteTime: {Type: field.TypeTime, Column: position.FieldDeleteTime},
			position.FieldStatus:     {Type: field.TypeEnum, Column: position.FieldStatus},
			position.FieldCreateBy:   {Type: field.TypeUint32, Column: position.FieldCreateBy},
			position.FieldUpdateBy:   {Type: field.TypeUint32, Column: position.FieldUpdateBy},
			position.FieldRemark:     {Type: field.TypeString, Column: position.FieldRemark},
			position.FieldName:       {Type: field.TypeString, Column: position.FieldName},
			position.FieldCode:       {Type: field.TypeString, Column: position.FieldCode},
			position.FieldParentID:   {Type: field.TypeUint32, Column: position.FieldParentID},
			position.FieldSortID:     {Type: field.TypeInt32, Column: position.FieldSortID},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: role.FieldID,
			},
		},
		Type: "Role",
		Fields: map[string]*sqlgraph.FieldSpec{
			role.FieldCreateTime: {Type: field.TypeTime, Column: role.FieldCreateTime},
			role.FieldUpdateTime: {Type: field.TypeTime, Column: role.FieldUpdateTime},
			role.FieldDeleteTime: {Type: field.TypeTime, Column: role.FieldDeleteTime},
			role.FieldStatus:     {Type: field.TypeEnum, Column: role.FieldStatus},
			role.FieldCreateBy:   {Type: field.TypeUint32, Column: role.FieldCreateBy},
			role.FieldUpdateBy:   {Type: field.TypeUint32, Column: role.FieldUpdateBy},
			role.FieldRemark:     {Type: field.TypeString, Column: role.FieldRemark},
			role.FieldName:       {Type: field.TypeString, Column: role.FieldName},
			role.FieldCode:       {Type: field.TypeString, Column: role.FieldCode},
			role.FieldParentID:   {Type: field.TypeUint32, Column: role.FieldParentID},
			role.FieldSortID:     {Type: field.TypeInt32, Column: role.FieldSortID},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreateBy:      {Type: field.TypeUint32, Column: user.FieldCreateBy},
			user.FieldUpdateBy:      {Type: field.TypeUint32, Column: user.FieldUpdateBy},
			user.FieldCreateTime:    {Type: field.TypeTime, Column: user.FieldCreateTime},
			user.FieldUpdateTime:    {Type: field.TypeTime, Column: user.FieldUpdateTime},
			user.FieldDeleteTime:    {Type: field.TypeTime, Column: user.FieldDeleteTime},
			user.FieldRemark:        {Type: field.TypeString, Column: user.FieldRemark},
			user.FieldStatus:        {Type: field.TypeEnum, Column: user.FieldStatus},
			user.FieldUsername:      {Type: field.TypeString, Column: user.FieldUsername},
			user.FieldPassword:      {Type: field.TypeString, Column: user.FieldPassword},
			user.FieldNickName:      {Type: field.TypeString, Column: user.FieldNickName},
			user.FieldRealName:      {Type: field.TypeString, Column: user.FieldRealName},
			user.FieldEmail:         {Type: field.TypeString, Column: user.FieldEmail},
			user.FieldMobile:        {Type: field.TypeString, Column: user.FieldMobile},
			user.FieldTelephone:     {Type: field.TypeString, Column: user.FieldTelephone},
			user.FieldAvatar:        {Type: field.TypeString, Column: user.FieldAvatar},
			user.FieldGender:        {Type: field.TypeEnum, Column: user.FieldGender},
			user.FieldAddress:       {Type: field.TypeString, Column: user.FieldAddress},
			user.FieldRegion:        {Type: field.TypeString, Column: user.FieldRegion},
			user.FieldDescription:   {Type: field.TypeString, Column: user.FieldDescription},
			user.FieldAuthority:     {Type: field.TypeEnum, Column: user.FieldAuthority},
			user.FieldLastLoginTime: {Type: field.TypeInt64, Column: user.FieldLastLoginTime},
			user.FieldLastLoginIP:   {Type: field.TypeString, Column: user.FieldLastLoginIP},
			user.FieldRoleID:        {Type: field.TypeUint32, Column: user.FieldRoleID},
			user.FieldOrgID:         {Type: field.TypeUint32, Column: user.FieldOrgID},
			user.FieldPositionID:    {Type: field.TypeUint32, Column: user.FieldPositionID},
			user.FieldWorkID:        {Type: field.TypeUint32, Column: user.FieldWorkID},
		},
	}
	graph.MustAddE(
		"parent",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
		},
		"Menu",
		"Menu",
	)
	graph.MustAddE(
		"children",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
		},
		"Menu",
		"Menu",
	)
	graph.MustAddE(
		"parent",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
			Bidi:    false,
		},
		"Organization",
		"Organization",
	)
	graph.MustAddE(
		"children",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
		},
		"Organization",
		"Organization",
	)
	graph.MustAddE(
		"parent",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   position.ParentTable,
			Columns: []string{position.ParentColumn},
			Bidi:    false,
		},
		"Position",
		"Position",
	)
	graph.MustAddE(
		"children",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.ChildrenTable,
			Columns: []string{position.ChildrenColumn},
			Bidi:    false,
		},
		"Position",
		"Position",
	)
	graph.MustAddE(
		"parent",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   role.ParentTable,
			Columns: []string{role.ParentColumn},
			Bidi:    false,
		},
		"Role",
		"Role",
	)
	graph.MustAddE(
		"children",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.ChildrenTable,
			Columns: []string{role.ChildrenColumn},
			Bidi:    false,
		},
		"Role",
		"Role",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (dq *DictQuery) addPredicate(pred func(s *sql.Selector)) {
	dq.predicates = append(dq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the DictQuery builder.
func (dq *DictQuery) Filter() *DictFilter {
	return &DictFilter{config: dq.config, predicateAdder: dq}
}

// addPredicate implements the predicateAdder interface.
func (m *DictMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the DictMutation builder.
func (m *DictMutation) Filter() *DictFilter {
	return &DictFilter{config: m.config, predicateAdder: m}
}

// DictFilter provides a generic filtering capability at runtime for DictQuery.
type DictFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *DictFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *DictFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(dict.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *DictFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(dict.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *DictFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(dict.FieldUpdateTime))
}

// WhereDeleteTime applies the entql time.Time predicate on the delete_time field.
func (f *DictFilter) WhereDeleteTime(p entql.TimeP) {
	f.Where(p.Field(dict.FieldDeleteTime))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *DictFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(dict.FieldStatus))
}

// WhereCreateBy applies the entql uint32 predicate on the create_by field.
func (f *DictFilter) WhereCreateBy(p entql.Uint32P) {
	f.Where(p.Field(dict.FieldCreateBy))
}

// WhereUpdateBy applies the entql uint32 predicate on the update_by field.
func (f *DictFilter) WhereUpdateBy(p entql.Uint32P) {
	f.Where(p.Field(dict.FieldUpdateBy))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *DictFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(dict.FieldRemark))
}

// WhereKey applies the entql string predicate on the key field.
func (f *DictFilter) WhereKey(p entql.StringP) {
	f.Where(p.Field(dict.FieldKey))
}

// WhereCategory applies the entql string predicate on the category field.
func (f *DictFilter) WhereCategory(p entql.StringP) {
	f.Where(p.Field(dict.FieldCategory))
}

// WhereCategoryDesc applies the entql string predicate on the category_desc field.
func (f *DictFilter) WhereCategoryDesc(p entql.StringP) {
	f.Where(p.Field(dict.FieldCategoryDesc))
}

// WhereValue applies the entql string predicate on the value field.
func (f *DictFilter) WhereValue(p entql.StringP) {
	f.Where(p.Field(dict.FieldValue))
}

// WhereValueDesc applies the entql string predicate on the value_desc field.
func (f *DictFilter) WhereValueDesc(p entql.StringP) {
	f.Where(p.Field(dict.FieldValueDesc))
}

// WhereValueDataType applies the entql string predicate on the value_data_type field.
func (f *DictFilter) WhereValueDataType(p entql.StringP) {
	f.Where(p.Field(dict.FieldValueDataType))
}

// WhereSortID applies the entql int32 predicate on the sort_id field.
func (f *DictFilter) WhereSortID(p entql.Int32P) {
	f.Where(p.Field(dict.FieldSortID))
}

// addPredicate implements the predicateAdder interface.
func (mq *MenuQuery) addPredicate(pred func(s *sql.Selector)) {
	mq.predicates = append(mq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the MenuQuery builder.
func (mq *MenuQuery) Filter() *MenuFilter {
	return &MenuFilter{config: mq.config, predicateAdder: mq}
}

// addPredicate implements the predicateAdder interface.
func (m *MenuMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the MenuMutation builder.
func (m *MenuMutation) Filter() *MenuFilter {
	return &MenuFilter{config: m.config, predicateAdder: m}
}

// MenuFilter provides a generic filtering capability at runtime for MenuQuery.
type MenuFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *MenuFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int32 predicate on the id field.
func (f *MenuFilter) WhereID(p entql.Int32P) {
	f.Where(p.Field(menu.FieldID))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *MenuFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(menu.FieldStatus))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *MenuFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(menu.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *MenuFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(menu.FieldUpdateTime))
}

// WhereDeleteTime applies the entql time.Time predicate on the delete_time field.
func (f *MenuFilter) WhereDeleteTime(p entql.TimeP) {
	f.Where(p.Field(menu.FieldDeleteTime))
}

// WhereCreateBy applies the entql uint32 predicate on the create_by field.
func (f *MenuFilter) WhereCreateBy(p entql.Uint32P) {
	f.Where(p.Field(menu.FieldCreateBy))
}

// WhereUpdateBy applies the entql uint32 predicate on the update_by field.
func (f *MenuFilter) WhereUpdateBy(p entql.Uint32P) {
	f.Where(p.Field(menu.FieldUpdateBy))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *MenuFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(menu.FieldRemark))
}

// WhereParentID applies the entql int32 predicate on the parent_id field.
func (f *MenuFilter) WhereParentID(p entql.Int32P) {
	f.Where(p.Field(menu.FieldParentID))
}

// WhereType applies the entql string predicate on the type field.
func (f *MenuFilter) WhereType(p entql.StringP) {
	f.Where(p.Field(menu.FieldType))
}

// WherePath applies the entql string predicate on the path field.
func (f *MenuFilter) WherePath(p entql.StringP) {
	f.Where(p.Field(menu.FieldPath))
}

// WhereRedirect applies the entql string predicate on the redirect field.
func (f *MenuFilter) WhereRedirect(p entql.StringP) {
	f.Where(p.Field(menu.FieldRedirect))
}

// WhereAlias applies the entql string predicate on the alias field.
func (f *MenuFilter) WhereAlias(p entql.StringP) {
	f.Where(p.Field(menu.FieldAlias))
}

// WhereName applies the entql string predicate on the name field.
func (f *MenuFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(menu.FieldName))
}

// WhereComponent applies the entql string predicate on the component field.
func (f *MenuFilter) WhereComponent(p entql.StringP) {
	f.Where(p.Field(menu.FieldComponent))
}

// WhereMeta applies the entql json.RawMessage predicate on the meta field.
func (f *MenuFilter) WhereMeta(p entql.BytesP) {
	f.Where(p.Field(menu.FieldMeta))
}

// WhereHasParent applies a predicate to check if query has an edge parent.
func (f *MenuFilter) WhereHasParent() {
	f.Where(entql.HasEdge("parent"))
}

// WhereHasParentWith applies a predicate to check if query has an edge parent with a given conditions (other predicates).
func (f *MenuFilter) WhereHasParentWith(preds ...predicate.Menu) {
	f.Where(entql.HasEdgeWith("parent", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasChildren applies a predicate to check if query has an edge children.
func (f *MenuFilter) WhereHasChildren() {
	f.Where(entql.HasEdge("children"))
}

// WhereHasChildrenWith applies a predicate to check if query has an edge children with a given conditions (other predicates).
func (f *MenuFilter) WhereHasChildrenWith(preds ...predicate.Menu) {
	f.Where(entql.HasEdgeWith("children", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (oq *OrganizationQuery) addPredicate(pred func(s *sql.Selector)) {
	oq.predicates = append(oq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the OrganizationQuery builder.
func (oq *OrganizationQuery) Filter() *OrganizationFilter {
	return &OrganizationFilter{config: oq.config, predicateAdder: oq}
}

// addPredicate implements the predicateAdder interface.
func (m *OrganizationMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the OrganizationMutation builder.
func (m *OrganizationMutation) Filter() *OrganizationFilter {
	return &OrganizationFilter{config: m.config, predicateAdder: m}
}

// OrganizationFilter provides a generic filtering capability at runtime for OrganizationQuery.
type OrganizationFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *OrganizationFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *OrganizationFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(organization.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *OrganizationFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(organization.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *OrganizationFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(organization.FieldUpdateTime))
}

// WhereDeleteTime applies the entql time.Time predicate on the delete_time field.
func (f *OrganizationFilter) WhereDeleteTime(p entql.TimeP) {
	f.Where(p.Field(organization.FieldDeleteTime))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *OrganizationFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(organization.FieldStatus))
}

// WhereCreateBy applies the entql uint32 predicate on the create_by field.
func (f *OrganizationFilter) WhereCreateBy(p entql.Uint32P) {
	f.Where(p.Field(organization.FieldCreateBy))
}

// WhereUpdateBy applies the entql uint32 predicate on the update_by field.
func (f *OrganizationFilter) WhereUpdateBy(p entql.Uint32P) {
	f.Where(p.Field(organization.FieldUpdateBy))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *OrganizationFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(organization.FieldRemark))
}

// WhereName applies the entql string predicate on the name field.
func (f *OrganizationFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(organization.FieldName))
}

// WhereParentID applies the entql uint32 predicate on the parent_id field.
func (f *OrganizationFilter) WhereParentID(p entql.Uint32P) {
	f.Where(p.Field(organization.FieldParentID))
}

// WhereSortID applies the entql int32 predicate on the sort_id field.
func (f *OrganizationFilter) WhereSortID(p entql.Int32P) {
	f.Where(p.Field(organization.FieldSortID))
}

// WhereHasParent applies a predicate to check if query has an edge parent.
func (f *OrganizationFilter) WhereHasParent() {
	f.Where(entql.HasEdge("parent"))
}

// WhereHasParentWith applies a predicate to check if query has an edge parent with a given conditions (other predicates).
func (f *OrganizationFilter) WhereHasParentWith(preds ...predicate.Organization) {
	f.Where(entql.HasEdgeWith("parent", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasChildren applies a predicate to check if query has an edge children.
func (f *OrganizationFilter) WhereHasChildren() {
	f.Where(entql.HasEdge("children"))
}

// WhereHasChildrenWith applies a predicate to check if query has an edge children with a given conditions (other predicates).
func (f *OrganizationFilter) WhereHasChildrenWith(preds ...predicate.Organization) {
	f.Where(entql.HasEdgeWith("children", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (pq *PositionQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PositionQuery builder.
func (pq *PositionQuery) Filter() *PositionFilter {
	return &PositionFilter{config: pq.config, predicateAdder: pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PositionMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PositionMutation builder.
func (m *PositionMutation) Filter() *PositionFilter {
	return &PositionFilter{config: m.config, predicateAdder: m}
}

// PositionFilter provides a generic filtering capability at runtime for PositionQuery.
type PositionFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PositionFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *PositionFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(position.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *PositionFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(position.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *PositionFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(position.FieldUpdateTime))
}

// WhereDeleteTime applies the entql time.Time predicate on the delete_time field.
func (f *PositionFilter) WhereDeleteTime(p entql.TimeP) {
	f.Where(p.Field(position.FieldDeleteTime))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *PositionFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(position.FieldStatus))
}

// WhereCreateBy applies the entql uint32 predicate on the create_by field.
func (f *PositionFilter) WhereCreateBy(p entql.Uint32P) {
	f.Where(p.Field(position.FieldCreateBy))
}

// WhereUpdateBy applies the entql uint32 predicate on the update_by field.
func (f *PositionFilter) WhereUpdateBy(p entql.Uint32P) {
	f.Where(p.Field(position.FieldUpdateBy))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *PositionFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(position.FieldRemark))
}

// WhereName applies the entql string predicate on the name field.
func (f *PositionFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(position.FieldName))
}

// WhereCode applies the entql string predicate on the code field.
func (f *PositionFilter) WhereCode(p entql.StringP) {
	f.Where(p.Field(position.FieldCode))
}

// WhereParentID applies the entql uint32 predicate on the parent_id field.
func (f *PositionFilter) WhereParentID(p entql.Uint32P) {
	f.Where(p.Field(position.FieldParentID))
}

// WhereSortID applies the entql int32 predicate on the sort_id field.
func (f *PositionFilter) WhereSortID(p entql.Int32P) {
	f.Where(p.Field(position.FieldSortID))
}

// WhereHasParent applies a predicate to check if query has an edge parent.
func (f *PositionFilter) WhereHasParent() {
	f.Where(entql.HasEdge("parent"))
}

// WhereHasParentWith applies a predicate to check if query has an edge parent with a given conditions (other predicates).
func (f *PositionFilter) WhereHasParentWith(preds ...predicate.Position) {
	f.Where(entql.HasEdgeWith("parent", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasChildren applies a predicate to check if query has an edge children.
func (f *PositionFilter) WhereHasChildren() {
	f.Where(entql.HasEdge("children"))
}

// WhereHasChildrenWith applies a predicate to check if query has an edge children with a given conditions (other predicates).
func (f *PositionFilter) WhereHasChildrenWith(preds ...predicate.Position) {
	f.Where(entql.HasEdgeWith("children", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (rq *RoleQuery) addPredicate(pred func(s *sql.Selector)) {
	rq.predicates = append(rq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the RoleQuery builder.
func (rq *RoleQuery) Filter() *RoleFilter {
	return &RoleFilter{config: rq.config, predicateAdder: rq}
}

// addPredicate implements the predicateAdder interface.
func (m *RoleMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the RoleMutation builder.
func (m *RoleMutation) Filter() *RoleFilter {
	return &RoleFilter{config: m.config, predicateAdder: m}
}

// RoleFilter provides a generic filtering capability at runtime for RoleQuery.
type RoleFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *RoleFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *RoleFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(role.FieldID))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *RoleFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(role.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *RoleFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(role.FieldUpdateTime))
}

// WhereDeleteTime applies the entql time.Time predicate on the delete_time field.
func (f *RoleFilter) WhereDeleteTime(p entql.TimeP) {
	f.Where(p.Field(role.FieldDeleteTime))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *RoleFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(role.FieldStatus))
}

// WhereCreateBy applies the entql uint32 predicate on the create_by field.
func (f *RoleFilter) WhereCreateBy(p entql.Uint32P) {
	f.Where(p.Field(role.FieldCreateBy))
}

// WhereUpdateBy applies the entql uint32 predicate on the update_by field.
func (f *RoleFilter) WhereUpdateBy(p entql.Uint32P) {
	f.Where(p.Field(role.FieldUpdateBy))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *RoleFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(role.FieldRemark))
}

// WhereName applies the entql string predicate on the name field.
func (f *RoleFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(role.FieldName))
}

// WhereCode applies the entql string predicate on the code field.
func (f *RoleFilter) WhereCode(p entql.StringP) {
	f.Where(p.Field(role.FieldCode))
}

// WhereParentID applies the entql uint32 predicate on the parent_id field.
func (f *RoleFilter) WhereParentID(p entql.Uint32P) {
	f.Where(p.Field(role.FieldParentID))
}

// WhereSortID applies the entql int32 predicate on the sort_id field.
func (f *RoleFilter) WhereSortID(p entql.Int32P) {
	f.Where(p.Field(role.FieldSortID))
}

// WhereHasParent applies a predicate to check if query has an edge parent.
func (f *RoleFilter) WhereHasParent() {
	f.Where(entql.HasEdge("parent"))
}

// WhereHasParentWith applies a predicate to check if query has an edge parent with a given conditions (other predicates).
func (f *RoleFilter) WhereHasParentWith(preds ...predicate.Role) {
	f.Where(entql.HasEdgeWith("parent", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasChildren applies a predicate to check if query has an edge children.
func (f *RoleFilter) WhereHasChildren() {
	f.Where(entql.HasEdge("children"))
}

// WhereHasChildrenWith applies a predicate to check if query has an edge children with a given conditions (other predicates).
func (f *RoleFilter) WhereHasChildrenWith(preds ...predicate.Role) {
	f.Where(entql.HasEdgeWith("children", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *UserFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreateBy applies the entql uint32 predicate on the create_by field.
func (f *UserFilter) WhereCreateBy(p entql.Uint32P) {
	f.Where(p.Field(user.FieldCreateBy))
}

// WhereUpdateBy applies the entql uint32 predicate on the update_by field.
func (f *UserFilter) WhereUpdateBy(p entql.Uint32P) {
	f.Where(p.Field(user.FieldUpdateBy))
}

// WhereCreateTime applies the entql time.Time predicate on the create_time field.
func (f *UserFilter) WhereCreateTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldCreateTime))
}

// WhereUpdateTime applies the entql time.Time predicate on the update_time field.
func (f *UserFilter) WhereUpdateTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldUpdateTime))
}

// WhereDeleteTime applies the entql time.Time predicate on the delete_time field.
func (f *UserFilter) WhereDeleteTime(p entql.TimeP) {
	f.Where(p.Field(user.FieldDeleteTime))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *UserFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(user.FieldRemark))
}

// WhereStatus applies the entql string predicate on the status field.
func (f *UserFilter) WhereStatus(p entql.StringP) {
	f.Where(p.Field(user.FieldStatus))
}

// WhereUsername applies the entql string predicate on the username field.
func (f *UserFilter) WhereUsername(p entql.StringP) {
	f.Where(p.Field(user.FieldUsername))
}

// WherePassword applies the entql string predicate on the password field.
func (f *UserFilter) WherePassword(p entql.StringP) {
	f.Where(p.Field(user.FieldPassword))
}

// WhereNickName applies the entql string predicate on the nick_name field.
func (f *UserFilter) WhereNickName(p entql.StringP) {
	f.Where(p.Field(user.FieldNickName))
}

// WhereRealName applies the entql string predicate on the real_name field.
func (f *UserFilter) WhereRealName(p entql.StringP) {
	f.Where(p.Field(user.FieldRealName))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *UserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(user.FieldEmail))
}

// WhereMobile applies the entql string predicate on the mobile field.
func (f *UserFilter) WhereMobile(p entql.StringP) {
	f.Where(p.Field(user.FieldMobile))
}

// WhereTelephone applies the entql string predicate on the telephone field.
func (f *UserFilter) WhereTelephone(p entql.StringP) {
	f.Where(p.Field(user.FieldTelephone))
}

// WhereAvatar applies the entql string predicate on the avatar field.
func (f *UserFilter) WhereAvatar(p entql.StringP) {
	f.Where(p.Field(user.FieldAvatar))
}

// WhereGender applies the entql string predicate on the gender field.
func (f *UserFilter) WhereGender(p entql.StringP) {
	f.Where(p.Field(user.FieldGender))
}

// WhereAddress applies the entql string predicate on the address field.
func (f *UserFilter) WhereAddress(p entql.StringP) {
	f.Where(p.Field(user.FieldAddress))
}

// WhereRegion applies the entql string predicate on the region field.
func (f *UserFilter) WhereRegion(p entql.StringP) {
	f.Where(p.Field(user.FieldRegion))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *UserFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(user.FieldDescription))
}

// WhereAuthority applies the entql string predicate on the authority field.
func (f *UserFilter) WhereAuthority(p entql.StringP) {
	f.Where(p.Field(user.FieldAuthority))
}

// WhereLastLoginTime applies the entql int64 predicate on the last_login_time field.
func (f *UserFilter) WhereLastLoginTime(p entql.Int64P) {
	f.Where(p.Field(user.FieldLastLoginTime))
}

// WhereLastLoginIP applies the entql string predicate on the last_login_ip field.
func (f *UserFilter) WhereLastLoginIP(p entql.StringP) {
	f.Where(p.Field(user.FieldLastLoginIP))
}

// WhereRoleID applies the entql uint32 predicate on the role_id field.
func (f *UserFilter) WhereRoleID(p entql.Uint32P) {
	f.Where(p.Field(user.FieldRoleID))
}

// WhereOrgID applies the entql uint32 predicate on the org_id field.
func (f *UserFilter) WhereOrgID(p entql.Uint32P) {
	f.Where(p.Field(user.FieldOrgID))
}

// WherePositionID applies the entql uint32 predicate on the position_id field.
func (f *UserFilter) WherePositionID(p entql.Uint32P) {
	f.Where(p.Field(user.FieldPositionID))
}

// WhereWorkID applies the entql uint32 predicate on the work_id field.
func (f *UserFilter) WhereWorkID(p entql.Uint32P) {
	f.Where(p.Field(user.FieldWorkID))
}
