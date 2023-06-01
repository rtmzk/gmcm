package server

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/gin-gonic/gin"
	"gmcm/pkg/db"
	"gmcm/pkg/utils"
	"gmcm/static"
	"io"
	"os"
	"path/filepath"
)

type server struct {
	router *gin.Engine
}

func NewServer(e *gin.Engine) *server {
	return &server{
		router: e,
	}
}

func (s *server) Start(listen string) error {
	dbIns, err := db.SetupDB()
	if err != nil {
		return err
	}
	db.SetClient(dbIns)
	setupCephPackage()
	return s.router.Run(listen)
}

func setupCephPackage() error {
	if utils.IsExist("/tmp/ceph.tar.gz") {
		return nil
	}
	err := os.WriteFile("/tmp/ceph.tar.gz", static.CephPackage, 0644)
	if err != nil {
		os.Exit(1)
	}

	fr, err := os.Open("/tmp/ceph.tar.gz")
	if err != nil {
		os.Exit(1)
	}
	defer fr.Close()

	gr, err := gzip.NewReader(fr)
	if err != nil {
		os.Exit(1)
	}
	defer gr.Close()
	tr := tar.NewReader(gr)

	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			os.Exit(1)
		}

		fpath := filepath.Join("/tmp/ceph", h.Name)
		if h.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}
			fw, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			defer fw.Close()
			_, err = io.Copy(fw, tr)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}
	return nil
}
