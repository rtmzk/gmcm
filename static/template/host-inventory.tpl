[mons]
{{ range .Host -}}
{{ if find .NodeRole "mon" }}
{{- .IP }} {{ "" }}monitor_address={{ .IP }}
{{ end }}
{{- end }}

[mgrs]
{{ range .Host -}}
{{ if find .NodeRole "mon" }}
{{- .IP }}
{{ end }}
{{- end }}

[mdss]
#192.168.251.153


[rgws]
{{ range .Host -}}
{{ if find .NodeRole "rgw" }}
{{- .IP }} {{ "" }}radosgw_address={{ .IP }}
{{ end }}
{{- end }}

[osds]
{{- range $key, $value := .Device }}
{{ $key }} {{ "" }}devices="[{{ setDevice $value }}]" {{ if ne (setCacheDevice $value) "" }}{{ "" }}dedicated_devices="[{{ setCacheDevice $value }}]"{{end}}
{{- end }}
{{ "" }}

[all:vars]
ansible_ssh_user="root"
ansible_ssh_port={{ (index .Host 0).SSHPort }}