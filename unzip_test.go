package gz

import (
	"log"
	"testing"
)

func Test_unzip(t *testing.T) {

	log.Println()

	UnZip("dist/2022-07-27-210547.zip", "dist/2022-07-27-210547")

	log.Println()
}
