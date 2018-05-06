package img

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func NewImgScaler(sourceFilePath, destinationDirPath string) *ImgScaler {
	if !strings.HasSuffix(destinationDirPath, "/") {
		destinationDirPath += "/"
	}
	i := new(ImgScaler)
	i.sourceFilePath = sourceFilePath
	i.destinationDirPath = destinationDirPath
	return i
}

type ImgScaler struct {
	maxSizes           []int
	sourceFilePath     string
	destinationDirPath string
	destPaths          []string
}

func (i *ImgScaler) PrepareResizeTo(widths ...int) []string {
	i.maxSizes = widths
	return i.getPaths()
}

func (i *ImgScaler) Resize() {
	ic := newImageContainer(i.sourceFilePath)
	for _, ms := range i.maxSizes {
		ic.resizeToAndSaveAs(i.getPathFor(ms), ms)
	}
}

func (i *ImgScaler) getPaths() []string {
	pths := []string{}
	for _, ms := range i.maxSizes {
		pths = append(pths, i.getPathFor(ms))
	}
	return pths
}

func (i *ImgScaler) getPathFor(ms int) string {
	fileExtension := filepath.Ext(i.sourceFileName())
	basename := strings.TrimSuffix(i.sourceFileName(), fileExtension)
	newFilename := basename + i.getSizeTagForFilename(ms) + fileExtension
	return path.Join(i.destinationDirPath, newFilename)
}

func (i *ImgScaler) getSizeTagForFilename(w int) string {
	return fmt.Sprintf("-w%d", w)
}

func (i *ImgScaler) sourceFileName() string {
	return filepath.Base(i.sourceFilePath)
}
