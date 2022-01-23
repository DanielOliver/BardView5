// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../migrations/0000_init.down.sql (0)
// ../migrations/0000_init.up.sql (3.08kB)
// ../migrations/0001_world.down.sql (0)
// ../migrations/0001_world.up.sql (5.378kB)

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

	info := bindataFileInfo{name: "0000_init.down.sql", size: 0, mode: os.FileMode(0666), modTime: time.Unix(1634514631, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __0000_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\x4b\x8b\xdb\x3c\x14\xdd\x07\xf2\x1f\x2e\xb3\xb1\x03\x59\x7c\x9b\x6f\xd5\xd5\x30\x78\x4a\x68\xc9\x40\x92\x16\x4a\x29\x46\xb1\x6f\xa6\x6a\x6c\x29\xd5\x23\x89\xfb\xeb\x8b\x2d\xbf\x24\xbf\x12\x98\xa1\x77\x67\xf9\x1e\xe9\x9c\xfb\xd0\x55\x24\x90\x28\x04\x45\xf6\x09\xc2\x43\xc4\xd3\x94\xb3\x90\x44\x11\x4a\xf9\x30\x9f\xf9\xf3\x19\x00\x00\x23\x29\x82\x31\x85\x57\x65\x16\x73\x8b\x38\x93\x4a\x10\xca\x14\x58\xd0\xf0\x74\x84\x93\xa0\x29\x11\x19\x1c\x31\x5b\x1a\x84\x39\x2b\x0e\x89\x02\x45\x53\x94\x8a\xa4\x27\xb8\x50\xf5\x93\x6b\xb3\x02\x7f\x38\x43\x88\xf1\x40\x74\xa2\xc0\x67\xfc\xe2\x2f\x80\xb4\xff\x79\x5a\x45\xde\x02\x18\x57\xc0\x74\x92\xcc\x67\x8b\x0f\xf3\xd9\x7c\x66\xab\xd0\x12\x45\x43\x3e\xff\x0a\x69\x5c\x51\xde\xd3\x57\xca\xfa\x25\x14\x9e\xa7\x63\xf3\x2f\xb7\xae\x0c\xad\x9b\xdd\xba\x9f\x96\x55\x3c\x9d\x00\xec\xb3\x16\x95\x01\x64\x17\x45\x8c\xeb\x58\xe8\x6a\xbb\x23\x86\xe5\x29\x67\x14\x92\x72\x66\x87\x69\x5c\x55\x73\xca\x7f\x8b\x72\x17\x3c\x1c\x30\x52\xf4\x8c\x61\x5c\x24\xe4\x7d\xb8\x22\x8b\xcd\xfe\x30\x19\x91\x16\x8a\xca\x90\x14\xdc\x4a\x85\x9c\x27\x48\xd8\x8d\x0a\x95\xd0\x58\x89\xb4\x4a\xdd\xb4\xc4\x8d\x05\x80\x29\xa1\x49\xf3\xf7\x0e\x64\xab\x05\xef\x44\x16\x55\xad\xc8\xab\x6c\x90\xdf\x7f\xdc\x84\x94\x99\x54\x98\xd6\xd8\x1b\x91\x06\xfb\xf4\xb2\xde\xee\x36\x8f\xab\xf5\x0e\x0e\xc7\xb0\xe0\x50\xd6\xf1\x3e\x6b\x1a\xec\xf9\x65\x13\xac\x3e\xae\xe1\x53\xf0\x0d\xfc\xa6\x3b\x16\x76\x0b\x6e\x82\xe7\x60\x13\xac\x9f\x82\x6d\xd9\xdb\xe0\x97\x4d\x5d\x65\xa4\xef\xb4\x22\x49\x26\x47\x43\x07\xb6\xf3\x38\x7c\xa6\x9d\x6e\x3f\xcf\xc4\xc2\xb9\x78\x34\xa3\xbf\x35\x02\x65\x31\x5e\x4d\xc0\x8b\x4c\x87\xda\xac\x70\x56\x13\x2f\xd6\x27\xb0\xf9\x8d\xd2\x03\xcd\x97\xfb\xae\x3b\xc1\x13\x0c\xa5\xde\xff\xc2\x48\xdd\x77\x67\xb7\x91\xff\xf6\xca\x2e\x98\xe4\xdd\xc9\x59\xaf\x84\x29\x11\x06\xda\xaf\xa1\xad\x72\xb4\x71\x6e\xb0\x81\x0b\xbd\xb8\x9a\xdf\x28\x40\x03\x2d\xd4\x96\x59\x8a\x19\x28\xeb\xb6\xde\x91\x4e\xb2\xca\xc6\x29\xeb\xd5\x7a\x1b\x6c\x76\xb0\x5a\xef\x5e\xec\xf0\x55\x6e\x5f\x1f\x3f\x7f\x09\xb6\xe0\x7b\x31\x8b\xff\xc7\x0b\x17\x49\xec\x0d\x40\x0d\x67\x83\x5c\x82\x43\xae\xde\x27\x25\x8c\xbc\xa2\xb7\x04\x6b\xc7\x65\x4d\xdf\xf7\xf8\x85\xa1\x18\x73\x38\x53\xbc\xb8\xff\x3b\x94\xfa\x9b\xb9\xa6\x71\x12\xf4\x4c\x14\xda\x1b\x13\x96\xe5\x2d\x68\x2f\x9e\xf4\x3e\xa1\x91\xd7\x57\xcc\x78\x55\x28\x18\x49\x42\xc9\xb5\x88\xb0\x29\x68\xe7\x47\xfd\x2a\x19\x79\x93\xb8\x90\xe9\xe7\x89\xf3\xc8\xa8\xad\x9c\xe8\xc3\x0f\x8b\xda\xde\xae\x90\xa1\xfb\xb0\xe8\xf0\xa9\x46\x6d\xdb\x7a\x1e\x16\x4e\x1c\x8e\x68\xe4\x95\xfd\xec\xce\x59\xc7\xbb\xe2\xd0\xef\xed\x4c\xc8\x26\x0e\x66\xe0\x8d\x4f\xc5\x09\x6f\x67\x6e\x5b\xde\xd3\x73\xd3\xd5\xf1\x1e\x23\xd4\xf4\xfc\xdf\x00\x00\x00\xff\xff\x5b\x8a\xa7\x34\x08\x0c\x00\x00")

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

	info := bindataFileInfo{name: "0000_init.up.sql", size: 3080, mode: os.FileMode(0666), modTime: time.Unix(1642385905, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8a, 0xc9, 0x93, 0x3e, 0xa8, 0xf3, 0x87, 0x2a, 0x70, 0xa9, 0x84, 0x9f, 0x44, 0xe9, 0x39, 0xcf, 0xce, 0xc7, 0x64, 0xc8, 0xef, 0x37, 0xfb, 0x6, 0xc9, 0x17, 0xac, 0xb9, 0x31, 0x19, 0x3f, 0xd6}}
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

	info := bindataFileInfo{name: "0001_world.down.sql", size: 0, mode: os.FileMode(0666), modTime: time.Unix(1640657667, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __0001_worldUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x57\x4d\x6f\xe3\x36\x10\xbd\x07\xc8\x7f\x18\xf8\x62\x19\x30\x8a\x64\x81\x9c\x7a\x0a\x16\xda\xad\xd1\x5d\xa7\xb0\xdd\x02\x8b\xa2\x10\x68\x69\xa2\x25\x42\x91\x06\x49\x25\xd1\xfe\xb7\x5e\xfa\xcb\x0a\x7d\x45\xa4\x44\x4a\x4a\x50\x6f\x0f\xab\x43\x02\x88\xcf\x33\x8f\x9a\x37\x8f\xc3\x58\x22\xd1\x08\x9a\x1c\x19\xc2\x22\xe1\xc9\x0d\x46\x4f\x42\xb2\x64\x71\x79\x11\x5c\x5e\x00\x00\x18\x2f\x23\x9a\x40\xf5\x1c\x69\x4a\xb9\xae\xd7\xcb\x27\x16\x5c\x69\x49\x28\xd7\x16\xfc\xf4\xd0\x41\xca\xe7\x24\x69\x46\x64\x01\x0f\x58\xac\xeb\x95\x3a\x7f\x12\x1d\x8b\x0e\x55\x07\x07\xe0\x39\x63\x3d\x18\xd1\x1d\x4c\xd3\x0c\x95\x26\xd9\x09\x9e\xa8\xfe\x2a\x72\x5d\xbd\x81\x6f\x82\x23\x24\x78\x4f\x72\xa6\x21\xe0\xe2\x29\x58\x01\x31\xd7\x96\xb9\x8e\x97\x2b\xe0\x42\x9b\x19\x1e\x51\x2a\x2a\xb8\xc9\xb6\x23\xd2\x40\xcd\xb5\x97\x0c\x57\xab\x26\x02\x55\x11\x89\x35\x7d\x44\x33\x82\x10\x0c\x09\x1f\x8f\xa0\x65\x8e\x6d\x90\x58\x64\x99\xe0\x11\x89\x63\x54\xaa\xdd\x28\x3e\x57\xdb\xee\x31\xce\x15\xca\x48\x93\x54\x19\x31\x4b\xe8\x9f\x7f\x0d\xa0\xaa\x50\x1a\x33\x1b\xec\x81\x26\x28\xe9\x23\x26\xd1\xbd\x14\x59\x5d\x45\x67\x41\x38\xc9\x10\x7a\x8f\x87\x67\x26\x92\x9c\xa1\x07\x6a\x26\x56\xb1\xa4\x27\x6d\x16\xc1\x13\x11\x9f\x35\x4a\x4e\x58\xa4\x44\x2e\x63\xac\x54\xe9\xe0\xd8\x87\x3d\x60\xd1\x4f\x5c\x03\xdf\xdf\x6d\xf7\x87\xdd\xed\x66\x7b\x80\xfb\x87\xc8\x14\x70\xa3\xbb\x63\xd1\xe9\xf8\xc3\xdd\x2e\xdc\x7c\xdc\xc2\xaf\xe1\x17\x08\x3a\xf9\xae\x6c\xa5\xef\xc2\x0f\xe1\x2e\xdc\xbe\x0f\xf7\xb0\x28\x0b\xb5\x80\xa0\xaa\x17\x4d\xda\x4a\x8f\x24\xad\x34\x50\x4b\xc0\x97\xd7\x94\x89\x3f\xb5\xad\xa6\xa0\x2c\xda\x74\x7a\x53\x01\x9e\xf4\x43\x91\x8c\x6c\xdf\x34\x15\x08\x6c\x37\x59\xcd\xa8\x81\x95\xac\x57\x53\x0f\xbf\xa1\x40\x46\xf8\xf5\xc0\x0b\xcf\xcf\x2f\x2f\x56\x3f\x97\x7f\xbd\x8e\x19\x11\xa5\x68\xca\x33\xe4\xba\x33\xcf\x9e\xbf\x35\x2a\xf5\x3b\xdb\x59\x4d\xad\xcd\xee\x30\x23\x87\x9f\x35\x7a\xf5\xfc\x78\xed\x3c\x1c\xdc\x20\x29\x18\x56\xde\x58\x53\x69\x7a\xd0\x00\x0d\x24\xe0\xfe\xac\xd6\x79\xf2\xdb\x6e\xf3\xf9\x76\xf7\xa5\x2e\x78\xc3\x75\x0d\xaf\x97\x97\x11\xbe\x8c\xe2\x51\x54\xdb\xbc\xff\x5d\x97\x1b\x79\xab\x17\xbe\x56\xb3\x37\xf4\xf6\x36\x9b\xcd\xc6\xa8\x96\x87\x93\x81\x18\x21\x64\xa0\x16\x8d\xf7\xb4\x3d\xe4\xee\xa2\xac\x1c\x22\xca\x33\xad\x38\xa1\xb3\x83\x46\xbb\xe7\x6c\x9d\xf3\xba\xae\x31\x0e\xc6\x9e\xce\x47\xa6\x25\x73\xe7\x13\x43\x93\xbf\x90\x56\x90\x73\x1c\x5c\x63\x0e\xa8\xe8\x37\x8c\x62\xa2\x31\x15\xb2\xf8\xb1\x8a\x67\x6d\x7d\xc6\xc8\xab\x4e\x24\x46\x67\x92\xd1\xf2\xda\x69\xbe\x77\x7d\x1b\x71\xd5\x95\x35\xad\xbf\x55\xdd\xcb\x41\x31\x79\x35\x68\x7f\xf1\xb6\xcb\xc1\x9c\x43\xb4\x7d\xce\x7f\x43\x78\xb5\xc2\x5c\x97\x29\xe7\x96\x5c\x03\xb6\x53\x31\xe0\xb9\x0a\x74\x23\xfe\x8c\xcb\x80\x17\x6c\x9a\xca\x24\x0d\xc2\x9a\x33\x64\x0e\x67\x4b\xcf\x53\xe0\x8c\x32\x46\xa3\xf8\x2b\x61\x0c\x79\x8a\x91\x24\x9a\xf2\xd4\x33\x6e\x30\xc2\xd3\x9c\xa4\x38\xeb\x6b\xb8\xae\x1d\x9e\xce\x9c\xb6\xde\xef\x7c\x5d\x68\xd3\xfe\xaf\xf3\x43\x4b\xc2\xe1\x52\x1e\x4a\x16\x66\x92\x91\x7d\xb2\x4c\x5f\x62\x4c\xc5\x7a\x08\x98\x90\xc9\xfc\xd6\x54\xd2\x9b\x63\x1a\xa7\xa4\x3c\xc1\x67\x70\xd4\x04\x04\xb7\x5f\xf7\x3f\xec\x7a\xe0\xa2\xf3\x02\x47\x95\x3d\x4c\x47\xaf\xd8\x02\xe5\x31\xcb\x13\x6c\x97\x3d\xc9\x2c\xbf\x6f\x7b\x68\x68\xf8\xed\x4a\x37\xee\x8f\x78\xfd\x0b\xf8\x2d\x66\x3f\xc7\xe7\xcf\x6d\xf1\x6f\x9e\x1f\x46\x7c\xc4\xa7\xdc\x97\x8f\x75\xce\xc3\x7d\xb3\xdd\x87\xbb\x03\x6c\xb6\x87\xbb\xb1\x0e\x5b\xd7\x63\xca\xea\xf2\xe2\x8f\xdb\x4f\xbf\x87\x7b\x08\x96\x07\xca\x8b\xe5\x1a\x96\xef\xfe\xf9\x1b\x8e\x05\x94\xff\xee\xf5\x4f\xcb\x76\xef\x00\x10\x2c\xf7\x19\x61\xac\x04\xdd\x94\x90\x9b\x21\xe0\x33\x26\x34\xcf\xc6\x10\x9f\x88\x4c\xb1\x04\x5c\x5f\x95\x88\xeb\xab\x21\xe4\x97\xbc\x41\x54\x31\xae\x1d\x41\x3e\x12\x99\x12\xae\x73\xc2\x2b\xc6\x55\xa4\x77\x55\x24\x10\x12\x58\x99\x42\x2e\xeb\x0f\xf2\x6f\x00\x00\x00\xff\xff\xd6\x4d\x05\xff\x02\x15\x00\x00")

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

	info := bindataFileInfo{name: "0001_world.up.sql", size: 5378, mode: os.FileMode(0666), modTime: time.Unix(1642977567, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc5, 0xca, 0x88, 0xb1, 0xa8, 0xfa, 0x4c, 0x47, 0x7c, 0xf8, 0xdf, 0xda, 0xf4, 0xe7, 0x28, 0x4c, 0xed, 0x80, 0x3, 0x65, 0xf0, 0x9, 0xfe, 0xbe, 0x5e, 0x6a, 0xf3, 0x31, 0xd3, 0x16, 0x69, 0xe0}}
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
