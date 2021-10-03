package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)
func main() {
	TestString := "Hi,hero"
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result :=Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n",Result)
	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Result = Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n",Result)

	
	TestFile :="READ.md"
	infile,inerr:=os.Open(TestFile)
	if inerr == nil {
		md5h :=md5.New()
		io.Copy(md5h,infile)
		fmt.Printf("%x %s\n",md5h.Sum([]byte("")),TestFile)
		sha1h := sha1.New()
		io.Copy(sha1h,infile)
		fmt.Printf("%x %s\n",sha1.Sum([]byte("")),TestFile)
	}else{
		fmt.Println(inerr)
		os.Exit(1)
	}
}