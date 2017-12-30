package img

import (
	"reflect"
	"testing"
)

func TestPrepareResizeReturnsCorrectPaths(t *testing.T) {
	i := NewImg(
		"/Users/drewing/Desktop/dt_02042017/blog_local/atthezoo.png",
		"/Users/drewing/Desktop")
	actual := i.PrepareResizeTo(800, 390)
	expected := []string{
		"/Users/drewing/Desktop/atthezoo-w800.png",
		"/Users/drewing/Desktop/atthezoo-w390.png"}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("Expected ", expected, " but got ", actual)
	}
}

func TestGetPathFor(t *testing.T) {
	img := new(Img)

	img.sourceFilePath = "/a/path/to/an/image.png"
	img.destinationDirPath = "/another/path/"
	actual := img.getPathFor(800)
	expected := "/another/path/image-w800.png"
	if actual != expected {
		t.Fatal("Expected ", expected, " but got ", actual)
	}
}

func TestGetSourceFileName(t *testing.T) {
	img := new(Img)

	img.sourceFilePath = "/a/path/to/an/image.png"
	actual := img.getSourceFileName()
	expected := "image.png"
	if actual != expected {
		t.Fatal("Expected ", expected, " but got ", actual)
	}
}
