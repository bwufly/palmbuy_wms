package order

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"strconv"
)

// UploadShow shows uploading simgle file page.
func UploadShow(r *ghttp.Request) {
	r.Response.Write(`
    <html>
    <head>
        <title>GF Upload File Demo</title>
    </head>
        <body>
            <form enctype="multipart/form-data" action="/order/upload" method="post">
                <input type="file" name="upload-file" />
                <input type="submit" value="upload" />
            </form>
        </body>
    </html>
    `)
}

// Upload uploads files to /tmp .
func Upload(r *ghttp.Request) {

	uploadFile := r.GetUploadFile("upload-file")
	fileName, err := uploadFile.Save(gfile.TempDir())
	if err != nil {
		r.Response.Write(err)
		return
	}
	filePath := gfile.TempDir()+"\\"+fileName
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := f.GetRows("SheetJS")
	if err != nil {
		log.Fatal(err)
	}
	orderInfos,orderSkuInfos,orderAddressInfos := extractOrderInfo(rows)
	err = g.DB().Transaction(func(tx *gdb.TX) error {
		_, err := tx.BatchInsert("customer_orders", orderInfos,len(orderInfos))
		if err != nil {
			return err
		}
		_, err = tx.BatchInsert("customer_order_products", orderSkuInfos,len(orderSkuInfos))
		if err != nil {
			return err
		}
		_, err = tx.BatchInsert("customer_order_addresses", orderAddressInfos,len(orderAddressInfos))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		r.Response.Write("上传失败")
	} else {
		r.Response.Write("上传成功")
	}

}

func extractOrderInfo(rows [][]string) (orderInfos []map[string]interface{},orderSkuInfos []map[string]interface{},orderAddressInfos []map[string]interface{}) {
	lastOrderId := ""
	OrderInfosMap := make(map[string]map[string]interface{})
	orderInfo := make(map[string]interface{})
	for _, row := range rows {
		// 判断订单状态，没待发货状态剔除
		if row[1] != "Waitting to delivery" {
			fmt.Println("aaa")
			continue
		}
		orderSkuInfo := make(map[string]interface{})

		totalNum, _ := strconv.Atoi(row[3])
		if row[0] == lastOrderId {
			orderInfo["total_num"] = orderInfo["total_num"].(int) + totalNum
		} else {
			orderInfo = initOrder(row, totalNum, orderInfo)
		}
		orderSkuInfo = initOrderSku(row)
		orderAddressInfo := initOrderAddress(row)
		fmt.Println(row[0], totalNum, orderInfo["total_num"])
		lastOrderId = row[0]
		fmt.Println()
		tmpOrderInfo := transOrderInfo(orderInfo)
		OrderInfosMap[lastOrderId] = tmpOrderInfo

		orderSkuInfos = append(orderSkuInfos, orderSkuInfo)
		orderAddressInfos = append(orderAddressInfos, orderAddressInfo)

	}
	for _, orderInfo := range OrderInfosMap {
		orderInfos = append(orderInfos, orderInfo)
	}
	return
}

func initOrderAddress(row []string) map[string]interface{} {
	orderSkuInfo := make(map[string]interface{})
	orderSkuInfo["order_id"] = row[0]
	orderSkuInfo["first_name"] = row[6]
	orderSkuInfo["last_name"] = ""
	orderSkuInfo["country"] = "Nigeria"
	orderSkuInfo["country_code"] = "AU"
	orderSkuInfo["state"] = row[11]
	orderSkuInfo["city"] = row[10]
	orderSkuInfo["street"] = row[8]+" " + row[7]
	orderSkuInfo["zipcode"] = row[11]
	orderSkuInfo["phone"] = row[5]
	orderSkuInfo["email"] = row[7]
	orderSkuInfo["created_at"] = row[15]

	return orderSkuInfo
}


func initOrderSku(row []string) map[string]interface{} {
	orderSkuInfo := make(map[string]interface{})
	orderSkuInfo["order_id"] = row[0]
	orderSkuInfo["product_id"] = 0
	orderSkuInfo["sku_id"] = row[2]
	orderSkuInfo["pic_value_id"] = 28163
	orderSkuInfo["origin_price"] = 0
	orderSkuInfo["unit_price"] = 6.69
	orderSkuInfo["num"] = row[3]
	orderSkuInfo["total_price"] = 6.69
	orderSkuInfo["stock_status"] = 0
	orderSkuInfo["lock_stock"] = 0
	orderSkuInfo["created_at"] = row[15]

	return orderSkuInfo
}

func transOrderInfo(info map[string]interface{}) map[string]interface{} {
	orderInfo := make(map[string]interface{})
	orderInfo["order_id"] = info["order_id"]
	orderInfo["user_id"] = info["user_id"]
	orderInfo["total_num"] = info["total_num"]
	orderInfo["final_price"] = info["final_price"]
	orderInfo["pay_channel"] = info["pay_channel"]
	orderInfo["pay_method"] = info["pay_method"]
	orderInfo["currency_code"] = info["currency_code"]
	orderInfo["abbreviation"] = info["abbreviation"]
	orderInfo["from_type"] = info["from_type"]
	orderInfo["status"] = info["status"]
	orderInfo["language"] = info["language"]
	orderInfo["created_at"] = info["created_at"]
	orderInfo["pay_at"] = info["pay_at"]
	orderInfo["stock_status"] = info["stock_status"]
	return orderInfo
}

func initOrder(row []string,totalNum int, orderInfo map[string]interface{}) map[string]interface{} {
	orderInfo["order_id"] = row[0]
	orderInfo["user_id"] = 11
	orderInfo["total_num"] = totalNum
	orderInfo["final_price"] = 30
	orderInfo["pay_channel"] = "shopify"
	orderInfo["pay_method"] = "pay"
	orderInfo["currency_code"] = "USD"
	orderInfo["abbreviation"] = "US"
	orderInfo["from_type"] = 6
	orderInfo["status"] = 3
	orderInfo["language"] = "en"
	orderInfo["created_at"] = row[15]
	orderInfo["pay_at"] = row[15]
	orderInfo["stock_status"] = 0

	return orderInfo
}