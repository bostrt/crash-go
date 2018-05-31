```
# go build
# ./crash-go 
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x44c152]

goroutine 1 [running]:
main.doBadThings(...)
	/home/user/crash-go/main.go:14
main.main()
	/home/user/crash-go/main.go:9 +0x2

```
