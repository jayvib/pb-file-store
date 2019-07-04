package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/jayvib/readfile-tutorial/pb"
)

// This practice demonstrates how to use protocol buffer
// to be the data serialization.
//
// I will use it for data persistence.

func main() {
	file, err := os.Create("test.pb")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileDB := &db{
		offsets: make(map[int64]int64),
		f:       file,
	}

	article := &pb.Article{
		Id:      1,
		Title:   "Item 1",
		Content: "This is a testing",
	}

	err = fileDB.Set(article)
	if err != nil {
		log.Fatal(err)
	}

	articleResult, err := fileDB.Get(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(articleResult.Title)
}

type db struct {
	// lock protects the offsets
	lock sync.RWMutex
	// offsets is use for caching
	offsets map[int64]int64

	f io.ReadWriteSeeker
}

func (d *db) pbAppend(article *pb.Article) (int64, error) {
	entityBytes, err := proto.Marshal(article)
	if err != nil {
		return 0, fmt.Errorf("pb marshall error %v", err)
	}
	byteBuffer := writeBinaryBufferLength(entityBytes)
	// get the offset number of the EOF
	offset, err := d.f.Seek(0, 2)
	if err != nil {
		return 0, fmt.Errorf("file seek error %v", err)
	}
	// write the bytes to the buffered
	_, err = byteBuffer.Write(entityBytes)
	if err != nil {
		return 0, fmt.Errorf("buffer write error %v", err)
	}
	// write to the file
	if err != nil {
		return 0, fmt.Errorf("file write error %v", err)
	}
	return offset, nil
}

// writeBinaryBufferLength writes the binary representation
// to bytes.Buffer with the length of b being used.
func writeBinaryBufferLength(b []byte) *bytes.Buffer {
	var length = uint64(len(b))
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, length)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}

func (d *db) Set(article *pb.Article) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	// what .pbAppend do?
	offset, err := d.pbAppend(article)
	if err != nil {
		return err
	}
	d.offsets[article.Id] = offset
	return nil
}

var ErrNotFound = errors.New("item not found")

func (d *db) Get(id int64) (*pb.Article, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	offset, ok := d.offsets[id]
	if !ok {
		return nil, ErrNotFound
	}
	_, err := d.f.Seek(offset, 0)
	if err != nil {
		return nil, err
	}
	size, err := d.readSize()
	if err != nil {
		return nil, fmt.Errorf("read size errror %v", err)
	}
	entity, err := d.readPbData(size)
	if err != nil {
		return nil, fmt.Errorf("key readData error %v", err)
	}
	return entity, nil
}

func (d *db) readSize() (uint64, error) {
	intSize := 8
	byteBuffer := make([]byte, intSize)
	_, err := d.f.Read(byteBuffer)
	if err != nil {
		return 0, err
	}
	reader := bytes.NewReader(byteBuffer)
	var readSize uint64
	err = binary.Read(reader, binary.LittleEndian, &readSize)
	if err != nil {
		return 0, err
	}
	return readSize, nil
}

func (d *db) readPbData(size uint64) (*pb.Article, error) {
	dataBuf := make([]byte, size)
	_, err := d.f.Read(dataBuf)
	if err != nil {
		return nil, err
	}
	var article pb.Article
	err = proto.Unmarshal(dataBuf, &article)
	if err != nil {
		return nil, err
	}
	return &article, nil

	return nil, nil
}

func (d *db) Delete(id string) error {
	return nil
}
