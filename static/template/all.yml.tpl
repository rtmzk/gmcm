dummy:
ceph_origin: distro
ceph_stable_release: octopus
configure_firewall: false
ntp_service_enabled: false
fsid: 78bcda49-b1c8-440f-8138-8440b44800a9
generate_fsid: false
fetch_directory: fetch/
cephx: true
copy_admin_key: false
ip_version: ipv4
public_network:{{" "}}{{ .PublicNetwork }}
{{- if ne .ClusterNetwork "" }}
cluster_network:{{" "}}{{ .ClusterNetwork -}}
{{ end }}
osd_objectstore: bluestore
osd_auto_discovery: false
osd_memory_target: 2147483648
osd_scenario: lvm
os_tuning_params:
  - { name: vm.min_free_kbytes, value: 1048576 }
ceph_mgr_modules: [dashboard,status,restful]
mds_max_mds: 1
cephfs: cephfs
cephfs_data: cephfs_data
cephfs_metadata: cephfs_metadata
cephfs_pools:
  - { name: "{{ "{{ cephfs_data }}" }}", size: "1" }
  - { name: "{{ "{{ cephfs_metadata }}" }}", size: "1" }
radosgw_num_instances: 1
radosgw_civetweb_port: 8080
radosgw_civetweb_num_threads: 500
rgw_override_bucket_index_max_shards: 768
rgw_multisite: true
rgw_realm: macrowing
rgw_zonegroup: edoc2
rgw_zonegroupmaster: true
rgw_zone: master
rgw_zone_user: admin
rgw_zonemaster: true
rgw_zonesecondary: false
rgw_zone_user_display_name: admin
rgw_multisite_endpoint_addr: "{{ "{{ radosgw_address }}" }}"
rgw_multisite_proto: http
system_access_key: 6kWkikvapSnHyE22P7nO
system_secret_key: MGecsMrWtKZgngOHZdrd6d3JxGO5CPWgT2lcnpSt
rgw_create_pools:
  "{{ "{{ rgw_zone }}" }}.rgw.buckets.data":
    pg_num: {{ .Pools.Sp.Root }}
    size: "{{ .Pools.Sp.Replicas }}"
    type: replicated
    rule_name: HDD
  "{{ "{{ rgw_zone }}" }}.rgw.buckets.index":
    pg_num: {{ .Pools.Sp.Index }}
    size: "{{ .Pools.Sp.Replicas }}"
    type: replicated
    {{- if .NoSSD }}
    rule_name: HDD
    {{ else }}
    rule_name: SSD
    {{ end }}
  "{{ "{{ rgw_zone }}" }}.rgw.control":
    pg_num: {{ .Pools.Sp.Control }}
    size: "{{ .Pools.Sp.Replicas }}"
    type: replicated
    rule_name: HDD
  "{{ "{{ rgw_zone }}" }}.rgw.meta":
    pg_num: {{ .Pools.Sp.Meta  }}
    size: "{{ .Pools.Sp.Replicas }}"
    type: replicated
    rule_name: HDD
  "{{ "{{ rgw_zone }}" }}.rgw.log":
    pg_num: {{ .Pools.Sp.Log }}
    size: "{{ .Pools.Sp.Replicas }}"
    type: replicated
    rule_name: HDD
  "{{ "{{ rgw_zone }}" }}.rgw.buckets.non-ec":
    pg_num: {{ .Pools.Sp.NonEc }}
    size: "{{ .Pools.Sp.Replicas }}"
    type: replicated


crush_rule_config: true

crush_rule_hdd:
  name: HDD
  root: default
  type: host
  class: hdd
  default: false

{{- if not .NoSSD }}
crush_rule_ssd:
  name: SSD
  root: default
  type: host
  class: ssd
  default: false
{{ end }}

crush_rules:
  - "{{ "{{ crush_rule_hdd }}" }}"
  {{- if not .NoSSD }}
  - "{{ "{{ crush_rule_ssd }}" }}"
  {{ end }}

#crush_rule_ssd:
#  name: SSD
#  root: default
#  type: host
#  class: ssd
#  default: false

dashboard_port: 7000
dashboard_enabled: True
dashboard_protocol: http
dashboard_admin_user: admin
dashboard_admin_password: 1qaz2WSX
node_exporter_container_image: "registry.edoc2.com:5000/ceph/node-exporter:v0.17.0"
node_exporter_port: 9100
grafana_admin_user: admin
grafana_admin_password: 1qaz2WSX
grafana_container_image: "registry.edoc2.com:5000/ceph/grafana:5.4.3"
grafana_container_memory: 1
grafana_port: 3000
prometheus_container_image: "registry.edoc2.com:5000/ceph/prometheus:v2.7.2"
prometheus_container_memory: 1
prometheus_port: 9092
alertmanager_container_image: "registry.edoc2.com:5000/ceph/alertmanager:v0.16.2"
alertmanager_container_memory: 1
alertmanager_port: 9093
alertmanager_cluster_port: 9094

ceph_conf_overrides: 
  global: 
    osd pool default size: 1
    osd pool default pg num: 8
    osd pool default pgp num: 8
    rgw num rados handles: 2
    rgw multipart part upload limit: 50000
  mon:
    mon allow pool delete: true
    mon max pg per osd: 500
    mon warn on pool no redundancy: False
  mds:
    mds standby replay: True
    mds cache memory limit: 4294967296
  osd:
    osd scrub begin hour: 23
    osd scrub end hour: 6
    osd scrub sleep: 3 
    osd scrub chunk min: 1
    osd scrub chunk max: 1
