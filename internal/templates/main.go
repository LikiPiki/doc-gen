// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package main generated by go-bindata.// sources:
// templates/full/main.tex
// templates/full/text.tex
// templates/full/titul.tex
// templates/simple/main.tex
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesFullMainTex = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\xcd\x8e\xdb\x36\x10\x3e\x57\x4f\xa1\x4b\x6f\x82\x1b\x6f\x13\xa0\x17\xdf\x5a\xf4\x01\x7a\x5c\xf9\x40\x49\x23\x99\x5d\x91\x22\xc8\x51\x76\xbd\x84\x80\x5d\xa7\x48\x0f\x09\xd0\x6b\x4f\x29\x1a\xf4\x01\xd4\xed\x6e\xeb\xd6\x6b\xfb\x15\x86\x6f\x54\x90\x96\x13\xdb\x2d\x82\xbd\x98\x1f\xe7\xe7\x9b\x99\x8f\x63\xa5\x45\x93\xb7\x02\x24\xe6\x35\x33\xe6\x9c\x3d\x57\x4c\x81\x4e\xe2\xf1\x99\xc2\xa9\x65\x1a\x79\x5e\x43\x17\x45\x69\x6b\x40\xb1\xfc\x82\x55\x70\xde\x62\xf9\xd5\xd4\x72\xa9\x5a\x04\x99\x77\x47\x3e\x90\x55\xcd\xcd\x2c\x89\x75\x6b\x0c\x67\x72\x6a\x33\x96\x41\x7d\x14\x64\xb9\x2c\x40\x62\xc9\xb5\xc1\x63\xc7\x6c\xae\x40\x6b\x28\x8f\x0b\xda\x0a\x1a\x01\xa8\xe7\x5d\x94\xee\xa1\x7d\x09\x3a\x6b\x0c\x24\xfb\x8e\x51\x30\x5d\x71\x39\x39\xcb\x45\x92\x1d\xe0\x7a\xc0\x5f\xe6\x22\xd1\x03\x1e\x8f\x5e\xe4\xe2\xa4\x46\x59\x37\xec\xa4\x9b\xba\x91\x15\xb2\xcc\xcf\x7f\x38\x62\x30\xa9\xc6\x70\xe4\x8d\x9c\x64\x50\x37\x97\x53\x9b\x33\xe5\xaf\x5d\x94\x0e\xc8\x00\xb6\xea\xfc\x03\xc5\xd4\x9a\x0b\xae\x26\x63\x38\xad\x0b\xb2\x15\x1c\xbd\x39\x35\x80\x1e\xf1\x6b\xb0\xfe\x34\xa0\x26\xcf\x14\xee\x1c\x3e\x0c\x34\xc3\x13\xd7\x11\x15\x13\x46\x30\x9c\x25\x31\x13\xa6\x6c\x24\x9a\x80\xcc\x5c\x64\x5d\x94\x4a\xb8\xc4\x19\x34\x1a\x84\x1d\xce\xce\xd2\x7b\x7a\xa0\x8d\xbb\xa1\x07\x7a\xa4\xfe\x38\x48\xcc\x75\x5b\x43\x67\xe9\x17\x77\x43\x3d\xdd\xd1\x92\x56\xb4\x39\xa9\xa8\xaa\x52\xd5\x0d\x9a\x13\xb3\x28\x4a\xcd\x04\x14\x5d\x94\x8a\xa2\x2c\xa0\xe4\x12\x0c\xce\x6b\xd8\x97\x36\xdf\xf9\x5b\x67\xa3\xcf\xb8\x94\xa0\xb1\x51\xc3\xcb\xbc\x50\x98\x0c\xc6\xac\x41\x6c\xc4\xfe\xc5\x9e\x29\x8c\x76\x0d\x8a\x62\x20\x01\xf9\xf2\x3c\xb0\x4e\x8e\x58\xa7\x36\xfb\xe4\x80\x4f\xc8\x7f\xc2\xe4\x97\x9a\xa9\x92\x57\xc7\x0b\x53\x69\xa6\x66\x3c\xbf\xf2\x7b\xba\x83\x46\x31\x9c\x59\xcb\x05\xab\xc0\x7c\xd1\x75\x51\xfa\x35\xe4\x35\xd3\xf0\xed\xe0\xff\xe6\x0a\x41\x1a\xbf\x2f\x76\xa4\x8a\x32\x19\x29\x59\x25\xa3\xef\x55\xd5\x45\xd1\xe7\x31\xbd\x73\x0b\x77\xeb\x16\xee\x15\x6d\x63\xba\x8b\xc7\xee\x96\x1e\x3d\xa0\x35\xf5\xee\x47\xea\x69\x45\x0f\x31\x6d\xa9\x0f\xad\xfe\xe1\x7f\xdd\x0f\xd4\x87\x95\xa9\x41\x56\x38\x4b\x15\xd3\xbb\xff\x9b\x1d\xff\x67\xe9\x91\x5f\x5c\x7b\x53\x06\x15\x97\x76\xff\x29\xe8\xa2\xb4\x80\x32\xcd\x59\x69\xc7\xcf\x87\x8b\xd2\xa0\x9a\xc2\xd2\x6f\xf4\x33\xbd\x8b\x69\x1b\x34\xdd\xd2\x86\xee\x83\xae\x1f\x43\xae\x2d\xdd\xd3\x86\x56\xf4\x27\xad\x69\x13\x9a\x7f\x7b\x12\xff\x21\xa3\x66\x99\x64\x02\xac\x9f\x86\xfe\xa2\x3b\xea\x69\x4d\x4b\x3f\xd1\x8a\x7a\xfa\x3d\xbc\x5c\xef\x16\xe1\x5c\xd3\x86\xfe\x8e\xc3\x98\xde\xb1\x70\x6f\x3e\x92\x14\xdc\xe0\xff\x90\xdc\xd3\xd2\xdd\xba\xd7\xb4\xa4\x2d\xad\x68\x49\x6b\x9f\xe3\x55\x7d\x4f\xcb\xa0\xe9\xca\xbd\x0d\x42\xfe\x14\x87\x3e\x6f\x76\xa9\xee\xb5\x17\x30\x7c\xde\x2c\x72\x6c\xeb\x11\xc2\x95\x57\xc9\x00\xe6\x4d\x2b\x11\xb4\x55\xac\x82\xce\x9e\xed\xe8\x7e\x3d\xcc\x75\x6f\xe2\x20\xc0\x3f\xee\x15\x3d\xd2\x03\xad\xdd\xe2\x80\x0e\xae\x70\xcf\x86\x1c\x6b\xf0\x3c\x51\x94\x82\x2c\x0e\xd4\xff\x37\x00\x00\xff\xff\xc2\x9e\x15\xd6\x94\x05\x00\x00")

func templatesFullMainTexBytes() ([]byte, error) {
	return bindataRead(
		_templatesFullMainTex,
		"templates/full/main.tex",
	)
}

func templatesFullMainTex() (*asset, error) {
	bytes, err := templatesFullMainTexBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/full/main.tex", size: 1428, mode: os.FileMode(420), modTime: time.Unix(1600369955, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFullTextTex = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\x29\x49\xad\x28\x49\x4a\xab\x0e\xce\xcf\x4d\x55\x00\xb1\x15\x32\x52\x8b\x52\xf5\xf4\xf4\x6a\x01\x01\x00\x00\xff\xff\xae\x03\xeb\x47\x1a\x00\x00\x00")

func templatesFullTextTexBytes() ([]byte, error) {
	return bindataRead(
		_templatesFullTextTex,
		"templates/full/text.tex",
	)
}

func templatesFullTextTex() (*asset, error) {
	bytes, err := templatesFullTextTexBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/full/text.tex", size: 26, mode: os.FileMode(420), modTime: time.Unix(1596735266, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFullTitulTex = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\xdd\x6e\x13\x47\x14\xc7\xaf\x77\x9f\x62\xae\x72\x05\x56\x92\x06\xa4\xf6\x55\xba\xbd\x70\x9c\x85\x5a\x72\x4c\xe4\x04\x54\x61\x59\x5a\x3b\x51\x5b\x09\x24\x36\x09\x86\xc5\xc1\x8e\x3f\x10\x2d\xa2\x55\x1d\x93\x05\xab\x26\x6b\x89\x27\xf8\xcd\x2b\xe4\x49\xaa\x73\xbc\x4e\x4c\xec\x20\x5a\x75\x6f\xd6\x3b\x67\xce\xc7\xfc\xcf\xef\x8c\xbd\x75\xff\x6e\xbe\x58\xde\xc9\xef\x14\xfc\xad\xec\x5d\xbf\xe2\x3a\xe9\x52\xce\x2f\xee\xf8\xa5\x8a\xeb\x38\xbc\x24\xa2\x49\x44\x87\x2e\x75\xda\xfa\x3e\xa0\x65\x38\xe0\x2d\x1d\x7e\xa7\xce\xa1\x7c\xb6\xd8\xa7\x4d\xc8\x73\x5a\x1c\x10\xaa\xcf\x5f\x86\xc8\xd0\x24\xa4\x47\x43\x7e\xb6\x69\xd1\xa1\x43\xc4\x0b\x3a\x34\x68\xf1\xc2\xf0\x8a\x3a\x4f\x35\x76\xc8\x6b\x22\x22\x57\xf2\xbe\x22\xe6\x94\xd8\x06\xf4\x19\xd9\xc7\x9c\x91\x10\x1b\xde\x91\xd8\xaa\xdd\xe5\x94\xbe\x0d\x6c\xd5\xd6\x18\x10\x73\x36\x35\xf7\x19\xd8\x1a\x89\x7e\x7e\x9c\x2e\x26\x9c\x68\x98\x0f\x24\x0c\xe8\xdb\x1a\xf1\x4c\x48\xbb\x6b\x7f\xb1\x01\x31\xef\x25\x1f\x67\x0c\xc5\x67\x60\x1f\xd9\xaa\xfd\x95\x58\x32\x1a\xc6\x36\x20\xb1\x7b\xc4\xb6\x6a\xab\x0c\x35\xc3\x65\x5d\xba\xe5\x4a\x16\x09\x64\x9f\xc8\x51\x3e\x89\x4e\x22\x48\x83\xee\x4d\x8e\xa9\xa7\x4a\xee\xd3\xa3\xcd\xa1\x0a\x11\x89\x10\x87\xaa\x4e\x8f\xa7\x84\x17\x42\xd7\x69\xd2\xe4\xad\x98\x7b\xaa\xe9\x41\xda\x85\x48\xa3\x74\x0d\x21\x7f\xaa\xae\x0d\xf5\x96\x76\xfd\x46\x3d\x55\x57\x3b\x73\x4c\x9b\x88\x7d\x5a\xa9\xfc\x5d\x7d\xd7\x27\x1d\xfa\x34\x72\x1d\xc7\x7b\xb0\xbd\x95\xcd\xf9\xe5\x95\xcc\x2d\x7f\xb3\xa2\xf2\x37\x08\xd3\xc6\xb4\x09\xcd\xf9\xde\x33\x2f\x97\xbd\x33\xb3\x75\x4d\x36\x3a\x9e\x5f\xdc\xb8\x80\xc5\x75\xb4\x85\x92\xa8\x4b\x68\x78\x4e\xc8\x1b\x22\xde\x68\xae\xd0\xd0\x11\x4a\x5e\xeb\x97\xb6\xde\xbd\x88\xb6\x3c\x4d\xac\xc5\xd6\x39\xa6\xa5\x2a\x08\x48\x72\xcc\x23\xfe\xb8\xdc\xbc\x32\xd9\xea\x15\xef\x95\x36\xb3\x85\xed\xfc\x43\xbf\x2c\x85\xa5\x30\x67\xd7\xef\x17\xb2\xa5\x4a\x39\x97\xcb\x09\xbf\x8e\xb7\x55\xf2\xb7\xee\x6d\x3c\x34\x57\x9f\xa5\xb9\x95\x19\x53\xea\x35\x6f\xf2\x3c\xf3\xfd\xcd\x15\xff\xa7\x1f\x34\xf6\xfd\xe2\x86\x5f\x2a\xe4\x8b\x7e\xd9\xfb\x71\x52\xdc\xad\xdc\x66\xa5\x22\x01\xfe\x83\xc9\xf3\x24\x28\xa7\x24\x8c\x78\x2f\x64\x09\xe0\xf6\xf1\x0d\x65\x34\x63\x14\xf7\x98\xb1\x50\x2a\xab\x66\xc9\x30\x26\xe1\x94\x31\x43\x5b\x95\x25\x9d\x8b\x1a\xfd\xcf\x4f\xc3\x50\x69\xfc\x99\xa1\x42\xfb\xe8\x86\xb1\x7b\xf4\xf9\xc8\x90\x91\x50\x7a\x79\x38\x2d\xe0\xc3\x94\x5f\xe2\x85\xe2\x7c\x41\xb7\x6b\x1f\x0d\xad\xb8\x4c\x3b\xe4\x3a\x15\x59\x9c\x76\xf5\xf6\x04\xa8\xb9\xfb\xe7\x88\x3e\x27\x24\x32\x5a\x32\xd9\x36\x90\xc9\xb3\x4f\x8c\xce\x9a\x18\xe4\xb4\xe7\x7b\xcf\x56\x67\x39\x4e\x29\xf6\x0a\xd9\xf5\x62\x76\xd3\x5f\x60\x12\xe1\x44\x2d\xd1\x4d\x74\x19\x8b\x12\x9c\x11\x7f\x67\xce\x83\x97\xe2\xb8\x91\xdf\xde\x39\x0f\x9a\x73\x90\x4f\x23\x7d\xab\x05\xcf\x20\xdf\x9b\xdc\x89\xc2\xee\x91\x0e\xd7\xd1\x97\x98\x75\x17\x11\xab\xc8\xea\x88\xf6\xf4\x4a\x6c\xca\x80\x1f\xd2\xce\x98\x25\xb3\x72\x7b\x6d\xf5\xdf\xaa\xbe\x64\xe8\x64\x08\x33\x86\xfd\x09\x17\x8c\x88\x19\x5c\x4b\xf4\xe7\xbe\x0b\x40\xfd\xe6\x7a\x86\xd7\xbe\x0a\xef\xaf\x2f\xfc\x7f\x21\xdb\xbd\xca\x9c\xf7\xe0\x4e\xbe\x50\x58\x04\x5a\x47\x99\xff\xdb\xd6\xe4\x8e\x8e\x65\xd2\x6c\xc0\x89\xdd\xb5\x01\xef\x84\x98\xd5\xe5\xd5\xe5\x2b\x30\x4c\x62\x5f\xfe\x7d\xfe\x13\x00\x00\xff\xff\x80\xc6\x37\xa9\x51\x07\x00\x00")

func templatesFullTitulTexBytes() ([]byte, error) {
	return bindataRead(
		_templatesFullTitulTex,
		"templates/full/titul.tex",
	)
}

func templatesFullTitulTex() (*asset, error) {
	bytes, err := templatesFullTitulTexBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/full/titul.tex", size: 1873, mode: os.FileMode(420), modTime: time.Unix(1596734959, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSimpleMainTex = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\x31\xae\x03\x21\x0c\x04\xd0\x9e\x53\xfc\x03\xfc\x26\xa9\x72\x97\xdd\x2d\x0c\x0c\xac\x15\xe2\x20\x6c\x8a\x08\x71\xf7\x28\x45\xa4\xd0\xce\x1b\x69\x66\x8f\xcf\xd0\x1f\x10\x0b\x85\x54\xb7\xcb\xb5\xda\x31\xa8\x19\x87\x82\xe9\xdc\xde\x15\x95\xc2\x9d\x32\xb6\x6e\xe9\x76\x0c\x96\xda\x0d\x12\xe6\x62\x90\x5c\x58\xcf\xff\xbf\xd6\x55\x99\xe4\x18\x9e\x3c\xca\x52\x1a\x2c\x11\x62\x89\x9b\xda\x0a\xe7\xab\xa2\x35\xa4\xcf\xa0\x47\x66\x19\xdf\x57\xd3\x39\xb7\x43\xe2\x4f\xf0\x0e\x00\x00\xff\xff\x44\x74\xac\x1c\xb2\x00\x00\x00")

func templatesSimpleMainTexBytes() ([]byte, error) {
	return bindataRead(
		_templatesSimpleMainTex,
		"templates/simple/main.tex",
	)
}

func templatesSimpleMainTex() (*asset, error) {
	bytes, err := templatesSimpleMainTexBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/simple/main.tex", size: 178, mode: os.FileMode(420), modTime: time.Unix(1599245826, 0)}
	a := &asset{bytes: bytes, info: info}
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
	"templates/full/main.tex":   templatesFullMainTex,
	"templates/full/text.tex":   templatesFullTextTex,
	"templates/full/titul.tex":  templatesFullTitulTex,
	"templates/simple/main.tex": templatesSimpleMainTex,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
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
	"templates": &bintree{nil, map[string]*bintree{
		"full": &bintree{nil, map[string]*bintree{
			"main.tex":  &bintree{templatesFullMainTex, map[string]*bintree{}},
			"text.tex":  &bintree{templatesFullTextTex, map[string]*bintree{}},
			"titul.tex": &bintree{templatesFullTitulTex, map[string]*bintree{}},
		}},
		"simple": &bintree{nil, map[string]*bintree{
			"main.tex": &bintree{templatesSimpleMainTex, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
