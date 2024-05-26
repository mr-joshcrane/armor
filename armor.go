package armor

import (
	"fmt"
	"io"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/json"
)

func getSchema(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Validate(policy string) error {
	ctx := cuecontext.New()

	cueSource, err := getSchema("cloudformation.cue")
	if err != nil {
		return fmt.Errorf("failed to read schema: %w", err)
	}
	schema := ctx.CompileString(cueSource).LookupPath(cue.ParsePath("#Policy"))
	if schema.Err() != nil {
		return fmt.Errorf("failed to compile schema: %w", schema.Err())
	}

	jsonExpr, err := json.Extract("", []byte(policy))
	if err != nil {
		return fmt.Errorf("failed to extract JSON: %w", err)
	}
	policyAsCue := ctx.BuildExpr(jsonExpr)
	unified := schema.Unify(policyAsCue)
	if err := unified.Validate(); err != nil {
		return fmt.Errorf("policy is invalid: %w", err)
	}
	return nil
}
