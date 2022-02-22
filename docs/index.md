# Wordle Provider

Base the names of your infrastructure on the Wordle of the day! Powered by [Wordle](https://www.nytimes.com/games/wordle/index.html)

## Example Usage

```hcl
terraform {
  required_providers {
    wordle = {
      source  = "firebovine/mcbroken"
    }
  }
}

provider "wordle" {}

data "wordle_word" "this" {}

data "wordle_word" "that" {
  date = "2022-01-01T00:00:00Z"
}

output "example_this" {
  value = data.wordle_word.this.word
}

output "example_that" {
  value = data.wordle_word.that.word
}

# Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

# Outputs:

# example_that = "rebus"
# example_this = "other"
```
