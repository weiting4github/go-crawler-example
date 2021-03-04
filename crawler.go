package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/urfave/cli/v2"
)

const domain = "https://www."

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "top",
				Usage: "show top <number> sites URL on www.alexa.com/topsites/",
			},
			&cli.StringFlag{
				Name:  "country",
				Usage: "show top 20 sites URL on www.alexa.com/topsites/ by country",
			},
		},
		Action: func(c *cli.Context) error {
			// fmt.Println("IsSet:", c.IsSet("top"))
			// fmt.Println(showTOP(c.String("top")))
			if c.IsSet("top") {
				fmt.Println(showTOP(c.String("top")))
				return nil
			}

			if c.IsSet("country") {
				fmt.Println(showTOP20(c.String("country")))
				return nil
			}

			fmt.Println("Use clawer --help for more information.")
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func showTOP(topNum string) string {
	r := "nothing has been found"
	c := colly.NewCollector()
	c.OnHTML("div.tr.site-listing ", func(e *colly.HTMLElement) {
		// fmt.Println(e.ChildText("div:nth-child(2)"))
		// fmt.Println(topNum)
		if e.ChildText("div:first-child") == topNum { // number
			// fmt.Println(e)
			e.ForEach("div:nth-child(2) a[href]", func(_ int, el *colly.HTMLElement) {
				// fmt.Println(domain + el.Attr("href"))
				r = domain + el.Text

			})
		}
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit("https://www.alexa.com/topsites/")
	return r
}

func showTOP20(country string) string {
	count := 0
	c := colly.NewCollector()
	c.OnHTML("div.tr.site-listing ", func(e *colly.HTMLElement) {
		if count >= 20 {
			return
		}
		fmt.Println(e.ChildText("div:nth-child(2)"))
		// fmt.Println(e.Index)

		// fmt.Println(e.ChildText("div:nth-child(2)"))
		count++
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit("https://www.alexa.com/topsites/countries/" + country)
	if count == 0 {
		return "nothing has been found"
	}

	return ""
}
