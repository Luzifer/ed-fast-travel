// Code generated by go-bindata.
// sources:
// assets/application.coffee
// assets/application.js
// assets/frontend.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsApplicationCoffee = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x59\x5d\x6f\xeb\xb8\xd1\xbe\xd7\xaf\x98\x55\xce\xbe\x92\xdf\xda\x4a\xd2\x62\x6f\xd2\x38\xc5\x9e\xa0\x8b\x9e\x76\xb1\x5d\x9c\x74\xb1\x17\x07\x41\x40\x8b\x63\x89\x8d\x4c\xaa\x24\x15\xc7\x3d\xc8\x7f\x2f\xf8\x25\x51\xb2\x9c\x64\x2f\xda\xab\xc8\x9c\x4f\x0e\x67\x9e\x19\x32\x0a\x89\x2c\xeb\x7f\xb0\x1d\x4a\x58\x03\xef\x9a\x26\x51\x07\xa5\x71\x77\x4b\xca\x1a\xc3\x92\x14\x9d\x36\x3f\xbe\xdc\x27\x4a\x94\x8f\x61\x79\xcb\xa4\xd2\xb7\x82\x73\x2c\x35\xac\x41\xcb\x0e\x13\x80\xe4\x03\xac\x6e\x12\x80\x0d\xe3\xf4\x97\x4f\xf9\x22\x01\x68\x04\xa1\x77\xa8\x35\xe3\x95\xb2\x0b\xaa\x25\x7b\x7e\x27\xca\x47\xd4\xf9\x22\x49\x1c\x2b\xac\x21\x5f\x38\xd9\x0f\x79\x76\xa6\xbc\xc0\x4a\x91\x27\xcc\x16\x85\x61\x82\xac\x6c\x58\xf9\x98\x2d\xc1\x2c\x06\x95\x5e\x80\xf1\xb6\xd3\x77\xd6\xfb\x9e\xdd\xae\x65\x4b\xa8\x09\xa7\x0d\x3a\xe2\x27\xb3\xf6\xe7\x27\xe4\xfa\x15\xc1\x47\x3c\x74\xed\x44\xf0\x6f\x78\xf8\xa5\x8d\x05\x09\xa5\x2b\x17\xae\xd5\x46\xf3\x23\x1f\x63\xd9\xef\x29\xf5\x42\x5d\x4b\x89\x46\x2f\x77\x42\xe6\x17\xcb\x13\x02\x51\x77\x9a\x8a\x3d\x7f\x5d\xe4\xce\x73\x59\xa1\x3d\xe3\x54\xec\x03\x63\x4d\x54\x5d\xd6\x84\x57\xd8\x73\xff\x85\xa8\xfa\xd6\x2e\x25\x49\x7c\x3a\xd1\x19\xa8\x61\xa9\x11\x25\x69\xee\xb4\x90\xa4\xc2\xa2\x42\xfd\x49\xe3\x0e\xb2\xc0\x90\x19\x66\x58\x27\x00\x00\x4a\x8b\xf6\x41\x1a\xbd\x57\xf0\xdd\xc5\x85\x5b\xab\xc5\xfe\xa1\x14\x42\x52\xc6\x89\x46\x75\xe5\x32\x25\x01\x60\xdb\xc1\xca\x37\x3e\xa9\xac\x04\xac\xe1\xaf\x77\x7f\xff\xa9\x68\x89\x54\xd8\xf3\x24\x21\x20\x5a\xb4\x2b\x6b\x24\x5b\x14\x4f\xa4\x01\x55\x0c\x76\xfb\xa0\x89\xfd\xca\x1a\x55\xd9\xa2\x68\xa5\x68\x21\x2b\x6b\x2c\x1f\x91\x9a\xec\x29\xa6\x4e\x39\x6f\xbe\x99\x25\x58\x8d\x1b\x41\x0f\xd9\xa2\x20\x94\xde\x36\x44\x29\xc8\xb8\x08\xfa\x13\x00\x6c\x14\x4e\x38\x25\xee\xc4\x13\x1e\x33\x27\x71\xea\xc6\xf1\x9e\x09\xe1\xdc\x6e\x6d\x01\xcd\x45\xf5\xad\x6d\x27\xb6\x12\xa3\x93\x54\x47\x27\xb9\x74\x61\x57\x5a\x32\x5e\xb1\xed\x21\x57\x8b\x49\x35\x66\x8b\x62\x27\x28\x69\x20\xab\x19\xc5\x2c\xf9\x5f\xc6\x0d\xb5\x2d\xdd\x3b\x4d\x2c\x1a\xe5\xca\x7c\xf8\xe8\xb1\x52\x70\x27\xb3\x06\xb5\x67\xba\xac\xc1\x92\xad\xf2\x7d\x8d\xdc\xec\xd2\x60\x1d\xe3\x55\x06\xda\x2e\x6c\x09\x6c\xc9\x4a\xb5\x8c\x73\x94\xe6\xb3\xed\x1a\x85\x59\x24\x82\x52\x0a\x39\x66\xc7\xe7\xb2\x21\x3b\xa2\x99\xe0\x2b\x2d\x19\xe1\x55\x33\x12\x51\x5d\x59\xa2\x52\x63\x21\x7b\x04\x8e\xcb\xec\x38\x2c\xb7\xc8\x4b\xd6\x64\xc7\x68\x62\x76\x63\x82\xa6\xb5\x34\x75\x4e\x94\x39\x9a\x7e\x8b\x36\x14\xbf\x12\xc9\x19\xaf\x4c\x1c\xf6\xee\x33\x44\x62\x0b\xfb\x40\x5b\x43\x9a\x86\xf0\x9e\xf9\xd5\x6c\x51\x98\xa3\xb3\x69\x14\x87\x3f\xd0\x57\x1a\x9f\x75\xb6\x28\xcc\x9f\xa0\x69\x46\x87\x39\x6f\x83\xdd\xf3\xc8\x3a\xc6\x91\xe8\xd8\xb2\xbd\x64\x26\x91\x66\x36\xed\x20\xd4\xe5\x2c\x65\x8a\x6c\x1a\x5b\xab\x23\xb0\x18\xba\x55\x8c\x17\x0e\xee\x8a\xb2\x41\x22\x0d\x55\x74\x3a\xe6\x75\xf5\x72\xdc\xe9\x9c\xd2\xe3\x0e\x60\x8b\xcc\x18\xc8\xb2\x19\x51\x6f\x4c\xa1\x1e\x9b\x72\xe2\x1f\x0f\x3f\x91\x1d\x2e\xe1\xd2\xa0\x5f\x32\xdf\x3d\x4c\x70\xf0\x49\x0f\xe7\x85\x4f\xba\x78\xc4\xc3\xad\xa0\x68\xce\xec\xf2\x0f\x7d\xbc\x8f\xc2\xa3\x25\xab\x2a\x94\x01\xfd\x4d\x2a\x4c\x8d\x9f\x8e\xfd\x50\x01\x49\xd8\xd5\x9d\x2d\x75\x58\x9f\x0a\x83\xc5\x5c\x83\xf9\x90\x9e\x93\x96\x9d\x07\x67\x0e\x2b\x4e\x76\xf8\x27\xf7\xf3\xc1\x7c\xaf\xcf\xbe\xc6\x2a\x5f\xd2\x25\xe4\x94\x68\xe2\x5d\xb1\xce\x84\xa4\x35\xeb\x85\x2d\xae\x87\x1d\x2a\x45\x2a\x97\x85\x06\x49\x2c\xc9\x17\x91\x5d\x3c\xda\x85\x2b\x4a\x4f\x9b\x04\xa9\x92\xa2\x6b\x23\xb8\xc9\x4d\xf3\x5b\x39\x89\x45\x5f\x7f\xf3\x7a\x43\xe5\xbe\xae\x39\x82\xa7\x63\xe5\x73\xf3\x84\x69\x50\x76\x53\x6e\x67\x76\xb9\x30\x01\x9b\xb7\x73\xa2\x0a\xb6\x24\xf2\x7b\x34\xa1\x4d\x75\x8f\xb3\xee\x7b\x4a\xa3\x7c\xf0\xe9\xfe\x96\x39\x10\x72\x6c\x24\xaa\x35\x89\xba\x93\xdc\xd7\x8e\x1d\x0c\xbf\x5c\xdc\x1b\x8e\x2c\x03\xc2\xa9\x5b\x2a\x1a\xe4\x95\xae\x6d\x32\x3b\xa9\x61\x82\x8c\x25\x47\xcc\x2b\xb8\xbc\x37\x55\x17\x59\x2e\x18\x1d\xc4\x8b\xb6\x53\xf5\x1c\xd5\x34\x35\x83\xc7\x85\x19\x74\x60\xed\xb9\xff\x29\x18\xcf\xb3\x65\xb6\x48\x66\xc7\x3c\x73\x2c\x59\x36\x4b\xea\x6b\xcc\x4d\x8f\xc9\x34\xe2\x0e\x3c\xa2\x19\x36\x8a\x30\x72\xda\x0a\x66\x6b\xdc\xb7\xa1\xde\xbb\x56\x0a\x2d\x4a\xd1\x0c\xcd\x22\xad\xb5\x6e\xd5\x55\xea\x7a\x45\xba\x57\xea\xea\xfc\xfc\xec\xeb\xb0\x1f\xa1\xf4\x8b\x2d\x3b\xbb\xa5\x74\x68\x20\xe9\xfe\x2d\x56\xe3\xb5\x9f\xd5\x71\x0f\xbf\xe2\xc6\x8f\xdb\xc1\xc1\x85\x67\x28\x04\x2f\x1b\xa1\xcc\xbe\x8e\xeb\x34\xbb\x15\xbb\x5d\xc7\x99\x33\x02\x7b\xa6\x6b\x50\x28\x9f\x50\x9a\x8d\x22\x85\x8e\xe3\x73\x8b\xa5\x46\xda\x1c\x0a\xf8\xb9\x41\xa2\x10\x24\x9a\x91\x12\x74\xcd\x14\xb4\xa4\x42\xd0\x02\x4a\xc1\x35\xe3\x1d\xc2\x41\x74\x12\x4a\xd2\x94\x5d\xe3\x3c\xcf\x06\x47\x3c\x16\xc0\xda\x4f\xa9\x9f\xcd\x56\x3e\xa3\xea\x1a\x3d\x30\x89\x16\x79\x14\x70\x9b\x4d\xf1\x4d\xa4\x2f\xab\x30\x04\xf7\xc7\x19\xcd\xc1\x9e\x69\x72\x83\x71\x45\x96\x1c\x19\x1f\x03\xb6\xa9\xb5\xd1\x80\x6a\x68\xb6\x02\x17\x61\x1c\x3a\x02\xb1\xb7\xb1\x6f\xa8\xab\x76\x43\xe2\x16\xb5\x15\x12\xb0\x01\xc6\x4d\xa6\x6a\x59\x90\x4e\x8b\x15\xa1\x14\x69\x61\x38\x3d\xf2\xd8\xca\xc6\x66\x61\xfd\xc8\x33\x89\xff\xea\x50\xe9\x07\x46\x33\xdb\xc9\xd2\xb3\xaf\xd6\xae\x4d\x8e\x87\x81\xfa\x92\xfa\x48\x84\xf3\xb1\x3f\xbd\x07\x68\xed\x23\x65\xfa\xb3\xd8\xbf\xcb\xa3\xff\x86\x33\x13\x65\xbe\xe3\x1c\xe9\xb2\x07\x55\x28\x4d\xe4\x83\xc7\xd8\x53\x0a\x87\x0d\xa1\xab\xc5\x8d\x44\xf2\xe8\x8f\xae\x27\x46\x90\x37\x08\x54\xe8\x51\xe2\x47\xc6\xd1\x0f\xe2\x1f\x72\x4f\x5f\x14\x8c\x2b\x94\xfa\x23\x6e\x85\xc4\xdc\x04\x31\x24\x84\x75\xb1\x14\x1d\xd7\x28\xe1\x06\x2e\xa6\x82\x5b\xc6\x69\x9e\xb9\x5b\x0c\x17\x7e\xf8\xca\x63\x29\x37\x87\x1f\x09\x0c\xed\x77\x24\x34\x13\x0d\xc3\x33\xaf\x24\x1a\xda\x83\x92\x52\xec\x5a\xd6\xe0\xed\x40\x39\xa9\xd7\x8d\xe7\x8b\x79\xd5\xdb\x86\x55\xb5\x7e\xa0\x4c\x69\xc2\xcb\xde\xc7\xc9\xa1\x4d\xb8\x0a\x2d\x7e\x60\xcf\x48\xf3\xdf\x2f\x5e\xe0\xc7\x43\x3a\xaf\x5a\x0b\x4d\x9a\x87\xf7\x19\x98\xe5\x7d\x9f\x99\xb2\x61\x6d\x3f\x89\x1b\x9d\x2b\xb3\xb2\x11\x44\x52\x37\x28\x2f\xe1\xb5\x88\x27\x00\x86\xdf\xa3\xf0\x6d\x10\x3d\x69\xe8\xcb\xc5\xbd\x17\x29\x44\x74\x99\x58\x42\x8e\x3d\xde\x69\x51\x55\x0d\x7e\xec\xb4\x0e\x57\x1e\x0c\x18\xb7\x84\x6c\xa3\xf9\x8a\xe2\x96\x74\x8d\xf1\xcd\xfe\x1c\x0d\x36\x14\x1b\x72\xb0\xe3\xe9\x32\xc2\xd0\xf7\x68\x1d\x9c\x19\x19\x89\x53\xdc\x47\xa1\x95\xa2\x92\xa8\x14\x5c\xfb\xf6\xff\xc1\x55\x43\xd8\x6b\xa0\xaf\x2c\x7e\x15\xa5\x19\xa5\xf6\x8c\xea\x3a\x5b\x4e\x0b\xba\x57\xf5\xff\x70\x79\x71\xf1\xf2\x6d\x3a\xb9\xb6\x38\xbd\x6e\x28\x33\xd3\xea\x71\xe6\x1a\xf0\xf6\x49\xea\x76\x9b\xe6\x67\x5f\xdd\x42\xf1\xdc\x27\xc1\x77\x8b\x17\x38\x87\x9e\x70\x38\x45\xf8\x77\x4c\x58\xa4\xa1\x59\x0c\xef\x29\xe3\x57\xac\x11\x36\x0e\x6e\x0e\x03\xd1\x68\x76\x29\x54\xb7\x71\x77\xef\xfc\x72\x51\xa8\xb6\x61\xba\x9f\x60\xd8\xf6\x15\xde\xe8\xa6\x37\xf4\x11\xbd\x6b\xc3\x30\x64\x72\xca\x74\x31\x03\x3c\xfb\x9a\x35\x68\x88\x61\xee\xba\x09\x33\x5a\x0f\xc4\x60\x71\x55\xef\xda\x2f\x17\xf7\x2f\x57\x57\xee\xf3\xd2\x7d\x96\xfa\xc5\x0e\x17\x60\xee\xd7\x52\x9f\x04\x45\x4b\x75\xa0\x0d\x71\x07\x58\x46\x76\xe6\x58\x07\x7c\x5f\x82\xf3\x60\xcc\x15\x63\xec\x74\x76\xfb\x2c\xf6\x2e\x58\x7d\x0b\xab\x50\xff\xec\x33\xe8\x23\x91\xbd\x73\x2e\x6b\xde\xe3\x9b\xe3\x7c\xa7\x51\x4d\x64\x85\xa7\x43\xe2\xc8\xef\xb3\x3b\xe6\x9d\x06\xe5\xf2\x7e\xc2\xf6\x4e\x07\x77\xaa\xf2\xaf\x4c\xfe\xf8\x1e\x7a\xcd\x57\x60\x67\x99\x4f\x5c\xe7\x2e\xec\xe1\x46\xe3\x4c\x9c\x64\xbc\xec\x19\xa7\xed\xfc\x6a\xba\x2b\xff\xb6\x15\x00\xd8\x2b\xfa\xa1\x11\x44\xe7\xf3\x2f\x5d\xfe\xa9\xcb\x0c\x7e\x0a\x39\xcd\x27\xaf\x53\x3b\x55\x79\x16\x97\xeb\xd6\x9f\xa2\x70\xc1\xb1\x09\x5f\x6a\xf8\x1d\x5c\x26\xc9\x38\x0f\xc6\x35\x7a\xad\x25\xd8\x37\x96\x75\x3a\x54\xab\x4d\xa0\x34\x80\x23\x5c\x6b\x1a\x78\xa4\xe9\x22\xe9\xcd\xff\xf1\x8d\x6a\xff\x78\x7d\xae\x69\xcc\xd4\x7f\x03\x5c\x53\xf6\x14\x64\x02\x8a\xa5\x11\x7d\x9e\xc3\x40\x22\xc4\x3f\x56\x66\xbb\x2d\x52\x20\xa5\x66\x4f\x98\x82\x14\x0d\x0e\xfc\xc6\x4b\x50\xfa\x60\xd6\x2c\x84\x5e\xc1\xc5\xb7\xe9\xcd\xf5\x39\x65\x4f\xb1\x33\xa3\xdf\x53\xaf\xdf\xb5\xb5\xb7\x98\xae\xcf\xb5\xbc\x31\x89\x36\x4a\xfe\x77\x84\xfa\x74\x94\xc1\x0f\x44\x66\x3f\xa7\xe3\xbc\xb1\x3d\x2b\x08\x6e\x34\x87\xa8\x3d\xd9\xef\x67\x65\x5b\x6a\x7a\x73\xcd\x02\x9b\x7f\x97\x0b\x0d\x39\x05\x22\x19\x59\xd5\x8c\x52\xe4\xeb\x54\xcb\x0e\x8d\x51\x76\x73\x7d\xee\xd4\xc7\x06\x55\x4b\x7a\x73\xd1\x04\x66\x04\x0c\xe9\x14\x6b\x34\x67\x1d\xb1\xbe\x1a\x6c\x98\x4c\x2e\xc7\xe1\x18\xb3\xcf\x8e\x3b\x23\xa1\xfe\xac\xc6\xff\x3b\x88\x0f\xcb\xbe\xf9\x64\xf6\x46\x69\xa6\x67\x29\x9a\xf3\xf0\x8f\x88\xec\xb7\xbe\xec\x24\xf1\x7f\x35\xde\x30\xe2\xfe\x3d\xf2\x1b\x4d\x24\xc7\xd3\xcb\x1a\x72\x77\x72\x06\x5c\x4d\xdb\x5d\x02\xa1\xb4\x4f\x45\x47\x1b\xbf\x36\xbb\xef\x11\xb9\x7f\xb6\x26\x94\x26\x89\x9b\x9d\xd6\x90\x6b\xb6\xc3\x25\x6c\x1f\xc3\x95\xf0\xf8\x4d\x70\xfb\xa8\x97\x60\xd8\x92\xff\x04\x00\x00\xff\xff\x07\x19\x3e\x36\x63\x1b\x00\x00")

func assetsApplicationCoffeeBytes() ([]byte, error) {
	return bindataRead(
		_assetsApplicationCoffee,
		"assets/application.coffee",
	)
}

func assetsApplicationCoffee() (*asset, error) {
	bytes, err := assetsApplicationCoffeeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/application.coffee", size: 7011, mode: os.FileMode(436), modTime: time.Unix(1473767028, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsApplicationJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x3a\x5b\x73\xe3\xb6\xd5\xef\xfb\x2b\x4e\x94\x4d\x48\x7e\x4b\x51\xd2\xd7\xc9\x8b\x6d\xb9\x93\xf5\x34\xed\xb6\x99\x34\xb3\x6e\x26\x0f\x3b\x1e\x0f\x44\x1c\x49\xb0\x29\x80\x05\x40\xcb\x6a\xb3\xff\xbd\x03\x80\x04\xc1\x9b\xec\xed\x74\x3a\xf5\x83\x25\x02\xe7\x86\x73\x3f\xa0\x16\x0b\xf8\x23\x72\x94\x44\x23\x85\xcd\x09\x6e\xc4\x76\x8b\x78\x9b\x4b\x56\x6a\x58\x65\xab\x65\xb6\x7c\x13\x6f\x2b\x9e\x6b\x26\x78\x9c\xc0\x3f\xdf\x00\x3c\x11\x09\x1b\xc6\xe9\x2f\x1f\x52\xc8\xc5\xa1\x64\x05\xde\x08\x21\x29\xe3\x44\xa3\x4a\x81\x62\x41\x4e\x29\x6c\x99\x54\xfa\x46\x70\x8e\xb9\x4e\x61\x87\xfa\x67\x29\x76\x12\x95\x7a\x4f\xa4\x7d\xbe\x3d\x29\x8d\x87\x1f\x19\xc7\x14\xf6\x84\xd3\x02\xff\x44\xd4\xfe\x66\x4f\xf8\xce\xaf\x7c\x14\x95\xc6\x8f\xa8\xaa\x42\x37\x4b\xb7\xfb\x4a\x53\x71\xe4\xfe\xd9\x92\xf9\x9e\xd2\xee\xc2\x07\x5e\x56\xfa\x0f\x4f\xc8\x75\x77\xfd\x2f\x78\xfa\xa5\xec\xac\xff\x52\x52\xa2\x31\x85\x42\x10\x7a\x8b\x5a\x33\xbe\x53\x29\x48\xc3\x39\x05\x45\x9e\xb0\x5d\x54\x48\x64\xbe\x77\x84\xde\x9f\x7e\x22\x07\x6c\xd6\xfe\xc6\x0e\x28\xcd\x83\xb6\x8c\x6f\xb5\x25\xa9\x50\xff\x4a\x24\x67\x7c\x97\x82\x12\xf9\x63\x0a\xaa\x24\x47\x7e\x2b\xf2\x47\xd4\x29\x28\x4b\xe8\x86\xe4\x7b\x4c\x41\x8b\xdd\xae\xc0\xf7\x95\xd6\x82\xdf\x14\x44\xa9\xcb\x37\x6f\x20\xa4\x0e\x6b\xe0\x55\x51\xb8\xe5\x16\x33\x5c\xb6\x42\xc3\x1a\x3e\xdd\x39\x28\x91\x3f\x86\xdb\xa1\x49\x60\x0d\x5a\x56\x68\xd7\xdf\xf6\x4d\x0c\xb5\x81\xe3\xe4\xd2\x3e\x85\xaa\x69\xd6\x24\xea\x4a\xf2\xf0\x40\x6e\xe7\x73\x62\x69\x3a\x02\xb0\x86\x01\xe9\xb7\x71\xf4\xb5\xaa\x89\xcd\x8d\x7e\xa3\x24\x33\xd0\x71\x94\x17\x2c\x7f\x8c\xba\x4a\xaf\xb9\x19\x24\x66\x35\x6b\x4f\xee\x51\xec\x5a\x34\x65\xfa\x17\x90\x1f\xf1\x54\x95\xd1\x94\x7f\x04\xc8\x84\xd2\xb9\x53\xf9\x7c\xa3\xf9\x40\xde\x9e\x23\x06\x88\x95\xf5\xad\x1a\x77\x02\xcf\xf9\x5f\x80\xa4\x6a\x0f\x3f\x8f\xd6\xc4\x41\xd7\x1c\x6f\xe3\x23\xe3\x54\x1c\x1b\x9c\x3d\x51\xfb\xdc\x86\x54\x34\x8c\x32\x67\x2f\x6b\xae\xd0\xc2\x63\x46\x33\x41\x6f\x03\xc0\x81\x38\x9e\xaa\x45\x28\x44\x4e\x8a\x5b\x2d\x24\xd9\x61\xb6\x43\xfd\x41\xe3\x21\x8e\x1a\x80\xa8\x16\xd2\x40\x3a\x7a\x00\x4a\x8b\xf2\x5e\x1a\x31\x2e\xe0\xbb\xe5\x32\x6d\x96\xf7\xe2\x78\x9f\xb7\xf9\xe4\xc2\xba\xa9\xdd\xfc\xec\x88\xb0\x2d\xc4\x9e\xf3\x57\x6b\xe7\xdf\x49\x4b\x17\xd6\xf0\xe7\xdb\xbf\xfe\x94\x95\x44\x2a\xf4\x90\xb5\x04\x9f\x5b\x2d\x6b\x51\xce\x2d\xff\x28\xc9\x9e\x48\x11\xab\xac\x15\xa9\x63\x0d\x71\x9c\x5b\x81\x54\x94\x64\xa5\x14\x65\x1c\xe5\x7b\xcc\x1f\x91\x1a\x57\xcd\xfa\x02\x27\xad\x94\x5f\x8d\xec\x7a\x39\xbd\xc1\xa2\x8d\xa0\xa7\x28\xc9\x08\xa5\x36\xee\xe3\x88\x0b\xcf\xb0\x96\x1a\xb0\x50\x38\x8d\x2a\xf1\x20\x9e\x70\x12\xdb\x5b\x39\x0c\xac\x49\x2b\x9f\x37\xd5\xa8\xea\x92\x69\xeb\xbd\xa4\xc2\x24\x34\x6d\xc7\x8b\xd4\xc0\x8b\x52\x67\x58\xa5\x25\xe3\x3b\xb6\x3d\xc5\x2a\x09\x0d\xe5\x9d\x2d\x3b\x08\x4a\x8a\x38\xda\x33\x8a\xd1\xff\xb4\x3d\xc2\x72\x11\x1a\x44\x99\x85\xd0\x2a\x2c\xf7\x55\xc1\x9e\xa6\x79\x84\x35\x0c\x93\x37\x80\x3a\x32\x9d\xef\xa1\x4b\xc7\xfc\xe5\x44\x21\x44\xae\xaa\x30\xbe\x8b\x2e\xfc\x8e\x3f\x46\xb4\x25\xb0\x25\x73\x55\x32\xce\x51\x9a\xaf\x65\x55\x28\x8c\x2e\x7b\x44\x50\x4a\x21\xcf\x10\xc0\xe7\xbc\x20\x07\x62\x24\x9b\x6b\xc9\x08\xdf\x15\x43\x22\xaa\xca\x73\x54\xea\x0c\x19\xeb\x28\x01\x1e\xc5\x2d\xa9\x0a\x3d\x8d\x50\x22\xcf\x59\xe1\x31\x5c\xc4\x7f\x4e\xe2\x7e\xa2\xec\xa4\x75\xa3\x50\x63\x73\xad\xa5\x49\xb4\x44\x19\x6f\xf3\x5a\x0e\x12\x65\x5b\xd2\x43\x7b\x1d\xdd\x52\xa3\x69\xe3\x6c\xc7\x06\x6a\xbd\x86\xd9\x6c\xcc\xcf\xbe\xae\x41\xa2\x24\x33\x9e\x1a\x8f\x3b\x57\x00\x38\xd7\xf8\xac\xa3\x24\x33\x1f\x9e\xe5\xe5\x39\xba\xc6\xdd\xe3\x81\xdb\x8d\x17\xcb\xb1\x84\xd0\x71\xd0\x38\x3a\x4a\xa6\x2d\xe1\x33\xa5\xd1\x85\x37\x65\x8a\x6c\x0a\x9b\x22\x4d\xfe\x4e\xc2\xd4\xdd\x76\x34\xc3\xec\xed\xaa\x57\x96\x17\x48\xa4\x81\x11\x95\x0e\x31\xfc\x69\xc7\xfa\xa2\xd6\xda\x86\xcf\xb0\xe8\xdb\x4c\x65\x79\x46\xd1\xc0\x1e\x5d\x82\xb5\x18\x0a\x75\x57\x88\x6e\xeb\xb7\xfa\x6e\xb9\x3c\xab\xdc\xb6\x99\x08\x95\x8b\x4f\x3a\x74\x14\x7c\xd2\xd9\x23\x9e\x6e\x04\x45\xeb\x2c\xab\xdf\x8d\x3a\xcb\x40\xd1\x5a\xb2\xdd\x0e\x65\xd3\x17\x8c\x24\x97\xbe\xc4\x93\x29\xdf\x41\xda\xb4\x7a\x39\x66\xf7\x36\x5f\x34\x85\x3c\xc0\x80\xf5\x48\x83\x65\x75\xdd\x0d\x38\xd3\x16\xc4\xb3\x05\x29\xd9\xa2\x39\xc7\x69\xce\xc9\x01\x7f\xef\x1e\xef\xcd\xf7\xf5\x0c\xde\x75\xa8\xa7\xad\xc8\x94\x68\x12\xa4\x38\x1f\x88\x76\x23\xb3\xe9\xe8\xfe\x80\x4a\x11\x5f\xbf\xeb\xb4\x6f\xf7\xeb\x4c\x13\x26\xc3\xde\x29\x5d\x42\x4b\xda\x54\x33\xae\xfd\x9d\x14\x55\xd9\xa9\x0f\x7b\xa2\xe6\x3d\xe4\x5e\x10\x0f\x35\x5a\xe7\xbd\x80\xdb\x04\x9b\x4e\x2d\x19\xe1\x34\xd6\xde\x1a\xed\xdb\x43\xbb\x93\xdb\xe5\xcc\xa8\x77\x9a\xdd\x44\xf0\x6e\x49\xa1\x70\xa8\x92\xee\xfc\xd1\x67\xd5\x4f\xbd\x97\xa3\xc1\xf1\x3d\xa5\x63\x0e\xd9\xc4\xee\x4b\xb2\xc1\x6f\xbf\x75\xc5\x68\x72\x49\x3f\x7a\xfa\x89\xc1\x0e\x49\x9f\x96\x77\x16\x23\x8a\xe0\xdb\x6f\xdd\xdc\x94\x15\xc8\x77\x7a\xef\x62\x30\x20\x12\xcc\x54\x63\x84\x3a\xb8\x73\x58\xdd\xd9\x0c\x13\x48\x96\x31\xda\xa3\x96\x95\x95\xda\xc7\x3d\x90\x46\x69\xa6\x07\x32\x0a\xc9\x4c\xeb\x0e\xeb\x1a\xe3\x41\x30\x1e\x47\x69\x34\x68\x63\x87\x86\x8f\xa2\xc9\xb1\xc7\xa7\x0c\x37\x36\xf5\x26\xb9\x91\x99\xd2\xe7\x92\x76\xca\x9b\x4a\x22\xc8\x69\x29\x18\xd7\x8e\x66\xf3\xf4\x42\x7f\xe2\x0f\x5b\x4a\xa1\x45\x2e\x8a\x41\xaf\x32\xdb\x6b\x5d\xaa\x8b\xd9\x48\xbd\x9f\x1d\x95\xba\x58\x2c\x4c\xc6\x68\x95\x26\x94\x86\x77\xe0\xd2\x8c\x55\xdd\xec\x35\xad\xc3\xec\xf8\x25\x94\xfa\x2d\x45\x33\x67\xe3\x11\x7e\xc5\x4d\x3d\x0c\x37\x1a\x08\x60\x32\xc1\xf3\x42\xa8\xd1\x3c\x1c\x14\x22\x9f\xd7\xa2\x1b\x71\x38\x54\x9c\x39\x89\xe0\xc8\xf4\x1e\x14\xca\x27\xb4\xda\x46\x0a\x15\xc7\xe7\x12\x73\x8d\xb4\x38\x65\xf0\x73\x81\x46\x63\x12\xcd\x30\x07\x7a\xcf\x14\x94\x64\x87\xa0\x05\xe4\x82\x6b\xc6\x2b\x84\x93\xa8\x24\xe4\xa4\xc8\xab\xc2\x1d\xd3\x7b\x54\x47\xcc\x3a\x8d\xc2\x7a\x78\x0f\xd3\xf5\x19\x07\x2e\x4a\xe4\xe3\x87\x32\x71\x12\x5e\x3b\x84\xf6\x6d\x67\x55\xef\x97\xc1\xb8\x3a\x4c\x3a\xbd\xeb\x0b\x9b\x9a\x7a\x26\xe9\x67\x9a\x40\xee\x89\x0a\x6c\x5c\x37\x2f\x58\x99\xda\x2c\x96\x02\x52\xa6\x3f\x8a\x63\x0a\x58\xa4\xc0\x52\x78\x48\xa1\x40\x6e\xff\xad\x52\x28\x37\x44\xa6\x20\x71\x6b\xff\xad\x1c\x77\x83\xd8\x1d\x3a\x4d\x49\xb7\xd5\x2a\x98\x3d\xc6\x8b\xd0\xeb\xaa\xd8\x30\x91\x19\x41\x3a\x8d\x8f\xc4\xad\x2b\xc4\x5a\x66\xa4\xd2\x62\x4e\x28\x45\x9a\x19\xb8\x46\x95\x5b\x21\x21\x66\xb0\x86\xa5\x3d\x8e\xc9\x2d\xb8\xad\x33\xd7\x25\x30\xb8\x32\xab\x97\xc0\xde\xbd\x6b\xe5\xc3\xc2\x81\x7d\x62\x77\x61\x49\x7d\x1b\x63\x91\xd8\x23\xc6\x91\xc4\xbf\x57\xa8\xf4\x3d\xa3\x91\xeb\xae\xe2\x99\x09\x24\x7b\x1c\x1b\x37\xf7\x2d\x44\xd2\x89\xef\xda\x29\xbb\x36\xf4\x67\xc3\x4e\x4b\x57\xdb\xa5\x7f\xe6\xd5\xc8\xa1\x3b\xe7\x7d\xf0\xe7\x5d\xb9\x93\xac\xfc\x89\x1f\xdc\x89\x57\x97\xf0\x30\x76\xe4\xd5\xa7\x87\xff\xde\x99\xfb\x0c\xea\x96\x68\x94\xbe\x75\xe7\x4c\x69\x22\xef\xeb\xd2\xfe\x4a\x26\xad\x0e\x1b\xdd\x02\x6c\x24\x92\xc7\x7e\x65\xf3\x80\x83\xfe\xbc\x25\xd1\xb9\xc1\x8d\xbd\xab\xbe\x6d\x90\x93\x8c\x71\x85\x52\xbf\xc7\xad\x90\x18\x1b\xa3\x26\x7d\x3e\xf6\x3c\xb9\xa8\xb8\x46\x09\xd7\xb0\x4c\x82\xa9\xc7\x93\xd9\xda\x8b\x2c\x77\x39\xc3\x45\x33\xfe\x84\xa8\xbd\xa2\x38\xc0\x6c\xdb\xcb\x0e\xf6\x88\x22\x83\x1e\x69\x48\x27\xb8\x36\x68\xe8\x0c\xef\xbf\x27\x49\xbb\x3b\x80\x64\x92\xfa\xb6\x60\xbb\xbd\xbe\xa7\x4c\x69\xc2\x73\x2f\x69\x87\x5e\x0f\x26\xd3\xe2\x07\xf6\x8c\x34\xfe\xff\x24\x31\xa5\x0a\x7e\x3c\xcd\x26\xe9\x6b\xa1\x49\x71\xff\x1a\x2e\xa3\x90\x5f\xc4\xcb\x24\x54\x3f\x4d\x1b\xd2\x73\xb3\xb2\x11\x44\x52\x37\xc3\xa6\x93\xae\x1c\x58\xc0\xe0\xd4\x65\xf5\xa6\x41\x8f\xa7\x98\x7d\x5a\xde\x05\x68\x99\xe0\x6d\x8f\x1d\x8c\x10\xc1\x5d\xc8\xe0\xbe\x3d\xc6\xa6\x0e\xa5\x10\x6d\x34\x9f\xd7\xed\x42\x54\x3f\x0e\x5a\xf6\xba\x2c\xd9\xf7\x1d\xb1\x99\x05\xd3\xb1\x0a\xe8\xe1\x5e\x66\xd8\x0a\xdc\xe1\xdf\x0e\x15\x8d\xa3\x07\x75\x25\xd4\x63\x59\xbf\x61\x81\xab\x4e\x03\xdb\x4c\x31\x36\x04\x1b\xad\x35\xb0\x73\x5b\x20\xb2\xdc\x4c\x17\x47\x46\xf5\x3e\x4a\x27\x88\xfe\x1f\xac\x96\x4b\x6b\xfb\x6f\x66\x2f\xdc\x83\x39\x4e\x6e\x72\x19\x5e\x3f\x0c\x83\x26\xac\xcd\x75\x9c\xd4\x44\x9b\xfe\x2c\x36\xe9\xaf\xde\xcb\x9e\xbd\x33\x7e\x57\x3b\xe3\x02\xc2\xfd\xd3\x0b\xfb\xff\xe8\xef\x27\xb3\x7e\xe3\xd0\x5e\x92\x4f\xf5\xbb\xb9\x4e\xe1\xa0\x76\x6d\x4f\xd0\xe4\x7b\xf3\xdd\x68\x4e\xa5\x60\x1c\x5b\xa7\xa0\x89\xdc\xa1\xf9\x3c\x94\xbe\x33\xef\xd4\xab\x9e\xaa\x9a\x91\xa3\x33\x0a\x64\xaa\xda\xb8\x9b\xcf\x78\x95\x64\xaa\x2c\x98\x0e\x06\x02\xe3\x0a\xd3\xe0\x13\x97\x50\x61\xda\xd4\x87\xb2\x19\x36\xea\x30\x32\x1d\xd3\xb2\x29\xb4\xf6\x40\xc1\x14\x74\xdc\xb3\x02\x21\xd6\x87\xb2\x19\x7d\xae\xbb\x3e\xd7\x68\x03\xd6\x86\xb4\x19\xb7\xde\xc1\xec\xe2\xc2\x98\xc1\x3c\xaf\x82\xe7\x5c\xfb\x5b\x1d\xa3\xaf\x73\xb5\xc5\x02\x8c\x14\xe1\x50\xff\x53\xe0\x6d\x49\x4d\x6b\x91\x86\x90\x9d\x92\xd5\x9b\x9f\x3e\x8a\x63\x94\x78\x94\xba\x49\xe9\xbe\xd7\x0c\x45\x75\x31\xf0\x5a\x49\x1d\xf4\xeb\xd9\x3b\x97\x3a\xa7\x2a\x07\xf1\x7a\x09\xba\xf0\x7d\x65\xad\xee\x46\x40\x5f\x2f\xee\x41\xed\x82\xf7\x09\xb5\xa1\xef\x3d\x93\x0b\xb0\x7d\xf3\x07\xae\xe3\xda\x32\xa9\x07\x75\xcc\x26\x61\x57\x21\x6c\xbf\xf5\xba\x08\xc3\x32\x60\x2e\x4a\x5f\xd7\x6a\x72\x3f\x14\x82\xe8\x78\xfc\xe5\x46\xd2\xa4\x5f\xef\xa6\x66\xea\x51\xc8\x69\xdc\x7b\x1f\x71\x50\xbb\xc0\x44\x36\xa2\x4c\x84\xa8\x82\xe5\x18\xaf\x82\xc2\x61\xc3\xc9\x5d\x03\xd8\x38\xcb\xcd\xa8\xb9\xea\xf4\x31\x75\xe6\xab\x61\xdb\xfc\xd4\xf5\xb8\xb1\xec\xd4\x5e\x5b\x5d\x69\x33\xdc\x10\xa5\xd6\xb3\x36\xdd\x58\xd7\x9d\x5d\xc3\x95\xa6\xcd\xa6\x34\xa5\x7e\x76\xfd\x2d\xdf\xa8\xf2\xf2\x6a\xa1\xa9\xdd\xbd\x86\x2b\xca\x9e\x1a\x98\xa6\x0e\xcc\xc6\x97\x4d\x1d\x81\xf0\x61\x6e\xd4\x52\x22\x05\x92\x6b\xf6\x84\x33\x90\xa2\xc0\x16\xde\xc8\x00\x4a\x9f\xcc\x9a\xad\x3b\x17\xb0\xfc\x66\x76\x7d\xb5\xa0\xec\xe9\x1a\xfc\x47\x2d\xca\x79\x41\xcf\xed\x2e\xb4\xbc\x8e\x92\x8e\xfa\xda\x78\xf9\x37\xb4\x37\xa2\x38\xa8\x5b\x53\x23\x7d\xab\xba\x8d\x2d\xf4\x0d\xe0\x46\x73\x08\x6a\xba\xfd\xfe\xac\x6c\xaf\x32\xbb\xbe\x62\x0d\x58\xfd\xca\xa3\xe9\x76\x66\x40\x24\x23\xf3\x3d\xa3\x14\xf9\x7a\xa6\x65\x85\x86\x09\xbb\xbe\x5a\x38\xf2\xd7\x70\xa5\x4a\xe2\xb9\x04\x9d\xae\x81\x33\x5b\x3d\x88\xa0\x87\x0d\x20\x46\x15\x09\xbd\x0e\x30\x38\x5e\x17\x6e\xb4\x5f\xf4\xd0\x7d\x03\x74\xdf\x5e\x9f\xb3\x80\xbd\x30\x8e\xec\xfd\x8b\x99\x68\xa4\x28\x16\xcd\xab\xf1\x68\xf2\x5a\x78\x78\x8b\x32\x39\x57\x0f\x2f\x26\xdd\xab\xf8\x2f\x94\xc9\xbd\xe3\xff\x8f\x4b\x34\x68\x17\x43\xb1\x9c\xf5\x4d\x26\x37\x7d\x43\x0a\x84\xd2\xf6\xe7\x15\x6e\xb3\x7b\x73\xec\xbe\xf7\x5f\x81\x35\x90\xfe\x2a\x9b\xd4\xbf\x5e\x70\x22\xd8\xce\x36\x64\xab\xd9\x01\x53\xd8\x3e\xea\x9e\x5a\x86\xaf\x4b\xb6\x8f\xa6\xdf\x61\x87\xf6\x37\x06\x9f\x93\x2c\x27\x45\x11\xeb\x3d\x53\xc9\xe5\x9b\x7f\x05\x00\x00\xff\xff\x4b\x14\x7c\x3c\x7a\x24\x00\x00")

func assetsApplicationJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsApplicationJs,
		"assets/application.js",
	)
}

func assetsApplicationJs() (*asset, error) {
	bytes, err := assetsApplicationJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/application.js", size: 9338, mode: os.FileMode(436), modTime: time.Unix(1473767033, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsFrontendHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x6d\x73\xdb\x36\xf2\x7f\xed\x7c\x8a\x0d\x33\xd3\x26\x19\x93\xb4\xfc\x6c\x85\xd2\x7f\xf2\x77\x3c\xad\xaf\x4e\xae\x17\xa7\xd7\xeb\x74\xfa\x02\x24\x96\x22\x1c\x10\x60\x01\x50\xb2\xea\xc9\x77\xbf\x01\x48\x4a\x24\x25\xfa\x21\x9d\xeb\xbd\x39\xcf\xd8\x16\xc0\xc5\x3e\xef\x0f\x58\x50\xd1\xf3\x77\x7f\x3f\xff\xf4\xcb\x8f\x17\x90\x99\x9c\x4f\x9f\x45\xf6\x1f\x70\x22\x66\x13\x0f\x85\x37\x7d\x06\x10\x65\x48\xa8\xfd\x00\x10\xe5\x68\x08\x24\x19\x51\x1a\xcd\xc4\x2b\x4d\xea\x9f\x7a\xed\x47\x99\x31\x85\x8f\xbf\x97\x6c\x3e\xf1\xfe\xe5\xff\xf4\xd6\x3f\x97\x79\x41\x0c\x8b\x39\x7a\x90\x48\x61\x50\x98\x89\x77\x79\x31\x41\x3a\xc3\xce\x4a\x41\x72\x9c\x78\x73\x86\x8b\x42\x2a\xd3\x22\x5e\x30\x6a\xb2\x09\xc5\x39\x4b\xd0\x77\x83\x5d\x60\x82\x19\x46\xb8\xaf\x13\xc2\x71\x32\x6a\x18\x3d\xf7\x7d\xf8\x94\x21\x90\x58\xce\x11\x0e\xc0\x31\x36\x64\xa6\xe1\x75\x5e\x6a\xf3\x1a\x12\x99\x23\xa4\x4c\x69\x03\x4c\x80\xc9\x10\xac\x6d\x6f\x80\x88\x25\x48\x93\xa1\x72\xe3\x46\x36\xd8\x45\xd5\x9a\xd7\x24\x35\xa8\x5e\xdb\x25\x1a\x2b\x96\xbe\x5f\x4b\x35\xcc\x70\x9c\x22\xf5\x53\xa2\x8d\x6f\x14\x99\x23\x07\x1f\x2e\x38\x33\x38\x86\x77\x44\xcc\x50\xc9\x52\x03\x97\x62\xe6\x53\xa6\x0d\x11\x09\x42\x4d\x57\x70\x22\x04\xaa\x28\xac\xb8\x3c\x5b\x1b\xf2\xff\x52\x1a\x6d\x14\x29\xd6\x92\x38\x13\x9f\x21\x53\x98\x4e\xbc\x30\x4c\xa8\xb8\xd1\x41\xc2\x65\x49\x53\x4e\x14\x06\x89\xcc\x43\x72\x43\x6e\x43\xce\x62\x1d\x9a\x05\x33\x06\x95\x1f\x37\x6c\xc2\x83\xe0\x20\x38\x09\x13\xad\xc3\xd5\x5c\x90\x33\x11\x24\x5a\x7b\xa0\x90\x4f\x3c\x6d\x96\x1c\x75\x86\x68\x5a\x1e\xed\x4a\x75\x4b\x17\xc4\x24\x99\x93\x57\x90\x02\xd5\xa3\xf8\x6d\x37\xa2\xc7\x2e\xe5\xc4\xf0\xe5\x13\xf4\x7b\x8a\x47\x52\x29\x8c\x4f\x16\xa8\x65\x8e\xe1\x61\x70\x1c\x1c\x38\x67\xb4\xa7\xef\x93\x57\x09\x74\x53\x95\x70\x80\x58\xd2\xa5\x4d\xc6\xa2\x34\xbb\x10\x97\xc6\x48\x01\x77\xe0\x18\x6a\xf6\x07\x8e\x61\x74\x54\xdc\xbe\x81\x2f\x35\x79\xa0\xd8\x2c\x33\x70\x07\x06\x6f\x8d\x4f\x38\x9b\x89\x31\xb8\xb9\x16\x4d\x2a\x55\xee\xcf\x94\x2c\x0b\xb8\x83\x9c\xa8\x19\x13\x7e\x2c\x8d\x91\xf9\x18\xf6\x5a\x74\x89\x94\x8a\x32\x41\x0c\x6a\xb8\x83\x44\x72\xa9\xc6\xa0\x66\x31\x79\x79\x78\xb8\x0b\xc7\xfb\xbb\x70\xba\xb7\x0b\x7b\xc1\xe1\xab\xf5\x22\xab\x6f\x20\xa4\xef\xd6\xea\x3e\x0f\xca\x74\xc1\xc9\x72\x0c\x42\x0a\x6c\x49\x2a\x94\x9c\x29\xd4\x7a\xad\x8f\x91\xc5\x18\x8e\xad\x69\x7d\x05\xdb\xe6\xbe\x58\x10\x25\x98\x98\x0d\xb3\xae\x5c\x16\x24\x9c\xb5\x8c\x75\x0e\x19\x43\xc7\x73\x25\x0f\x04\x99\x4f\x39\x9b\x12\x6b\x6c\xa9\xb4\x54\xe3\x42\x32\x61\x50\x35\x54\x51\x58\xc7\x66\x5d\x42\xdf\x7f\x7a\x7f\x75\x04\x3a\x63\x39\x10\x41\xe1\x23\xea\x42\x0a\x1a\xdc\x68\x48\xa5\x82\xcb\x8b\x53\xd0\x65\x61\xe1\x06\x64\x5a\x13\x23\xc7\x1c\x85\xd1\x6e\x41\x8e\x94\x11\xf8\xbd\x44\xc5\xb0\x55\xf0\x96\xf5\xcf\x6f\x3f\x7e\xb8\xfc\xf0\xdd\xb8\xcd\x94\x4a\xd4\xe2\x5b\x03\x0b\xa9\x3e\x03\x4b\x61\x29\x4b\xb0\x80\xe6\x80\xa6\x20\x33\x84\x39\x23\x90\x32\x8e\xe3\x30\xec\xb0\xfb\x95\xa5\xc0\x0d\x5c\x5e\xc0\xd9\x6f\x4d\x72\x45\x3a\x51\xac\x30\xa0\x55\x32\xf1\x2c\xae\xea\x71\x18\x4a\xad\x83\x9c\xdc\x26\x54\xb8\xe4\xb6\x60\x7d\xa4\x33\x36\x0f\x0f\x82\x93\x60\x7f\x3d\x76\x89\x7c\xa3\xbd\x69\x14\x56\x6c\x9e\xc2\x55\x55\x26\x85\xa3\xe0\x30\xd8\x6f\x46\x03\x1c\xa3\xe7\xbf\xa2\xa0\x2c\xfd\xad\x32\x27\x0a\x9b\xcd\x22\xb2\xc9\x36\x7d\xb6\xb3\x13\x09\x32\x87\x84\x13\xad\x27\x9e\x20\xf3\x98\x28\xa8\xfe\xf9\x14\x53\x52\x72\x5b\x5b\x3b\x3b\x3b\x11\x65\x2b\x32\x8b\xbf\x84\x09\x54\x7e\xca\x4b\x46\x2b\x82\x9d\x0a\x15\x95\x0d\x8c\xfd\x35\x72\x36\xe3\x08\x33\x34\xe0\xca\x05\xa9\x8b\x6a\x8c\x16\xf6\x20\x97\x31\xe3\xd8\x24\x9e\x73\xf6\x4e\x4f\x4a\xad\x85\x55\x18\x55\x2d\x63\x27\xaa\xcb\xd8\x2c\x0b\x9c\x78\xd5\xc0\xeb\xad\xa8\x25\x27\x92\x73\x52\x68\xa4\x1e\x50\x62\x48\x3d\x6d\xb5\xaf\xe6\x9b\x69\xa2\x66\x76\xb3\x7c\x11\x6b\x1f\x6f\x49\x5e\x70\xf4\x6b\x46\x0d\xa5\x3f\xf2\x80\x28\x46\x7c\xbc\x2d\x88\xa0\x48\x27\x5e\x4a\xb8\xc6\x46\xa9\x9d\x48\x17\x44\x34\x5a\x68\xe5\x4b\xc1\x97\xde\xf4\x53\xa5\x87\x20\x73\x36\x23\x86\x49\x11\x85\x96\x6e\xeb\x22\x96\x48\xe1\xc7\x44\xb9\xf8\xfd\x27\x88\xa2\xb0\x72\x56\x33\x24\x3d\xa7\xc5\x36\x70\x5e\x6f\x9b\x8c\x42\x52\xc7\x25\xa4\x6c\x3e\x7d\xb6\x8e\xf3\xb9\xe4\x1c\x13\xe3\x4a\xc7\x26\x90\x05\x7a\xbd\x6b\x23\x9c\xeb\x5d\x17\xff\x6a\xbf\x6e\xb6\x6a\x1b\x7a\x17\x00\x0b\x38\xdb\xa2\xdd\xf8\x1a\x7a\xbe\xf7\x80\xd1\x89\x77\x6f\x6c\x1a\x93\x4a\xde\xb2\xa9\xe1\x23\xc8\x7c\x45\x10\x96\xfc\x21\xda\xe6\xa3\x43\xb9\x75\x7c\x39\x9b\x46\x04\xba\x79\x94\x4b\x4a\x78\x3f\x89\x34\x1a\xc3\xc4\x4c\x7b\xdd\x0c\x9d\x46\xac\x91\x97\x12\x48\x89\x3f\x43\xa2\xea\xa4\xca\x18\xa5\x28\x26\x9e\x51\x25\xda\xa0\xb1\x29\x5c\xd7\x5c\xac\xff\xa3\x90\xb3\x9e\x1e\xd5\x86\xfa\xa2\x72\x8d\x51\x44\x67\xbe\x92\xa5\xc1\x4d\x31\xee\xe1\xb0\x9c\x73\x8e\x44\xc1\x47\xbb\x76\x2d\x0a\x5a\x3f\x77\x77\x16\x28\x85\x34\x10\x50\xa6\x49\xcc\xf1\x5a\xa6\x66\x41\x14\x9e\x4b\x61\x94\xe4\xf0\xe5\x4b\xcf\x47\x56\xa7\xb2\xa0\xc4\xa0\xaf\x97\xda\x60\xbe\xa9\x15\x95\x0b\xc1\x25\xa1\xc3\x8a\xfd\xe4\x18\x40\x23\x6c\xc0\x0f\x4e\x96\xce\x4a\x63\x19\x0e\x4a\x2b\xe4\x02\x95\x2f\xd3\xf4\x1e\x7f\xd7\x3c\x06\x9d\x80\x82\x6e\x5a\xfa\x70\x36\x30\x91\xca\x5e\x26\x3c\xc6\x3f\xd5\xb2\x21\x65\x2f\x45\x2a\x7b\x0e\x59\x67\x76\x55\xa8\xae\x44\xc3\xa0\x57\x2a\x4d\xdd\x75\x68\x7a\x58\x5e\xd3\x44\xa1\xdd\xcf\x6d\xbd\x6f\xc5\xfd\xe6\xb8\x05\xd0\x7e\xac\xe4\xc2\x5b\x7b\xae\x57\xdc\x7e\x4e\xfd\xd1\x1e\xd4\x9f\x64\x9a\x6a\x34\xfe\xc8\x6b\x7b\xba\xbd\x82\x70\x54\x06\xdc\x5f\xbf\x3e\xa7\x78\xa0\xa4\xf5\xb4\x9b\xac\xfc\xd8\x3c\xe9\xc6\x2b\xd2\x46\x49\x31\x9b\xfe\x5c\x3d\x7d\x6e\x4f\x1e\x6e\x02\x2a\xa0\x6c\xad\xf4\xed\x51\x6f\x8d\x94\x2d\x1e\x15\xe0\x6d\x1d\x36\x60\xf8\x35\x1e\xd8\xef\x5a\x6c\x6c\x49\x35\x34\xd5\xc0\xfd\xf5\xb5\x51\xac\x40\xda\xb7\xcc\xa8\xee\x84\x9d\xca\xfa\x32\xa0\x06\xaf\x17\x51\x68\xb2\x87\xe9\x4f\xbd\xe9\xb5\x4b\xc4\x07\xc8\x6b\xae\xef\xea\x0e\xe9\x71\xd4\x9f\xa4\x21\x1c\x86\xd7\x44\xa1\x35\xa9\x3b\xf5\xdc\xf7\x1f\xb6\x9a\xf6\x04\x7d\x23\x62\x5d\xbc\x89\x42\x43\xb7\xd0\x4e\xff\x89\x8c\xc3\xcf\xa8\x0d\x5c\x63\x62\xa4\x82\x77\x57\xfe\x2f\x40\x8f\x4f\x07\x16\x3c\x85\xf9\x23\x69\xb7\x59\xfa\x18\xc3\x46\x43\x62\x5d\x47\x3a\xf1\xce\x39\x4b\x3e\x83\x91\x90\xc8\x62\x09\x15\xa4\xb8\x6e\xdd\x5b\xc7\xb9\x58\xae\xb0\xa6\xc7\x08\xa0\x73\x78\x70\xa4\x55\x4d\xfc\xa8\xa4\x44\x78\xa7\x96\x08\xef\x2f\xfc\x1f\x20\x19\x8d\xfc\xd1\xd1\x66\xa5\x6c\xe5\xb2\xea\x5b\xbc\xe9\x4b\x7f\x74\xba\x77\x12\xec\xd9\x1f\x08\xc1\x1f\x9d\x1e\x04\xa7\x27\x47\x6e\x70\xb2\x77\x16\x1c\xd9\x07\xaf\xb6\xf3\x7d\x9c\xc7\x8f\xf6\x46\xc1\xf1\x08\xae\x96\x7f\x9a\xfc\x6b\x63\xb4\xff\xdf\x8f\xd1\xc5\xf7\xfe\x4f\x80\x07\xfe\xe8\xf8\x6b\x42\xb4\xbf\x3f\x3a\x0e\x46\x47\xc7\xfb\x47\x2e\x44\xc7\x67\xc1\xc9\xe9\xc8\x0d\xce\x4e\x4f\x82\x83\xc3\x83\x93\xa3\x3f\x15\xa2\xc3\xb3\xc3\xe0\xe4\xf4\xd1\x21\x3a\x3b\x3b\x0e\x0e\xce\xee\x09\x51\x7b\xc6\xee\x58\xfd\x90\x39\x8c\x77\x8d\x7f\x05\x6d\x1f\x3b\xc0\xbc\x12\xfb\x00\x6a\x6c\xba\xb0\x85\xe8\xee\x22\x20\x93\x8a\xfd\x61\x77\x45\xbe\x25\x6a\x5b\xe8\x5d\x27\x54\xed\x5d\x84\xd2\x3a\xe0\xf5\xec\x36\x06\x5d\x16\xce\xa0\x7b\xa9\x07\xe9\x7d\x42\xa9\x14\x83\xab\x00\x36\x4e\x4d\x28\x12\xc6\x37\x54\xb5\xdd\xc5\xe0\xd1\x64\x48\xa5\xee\x5e\xda\x97\x6b\x95\xac\xcf\x48\x2e\xab\x3b\x0e\x4b\xaa\x63\xa6\xd7\x0f\xa8\x07\x05\x27\x09\x66\x92\x53\x54\x13\xaf\x9a\x84\x0f\xb6\xa4\x9e\xea\x99\xd8\xdc\xeb\x97\xba\xe1\xec\xf9\xc1\x2e\x6a\x98\xc5\x46\x40\x6c\xc4\xaa\x5f\x86\xfa\xa8\x4c\xa7\x6f\x29\x5d\xf7\x5d\x4f\x76\xce\xe0\xa3\x81\x07\x5b\xa7\x87\x72\xfb\xfe\xcc\xbf\x7f\x13\xeb\x8c\xad\xa5\xfd\x83\x12\xac\x0e\x97\xf5\xa9\x67\x75\x9b\xd2\xa7\x50\x72\xb1\x2e\xe0\xfe\xe2\xfa\xb4\x59\x11\x74\x8f\xa2\xee\xb8\x0d\x29\xa1\xe8\xd9\x13\x13\x13\x14\x6f\x27\x9e\x6d\xd4\xab\x43\x22\x65\x84\xcb\x59\x9d\xa9\x9c\xc4\xc8\x39\xd2\x78\x39\xf1\x66\x8a\xd1\x2a\x5b\xde\x5b\x16\x57\xf6\x51\x95\x5d\xab\xd6\x6d\xe3\xbe\xc3\x09\xf3\x2b\x96\x50\x0d\x74\xbe\x92\x24\x93\x32\x47\xd1\x34\x8b\x9b\xeb\xea\x16\x78\xd5\x85\x6e\x10\x74\x6f\x38\xee\xbf\xe2\x48\xb8\x5c\xdd\x5c\x50\xa6\x73\xb6\x62\xd3\xb6\xd5\xee\x38\x96\x6e\x5a\x61\xfe\x96\x72\xfd\xc6\xb0\x1c\xf5\x9b\x1a\xd3\x7b\x97\x03\x3b\x51\x76\xd8\x55\xd1\xed\x63\x95\x9b\xb6\x3a\x70\xba\xee\x58\xb3\xc3\x55\x53\xe2\xd2\x71\xc0\xe8\x58\xd2\x65\xff\x74\x6b\x89\xda\x91\xf0\x2d\x04\x74\xf0\xc0\xeb\x41\xfd\x20\xc4\x6e\xa9\x0e\xe7\x19\x48\xa5\x9a\x78\xda\xc8\xc2\x57\x44\xcc\x70\x7d\x9e\x85\x18\xcd\x02\x51\x80\x7d\xa8\xa3\xd0\x91\xdf\xbf\x05\x3c\x84\xc7\x1d\x6c\x13\x65\x1e\xa3\xba\x07\xdd\x5a\x4a\xc1\x9c\xf0\x12\x27\xde\xd1\xde\xde\x83\xfb\xca\x16\x90\xb7\xfb\xe6\x53\x10\xa2\xdd\xce\x6c\x13\x91\x64\x98\x7c\x8e\xe5\xed\xa0\x57\x1f\x34\x7e\xc5\xa1\x6e\xd9\xe5\xa2\xbe\xe4\xf6\x6c\xf3\x2d\x17\xd0\x3a\x8f\x6c\xd1\x7a\xab\x90\x6d\x7a\x77\x72\xae\x6f\xec\x66\x12\xa6\x52\x1a\x57\x79\x1d\x1e\xf7\xd5\x5f\x83\xf5\x85\x62\x39\x51\xcb\x2e\x72\xf8\x9a\xcc\xd1\x9b\x5e\x93\x39\x6e\x83\xfd\x96\x72\x9d\x16\xbc\x03\x13\x5b\x9b\xf4\x0e\x00\x35\x1d\x7a\xff\xf9\x5f\x04\x93\xee\x72\xe2\x5e\x88\xfc\x1f\x32\xb6\x91\xf1\x52\xd8\x52\xaf\xef\x7f\xff\x04\x38\x16\x8d\x0a\xeb\x0b\x8e\xde\x3b\xcd\xbb\x3b\x08\xe6\xa8\x34\x93\x02\xbe\x7c\x59\xdd\x7a\xf4\x0a\xa4\xe8\x4d\x94\x9b\x85\xc5\xd9\xf4\x8a\x69\x83\x14\x98\x80\xd5\x7d\x63\x66\x4c\x31\x0e\x43\xa4\x89\xa4\x78\x1b\xd8\x3c\x08\xff\x2f\x9f\x18\x29\xb9\xfe\x06\x85\x51\xcb\xc9\xc1\xd9\x99\x37\xbd\x78\x77\x6e\x09\xdc\xdd\xf1\x26\xe3\xfa\xa4\x66\xc3\x04\xa9\x92\x79\x97\xbf\x1e\x87\xe1\x62\xb1\x08\x90\xea\x3c\x10\x68\x42\xcb\xee\xfa\xfd\x00\xaf\x8f\xc8\x91\x68\xa4\x50\x0a\x8a\x6a\x93\xd3\x8c\x99\xac\x8c\xdd\x3b\x93\xab\xf2\x0f\x96\xa2\x0a\xbb\x2e\x0b\x63\x2e\xe3\x30\x27\xda\xa0\x0a\xaf\x2e\xcf\x2f\x3e\x5c\x5f\x78\xd3\xb7\x05\x49\x32\x84\xfd\x60\x0f\xae\x58\x82\x42\xe3\x90\x2d\xb2\x54\x09\x82\xf5\x07\x90\x39\x61\xdc\x5d\xe0\x48\xf1\x15\x9a\x78\xd3\xef\x98\xf9\xbe\x8c\x07\x24\xbd\x27\xac\x3e\x0c\x8d\x37\x99\xf3\x8a\x63\xc0\x64\xe8\x4d\x7f\x10\xa5\x81\xb7\x19\x47\xe5\x6e\x8f\xe1\xe5\xf9\xfb\x77\x1f\xa1\x16\x5a\xea\x57\xbd\x6c\x58\xdf\x88\xff\x15\xe0\x04\xcd\xeb\xb9\x9b\x7f\x94\xa8\x96\xf0\x52\x60\x82\x5a\x13\xb5\x74\x2f\x08\x56\xaf\xd4\xbf\xd5\xf0\x37\x32\x27\xd7\xd5\xbb\xb0\x82\x97\x33\x26\xf4\xab\xf5\x2b\xb9\xf6\x4b\xb2\x07\xdf\x2b\xdf\xfc\x6e\x65\x85\xa3\x60\xb4\x1f\x1c\xd6\xa3\xc1\xb7\x65\xbe\x0f\x97\x22\xe1\xa5\x8d\x28\xe7\x90\xc8\xbc\x60\x1c\x69\xa3\x03\xbc\x8c\x91\xcb\xc5\xab\x5d\x90\x0a\x58\x4d\xc8\x04\x65\x73\x46\x4b\x8b\xb4\x8c\xa3\x06\xa2\x41\x20\x52\xa4\x5f\xa9\xf1\xd0\x77\x03\x6e\xfa\x5f\x0d\xe8\x5a\xf0\x15\x92\x12\xce\x8a\x58\x12\x45\x83\x1b\x1d\x8e\x82\xa3\x60\xb4\xdf\x9a\x1b\x70\x52\x47\x02\xd1\x1a\x8d\x0e\x49\x51\x70\x96\x38\x94\xdb\x58\x12\x85\xd5\xbb\xc6\x28\xac\xbe\xc3\xf2\xef\x00\x00\x00\xff\xff\xc1\x46\xb2\x80\xd4\x22\x00\x00")

func assetsFrontendHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsFrontendHtml,
		"assets/frontend.html",
	)
}

func assetsFrontendHtml() (*asset, error) {
	bytes, err := assetsFrontendHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/frontend.html", size: 8916, mode: os.FileMode(436), modTime: time.Unix(1473764711, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/application.coffee": assetsApplicationCoffee,
	"assets/application.js": assetsApplicationJs,
	"assets/frontend.html": assetsFrontendHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"application.coffee": &bintree{assetsApplicationCoffee, map[string]*bintree{}},
		"application.js": &bintree{assetsApplicationJs, map[string]*bintree{}},
		"frontend.html": &bintree{assetsFrontendHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

