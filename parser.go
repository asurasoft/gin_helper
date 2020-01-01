package gin_helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"reflect"
)

type Checker interface {
	CheckValid() error
}

func ParseData(c *gin.Context, obj interface{}) error{
	if reflect.TypeOf(&obj).Kind() != reflect.Struct{
		return errors.New("obj必须是ptr -> struct")
	}
	requestFormat := c.Request.Header.Get("Content-Type")
	if requestFormat == binding.MIMEJSON{//使用声明好的常量
		c.ShouldBindJSON(&obj)
		if checker, ok := obj.(Checker); ok{//判断一个对象是否响应响应的接口
			return checker.CheckValid()
		}
		return nil
	}else if requestFormat == binding.MIMEPOSTForm || requestFormat == binding.MIMEMultipartPOSTForm{ //form请求
		c.ShouldBind(&obj)
		if checker, ok := obj.(Checker); ok{//判断一个对象是否响应响应的接口
			return checker.CheckValid()
		}
	}
	return nil
}
