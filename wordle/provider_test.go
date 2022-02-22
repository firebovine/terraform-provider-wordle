/* straight up stolen from https://github.com/circa10a/terraform-provider-mcbroken/blob/main/mcbroken/provider_test.go
*/
package wordle

import (
  "testing"

  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider
var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
    testAccProvider = Provider()
    testAccProviders = map[string]*schema.Provider{
        "wordle": testAccProvider,
    }
    testAccProviderFactories = map[string]func() (*schema.Provider, error){
        "wordle": func() (*schema.Provider, error) {
            return testAccProvider, nil
        },
    }
}

func TestProvider(t *testing.T) {
    if err := testAccProvider.InternalValidate(); err != nil {
        t.Fatalf("err: %s", err)
    }
}

func TestProvider_Impl(t *testing.T) {
    var _ *schema.Provider = Provider()
}
