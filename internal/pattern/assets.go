// Code generated by go-bindata. DO NOT EDIT. @generated
// sources:
// DefaultHttpPattern.json
// DefaultChannelPattern.json
package pattern

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

var _defaulthttppatternJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\xcf\x6f\xdb\x3a\x0c\xbe\xe7\xaf\x60\x85\xa0\x78\x05\x5a\xa7\xef\xbd\x9d\x02\xe4\xb0\x76\xc5\x86\x6e\x05\x86\x22\x5b\x0f\xc3\x0e\x8c\xcc\xc4\x6a\x14\xcb\x93\xe8\x64\x59\x9b\xff\x7d\xb0\x62\xbb\x8e\xeb\x38\x69\x33\xec\xd2\x1f\x22\x29\x7d\x1f\xf9\x91\xf4\x43\x07\x40\xc4\x38\x23\xd1\x07\xf1\x8e\xc6\x98\x6a\xfe\xc0\x9c\x7c\x46\x66\xb2\xb1\x38\xcd\xec\x8e\x29\x71\xa2\x0f\xdf\x3a\x00\x00\x0f\xfe\x27\x80\x50\xe3\x2c\xa8\x1b\x48\x13\x8f\x83\xd4\xd1\x2d\x32\x7d\x52\x33\xc5\x64\x61\x30\x00\xb6\x29\xf9\x78\xef\xec\xc8\xce\x95\xf4\xcf\x54\xfc\x9e\xec\x2a\x4e\x52\x16\xfd\xf2\x76\x00\xc1\x66\x4a\x71\x16\x30\xd1\x66\x84\x5a\xe4\x96\x95\xff\xbd\x3a\x6d\x07\x73\x7d\x37\x6c\x05\x71\x7d\x37\xfc\x8a\x5a\x85\xc8\x66\x4f\x14\x83\x6e\x90\xe0\x52\x1b\x0c\x83\x88\x30\x24\xeb\x82\xb7\x29\x47\xc6\xaa\x5f\xc8\xca\xc4\xe5\x2d\x00\x62\x4a\xcb\x3c\xc4\xe3\xb9\x5f\xf0\x47\x5a\xbe\x8c\xc0\xa5\xb2\x32\x55\x7c\x61\x09\xa7\x3b\x12\xba\xe9\x2a\x5e\x92\x9f\x31\x6a\x47\xf0\xf8\x08\xdd\xa0\x9a\x91\xc0\xa4\x9c\xa4\xec\x82\x79\x76\xd2\xfa\x78\xa6\x96\x0b\x94\x53\x8a\xc3\x27\x7b\x84\x3a\x4b\xa3\xf8\xa7\x1b\x54\xec\x01\x59\x6b\x2c\x1c\x0d\x20\x56\xfa\x04\x8e\x8f\xe1\xc8\x9f\x04\xca\xc5\xc4\xfe\xcf\xa6\x80\x93\xfd\x18\x35\x27\x2c\x7b\x65\x3b\x88\x7d\xd2\xd9\x26\x0e\x93\x90\x5d\x17\xbf\x0f\x42\x9a\x34\xe6\x22\xff\x87\x95\x79\x1b\xea\xc1\x9f\x47\x6d\xc9\x11\xd7\x30\x77\x00\xbe\xfb\xce\xb7\xe4\x12\x13\x3b\x6a\xe9\xfe\x4a\x37\x97\xaa\xd1\xd9\xff\xb7\x84\x32\xa2\x06\xf1\x78\x2a\xa2\xef\x4f\xcb\xc3\x75\xe8\x26\x4e\x69\xc2\x8c\xda\x9b\xf3\xff\x2b\xad\x15\x22\xe3\x86\x9b\x1f\x50\xc8\xa9\x2b\x66\x0b\x78\x38\x70\xf5\x53\x12\x85\x14\xc2\x19\x0c\x23\x82\x3c\x55\xb0\x34\x29\x44\x38\x27\xb0\xf4\x23\x25\xc7\x14\x82\x72\x60\xe6\x64\x81\x23\x02\xd4\xda\x2c\x28\x04\x4f\x21\x10\xe5\x2b\xab\xd7\x4e\x9f\x75\x29\xdb\x9b\xcb\x77\xe1\x21\x09\xfa\x77\x47\x82\x8a\x1b\xb3\x89\xd4\x08\xc5\x3b\xdc\x90\x73\x38\xa1\xd7\x92\x6e\x93\xf2\xa6\xad\x7c\x96\xad\x4a\x92\xbf\xa3\x91\x32\x05\x72\x0d\x05\x46\x39\xce\x1c\xc3\xbe\xac\x8b\x7b\x7c\xcd\xf6\x81\xf6\xdf\xf9\xf9\x73\x68\x59\x21\xaa\xed\x5d\x24\xc4\x9b\xb7\x34\x63\xae\xe0\x86\x5e\x2c\x16\x78\xe3\x66\x0d\xc9\x49\xab\x92\xa2\xdf\x7d\x87\xe8\xba\x93\x25\x5f\xcc\x89\xe2\x28\x1d\x05\xd2\xcc\x7a\x89\x35\xf7\x24\xf9\x6c\xac\xcd\xc4\xf4\x66\x4a\x5a\x33\x41\xa6\x05\x2e\x7b\x28\x59\xcd\x15\x2f\x7b\x16\x99\x9e\x5d\xe5\x88\x59\xc5\x13\xb7\x99\x0c\xef\x56\xdd\x89\xb6\xc0\xda\x3e\x2f\x0b\x6a\xcd\xfb\xba\xc6\x2d\xf7\x20\x70\x66\x46\xe0\x37\xb7\x3b\x94\xe3\xfd\x82\x9b\xf7\x4f\x81\x6c\xdb\xe0\xad\x61\xbb\xac\xc9\xae\x18\xdd\x07\xc2\xcb\xd5\x3c\xaa\xbf\xde\x5c\x85\xd9\x5a\x92\xa2\xae\xb1\x2d\xd4\x1a\x77\x7b\x8d\xd7\x0d\x4e\x09\x30\x86\x88\x39\x01\x89\x5a\x03\x9b\x6c\xc8\x5a\x18\xd5\x23\x77\x31\x94\x26\x66\xab\x46\x15\x7d\x91\xe3\x9d\x94\x88\x23\x13\x66\xf7\xbe\xbf\x1a\x56\xbf\xc1\x52\xab\xaa\x7a\xcb\xd1\x7c\xb1\xf5\x0f\xc9\xac\xbf\x3a\xab\xce\xef\x00\x00\x00\xff\xff\xa7\x98\x63\xb1\x0a\x0b\x00\x00")

func defaulthttppatternJsonBytes() ([]byte, error) {
	return bindataRead(
		_defaulthttppatternJson,
		"DefaultHttpPattern.json",
	)
}

func defaulthttppatternJson() (*asset, error) {
	bytes, err := defaulthttppatternJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "DefaultHttpPattern.json", size: 2826, mode: os.FileMode(420), modTime: time.Unix(1541189812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _defaultchannelpatternJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x55\xcd\x6e\xdb\x3c\x10\xbc\xfb\x29\x36\x84\x91\x53\x3e\x39\x5f\xdb\x93\x01\x1f\x92\xf4\x94\x22\x40\x81\x06\xcd\xa1\xe8\x61\x4d\xad\x6d\xda\x32\x29\x90\x4b\x1b\x6a\xe2\x77\x2f\x44\xfd\xd8\x12\x64\x45\x81\x7b\xb1\x01\x72\x7f\x66\x66\x67\xc5\xd7\x11\x80\xd0\xb8\x25\x31\x05\xf1\x95\x16\xe8\x13\x7e\x58\xa1\xd6\x94\x7c\x47\x66\xb2\x5a\xdc\xe4\x21\x8e\x29\x75\x62\x0a\xbf\x46\x00\x00\xaf\xe1\x17\x40\xa8\x45\x9e\x37\x8e\xa4\xd1\x8b\xc8\x3b\x7a\x7c\x79\x86\xd9\x0c\xd8\x7a\x0a\x79\x21\xc8\x91\xdd\x29\x19\x3a\x3c\xbe\x3c\xff\xc4\x44\xc5\xc8\xc6\x1e\x03\x94\x4e\x3d\x8b\x69\x5d\x16\x40\xb0\xd9\x90\xce\x33\x66\xe3\x28\xc5\x2c\x31\x18\x47\x2b\xc2\x98\xac\x8b\xee\x3c\xaf\x8c\x55\x7f\x90\x95\xd1\x75\x15\x00\xb1\xa1\xac\x4c\x09\x78\xd6\x7b\xfe\x46\x99\x28\xef\x0f\xe1\xff\x70\xd3\x4f\xe0\x41\x59\xe9\x15\xdf\x5b\xc2\x0d\xd9\x5e\x2e\xcd\x50\x31\xa8\x7c\xa9\xcf\x02\x13\x47\xf0\xf6\x06\xe3\xe8\x54\x91\xc8\x78\x4e\x3d\xbb\x68\x97\x9f\xf4\x36\xbf\x93\x39\xf9\x7b\x94\x1b\xd2\x71\xaf\x92\xe1\xc8\x35\xce\x00\x84\x2c\x66\x9c\x97\x62\x72\x7c\xa2\x22\x80\xd8\x61\xe2\xa9\xbe\xaa\x6f\x0e\x95\x92\x75\xbb\x15\x26\x79\x37\x11\xf8\x0c\x53\xa0\x5b\x60\xb8\xbe\x86\x71\xd4\x20\x15\x91\xb5\xc6\xc2\xd5\x0c\xb4\x4a\x86\x0c\xa0\x4f\x04\x93\x92\x2d\xec\x32\x05\x21\x8d\xd7\x5c\x4d\xec\x32\x63\x9c\xc7\x3d\xfb\xf7\xb8\x2d\x39\xe2\x16\xea\x11\xc0\xef\xb0\xa0\x96\x5c\x6a\xb4\xa3\x8f\x2d\x69\x81\xbf\xdf\x83\xc5\x70\x6b\x94\x81\x9e\x98\x86\xf4\xfa\xb0\x48\x6a\x62\x97\x26\xce\xe9\x7e\xb9\xfd\xff\x64\x45\x63\x64\x6c\x7b\xb1\xaa\x98\x2f\x6e\x27\x94\x10\xf0\x44\xce\xe1\x92\x3a\xec\x78\xf1\xfc\x9a\x77\x75\x5b\xb6\x2a\x4d\xa9\x63\x0f\x3f\x2e\xc1\xe7\xa1\x12\xc8\x02\x0a\xcc\x4b\x9c\x25\x86\xa1\xac\xab\x3a\x61\x66\x43\xa0\x7d\xba\xbd\x7d\x07\x5a\x65\xac\x1c\xdd\x0f\x2f\x25\x39\x77\x75\x16\x4e\x6d\xc7\xd2\xee\x1d\x6e\xac\x1e\x9b\xee\xa7\x20\x26\x27\xad\x4a\x2b\xcb\x97\x11\x04\xce\x6c\x09\xc2\xa3\xe0\x8e\xc1\x96\xc2\x90\x97\x8a\x57\x7e\x1e\x49\xb3\x9d\xa4\xd6\xac\x49\xf2\x7f\x8b\xc4\x2c\xcd\x64\xab\xa4\x35\x4b\x64\xda\x63\x36\x41\xc9\x6a\xa7\x38\x9b\xac\xf7\xdc\xfd\xa9\xaa\x90\x9d\xdb\xcf\x16\xb6\x87\xd6\xa8\xaa\x0d\xbf\x10\x5e\xe9\x80\x79\xbb\xbb\x23\x66\xa5\x97\xcd\x2f\xb9\xd8\x16\x63\x14\xd8\xff\x35\xab\xa8\x9d\x79\x36\x5a\xcc\x9e\x70\x43\x80\x1a\x30\x44\x83\xc4\x24\x01\x36\x90\x19\x6f\x61\xde\xce\x7d\x8f\xa5\x34\x9a\xad\x9a\x9f\x10\x2c\x9f\x9e\xa3\x63\x46\x87\xd1\xdf\x00\x00\x00\xff\xff\xd9\x8a\xd0\x63\x88\x08\x00\x00")

func defaultchannelpatternJsonBytes() ([]byte, error) {
	return bindataRead(
		_defaultchannelpatternJson,
		"DefaultChannelPattern.json",
	)
}

func defaultchannelpatternJson() (*asset, error) {
	bytes, err := defaultchannelpatternJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "DefaultChannelPattern.json", size: 2184, mode: os.FileMode(420), modTime: time.Unix(1541019310, 0)}
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
	"DefaultHttpPattern.json": defaulthttppatternJson,
	"DefaultChannelPattern.json": defaultchannelpatternJson,
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
	"DefaultChannelPattern.json": &bintree{defaultchannelpatternJson, map[string]*bintree{}},
	"DefaultHttpPattern.json": &bintree{defaulthttppatternJson, map[string]*bintree{}},
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

