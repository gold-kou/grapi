// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"github.com/izumin5210/clig/pkg/clib"
	"github.com/izumin5210/gex/pkg/tool"
	"github.com/izumin5210/grapi/pkg/cli"
	"github.com/izumin5210/grapi/pkg/excmd"
	"github.com/izumin5210/grapi/pkg/grapicmd"
	"github.com/izumin5210/grapi/pkg/grapicmd/internal/module"
	"github.com/izumin5210/grapi/pkg/grapicmd/internal/usecase"
	"github.com/izumin5210/grapi/pkg/protoc"
	"github.com/rakyll/statik/fs"
)

// Injectors from wire.go:

func NewUI(ctx *grapicmd.Ctx) cli.UI {
	io := grapicmd.ProvideIO(ctx)
	ui := cli.UIInstance(io)
	return ui
}

func NewCommandExecutor(ctx *grapicmd.Ctx) excmd.Executor {
	io := grapicmd.ProvideIO(ctx)
	executor := excmd.NewExecutor(io)
	return executor
}

func NewScriptLoader(ctx *grapicmd.Ctx) module.ScriptLoader {
	io := grapicmd.ProvideIO(ctx)
	executor := excmd.NewExecutor(io)
	scriptLoader := ProvideScriptLoader(ctx, executor)
	return scriptLoader
}

func NewToolRepository(ctx *grapicmd.Ctx) (tool.Repository, error) {
	fs := grapicmd.ProvideFS(ctx)
	execInterface := grapicmd.ProvideExecer(ctx)
	io := grapicmd.ProvideIO(ctx)
	rootDir := grapicmd.ProvideRootDir(ctx)
	config := protoc.ProvideGexConfig(fs, execInterface, io, rootDir)
	repository, err := protoc.ProvideToolRepository(config)
	if err != nil {
		return nil, err
	}
	return repository, nil
}

func NewProtocWrapper(ctx *grapicmd.Ctx) (protoc.Wrapper, error) {
	config := grapicmd.ProvideProtocConfig(ctx)
	fs := grapicmd.ProvideFS(ctx)
	execInterface := grapicmd.ProvideExecer(ctx)
	io := grapicmd.ProvideIO(ctx)
	ui := cli.UIInstance(io)
	rootDir := grapicmd.ProvideRootDir(ctx)
	gexConfig := protoc.ProvideGexConfig(fs, execInterface, io, rootDir)
	repository, err := protoc.ProvideToolRepository(gexConfig)
	if err != nil {
		return nil, err
	}
	wrapper := protoc.NewWrapper(config, fs, execInterface, ui, repository, rootDir)
	return wrapper, nil
}

func NewInitializeProjectUsecase(ctx *grapicmd.Ctx, path clib.Path) (usecase.InitializeProjectUsecase, error) {
	aferoFs := grapicmd.ProvideFS(ctx)
	execInterface := grapicmd.ProvideExecer(ctx)
	io := grapicmd.ProvideIO(ctx)
	rootDir := grapicmd.ProvideRootDir(ctx)
	config := protoc.ProvideGexConfig(aferoFs, execInterface, io, rootDir)
	ui := cli.UIInstance(io)
	fileSystem, err := fs.New()
	if err != nil {
		return nil, err
	}
	generator := ProvideGenerator(ctx, ui, aferoFs, fileSystem, path)
	executor := excmd.NewExecutor(io)
	initializeProjectUsecase := ProvideInitializeProjectUsecase(ctx, config, ui, aferoFs, generator, executor)
	return initializeProjectUsecase, nil
}
