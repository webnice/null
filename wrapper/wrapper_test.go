package wrapper // import "gopkg.in/webnice/nul.v1/wrapper"

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"testing"
)

func TestExistsWrapers(t *testing.T) {
	_ = &BoolWrapper{}
	_ = &BytesWrapper{}
}