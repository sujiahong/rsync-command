package main

import (
	"bytes"
	"os/exec"
)

func doSync(address string, srcDir string) error {
	srcPath := "rsync://" + address + ":8765/test/steamapps/common/" + srcDir
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

func main() {
	err := doSync("10.0.1.210", "TheLongDark")
	println("error: ", err)
}
