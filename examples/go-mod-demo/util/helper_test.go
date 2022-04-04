package util

import (
	"testing"
)

// 执行测试：
// go test -v util/helper_test.go util/helper.go

func Test_UnixTime(t *testing.T) {
	t.Log("---> Test_UnixTime")

	u_time := UnixTime()

	t.Log("---> Test_UnixTime : ", u_time)
}

func Test_UnixTime2DateTime(t *testing.T) {
	t.Log("---> Test_UnixTime2DateTime")

	input := 1612563956
	want := "2021-02-06 06:25:56"

	result := UnixTime2DateTime(int64(input))
	t.Log("---> Test_UnixTime2DateTime unixtime: ", input)
	t.Log("---> Test_UnixTime2DateTime datetime: ", result)
	if result != want {
		t.Errorf("Test_UnixTime2DateTime failed: got %v, want %v", result, want)
	}
}
