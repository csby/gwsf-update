package main

import (
	"fmt"
	"github.com/csby/gwsf/gpkg"
	"path"
	"path/filepath"
)

const (
	srcPath = "src/github.com/csby/gwsf-update"
)

var (
	goIgnore = []string{
		".git",
		".idea",
		".gitignore",
		"go.sum",
	}
)

type Pkg struct {
	binPath string
}

func (s *Pkg) Run() {
	// app
	fmt.Println("binary file path: ", s.binPath)
	binFolder := filepath.Dir(s.binPath)
	fmt.Println("binary folder path: ", binFolder)
	_, binName := filepath.Split(s.binPath)
	fmt.Println("binary file name: ", binName)
	binExt := path.Ext(binName)
	fmt.Println("binary file ext: ", binExt)
	appName := moduleName
	appFileName := fmt.Sprintf("%s%s", appName, binExt)
	fmt.Println("app file name: ", appFileName)

	tmpFolder := filepath.Dir(binFolder)
	srcFolder := filepath.Join(filepath.Dir(tmpFolder), srcPath)
	fmt.Println("source folder path: ", srcFolder)

	relFolder := filepath.Join(tmpFolder, "rel", moduleName)
	fmt.Println("output folder path: ", relFolder)

	c := &gpkg.Config{
		Version:     moduleVersion,
		Destination: relFolder,
		Source:      true,
		Apps: []gpkg.Application{
			{
				Enable: true,
				Name:   appName,
				Bin: gpkg.Binary{
					Root: binFolder,
					Files: map[string]string{
						binName: appFileName,
					},
				},
				Src: gpkg.Source{
					Enable: true,
					Root:   srcFolder,
					Ignore: goIgnore,
				},
			},
		},
	}

	p := gpkg.NewPacker(c)
	e := p.Pack()
	if e != nil {
		fmt.Println("打包失败: ", e)
	} else {
		fmt.Println("打包成功: ", p.OutputFolder())
	}
}
