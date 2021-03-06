package di

import (
	"go.uber.org/dig"
)

func MustContainerProvide(c *dig.Container, constructor interface{}, opts ...dig.ProvideOption) {
	err := c.Provide(constructor, opts...)
	if err != nil {
		//log.Fatalf(ctx,"Failed to provide constructor. err:[%v]", err)
	}
}

func MustContainerInvoke(c *dig.Container, function interface{}, opts ...dig.InvokeOption) {
	err := c.Invoke(function, opts...)
	if err != nil {
		//log.Fatalf(ctx, "Failed to invoke. err:[%v]", err)
	}
}
