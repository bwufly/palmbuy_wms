package service

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/lxzan/hasaki"
	"github.com/tidwall/gjson"
	"strings"
	"time"
	"wms/app/dao"
	"wms/app/model"
)

// 中间件管理服务
var User = new(serviceProduct)

type serviceProduct struct{}

// 用户注册
func (s *serviceProduct) Transfer(r *model.ProductTransferReq) error {
	// 昵称为非必需参数，默认使用账号名称
	if r.Id == "" {
		return errors.New("商品ID不能为空")
	}
	ids := strings.Split(r.Id,",")
	fmt.Println(ids)
	token := s.GetToken()
	for i := 0; i < len(ids); i++ {
		productInfo := s.FindProductInfo(ids[i], token)
		s.SaveProductInfo(ids[i],productInfo)
		time.Sleep(1*time.Second)
	}
	return nil
}

func (s *serviceProduct) GetToken() string {
	resp,err := hasaki.Post("https://oms.palmpay.app/mfront/ng/api/mfront/login/0/login").Set(hasaki.Form{
		"accept": " application/json, text/plain, */*",
	}).
		Send(hasaki.Any{
			"loginName": "fei.wu@transsnet.com",
			"loginPwd":    "33a35e36d224cb623c4d58d0076dfb84",
		}).GetBody()
	if err != nil {
		fmt.Println(err)
	}
	return gjson.ParseBytes(resp).Get("data.token").String()
}

func (s *serviceProduct) FindProductInfo(id, token string) string {
	resp,err := hasaki.
		Get(fmt.Sprintf("https://oms.palmpay.app/mfront/ng/api/mfront/commodity-center/10842/commodity/retail/%s",id)).
		Set(hasaki.Form{"m_token":token}).
		Send(nil).
		GetBody()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(resp)
}

func (s *serviceProduct) SaveProductInfo(id string, productInfo string) {
	// 把相关信息保存到数据库
	r, err := dao.Products.FindOne("id = ?",id)
	if err != nil {
		fmt.Println(err)
	}
	// 商品不存在，创建 商品表，sku表，库存表
	if r == nil {
		err := g.DB().Transaction(func(tx *gdb.TX) error {
			// products
			result, err := tx.Insert("products", g.Map{
				"id": gjson.Get(productInfo,"data.commodityId").String(),
				"category_id": gjson.Get(productInfo,"data.categoryId").String(),
				"category_path": "",
				"supplier_id": "1",
				"author_id": "1",
				"stall_id": 0,
				"code": gjson.Get(productInfo,"data.codeValue").String(),
				"price": 0,
				"supply_price": 0,
				"purchase_price": 0,
				"video": "",
				"title": gjson.Get(productInfo,"data.fullName").String(),
				"en_title": gjson.Get(productInfo,"data.fullName").String(),
				"main_icon": gjson.Get(productInfo,"data.commodityPictureList.0").String(),
				"pic_type": 1,
				"sort": 99,
				"status": 1,
				"platform_product_id": "",
				"from_platform": "",
				"created_at": time.Now().String(),
				"message_id": "1",
			})
			if err != nil {
				return err
			}
			// user_detail
			id, err := result.LastInsertId()
			if err != nil {
				return err
			}
			var skus []g.Map
			var skuStocks []g.Map
			gjson.Get(productInfo,"data.skuInfoList").ForEach(func(key, value gjson.Result) bool {
				icon := ""
				valueId := value.Get("standardOptionsRespDetailMap.2.standardOptionsId").String()

				gjson.Get(productInfo,"data.commodityOtherPicturesObject.standardOptionsRespDetailList").ForEach(func(key, value gjson.Result) bool {
					if value.Get("standardOptionsId").String() == valueId {
						icon = value.Get("imageUrl").String()
					}
					return true
				})
				if icon == ""{
					icon = gjson.Get(productInfo,"data.commodityPictureList.0").String()
				}
				fmt.Println("valueId",valueId,"icon",icon)
				skus = append(skus, g.Map{
					"id":       value.Get("skuId").String(),
					"product_id":       id,
					"pic_value_id":       0,
					"value_ids":       "",
					"supply_price":       value.Get("costPrice").String(),
					"purchase_price":       value.Get("costPrice").String(),
					"origin_price":       0,
					"price":       value.Get("sellPrice").String(),
					"supplier_code":       value.Get("supplierCommodityCode").String(),
					"weight":       0,
					"supplier_size":       "",
					"icon":       icon,
					"created_at":       time.Now().String(),
				})

				skuStocks = append(skuStocks, g.Map{
					"product_id":       id,
					"sku_id":       value.Get("skuId").String(),
					"created_at":       time.Now().String(),
				})
				return true
			})
			// 批量插入sku表
			_, err = tx.BatchInsert("product_skus",skus )
			if err != nil {
				return err
			}
			// 批量插入库存表
			_, err = tx.BatchInsert("stocks",skuStocks )
			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			fmt.Println(err)
		}
	}
}


