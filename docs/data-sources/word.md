# wordle_word Data Source

// Get wordle word of the day

## Example Usage

```hcl
// Data source to get current wordle word
data "wordle_word" "today" {}

// Data source to get word from 2022-01-01
data "wordle_word" "first" {
  date = "2022-01-01T00:00:00Z"
}

// Get wordle word
output "todays_word" {
    value = data.wordle_word.today.word
}
```

## Argument Reference

* `date` - (Optional) RFC3339 timestamp of the desired day's wordle

## Attribute Reference

* `word` - wordle word of the day
