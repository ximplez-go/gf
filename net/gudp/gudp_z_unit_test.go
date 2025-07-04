// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gudp_test

import (
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/ximplez-go/gf/net/gudp"
	"github.com/ximplez-go/gf/os/glog"
	"github.com/ximplez-go/gf/test/gtest"
	"github.com/ximplez-go/gf/util/gconv"
)

var (
	simpleTimeout = time.Millisecond * 100
	sendData      = []byte("hello")
)

func startUDPServer(addr string) *gudp.Server {
	s := gudp.NewServer(addr, func(conn *gudp.ServerConn) {
		defer conn.Close()
		for {
			data, remote, err := conn.Recv(-1)
			if err != nil {
				if err != io.EOF {
					glog.Error(context.TODO(), err)
				}
				break
			}
			if err = conn.Send(data, remote); err != nil {
				glog.Error(context.TODO(), err)
			}
		}
	})
	go s.Run()
	time.Sleep(simpleTimeout)
	return s
}

func Test_Basic(t *testing.T) {
	var ctx = context.TODO()
	s := gudp.NewServer(gudp.FreePortAddress, func(conn *gudp.ServerConn) {
		defer conn.Close()
		for {
			data, remote, err := conn.Recv(-1)
			if len(data) > 0 {
				if err = conn.Send(append([]byte("> "), data...), remote); err != nil {
					glog.Error(ctx, err)
				}
			}
			if err != nil {
				break
			}
		}
	})
	go s.Run()
	defer s.Close()

	time.Sleep(100 * time.Millisecond)

	// gudp.Conn.Send
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			conn, err := gudp.NewClientConn(s.GetListenedAddress())
			t.AssertNil(err)
			t.Assert(conn.Send([]byte(gconv.String(i))), nil)
			t.AssertNE(conn.RemoteAddr(), nil)
			result, _, err := conn.Recv(-1)
			t.AssertNil(err)
			t.AssertNE(conn.RemoteAddr(), nil)
			t.Assert(string(result), fmt.Sprintf(`> %d`, i))
			conn.Close()
		}
	})
	// gudp.Conn.SendRecv
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			conn, err := gudp.NewClientConn(s.GetListenedAddress())
			t.AssertNil(err)
			result, err := conn.SendRecv([]byte(gconv.String(i)), -1)
			t.AssertNil(err)
			t.Assert(string(result), fmt.Sprintf(`> %d`, i))
			conn.Close()
		}
	})

	// gudp.Send
	gtest.C(t, func(t *gtest.T) {
		for i := 0; i < 100; i++ {
			err := gudp.Send(s.GetListenedAddress(), []byte(gconv.String(i)))
			t.AssertNil(err)
		}
	})
}

// If the read buffer size is less than the sent package size,
// the rest data would be dropped.
func Test_Buffer(t *testing.T) {
	var ctx = context.TODO()
	s := gudp.NewServer(gudp.FreePortAddress, func(conn *gudp.ServerConn) {
		defer conn.Close()
		for {
			data, remote, err := conn.Recv(-1)
			if len(data) > 0 {
				if err = conn.Send(data, remote); err != nil {
					glog.Error(ctx, err)
				}
			}
			if err != nil {
				break
			}
		}
	})
	go s.Run()
	defer s.Close()
	time.Sleep(100 * time.Millisecond)
	gtest.C(t, func(t *gtest.T) {
		result, err := gudp.SendRecv(s.GetListenedAddress(), []byte("123"), -1)
		t.AssertNil(err)
		t.Assert(string(result), "123")
	})
	gtest.C(t, func(t *gtest.T) {
		result, err := gudp.SendRecv(s.GetListenedAddress(), []byte("456"), -1)
		t.AssertNil(err)
		t.Assert(string(result), "456")
	})
}

func Test_NewConn(t *testing.T) {
	s := startUDPServer(gudp.FreePortAddress)

	gtest.C(t, func(t *gtest.T) {
		conn, err := gudp.NewClientConn(s.GetListenedAddress(), fmt.Sprintf("127.0.0.1:%d", gudp.MustGetFreePort()))
		t.AssertNil(err)
		conn.SetDeadline(time.Now().Add(time.Second))
		t.Assert(conn.Send(sendData), nil)
		conn.Close()
	})

	gtest.C(t, func(t *gtest.T) {
		conn, err := gudp.NewClientConn(s.GetListenedAddress(), fmt.Sprintf("127.0.0.1:%d", 99999))
		t.AssertNil(conn)
		t.AssertNE(err, nil)
	})

	gtest.C(t, func(t *gtest.T) {
		conn, err := gudp.NewClientConn(fmt.Sprintf("127.0.0.1:%d", 99999))
		t.AssertNil(conn)
		t.AssertNE(err, nil)
	})
}

func Test_GetFreePorts(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ports, err := gudp.GetFreePorts(2)
		t.AssertNil(err)
		t.AssertEQ(len(ports), 2)
	})
}

func Test_Server(t *testing.T) {
	var ctx = context.TODO()
	gudp.NewServer(gudp.FreePortAddress, func(conn *gudp.ServerConn) {
		defer conn.Close()
		for {
			data, remote, err := conn.Recv(-1)
			if len(data) > 0 {
				if err = conn.Send(data, remote); err != nil {
					glog.Error(ctx, err)
				}
			}
			if err != nil {
				break
			}
		}
	}, "GoFrameUDPServer")

	gtest.C(t, func(t *gtest.T) {
		server := gudp.GetServer("GoFrameUDPServer")
		t.AssertNE(server, nil)
		server = gudp.GetServer("TestUDPServer")
		t.AssertNE(server, nil)
		server.SetAddress("127.0.0.1:8888")
		server.SetHandler(func(conn *gudp.ServerConn) {
			_ = conn.Send([]byte("OtherHandle"), nil)
		})
	})
}
