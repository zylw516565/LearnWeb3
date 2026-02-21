package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

const protocol = "tcp"
const nodeVersion = 1
const commandLength = 12

type version struct {
	Version    int
	BestHeight int
	AddrFrom   string
}

type getblocks struct {
	AddrFrom string
}

type inv struct {
	AddrFrom string
	Type     string
	Items    [][]byte
}

var nodeAddress string
var knownNodes = []string{"localhost:3000"}

func sendVersion(addr string, bc *BlockChain) {
	bestHeight := bc.GetBestHeight()
	payload := gobEncode(version{nodeVersion, bestHeight, nodeAddress})
	request := append(commandToBytes("version"), payload...)

	sendData(addr, request)
}

func sendGetBlocks(addr string) {
	payload := gobEncode(getblocks{nodeAddress})
	request := append(commandToBytes("getblocks"), payload...)

	sendData(addr, request)
}

func sendInv(address, kind string, items [][]byte) {
	payload := gobEncode(inv{nodeAddress, kind, items})
	request := append(commandToBytes("inv"), payload...)

	sendData(address, request)
}

func handleConnection(conn net.Conn, bc *BlockChain) {
	request, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Panic(err)
	}

	command := bytesToCommand(request[:commandLength])
	fmt.Printf("Received %s command\n", command)

	switch command {
	case "version":
		handleVersion(request, bc)
	case "getblocks":
		handleGetBlocks(request, bc)
	case "inv":
		handleInv(request, bc)

	default:
		fmt.Println("Unknown command!")
	}

}

func handleInv(request []byte, bc *BlockChain) {
	var buff bytes.Buffer
	var payload inv

	buff.Write(request[commandLength:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Recevied inventory with %d %s\n", len(payload.Items), payload.Type)

	if "block" == payload.Type {

	} else if "tx" == payload.Type {

	}

}

func handleVersion(request []byte, bc *BlockChain) {
	var buff bytes.Buffer
	var payload version

	buff.Write(request[commandLength:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	myBestHeight := bc.GetBestHeight()

	if myBestHeight > payload.BestHeight {
		sendVersion(payload.AddrFrom, bc)
	} else if myBestHeight < payload.BestHeight {
		sendGetBlocks(payload.AddrFrom)
	}

	if !nodeIsKnown(payload.AddrFrom) {
		knownNodes = append(knownNodes, payload.AddrFrom)
	}
}

func handleGetBlocks(request []byte, bc *BlockChain) {
	var buff bytes.Buffer
	var payload getblocks

	buff.Write(request[commandLength:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blocks := bc.GetBlockHashes()
	sendInv(payload.AddrFrom, "block", blocks)
}

func nodeIsKnown(addr string) bool {
	for _, node := range knownNodes {
		if node == addr {
			return true
		}
	}

	return false
}

func sendData(addr string, data []byte) {
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		fmt.Printf("%s is not available\n", addr)
		var updatedNodes []string

		for _, node := range knownNodes {
			if node != addr {
				updatedNodes = append(updatedNodes, node)
			}
		}

		knownNodes = updatedNodes

		return
	}
	defer conn.Close()

	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Panic(err)
	}
}

func gobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func commandToBytes(cmd string) []byte {
	var result [commandLength]byte
	for i, c := range cmd {
		result[i] = byte(c)
	}

	return result[:]
}

func bytesToCommand(bytes []byte) string {
	var command []byte
	for _, b := range bytes {
		if b != 0x0 {
			command = append(command, b)
		}
	}

	return fmt.Sprintf("%s", command)
}

func StartServer(nodeID, minerAddress string) {
	nodeAddress = fmt.Sprintf("localhost:%s", nodeID)
	ln, err := net.Listen(protocol, nodeAddress)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	bc := NewBlockchain(nodeID)

	if nodeAddress != knownNodes[0] {
		sendVersion(knownNodes[0], bc)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handleConnection(conn, bc)
	}
}
