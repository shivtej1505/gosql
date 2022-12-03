package engine

import (
	"fmt"

	"github.com/shivtej1505/gosql/engine/context"
	"github.com/shivtej1505/gosql/engine/executor"
	"github.com/shivtej1505/gosql/engine/processor"
	"github.com/spf13/viper"
)

type Engine struct {
	Config    EngineConfig
	processor *processor.Processor
	executor  executor.Executor
}

type EngineConfig struct {
	Directory string `json:"directory"`
}

func NewEngine() Engine {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	var config EngineConfig

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}

	executorImpl := executor.NewExecutor(executor.ExecutorConfig{
		Directory: config.Directory,
	})

	return Engine{
		Config:    config,
		processor: processor.NewProcessor(processor.ProcessorConfig{}),
		executor:  executorImpl,
	}
}

func (engine Engine) RunQuery(query string) error {
	queryWithCtx := engine.processor.ParseQuery(query)

	fmt.Println(queryWithCtx)
	fmt.Println("-------")

	switch v := queryWithCtx.(type) {
	case context.InsertContext:
		return engine.executor.Insert(v)
	case context.SelectContext:
		return engine.executor.Select(v)
	case context.CreateTableContext:
		return engine.executor.CreateTable(v)
	}

	return nil
}
