package static

import _ "embed"

//go:embed .env
var DefaultConfigContent []byte

//go:embed ceph.tar.gz
var CephPackage []byte

//go:embed initdata.sql
var InitDatabaseSql []byte

//go:embed template/all.yml.tpl
var ALL_YML_TPL []byte

//go:embed template/ceph_install.conf.tpl
var CEPH_INSTALL_CONF_TPL []byte

//go:embed template/host-inventory.tpl
var HOST_INVENTORY_TPL []byte

//go:embed check_rule/check_rule.json
var CHECK_RULES []byte

//go:embed check_rule/envcheck.sh
var CHECK_SCRIPTS []byte
