// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package products

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table products.
type Entity struct {
    Id                uint        `orm:"id,primary"          json:"id"`                  // 自增ID                     
    CategoryId        uint        `orm:"category_id"         json:"category_id"`         // 商品所属类目               
    CategoryPath      string      `orm:"category_path"       json:"category_path"`       // 商品所属类目路径           
    SupplierId        int         `orm:"supplier_id"         json:"supplier_id"`         // 属所门店ID                 
    AuthorId          int         `orm:"author_id"           json:"author_id"`           // 所属用户ID                 
    StallId           int         `orm:"stall_id"            json:"stall_id"`            // 档口ID                     
    Code              string      `orm:"code"                json:"code"`                // 货号                       
    Price             float64     `orm:"price"               json:"price"`               // 售价                       
    SupplyPrice       float64     `orm:"supply_price"        json:"supply_price"`        // 供货价                     
    PurchasePrice     float64     `orm:"purchase_price"      json:"purchase_price"`      // 采购价                     
    Video             string      `orm:"video"               json:"video"`               // 视频                       
    Title             string      `orm:"title"               json:"title"`               // 商品名称                   
    EnTitle           string      `orm:"en_title"            json:"en_title"`            //                            
    MainIcon          string      `orm:"main_icon"           json:"main_icon"`           // 商品主图                   
    PicType           int         `orm:"pic_type"            json:"pic_type"`            // 图片类型1-正方形 2-长方形  
    Sort              int         `orm:"sort"                json:"sort"`                // 商品排序                   
    Status            int         `orm:"status"              json:"status"`              // 1-上架 2-下架 3-待上架     
    PlatformProductId string      `orm:"platform_product_id" json:"platform_product_id"` //                            
    FromPlatform      string      `orm:"from_platform"       json:"from_platform"`       //                            
    CreatedAt         *gtime.Time `orm:"created_at"          json:"created_at"`          // 创建时间                   
    ShelfAt           *gtime.Time `orm:"shelf_at"            json:"shelf_at"`            // 上架时间                   
    DeletedAt         *gtime.Time `orm:"deleted_at"          json:"deleted_at"`          //                            
    UpdatedAt         *gtime.Time `orm:"updated_at"          json:"updated_at"`          // 修改时间                   
    MessageId         string      `orm:"message_id"          json:"message_id"`          //                            
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
// Deprecated.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
// Deprecated.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
// Deprecated.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
// Deprecated.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
// Deprecated.
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
// Deprecated.
func (r *Entity) Update() (result sql.Result, err error) {
	where, args, err := gdb.GetWhereConditionOfStruct(r)
	if err != nil {
		return nil, err
	}
	return Model.Data(r).Where(where, args).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
// Deprecated.
func (r *Entity) Delete() (result sql.Result, err error) {
	where, args, err := gdb.GetWhereConditionOfStruct(r)
	if err != nil {
		return nil, err
	}
	return Model.Where(where, args).Delete()
}