terraform {
  required_providers {
    wordle = {
      version = "0.1"
      source  = "hashicorp.com/nwheeler-splunk/wordle"
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
