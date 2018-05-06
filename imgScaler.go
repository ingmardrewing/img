package img

import (
	"fmt"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type command struct {
	name      string
	arguments []string
}

func (c *command) setArgs(args ...string) {
	for _, a := range args {
		c.arguments = append(c.arguments, a)
	}
}

const (
	jpgImg = iota
	pngImg = iota
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
	sourceFilePath     string
	destinationDirPath string
	destPaths          []string
	commands           []*command
}

func (i *ImgScaler) PrepareResizeTo(widths ...int) []string {
	i.configureCommands(widths...)
	return i.destPaths
}

func (i *ImgScaler) Resize() {
	for _, c := range i.commands {
		err := exec.Command(c.name, c.arguments...).Run()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (i *ImgScaler) configureCommands(widths ...int) {
	for _, w := range widths {
		path := i.getPathFor(w)
		i.destPaths = append(i.destPaths, path)

		c := new(command)
		c.name = "convert"
		c.setArgs(i.sourceFilePath, "-resize", strconv.Itoa(w), path)

		i.commands = append(i.commands, c)
	}
}

func (i *ImgScaler) getPathFor(w int) string {
	return path.Join(i.destinationDirPath, i.assembleFileNameFor(w))
}

func (i *ImgScaler) assembleFileNameFor(w int) string {
	fileExtension := filepath.Ext(i.sourceFileName())
	basename := strings.TrimSuffix(i.sourceFileName(), fileExtension)

	return basename + i.getSizeTagForFilename(w) + fileExtension
}

func (i *ImgScaler) getSizeTagForFilename(w int) string {
	return fmt.Sprintf("-w%d", w)
}

func (i *ImgScaler) sourceFileName() string {
	return filepath.Base(i.sourceFilePath)
}
