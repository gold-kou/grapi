// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package svcgen

import (
	"github.com/izumin5210/grapi/pkg/cli"
	"github.com/izumin5210/grapi/pkg/gencmd"
	"github.com/izumin5210/grapi/pkg/grapicmd"
	"github.com/izumin5210/grapi/pkg/protoc"
)

import (
	_ "github.com/izumin5210/grapi/pkg/svcgen/template"
)

// Injectors from wire.go:

func NewApp(command *gencmd.Command) (*App, error) {
	ctx := gencmd.ProvideCtx(command)
	grapicmdCtx := gencmd.ProvideGrapiCtx(ctx)
	config := grapicmd.ProvideProtocConfig(grapicmdCtx)
	fs := grapicmd.ProvideFS(grapicmdCtx)
	execInterface := grapicmd.ProvideExecer(grapicmdCtx)
	io := grapicmd.ProvideIO(grapicmdCtx)
	ui := cli.UIInstance(io)
	rootDir := grapicmd.ProvideRootDir(grapicmdCtx)
	gexConfig := protoc.ProvideGexConfig(fs, execInterface, io, rootDir)
	repository, err := protoc.ProvideToolRepository(gexConfig)
	if err != nil {
		return nil, err
	}
	wrapper := protoc.NewWrapper(config, fs, execInterface, ui, repository, rootDir)
	grapicmdConfig := grapicmd.ProvideConfig(grapicmdCtx)
	builder := ProvideParamsBuilder(rootDir, config, grapicmdConfig)
	app := &App{
		ProtocWrapper: wrapper,
		ParamsBuilder: builder,
	}
	return app, nil
}
