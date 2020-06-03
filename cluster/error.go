package cluster

import (
	"fmt"
)

func NewError(name string, err error) error {
	return fmt.Errorf("err: %v, name: %s, ", err, name)
}

func createPoolErr(err error) error {
	return fmt.Errorf("CreatePoolErr %v", err)
}

func getConnErr(err error) error {
	return fmt.Errorf("GetConnFromPoolErr %v", err)
}

func unexpectedPkgLenErr(receive, expect int) error {
	return fmt.Errorf("UnexpectedLenErr received pkg length %d != expected %d", receive, expect)
}

func wrongFidErr(fid string) error {
	return fmt.Errorf("WrongFidErr fid is not with format group/filename: %s", fid)
}
