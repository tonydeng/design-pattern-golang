package abstract_factory

import "fmt"

//订单主记录的Dao
type OrderMainDao interface {
	SaveOrderMain()
}

// 订单详情记录的Dao
type OrderDetailDao interface {
	SaveOrderDetail()
}

//抽象模式工厂接口
type DaoFactory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

//关系型数据库的MainDao实现
type RdbMainDao struct {
}

func (*RdbMainDao) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

// 关系型数据库的DetailDao实现
type RdbDetailDao struct {
}

func (*RdbDetailDao) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

//RDB的抽象工厂实现
type RdbDaoFactory struct {
}

func (*RdbDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &RdbMainDao{}
}

func (*RdbDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &RdbDetailDao{}
}

//XML MainDao存储
type XmlMainDao struct {
}

func (*XmlMainDao) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

// XML DetailDao实现
type XmlDetailDao struct {
}

func (*XmlDetailDao) SaveOrderDetail() {
	fmt.Print("xml detail save\n")
}

// XML的抽象工厂实现
type XmlDaoFactory struct {
}

func (*XmlDaoFactory) CreateOrderMainDao() OrderMainDao {
	return &XmlMainDao{}
}

func (*XmlDaoFactory) CreateOrderDetailDao() OrderDetailDao {
	return &XmlDetailDao{}
}
