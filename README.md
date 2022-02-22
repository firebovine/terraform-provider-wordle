# terraform-provider-wordle <img src="https://i.imgur.com/fAS7XqO.png" height="5%" width="5%" align="left"/>

Base the names of your infrastructure on the Wordle of the day! Powered by [Wordle](https://www.nytimes.com/games/wordle/index.html)

- [terraform-provider-wordle](#terraform-provider-wordle)
  * [Usage](#usage)
  * [Development](#development)
    + [Linting](#linting)
    + [Testing](#testing)
    + [Mac](#mac)
    + [Linux](#linux)
    + [Windows](#windows)

## Usage

- View the provider on the [Hashicorp Registry](https://registry.terraform.io/providers/firebovine/wordle/latest/docs)

```hcl
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

# Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

# Outputs:

# example_that = "rebus"
# example_this = "other"
```

## Development

### Linting
```bash
make lint
```

### Testing
```bash
make test
```

### Mac

```bash
echo no
```

### Linux
```bash
make install && \
cd examples && \
terraform init && \
terraform apply
```

### Windows

![alt text](https://media.giphy.com/media/4cuyucPeVWbNS/giphy.gif)
