// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTable interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Table
	GetCreatedAt() *int64
	SetCreatedBy(v string) *Table
	GetCreatedBy() *string
	SetId(v string) *Table
	GetId() *string
	SetIsExternal(v bool) *Table
	GetIsExternal() *bool
	SetName(v string) *Table
	GetName() *string
	SetOwner(v string) *Table
	GetOwner() *string
	SetPath(v string) *Table
	GetPath() *string
	SetSchema(v *Schema) *Table
	GetSchema() *Schema
	SetSchemaId(v int64) *Table
	GetSchemaId() *int64
	SetUpdatedAt(v int64) *Table
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *Table
	GetUpdatedBy() *string
}

type Table struct {
	CreatedAt  *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CreatedBy  *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	IsExternal *bool   `json:"isExternal,omitempty" xml:"isExternal,omitempty"`
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	Owner      *string `json:"owner,omitempty" xml:"owner,omitempty"`
	Path       *string `json:"path,omitempty" xml:"path,omitempty"`
	Schema     *Schema `json:"schema,omitempty" xml:"schema,omitempty"`
	SchemaId   *int64  `json:"schemaId,omitempty" xml:"schemaId,omitempty"`
	UpdatedAt  *int64  `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	UpdatedBy  *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Table) String() string {
	return dara.Prettify(s)
}

func (s Table) GoString() string {
	return s.String()
}

func (s *Table) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Table) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *Table) GetId() *string {
	return s.Id
}

func (s *Table) GetIsExternal() *bool {
	return s.IsExternal
}

func (s *Table) GetName() *string {
	return s.Name
}

func (s *Table) GetOwner() *string {
	return s.Owner
}

func (s *Table) GetPath() *string {
	return s.Path
}

func (s *Table) GetSchema() *Schema {
	return s.Schema
}

func (s *Table) GetSchemaId() *int64 {
	return s.SchemaId
}

func (s *Table) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Table) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *Table) SetCreatedAt(v int64) *Table {
	s.CreatedAt = &v
	return s
}

func (s *Table) SetCreatedBy(v string) *Table {
	s.CreatedBy = &v
	return s
}

func (s *Table) SetId(v string) *Table {
	s.Id = &v
	return s
}

func (s *Table) SetIsExternal(v bool) *Table {
	s.IsExternal = &v
	return s
}

func (s *Table) SetName(v string) *Table {
	s.Name = &v
	return s
}

func (s *Table) SetOwner(v string) *Table {
	s.Owner = &v
	return s
}

func (s *Table) SetPath(v string) *Table {
	s.Path = &v
	return s
}

func (s *Table) SetSchema(v *Schema) *Table {
	s.Schema = v
	return s
}

func (s *Table) SetSchemaId(v int64) *Table {
	s.SchemaId = &v
	return s
}

func (s *Table) SetUpdatedAt(v int64) *Table {
	s.UpdatedAt = &v
	return s
}

func (s *Table) SetUpdatedBy(v string) *Table {
	s.UpdatedBy = &v
	return s
}

func (s *Table) Validate() error {
	return dara.Validate(s)
}
