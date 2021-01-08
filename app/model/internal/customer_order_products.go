// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// CustomerOrderProducts is the golang structure for table customer_order_products.
type CustomerOrderProducts struct {
    Id                int         `orm:"id,primary"          json:"id"`                  //                    
    OrderId           string      `orm:"order_id"            json:"order_id"`            //                    
    ProductId         int         `orm:"product_id"          json:"product_id"`          //                    
    SkuId             int         `orm:"sku_id"              json:"sku_id"`              //                    
    Abbreviation      string      `orm:"abbreviation"        json:"abbreviation"`        //                    
    PicValueId        int         `orm:"pic_value_id"        json:"pic_value_id"`        //                    
    OriginPrice       float64     `orm:"origin_price"        json:"origin_price"`        // 原商品单价         
    UnitPrice         float64     `orm:"unit_price"          json:"unit_price"`          // 最终商品单价       
    PreferentialPrice float64     `orm:"preferential_price"  json:"preferential_price"`  // 促销活动优惠价     
    CodePrice         float64     `orm:"code_price"          json:"code_price"`          // 优惠券的优惠金额   
    Num               int         `orm:"num"                 json:"num"`                 // 购买数量           
    RefundNum         int         `orm:"refund_num"          json:"refund_num"`          // 退货数量           
    TotalPrice        float64     `orm:"total_price"         json:"total_price"`         // 单个商品总价       
    MultiPromotionId  int         `orm:"multi_promotion_id"  json:"multi_promotion_id"`  // 促销活动id         
    SinglePromotionId int         `orm:"single_promotion_id" json:"single_promotion_id"` // 单品促销ID         
    StockStatus       int         `orm:"stock_status"        json:"stock_status"`        // 1:缺货；0：配送中  
    IsGift            int         `orm:"is_gift"             json:"is_gift"`             //                    
    LockStock         uint        `orm:"lock_stock"          json:"lock_stock"`          // 锁定库存           
    DeletedAt         *gtime.Time `orm:"deleted_at"          json:"deleted_at"`          //                    
    CreatedAt         *gtime.Time `orm:"created_at"          json:"created_at"`          //                    
    UpdatedAt         *gtime.Time `orm:"updated_at"          json:"updated_at"`          //                    
    MessageId         string      `orm:"message_id"          json:"message_id"`          //                    
}