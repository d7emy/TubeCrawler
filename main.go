package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	Cache          *cache
	channelChannel = make(chan browserStruct)

	videoDataSet = map[string]bool{}
)

const (
	resultPath = "results/result.csv"
)

type browserStruct struct {
	browseID    string
	channelUrl  string
	channelName string
}

func RandomString(Length int, br ...string) string {
	r := "mznxbcvalskdjfhgpqowieuryt1029384756"
	I := 36
	if br != nil {
		r = "QPWOEIRUTYALSKDJFHGMZNXBCVmznxbcvalskdjfhgpqowieuryt1029384756" + br[0]
		I = len(r)
	}
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, Length)
	for i := range b {
		b[i] = r[seededRand.Intn(I)]
	}
	return string(b)
}

// pesron key
var searchKeyword = "عبودي باد"
var whiteKeys = [][]string{ //keys
	//example
	{"abodybad"},
	{"abody", "bad"},
	{"عبودي باد"},
	{"باد", "عبودي"},
}

func main() {
	Cache = NewCache()

	//for anti-duplicate
	if fileExists(resultPath) {
		for _, line := range readCSV(resultPath) {
			Cache.Set(line[0])
		}
	}

	//start searching for the person using the searchKeyword variable
	startSearch()

	//crawling all the channels looking for the videos contains whitekey
	threads := 10
	for i := 0; i < threads; i++ {
		go startScrapFromChannel()
	}
	channelCache := NewCache()
	for _, line := range readCSV(resultPath) {
		if !channelCache.SetCheck(line[7]) {
			channelChannel <- browserStruct{line[7], line[3], line[2]}
		}
	}

	select {}
}

func isContainsKeys(s string) bool {
	if len(whiteKeys) == 0 {
		return true
	}
	s = strings.ToLower(s)
	for _, key := range whiteKeys {
		res := true
		for _, word := range key {
			res = res && strings.Contains(s, word)
			if !res {
				break
			}
		}
		if res {
			return res
		}
	}
	return false
}

// for browsing all channels for the searchkey
func startScrapFromChannel() {

	//contentIndex := 0

	for c := range channelChannel {
		continuation := ""
		requests := 0
		id := c.browseID
		bUser := c.channelUrl
		bName := c.channelName
		contents := []browseContent{}
		zeroContents := 0
		for {
			exit := false
			requests++
			//ag:
			if continuation == "" {
				resp := browse(id)
				for _, tab := range resp.Contents.TwoColumnBrowseResultsRenderer.Tabs {
					if tab.TabRenderer.Title == "Videos" {
						contents = tab.TabRenderer.Content.RichGridRenderer.Contents
						continuation = ""
						for _, content := range contents {
							if content.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token != "" {
								continuation = content.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token
							}
						}
						if continuation == "" {
							exit = true
						}
					}
				}
			} else {
				resp := browseByCont(id, continuation)
				if len(resp.OnResponseReceivedActions) == 0 {
					break
				}
				contents = resp.OnResponseReceivedActions[0].AppendContinuationItemsAction.ContinuationItems

				continuation = ""
				for _, content := range contents {
					if content.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token != "" {
						continuation = content.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token
					}
				}
				if continuation == "" {
					fmt.Printf("Done, %s Channel, %d Requests\n", id, requests)
					exit = true
				}
			}
			if len(contents) == 0 {
				zeroContents++
				if zeroContents > 5 {
					break
				}
			}
			for _, content := range contents {
				if content.ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token != "" {
					continue
				}
				//fmt.Printf("Video https://www.youtube.com/watch?v=%s\n", content.VideoRenderer.VideoID)
				ID := content.RichItemRenderer.Content.VideoRenderer.VideoID
				Title := content.RichItemRenderer.Content.VideoRenderer.Title.Runs[0].Text
				Channel_Name := bName
				Channel_Url := bUser
				Published_Time := content.RichItemRenderer.Content.VideoRenderer.PublishedTimeText.SimpleText
				Long := content.RichItemRenderer.Content.VideoRenderer.LengthText.SimpleText
				View_Count := content.RichItemRenderer.Content.VideoRenderer.ViewCountText
				Browse_ID := id
				fmt.Println(ID, Title, Channel_Name, Channel_Url, Published_Time, Long, View_Count.SimpleText, Browse_ID)
				if isContainsKeys(Title) && !Cache.SetCheck(ID) {
					appendMultiCSV(resultPath, [][]string{{ID, Title, Channel_Name, Channel_Url, Published_Time, Long, View_Count.SimpleText, Browse_ID}})
				}
			}
			if exit {
				fmt.Printf("Scrapped All %s Channel, from first Request\n", id)
				break
			}
		}

	}
	//fmt.Println(continuation)

}
