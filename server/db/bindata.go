// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../migrations/0000_init.down.sql (0)
// ../migrations/0000_init.up.sql (8.419kB)
// ../migrations/0001_world.down.sql (0)
// ../migrations/0001_world.up.sql (4.893kB)

package db

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var __0000_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _0000_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0000_initDownSql,
		"0000_init.down.sql",
	)
}

func _0000_initDownSql() (*asset, error) {
	bytes, err := _0000_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0000_init.down.sql", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1642271179, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __0000_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6f\x6f\x9b\xc8\x13\x7e\xcf\xa7\x18\x59\x91\x80\x9f\xa8\xd4\x46\x3f\xa9\xba\x54\xad\x64\xe5\x48\x2e\xba\x9c\x7d\x87\x9d\xfb\xa3\xd3\x09\xad\x61\xed\x6c\x03\x8b\xcb\x2e\x49\x7d\x6d\xbf\xfb\x69\x77\xf9\xb3\x60\xc0\x70\x67\xa7\x6f\xca\xab\x18\x76\xe6\x99\x99\x9d\x79\x9e\x85\x04\x29\x46\x1c\x03\x47\xab\x08\xc3\x24\x48\xe2\x38\xa1\x3e\x0a\x02\xcc\xd8\xc4\xb0\x0c\x00\x00\x8a\x62\x0c\xea\xe2\xf8\x23\x37\xf2\xbf\x21\x48\x28\xe3\x29\x22\x94\x43\xcd\xce\xdf\x3e\xc0\x36\x25\x31\x4a\x77\xf0\x80\x77\x8e\x34\x50\x38\xa1\x8f\x38\x70\x12\x63\xc6\x51\xbc\x85\x27\xc2\xef\x93\x4c\xdd\x81\xbf\x13\x8a\x21\xc4\x6b\x94\x45\x1c\x2c\x9a\x3c\x59\x36\x20\xfd\x99\x99\xf1\xc0\xb4\x81\x26\x1c\x68\x16\x45\x86\xfd\xc6\x30\xea\xe1\x67\x0c\xa7\x45\xd4\xe2\x6f\x9f\x84\x45\xb0\x2b\xb2\x21\xb4\x35\x76\xb9\x70\xfb\x50\x3e\x12\xd7\x5e\xf8\x59\x56\xb9\xda\xff\x59\xbb\x8a\xf8\xea\x79\xaf\x76\x5a\x18\x1d\x86\x7b\x46\x48\xad\xec\x2b\x58\x79\x8d\xa8\x9c\x02\x79\xc4\x29\x23\x09\xad\x17\xa8\x3f\xa5\x0a\xe4\xa5\xad\x9c\xe0\xf5\x1a\x07\x9c\x3c\x62\x3f\x94\xfb\x70\x92\x48\x31\x0d\x95\x7b\x38\x58\x8e\xca\x88\x30\x1f\xc9\xc8\xf2\xf4\x92\x24\xc2\x88\x0e\x4c\x8f\xa7\x19\xce\x33\xac\xb5\xb6\x9a\x80\x61\x3b\x8f\x63\x44\xa2\xea\xe1\x70\x43\x6d\xde\xc6\x19\xca\x4e\xe6\x68\xc3\x2a\xc3\x3f\xff\x1a\x62\xc8\x76\x8c\xe3\xb8\x34\x1d\x66\x28\x2d\x2f\xe7\xb3\xc5\xd2\x9b\xde\xcc\x96\xb0\x7e\xf0\x65\x00\x79\xf3\xae\x76\xe5\x44\x5d\xcd\x3d\xf7\xe6\x7a\x06\x3f\xba\x7f\x80\x55\x0d\x84\x5d\x1b\x39\xcf\xbd\x72\x3d\x77\x76\xe9\x2e\xf2\x39\x06\x2b\x1f\xe1\x7c\x23\xda\xa0\xe4\xde\xa8\xad\xe9\x40\xd3\x77\xaf\x13\xb0\xbe\xc7\x96\xd8\x00\x5b\xa7\x97\x8c\x92\x0f\x19\x06\x42\x43\xfc\x51\x55\x59\x6e\xae\x9f\xa9\x3b\x09\x2d\x43\x96\xf7\x7b\x2d\x05\x75\xb4\x18\x8a\xdb\xfb\x84\x96\x26\x11\xf6\xf9\x6e\x8b\x0b\x56\x2b\x6f\xf8\x2d\x04\xd4\xcd\x71\x95\xd9\x21\xa2\x6b\x50\x4f\xfd\x3a\x12\x73\xb7\xb6\x79\x03\xa9\xa7\xe9\x07\x5c\x75\xa4\x38\x8b\x38\xd9\x46\xd8\x47\x8c\x91\x0d\x8d\x31\xe5\xcc\x47\x51\x94\x3c\xe1\xb0\x97\x18\x46\x23\xe5\x83\x14\x23\x8a\x36\xb8\xb9\x41\x47\x42\xea\x6c\x4c\xb1\xc9\x72\x8f\x45\x65\xf5\x16\xab\xba\x28\xef\xed\xd6\x3e\xab\xb5\xd8\x10\xe1\x94\x0b\x87\xf6\xd3\x00\xfd\x1b\x52\x83\x7f\x23\x91\xe3\xf5\x66\xb8\xa8\x9d\x52\xcb\x46\xd5\x63\x84\xdc\x0d\xf1\xdb\xa7\x88\x23\xf4\x69\x04\x94\xd3\xc6\x6f\xaa\x5d\xd4\xa3\x4a\xd8\x4a\xd8\x2e\x91\x1a\x05\xdb\x22\x2e\x32\x8a\x62\x9c\xda\x85\x45\x0f\xb4\x53\x57\xca\x45\x8d\xf5\x6d\x82\x26\x17\x1c\x5f\x3b\x5b\x4e\xc9\x12\xa9\xa2\xc2\xda\xdc\x57\xb7\xc5\x06\x1c\x98\x7d\x6d\xf1\x48\x1a\x80\x13\x52\x01\x9c\x98\x0e\x8e\xef\xbf\x4e\x09\x87\xfc\x8f\xaa\x4d\x83\x16\xe0\xc4\xd4\xd0\xd0\x0e\x38\xc2\x36\xef\x9f\x70\x4f\xe8\xbf\xc1\x31\xf0\x0c\x3c\xa3\x0d\x91\xf8\xd9\xc3\x36\x07\x88\xa6\x5a\xd5\x45\x2f\x1a\x94\x28\x64\x3b\x54\x49\x1c\x1d\x50\x43\x0e\xe7\x0d\xb0\x67\x64\xb5\x80\x93\x84\x8e\xf8\x70\xa1\x59\x7d\xbd\xcf\x16\x32\x08\x96\xad\xde\xe3\x80\x8f\x8d\x3d\x37\xfb\xca\xc1\x6f\x71\x1a\x13\xc6\xb4\xda\x37\x6e\x0f\x50\x13\x6d\xf1\x37\x35\xf9\xa6\x26\xcf\xa0\x26\x6a\xee\x6b\x0b\x8e\x7a\x98\xcd\x67\xf3\xd4\xfe\xf5\x12\x1d\x6f\x0c\x12\x1a\x12\x51\x9e\x4a\x0f\xdf\xb3\x84\xae\x8e\x15\xfb\x9a\xe0\x28\xac\x6b\xed\x7f\x55\xda\x3e\x95\xd5\xc8\xe5\xd4\x2a\xab\x41\xa9\x06\x6b\x07\x53\xcf\xfa\x5f\x1d\xf2\x06\x55\x2f\xed\x03\xf0\xf2\x86\x68\x07\xcc\x1f\xf6\x23\x16\x2d\xdb\x0f\x79\x5a\x6d\x97\x22\x73\x33\x5b\xb8\xde\x12\x6e\x66\xcb\x79\xd7\x6b\x94\x23\x45\xd2\xe9\xfd\xbc\xe3\x34\x3e\xc9\xd8\xc6\xc2\xbd\x75\x2f\x97\xf0\xca\x01\xf3\x8e\xe1\x14\xbc\x24\xc2\xa6\x03\x6b\x14\x31\xec\x80\x60\x1f\xe3\xb7\x1f\x5c\xcf\x85\xd9\x7c\x09\xee\xef\x37\x8b\xe5\xc2\x2a\x6c\xe0\xca\x9b\xff\x54\xfb\x9e\xa2\x96\x4a\xb1\x7e\xab\x3b\x7c\x8e\x14\xce\xf5\x14\x1c\xb8\x8e\x92\x15\x8a\x4c\x95\xc3\xd1\x32\x29\xfd\x8a\x8c\x2e\x3d\x77\xba\x74\x61\xee\x81\xe7\xfe\x7c\x3b\xbd\x74\xe1\xea\x6e\x76\xb9\xbc\x99\xcf\x60\x83\xd5\xa1\xd2\xd7\xd3\xb3\x6c\xf0\xdc\xe5\x9d\x37\x5b\x14\xe2\x3f\x5d\x18\x67\x67\x45\x02\xfa\x52\xa3\x19\x90\xd1\x55\x5a\xe1\x40\xb4\xd0\xed\x74\x76\x7d\x37\xbd\x76\x61\xf1\xcb\xed\x98\xd8\x36\x32\x9f\x53\x86\x58\xd6\xac\x23\xd4\xe2\x04\x95\xa4\x90\xe2\x6d\x84\x02\x0c\xeb\x8c\xaa\x59\xdf\x60\x8a\x53\xc4\xb1\xff\xc0\x32\x11\x9d\x12\x42\xcc\xb3\x94\x32\x08\xee\x51\x6a\x9d\xbf\x56\x37\x23\x44\x37\x19\xda\x60\x60\x1f\x22\x03\x31\x81\xc5\x70\x24\x86\x97\x65\x2b\xc6\x53\x42\x37\x96\xd1\xe0\xc8\x1c\xce\xe2\x89\x2f\x7d\x05\x51\x12\x3c\xf8\xe5\xf9\xc0\xb2\x1d\x30\x77\xbb\xdd\x2e\x8e\xc3\xf0\xfe\xfe\xfc\xff\x82\x58\x32\x66\xda\x4d\x3f\xfa\xf5\xf9\x33\x94\xfe\x52\x44\xc3\x24\xb6\x6c\xf8\x1f\xbc\xc2\xdf\x39\x60\xbe\x2c\xae\x7e\x1f\x02\x17\x4c\x07\x4c\xd3\x76\xc4\x68\x9e\xbf\xb6\x2f\x2e\x8a\x6c\xdf\x18\x67\x67\x2d\xe3\x54\x32\x71\x31\x44\xf5\xc1\x12\xaf\x72\xb6\xf1\xeb\xf4\xf6\xce\x5d\x80\xf5\xaa\x63\x56\xfa\x5b\xc3\x01\xf3\xd3\x97\xf6\x59\xae\x13\x65\x89\x63\x0a\x67\xed\x16\x35\x32\xaf\x0c\xd4\x54\x0b\x93\xa1\x16\xc9\x13\x55\x18\x43\x0d\x1e\x09\x7e\x1a\xb3\x5e\xb5\xe7\x18\x8b\x6d\xb6\x8a\x48\xe0\x17\x40\x35\xbb\xd6\x7f\xae\x54\x96\x29\x79\x6c\x01\x3b\x60\x24\xe1\xda\xcb\x5c\xc9\xa1\xb5\xff\x0e\xe2\x40\xd9\x33\x2a\x13\xa7\x38\x44\x39\xda\x89\xa7\xd6\x36\x96\xce\x04\x24\xd4\x88\x73\x00\x67\x3a\x50\xec\xaf\x03\xaa\x35\x44\x4b\x19\xa0\x44\xcf\x27\xe1\xe4\x02\x3e\xc9\xc1\x98\x24\xdb\xc9\x05\x4c\xf0\x87\x89\x12\xdd\x89\x3c\x20\x89\x5b\x0c\x17\xd1\x4f\x0c\x80\x2f\x86\x6a\xc8\x1e\x0a\xc1\x8f\x28\xca\x04\x85\xe4\xff\x8e\x17\x50\x56\x99\x9d\x3a\xc6\x39\x50\xb9\x2d\xbe\xb2\x96\x5f\x57\xd4\xef\x3a\xf5\xe4\xe7\xfc\x03\xcc\xa3\xc1\xbc\x78\x07\x26\x09\x4d\x78\xf1\xee\x1d\x98\x22\x22\x6c\xda\x17\x17\xf9\xf9\xf4\x6d\x81\xa5\x93\xc2\xdc\x03\xab\xcb\x81\xac\x86\x69\x8b\x32\x57\x81\x9b\xa5\xf5\x74\xf6\x7d\x19\xfd\x5b\x2d\x35\x5b\x50\xc7\x3f\x01\x00\x00\xff\xff\x24\x32\xe5\xde\xe3\x20\x00\x00")

func _0000_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0000_initUpSql,
		"0000_init.up.sql",
	)
}

func _0000_initUpSql() (*asset, error) {
	bytes, err := _0000_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0000_init.up.sql", size: 8419, mode: os.FileMode(0644), modTime: time.Unix(1642285202, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8d, 0x83, 0xbd, 0x90, 0xa, 0xa, 0xed, 0x24, 0xfd, 0x34, 0x77, 0x1d, 0x69, 0x9c, 0x85, 0xbd, 0x61, 0xb7, 0xf9, 0xd1, 0x97, 0x7c, 0x5c, 0xf0, 0x15, 0x3f, 0x2a, 0xb7, 0xfd, 0xa8, 0xb0, 0xa1}}
	return a, nil
}

var __0001_worldDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _0001_worldDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_worldDownSql,
		"0001_world.down.sql",
	)
}

func _0001_worldDownSql() (*asset, error) {
	bytes, err := _0001_worldDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_world.down.sql", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1642271179, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __0001_worldUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x57\xcd\x6a\xeb\x38\x14\xde\xfb\x29\x0e\xd9\xd8\x81\x30\xb4\x17\xba\x9a\xd5\xe5\xe2\x76\xc2\x74\x52\x48\x32\x03\x43\x29\x46\xb1\x15\x57\x54\x96\x82\x2c\xb7\x75\xdf\x6d\x36\xf3\x64\x83\xfc\x17\xc9\x96\x7f\x12\x26\xd9\xdc\x6c\x02\xd2\xa7\xf3\x7f\xbe\x73\x1c\x0a\x8c\x24\x06\x89\x76\x14\xc3\x2c\x62\xd1\x1d\x0e\x3e\xb8\xa0\xd1\xcc\xf1\x1c\x00\x00\xed\x28\x20\x91\x3a\x81\x1d\x89\x09\x93\xc5\xad\xfa\x85\x9c\xa5\x52\x20\xc2\xa4\x01\x3e\xbc\x35\x08\xf5\x3b\x08\x92\x20\x91\xc3\x1b\xce\x17\xc5\x45\xa9\x39\x0a\x76\x79\x83\x29\x05\x03\xb0\x8c\x52\x13\x84\x64\x03\x92\x24\xc1\xa9\x44\xc9\x01\x3e\x88\x7c\xe5\x99\x2c\x4e\xe0\x8b\x33\x0c\x11\xde\xa3\x8c\x4a\xf0\x18\xff\xf0\xe6\x80\xf4\x3b\x37\x93\xa1\x3b\x07\xc6\xa5\x26\xff\x1d\x8b\x94\x70\xa6\x1b\xda\x18\x51\x01\xf5\xbb\x46\xfe\xcd\xbc\x7c\x4f\xd2\x00\x85\x92\xbc\x63\xed\x3d\xe7\x14\x23\x36\xfc\x5e\x8a\x0c\x57\x22\x42\x9e\x24\x9c\x05\x28\x0c\x71\x9a\x56\x2e\xe2\xcf\xc2\x5f\xd3\xd6\x2c\xc5\x22\x90\x28\x4e\x8f\xf2\x14\xf0\xf9\xa5\x0d\x4c\xf3\x54\xe2\xc4\x80\xda\x81\x11\x16\xe4\x1d\x47\xc1\x5e\xf0\xa4\xcc\x9a\x25\x05\x0c\x25\x18\xcc\x5f\xd7\xbe\x02\xf9\xe3\x69\xb5\xd9\xae\xbf\x2f\x57\x5b\xd8\xbf\x05\x7a\x2d\x54\x69\xdc\xe5\x4d\x49\xdc\x3f\xad\xfd\xe5\xc3\x0a\x7e\xf7\xff\x06\xef\x58\x0a\x73\xa3\x66\xd6\xfe\xbd\xbf\xf6\x57\x3f\xfc\x0d\xcc\x94\xf7\x33\xf0\x8a\x20\x90\xa8\x8a\xdd\x80\xc6\x22\xaa\x65\x50\x7b\x94\xea\x71\xef\xd5\x6b\x66\xc7\x53\xc1\x18\xd5\xad\x87\xd5\xae\xbb\x1b\xf8\x7e\xc7\xf5\x9e\x04\xcf\x6c\xc7\xb9\xe3\xcc\x7f\x75\x1c\xc7\xd6\xc5\x89\x6a\x4b\x55\x31\xf9\x01\xd7\xcd\xac\xf5\x5c\x95\x68\x6b\xab\x5d\xaa\xc5\x6a\x9d\x96\xce\xe8\xb6\x96\x56\x77\x55\xbd\xd5\xef\xfa\x99\x47\xf7\x79\x90\x80\x7a\x33\x68\x48\xf8\xff\xcb\xb6\x48\x97\x2d\x5b\x29\xf9\xc2\x41\x88\x24\x8e\xb9\xc8\x7f\x96\x74\x19\x4e\x8f\x0e\x8c\xf4\x80\x42\x6c\xd5\x30\x90\x4f\x53\xc5\x15\x13\x5a\x15\x92\x4a\xa5\x36\x47\xeb\xf2\xaa\x26\xe9\xf8\x2c\xad\x1f\x9c\x31\x4d\xf5\x14\x0e\x0d\x54\xb8\xc2\x50\x3d\xb9\x9c\xf6\x44\xa4\xd2\x5c\x3b\xac\xfe\x58\xa6\x93\xb5\x40\x8a\x43\x63\x76\x1a\xd8\xe7\x97\x16\x56\xe7\x81\x7e\x89\x88\x92\x98\x25\x98\xc9\x09\xda\x8d\x4a\xec\x45\x25\x84\x52\x12\x84\xaf\x88\x52\xcc\x62\x1c\x08\x24\x09\x8b\xdb\xb1\x2b\xb1\x14\xb1\x38\x43\x31\x6e\x4f\xf9\xf6\x90\x4f\x43\x41\x0e\xb2\x95\x91\x6e\x0b\x8d\x72\xe2\x35\xa7\x78\xad\xb3\xa8\x00\xbb\x3e\xb3\x44\xce\x1c\xa0\xe3\x16\x58\x78\xc4\x6e\x8f\x01\x19\x33\xc7\xe4\xfb\xd1\xc5\x42\x2f\x47\xbb\x76\x1d\x31\xa6\xdc\x58\x0d\x2a\xdd\xbd\x44\x46\xd8\x2b\xda\x11\x89\x98\x6c\x71\xd9\xf1\x42\x75\xe8\x18\x91\x69\xe8\x73\xb8\x6c\xd2\xa7\xc1\xe5\x69\xec\xa4\x8f\x03\xcb\xc2\x3e\x7d\x63\xef\x5b\xd9\xbb\x9f\x63\x1d\xa3\x16\xbd\x03\xc7\x8e\xe4\x42\x1d\x23\x5a\x7d\x02\x94\x32\xa7\x7c\xc5\xec\x11\x4d\x55\xd5\xf6\x96\xad\x96\xf4\x6b\x92\x87\xa6\x76\x80\x3f\xda\xab\xf4\x70\xcf\x9c\xcc\x1f\x9a\x11\x55\x12\x86\xcc\x38\xe6\x69\x62\xf3\xce\x6c\x2f\xb5\x1e\x26\x2c\xc2\x9f\xdd\xde\x2b\x0d\xaf\xde\x00\x67\x1d\x44\xdb\xc3\x45\xa7\x90\xfa\x88\xa2\x1e\x47\x2d\x9a\xa8\x8f\xa7\x90\x44\x83\x3d\x9d\x22\x26\x6c\x3a\x17\x66\x87\x73\xd7\x65\xad\xdf\xa7\xed\xb4\x4d\x94\x2e\xb4\xce\x2e\x57\x1b\x7f\xbd\x85\xe5\x6a\xfb\x34\x34\xae\x16\xe5\x46\x3e\x77\xfe\xfa\xfe\xf8\xa7\xbf\x01\xcf\xdd\x12\x96\xbb\x0b\x70\xbf\xfd\xfb\x0f\xec\x72\x50\x7f\x7b\xf9\x8b\x5b\xb9\x0c\x00\x9e\xbb\x49\x10\xa5\x0a\x73\xa7\x10\x77\x9d\xfb\x3f\x70\x44\xb2\x64\x00\xf0\x88\x44\x8c\xd5\xfd\xed\x8d\x02\xdc\xde\x74\x10\xbf\x65\x15\xa0\x90\x70\xdb\x15\xf1\x80\x44\x8c\x98\xcc\x10\x2b\x8c\x2d\xe4\x7c\x2b\xe4\x00\x17\x40\x95\x02\xe1\xaa\x40\xfc\x17\x00\x00\xff\xff\xdf\x39\x25\x66\x1d\x13\x00\x00")

func _0001_worldUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0001_worldUpSql,
		"0001_world.up.sql",
	)
}

func _0001_worldUpSql() (*asset, error) {
	bytes, err := _0001_worldUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0001_world.up.sql", size: 4893, mode: os.FileMode(0644), modTime: time.Unix(1642271179, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8a, 0x9e, 0x2c, 0xce, 0x54, 0x84, 0xba, 0x68, 0x1f, 0x23, 0x82, 0x8c, 0x1f, 0xc, 0x61, 0xf, 0xad, 0xb6, 0x5e, 0x97, 0x7a, 0x6f, 0x86, 0xfb, 0xf0, 0xb3, 0x79, 0xe7, 0xc2, 0xec, 0xdf, 0x48}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"0000_init.down.sql":  _0000_initDownSql,
	"0000_init.up.sql":    _0000_initUpSql,
	"0001_world.down.sql": _0001_worldDownSql,
	"0001_world.up.sql":   _0001_worldUpSql,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"0000_init.down.sql": {_0000_initDownSql, map[string]*bintree{}},
	"0000_init.up.sql": {_0000_initUpSql, map[string]*bintree{}},
	"0001_world.down.sql": {_0001_worldDownSql, map[string]*bintree{}},
	"0001_world.up.sql": {_0001_worldUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
