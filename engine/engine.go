package engine

import (
	"fmt"

	"github.com/shivtej1505/gosql/engine/context"
	"github.com/shivtej1505/gosql/engine/executor"
	"github.com/shivtej1505/gosql/engine/processor"
	"github.com/spf13/viper"
)

type Engine struct {
	Config
	Processor processor.Processor
	Executor  executor.Executor
}

type Config struct {
	Directory string `json:"directory"`
}

func NewEngine() Engine {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	var config Config

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}

	fmt.Println(viper.GetString("directory"))

	return Engine{
		Config: config,
	}
}

func (engine Engine) RunQuery(query string) error {
	queryWithCtx := engine.Processor.ParseQuery(query)

	fmt.Println(queryWithCtx)
	fmt.Println("-------")

	switch v := queryWithCtx.(type) {
	case context.InsertContext:
		return engine.Executor.Insert(v)
	case context.SelectContext:
		return engine.Executor.Select(v)
	case context.CreateTableContext:
		return engine.Executor.CreateTable(v)
	}

	return nil
}
