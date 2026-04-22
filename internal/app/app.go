package app

func Run() error {
	cfg := loadConfig()

	database, storage, err := initInfrastructure(cfg)
	if err != nil {
		return err
	}

	deps := initDependencies(cfg, database, storage)

	router := setupRouter(deps)

	return router.Run(":8080")
}
