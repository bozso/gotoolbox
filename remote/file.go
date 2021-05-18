package remote

import (
	"bufio"

	"github.com/gofiber/fiber"

	"github.com/bozso/gotoolbox/path"
	"github.com/bozso/gotoolbox/services"
)

type FileGetter interface {
	GetRemoteFile(remote path.ValidFile) (local path.ValidFile, err error)
}

type HTTPFileService interface {
	SetupRemoteFileService(*fiber.Ctx) error
	GetRemoteFile(*fiber.Ctx) error
	GetRemoteContents(*fiber.Ctx) error
}

type FileService struct {
	remote FileGetter
}

func (f *FileService) SetupRemoteFileService(ctx *fiber.Ctx) (err error) {
	return
}

func (fs *FileService) get(g parsing.Getter) (local path.ValidFile, err error) {
	remote, err := services.MustGet(g, "remote_path")
	if err != nil {
		return
	}

	file, err := path.New(remote).ToValidFile()
	if err != nil {
		return
	}

	local, err := fs.remote.GetRemoteFile(file)
	if err != nil {
		return
	}
}

func (fs *FileService) GetRemoteFile(ctx *fiber.Ctx) (err error) {
	query := services.UrlQuery.FromFiberCtx(ctx)

	local, err := fs.get(query)
	if err != nil {
		return
	}

	ctx.SendString(local.String())
}
