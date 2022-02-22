package wordle

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceWordle() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWordleRead,
		Schema: map[string]*schema.Schema{
			"date": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsRFC3339Time,
				DefaultFunc: func() (interface{}, error) {
					var t = time.Now()
					midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
					return midnight.Format(time.RFC3339), nil
				},
			},
			"word": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceWordleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	date := d.Get("date").(string)

	var diags diag.Diagnostics

	mainjs, err := getMainJS()
	if err != nil {
		return diag.FromErr(err)
	}

	idx, err := getIdxFromTime(date, mainjs)
	if err != nil {
		return diag.FromErr(err)
	}

	word, err := getWordleWord(mainjs, idx)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("word", word); err != nil {
		return diag.FromErr(err)
	}

	// always run
	numDecimals := 10
	d.SetId(strconv.FormatInt(time.Now().Unix(), numDecimals))

	return diags
}

func getIdxFromTime(date string, mainjs string) (result int, err error) {
	userDate, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return -1, err
	}

	userMidnight := time.Date(userDate.Year(), userDate.Month(), userDate.Day(), 0, 0, 0, 0, time.Local)
	var msInDay int64 = 86400000

	wordleEpoch := getWordleEpoch(mainjs)

	if userMidnight.UnixMilli() < wordleEpoch {
		swordleEpoch := time.UnixMilli(wordleEpoch).Format(time.RFC3339)
		return -1, fmt.Errorf("time must be after the wordle epoch: %s", swordleEpoch)
	}

	delta := userMidnight.UnixMilli() - wordleEpoch
	var idx = int(math.Round(float64(delta / msInDay)))

	return idx, nil
}

func getMainJS() (js string, err error) {
	// get the main js file
	urlBase := "https://www.nytimes.com/games/wordle/"
	urlIndex := urlBase + "index.html"
	res, err := http.Get(urlIndex)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}
	mainjs := urlBase
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		src, _ := s.Attr("src")
		r := regexp.MustCompile(`^main\.[a-f0-9]+\.js$`)
		if r.MatchString(src) {
			mainjs = mainjs + src
		}
	})

	res, err = http.Get(mainjs) //nolint:gosec
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	contents, _ := ioutil.ReadAll(res.Body)
	return string(contents), nil
}

func getWordleEpoch(mainjs string) int64 {
	// calebwhy
	r := regexp.MustCompile(`Date\(([0-9]{4}),([0-9]{1,2}),([0-9]{1,2}),[0-9,]+\)`)
	epoch := r.FindStringSubmatch(mainjs)
	year, _ := strconv.Atoi(epoch[1])
	_month, _ := strconv.Atoi(epoch[2])
	_month++
	month := time.Month(_month)
	day, _ := strconv.Atoi(epoch[3])
	wordleEpoch := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return wordleEpoch.UnixMilli()
}

func getWordleWord(mainjs string, idx int) (word string, err error) {
	// calebwhy
	r := regexp.MustCompile(`[A-Za-z]+=\[([a-zA-Z,\"]+)\]`)
	var result string
	matches := r.FindAllStringSubmatch(mainjs, -1)
	for _, match := range matches {
		value := match[1]
		// a token value, get it?
		if strings.Contains(value, "token") {
			words := strings.Split(value, ",")
			midx := idx % len(words)
			result = strings.Trim(words[midx], "\"")
			return result, nil
		}
	}
	return "", fmt.Errorf("could not find a wordle word")
}
