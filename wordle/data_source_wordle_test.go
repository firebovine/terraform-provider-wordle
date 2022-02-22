package wordle

import (
  "fmt"
  "testing"

  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccWordle(t *testing.T) {
  resource.ParallelTest(t, resource.TestCase{
	PreCheck: func() { /* no precheck needed testAccPreCheck(t) */ },
	ProviderFactories: testAccProviderFactories,
	Steps: []resource.TestStep{
		{
			Config: testAccDate("2022-01-01T00:00:00Z"),
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr(
					"data.wordle_word.test", "word", "rebus"),
				resource.TestCheckResourceAttrSet(
					"data.wordle_word.test", "date"),
			),
		},
	},
  })
}

func testAccDate(date string) string {
  return fmt.Sprintf(`
data "wordle_word" "test" {
  date = "%[1]v"
}
`, date)
}
