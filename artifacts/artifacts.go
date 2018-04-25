// Code generated by go-bindata.
// sources:
// artifacts/backup-pvc.yaml
// artifacts/backupschedule-crd.yaml
// artifacts/backupschedule-cronjob.yaml
// artifacts/cluster-configmap.yaml
// artifacts/cluster-crd.yaml
// artifacts/cluster-service-read.yaml
// artifacts/cluster-service.yaml
// artifacts/cluster-statefulset.yaml
// DO NOT EDIT!

package artifacts

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

var _artifactsBackupPvcYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8d\x31\x8a\xc3\x30\x10\x45\x7b\x9d\xe2\x5f\x60\x0d\xdb\xaa\x75\xed\xc5\xac\xc1\xa9\x07\xe9\x13\x44\x2c\xc9\xd1\xc8\x81\x60\x7c\xf7\xa0\x38\x69\x06\xde\xfc\x07\xef\x16\x92\xb7\x18\x59\x34\x68\x65\xaa\x73\x5e\xb6\xc8\x7e\x91\x10\x8d\xac\x61\x6e\x43\x4e\x16\x8f\x5f\x13\x59\xc5\x4b\x15\x6b\x80\x24\x91\x16\xfb\x8e\x71\xee\xff\x24\x12\xdd\xfb\x1e\x87\xd1\x95\xae\x19\xe2\x1c\x55\x87\xec\xa9\x0d\x81\x1f\xfc\x53\xfc\xa5\x84\xca\x41\xd2\xd3\x00\x85\x9a\xb7\xe2\xbe\x42\xe1\x7d\xa3\xd6\x0f\x01\x5a\x73\x91\xeb\x99\xe9\xa6\x95\xae\x9b\xce\x4f\xcb\xbc\x02\x00\x00\xff\xff\x31\xd3\xda\x23\xb8\x00\x00\x00")

func artifactsBackupPvcYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsBackupPvcYaml,
		"artifacts/backup-pvc.yaml",
	)
}

func artifactsBackupPvcYaml() (*asset, error) {
	bytes, err := artifactsBackupPvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/backup-pvc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsBackupscheduleCrdYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xbb\x4e\xec\x30\x10\x86\x7b\x3f\xc5\xbc\xc0\xf1\x51\x3a\xe4\x12\xe8\x58\x90\x60\x25\xfa\x89\x33\x24\xa3\xf8\x86\x67\x1c\xb1\x6f\x8f\xe2\x80\x68\x16\x51\xda\xff\xe7\xff\x62\x2c\xfc\x4a\x55\x38\x27\x07\x58\x98\x3e\x94\xd2\x7e\x12\xbb\xde\x88\xe5\xfc\x7f\x1b\x46\x52\x1c\xcc\xca\x69\x72\x70\xd7\x44\x73\x7c\x21\xc9\xad\x7a\xba\xa7\x37\x4e\xac\x9c\x93\x89\xa4\x38\xa1\xa2\x33\x00\x09\x23\x39\x88\x17\x79\x0f\x23\xfa\xb5\x15\xf1\x0b\x4d\x2d\x90\x58\x5f\x6d\xbf\xcf\x85\x2a\x6a\xae\x76\xae\x1a\xec\xcc\xba\xb4\xd1\xfa\x1c\x8d\x14\xf2\xbb\xc7\x5c\x73\x2b\x0e\xfe\xe4\x8f\x34\xd9\x9f\x00\x1c\x1d\x1f\x2f\xe7\xe7\xd3\x6d\x0f\x3e\x7f\x05\x77\x35\xb0\xe8\xc3\x6f\xc4\x89\x45\x3b\x25\x9c\xe6\x16\xb0\x5e\x1d\xd0\x89\x12\x5a\xc5\x70\x7d\xe0\x61\xb1\xe4\xaa\x4f\x3f\xb5\xfe\x41\x1c\x77\x45\x7c\x2e\xe4\xa0\x2b\x05\x3d\x4d\x06\x60\xfb\xfe\xfc\x6d\x30\x9f\x01\x00\x00\xff\xff\xa4\x65\x5f\x41\x8c\x01\x00\x00")

func artifactsBackupscheduleCrdYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsBackupscheduleCrdYaml,
		"artifacts/backupschedule-crd.yaml",
	)
}

func artifactsBackupscheduleCrdYaml() (*asset, error) {
	bytes, err := artifactsBackupscheduleCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/backupschedule-crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsBackupscheduleCronjobYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xc1\x6e\xd3\x40\x10\xbd\xfb\x2b\x46\x81\xa8\x07\xb4\x76\xa2\x1e\x90\x2c\xf5\x00\x29\x08\x10\x29\x81\x96\x4a\xdc\x3a\x5e\x4f\x93\xa5\xde\x5d\x77\x77\x1c\x12\xa5\xf9\x77\xb4\x76\x52\xe2\x6d\x23\x95\x39\xf8\x30\xef\xcd\xf8\xcd\x9b\x1d\xac\xd5\x35\x39\xaf\xac\xc9\xa1\x40\x96\x8b\x6c\x39\x2e\x88\x71\x9c\xdc\x29\x53\xe6\x30\x71\xd6\x7c\xb1\x45\xa2\x89\xb1\x44\xc6\x3c\x01\x30\xa8\x29\x87\xcd\x66\x0f\x5e\xa0\x26\x48\xdb\xef\x76\x9b\x00\xd8\x3f\x86\xdc\x0f\xba\x25\x47\x46\x92\x0f\x25\x00\x02\xba\x86\xd3\xf5\xe5\xf7\xaf\xef\x51\xde\x35\xf5\xa5\x5c\x50\xd9\x54\xd4\xe2\x00\x87\x52\xa4\x4b\xf5\xda\xdf\x57\xb6\x26\x87\x6c\x5d\x3a\x77\x5c\xa5\x73\xc5\x8b\xa6\x48\xa5\xd5\xd9\x72\xbc\xab\x7a\x14\x73\x20\x20\x44\xa3\xca\x2e\xfd\xf3\xf3\x79\xc8\xfa\x9a\x64\x50\xe2\x77\x3f\xcd\x61\x10\xd0\xcb\x9a\x64\x7a\xa5\xda\xca\x41\x02\xf0\xdb\x16\x57\xa4\xeb\x0a\x99\x3a\xdd\xfb\xba\x10\xdc\x43\x62\x34\x84\x23\xcf\xe8\x78\x66\x2b\x25\xd7\x39\x7c\x33\x1f\x51\x55\x8d\xa3\x03\xca\xd2\x56\x8d\xde\xbb\xd2\x85\xd8\x4d\x51\xb4\xb6\x1c\x00\x00\x75\x70\xc4\x33\x19\xbe\x6e\xeb\x26\x15\x2a\x9d\xf7\x28\x00\x32\x24\x2f\xf6\x46\xcc\xae\x27\xf1\x46\xe2\x1f\x49\x6b\x6e\x7b\x3d\x48\xd7\xbc\x3e\x57\x2e\x87\xcd\x31\xbe\x9a\x0b\x8d\x7d\x71\x5d\x7a\x8a\x75\x2c\xa8\xab\x6a\x37\x98\xf4\xe8\x8c\xca\x90\x7b\xe9\xf0\x4a\xe3\x9c\x72\x08\xbb\xcf\x56\xec\xb0\xa3\xe4\x61\x05\x9e\x23\x25\x5a\xa3\x29\xfb\x3a\x04\x14\xe8\x17\x51\x6a\x20\xe4\x20\x4a\x3d\x44\xea\x3d\x31\x08\x5a\x25\x51\xfa\x15\x4c\x2a\x6b\x08\xc2\x19\xc0\xad\xb3\x1a\x34\x7a\x26\x97\x46\xbc\xf3\x77\x57\x1f\xce\x6e\x4a\x64\x82\x93\x37\xc3\x5f\x62\xa8\xc5\xb0\x14\xc3\x4f\x62\x38\x3d\xb9\x89\xb8\xfa\xae\x54\x0e\xb2\xd6\xa9\xac\x1b\x2f\x3b\x78\xcb\xd9\xeb\x4d\xe8\xb6\x8d\xed\x95\xc8\x20\x84\x23\xb9\x14\xd6\x54\x6b\x78\x7c\xc9\x93\xaa\x09\x92\x60\xbb\x15\xa3\xf4\x99\x2c\x9c\x9e\x8e\xde\xc2\x03\xac\x0a\xcf\x8e\x50\x83\x58\x81\x98\xbc\x40\xc0\x13\x2f\x66\x8e\x6a\x74\x04\xbc\xa0\xdd\xea\x62\x1f\xfe\x6d\x0c\x84\xa8\x77\x74\x21\x18\xdd\x9c\x58\x94\xca\x9d\xfd\xd7\xd8\xdd\xdd\x4c\x6d\x63\xd8\xc7\x7b\x3e\xfa\x82\x00\x74\x28\x98\x21\x2f\xf2\xe3\x53\x3e\xdb\xed\xc9\x8d\xf4\x7b\x11\xcb\x5d\xbf\x40\x4c\xcb\x88\xea\xc8\xdb\xc6\x49\xf2\xf1\x69\x04\xe8\xbe\x21\x1f\x0f\xd1\x85\xac\x9b\x1c\xc6\xa3\x91\x7e\x06\xd3\xa4\xad\x5b\xb7\xf0\x54\xfd\x0d\x00\x00\xff\xff\x38\xda\xd5\x94\xb8\x05\x00\x00")

func artifactsBackupscheduleCronjobYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsBackupscheduleCronjobYaml,
		"artifacts/backupschedule-cronjob.yaml",
	)
}

func artifactsBackupscheduleCronjobYaml() (*asset, error) {
	bytes, err := artifactsBackupscheduleCronjobYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/backupschedule-cronjob.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsClusterConfigmapYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\xb1\x0a\xc2\x30\x10\xc6\xf1\x3d\x4f\xf1\x81\x73\x05\xd7\x6c\xe2\xec\xea\x22\x0e\xd7\xe6\xda\x06\x93\x4b\xcc\xc5\x42\xc1\x87\x97\x06\xba\x39\x38\x26\xf7\xbb\xff\x51\xf6\x37\x2e\xea\x93\x58\x2c\x27\xf3\xf4\xe2\x2c\x2e\x49\x46\x3f\x5d\x29\x9b\xc8\x95\x1c\x55\xb2\x06\x10\x8a\x6c\x11\x57\x7d\x05\x03\x04\xea\x39\xe8\xf6\x0f\x50\xce\xfb\x60\xc7\x91\xb4\x72\x39\x0e\x32\x5a\x7c\x1a\x3a\xe0\x9c\x73\x58\x51\x67\xaf\x18\xda\x05\x24\x09\x2b\x92\xa0\xce\xbc\x6f\x34\x7b\x6f\x31\xf7\x68\x8f\x90\xa6\xae\xf7\x62\x00\x0d\xb4\xf0\xbf\xcd\x86\xf5\x47\x4f\xdf\x99\x4b\x57\x98\x5c\xb7\x59\xf3\x0d\x00\x00\xff\xff\x1c\x35\x5e\xba\x03\x01\x00\x00")

func artifactsClusterConfigmapYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsClusterConfigmapYaml,
		"artifacts/cluster-configmap.yaml",
	)
}

func artifactsClusterConfigmapYaml() (*asset, error) {
	bytes, err := artifactsClusterConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/cluster-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsClusterCrdYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x50\xbd\x4e\x03\x31\x0c\xde\xf3\x14\x7e\x01\x82\xba\xa1\xac\x65\xa3\x20\x01\x12\xbb\x9b\x9a\xab\xd5\x24\x0e\xb6\x53\xd1\xb7\x47\x77\x27\x40\xc0\xc0\x66\x7f\x3f\xfa\x3e\x1b\x3b\xbf\x90\x1a\x4b\x4b\x80\x9d\xe9\xdd\xa9\xcd\x9b\xc5\xd3\x8d\x45\x96\xeb\xf3\x66\x4f\x8e\x9b\x70\xe2\x76\x48\xb0\x1d\xe6\x52\x9f\xc8\x64\x68\xa6\x5b\x7a\xe5\xc6\xce\xd2\x42\x25\xc7\x03\x3a\xa6\x00\xd0\xb0\x52\x82\x7a\xb1\xb7\x92\xcb\x30\x27\xb5\x98\x35\x2e\x80\x74\x52\x74\xd1\x38\xa9\x97\x38\xb1\x1f\xc7\x3e\x66\xa9\xc1\x3a\xe5\xd9\x3c\xa9\x8c\x9e\xe0\x5f\xfd\x1a\x63\xb3\x05\x60\x2d\x77\x7f\x79\x7e\xdc\x6d\xd7\xc4\x05\x2e\x6c\x7e\xf7\x87\xda\xb1\xf9\x42\x1b\xb7\x69\x14\xd4\x9f\x5d\x17\xaa\x97\xa1\x58\x7e\x1d\xb1\x9a\x8e\xa2\xfe\xf0\x1d\x7d\x05\x35\x7f\x0d\xb3\xc4\xb2\x74\x4a\xb0\x48\x3a\x66\x3a\x04\x80\xf3\xe7\x8b\xcf\x9b\xf0\x11\x00\x00\xff\xff\xdd\x99\x29\xdc\x72\x01\x00\x00")

func artifactsClusterCrdYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsClusterCrdYaml,
		"artifacts/cluster-crd.yaml",
	)
}

func artifactsClusterCrdYaml() (*asset, error) {
	bytes, err := artifactsClusterCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/cluster-crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsClusterServiceReadYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xc1\x4e\xc3\x30\x10\x44\xef\xf9\x8a\xfd\x81\x1a\xaa\x4a\x1c\x72\x85\x0b\x12\xa0\xd2\x0a\xee\x5b\x67\x08\x16\x8e\x6d\xd6\x9b\xa0\xaa\xea\xbf\x23\x3b\x01\x7a\xc1\x07\x1f\x76\xde\xec\xcc\x72\x72\xaf\x90\xec\x62\x68\x69\x5a\x37\x1f\x2e\x74\x2d\xed\x21\x93\xb3\x68\x06\x28\x77\xac\xdc\x36\x44\x81\x07\xb4\x74\x3a\xd1\x0e\xdc\x2d\xc0\x13\x0f\x20\x53\xff\xf3\xb9\x21\xf2\x7c\x80\xcf\x85\x26\xe2\x94\x2a\x7e\x21\xc7\xaf\x00\xd9\xe1\x0d\x82\x60\xb1\x70\x2b\x9a\x33\x1f\x8f\xfb\xe7\x87\x5b\x3f\x66\x85\x54\xa1\xac\xf8\xeb\x66\xc5\x0c\xc7\xfc\xe9\x63\x82\xb0\x46\x31\xbd\xa8\x37\xbd\xd3\xf7\xf1\x60\x6c\x1c\xae\xa6\xf5\xe2\xfa\x2d\x7a\x91\x5c\xde\xe8\xba\x79\xfc\x72\x7f\x57\xa6\x39\xc1\x96\x0a\x29\x8a\xd6\x2e\xab\xc5\x5a\x73\xaa\xa9\x48\xb3\x67\x9f\x60\xcd\x36\x8a\xfe\xec\x53\x96\x1e\xba\xad\xc0\x66\x73\x7d\xd3\x10\x65\x78\x58\x8d\xf2\xcf\xfd\xdf\x01\x00\x00\xff\xff\xd5\x59\xa5\xd3\x6a\x01\x00\x00")

func artifactsClusterServiceReadYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsClusterServiceReadYaml,
		"artifacts/cluster-service-read.yaml",
	)
}

func artifactsClusterServiceReadYaml() (*asset, error) {
	bytes, err := artifactsClusterServiceReadYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/cluster-service-read.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsClusterServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\x4e\xc3\x30\x10\x44\xef\xfe\x8a\xfd\x81\x1a\xaa\x4a\x1c\x72\x85\x4b\x25\xa8\x0a\x11\xdc\x5d\x67\x08\x16\x8e\x6d\xd6\x9b\xa0\xaa\xea\xbf\x23\x3b\x29\xf4\x82\x0f\x3e\xec\xcc\xec\x3c\xad\x49\xee\x0d\x9c\x5d\x0c\x0d\x4d\x6b\xf5\xe9\x42\xd7\x50\x0b\x9e\x9c\x85\x1a\x20\xa6\x33\x62\x1a\x45\x14\xcc\x80\x86\x4e\xa7\x8b\xb8\x33\x03\x48\xd7\xff\x7c\x56\x44\xde\x1c\xe0\x73\x71\x12\x99\x94\xaa\xf5\x4a\x8e\xdf\x01\xfc\x82\x77\x30\x82\xc5\xe2\x5b\xd1\xdc\xf7\x74\x6c\x9f\x1f\xef\xfd\x98\x05\x5c\x85\xb2\xe2\x8f\xcb\xb2\x1e\x8e\xf9\xcb\xc7\x04\x36\x12\x59\xf7\x2c\x5e\xf7\x4e\x3e\xc6\x83\xb6\x71\xb8\x99\xd6\x4b\xea\x17\xf2\xaa\xb9\xbc\xd1\x75\xf3\xf8\x75\xfb\x50\xa6\x39\xc1\x16\x84\x14\x59\x2a\xcb\x6a\x89\xd6\x9e\x1a\x2a\xd2\x9c\x69\x13\xac\xde\x47\x96\xcb\x3e\x31\xdc\x43\xf6\xd5\xb0\xd9\xdc\xde\x29\x22\x3b\xc3\x6f\xf7\x0d\xed\x62\x80\x22\xca\xf0\xb0\x12\xf9\x9f\x8b\xfc\x04\x00\x00\xff\xff\x97\xdb\xcc\x3a\x78\x01\x00\x00")

func artifactsClusterServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsClusterServiceYaml,
		"artifacts/cluster-service.yaml",
	)
}

func artifactsClusterServiceYaml() (*asset, error) {
	bytes, err := artifactsClusterServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/cluster-service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _artifactsClusterStatefulsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\xe1\x72\x1a\x39\x12\xfe\xef\xa7\xe8\x22\xae\x35\xbe\x58\x18\x92\xca\xa5\x8e\x84\x5c\x79\x09\x49\x5c\x67\x1b\x2f\xc3\x26\xb7\xe5\x70\xac\xd0\x34\xa0\xf2\x8c\x34\x91\x34\xd8\x9c\xe3\x7b\xf6\x2b\x8d\x06\xd0\x0c\x83\xe3\xa4\x6a\xaf\x6a\x8f\x1f\x36\x23\xa9\x5b\x5f\x7f\xea\x6e\x75\x0f\x34\xe1\x1f\x51\x69\x2e\x45\x1b\x68\x92\xe8\xe3\x45\x6b\xef\x9a\x8b\xb0\x0d\x81\xa1\x06\xa7\x69\x14\xa0\xd9\x8b\xd1\xd0\x90\x1a\xda\xde\x03\x10\x34\xc6\x36\xdc\xdd\xf9\x0b\x2e\x68\x8c\xd0\xc8\xfe\xde\xdf\xef\x01\xc8\x1b\x81\x6a\x80\x53\x54\x28\x18\x6a\x2b\x06\x40\xc0\x29\x3e\x5f\x06\xbf\x9c\x75\xa3\x54\x1b\x54\xd9\x04\x80\x8f\x82\xa9\x46\xbc\xd4\x5f\x22\x99\xa0\xa2\x46\xaa\xc6\x4c\x99\xa8\x31\xe3\x66\x9e\x4e\x1a\x4c\xc6\x16\xa1\x93\x5a\x23\xf1\x76\xb6\x9f\x94\x87\x6e\xf8\xd7\xd3\xb7\x76\x54\x27\xc8\x2c\x04\x8d\x11\x32\x23\x95\x83\x13\x53\xc3\xe6\x67\x74\x82\x51\x8e\xcf\xc2\x48\xca\xfa\x34\xaa\x05\x67\x78\x51\xb1\x95\xc2\x24\xe2\x8c\x6a\x37\x1e\x24\xc8\x1a\x83\x7c\xc8\x2d\x30\x18\x27\x11\x35\x98\xef\xe7\x71\x68\x3f\x51\x61\xeb\xca\xcd\x01\x56\xd0\xed\x87\x0b\x6e\xba\x52\x18\xca\x05\xaa\xb5\x20\xc9\x69\xb0\xb3\x24\x23\x6e\xad\x91\xc7\x74\x86\x1e\xba\x53\xfb\xbc\xa1\x09\x80\xc9\x38\xa6\x22\xdc\x60\x20\x30\xa1\x7a\xee\x3d\xd6\x08\xab\x79\x8f\x5f\xd7\xdf\x2d\x33\x06\x08\xde\x7a\x23\x4f\xe0\x3d\x0a\x7b\x68\x08\x19\x90\x8c\x3c\x54\x84\x87\x30\x55\x32\x86\x44\x86\x20\x55\xc8\x05\x8d\x80\x8b\x10\x6f\x1b\x9e\xf0\xd5\x15\xfc\x3e\x97\xda\x58\x6b\x7e\x87\xce\x7f\x80\xd4\xaf\x9a\xe4\x6f\xa3\xa7\x87\xfb\x30\x1a\xc1\xd7\xaf\x80\xb7\xdc\x40\xcb\x13\xc9\x75\x75\xf6\xef\x7e\x3e\x09\x3e\x8c\x07\xbd\xf3\x93\x61\xf7\xc3\x55\x6b\x74\xef\x2d\x42\x36\x97\x70\x95\xe1\x09\x47\xf0\x06\x8e\x63\x61\x8e\x99\x14\xd3\x46\x78\xbc\xc6\xd7\x60\x62\x5a\x30\xe4\x24\x0c\x81\x0a\x90\xd3\xa9\x35\xd3\x48\xa0\x0b\xc9\x43\x50\x98\x89\x84\x1b\xcb\x3a\x4d\x58\xd0\x28\xc5\x46\x79\xcb\xcd\x8a\xfd\x7a\xbd\xd5\x6c\xc2\x53\xd8\xcf\x01\x1f\x1e\xc2\x9b\xc7\x02\xe9\xca\x64\x69\x5d\x43\xc9\x44\x71\xcb\xac\x93\x80\x29\x8f\x50\x3b\x5a\xed\x08\x9f\x91\x98\x26\x16\x28\xc6\x89\x59\xbe\xe5\xca\x07\xc4\xa7\x96\xde\xd5\xf6\x40\xf0\x0b\x34\x61\x34\x7a\x05\x66\x8e\xc2\x5b\x07\xc0\x92\x0d\x2e\xa7\xf3\x38\xa6\x36\x50\x2d\xb0\x02\x64\xdf\xde\x48\xe3\xb7\xb4\xe8\x88\x2e\xf0\x21\x25\x53\xbe\x7e\x58\xc8\x28\x8d\xf1\x5c\xa6\xc2\x68\xdf\x37\x9d\xa3\x5b\x51\x4f\x2e\xb6\xcb\x2e\xa9\x99\xb7\x7d\xcd\x95\x52\x0e\xca\x37\x64\x0b\x8b\xd6\xd2\x91\x14\x58\x1d\x5d\x36\x3b\x1d\xdf\x1a\x45\x27\x94\x5d\xa7\x49\xdb\x06\xbc\x36\x7f\x5c\x88\x05\xd7\x3c\xb1\xe7\xe6\x30\xd9\xa3\xb5\x49\x05\x68\xa4\x90\x86\x4b\x1b\x24\xda\xe8\x52\x60\x91\x10\x8e\x17\x54\x1d\x47\x7c\x72\x9c\x19\xe1\xfe\xda\xb8\xfa\xe9\x27\x17\x57\xcd\xff\x41\x28\x3e\x81\x01\x6a\x23\x15\x3a\xc8\x99\xf3\x3a\xd6\x40\x0a\xc8\xfd\xec\x87\xdc\xf6\xee\xce\x2e\x16\x98\x27\xba\x77\x4a\xc6\x3f\x67\x8a\x1b\xee\x5f\x96\x52\x6b\x35\x3f\xf7\xfd\x08\x20\x80\xcd\x41\x03\x21\x89\xc2\x84\x2a\x04\x42\x0c\x55\x33\x34\x24\xe4\xaa\x93\x13\xec\x16\x1d\xef\x16\x66\x32\x59\x12\xfb\xf4\x68\xf1\xbb\x3b\x40\x11\x6e\xdb\x50\xf2\x88\x35\x72\xa8\x17\x92\x2d\x34\x0f\x8b\xb6\x6c\x9d\xbb\x17\x82\x59\xea\xc9\xd4\x6d\x98\x49\x14\x2e\xb8\x4c\x35\x24\x58\xa4\x45\x30\x6a\x80\x10\x85\x6c\x41\xa4\x88\x96\xfe\x2d\x46\xf6\xeb\xf5\xd5\x09\x92\xd6\xe1\x61\xc3\x9b\x83\xe7\xcf\x9b\x2f\xe1\x2b\xdc\x4e\xb4\x51\x48\x63\x20\xb7\x40\xba\x25\x57\x2d\x40\xba\xcc\x19\xb7\xd6\x3a\x8a\x7c\x1c\x8f\x38\x9c\x6a\xd5\x0f\x27\x1c\xcb\xc0\x8e\xa4\xb1\x0b\xa9\x4e\x27\x6e\x45\x71\xe2\x11\x29\x0c\x0d\xcb\x5d\xa0\x94\xc8\x7e\xc0\xc5\x57\xdb\x29\xe7\xe4\xbb\x12\x9f\xe7\x70\x55\x26\xac\xcb\x07\x6f\xcb\x53\xa1\x0d\x15\xac\x50\x4c\x94\xdd\x93\xed\x2c\x55\xbe\xb3\x4a\x41\xb1\xd8\x3e\x94\xf3\xdf\x82\x5f\xce\xc6\x83\x7e\x7f\x38\xbe\x3c\x09\x82\x4f\xfd\xc1\x5b\x0f\x7c\x76\x25\x5b\xbc\xed\x82\xc7\x6b\x64\x0a\xcd\x3f\x70\x39\xc0\x69\x71\xc6\x2f\x26\x33\x14\x41\xb6\xb4\x1c\x6c\x00\xd7\xb8\x6c\x43\x42\xb5\xbe\x91\x6a\x73\x38\x89\x54\x55\x9e\x53\x76\x8b\x35\x25\x97\x52\x99\xb6\x75\xff\xbf\xfe\xa9\x7c\x50\xa1\x96\xa9\x5a\x17\xf4\xab\xc1\x2f\x29\x6a\xa3\xdb\xa5\x12\x20\x6d\xc3\x8b\x66\x33\x2e\x8c\xc6\x18\x4b\xb5\x6c\x43\xeb\xfd\x26\xd5\x44\x7c\x81\x02\xb5\xbe\x54\x72\x82\xbe\x12\xbc\xdd\xd4\xbe\x2b\xfe\xdc\x65\x0a\x57\x35\x7b\x89\xd6\x8e\xb2\xcb\xf3\x08\x6a\x19\x5a\x1a\xc6\x5c\x40\xc2\xc5\x0c\x48\xf2\xb9\xb6\x7f\x57\xe1\x23\xf7\x9f\x6b\xb5\x91\x7f\xc5\x08\x6e\x38\x8d\xde\x62\x44\x97\x01\x32\x29\x42\xdd\x86\xe7\x7e\x52\x4c\x50\x71\x19\xae\xe7\x5a\xfe\x9c\xe1\x31\xca\xd4\xac\x27\x5f\x78\x4c\xd1\x90\x3f\xd2\xaa\x27\xd0\x9d\x23\xbb\x86\x1b\x04\x46\x45\xb6\x20\x35\x08\x5f\x52\x54\x1c\x35\xc8\x05\x2a\x18\x76\x2f\xa1\xae\xaf\x79\x42\x04\x9a\x1b\xa9\xae\xad\x95\x5c\xdb\x02\xb5\x94\xd5\x1f\xe6\x08\xc8\x1c\x5a\xcf\x5e\x36\x9a\x8d\x66\xa3\xf5\x20\x4d\x40\x10\x0e\x82\xde\x59\xaf\x3b\x84\xd6\xc1\xb7\x49\x7b\xb1\x9b\xb3\x67\x0f\x50\xd6\x2a\xe5\x86\x4d\x0e\xff\xce\x42\xeb\xcf\x9b\x25\x2a\x4c\xae\x4a\x15\x2f\xff\xa8\x9a\x92\x95\xab\xc3\xbd\xc2\x9d\xfb\x16\x0d\xaa\x98\x0b\x84\x09\x17\x91\x9c\x41\x22\x35\x37\x5c\xda\xee\xc8\x95\x1c\x61\x96\x9e\x8e\xec\xfd\x44\xc5\x72\xbb\x82\x23\x53\xcf\xc6\x71\xd6\x09\x8c\xb9\x98\xca\xea\x5a\xee\x09\xfc\xd3\x28\xea\x2e\x9a\x75\x59\x3b\xcb\x1b\xcb\x10\x28\x24\x54\x59\xf7\x83\x5a\xf7\xc3\xc9\xc5\xfb\x1e\x9c\x9f\x04\xc3\xde\x00\x86\xfd\x5a\x16\x33\xcb\x92\xb6\x09\x32\x9a\x6a\x84\x1b\x3c\x50\xae\x46\xb2\xb1\x93\x55\x34\x59\xb4\x71\x6d\xec\x80\x6b\x50\x8a\xd9\x6a\xb1\x03\x37\x9b\x53\x31\xc3\xb1\x2b\xb3\xc6\x46\x36\xf4\x97\xa8\xc1\xcb\x66\x9c\xce\x84\x2d\x2d\x3d\x15\x8e\x40\xa7\x83\x0b\x30\x73\xae\x81\x51\x8d\x50\xe7\xe6\x40\x43\xaa\x31\x42\xad\x4b\x01\xad\xe2\x12\x81\x9e\x96\x42\x13\x56\x45\xb6\xbf\xe3\x0e\xb6\x3f\x15\x78\x09\xb9\x42\x66\xa2\xa5\x23\x28\x2f\x81\xe1\x92\x2a\xbd\x75\xfe\x45\x98\xb6\x69\xb0\x75\x60\xf5\xe6\x59\x1b\xf1\xaf\x7a\xe3\x2f\x7f\x3f\xbc\xba\x6a\xeb\x84\x32\x6c\x8f\x46\x4f\xb3\x81\x9d\x4d\x45\x66\xfb\x37\x0d\xcf\xfb\xed\x2d\x6f\xc8\xbf\x8d\xcf\xfa\xef\xc7\xef\x4e\xcf\x7a\x9d\x83\xed\xe6\xe4\xe0\xe8\x73\x29\x72\xed\xc7\x13\xbc\xec\x07\xe5\x9e\xe6\xd9\xe8\xbe\x06\x6f\x1e\xe1\x03\x53\x5e\x0c\x24\x97\xe5\xf9\xd4\x26\x7a\x81\x18\xda\x6e\x9d\xc9\x38\x89\xd0\x20\xd0\xbc\x7a\x9f\x2c\x41\x1b\xeb\xe0\x62\xb6\x7a\xc1\x54\xe6\x7a\x7d\xce\x3b\x20\x54\x1f\xb4\x23\xe9\x13\xe5\x99\xea\xa9\x54\xae\x2e\xc8\x50\x4c\x10\x5c\x98\xd5\x29\x63\x98\x64\x2b\x98\x14\x02\x99\xdd\x5b\x1f\xd6\x0a\x9a\x52\x61\x78\x04\x95\xd7\xc9\x8e\xdb\x24\xbb\x4c\x6a\xab\xcb\xa4\xf6\x0a\x42\x09\x3a\x42\x4c\xa0\x65\xbf\x0b\xdc\xab\x80\x7a\xea\x6e\x19\xfe\xef\x12\x15\xf9\x4b\x8f\x8c\xad\x95\x2b\xd6\xca\xa1\x27\x5c\x5c\xd9\x0c\xb5\x4a\xa4\x59\x19\x4c\x95\x39\x02\x6a\x0c\xc6\x89\x71\xf1\x47\x0d\x89\xa5\x36\x44\x0a\xb6\x1d\xfe\xbb\x18\xae\x1c\x97\x8a\xcf\x8a\x0a\xbe\x8b\xa2\xd7\xaf\x7b\xfd\x77\x9e\xfc\x7e\xfd\xf5\xce\x6d\x0e\x8f\xf6\x2a\x1c\xf6\x43\x3f\x18\x76\x0e\xfc\xe6\xab\xe9\xb7\x5b\x07\x95\x42\xbf\x06\xbd\x41\xe7\x40\x49\x69\xaa\xe7\x57\x08\x6d\xf8\x54\x01\xaf\x96\xea\xf6\x2f\x2e\x7a\xdd\xe1\x78\xd0\x1b\x0e\x7e\xeb\xb4\x9a\xaf\xbc\x45\xc1\xf0\x64\x30\x84\xe0\xec\xe4\x63\xcf\x1f\x2e\x5a\x5f\x0e\x9e\xc0\x1e\x1d\xd0\xfc\xb5\x9a\x75\x5a\x6d\xbb\x0d\x97\x19\x34\xdc\xcc\x51\xac\x2a\x51\x0c\x6d\x14\xd9\x26\xb5\xf0\x1e\xc4\xd6\x55\xab\x56\x35\xe2\xda\xa0\x00\x42\xae\x11\x13\x22\x93\xec\xbb\x55\xe8\xfa\x57\x42\x62\x7a\x4b\x6c\x04\xe8\x4e\xcb\x75\xa9\x84\x41\x31\x55\xd4\x0a\xdd\xe6\xfa\x4b\x76\x4b\x90\x2c\xdf\x12\xe2\xda\xda\xce\xa6\xbf\x25\x73\xa9\x4d\xc7\x73\x07\x92\x6a\x54\x1d\xcb\xbe\x6d\x59\xf3\x82\xa1\xf3\x40\xf1\xfa\x7f\xdd\x2f\xb4\x76\xf6\x0b\xcd\xe6\xf9\xaa\x63\x70\x86\x6f\x35\x95\x05\x7c\xab\x97\xa0\x6d\xb8\xbb\xaf\x58\x57\x7a\x11\xe8\x86\xce\x69\xe2\x43\xda\xee\xe0\xbe\xbb\x03\xdf\xd5\x7f\x27\xa8\x74\xe6\x80\xe6\x63\x66\x4b\x37\xa2\xbc\x50\x83\x32\x3b\x70\x51\x2c\x35\xab\xb7\x5b\xef\xe5\x77\xdf\x8b\x8d\xd6\x61\xfe\x13\x44\x46\x17\xd9\xfa\x0d\x62\xcb\x59\xbe\xfb\x47\x09\x7b\x61\x68\x7d\x2e\x43\xd4\xb6\xe5\x18\x20\x0d\x3f\x29\x6e\xb0\x2f\x18\xae\x9b\x86\x0a\x3f\xa8\xf2\x02\xcb\x52\xf1\x5d\x40\xe0\x46\xec\xc6\xff\x0d\x00\x00\xff\xff\xac\x14\x1d\x9b\xa9\x1a\x00\x00")

func artifactsClusterStatefulsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_artifactsClusterStatefulsetYaml,
		"artifacts/cluster-statefulset.yaml",
	)
}

func artifactsClusterStatefulsetYaml() (*asset, error) {
	bytes, err := artifactsClusterStatefulsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "artifacts/cluster-statefulset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"artifacts/backup-pvc.yaml":             artifactsBackupPvcYaml,
	"artifacts/backupschedule-crd.yaml":     artifactsBackupscheduleCrdYaml,
	"artifacts/backupschedule-cronjob.yaml": artifactsBackupscheduleCronjobYaml,
	"artifacts/cluster-configmap.yaml":      artifactsClusterConfigmapYaml,
	"artifacts/cluster-crd.yaml":            artifactsClusterCrdYaml,
	"artifacts/cluster-service-read.yaml":   artifactsClusterServiceReadYaml,
	"artifacts/cluster-service.yaml":        artifactsClusterServiceYaml,
	"artifacts/cluster-statefulset.yaml":    artifactsClusterStatefulsetYaml,
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
	"artifacts": {nil, map[string]*bintree{
		"backup-pvc.yaml":             {artifactsBackupPvcYaml, map[string]*bintree{}},
		"backupschedule-crd.yaml":     {artifactsBackupscheduleCrdYaml, map[string]*bintree{}},
		"backupschedule-cronjob.yaml": {artifactsBackupscheduleCronjobYaml, map[string]*bintree{}},
		"cluster-configmap.yaml":      {artifactsClusterConfigmapYaml, map[string]*bintree{}},
		"cluster-crd.yaml":            {artifactsClusterCrdYaml, map[string]*bintree{}},
		"cluster-service-read.yaml":   {artifactsClusterServiceReadYaml, map[string]*bintree{}},
		"cluster-service.yaml":        {artifactsClusterServiceYaml, map[string]*bintree{}},
		"cluster-statefulset.yaml":    {artifactsClusterStatefulsetYaml, map[string]*bintree{}},
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
