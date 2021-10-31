// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../migrations/0_init.down.sql (0)
// ../migrations/0_init.up.sql (8.422kB)

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

var __0_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _0_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0_initDownSql,
		"0_init.down.sql",
	)
}

func _0_initDownSql() (*asset, error) {
	bytes, err := _0_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0_init.down.sql", size: 0, mode: os.FileMode(0644), modTime: time.Unix(1635640494, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55}}
	return a, nil
}

var __0_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x6f\x6f\x9b\xc8\x13\x7e\xcf\xa7\x18\x59\x91\x80\x9f\xa8\xd4\x46\x3f\xa9\xba\x54\xad\x64\xe5\x48\x2e\xba\x9c\x7d\x87\x9d\xfb\xa3\xd3\x09\xad\x61\xed\x6c\x03\x8b\xcb\x2e\x49\x7d\x6d\xbf\xfb\x69\x77\xf9\xb3\x60\xc0\x70\x67\xa7\x6f\xca\xab\x18\x76\xe6\x99\x99\x9d\x79\x9e\x85\x04\x29\x46\x1c\x03\x47\xab\x08\xc3\x24\x48\xe2\x38\xa1\x3e\x0a\x02\xcc\xd8\xc4\xb0\x0c\x00\x00\x8a\x62\x0c\xea\xe2\xf8\x23\x37\xf2\xbf\x21\x48\x28\xe3\x29\x22\x94\x43\xcd\xce\xdf\x3e\xc0\x36\x25\x31\x4a\x77\xf0\x80\x77\x8e\x34\x50\x38\xa1\x8f\x38\x70\x12\x63\xc6\x51\xbc\x85\x27\xc2\xef\x93\x4c\xdd\x81\xbf\x13\x8a\x21\xc4\x6b\x94\x45\x1c\x2c\x9a\x3c\x59\x36\x20\xfd\x99\x99\xf1\xc0\xb4\x81\x26\x1c\x68\x16\x45\x86\xfd\xc6\x30\xea\xe1\x67\x0c\xa7\x45\xd4\xe2\x6f\x9f\x84\x45\xb0\x2b\xb2\x21\xb4\x35\x76\xb9\x70\xfb\x50\x3e\x12\xd7\x5e\xf8\x59\x56\xb9\xda\xff\x59\xbb\x8a\xf8\xea\x79\xaf\x76\x5a\x18\x1d\x86\x7b\x46\x48\xad\xec\x2b\x58\x79\x8d\xa8\x9c\x02\x79\xc4\x29\x23\x09\xad\x17\xa8\x3f\xa5\x0a\xe4\xa5\xad\x9c\xe0\xf5\x1a\x07\x9c\x3c\x62\x3f\x94\xfb\x70\x92\x48\x31\x0d\x95\x7b\x38\x58\x8e\xca\x88\x30\x1f\xc9\xc8\xf2\xf4\x92\x24\xc2\x88\x0e\x4c\x8f\xa7\x19\xce\x33\xac\xb5\xb6\x9a\x80\x61\x3b\x8f\x63\x44\xa2\xea\xe1\x70\x43\x6d\xde\xc6\x19\xca\x4e\xe6\x68\xc3\x2a\xc3\x3f\xff\x1a\x62\xc8\x76\x8c\xe3\xb8\x34\x1d\x66\x28\x2d\x2f\xe7\xb3\xc5\xd2\x9b\xde\xcc\x96\xb0\x7e\xf0\x65\x00\x79\xf3\xae\x76\xe5\x44\x5d\xcd\x3d\xf7\xe6\x7a\x06\x3f\xba\x7f\x80\x55\x0d\x84\x5d\x1b\x39\xcf\xbd\x72\x3d\x77\x76\xe9\x2e\xf2\x39\x06\x2b\x1f\xe1\x7c\x23\xda\xa0\xe4\xde\xa8\xad\xe9\x40\xd3\x77\xaf\x13\xb0\xbe\xc7\x96\xd8\x00\x5b\xa7\x97\x8c\x92\x0f\x19\x06\x42\x43\xfc\x51\x55\x59\x6e\xae\x9f\xa9\x3b\x09\x2d\x43\x96\xf7\xf7\x89\x29\x4d\x22\xec\xf3\xdd\x16\x17\xec\x54\xde\xf0\x5b\x88\xa4\x9b\xab\x2a\xb3\x43\x84\xd5\xa0\x90\xfa\x75\x24\x06\x6e\x6d\xd7\x06\x52\x4f\xf3\x0e\xb8\xea\x48\x71\x16\x71\xb2\x8d\xb0\x8f\x18\x23\x1b\x1a\x63\xca\x99\x8f\xa2\x28\x79\xc2\x61\xef\x80\x8f\x46\xca\x07\x22\x46\x14\x6d\x70\x73\x83\x8e\x84\xd4\xd9\x60\x62\x93\xe5\x1e\x8b\xca\xea\x3d\x56\x75\x51\xde\xa3\xad\x7d\x56\x6b\xb1\x21\x02\x28\x17\x0e\xed\xa7\x01\x3a\x36\xa4\x06\xff\x46\xea\xc6\xeb\xc6\x70\x71\x3a\xa5\x26\x8d\xaa\xc7\x08\xd9\x1a\xe2\xb7\x4f\xd9\x46\xe8\xcc\x08\x28\xa7\x8d\xdf\x54\xbb\xa8\x47\x95\x40\x95\xb0\x5d\x62\x33\x0a\xb6\x45\x24\x64\x14\xc5\x38\xb5\x0b\x84\x1e\x68\xa7\x3e\x94\x8b\x1a\xeb\xdb\x84\x49\x2e\x38\xbe\x06\xb6\x9c\x76\x25\x52\x45\x85\xb5\xb9\xaf\x6e\x8b\x0d\x38\x30\xfb\xda\xe2\x91\x34\x00\x27\xa4\x02\x38\x31\x1d\x1c\xdf\x7f\x9d\x12\x0e\xf9\x1f\x55\x9b\x06\x2d\xc0\x89\xa9\xa1\xa1\x1d\x70\x84\x6d\xde\x3f\xa9\x9e\xd0\x7f\x83\x63\xe0\x19\x78\x46\x1b\x22\xf1\xb3\x87\x6d\x0e\x10\x4d\xb5\xaa\x8b\x5e\x34\x28\x51\xc8\x76\xa8\x92\x38\x3a\xa0\x86\x1c\xb2\x1b\x60\xcf\xc8\x6a\x01\x27\x09\x1d\xf1\x01\x42\xb3\xfa\x7a\x9f\x1f\x64\x10\x2c\x5b\xbd\xc7\x01\x1f\x1b\x7b\x6e\xf6\x95\x83\xdf\xe2\x34\x26\x8c\x69\xb5\x6f\xdc\x1e\xa0\x26\xda\xe2\x6f\x6a\xf2\x4d\x4d\x9e\x41\x4d\xd4\xdc\xd7\x16\x1c\xf5\x30\x9b\xcf\xe6\xa9\xfd\xeb\x25\x3a\xde\x18\x24\x34\x24\xa2\x3c\x95\x1e\xbe\x67\x09\x5d\x1d\x2b\xf6\x35\xc1\x51\x58\xd7\xda\xff\xaa\xb4\x7d\x2a\xab\x91\xcb\xa9\x55\x56\x83\x52\x0d\xd6\x0e\xa6\x9e\xf5\xbf\x3a\xe4\x0d\xaa\x5e\xda\x07\xe0\xe5\x0d\xd1\x0e\x98\x3f\xec\x47\x2c\x5a\xb6\x1f\xf2\xb4\xda\x2e\x45\xe6\x66\xb6\x70\xbd\x25\xdc\xcc\x96\xf3\xae\xd7\x28\x47\x8a\xa4\xd3\xfb\x79\xc7\x69\x7c\x92\xb1\x8d\x85\x7b\xeb\x5e\x2e\xe1\x95\x03\xe6\x1d\xc3\x29\x78\x49\x84\x4d\x07\xd6\x28\x62\xd8\x01\xc1\x3e\xc6\x6f\x3f\xb8\x9e\x0b\xb3\xf9\x12\xdc\xdf\x6f\x16\xcb\x85\x55\xd8\xc0\x95\x37\xff\xa9\xf6\x3d\x45\x2d\x95\x62\xfd\x56\x77\xf8\x1c\x29\x9c\xeb\x29\x38\x70\x1d\x25\x2b\x14\x99\x2a\x87\xa3\x65\x52\xfa\x15\x19\x5d\x7a\xee\x74\xe9\xc2\xdc\x03\xcf\xfd\xf9\x76\x7a\xe9\xc2\xd5\xdd\xec\x72\x79\x33\x9f\xc1\x06\xab\x43\xa5\xaf\xa7\x67\xd9\xe0\xb9\xcb\x3b\x6f\xb6\x28\xc4\x7f\xba\x30\xce\xce\x8a\x04\xf4\xa5\x46\x33\x20\xa3\xab\xb4\xc2\x81\x68\xa1\xdb\xe9\xec\xfa\x6e\x7a\xed\xc2\xe2\x97\xdb\x31\xb1\x6d\x64\x3e\xa7\x0c\xb1\xac\x59\x47\xa8\x9d\x5f\x87\xb3\x8c\x84\xf9\x87\x3b\x69\xa7\x7d\x20\x16\x8f\xb4\xc3\x57\x92\x42\x8a\xb7\x11\x0a\x30\xac\x33\xaa\x68\x62\x83\x29\x4e\x11\xc7\xfe\x03\xcb\x44\x62\x4a\x43\x31\xcf\x52\xca\x20\xb8\x47\xa9\x75\xfe\x5a\xdd\x8c\x10\xdd\x64\x68\x83\x81\x7d\x88\x0c\xc4\x44\x98\x0c\x47\x62\xee\x59\xb6\x62\x3c\x25\x74\x63\x19\x0d\x7a\xcd\xe1\x2c\x9e\xf8\xd2\x57\x10\x25\xc1\x83\x5f\x1e\x2d\x2c\xdb\x01\x73\xb7\xdb\xed\xe2\x38\x0c\xef\xef\xcf\xff\x2f\x38\x29\x63\xa6\xdd\xf4\xa3\x5f\x9f\x3f\x43\xe9\x2f\x45\x34\x4c\x62\xcb\x86\xff\xc1\x2b\xfc\x9d\x03\xe6\xcb\xe2\xea\xf7\x21\x70\xc1\x74\xc0\x34\x6d\x47\x4c\xf5\xf9\x6b\xfb\xe2\xa2\xc8\xf6\x8d\x71\x76\xd6\x32\x89\x25\x89\x17\xf3\x57\x9f\x49\xf1\x16\x68\x1b\xbf\x4e\x6f\xef\xdc\x05\x58\xaf\x3a\xc6\xac\xbf\xab\x1c\x30\x3f\x7d\x69\xa7\x81\x3a\xc7\x96\x38\xa6\x70\xd6\x6e\x51\xd3\x81\xca\x40\x11\x82\x30\x19\x6a\x91\x3c\x51\x85\x31\xd4\xe0\x91\xe0\xa7\x31\xeb\x55\x7b\x8e\xb1\xd8\x66\xab\x88\x04\x7e\x01\x54\xb3\x6b\xfd\xff\x4a\x65\x99\x92\xc7\x16\xb0\x03\x46\x12\xae\xbd\xcc\x95\x92\x5a\xfb\xaf\x2f\x0e\x94\x3d\xa3\x32\x71\x8a\xf3\x97\xa3\x1d\x96\x6a\x6d\x63\xe9\x24\x42\x42\x8d\x73\x07\xd0\xad\x03\xc5\xfe\x3a\xa0\x5a\x43\xb4\x94\x01\x8a\x11\x7c\x12\x4e\x2e\xe0\x93\x1c\x8c\x49\xb2\x9d\x5c\xc0\x04\x7f\x98\x28\xbd\x9e\xc8\xb3\x95\xb8\xc5\x70\x11\xfd\xc4\x00\xf8\x62\xa8\x86\xec\xa1\x10\xfc\x88\xa2\x4c\x50\x48\xfe\x1f\x79\x01\x65\x95\xd9\xa9\x13\xa0\x03\x95\xdb\xe2\x03\x6d\xf9\x61\x46\xfd\xae\x53\x4f\xfe\x8a\x70\x80\x79\x34\x98\x17\xef\xc0\x24\xa1\x09\x2f\xde\xbd\x03\x53\x44\x84\x4d\xfb\xe2\x22\x3f\xda\xbe\x2d\xb0\x74\x52\x98\x7b\x60\x75\x39\x90\xd5\x30\x6d\x51\xe6\x2a\x70\xb3\xb4\x9e\xce\xbe\x2f\xa3\x7f\xab\xa5\x66\x0b\xea\xf8\x27\x00\x00\xff\xff\xe9\x25\x08\x77\xe6\x20\x00\x00")

func _0_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0_initUpSql,
		"0_init.up.sql",
	)
}

func _0_initUpSql() (*asset, error) {
	bytes, err := _0_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0_init.up.sql", size: 8422, mode: os.FileMode(0644), modTime: time.Unix(1635644171, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb6, 0x28, 0xa8, 0x1e, 0xc6, 0xb8, 0x3, 0xc0, 0x32, 0xf1, 0xd9, 0x81, 0x2c, 0x1f, 0x8a, 0x26, 0xa6, 0x2a, 0x6c, 0x4c, 0x5f, 0x9b, 0x24, 0x62, 0x1e, 0x35, 0xc7, 0xd0, 0x62, 0xf9, 0x29, 0x10}}
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
	"0_init.down.sql": _0_initDownSql,
	"0_init.up.sql":   _0_initUpSql,
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
	"0_init.down.sql": {_0_initDownSql, map[string]*bintree{}},
	"0_init.up.sql": {_0_initUpSql, map[string]*bintree{}},
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
