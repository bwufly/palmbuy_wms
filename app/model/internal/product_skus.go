// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// ProductSkus is the golang structure for table product_skus.
type ProductSkus struct {
    Id            uint        `orm:"id,primary"     json:"id"`             // 自增ID        
    ProductId     uint        `orm:"product_id"     json:"product_id"`     // 商品ID        
    PicValueId    uint        `orm:"pic_value_id"   json:"pic_value_id"`   // 图片属性值ID  
    ValueIds      string      `orm:"value_ids"      json:"value_ids"`      //               
    SupplyPrice   float64     `orm:"supply_price"   json:"supply_price"`   // 供货价        
    PurchasePrice float64     `orm:"purchase_price" json:"purchase_price"` // 采购价        
    OriginPrice   float64     `orm:"origin_price"   json:"origin_price"`   // 原价          
    Price         float64     `orm:"price"          json:"price"`          //               
    SupplierCode  string      `orm:"supplier_code"  json:"supplier_code"`  //               
    Weight        float64     `orm:"weight"         json:"weight"`         // 重量          
    SupplierSize  string      `orm:"supplier_size"  json:"supplier_size"`  //               
    Icon          string      `orm:"icon"           json:"icon"`           //               
    CreatedAt     *gtime.Time `orm:"created_at"     json:"created_at"`     //               
    DeletedAt     *gtime.Time `orm:"deleted_at"     json:"deleted_at"`     //               
    UpdatedAt     *gtime.Time `orm:"updated_at"     json:"updated_at"`     //               
    MessageId     string      `orm:"message_id"     json:"message_id"`     //               
}