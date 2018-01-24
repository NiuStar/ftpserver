package ftpserver1

import (
	"fmt"
	"github.com/NiuStar/filedriver1"
	"github.com/NiuStar/ldbauth"
	"github.com/goftp/server"
	"os"
)

func StartFtp(name, pass, path string, port int) {
	go func() {
		perm := server.NewSimplePerm("root", "root")
		var userList map[string]string = make(map[string]string)
		userList[name] = pass
		var auth = &ldbauth.LDBAuth{Users: userList}
		var factory server.DriverFactory

		rootPath := path //"../"
		_, err := os.Lstat(rootPath)
		if os.IsNotExist(err) {
			os.MkdirAll(rootPath, os.ModePerm)
		} else if err != nil {
			fmt.Println(err)
			return
		}
		factory = &filedriver1.FileDriverFactory{
			rootPath,
			perm,
		}
		ser := server.NewServer(&server.ServerOpts{Factory: factory, Auth: auth, Port: port})
		err = ser.ListenAndServe()
		fmt.Println(err)
	}()
}
