// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cmd

import (
	grapicmd "github.com/izumin5210/grapi/pkg/grapicmd"
	di "github.com/izumin5210/grapi/pkg/grapicmd/di"
	cobra "github.com/spf13/cobra"
)

// Injectors from wire.go:

func NewGrapiCommand(cfg *grapicmd.Config) *cobra.Command {
	config := di.ProvideGexConfig(cfg)
	ui := di.ProvideUI(cfg)
	generator := di.ProvideGenerator(cfg, ui)
	initializeProjectUsecase := di.ProvideInitializeProjectUsecase(cfg, config, ui, generator)
	initCmd2 := provideInitCommand(cfg, initializeProjectUsecase)
	executor := di.ProvideCommandExecutor(cfg, ui)
	executeProtocUsecase := di.ProvideExecuteProtocUsecase(cfg, config, ui, executor, generator)
	genSvcCmd2 := provideGenerateServiceCommand(cfg, ui, generator, executeProtocUsecase)
	genScaffoldSvcCmd2 := provideGenerateScaffoldServiceCommand(cfg, ui, generator, executeProtocUsecase)
	genCmdCmd2 := provideGenerateCommandCommand(cfg, generator)
	generateCmd2 := provideGenerateCommand(genSvcCmd2, genScaffoldSvcCmd2, genCmdCmd2)
	destroySvcCmd2 := provideDestroyServiceCommand(cfg, generator)
	destroyCmdCmd2 := provideDestroyCommandCommand(cfg, generator)
	destroyCmd2 := provideDestroyCommand(destroySvcCmd2, destroyCmdCmd2)
	protocCmd2 := provideProtocCommand(cfg, executeProtocUsecase)
	scriptLoader := di.ProvideScriptLoader(cfg, executor)
	buildCmd2 := provideBuildCommand(cfg, ui, scriptLoader)
	versionCmd2 := provideVersionCommand(cfg)
	userDefinedCmds2 := provideUserDefinedCommands(cfg, ui, scriptLoader)
	command := provideGrapiCommand(cfg, initCmd2, generateCmd2, destroyCmd2, protocCmd2, buildCmd2, versionCmd2, userDefinedCmds2)
	return command
}
