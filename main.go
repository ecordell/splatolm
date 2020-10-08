package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/operator-framework/api/pkg/operators/v1alpha1"
	"github.com/operator-framework/operator-lifecycle-manager/pkg/controller/registry/resolver"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "splatolm <csv file>",
		Short: "splatolm takes a CSV and generates the resources OLM would apply to a cluster",
		Long:  "splatolm takes a CSV and generates the resources OLM would apply to a cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("must specify csv to parse")
			}

			file, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer file.Close()

			var csv v1alpha1.ClusterServiceVersion
			if err := yaml.NewYAMLOrJSONDecoder(file, 50).Decode(&csv); err != nil {
				return err
			}

			steps, err := resolver.NewServiceAccountStepResources(&csv, "", "")
			if err != nil {
				return err
			}
			for _, s := range steps {
				var manifest map[string]interface{}
				if err := json.Unmarshal([]byte(s.Manifest), &manifest); err != nil {
					return err
				}
				manifest["apiVersion"] = strings.Join([]string{s.Group, s.Version}, "/")
				manifest["kind"] = s.Kind
				out, err := json.Marshal(manifest)
				if err != nil {
					return err
				}
				fmt.Println(string(out))
			}
			return nil
		},
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
