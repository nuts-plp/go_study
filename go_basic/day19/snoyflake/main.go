package main

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// 需要传入当前机器id
func Init(startTime string, machineID uint16) (err error) {
	sonyMachineID = machineID
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}
func GenID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake init failed")
		return
	}
	id, err = sonyFlake.NextID()
	return
}

func main() {
	if err := Init("2022-07-01", 1); err != nil {
		fmt.Printf("Init snoyflake failed  err:%v\n", err)
		return
	}
	id, _ := GenID()
	fmt.Println(id)
}
