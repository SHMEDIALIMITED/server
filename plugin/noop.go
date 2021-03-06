package plugin

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gocraft/web"
	"net/http"
)

const (
	PLUGIN_NOOP string = "noop"
)

type NoopPlugin struct{}

func (p *NoopPlugin) Bootstrap(config *Config) (Interface, error) {
	log.Warn("NoopPlugin::Bootstrap")
	var err error
	return p, err
}

func (p *NoopPlugin) Inbound(rw web.ResponseWriter, req *web.Request) (int, error) {
	log.Warn("NoopPlugin::Inbound")
	var err error
	return http.StatusOK, err
}

func (p *NoopPlugin) Outbound(res *http.Response) (int, error) {
	log.Warn("NoopPlugin::Outbound")
	var err error
	return http.StatusOK, err
}
