// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package asset

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 12, 8, 0, 40, 8, 833647412, time.UTC),
		},
		"/etc": &vfsgen۰DirInfo{
			name:    "etc",
			modTime: time.Date(2018, 12, 8, 0, 40, 8, 835372231, time.UTC),
		},
		"/etc/docker": &vfsgen۰DirInfo{
			name:    "docker",
			modTime: time.Date(2018, 12, 8, 0, 40, 8, 833791262, time.UTC),
		},
		"/etc/docker/daemon.json": &vfsgen۰CompressedFileInfo{
			name:             "daemon.json",
			modTime:          time.Date(2018, 12, 8, 0, 40, 8, 834130802, time.UTC),
			uncompressedSize: 239,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\xca\xcc\x2b\x4e\x4d\x2e\x2d\x4a\xd5\x2d\x4a\x4d\xcf\x2c\x2e\x29\xca\x4c\x2d\x56\xb2\x52\x88\xe6\x52\x50\x50\x50\x50\x4a\xc9\x4f\xce\x4e\x2d\x52\xd2\x41\xe6\xe9\x65\xe6\xc3\x04\xa0\x5a\x2a\x75\x0d\xf5\x30\xe4\xd2\x93\x91\x79\xd9\x16\xc5\x7a\xa8\x22\x85\xa5\x89\x95\x48\xdc\xea\x6a\xbd\x80\xa2\xfc\x8a\x4a\xcf\x82\xda\x5a\x2b\x4b\x13\x03\x43\x5c\x12\x46\xb8\x24\x8c\x71\x49\x98\x28\x71\x29\x28\xc4\x72\xd5\x02\x02\x00\x00\xff\xff\xc3\x6d\x8a\xf6\xef\x00\x00\x00"),
		},
		"/etc/nginx": &vfsgen۰DirInfo{
			name:    "nginx",
			modTime: time.Date(2018, 12, 8, 0, 40, 8, 834327449, time.UTC),
		},
		"/etc/nginx/nginx.conf": &vfsgen۰CompressedFileInfo{
			name:             "nginx.conf",
			modTime:          time.Date(2018, 12, 8, 0, 40, 8, 834571974, time.UTC),
			uncompressedSize: 2505,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x92\xcf\x8e\x9b\x30\x10\x87\xef\x7e\x8a\x39\xe4\x1a\x48\xb2\xa8\xca\x2e\xf7\xb6\x7b\x5b\xf5\xd4\x9b\xe5\xc5\x03\xb1\x00\x0f\x9d\x31\xf9\xd3\x88\x77\xaf\x4c\x48\xb5\xd9\x36\x3d\xe5\xd2\x36\x48\x48\x63\x7e\xdf\xcc\x58\xe2\xeb\x05\x19\xc0\x57\xce\xef\x73\xb5\x23\xae\x91\x75\xc7\x54\xa0\x08\x0a\x80\xe9\x03\xe5\x4a\x21\x33\xb1\x6e\xa8\x02\x48\xb7\x86\xd3\x86\xaa\x74\xec\x49\xc7\x24\x89\xc9\xce\xb0\xcf\x55\xe7\x2c\x4c\xcf\x48\x72\xef\x4f\x64\xd2\x39\x1b\x27\x6d\xd1\x07\x81\xa3\x02\x98\xb6\x15\xe4\x3d\x16\xc1\x91\x17\x80\xe5\x62\x95\xe5\x6a\x50\x6a\x13\x42\x37\x52\x82\xbc\x45\x1e\xcb\xf3\x41\x7b\xd3\x22\x30\x56\x4e\x02\x1f\xe6\xcb\xc4\x52\x51\x23\x27\x8e\xf2\x91\x6a\x9c\x04\xf4\xb0\x5e\x4c\x47\x2a\x4c\x1c\x0f\xe9\x34\x05\xa0\x63\xda\x1f\x74\x67\x44\x20\x2e\x7a\x4a\xd3\xe3\x31\x79\x89\x1f\x9f\xbb\x61\x78\x7a\xcc\x16\xcb\xfc\x02\x65\xb4\x8e\xb1\x08\x40\x65\x79\x99\x08\x06\xbd\x41\x63\x91\xe1\x33\x49\x80\xd9\x86\x24\x5c\x45\xbe\xce\xbf\xa0\x69\xe6\xcf\x2f\x30\x63\x6c\x29\xa0\x36\xd6\xf2\x1f\xf0\x8f\xc4\x3b\xc3\x16\x6d\xac\x60\x76\x22\x8c\xb5\x7a\xaf\xcb\x73\x14\xab\xf3\x88\xa2\x71\xe8\x83\x6e\xcd\x5e\xbf\x92\x3d\x68\x71\xdf\x11\x96\x8b\x4f\xef\xf2\x31\x7b\xed\xcb\x12\x79\x42\x56\xeb\xfa\xf2\x1a\xd3\x7f\xd1\xc1\xb5\x48\x7d\x80\xc7\xc5\xfb\x6b\x7a\x7b\x35\x64\x34\xd7\xc3\xb7\x7b\xb3\xfa\x77\x99\x40\x06\x0f\xab\x5f\x22\xf9\x99\x9f\x9a\x3f\xc4\x6e\x75\x01\x05\x6c\x3b\x5d\xba\x06\xf5\x8e\x5d\xc0\x37\x5c\x64\x06\x75\x7a\xaf\x4a\x55\x15\xb7\xd5\x68\x75\xd7\xe8\x2f\xd1\xe8\x86\x12\xd5\x6b\x49\x6e\x2d\xd2\xc3\x5d\xa4\xff\x4f\xa4\x6f\xbd\x39\xdc\xd4\xa2\xec\x6e\xd1\x3f\x6a\xd1\xf0\x23\x00\x00\xff\xff\x44\x5d\xa6\xe7\xc9\x09\x00\x00"),
		},
		"/etc/sysctl.d": &vfsgen۰DirInfo{
			name:    "sysctl.d",
			modTime: time.Date(2018, 12, 8, 0, 40, 8, 834778400, time.UTC),
		},
		"/etc/sysctl.d/k8s.conf": &vfsgen۰CompressedFileInfo{
			name:             "k8s.conf",
			modTime:          time.Date(2018, 12, 8, 0, 40, 8, 835252441, time.UTC),
			uncompressedSize: 78,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\x4b\x2d\xd1\x4b\x2a\xca\x4c\x49\x4f\x85\x52\xba\x79\x69\xba\xc9\x89\x39\x39\xba\x99\x05\x66\x25\x89\x49\x39\xa9\xc5\x0a\xb6\x0a\x86\x5c\xf8\xd4\x21\x94\x01\x02\x00\x00\xff\xff\x7a\x4e\x4e\x73\x4e\x00\x00\x00"),
		},
		"/etc/yum.repos.d": &vfsgen۰DirInfo{
			name:    "yum.repos.d",
			modTime: time.Date(2018, 12, 8, 0, 40, 8, 835449464, time.UTC),
		},
		"/etc/yum.repos.d/bootstrap.repo": &vfsgen۰CompressedFileInfo{
			name:             "bootstrap.repo",
			modTime:          time.Date(2018, 12, 8, 0, 40, 8, 835807052, time.UTC),
			uncompressedSize: 169,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\xae\xae\xd6\x73\xca\xcf\x2f\x29\x2e\x29\x4a\x2c\xf0\x2c\xa8\xad\x8d\x47\x16\x08\xc8\x2f\x2a\xa9\xad\x8d\xe5\xca\x4b\xcc\x4d\xb5\x4d\x4c\x49\x49\x4d\x51\x48\x2b\xca\xcf\xb5\x52\xc8\x28\x29\x29\xb0\xd2\xd7\x47\xd7\x6c\x85\xa9\x99\x2b\x29\xb1\x38\xb5\xb4\x28\xc7\x96\x04\x2d\xa9\x79\x89\x49\x39\xa9\x29\xb6\x86\x5c\xe9\x05\xe9\xc9\x19\xa9\xc9\xd9\xb6\x06\x80\x00\x00\x00\xff\xff\x5b\x55\x87\x86\xa9\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc"].(os.FileInfo),
	}
	fs["/etc"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/docker"].(os.FileInfo),
		fs["/etc/nginx"].(os.FileInfo),
		fs["/etc/sysctl.d"].(os.FileInfo),
		fs["/etc/yum.repos.d"].(os.FileInfo),
	}
	fs["/etc/docker"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/docker/daemon.json"].(os.FileInfo),
	}
	fs["/etc/nginx"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/nginx/nginx.conf"].(os.FileInfo),
	}
	fs["/etc/sysctl.d"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/sysctl.d/k8s.conf"].(os.FileInfo),
	}
	fs["/etc/yum.repos.d"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/etc/yum.repos.d/bootstrap.repo"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
