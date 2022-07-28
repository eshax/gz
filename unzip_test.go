package gz

import (
	"log"
	"testing"
)

func Test_unzip(t *testing.T) {

	log.Println()

	UnZip("dist/2022-07-22-101209.zip", "dist/2022-07-22-101209")

	log.Println()
}
