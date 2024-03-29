// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../migrations/0000001_init.down.sql (0)
// ../migrations/0000001_init.up.sql (9.92kB)

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

var __0000001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _0000001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0000001_initDownSql,
		"0000001_init.down.sql",
	)
}

func _0000001_initDownSql() (*asset, error) {
	bytes, err := _0000001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0000001_init.down.sql", size: 0, mode: os.FileMode(0666), modTime: time.Unix(1634514631, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __0000001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x59\x4f\x6f\xdb\x36\x14\xbf\x07\xc8\x77\x78\xc8\xc5\x36\xa0\x04\x49\x87\x9e\x76\x0a\x32\xb7\x0b\xda\x26\x83\x93\x15\x28\x86\x41\xa0\xa5\x67\x87\x8b\x44\x6a\x22\x95\xd4\x1d\xf6\xd1\x76\xd9\x27\x1b\x48\xfd\x23\x25\x52\x96\x03\x3b\x3d\x4c\x97\xd6\xd2\xfb\xc7\xf7\xf7\xc7\x97\x28\x47\x22\x11\x24\x59\x26\x08\x27\x11\x4f\x53\xce\x42\x12\x45\x28\xc4\xc9\xf1\xd1\xf4\xf8\x08\x00\x80\x91\x14\xa1\x7c\x24\x7e\x95\xe5\x4b\xf5\x44\x9c\x09\x99\x13\xca\x24\x58\xac\x61\xf6\x08\x59\x4e\x53\x92\x6f\xe0\x11\x37\x41\xc9\x51\xea\x8a\x43\x22\x41\xd2\x14\x85\x24\x69\x06\xcf\x54\x3e\xf0\xa2\x7c\x03\xdf\x38\x43\x88\x71\x45\x8a\x44\xc2\x94\xf1\xe7\xe9\x0c\x88\xf9\x6d\x52\xc8\x68\x32\x03\xc6\x25\xb0\x22\x49\x8e\x8f\x66\x3f\x1e\x1f\x1d\x1f\x5d\xdf\xdc\xcd\x17\xf7\x70\x7d\x73\x7f\x6b\x1b\x02\x53\x65\xfb\xec\xf8\xe8\xf3\xe5\xc7\x5f\xe7\x77\x30\x9d\x64\x39\x7d\x22\x12\x27\xb3\xa0\x39\xc6\x74\x42\xd8\xa6\x10\x98\xdb\x2f\xb3\x62\x99\xd0\x68\x52\x6a\xb0\xfd\xa4\x88\x5b\xf7\xa8\x5f\x21\x8d\x6b\xa7\x2c\xe9\x9a\x32\xb7\x93\x34\x65\xf6\xd8\x7e\x53\x4f\xdf\x51\x45\xd1\x4a\xeb\xff\xb4\x9e\xda\x13\x1d\x17\x2f\x37\x86\x29\x1e\xce\x3e\x17\x29\x49\x87\x82\xd3\x3c\x3b\x44\xa9\xd2\xf2\x84\xb9\xa0\x9c\xd9\x6e\x1a\x3e\x55\xab\xe5\xbc\x0e\x0d\xae\x56\x18\x49\xfa\x84\x61\xac\x03\x72\x18\x5b\x91\xc5\xa5\x7c\xd8\xea\x11\x83\x8b\x8a\x90\x68\xdb\xaa\x13\x72\x9e\x20\x61\x23\x4f\x28\xf3\x02\xeb\x43\xda\x39\xac\x8b\x6e\x64\x02\x60\x4a\x68\xd2\x7e\xdd\x81\xd3\x28\xf2\x1d\x39\x75\x56\x4b\xb2\x16\x2d\xe7\x6f\xbf\x8f\xe2\x14\x1b\x21\x31\x6d\x78\x47\x72\x96\xbc\x57\xb7\x37\x77\xf7\x8b\xcb\xeb\x9b\x7b\x58\x3d\x86\xda\x86\x2a\x8f\x97\x9b\xb6\xc0\xde\xdd\x2e\xe6\xd7\xef\x6f\xe0\xc3\xfc\x0b\x4c\xdb\xea\x98\xd9\x25\xb8\x98\xbf\x9b\x2f\xe6\x37\x57\xf3\xbb\xaa\xb6\x61\x5a\x15\x75\x1d\x11\x97\x36\x1d\xa4\x32\x46\x3e\x85\x66\x1c\xfd\x3a\xdd\x2d\xcb\x6a\x3c\x05\xa3\x7f\x16\x08\x94\xc5\xf8\xb5\x74\xb8\x8e\x74\x58\x94\x6f\x38\x6b\x0c\xd7\xef\xb7\xf0\xaa\x8e\xe2\x60\x55\xaf\x5d\xed\x2e\xe7\x09\x86\xa2\x58\xfe\x81\x91\xdc\x6d\x2a\x98\x9c\xdf\x73\x28\x98\x76\xf4\x67\x42\xcc\xe2\xb7\x28\x50\x4a\xca\xd6\xce\x7e\xaf\xd9\xe5\x26\xc3\x17\x9c\x5e\xb1\x7d\xf7\xa3\x2b\x23\xfa\xe7\xa6\x8c\xf5\x46\x61\x54\x08\xc9\x53\xbf\x17\x54\x83\xe3\xcc\xe9\x87\xe1\xa6\x31\xe2\xe9\x74\x07\x2b\x68\xfb\x15\x6d\x4d\xbc\x3d\x39\xbf\xe9\x4c\xdd\x0c\x28\x5d\xd6\xc9\x81\x32\x1a\x81\x75\xc8\x99\xa7\xb9\x99\x52\x2a\x52\x4f\xc3\xb1\xa4\xf9\x7b\x9c\x55\xd0\x9d\x86\xd3\xcb\x9d\x52\xb1\xd3\x5e\x23\x97\x52\xc2\xc8\x1a\x27\x01\x74\xaa\xc9\xcc\x2d\xfe\xcc\x30\x1f\x26\x79\xa2\xf8\xdc\xa7\xf0\x24\x63\x9b\x85\xda\x2a\x13\x25\x0d\x80\x30\x4d\xbb\x1d\x84\x75\xe6\xe1\xbe\x93\xbb\x83\xd4\x86\xf1\xd0\x18\xf1\x5e\x38\xb7\xcf\x04\x87\xda\xd7\xba\xa1\x1c\xc4\x31\x56\xd5\xef\x5f\xfc\xe9\x29\x5c\xaf\x20\xe3\x59\x91\x28\x27\x05\xf0\xee\x03\x48\x0e\x04\x44\x86\x11\x5d\xd1\x08\x72\x8c\x78\x1e\x9f\x55\x08\x25\xe2\x99\x95\x59\xfb\x0b\xd3\xe9\x69\xa5\xdb\xa8\xb2\x2a\x80\x24\x23\x4b\x9a\x50\x49\x51\x34\x2e\xf0\x01\xa3\xdd\x5d\x40\x84\xa0\x6b\x16\x72\x56\x01\xa6\x41\xa0\xda\xe4\xc7\x8a\x24\x02\x67\x3b\x89\x27\xb1\xf6\xdb\xfe\xc4\x7b\x9b\xe3\xab\x21\xbf\x76\xa4\xaf\x1e\x87\x7a\xb0\x22\xf1\xab\xea\xce\x64\xbf\xaa\xef\xd5\xf5\x2b\xa1\x34\x0e\x5a\x63\xed\xe6\x1f\x34\xb5\x11\x58\x09\x1b\x40\x39\x28\xba\x59\x16\xd8\x89\x61\xcc\x8e\x8b\xf3\x00\x6a\x2c\xd2\x6d\xfd\x41\x19\x7a\x98\xfc\x05\x27\x6a\x36\x9c\x04\x70\x52\x8e\x1a\xf5\x3f\x3d\x50\x4e\xe0\x6f\xc5\x76\x5b\x0d\x17\x75\x8f\x0a\xa0\x4c\x27\x63\xb2\xbc\xd9\x8b\x92\xcb\x38\xa5\x6c\x52\x89\x77\x68\xf9\x61\x17\x2d\xa5\xc8\xcf\x14\x9f\xb5\xe1\x95\x4c\x7d\x0f\xf4\xa2\x2f\xed\xc2\x14\x99\x81\xc3\xad\x49\x52\xb5\xa7\xc3\x43\x1d\xe8\xde\xe8\x6b\xcd\xf5\xc5\xd6\x59\xe5\xcd\x35\xde\xda\x99\x74\x58\x03\xc7\x4c\x77\x93\x34\x2d\xb4\xd7\xbe\x03\x10\x45\xf4\x00\x44\xa8\x4f\x65\x00\xce\xe0\xa7\xd2\x0c\xa1\x38\xe4\x03\x15\x40\x57\x5a\x62\xc5\x8b\xb1\xba\x0c\x2d\x78\x82\x67\x8d\xf8\xab\xdb\xcb\x8f\xf3\xbb\xab\xb9\x2e\x87\xb3\x36\xe3\x3b\xd1\x68\xbe\xcc\x1c\x63\xa3\x6f\x7b\xaf\xdc\x3b\xe2\x2c\x78\xf2\xcb\xe2\xfa\xd3\xe5\xe2\x4b\x59\xea\x95\xdf\xcc\xe2\xab\xfc\x34\x0c\x1e\x5b\xd1\x4a\x82\xa7\x8f\xd4\x9d\x6f\x0f\x2d\xd2\x50\xa8\x7e\x0f\x35\xae\x41\x85\x1a\xe6\x99\x84\xae\xd2\xd0\x65\x16\x56\x61\x6e\x0b\xc3\x7a\xad\x82\x31\x00\x0a\x6d\xda\xed\xe8\xd0\x0b\xdf\x5e\x0f\x87\x75\xf6\x69\x86\x09\x8e\x12\xec\x57\x60\x67\x59\xd5\x8e\xe9\x41\xf6\x81\x2d\x55\x8b\xd4\x86\xf7\x43\x2d\x9c\x19\xde\x06\x79\xe9\xbc\xc8\xbc\x43\x97\xf2\xb8\x48\x5c\x3b\x2d\x83\x26\x46\x11\xe5\x34\x93\x8d\x23\x1d\xb2\x5c\x49\x6e\xe7\xcb\xab\xc1\x8f\x8e\xda\xd7\xdc\x40\x59\xb5\x96\xaa\xca\x51\x51\xb5\x76\x22\x46\x51\x0c\xce\xa1\x83\xd5\xc1\x6e\x33\xa8\xb3\xc2\x31\x82\x3e\xd0\x22\xcc\x93\x6f\xe9\x13\xfe\x00\x5a\x42\x0e\x91\x3c\x83\x7d\x92\x7e\xc3\x30\x22\x12\xd7\x3c\xdf\xfc\xbf\x82\x67\x1d\x7d\x44\x97\x17\x19\x89\xd0\xa9\x64\xb8\x2d\x58\x6a\x0e\x19\x5f\x13\xb6\x3b\xc3\x5b\xaf\x6d\xf4\x49\xcc\x7d\xcd\x3d\x65\x1b\x05\x3f\xdf\xfc\xfb\x0f\x2c\x37\xa0\xfe\x59\xc9\x33\x7b\x19\x73\x97\x92\x24\x51\x44\x6f\x15\xc9\xdb\x3e\xc1\x27\x8c\x69\x91\x0e\x51\x7c\x24\x79\xb9\x15\xba\x38\x57\x14\x17\xe7\x7d\x92\x9f\x8b\x8a\x42\xcb\xb8\x70\x08\x79\x4f\xf2\x35\x61\xb2\x20\x4c\x5b\xac\x25\xbd\xd1\x92\x80\xe7\x90\x28\x15\xb9\x73\x49\x64\x55\x5b\x17\x18\xd4\x45\xd8\x5c\xee\xb7\xc2\x83\x9a\xe3\x05\xf0\xc0\xd4\x30\x0c\x12\xaa\x74\x3b\x2c\x54\xe8\x98\x33\x0a\x30\xf4\xe0\x94\x5b\x88\x6f\x48\x5b\x65\xd4\x1b\xfd\xbc\xc8\x23\x14\x2e\x6a\x07\x00\xe8\x01\x8a\x41\xea\x84\xb0\x75\x41\xd6\x5d\xe9\x1e\x6a\x64\x4f\x34\xe7\x1a\xb9\x8a\xed\xd4\x54\x84\x09\xae\x91\xc5\x2a\xfc\x86\x63\xc7\x40\x29\xeb\xf2\x48\x45\x58\xfd\x99\xc6\x8e\xd0\xce\x82\xcc\xd1\xe2\xf2\xbb\xb1\xa7\x49\x2a\x84\xee\x89\x90\x11\x1d\xb3\xa5\x0c\x52\xa6\x34\x49\x68\x18\x3d\x90\x24\x41\xb6\xc6\x30\x27\x2a\x59\x5c\x79\x4f\xf2\x94\xe7\x61\x94\x10\x61\xc5\xa5\xde\xad\x19\x94\x0f\x54\x86\x19\xa7\x76\x3c\x9c\x94\x3d\x30\xe7\x3f\x91\xcc\x43\x11\xf1\xbc\xe3\x6e\xa7\xcc\xaf\x23\x29\x29\x93\x23\x29\x9f\xa9\x18\x49\x19\x71\x36\x96\xf2\x81\x8c\xa0\xdc\x0e\x49\x5e\x19\xca\xd6\x6a\xab\xb6\xe2\x51\xda\x6d\x3d\x03\xaa\xed\xbb\xa0\x83\x75\x84\x31\x8e\x41\xea\x31\xcc\xa2\xd9\x6e\x95\x63\x34\x8f\xb1\x47\xd5\xb2\xc7\x00\x93\x64\xab\x7e\x0b\xb5\xbb\x51\x7e\xf9\x37\x61\x67\x6c\x80\x33\xfb\x43\xdf\xb9\x41\x6f\xb4\x8e\x15\x1e\xea\x79\x31\x46\x83\xb6\x1a\x28\x8b\x92\x22\xc6\x9a\xc0\xa3\xd0\x82\x01\xf5\x1c\xe8\xe2\x80\xfa\xfd\xa8\x0d\x41\x43\xfc\x12\x0c\xb0\x65\x39\xa7\x9f\x43\x0f\xfe\x17\xa3\xec\xda\xbe\x1d\x70\x70\xe3\xac\xc3\x41\xe0\xff\x02\x00\x00\xff\xff\x0f\xdf\xec\x67\xc0\x26\x00\x00")

func _0000001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0000001_initUpSql,
		"0000001_init.up.sql",
	)
}

func _0000001_initUpSql() (*asset, error) {
	bytes, err := _0000001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0000001_init.up.sql", size: 9920, mode: os.FileMode(0666), modTime: time.Unix(1646544833, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xab, 0xde, 0x1a, 0xb8, 0x8a, 0x7b, 0xed, 0x6c, 0x6b, 0x1d, 0xb7, 0xa7, 0x88, 0x2e, 0x9d, 0x6c, 0x29, 0x73, 0x77, 0x34, 0x55, 0x13, 0x49, 0xa6, 0x81, 0x15, 0x6f, 0x6c, 0x88, 0xb8, 0xd2, 0x81}}
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
	"0000001_init.down.sql": _0000001_initDownSql,
	"0000001_init.up.sql":   _0000001_initUpSql,
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
	"0000001_init.down.sql": {_0000001_initDownSql, map[string]*bintree{}},
	"0000001_init.up.sql": {_0000001_initUpSql, map[string]*bintree{}},
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
