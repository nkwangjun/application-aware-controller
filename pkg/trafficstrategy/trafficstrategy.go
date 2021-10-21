package trafficstrategy

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"

	"k8s.io/klog/v2"
)

type WebServer struct {
	trafficstrategy string
	forecastWindow  int64
	nodepoolsConf   *NodepoolsConf
}

// NodepoolsConf 类型
type NodepoolsConf struct {
	Nodepools []NodepoolConf `yaml:"nodepools"`
}

// NodepoolConf 类型
type NodepoolConf struct {
	Name        string   `yaml:"name"`
	MinSize     int32    `yaml:"minSize"`
	MaxSize     int32    `yaml:"maxSize"`
	NodeFlavors []string `yaml:"nodeFlavors"`
}

func NewWebServer(trafficStrategyAddr string, forecastWindow int64, nodepoolConfig string) (*WebServer, error) {
	config, err := LoadFromFile(nodepoolConfig)
	if err != nil {
		return nil, err
	}

	ws := &WebServer{
		trafficstrategy: trafficStrategyAddr,
		forecastWindow:  forecastWindow,
		nodepoolsConf:   config,
	}

	return ws, nil
}

func LoadFromFile(nodepoolConfig string) (*NodepoolsConf, error) {
	nodepoolConfigBytes, err := ioutil.ReadFile(nodepoolConfig)
	if err != nil {
		return nil, err
	}

	config := &NodepoolsConf{}
	err = yaml.Unmarshal(nodepoolConfigBytes, config)
	if err != nil {
		return nil, err
	}

	return config, err
}

func (ws *WebServer) Serve() {
	r := mux.NewRouter()
	r.HandleFunc("/trafficstrategy", ws.trafficStrategy)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ws.trafficstrategy,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	klog.Fatal(srv.ListenAndServe())
}

// 分流策略实现函数
func (ws *WebServer) trafficStrategy(w http.ResponseWriter, r *http.Request) {
	// TODO(wangjun): 到GSCO获取集群的资源预测总量

	// TODO(wangjun): 到CA获取NodeGroup的Node已发放量

	// TODO(wangjun): 根据资源总量到GlobalScheduler获取资源可发放量

	// TODO(wangjun): 返回各个group的分流资源数
}
