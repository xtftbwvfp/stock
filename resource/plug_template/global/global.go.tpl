package global

{{- if .HasGlobal }}

import "github.com/ykstudy/stock/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}