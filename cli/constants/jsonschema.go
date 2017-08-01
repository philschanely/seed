// Code generated by go-bindata.
// sources:
// ../spec/schema/seed.manifest.schema.json
// ../spec/schema/seed.metadata.schema.json
// DO NOT EDIT!

package constants

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

var _SpecSchemaSeedManifestSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x6f\xe3\x36\x10\xbe\xfb\x57\x08\xea\x1e\xfc\xa8\xd6\x5e\xa0\x97\xe4\x52\x2c\xb0\x3e\xf4\x90\xb4\x68\xda\x1c\x62\xbb\x01\x2d\x8d\x65\x26\x92\xa8\x52\x54\x0b\xb7\xf1\x7f\x2f\x28\x39\x89\x1e\xc3\x87\x9c\xd8\x69\x10\xef\x65\x9d\xe1\x3c\xbe\x21\xe7\x25\x89\xff\xf6\x1c\xc7\xfd\x94\xf9\x6b\x88\x89\x7b\xee\xb8\x6b\x21\xd2\xf3\xf1\xf8\x2e\x63\x89\x57\x52\x3f\x33\x1e\x8e\x03\x4e\x56\xc2\x9b\xfc\x30\x2e\x69\xdf\xb9\xdf\x4b\x39\xb1\x49\x41\x0a\xb1\xe5\x1d\xf8\xa2\xa4\x91\x20\xa0\x82\xb2\x84\x44\xbf\x70\x96\x02\x17\x14\x32\xf7\xdc\x59\x91\x28\x83\x82\x21\xad\x92\xa5\x79\xc7\x71\x33\x80\xe0\x1a\x78\x46\x59\xf2\x44\xac\xe8\xcf\x04\xa7\x49\x58\xe8\x2f\xe8\x29\x11\x02\xb8\x64\x75\xff\xb8\x9a\x4e\xbf\xdd\x5e\x4f\x7f\xbd\xfa\xe9\xe7\xcb\x4f\x6e\xc1\xb1\x2d\x19\xdd\x3b\xb6\xc4\xb4\x55\xd0\x16\x74\x13\xe2\xd2\x64\x1b\x75\x41\x4f\x48\x0c\x35\x8a\x1a\x76\x0b\xfa\x8c\x78\xff\x4c\xbc\xb3\x5b\x6f\x31\xda\x21\xaf\xa0\x2f\xa1\x45\x21\xe3\x54\xac\xe3\xf6\xe6\x74\xb2\xd4\x9f\x3c\xcc\xbe\x78\x67\x8b\xd9\xc4\x3b\x5b\x0c\x07\xf3\xf9\x67\x23\xa5\xef\xd5\x09\x0f\xe5\x7f\x12\xf3\x57\xef\xc6\x2b\xa8\x8f\xbf\x87\x83\x7e\x4b\x81\x9e\x7f\x30\x1c\xfc\xd8\x9f\xcf\x47\x55\xea\x48\x2a\xa9\x11\x24\x97\x62\x67\x52\xe2\xdf\x93\x10\x4e\xfb\xd2\xd8\x17\x41\x45\x64\x0a\x48\x54\x30\x80\xcc\xe7\x34\x15\xe6\xdd\xc4\xed\x92\x50\x25\x46\x38\x27\x9b\xfa\x19\x50\x01\x71\xd6\xe0\xd7\x18\x72\x9c\x2d\x9e\x1e\xb9\x58\x33\x7e\x69\x4e\x41\x8d\xf4\x34\x26\x34\xda\x5f\xfc\x77\xbe\x97\xb0\xa0\x31\xb0\x5c\xa8\x44\x69\x22\x20\x04\x8e\xcb\xca\x45\xbe\x22\xbe\xd2\xe9\x46\x81\x73\x6c\x8b\x9c\xa3\x29\x74\xc5\x9a\x1f\x07\x2d\xa2\xf6\xd8\x6a\xc0\x77\xe0\xd3\x5c\x7c\x23\x82\xe8\x14\x21\x0e\x74\x72\xc2\xe4\x48\xb1\xbe\xa2\x91\x62\x49\x1b\xbb\xcf\x9e\xa0\x31\x6c\xeb\xcb\x1e\x1e\xd9\xf9\xb5\xe3\x42\xfa\x92\x02\x20\x52\x23\x9b\x16\x1b\x3d\xeb\xab\x77\xd3\xe8\x59\xf5\x7f\x5b\x95\x2e\x97\xc3\x9f\x39\xe5\x80\x45\x11\x82\x6c\xc9\x58\x04\x24\xd1\x41\x0b\x60\x45\xf2\x48\xe6\x91\xe0\x39\x74\xc6\x13\x43\x40\xc9\x6f\xa5\x39\x0b\x40\xea\x68\xd8\x31\xea\x63\xa2\xa1\x0d\x4b\x97\x1a\xec\xee\xee\xe4\x91\xa0\x69\xab\xfe\x2b\xec\x77\xda\xde\x22\x22\x55\x80\x50\xba\x02\x66\x35\x06\x66\xda\xe8\xb5\x38\x35\x94\x63\x81\x50\xdb\x10\x11\x78\xae\x9c\x7a\x4f\x15\xe1\x43\x57\x04\x61\x5d\x0c\xcc\xbb\x04\x49\x1e\x6b\xa2\xbc\xe0\x31\x95\x14\xc7\xc6\x69\xe7\x79\x62\xd0\x32\x25\x79\xbc\x34\xf1\x68\xa3\x73\xc7\x63\xaa\x5c\x58\x06\x3a\x6f\x50\x28\xc4\x4b\x6b\x44\x4f\xf7\x77\x73\xb4\x61\xb9\x38\xcd\x36\x7b\x7a\x64\xe7\xd7\x8e\xeb\xff\x59\xc9\x3a\xce\x12\xfa\x1c\x52\x9b\x79\x06\x7c\x40\x23\x3e\xcb\x93\xe6\xc3\x89\xde\x84\x55\x45\x76\xbf\xd8\x1e\x45\xbf\x78\x4a\x1e\x3d\xcc\xe7\xc3\xc1\xfb\xe9\x2a\x6f\x36\x07\x19\xa3\xe5\x34\x28\xbd\xf7\xf2\x72\x0f\x9b\xc3\xe6\xfc\x69\xf2\x39\xe4\xe4\xf3\xc1\x4b\xd5\x71\x27\xb1\x58\xb6\x2f\x2c\xc7\x0d\xc5\x49\x53\x98\xcc\x45\xa9\x63\x41\x32\x17\x23\x7d\x21\x3a\x50\x11\x42\x43\x40\x6a\x58\x5b\x23\xb1\xd6\x1a\xb3\x40\xe3\x9f\xa1\x86\xb8\x9c\xa9\xc3\x8d\xff\xad\x08\x36\x45\x80\x57\x46\x14\xce\x50\x07\xac\xda\xa2\x21\x4f\xd4\x39\x52\x6e\x70\x6b\xa1\x99\x1b\xfa\xa8\xe7\x90\xb1\x9c\xfb\x68\x3c\x1d\xf7\xf1\x23\xf3\x49\x44\xf8\x69\x40\x38\xea\x80\xf0\x17\x89\x72\x4b\x58\xbb\x6e\xd8\xd9\x44\xf1\xf2\xfe\xa2\x7c\xc9\x48\x41\x75\xc0\xdd\x8c\x1d\xb9\x11\x95\xbb\x64\xdf\x89\xd0\xf0\x34\xc2\x78\xcc\x00\x64\xb1\x6d\xa5\x5b\x6f\xcb\x40\x08\x9a\x84\xa7\xee\xf6\x3a\x9d\x2d\x03\x9f\x83\xfa\x51\xd7\x72\xf0\xb2\x78\x51\xfe\x8a\x1d\xa4\x6b\x9f\xe8\x61\xbf\xab\x5f\x14\x81\x73\xc6\x2f\x48\x9a\xca\xad\x7e\xb5\xaf\xb8\x68\x2c\x75\x88\x23\x7d\x0c\xb9\xbe\x6a\x7a\x50\x7f\x45\x6d\x79\xfe\x28\x80\x7c\x34\x6f\xe8\xc2\x67\x9b\xb6\x2a\xf5\x67\xf4\x3d\x15\xfa\x44\x40\xc8\x38\xfe\xf0\x67\x4e\x93\xea\x68\xf3\x74\x9b\x04\xe3\xd3\x8c\x5b\x5a\x41\x69\x82\x08\x82\xaf\x64\x9b\x4c\x40\xfc\xf2\xc1\x46\x99\x14\x65\x14\x28\x0e\xb4\xa7\x32\x59\xc9\x82\x5e\xc3\x24\x6e\xac\xd9\x54\xda\x17\x73\xd4\x57\x53\x5a\x97\x33\x14\x97\x2e\xf0\x7b\x0d\xc8\x9d\x01\xf4\x2a\x40\xaf\xea\xa5\xf4\xaa\xf0\xa8\xed\x4d\xed\xaa\x55\xa9\xa9\xb8\x27\xd5\x93\xb2\xdb\xff\x02\x00\x00\xff\xff\x23\x27\xca\x8f\x0e\x26\x00\x00")

func SpecSchemaSeedManifestSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_SpecSchemaSeedManifestSchemaJson,
		"../spec/schema/seed.manifest.schema.json",
	)
}

func SpecSchemaSeedManifestSchemaJson() (*asset, error) {
	bytes, err := SpecSchemaSeedManifestSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../spec/schema/seed.manifest.schema.json", size: 9742, mode: os.FileMode(448), modTime: time.Unix(1501609041, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _SpecSchemaSeedMetadataSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\xdf\x4f\xdb\x30\x10\x7e\xcf\x5f\x71\x32\x7d\x00\xd1\x52\x84\x78\x59\x5f\x10\xd2\xd0\xc4\xb4\x51\x44\x27\x1e\x06\xdd\x64\x9a\x4b\x6b\x94\xd8\xc1\x76\x85\x2a\xd6\xff\x7d\xb2\xf3\xa3\x76\x9a\x42\x18\x1d\xe4\x09\xe7\x7c\x77\xfe\xee\xbb\xb3\xfb\xf1\x14\x00\x90\x8e\x9a\xcc\x30\xa1\x64\x00\x64\xa6\x75\x3a\xe8\xf7\xef\x95\xe0\xbd\xcc\x7a\x20\xe4\xb4\x1f\x4a\x1a\xe9\xde\xe1\x71\x3f\xb3\xed\x90\xae\x89\x63\xe1\x33\x21\x53\x14\xc6\x64\xfe\x26\xa8\xe5\xe2\xc0\x7c\xe5\x81\x9a\xe9\x18\x4d\x6c\xb1\x99\x99\x43\x54\x13\xc9\x52\xcd\x04\x37\x9b\x43\x8e\x50\x38\x00\x55\x10\x62\xc4\x38\x86\x70\xb7\x80\x2f\x28\xbe\x8e\x86\x17\x79\xb6\x45\x6a\x93\x89\xbb\x7b\x9c\xe8\xcc\x96\x4a\x91\xa2\xd4\x0c\x15\x19\x80\x29\x12\x80\x28\xc4\xf0\x1a\xa5\xca\xd2\x67\x46\x27\x5c\x69\xc9\xf8\xd4\x86\x5b\x7b\x4a\xb5\x46\x69\x91\xfc\x1a\x9d\x9d\x7d\xfe\x7d\x7d\x76\x35\x3a\x1f\x5e\x74\x88\xf5\x58\x66\x8e\xab\x12\x6a\x52\x3a\x88\xac\x5d\xe2\xc3\x9c\x49\x34\xb4\xdd\xe4\xb6\xc2\xbb\xbb\xfa\x9e\x08\x21\x43\xc6\xa9\x46\x45\x72\xeb\xb8\x4c\x21\x38\x0e\x23\x2f\xfe\xa9\x5c\xb9\xcc\x5e\x0a\xc6\xb5\x93\x15\x80\xd0\x30\x64\x86\x5b\x1a\x5f\xba\xec\x44\x34\x56\xe8\x39\xd6\x90\xe7\x83\xad\x5a\x01\x08\xf2\x79\xe2\xc1\x2a\x77\x32\x24\x15\xfb\xd8\xfb\x5e\x76\xfd\x43\x5c\x06\x6a\xce\xea\x48\x34\x14\x90\x9d\xbe\x1d\x09\x5b\x94\xea\xa7\x42\xd9\x95\x7f\xd4\x32\xa8\x5b\x3b\x07\xd6\xd3\xf7\x7d\x1e\x6b\xd6\x1e\x0e\x1d\x38\xef\x47\xe4\xa9\x94\x74\xb1\x15\x36\xbf\x31\x8e\x23\xff\x7a\x7d\x24\x9b\x0e\x9c\xff\xcf\x66\xbc\xe1\xb0\xb7\x0c\x66\xcb\xf8\xac\x62\xda\x26\xa9\xc5\x53\x4a\xed\x30\x76\xab\xdb\x4c\x63\x52\x17\xf7\xea\x76\xf8\x0d\xf9\xe7\xf6\x5c\x8a\x78\x31\x15\xbc\x15\x6d\x29\xb0\xbc\xc7\x8b\x51\x73\xd2\xdb\x5e\xde\xf6\xd0\xe8\x01\x6a\xf3\x68\xd7\x76\xa1\xe1\x5c\x07\x6e\x3d\x85\xb4\xd1\x2c\xc1\x26\xb2\xa6\x49\x63\x36\x35\x85\x28\x4d\xa5\xae\x94\xb9\x49\x93\x41\x55\x97\xdd\x1c\xf6\x3e\x8d\x9f\x8e\x97\xbd\x6c\x71\xb4\x5a\xfc\x28\x16\x83\xb5\xc5\xee\xed\xed\x81\x5d\xef\xef\x9d\xec\xee\xde\xec\xf7\xc6\x6b\x2e\x7b\x7f\x7e\xee\x9d\x74\x48\xdd\xd4\x12\xe4\x61\xeb\xe1\x06\x15\xd8\x1b\x94\x67\xc6\x7d\xa5\x36\x7f\x14\x82\x3c\xcb\x7a\x06\x4f\x4f\x07\xb9\x46\x25\xce\x40\xae\x94\x77\x29\xcd\x9c\x61\xaa\x88\xfd\x53\x50\x8c\x4f\x63\x84\xd2\xb7\x5b\x1d\x3b\xff\x9e\x90\x84\xf1\xf3\xfc\x8a\x1c\x95\xc6\xe2\xd2\x6c\x92\xc7\x79\x2a\x3e\x4f\xee\x50\xd6\xf6\xb7\x99\xff\x9a\x2e\x5f\xdd\x81\x02\x94\x1d\x7f\xef\x3e\xf9\xc2\xea\x19\x2e\x38\xd8\x5a\x41\x44\x25\x1d\xea\x45\x3e\xd6\xdf\x8b\x86\x42\x79\xe9\x81\x74\x7e\x20\x9b\x21\xd4\x8f\x02\x84\x84\x44\x48\xac\x43\x4b\xe3\xf8\xb9\x7f\x58\x5e\x25\x41\x37\xb5\xc9\x1d\x85\x17\xde\x34\x53\x1e\x95\x57\xcd\xcb\x8b\xc4\x5c\xae\xea\x82\xc7\x19\x4a\x04\x3d\x43\x88\x98\x54\x1a\xf0\x61\x4e\x63\x65\x0d\x31\x55\xfa\xc3\xea\x3e\x7e\xa1\xee\xe2\xc7\xa1\x59\xd1\x19\x4b\x60\xa6\x60\x8b\x83\xe7\x70\xef\x8f\x9e\x79\x64\x82\x65\xf0\x37\x00\x00\xff\xff\xbb\x44\x9e\x3d\x8d\x10\x00\x00")

func SpecSchemaSeedMetadataSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_SpecSchemaSeedMetadataSchemaJson,
		"../spec/schema/seed.metadata.schema.json",
	)
}

func SpecSchemaSeedMetadataSchemaJson() (*asset, error) {
	bytes, err := SpecSchemaSeedMetadataSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../spec/schema/seed.metadata.schema.json", size: 4237, mode: os.FileMode(448), modTime: time.Unix(1501609041, 0)}
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
	"../spec/schema/seed.manifest.schema.json": SpecSchemaSeedManifestSchemaJson,
	"../spec/schema/seed.metadata.schema.json": SpecSchemaSeedMetadataSchemaJson,
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
	"..": &bintree{nil, map[string]*bintree{
		"spec": &bintree{nil, map[string]*bintree{
			"schema": &bintree{nil, map[string]*bintree{
				"seed.manifest.schema.json": &bintree{SpecSchemaSeedManifestSchemaJson, map[string]*bintree{}},
				"seed.metadata.schema.json": &bintree{SpecSchemaSeedMetadataSchemaJson, map[string]*bintree{}},
			}},
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

