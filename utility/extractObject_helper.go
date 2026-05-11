package utility

import (
	"fmt"
	"strings"

)


func ExtractObjectPath(fileURL, bucket string) string {
	prefix := fmt.Sprintf("/storage/v1/object/public/%s/", bucket)
	idx := strings.Index(fileURL, prefix)
	if idx == -1 {
		return ""
	}
	return fileURL[idx+len(prefix):]
}