package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shivtej1505/gosql/engine"
)

type Prompt struct {
	Engine engine.Engine
}

func NewPrompt(engine engine.Engine) Prompt {
	return Prompt{
		Engine: engine,
	}
}

// TODO: Implement history
func (prompt Prompt) StartPrompt() {
	//scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	var query string
	var delim byte
	var err error
	delim = ';'
	for {
		fmt.Printf("gosql>")
		query, err = reader.ReadString(delim)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", query)
		query = strings.TrimPrefix(query, "\n")
		query = strings.TrimSuffix(query, ";")
		prompt.Engine.RunQuery(query)
	}
}
