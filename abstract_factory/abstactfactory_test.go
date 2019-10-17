package abstract_factory

import "testing"

func getMainAndDetail(factory DaoFactory) {
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
}

func TestExampleRdbFactory(t *testing.T) {
	var factory DaoFactory

	factory = &RdbDaoFactory{}
	getMainAndDetail(factory)
}

func TestExampleXmlFactory(t *testing.T) {
	var factory DaoFactory

	factory = &XmlDaoFactory{}
	getMainAndDetail(factory)
}
