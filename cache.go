package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
)

var routeNotFoundError = errors.New("Route not found in cache")

type routeCache interface {
	StoreRoute(from, to starSystem, distance int64, hops []traceResult) error
	GetRoute(from, to starSystem, distance int64) ([]traceResult, error)
}

func getCache(uri string) (routeCache, error) {
	parsed, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	switch parsed.Scheme {
	case "file":
		return newFileSystemCache(strings.Replace(uri, "file://", "", -1))
	case "none":
		return noCache{}, nil
	}

	return nil, errors.New("Specified route cache was not found")
}

type fileSystemCache struct {
	path string
}

func newFileSystemCache(storePath string) (*fileSystemCache, error) {
	if storePath, err := homedir.Expand(storePath); err != nil {
		return nil, err
	} else {
		return &fileSystemCache{path: storePath}, nil
	}
}

func (f fileSystemCache) StoreRoute(from, to starSystem, distance int64, hops []traceResult) error {
	storePath := f.getFilename(from, to, distance)

	if err := os.MkdirAll(f.path, 0755); err != nil {
		return err
	}

	fp, err := os.Create(storePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	return json.NewEncoder(fp).Encode(hops)
}

func (f fileSystemCache) GetRoute(from, to starSystem, distance int64) ([]traceResult, error) {
	storePath := f.getFilename(from, to, distance)

	if stat, err := os.Stat(storePath); err != nil {
		return nil, routeNotFoundError
	} else {
		if time.Now().Sub(stat.ModTime()) > cfg.CacheTime {
			return nil, routeNotFoundError
		}
	}

	fp, err := os.Open(storePath)
	if err != nil {
		return nil, err
	}

	result := []traceResult{}
	return result, json.NewDecoder(fp).Decode(&result)
}

func (f fileSystemCache) getFilename(from, to starSystem, distance int64) string {
	return path.Join(f.path, fmt.Sprintf("%d_%d_%d.json", from.ID, to.ID))
}

type noCache struct{}

func (n noCache) StoreRoute(from, to starSystem, distance int64, hops []traceResult) error { return nil }
func (n noCache) GetRoute(from, to starSystem, distance int64) ([]traceResult, error) {
	return nil, routeNotFoundError
}
