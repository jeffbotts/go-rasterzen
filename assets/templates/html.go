// Code generated by go-bindata.
// sources:
// templates/html/index.html
// DO NOT EDIT!

package templates

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

var _templatesHtmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\xc1\x8e\xd3\x30\x10\xbd\xef\x57\x0c\xb9\xa4\xd5\x76\xed\x6d\x2b\xba\x55\x37\x29\x42\x42\x08\x44\x85\x38\xac\x38\x80\x10\x9a\xd6\x53\xe2\x5d\xbb\xb1\x6c\xd3\x4d\x36\xca\xbf\x23\xc7\x49\x69\x6f\x7b\xcb\xbc\x79\xef\xf9\x8d\xc7\xc9\x0a\xaf\xd5\xfa\x0a\x20\x2b\x08\x45\xf8\x00\xc8\xbc\xf4\x8a\xd6\x16\x9d\x27\x2b\x32\x1e\xcb\xab\xd8\x53\xf2\xf0\x04\x96\x54\x9e\x38\x5f\x2b\x72\x05\x91\x4f\xa0\xb0\xb4\xcf\x13\xbe\x73\x8e\xf7\x32\xb6\x73\x2e\x01\xbe\x0e\xa2\x57\x2a\x15\xe1\x5e\x91\x1f\x94\x51\xe5\x76\x56\x1a\x0f\xce\xee\xf2\x84\x3f\xe2\x11\x23\x70\x22\x3f\xba\x64\x9d\xf1\x08\xbe\x4e\x72\x53\xa0\x2b\x2e\x75\x5d\xc8\x5e\x1c\xb2\x81\xaf\x0d\xe5\x89\xa7\xca\x87\x64\x49\x34\x86\xa6\x01\xb9\x07\xf6\x81\xb6\x7f\xff\x40\xdb\x46\x90\x0d\xb6\x5e\x2a\x82\x06\xb6\xa5\x15\x64\x57\x20\xd0\x15\x24\xc0\x92\x80\xa9\xa9\xe0\x8d\xd4\xa6\xb4\x1e\x0f\xfe\x1e\xda\x93\x1d\x1d\xc4\x60\x94\xf1\xee\xe8\x6e\x17\x7c\x58\x46\x28\xb6\xa5\xa8\xfb\xc9\x84\x3c\x82\x14\x79\xa2\xd1\x84\xf4\x42\x1e\x63\xe3\x7c\xee\x3e\x2b\x1c\xd1\x02\x1a\xf9\xfb\x89\x6a\xc8\x21\x6d\x1a\x60\x5f\xa9\xf2\x2f\x74\x78\xff\xed\xf3\x17\xaa\xa1\x6d\xd3\xfb\x33\xee\xbe\xb4\x1a\xfd\x40\x7d\x90\x8a\x3e\x46\xa4\xe3\x9d\x11\x35\x1a\xc8\x61\xc3\x34\x9a\x51\xaa\xd1\xa4\x63\xe6\xc8\x7f\x97\xf4\x3c\xfa\x39\xbf\x63\x8b\xe9\x7c\xf9\x76\x39\x81\x9b\xe9\x6c\xc6\xe6\x77\xf3\xc5\xe2\xd7\x04\xa6\xf3\xf1\xf9\x59\x0a\x6b\xb2\x9d\x49\xb8\xb4\x4d\xa8\x46\x69\xe1\xbd\x59\x71\xae\xca\x1d\xaa\xa2\x74\x7e\xb5\xbc\x5d\xde\xf2\x14\xae\x87\x68\xd7\x90\xf2\xe6\xa5\xe5\x4d\xd5\xf2\xa6\x6e\xd9\x65\xeb\x5d\x3f\x6c\x1e\xe0\xfe\x7b\x02\x8d\xc6\xea\x47\x59\xea\x15\xcc\x66\xed\xf8\x34\x47\x17\x80\xa1\x10\x0f\xe5\x48\xa3\x39\x35\x42\xb8\xf0\x3c\x20\x87\x03\x3d\xc3\x86\x7d\x42\x57\xf4\x0c\xf8\x7f\xd3\x17\xef\xad\x5b\x58\x5c\x52\xc6\xe3\xef\xf4\x2f\x00\x00\xff\xff\xb4\xf6\xcb\xfc\x56\x03\x00\x00")

func templatesHtmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlIndexHtml,
		"templates/html/index.html",
	)
}

func templatesHtmlIndexHtml() (*asset, error) {
	bytes, err := templatesHtmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/index.html", size: 854, mode: os.FileMode(420), modTime: time.Unix(1566429250, 0)}
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
	"templates/html/index.html": templatesHtmlIndexHtml,
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
		"html": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{templatesHtmlIndexHtml, map[string]*bintree{}},
		}},
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
