package lmjaeger

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/logicmonitor/k8s-collectorset-controller/pkg/config"
	"github.com/logicmonitor/k8s-collectorset-controller/pkg/lmctx"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"gopkg.in/yaml.v2"
)

func Initialise(lmconf *config.Config) {
	conf := readConf()
	cl, _ := conf.InitGlobalTracer(lmconf.Account + "-csc")
	shutdownHook(cl)
}

func shutdownHook(closer io.Closer) {
	go func(cl io.Closer) {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Warn("Closing tracer on shutdown signal")
		cl.Close()
	}(closer)
}

func readConf() *jaegercfg.Configuration {
	configBytes, err := ioutil.ReadFile("/etc/collectorset-controller/jaeger-config.yaml")
	if err != nil {
		log.Fatalf("Failed to read jaeger config file: /etc/collectorset-controller/jaeger-config.yaml")
	}
	log.Debugf("jaeger config raw: %s", configBytes)
	conf := &jaegercfg.Configuration{}
	err = yaml.Unmarshal(configBytes, conf)
	if err != nil {
		log.Fatalf("Couldn't parse jaeger-config.yaml: %s", err.Error())
	}
	log.Infof("jaeger config read: %+v", conf)
	return conf
}

type LMSpan interface {
	opentracing.Span
	Info(alternatingKeyValues ...interface{})
	Error(alternatingKeyValues ...interface{})
	K8sObject(key string, object interface{})
}

type LMSpanObject struct {
	*jaeger.Span
}

func (span *LMSpanObject) Info(alternatingKeyValues ...interface{}) {
	alternatingKeyValues = append(alternatingKeyValues, "level")
	alternatingKeyValues = append(alternatingKeyValues, "info")
	span.LogKV(alternatingKeyValues...)
}
func (span *LMSpanObject) K8sObject(key string, object interface{}) {
	objJson, err := json.Marshal(object)
	if err != nil {
		span.LogEventWithPayload(key, object)
	} else {
		span.LogEventWithPayload(key, string(objJson))
	}
}

func (span *LMSpanObject) Error(alternatingKeyValues ...interface{}) {
	alternatingKeyValues = append(alternatingKeyValues, "level")
	alternatingKeyValues = append(alternatingKeyValues, "info")
	span.LogKV(alternatingKeyValues...)
	span.SetTag("error", true)
}

// StartSpan creates new span and injects it in lmcontext object
func StartSpan(lctx *lmctx.LMContext, operationName string, options ...opentracing.StartSpanOption) LMSpan {
	span := opentracing.StartSpan(operationName, options...)
	lmSpanObj := &LMSpanObject{
		Span: span.(*jaeger.Span),
	}
	lctx.Set("span", lmSpanObj)
	return lmSpanObj
}

// Span returns span object from lmcontext object
func Span(lctx *lmctx.LMContext) LMSpan {
	return lctx.Extract("span").(LMSpan)
}
