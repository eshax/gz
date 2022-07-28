package gz

import "testing"

func Test_Extract(t *testing.T) {
	Extract("dist/2022-07-27-204233.zip", "index.json", "dist/2022-07-27-204233")
}
