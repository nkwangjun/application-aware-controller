package recommendation

// NewController returns a new recommendation controller
func NewController() *Controller {
	// TODO(wangjun): 返回recommendation controller

	return nil
}

// Controller is the controller implementation for recommendation resources
type Controller struct {
}

func (c *Controller) syncHandler(key string) error {
	// TODO(wangjun): 启动一个后台任务, 执行以下流程

	// TODO(wangjun): 获取nodegroup config详情

	// TODO(wangjun): 到GSCO获取集群的资源预测总量

	// TODO(wangjun): 到CA获取NodeGroup的Node已发放量

	// TODO(wangjun): 根据资源总量到GlobalScheduler获取资源可发放量

	// TODO(wangjun): 创建nodegroup-scale-recommendation资源

	return nil
}