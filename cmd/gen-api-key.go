package cmd

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tkhq/tkcli/internal/apikey"
	"github.com/tkhq/tkcli/internal/clifs"
)

var genApiKey = &cobra.Command{
	Use:     "generate-api-key generates a Turnkey API key",
	Short:   "generate-api-key generates a Turnkey API key",
	Aliases: []string{"g", "gen"},
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.PersistentFlags().GetString("key")
		if err != nil {
			return errors.Wrap(err, "failed to read key name")
		}

		enc := json.NewEncoder(cmd.OutOrStdout())
		enc.SetIndent("", "   ")

		apiKey, err := apikey.NewTkApiKey()
		if err != nil {
			return errors.Wrap(err, "failed to create keypair")
		}

		if name == "-" {
			return enc.Encode(map[string]interface{}{
				"publicKey":  apiKey.TkPublicKey,
				"privateKey": apiKey.TkPrivateKey,
			})
		}

		if err = clifs.StoreKeypair(name, apiKey); err != nil {
			return errors.Wrap(err, "failed to store new keypair")
		}

		return enc.Encode(map[string]interface{}{
			"publicKey":      apiKey.TkPublicKey,
			"publicKeyFile":  clifs.PublicKeyFile(name),
			"privateKeyFile": clifs.PrivateKeyFile(name),
		})
	},
}
