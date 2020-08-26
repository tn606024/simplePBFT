package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

var urlName = "localhost:%d"

type Server struct {
	node *Node
	url	string
}

func nodeIdToPort(nodeId int) int{
	return nodeId + 8080
}

func NewServer(nodeId int) *Server{
	server :=  &Server{
		NewNode(nodeId),
		fmt.Sprintf(urlName, nodeIdToPort(nodeId)),
	}
	return server
}

func (s *Server) Start(){
	s.node.Start()
	ln, err := net.Listen("tcp", s.url)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Printf("server start at %s\n",s.url)
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn){
	req, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	s.node.msgQueue <- req
}
