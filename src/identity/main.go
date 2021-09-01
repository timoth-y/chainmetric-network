package main

import (
	"github.com/timoth-y/chainmetric-contracts/shared/core"
	"github.com/timoth-y/chainmetric-contracts/shared/server"
	"github.com/timoth-y/chainmetric-contracts/shared/utils"
	"github.com/timoth-y/chainmetric-contracts/src/identity/api/middleware"
	"github.com/timoth-y/chainmetric-contracts/src/identity/api/rpc"
	"github.com/timoth-y/chainmetric-contracts/src/identity/usecase/identity"
	"github.com/timoth-y/chainmetric-contracts/src/identity/usecase/privileges"
)

func init() {
	core.InitCore()
	core.InitMongoDB()
	privileges.Init()
	utils.MustExecute(identity.Init, "failed to initialize identity package")
	utils.MustExecute(func() error {
		return server.Init(
			server.WithUnaryMiddlewares(middleware.AuthWithJWTForUnaryGRPC(
				"UserService/register",
				"AccessService/requestFabricCredentials",
				"AccessService/requestFabricCredentials",
			)),
			server.WithStreamMiddlewares(middleware.AuthWithJWTForStreamGRPC()),
			server.WithServiceRegistrar(rpc.RegisterAdminService),
			server.WithServiceRegistrar(rpc.RegisterAccessService),
		)
	}, "failed to initialize server")
}

func main() {
	if err := server.Serve(":8080"); err != nil {
		core.Logger.Fatal(err)
	}
}
