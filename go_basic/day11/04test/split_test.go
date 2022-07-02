package split

import (
	"reflect"
	"testing"
	// "fmt"
)
//单例测试     go test -cover -coverprofile=filename
func TestSplit1(t *testing.T){
	ret := Split("a:b:c",":")
	want :=[]string{"a","b","c"}
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("got %#v, want %#v", ret, want)
	}
}

func TestSplit2(t *testing.T){
	ret := Split("asdafsdsasadas","as")
	want :=[]string{"","dafsds","ad",""}
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("got %#v, want %#v", ret, want)
	}
}


func TestSplit3(t *testing.T){
	ret := Split("absbhasqbsadaf","a")
	want :=[]string{"","bsbh","sqbs","d","f"}
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("got %#v, want %#v", ret, want)
	}
}


//测试组
func TestSplit4(t *testing.T){
	type testCase struct {
		str string
		sep string
		want []string
	}
	testGroup := []testCase{
		testCase{"absbhasqbsadaf","a",[]string{"","bsbh","sqbs","d","f"}},
		testCase{"a:b:c",":",[]string{"a","b","c"}},
		testCase{"asdafsdsasadas","as",[]string{"","dafsds","ad",""}},
		testCase{"你好啊，亲爱的小潘！","亲",[]string{"你好啊，","爱的小潘！"}},
	}
	for _,tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got,tc.want){
			t.Errorf("got %#v, want %#v", got,tc.want)
			
		}
	}
}


//子测试
func TestSplit5(t *testing.T) {
	type testCase struct {
		str string
		sep string
		want []string
	}
	testGroup := map[string]testCase{
		"case_1":{"absbhasqbsadaf","a",[]string{"","bsbh","sqbs","d","f"}},
		"case_2":{"a:b:c",":",[]string{"a","b","c"}},
		"case_3":{"asdafsdsasadas","as",[]string{"","dafsds","ad",""}},
		"case_4":{"你好啊，亲爱的小潘！","亲",[]string{"你好啊，","爱的小潘！"}},
	}
	for name,tc := range testGroup{
		t.Run(name,func(t *testing.T){
			got :=Split(tc.str,tc.sep)
			if !reflect.DeepEqual(got,tc.want) {
				t.Errorf("got %#v, want %#v", got,tc.want)
			}
		})
	}
}

//并行测试
//RunParallel会创建多个gorutine，并将b.N分配给这些gorutine执行，gorutine默认为GOMAXPROCES值
//  修改使用SetParallelism    Setparallelism通常与-cpu搭配使用
func BenchmarkSplitParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河","沙")
	}
	})
	
}


//基准测试  时间和内存   go test -benchSplit -benchmem -benchtime -cpu
//基准测试重置时间  b.ResetTimer()
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c",":")
	}
}

//性能比较函数测试
func benchmarkFib(b *testing.B,n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
	
}


func BenchmarkFib1(b *testing.B){
	benchmarkFib(b,1)
}
func BenchmarkFib5(b *testing.B){
	benchmarkFib(b,5)
}
func BenchmarkFib10(b *testing.B){
	benchmarkFib(b,10)
}
func BenchmarkFib15(b *testing.B){
	//设置cpu核心数

	benchmarkFib(b,15)
}
func BenchmarkFib20(b *testing.B){
	benchmarkFib(b,20)
}
func BenchmarkFib50(b *testing.B){
	benchmarkFib(b,50)
}