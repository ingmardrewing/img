package img

import (
	"image"
	"log"

	"github.com/disintegration/imaging"
)

// returns a struct holding the source and target path
// of an image, with the source being the not
// scaled image and the target being the scaled
// image, while maxSideLength is the number of
// pixels on the longer edge of the target image
func newImageContainer(
	sourceImagePath string) *imgContainer {
	ic := new(imgContainer)
	ic.readSourceImage(sourceImagePath)
	return ic
}

type imgContainer struct {
	src image.Image
}

func (i *imgContainer) readSourceImage(sourceImagePath string) {
	src, err := imaging.Open(sourceImagePath)
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}
	if src == nil {
		log.Fatalf("Image is nil: %v", sourceImagePath)
	}
	i.src = src
}

func (i *imgContainer) resizeAndCropToAndSaveAs(path string, maxSideLength int) {
	dest := imaging.Fill(
		i.src,
		maxSideLength,
		maxSideLength,
		imaging.Center,
		imaging.Lanczos)

	err := imaging.Save(dest, path)
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}
}

func (i *imgContainer) resizeToAndSaveAs(path string, maxSideLength int) {
	dest := imaging.Fit(
		i.src,
		maxSideLength,
		maxSideLength,
		imaging.Lanczos)

	err := imaging.Save(dest, path)
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}
}
