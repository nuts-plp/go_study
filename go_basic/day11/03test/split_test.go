package split_string
import(
	"reflect"
	"testing"
	
)

func TestSplit(t *testing.T) {
	got := SplitStr("a:b:c",":")
	want :=[]string{"a","b","c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

