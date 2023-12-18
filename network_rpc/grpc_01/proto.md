>1、syntax  
> 声明proto文件版本

>2、option  go_package="./;service"
> 声明生成go代码

>3、message   
> protobuf中定义一个消息类形式通过该关键字定义，相当于go中的结构体，定义对象。消息就是传输数据的格式，在消息中承载数据对应每一个字段，每个字段都有名和类型。

>4、字段规则
> >required:消息体中必填字段，不设置会导致编码异常，再protobuf2中使用，再protobuf3中被删去
> >optional:消息体可选字段。protobuf3中没有了required、optional等关键字说明，都默认为optional
> >repeated:消息体中可重复字段，重复的值的顺序会被保留在go中，重复的会被定义为切片
> 消息号
> >在消息体中，每个字段都必须有一个唯一标识的标识号，是[1,2^29-1]范围中的一个整数
> 嵌套消息
> >可以在其他消息类型中定义，使用消息类型，如下示例      

```protobuf
syntax = "proto3";
message PersonInfo{
  message Person{
    string name = 1;
    int32 height = 2;
    repeated int32 weight = 3;
  }
  repeated Person info = 1;
}
```

>5、服务定义
> 如果想要将消息类型用在rpc系统中，可以在.proto文件中定义一个rpc服务接口，protocol buffer编译器会根据所选择的生成不同语言服务接口代码及存根  

```protobuf
syntax="proto3";
service ServiceSearch{
  rpc Search(SearchRequest)returns(SearchResponse);
  
}
```
>上述定义了一个rpc服务，该方法接受SearchRequest为参数返回SearchResponse


>两种认证方式
> > TLS
> > token