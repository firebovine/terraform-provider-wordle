terraform {
  required_providers {
    wordle = {
      source  = "firebovine/wordle"
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
