# 适配器模式

适配器模式用于转换一种接口适配另一种接口。

实际使用中`Adaptee`一般为接口，并且使用工厂函数生成实例。

在`Adapter`中匿名组合`Adaptee`接口，所以`Adapter`类也拥有`SpecificRequest`实例方法，又因为`Go`语言中非入侵式接口特征，其实`Adapter`也适配`Adaptee`接口。
