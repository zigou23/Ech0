// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	handler3 "github.com/lin-snow/ech0/internal/handler/common"
	handler6 "github.com/lin-snow/ech0/internal/handler/connect"
	handler2 "github.com/lin-snow/ech0/internal/handler/echo"
	handler4 "github.com/lin-snow/ech0/internal/handler/setting"
	handler5 "github.com/lin-snow/ech0/internal/handler/todo"
	"github.com/lin-snow/ech0/internal/handler/user"
	repository2 "github.com/lin-snow/ech0/internal/repository/common"
	repository5 "github.com/lin-snow/ech0/internal/repository/connect"
	repository3 "github.com/lin-snow/ech0/internal/repository/echo"
	"github.com/lin-snow/ech0/internal/repository/keyvalue"
	repository4 "github.com/lin-snow/ech0/internal/repository/todo"
	"github.com/lin-snow/ech0/internal/repository/user"
	"github.com/lin-snow/ech0/internal/service/common"
	service6 "github.com/lin-snow/ech0/internal/service/connect"
	service4 "github.com/lin-snow/ech0/internal/service/echo"
	service2 "github.com/lin-snow/ech0/internal/service/setting"
	service5 "github.com/lin-snow/ech0/internal/service/todo"
	service3 "github.com/lin-snow/ech0/internal/service/user"
	"gorm.io/gorm"
)

// Injectors from wire.go:

// BuildHandlers 使用wire生成的代码来构建Handlers实例
func BuildHandlers(db *gorm.DB) (*Handlers, error) {
	userRepositoryInterface := repository.NewUserRepository(db)
	commonRepositoryInterface := repository2.NewCommonRepository(db)
	commonServiceInterface := service.NewCommonService(commonRepositoryInterface)
	keyValueRepositoryInterface := keyvalue.NewKeyValueRepository(db)
	settingServiceInterface := service2.NewSettingService(commonServiceInterface, keyValueRepositoryInterface)
	userServiceInterface := service3.NewUserService(userRepositoryInterface, settingServiceInterface)
	userHandler := handler.NewUserHandler(userServiceInterface)
	echoRepositoryInterface := repository3.NewEchoRepository(db)
	echoServiceInterface := service4.NewEchoService(commonServiceInterface, echoRepositoryInterface)
	echoHandler := handler2.NewEchoHandler(echoServiceInterface)
	commonHandler := handler3.NewCommonHandler(commonServiceInterface)
	settingHandler := handler4.NewSettingHandler(settingServiceInterface)
	todoRepositoryInterface := repository4.NewTodoRepository(db)
	todoServiceInterface := service5.NewTodoService(todoRepositoryInterface, commonServiceInterface)
	todoHandler := handler5.NewTodoHandler(todoServiceInterface)
	connectRepositoryInterface := repository5.NewConnectRepository(db)
	connectServiceInterface := service6.NewConnectService(connectRepositoryInterface, echoRepositoryInterface, commonServiceInterface, settingServiceInterface)
	connectHandler := handler6.NewConnectHandler(connectServiceInterface)
	handlers := NewHandlers(userHandler, echoHandler, commonHandler, settingHandler, todoHandler, connectHandler)
	return handlers, nil
}

// wire.go:

// UserSet 包含了构建 UserHandler 所需的所有 Provider
var UserSet = wire.NewSet(repository.NewUserRepository, service3.NewUserService, handler.NewUserHandler)

// EchoSet 包含了构建 EchoHandler 所需的所有 Provider
var EchoSet = wire.NewSet(repository3.NewEchoRepository, service4.NewEchoService, handler2.NewEchoHandler)

// CommonSet 包含了构建 CommonHandler 所需的所有 Provider
var CommonSet = wire.NewSet(repository2.NewCommonRepository, service.NewCommonService, handler3.NewCommonHandler)

// SettingSet 包含了构建 SettingHandler 所需的所有 Provider
var SettingSet = wire.NewSet(keyvalue.NewKeyValueRepository, service2.NewSettingService, handler4.NewSettingHandler)

// TodoSet 包含了构建 TodoHandler 所需的所有 Provider
var TodoSet = wire.NewSet(repository4.NewTodoRepository, service5.NewTodoService, handler5.NewTodoHandler)

// ConnectSet 包含了构建 ConnectHandler 所需的所有 Provider
var ConnectSet = wire.NewSet(repository5.NewConnectRepository, service6.NewConnectService, handler6.NewConnectHandler)
