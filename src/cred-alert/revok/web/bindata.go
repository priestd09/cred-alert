package web

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _web_templates_index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x92\x31\x8f\xe3\x20\x10\x85\xfb\xfc\x0a\x8e\x3e\x4c\xa4\x48\x57\x9c\x30\xc5\xf9\xae\xbd\xa4\xb8\x66\x4b\x62\x66\x03\x5a\x0c\x11\x4c\x22\x79\x91\xff\xfb\xca\x26\xeb\x38\xa9\xfc\xfc\x3c\x33\xdf\xd3\x8c\xe5\x8f\x3f\x87\xf6\xff\xdb\xf1\x2f\xb3\xd4\x7b\xb5\x91\xd3\x83\x79\x1d\xce\x0d\xc7\xc0\xd5\x86\x31\x69\x51\x9b\x49\x30\x26\xc9\x91\x47\xd5\xb6\x47\x09\x55\x56\xdb\xbb\xf0\xc1\x12\xfa\x86\x67\x1a\x3c\x66\x8b\x48\x9c\xd9\x84\xef\x0d\xb7\x44\x97\x5f\x00\xc3\xd5\x89\x41\xdb\x18\xf5\xc5\x65\xd1\xc5\x1e\x2e\xd7\x84\xb0\x13\x3f\xc5\x6e\x96\xdb\xde\x05\xd1\xe5\x5c\x99\xf0\x0d\x95\xa7\x68\x86\x3b\xc6\xb8\x1b\x73\xa6\xe1\x5d\x0c\x84\x81\x78\xb5\xa7\x84\x7b\xd5\x26\x34\x18\xc8\x69\xcf\xda\x78\x0d\x94\xd9\xef\x81\x1d\xd2\x59\x07\xf7\xa9\xc9\xc5\x20\xc1\xee\x97\x06\xd2\x27\x8f\xac\xf3\x3a\xe7\x86\xcf\xf4\xea\x3c\xe4\xf6\x14\x93\xc1\x84\x66\xa1\x4c\x6d\x8f\x55\x2c\x8e\x7a\x86\x90\x7d\xfd\xfe\x9a\xec\xb9\x66\x7a\x9b\x87\x2e\x4e\x29\x49\x87\x33\x32\x31\x8e\x2b\x72\x7a\x1e\x6b\x94\xd4\xf7\x05\x43\x5c\x05\xc8\x50\x8a\xf8\xa7\x7b\x1c\x47\xae\x16\x29\x41\x2b\x09\x64\x5e\x67\x94\x22\x1e\xe9\xe6\x70\x53\xed\xba\x4e\xc2\x9a\x5c\x0a\x06\xb3\xc4\x92\x30\x6f\xea\x7e\x1c\x30\xee\x56\x4f\x57\x2f\x26\xa1\xfe\x51\x5f\x01\x00\x00\xff\xff\x90\xb7\x12\x2b\x62\x02\x00\x00")

func web_templates_index_html() ([]byte, error) {
	return bindata_read(
		_web_templates_index_html,
		"web/templates/index.html",
	)
}

var _web_templates_organization_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x92\x41\x8f\xdb\x20\x10\x85\xef\xfb\x2b\x28\xf7\x65\x22\x45\xea\xa1\xc2\x1c\xea\xf6\xda\x44\x55\x2f\x3d\x12\x33\x0d\xa8\x18\x2c\x98\xa4\x72\x91\xff\xfb\xca\x26\x71\x9c\x9c\xfc\xfc\x18\xe6\x7d\x9a\x41\x7e\xfa\x76\x68\x7f\xfd\x3e\x7e\x67\x96\x7a\xaf\xde\xe4\xfc\x61\x5e\x87\x73\xc3\x31\x70\xf5\xc6\x98\xb4\xa8\xcd\x2c\x18\x93\xe4\xc8\xa3\x6a\xdb\xa3\x84\x2a\xab\xed\x5d\xf8\xcb\x12\xfa\x86\x67\x1a\x3d\x66\x8b\x48\x9c\xd9\x84\x7f\x1a\x6e\x89\x86\x2f\x00\xe3\xc5\x89\x51\xdb\x18\xf5\xe0\xb2\xe8\x62\x0f\xc3\x25\x21\xec\xc4\x67\xb1\x5b\xe4\x7b\xef\x82\xe8\x72\xae\x99\x70\x0f\x95\xa7\x68\xc6\x5b\x8c\x71\x57\xe6\x4c\xc3\xbb\x18\x08\x03\xf1\x6a\xcf\x84\x7b\xd5\x26\x34\x18\xc8\x69\xcf\xda\x78\x09\x94\xd9\xd7\x91\xfd\xc4\x21\x66\x47\x31\x8d\x12\xec\x7e\x2d\x27\x7d\xf2\xc8\x3a\xaf\x73\x6e\xf8\x92\x5d\x9d\x87\x7c\x3f\xc5\x64\x30\xa1\x59\x33\xe6\x6b\x8f\x41\xac\x8e\xda\x46\x90\x7d\x3d\x7d\xa5\x7a\xae\x99\xff\x96\x96\xab\x53\x4a\xd2\xe1\x8c\x4c\x4c\xd3\x26\x37\x3d\xb7\x35\x4a\xea\xdb\x70\x21\xa6\xb3\x0e\xee\xbf\x26\x17\x43\x86\x52\xc4\xe1\x5f\xc0\x34\x4d\x90\xee\x5c\x0e\x17\xff\x87\xee\x71\x9a\xb8\x5a\xa5\x04\xad\x24\x90\x79\xed\x5d\x8a\x78\x50\x2f\xd0\x73\xed\xb6\x4e\xc2\x96\xa8\x14\x0c\x66\xc5\x95\xb0\xcc\xef\xb6\x30\x30\xee\x5a\xd7\x59\xb7\x28\xa1\xbe\xb2\x8f\x00\x00\x00\xff\xff\x75\xa1\xdc\xf8\x76\x02\x00\x00")

func web_templates_organization_html() ([]byte, error) {
	return bindata_read(
		_web_templates_organization_html,
		"web/templates/organization.html",
	)
}

var _web_templates_repository_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x92\x31\x6f\xfb\x20\x10\xc5\xf7\x7c\x8a\xfb\xb3\x07\x22\x45\xfa\x0f\x15\xf6\x10\xb7\x6b\x9b\xa1\x4b\x47\x62\xae\x01\x15\x43\x04\x97\x48\x16\xf2\x77\xaf\x6c\x12\x27\xb1\x3a\xf9\xf9\xdd\xf9\xfd\xce\x07\xf2\xdf\xeb\x47\xf3\xf9\xb5\x7f\x03\x43\x9d\xab\x57\x72\x7c\x80\x53\xfe\x58\x31\xf4\xac\x5e\x01\x48\x83\x4a\x8f\x02\x40\x92\x25\x87\x75\xd3\xec\xa5\x28\xb2\xd8\xce\xfa\x1f\x88\xe8\x2a\x96\xa8\x77\x98\x0c\x22\x31\x30\x11\xbf\x2b\x66\x88\x4e\x2f\x42\xf4\x67\xcb\x7b\x65\x42\x50\x27\x9b\x78\x1b\x3a\x71\x3a\x47\x14\x1b\xfe\x9f\x6f\x26\xb9\xee\xac\xe7\x6d\x4a\x85\x29\x6e\x50\x79\x08\xba\xbf\x62\xb4\xbd\x80\xd5\x15\x6b\x83\x27\xf4\xc4\x8a\x3d\x4e\xb8\xad\x9b\x88\x1a\x3d\x59\xe5\xa0\x09\x67\x4f\x09\x76\x3d\xec\xa2\xf2\xad\x91\xc2\x6c\xe7\x56\x52\x07\x87\xd0\x3a\x95\x52\xc5\x26\x6e\x71\xee\x72\x7d\x08\x51\x63\x44\x3d\xe7\x8f\x9f\xdd\x97\x30\x3b\xf5\x2d\x9e\xcc\xb2\xb2\x9c\xe6\xb9\x67\x7c\x9b\xe2\x66\x27\xe7\xa8\xfc\x11\x81\x0f\xc3\x03\x33\x3e\xc7\xea\x3a\x67\xfe\xae\x3a\x1c\x06\x29\x48\xff\x51\xbc\x63\x27\xea\xb2\x4f\x8a\xc7\xc8\x9c\xd1\xeb\x99\x27\xc5\xf4\xf3\xd7\x4d\x0b\x6d\x2f\xe5\x1c\xca\xfa\xa5\x28\xd7\xe3\x37\x00\x00\xff\xff\x04\x0e\x0c\xb5\x2f\x02\x00\x00")

func web_templates_repository_html() ([]byte, error) {
	return bindata_read(
		_web_templates_repository_html,
		"web/templates/repository.html",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"web/templates/index.html":        web_templates_index_html,
	"web/templates/organization.html": web_templates_organization_html,
	"web/templates/repository.html":   web_templates_repository_html,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"web": &_bintree_t{nil, map[string]*_bintree_t{
		"templates": &_bintree_t{nil, map[string]*_bintree_t{
			"index.html":        &_bintree_t{web_templates_index_html, map[string]*_bintree_t{}},
			"organization.html": &_bintree_t{web_templates_organization_html, map[string]*_bintree_t{}},
			"repository.html":   &_bintree_t{web_templates_repository_html, map[string]*_bintree_t{}},
		}},
	}},
}}
