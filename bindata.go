// Code generated by go-bindata.
// sources:
// templates/list.html
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

var _templatesListHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x58\x6b\x6f\xdb\xbc\x15\xfe\xec\xfd\x0a\x8e\x08\x60\x19\xb5\xa4\xc8\x97\x5c\x5c\xcb\x45\x6e\x6d\x52\x74\x69\x90\x16\xd8\xba\x61\x18\x68\x89\xb6\x98\x50\xa2\x4a\x52\x8e\x5d\x23\xff\x7d\x87\x12\x65\x2b\x89\x13\x74\x40\xf7\x7e\x78\x2d\xf2\xdc\x9e\x73\xe5\x49\xc7\x7f\x3d\xff\x7a\xf6\xfd\xc7\xcd\x05\x4a\x74\xca\x27\x7f\x19\x57\x3f\xad\x71\x42\x49\x0c\xbf\xad\x71\x4a\x35\x41\x51\x42\xa4\xa2\x3a\xc4\x85\x9e\xb9\x47\x78\x4b\xc8\x48\x4a\x43\xbc\x60\xf4\x21\x17\x52\x63\x14\x89\x4c\xd3\x0c\x18\x1f\x58\xac\x93\x30\xa6\x0b\x16\x51\xb7\x3c\x74\x11\xcb\x98\x66\x84\xbb\x2a\x22\x9c\x86\x41\x43\x4d\xa2\x75\xee\xd2\x9f\x05\x5b\x84\x78\xe9\x16\xc4\x8d\x44\x9a\x13\xcd\xa6\x9c\x36\x74\x32\x1a\xd2\x78\x4e\x41\xce\x08\x72\x96\xdd\xa3\x44\xd2\x59\xd8\x36\xe2\x6a\xe4\xfb\x33\xe0\x54\xde\x5c\x88\x39\xa7\x24\x67\xca\x03\x35\x7e\xa4\xd4\x87\x19\x49\x19\x5f\x85\x57\x19\xe8\x52\x82\x13\x4d\xda\x48\x52\x1e\xb6\x95\x5e\x71\xaa\x12\x4a\x75\x1b\xe9\x55\x4e\xc3\xb6\xa6\x4b\x6d\x64\xda\x93\x8d\x11\xc3\x89\xb7\x9c\xb8\xb2\x8a\x6b\xab\x29\x59\x46\x71\xe6\x4d\x85\xd0\x4a\x4b\x92\x9b\x83\x31\xbc\xb9\xf0\x07\xde\xbe\xb7\xef\x12\x9e\x27\xc4\xeb\x19\xe5\x5b\x9a\x97\x32\xe0\x56\x0a\x43\x74\x34\x9d\x4b\xa6\x57\x60\x2b\x21\xfd\xa3\x81\xbb\xea\xeb\xd9\xf2\xe4\x9f\xff\x28\x92\xc1\xe5\xc3\xb7\x1f\x2b\x3e\x3b\x7d\xf7\x39\xe8\x0d\xff\xb6\xbc\x52\x07\xe9\xed\xf0\xe3\xd7\x4b\x92\xde\x9c\x7e\xda\x3f\x18\xfc\x3a\x7d\x77\xf2\x91\xfe\xfd\xf2\x78\x70\x1d\x2f\xc8\xd9\x69\x7a\xf4\x33\x8b\x21\x74\x52\x28\x25\x24\x9b\xb3\x2c\xc4\x24\x13\xd9\x2a\x15\x85\xaa\x02\x5f\x3a\x54\x39\x8d\x6b\xa7\x4b\x4a\x6b\x2a\xe2\x15\x5a\x9b\xaf\xd6\x94\x44\xf7\x73\x29\x8a\x2c\x1e\x21\x39\x9f\x3a\xfd\x7e\x17\xf5\x0f\xbb\x68\xd0\xef\xbc\x2f\x19\x22\xc1\x85\xac\x68\xc1\x61\xd0\x45\xc1\xe1\x11\xfc\xef\x38\xb0\x64\x93\x11\xb7\x8a\xfe\x08\xb5\x9b\xf1\xef\x22\x7c\x49\xf9\x82\x6a\x16\x11\x74\x4d\x0b\x8a\x4b\x89\x47\x93\xdb\x96\x47\xa5\x14\xd2\xe5\x4c\x69\x0b\x24\x25\x12\x9c\x70\xb5\xc8\x47\xa8\xbf\x9f\x2f\x1b\xcc\x49\xb0\x83\xa7\x37\xb0\x3c\x15\x04\xc5\x7e\xd1\x11\x0a\xbc\xa1\xa4\xe9\xfb\x26\xf3\x54\x68\x2d\xd2\x06\xbf\x05\x10\x11\x19\xd7\x31\x10\x32\xa6\xd2\x6d\x78\xda\x1b\x74\x51\xef\x00\x22\x51\xbb\xd9\x0c\x13\x64\x35\x53\x39\x91\x50\xb2\xcf\x15\xba\xa6\xab\xa8\x7c\x25\xb6\x03\x08\xdf\x60\xd8\x45\xc3\x1d\xb1\xed\x05\xc6\x64\x00\x91\xef\xf5\x6a\xb2\xc5\x55\x79\xf0\x16\xbc\x1d\x08\x00\x60\x66\x61\x34\x33\x38\xec\x99\xe4\xed\x77\xd1\xf1\xf0\xa5\xe8\x0c\x6a\xf6\x55\xf0\x2f\x0a\xc3\xa2\x83\x64\xbc\x19\xb9\xb7\xeb\xc7\x9a\x9f\xea\xcc\x8d\x89\xbc\xdf\x61\xfb\xf6\xd3\xa9\x33\x84\xa0\x1d\x00\xe8\xc3\xe0\xa9\x6d\xab\xdb\xb0\xbc\x66\xd7\xd0\x82\x21\x10\x82\x83\x81\x31\x7e\xbc\xd3\xee\x28\x11\x8b\x57\x3d\x3f\x3c\xee\xa2\x23\xa3\x61\xff\x09\xe8\xa7\xa9\x6e\xa6\x0a\xea\x30\x5f\x22\x68\x03\x16\xbf\x91\xae\x5c\xd2\xd7\xea\x04\x5c\x1d\x80\xcc\xb0\xf7\x7b\x3d\x68\x8d\x4b\x12\xb3\x42\x8d\xd0\xb0\x6e\x8c\x9c\xc4\x31\xcb\xe6\x00\x07\xfa\xa9\x32\x0c\x63\xc1\x2f\xe7\x42\x35\x62\x55\x24\x59\xae\x9b\x33\xe2\x8e\x2c\x48\x75\x8b\x91\x92\xd1\x76\x08\xc2\xd0\xbb\x83\x79\xcb\x45\x11\xcf\x38\x14\x7f\x39\x01\xc9\x1d\x59\xfa\x9c\x4d\x95\x9f\x8a\x14\xfa\xc1\xbb\x53\x7e\xcf\x0b\x06\x5e\x50\x5f\x98\xe9\x77\x07\x63\x07\xcc\x96\x5a\x27\x7f\xd6\xec\xdd\xcf\x82\xca\x95\xdf\xf7\x02\x6f\xdf\x1e\xfe\xdf\x26\xcd\xcc\x32\x7e\x06\x5e\x0f\x6c\x96\xa7\xff\xdd\x62\x39\x86\xf7\x9c\x58\x44\x85\x09\x53\xc7\x93\x50\x4d\x2b\x67\x56\x64\x91\x66\x22\x73\x3a\xb6\x34\xf6\x1c\xec\x45\x40\xd3\x34\x76\x89\x76\x35\x4b\xa9\xd2\x24\xcd\x71\xc7\xa3\x24\x4a\x5e\x0a\xb4\x16\x44\x22\xca\x51\x88\xf6\x1c\x9d\x30\x05\x25\x52\xdd\x53\xee\x19\x10\x0e\xbe\xa5\x11\x65\x0b\x1a\x23\x8c\xde\xa1\x2a\x4b\x0e\x10\x89\xd6\xd2\xc1\x31\x58\x32\x46\x70\xa7\xe3\xcd\xa4\x48\xaf\xc5\x83\xd3\xb1\x65\xf6\x58\xeb\x32\xa0\x62\xca\xa9\x06\x36\x0f\x6c\xb7\x23\xce\xa2\x7b\x18\xf8\xbf\x8f\x86\xcd\x90\x03\x4f\xc5\x8c\xc9\xd4\xc1\x27\xd0\x09\x2b\x51\x20\x55\xd8\x8f\x07\x92\x41\xe0\x04\xbc\xcb\xd0\x44\x0b\x78\xc3\x40\x14\x95\x2f\xc6\x07\x00\x56\xeb\x6e\xed\x79\x26\x2b\x4e\x7d\x6c\x15\x92\xc3\x1b\xe4\x97\x8c\x7e\x1b\xbc\x03\xb7\xc0\x21\xe2\xb4\x59\xdc\xee\x74\x6b\x36\x93\x11\xe0\x3b\xbf\xf8\x72\xf1\xfd\xa2\xbd\xb9\x56\x45\x14\x51\x05\xfd\xf3\xd2\x0b\xf8\x8f\x8b\x88\x98\x4b\x48\x14\x17\x24\x76\x6c\x4c\x6c\x53\xd5\xe1\x69\x5c\x6c\xa2\x65\x42\x20\x72\x23\xab\x20\x0e\x9b\xb8\xf0\x82\x5e\xc3\x6e\x05\xf6\xfe\x65\x15\x18\x94\x35\x9a\x36\xc9\xf3\xed\x21\x27\x2b\x63\x74\x73\xb1\x36\x5b\x19\xb8\x60\x2b\xe3\x3f\x44\x43\xf0\x4d\x02\xe1\xae\xce\x60\xdb\xe2\xfa\x77\x85\xa6\x01\x06\xf6\x05\xfd\xc5\x3c\xbc\x21\xca\xe8\x03\x32\x9f\x4e\xbb\x0c\x9a\x02\x35\x16\xaa\x9d\x54\xe5\x4f\xa3\xac\xc7\xbe\xdd\x1b\xc7\x66\x87\x28\xeb\x3c\x66\x0b\xc4\xe2\x10\x57\x1a\xaa\xca\x1e\x57\xc3\xb1\xfc\xae\x38\x22\x4e\x94\x0a\xb1\x59\xf5\x08\xcb\xa8\xac\xf8\x9e\x12\xa5\x78\xa8\xaf\x9f\x09\x71\x37\x8d\xdd\xc1\x86\x08\xfa\x83\xc9\x85\xb1\x87\xce\x88\x8e\x92\xf1\x54\xfa\x93\xb1\x4a\x09\xe7\xb5\x8c\xa9\x75\x37\x2d\x20\x3a\x78\xf2\x5d\x16\x69\x4e\x35\xba\x82\x05\x4c\x12\xc8\x2d\x94\x14\x6c\x2a\x1e\xf8\x65\x44\xa0\x6d\x41\x5d\x6d\xd7\x07\xc3\x6f\x80\x38\x6a\x80\x98\x09\x99\x6e\x4e\x70\x64\x94\xc7\xb0\x44\xd7\x12\x86\xec\x9a\xb9\x9e\x6f\x65\xcc\xc6\x49\xa6\xd0\x11\x40\x84\x35\x90\x12\x19\x25\x78\xf2\xad\xfc\x1d\xfb\x25\xa9\xc9\xcb\xb2\xbc\x68\x8e\x10\xfc\x44\xb7\x09\xa6\x14\x1c\x59\x35\x28\xe7\x24\xa2\x89\xe0\x10\xf9\x10\x57\x3a\xb1\xdf\x00\xe8\xd7\x08\x1b\x77\x0d\x0f\xb7\x78\x91\x79\x17\x2b\xe4\x08\x0c\x80\x71\x7b\x20\x92\x11\xb7\x44\x09\x06\xcc\xdf\x03\x4d\xb0\xd3\x02\x1e\xbf\xcc\xa2\xad\x0e\x1b\xbc\xa6\xe6\x8c\x56\x54\xbf\xb8\x18\x99\xce\x74\xcd\x3d\xc4\x76\x53\xc7\x10\x8b\x92\x73\x85\x60\x4c\x31\x6a\xc6\xd4\x89\x1e\xfb\x95\xb6\x3f\x63\x0c\x3a\x6b\x6b\xe5\x24\xcf\x5f\x6a\x7f\x52\x03\x26\x6c\x8d\x3c\x37\x68\xdb\xcf\xaa\x2b\x6c\xbd\xef\x2c\x77\xb4\xdd\x77\xf1\xcb\xae\x68\x5c\xaf\xd7\xb0\x5d\xce\x29\xf2\x1e\x1f\x5f\x36\x88\xd9\xd2\x76\x77\xc8\x76\xf3\xc3\x93\xab\xf3\x11\x1a\x97\xfb\x9f\xa5\x32\x10\x5a\xaf\x91\x77\x75\x8e\x1e\x1f\xa1\xe8\x81\x34\x41\x27\x37\x37\xcf\xd8\xca\xc0\x18\x3e\x08\xca\x96\xf1\xf5\x86\x30\x36\xa7\x30\x17\xef\x1b\x2d\x61\x96\x1a\x4b\xb7\x43\xab\x52\x79\x53\x1d\x3e\x7f\xfb\x7a\x5d\xaa\x06\xbe\xad\xd0\x5b\xb9\x6c\xa6\x11\xd9\x47\xa7\xca\xa6\x99\x39\x1b\xaf\xf0\xe4\xd6\xbe\x15\xe5\x58\x78\x9e\xd3\xb7\x9d\xb0\x7b\xef\x76\x64\x54\x9f\x4b\xe5\xc2\x1f\x56\x49\xa3\xc8\xc7\x66\xb0\xa2\x7a\xc2\x56\xe6\xcf\xaa\xda\x3d\xd1\x06\xc5\x46\xed\x8e\x27\xfb\xa9\x52\xd4\xac\xf9\xe7\x7a\xc6\xbe\x91\xda\x05\xbe\xf1\xbd\x5e\x53\xae\x68\xb3\x4c\x36\x02\xf0\x87\xa7\xc8\xe6\x93\x4c\x54\x65\xa7\xcc\xce\x57\xde\xec\x52\x92\xc5\x56\x47\xb3\x9e\xed\x57\xfd\x01\xf1\x2c\xc7\x3d\x14\x7a\xf9\xcf\x07\xff\x0d\x00\x00\xff\xff\xd2\xb4\x80\x86\x56\x10\x00\x00")

func templatesListHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesListHtml,
		"templates/list.html",
	)
}

func templatesListHtml() (*asset, error) {
	bytes, err := templatesListHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/list.html", size: 4182, mode: os.FileMode(420), modTime: time.Unix(1468946085, 0)}
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
	"templates/list.html": templatesListHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"list.html": &bintree{templatesListHtml, map[string]*bintree{}},
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

