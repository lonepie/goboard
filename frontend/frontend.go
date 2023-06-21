package frontend

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/static"
)

//go:embed dist/*
var reactStatic embed.FS

// func GetFS() http.FileSystem {
// 	sub, err := fs.Sub(reactStatic, "dist")
// 	if err != nil {
// 		log.Fatalf("Error with embedded fs: %v", err)
// 	}
// 	return http.FS(sub)
// }

type embedFileSystem struct {
	http.FileSystem
	indexes bool
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	f, err := e.Open(path)
	if err != nil {
		return false
	}

	// check if indexing is allowed
	s, _ := f.Stat()
	if s.IsDir() && !e.indexes {
		return false
	}

	return true
}

func EmbedFolder(fsEmbed embed.FS, targetPath string, index bool) static.ServeFileSystem {
	subFS, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(subFS),
		indexes:    index,
	}
}

func GetFS(indexes bool) static.ServeFileSystem {
	return EmbedFolder(reactStatic, "dist", indexes)
}
