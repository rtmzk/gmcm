# Makefile helper functions #
# ========================= #

BLOCKER_TOOLS ?= golingci-lint statik
DEPEENDENCIES ?= ceph

STATIK ?= statik

CEPH_PACKAGE_URL ?= http://x/ceph_with_gmcm_15.2.5.tar.gz

COMMA := ,
SPACE :=
SPACE +=