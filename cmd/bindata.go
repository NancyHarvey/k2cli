// Code generated by go-bindata.
// sources:
// data/kraken_store
// DO NOT EDIT!

package cmd

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

var _dataKraken_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\xb1\x6e\xdb\x30\x10\xdd\xfd\x15\xd7\x2c\xb4\xa0\xf0\x98\x99\x41\x87\x14\x15\x90\x22\x7d\x1c\x92\x8c\x02\x74\x03\x07\x2e\x04\x07\x4d\x04\x8c\x7c\x7b\x71\x94\x9d\xc8\xb2\x82\x6a\x30\xcf\xba\xc7\xc7\xf7\xde\x11\xa2\x9d\xe7\xfd\x79\xa0\x97\xd7\xa7\x97\x21\x1c\x0e\xdb\xde\x9f\x77\x7a\x7e\x7a\xa3\x5f\xc3\x10\xe8\x75\xf8\x3b\x3c\xbd\x0d\xbf\x7f\xdc\xa0\x2e\xcf\x34\x4d\xfb\x3d\xb6\x86\x88\xc4\xf0\x6e\xdb\x7d\x56\xe3\x6e\xff\xf4\x59\x3d\xfe\xa7\x7f\xda\xef\x3f\x74\xf4\x41\xc7\x87\xee\x9b\xfe\xa8\x3f\xc6\x5a\xa6\x71\x5f\xbf\x63\xf7\xa8\xa7\xdf\x93\xe1\x2b\x8d\x6c\x75\xd7\x34\xb1\x21\xd7\x64\x9c\x94\x6b\x9c\x56\x3c\x6c\xd8\x18\xcb\x13\x5b\x23\x13\x3b\x22\x37\x2a\x0b\x8d\x34\x31\x75\x87\x2f\x0f\x27\xc5\x92\xe2\xb4\x3c\x35\x59\x23\x91\xe1\xce\x1d\x56\x52\x47\xa7\x30\x2d\x96\xbe\x92\x29\x7f\xb7\x92\x35\xba\x73\xa6\x1d\x75\x67\x14\xc9\x4f\x9e\x26\xf9\xb0\xeb\x11\x38\x15\xb3\xac\x44\xd2\x51\x77\xf9\x23\x77\x77\x2c\x8b\xd1\xfb\x26\x89\xcd\xa8\x34\xc7\xc6\x7b\x66\x3f\xd2\xf1\xc2\xb2\x70\x8a\xb5\x1f\xb2\x9c\xba\x84\x63\xd8\x90\xd8\xa6\x51\x2d\x9f\xe8\x78\x15\xef\xd1\xc9\x99\x47\xd7\x39\xf4\xe7\xf7\x6d\xbf\xb1\x9b\x61\xcb\xb2\x18\xab\xa3\xe8\xe1\xcf\x1c\xd6\x7c\x73\x1d\xc5\x7a\x17\x51\x9d\x67\xb9\x42\x94\xda\xb6\xda\x12\x73\x00\x90\x81\x10\x72\xea\xb9\x59\xa9\xee\x0a\x6b\xe7\x98\x9c\x8b\x01\xb1\xf6\xd6\x02\x99\x7d\x5f\x33\x42\xf5\x2e\xc5\x6a\x37\x27\xcb\x0c\x20\x34\x53\x0c\x84\x96\x41\x01\x80\x5e\x36\x48\x8e\x08\x01\x79\xf6\x42\x1e\x00\x11\xbb\x14\x54\x4a\xdd\x22\x13\x82\xe5\x92\x81\x38\x47\x00\xb5\x6a\xe9\xa4\x07\xe6\x6b\xa4\x05\x8a\x32\xf7\xcd\x54\xcc\x68\x38\x22\x8a\x08\x1b\x52\x07\x78\x61\xeb\xfb\x0c\x24\x11\xd6\x45\x63\xa2\x02\xf8\x15\x54\x44\xe5\xc7\x25\x27\x20\xf7\xbe\xb9\x01\x42\xe3\x77\xf2\xc5\xcb\x29\xa5\x14\x03\x50\xe7\xe2\xbc\xaf\x0b\x0e\xbd\xf7\xae\xa8\xf0\x10\x6b\x4d\x2b\x73\xe2\x3d\x13\xf9\x80\xb9\xa5\xa5\xce\x55\x4d\xbd\x84\xe3\xf4\xfa\xf8\x6d\x18\xe4\x33\x92\x25\x49\xc0\xdc\x70\x28\x24\x2e\x23\xb8\x9b\x3b\xc0\x35\x20\xd5\x86\x23\x92\x08\xcc\x11\xa1\xee\x7c\x7f\xb8\xa8\xbb\x54\xe7\x79\xae\x9a\x70\x28\xdb\xb1\x2e\x82\x63\x2e\xa5\xa6\x18\x63\xaa\xa5\xe4\x7c\x23\xad\x09\x0f\x76\x3d\x91\x10\xf7\xbe\x77\xb6\xba\x6b\x4f\xd5\x1e\xfe\x05\x00\x00\xff\xff\xd0\xf5\xe8\x83\x97\x05\x00\x00")

func dataKraken_storeBytes() ([]byte, error) {
	return bindataRead(
		_dataKraken_store,
		"data/kraken_store",
	)
}

func dataKraken_store() (*asset, error) {
	bytes, err := dataKraken_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/kraken_store", size: 1431, mode: os.FileMode(420), modTime: time.Unix(1504625989, 0)}
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
	"data/kraken_store": dataKraken_store,
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
	"data": &bintree{nil, map[string]*bintree{
		"kraken_store": &bintree{dataKraken_store, map[string]*bintree{}},
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

