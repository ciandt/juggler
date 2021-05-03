package internal

import (
	"fmt"
	"io"
	"net"
	"time"
)

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()

	copyBuffer(dest, src, nil)
}

func copyBuffer(dst io.Writer,
	src io.Reader,
	buf []byte) (written int64, err error) {
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*io.LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	for {
		nr, er := src.Read(buf)
		if nr > 0 {

			//TODO: Add behavior
			//Struct wrapper
			//Pass a interface/function
			//Use a specific buffer

			// rand.Seed(time.Now().UnixNano())
			// dr, _ := time.ParseDuration(fmt.Sprintf("%dms", rand.Intn(120)))

			// time.Sleep(dr)

			// fmt.Println(time.Now(), "Writing with delay ", dr)

			time.Sleep(time.Second * 3)
			fmt.Println(time.Now(), "Writing with delay ")

			nw, ew := dst.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = io.ErrShortWrite
				}
			}
			written += int64(nw)
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
