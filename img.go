package img

import (
	"fmt"
	"log"
	"os/exec"
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

func NewImg(sourceFilePath, destinationDirPath string) *Img {
	if !strings.HasSuffix(destinationDirPath, "/") {
		destinationDirPath += "/"
	}
	i := new(Img)
	i.sourceFilePath = sourceFilePath
	i.destinationDirPath = destinationDirPath
	return i
}

type Img struct {
	sourceFilePath     string
	destinationDirPath string
	destPaths          []string
	commands           []*command
}

func (i *Img) PrepareResizeTo(widths ...int) []string {
	i.configureCommands(widths...)
	return i.destPaths
}

func (i *Img) Resize() {
	for _, c := range i.commands {
		err := exec.Command(c.name, c.arguments...).Run()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (i *Img) configureCommands(widths ...int) {
	for _, w := range widths {
		path := i.getPathFor(w)
		i.destPaths = append(i.destPaths, path)
		c := new(command)
		c.name = "convert"
		c.setArgs(i.sourceFilePath, "-resize", strconv.Itoa(w), path)
		i.commands = append(i.commands, c)
	}
}

func (i *Img) getPathFor(w int) string {
	tag := "-" + i.getTagFor(w)
	sf := i.getSourceFileName()
	parts := strings.Split(sf, ".")
	n := strings.Join(parts[:len(parts)-1], "")
	n += tag + "." + parts[len(parts)-1]
	return i.destinationDirPath + n
}

func (i *Img) getTagFor(w int) string {
	return fmt.Sprintf("w%d", w)
}

func (i *Img) getSourceFileName() string {
	parts := strings.Split(i.sourceFilePath, "/")
	return parts[len(parts)-1]
}
