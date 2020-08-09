package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)


func main() {

	Db = NewOrm()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s", name)
	})

	r.GET("/products", getProducts)
	r.GET("/product", getProduct)

	r.GET("/handler", getHandler)

	r.Run(":8080")
}

func getProduct(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(id)
	var product= &Product{}
	Db.First(&product, 15)
	Db.GetById(15,product)
	c.JSON(200, *product)
}

func getHandler(c *gin.Context) {
	fmt.Println(c.Request.Host + c.Request.URL.Path)
}

var Db MyDb

func NewOrm() MyDb {
	db, err := gorm.Open("mysql", "dev:dev@tcp(192.168.40.160:3306)/product_db?charset=utf8")
	if err != nil {
		panic("orm conn err:" + err.Error())
	}
	return MyDb{db}
}

type MyDb struct {
	*gorm.DB
}

func (m MyDb)GetById(id,out interface{})  {
	m.First(&out, id)
}

type Product struct {
	Id                int    `json:"id"`
	MongoId           string `json:"mongo_id"`
	Type              int    `json:"type"`                //商品类型 0默认 2组合
	Name              string `json:"name"`                //商品售卖名
	BaseName          string `json:"base_name"`           //商品名
	Spec              string `json:"spec"`                //商品描述
	NetWeight         string `json:"net_weight"`          //商品净含量
	NetWeightUnit     string `json:"net_weight_unit"`     //商品净含量单位
	ViewWeight        string `json:"view_weight"`         //展示净含量
	OriginId          string `json:"origin_id"`           //产地ID
	QualityPeriod     int    `json:"quality_period"`      //保质期
	QualityPeriodUnit int    `json:"quality_period_unit"` //保质期单位
	SalePeriod        int    `json:"sale_period"`         //可售期
}

//表名称
func (Product) TableName() string {
	return "product"
}

func getProducts(c *gin.Context) {
	var (
		ps          []Product
		productsMap map[string]Product
	)
	fmt.Println(Db)
	if err := Db.Where("is_sale = ? and status = ?", 1, 1).Find(&ps).Error; err != nil {
		log.Println("商品数据初始化异常:" + err.Error())
	}

	productsMap = make(map[string]Product)

	for _, product := range ps {
		if product.MongoId == "" {
			fmt.Println(string(product.Id) + " mongoId 为空")
			continue
		}
		productsMap[product.MongoId] = product
	}
	c.JSON(200, productsMap)
}
