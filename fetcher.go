package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var (
	baseUrl = "https://myanimelist.net/topanime.php"
)

type AnimeTitle struct {
	En  string
	Jp  string
	Syn string
}

type Review struct {
	Name    string
	Content string
	Date    string
}

type AnimeData struct {
	Title         AnimeTitle
	ImageUrl      string
	Type          string
	Episodes      string
	Status        string
	Aired         string
	Studios       string
	Genres        []string
	Duration      string
	Rating        string
	Score         string
	UserVoteCount string
	Ranked        string
	Popularity    string
	Synopsis      string
	Reviews       []Review
}

func getAnimeDetail(c *colly.Collector, href string) AnimeData {
	result := map[string]string{}
	genres := []string{}
	reviews := []Review{}
	reviewCount := 3

	c.OnHTML("div.spaceit_pad", func(h *colly.HTMLElement) {
		class := strings.Split(h.Attr("class"), " ")
		if len(class) > 1 || class[0] != "spaceit_pad" {
			return
		}

		res := strings.Split(h.Text, ":")
		if len(res) != 2 {
			return
		}

		for i := 0; i < len(res); i++ {
			res[i] = strings.TrimSpace(res[i])
		}

		result[res[0]] = res[1]
	})

	c.OnHTML("h1.title-name.h1_bold_none", func(h *colly.HTMLElement) {
		result["English"] = h.ChildText("strong")
	})

	c.OnHTML("div.leftside", func(h *colly.HTMLElement) {
		result["ImageUrl"] = h.ChildAttr("img.lazyload", "data-src")
	})

	c.OnHTML("span", func(h *colly.HTMLElement) {
		if h.Attr("itemprop") != "genre" {
			return
		}

		genres = append(genres, h.Text)
	})

	c.OnHTML("div.fl-l.score", func(h *colly.HTMLElement) {
		result["UserVoteCount"] = h.Attr("data-user")
		result["Score"] = h.Text
	})

	c.OnHTML("span.numbers.ranked", func(h *colly.HTMLElement) {
		result["Ranked"] = h.ChildText("strong")[1:]
	})

	c.OnHTML("span.numbers.popularity", func(h *colly.HTMLElement) {
		result["Popularity"] = h.ChildText("strong")[1:]
	})

	c.OnHTML("p", func(h *colly.HTMLElement) {
		if h.Attr("itemprop") != "description" {
			return
		}

		result["Synopsis"] = h.Text
	})

	c.OnHTML("div.review-element", func(h *colly.HTMLElement) {
		if len(reviews) >= reviewCount {
			return
		}

		r := Review{
			Name:    h.ChildText("div.username"),
			Content: h.ChildText("div.text"),
			Date:    h.ChildText("div.update_at"),
		}

		reviews = append(reviews, r)
	})

	c.Visit(href)

	animeData := AnimeData{
		Title:         AnimeTitle{En: result["English"], Jp: result["Japanese"], Syn: result["Synonyms"]},
		ImageUrl:      result["ImageUrl"],
		Type:          result["Type"],
		Episodes:      result["Episodes"],
		Status:        result["Status"],
		Aired:         result["Aired"],
		Studios:       result["Studios"],
		Genres:        genres,
		Duration:      result["Duration"],
		Rating:        result["Rating"],
		Score:         result["Score"],
		UserVoteCount: result["UserVoteCount"],
		Ranked:        result["Ranked"],
		Popularity:    result["Popularity"],
		Synopsis:      result["Synopsis"],
		Reviews:       reviews,
	}

	return animeData
}

func GetAnime(c *colly.Collector) AnimeData {
	var animeData AnimeData
	limit := GetRandomPage()
	animeNum := GetRandomNumber(1, 50)
	counter := 0

	c.OnHTML("a[href].hoverinfo_trigger", func(h *colly.HTMLElement) {
		class := strings.Split(h.Attr("class"), " ")
		if len(class) > 1 || class[0] != "hoverinfo_trigger" {
			return
		}

		counter++
		if counter != animeNum {
			return
		}

		href := h.Attr("href")
		animeData = getAnimeDetail(c, href)
	})

	if limit != 0 {
		baseUrl = fmt.Sprintf("%s?limit=%d", baseUrl, limit)
	}
	c.Visit(baseUrl)

	return animeData
}
