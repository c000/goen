// Code generated by "esc -o bindata.go -pkg generator -private templates/"; DO NOT EDIT.

package generator

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/templates/context.tgo": {
		local:   "templates/context.tgo",
		size:    532,
		modtime: 1529737446,
		compressed: `
H4sIAAAAAAAC/5yQwW7CMBBE7/mKEaIViWg+IBKXNL1yKT8Q20sUKV2ovRFFUf692gCmRGoP9ckezz7t
jJyPhKp8PbDQlyCI761gSAAgaw7EefxMJnEY4GtuCEupTUcoNljmO70GjOPNsnQmkOwUXmxw9C3LHoun
UJXvJIvrbP7G0sr559hMR/bIujuJnb7GJNn3bLGlU9xz5dq6Iyvb+oM0UMvNGs4gC59dXpUpsnveS1Bn
rK45xf2NpIh0cnuS3jOeo+0C0ROlQpnrqP9V2ss11H+am8/O6ivAdJo3uHLGpg+raZc3klb6HQAA//9D
kKRFFAIAAA==
`,
	},

	"/templates/root.tgo": {
		local:   "templates/root.tgo",
		size:    510,
		modtime: 1529737446,
		compressed: `
H4sIAAAAAAAC/4SPzY6bMBSF936KI5TFzKL2fqpZNamK1PxIoQ9gzA22AjayL2kjxLtXhqQ/q9ldzv24
57NS+BIaQkueomZqUN9hmYf0plTr2I61NKFXV907Y10zqjaQ/4ztEYdjhd22rFB9K8/4Wn7fSaEUfiRC
uICtS0hhjIZgcoNLaMONol87NPZlhc4Z8onAVjOM9qgJlzD6Bs7nY2wJF9cRumAWu+A/tkOkISTHId6l
EIM2V90SpgkbeVo/DronzLMQrh9CZLwIAJmI2reEzSN+e8dGlsuc8Gmen9RjLx9nUPyTnTRbzHPxZMk3
y6+vQtx0RE+sz8ZSr/EOTz9fsrDc/0lfhfirwbruaLWo8piy8zSBqR86zYRiISS3oXjSK5Fb8/v+g03w
TL94xWXe/w4AAP//BUC9z/4BAAA=
`,
	},

	"/templates/table.tgo": {
		local:   "templates/table.tgo",
		size:    11531,
		modtime: 1530274031,
		compressed: `
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
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/templates": {
		isDir: true,
		local: "templates",
	},
}
