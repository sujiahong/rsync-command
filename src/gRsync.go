package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func doSync(address string, srcDir string) error {
	srcPath := "'rsync://" + address + ":8765/test/steamapps/common/" + srcDir + "'"
	var cmd = exec.Command("rsync", "-vzrhtopg", "--progress", "--delete", "--bwlimit=4000", srcPath, "/cygdrive/e/SteamLibrary/steamapps/common")
	var outBuf bytes.Buffer
	cmd.Stdout = &outBuf
	var errBuf bytes.Buffer
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		println("error: ", err, errBuf.String())
		return err
	}
	println("out: ", outBuf.String())
	return nil
}

func fileLock(f *os.File) {
	// var fd = f.Fd()
	// err := syscall.Flock(int(fd), syscall.LOCK_EX|syscall.LOCK_NB)
	// if err != nil {
	// 	fmt.Println("cannot lock file:  ", err)
	// 	return
	// }
	// return
	h, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		fmt.Println("err == ", err)
		return
	}
	defer syscall.FreeLibrary(h)
	addr, err := syscall.GetProcAddress(h, "LockFile")
	if err != nil {
		fmt.Println("22222  err == ", err)
		return
	}
	r1, r2, err := syscall.Syscall6(addr, 5, f.Fd(), 0, 0, 0, 1, 0)
	fmt.Println("rrr == ", r1, r2, err)
}

func fileUnlock(f *os.File) {
	h, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		fmt.Println("err == ", err)
		return
	}
	defer syscall.FreeLibrary(h)
	addr, err := syscall.GetProcAddress(h, "UnlockFile")
	if err != nil {
		fmt.Println("22222  err == ", err)
		return
	}
	r1, r2, err := syscall.Syscall6(addr, 5, f.Fd(), 0, 0, 0, 1, 0)
	fmt.Println("rrr == ", r1, r2, err)
}

func main() {
	// err := doSync("10.0.1.210", "a")
	// println("error: ", err)
	f, err := os.Open("../file.lock")
	if err != nil {
		fmt.Println("eee: ", err)
		return
	}
	fileLock(f)
	//fileLock(f)
	fileUnlock(f)
	go func() {
		f, err := os.Open("../file.lock")
		if err != nil {
			fmt.Println("eee: ", err)
			return
		}
		var buf = make([]byte, 5)
		n, err := f.Read(buf)
		fmt.Println("n == ", n, err, buf)
	}()
	time.Sleep(2000 * time.Millisecond)
}
