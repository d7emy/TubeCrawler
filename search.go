package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ID,Title,Channel Name,Channel url,Published Time,Long,View Count
func startSearch() {
	continuation := ""
	contents := []searchContent{}
	contentIndex := 0
	key := searchKeyword
	for {
		if continuation == "" {
			resp := search(key)
			contentIndex = 0
			contents = resp.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents[contentIndex].ItemSectionRenderer.Contents
			if len(contents) == 0 {
				contentIndex = 1
				contents = resp.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents[contentIndex].ItemSectionRenderer.Contents
			}
			continuation = resp.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents[contentIndex^1].ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token
		} else {
			resp := searchByCont(key, continuation)
			contentIndex = 0
			contents = resp.OnResponseReceivedCommands[0].AppendContinuationItemsAction.ContinuationItems[contentIndex].ItemSectionRenderer.Contents
			if len(contents) == 0 {
				contentIndex = 1
				contents = resp.OnResponseReceivedCommands[0].AppendContinuationItemsAction.ContinuationItems[contentIndex].ItemSectionRenderer.Contents
			}
			if len(resp.OnResponseReceivedCommands[0].AppendContinuationItemsAction.ContinuationItems) > 1 {
				continuation = resp.OnResponseReceivedCommands[0].AppendContinuationItemsAction.ContinuationItems[contentIndex^1].ContinuationItemRenderer.ContinuationEndpoint.ContinuationCommand.Token
			} else {
				return
			}
		}

		for _, content := range contents {
			if content.ChannelRenderer.ChannelID != "" {
			} else if content.VideoRenderer.VideoID != "" {
				//fmt.Printf("Video https://www.youtube.com/watch?v=%s\n", content.VideoRenderer.VideoID)
				ID := content.VideoRenderer.VideoID
				Title := content.VideoRenderer.Title.Runs[0].Text
				Channel_Name := content.VideoRenderer.LongBylineText.Runs[0].Text
				Channel_Url := content.VideoRenderer.LongBylineText.Runs[0].NavigationEndpoint.CommandMetadata.WebCommandMetadata.URL
				Published_Time := content.VideoRenderer.PublishedTimeText.SimpleText
				Long := content.VideoRenderer.LengthText.SimpleText
				View_Count := content.VideoRenderer.ViewCountText
				Browse_ID := content.VideoRenderer.LongBylineText.Runs[0].NavigationEndpoint.BrowseEndpoint.BrowseID
				if isContainsKeys(Title) && !Cache.SetCheck(ID) {
					fmt.Println(ID, Title, Channel_Name, Channel_Url, Published_Time, Long, View_Count.SimpleText, Browse_ID)
					appendMultiCSV(resultPath, [][]string{{ID, Title, Channel_Name, Channel_Url, Published_Time, Long, View_Count.SimpleText, Browse_ID}})
				}

			} else if len(content.ReelShelfRenderer.Items) > 0 {
			} else if len(content.ShelfRenderer.Content.VerticalListRenderer.Items) > 0 {
				for _, content := range content.ShelfRenderer.Content.VerticalListRenderer.Items {
					ID := content.VideoRenderer.VideoID
					Title := content.VideoRenderer.Title.Runs[0].Text
					Channel_Name := content.VideoRenderer.LongBylineText.Runs[0].Text
					Channel_Url := content.VideoRenderer.LongBylineText.Runs[0].NavigationEndpoint.CommandMetadata.WebCommandMetadata.URL
					Published_Time := content.VideoRenderer.PublishedTimeText.SimpleText
					Long := content.VideoRenderer.LengthText.SimpleText
					View_Count := content.VideoRenderer.ViewCountText
					Browse_ID := content.VideoRenderer.LongBylineText.Runs[0].NavigationEndpoint.BrowseEndpoint.BrowseID
					if isContainsKeys(Title) && !Cache.SetCheck(ID) {
						fmt.Println(ID, Title, Channel_Name, Channel_Url, Published_Time, Long, View_Count.SimpleText, Browse_ID)
						appendMultiCSV(resultPath, [][]string{{ID, Title, Channel_Name, Channel_Url, Published_Time, Long, View_Count.SimpleText, Browse_ID}})
					}
					//fmt.Printf("Video https://www.youtube.com/watch?v=%s\n", item.VideoRenderer.VideoID)
				}
			}
			//fmt.Println(content.ShelfRenderer.Content)
		}
		//fmt.Println(continuation)
	}

}

func searchByCont(key, continuation string) SearchContinuationStruct {
	for {

		data := fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"SA","remoteHost":"5.156.210.49","deviceMake":"","deviceModel":"","visitorData":"CgtoSmd5X1FKUk1YWSjKp9KrBjIICgJTQRICGgA%%3D","userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36,gzip(gfe)","clientName":"WEB","clientVersion":"2.20231208.01.00","osName":"Windows","osVersion":"10.0","originalUrl":"https://www.youtube.com/results?search_query=searchKey","screenPixelDensity":1,"platform":"DESKTOP","clientFormFactor":"UNKNOWN_FORM_FACTOR","configInfo":{"appInstallData":"CMqn0qsGEPX5rwUQvZmwBRCVz68FEMyu_hIQg4L_EhDT4a8FEKn3rwUQvPmvBRC2nbAFELfvrwUQ3ej-EhDcmbAFEL2KsAUQrtT-EhC9tq4FELiLrgUQsJ6wBRCmmrAFEKy3rwUQ_IWwBRDonbAFENShrwUQmZSwBRD9n7AFEKKSsAUQ5LP-EhClwv4SEKaBsAUQ7qKvBRDHg7AFEInorgUQyfevBRDnuq8FELfq_hIQzN-uBRDQjbAFEOLUrgUQ65OuBRComrAFENDirwUQlPr-EhDpjLAFEL75rwUQooGwBRDamLAFENnJrwUQ1YiwBRDNlbAFENfprwUQ4fKvBRCI468FEJaDsAUQnouwBRCa8K8FENyCsAUQ6-j-EhC7jrAFEOrDrwUQq4KwBRDbr68FEJj8_hIQiIewBRCZkbAFEL-dsAUQyfy3IhCKgP8SEI6F_xIQjpywBQ%%3D%%3D"},"mainAppWebInfo":{"graftUrl":"/results?search_query=searchKey","pwaInstallabilityStatus":"PWA_INSTALLABILITY_STATUS_UNKNOWN","webDisplayMode":"WEB_DISPLAY_MODE_BROWSER","isWebNativeShareAvailable":true},"timeZone":"Asia/Riyadh"},"user":{"lockedSafetyMode":false},"request":{"useSsl":true,"internalExperimentFlags":[],"consistencyTokenJars":[]},"clickTracking":{"clickTrackingParams":"CA0Q7VAiEwi7_Neg4YKDAxU-x0kHHb6TCIM="}},"query":"%s","continuation":"%s"}`, strings.ReplaceAll(key, "\"", "\\\""), continuation)
		req, err := http.NewRequest("POST", "https://www.youtube.com/youtubei/v1/search?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8&prettyPrint=false", strings.NewReader(data))
		if err == nil {
			req.Header.Set("Host", "www.youtube.com")
			req.Header.Set("accept", "*/*")
			req.Header.Set("accept-language", "en-US,en;q=0.9")
			req.Header.Set("content-type", "application/json")
			req.Header.Set("origin", "https://www.youtube.com")
			req.Header.Set("referer", "https://www.youtube.com/results?search_query=searchKey")
			req.Header.Set("sec-ch-ua", "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Google Chrome\";v=\"120\"")
			req.Header.Set("sec-ch-ua-arch", "\"x86\"")
			req.Header.Set("sec-ch-ua-bitness", "\"64\"")
			req.Header.Set("sec-ch-ua-full-version", "\"120.0.6099.71\"")
			req.Header.Set("sec-ch-ua-full-version-list", "\"Not_A Brand\";v=\"8.0.0.0\", \"Chromium\";v=\"120.0.6099.71\", \"Google Chrome\";v=\"120.0.6099.71\"")
			req.Header.Set("sec-ch-ua-mobile", "?0")
			req.Header.Set("sec-ch-ua-model", "\"\"")
			req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
			req.Header.Set("sec-ch-ua-platform-version", "\"10.0.0\"")
			req.Header.Set("sec-ch-ua-wow64", "?0")
			req.Header.Set("sec-fetch-dest", "empty")
			req.Header.Set("sec-fetch-mode", "same-origin")
			req.Header.Set("sec-fetch-site", "same-origin")
			req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
			req.Header.Set("x-goog-visitor-id", "CgtoSmd5X1FKUk1YWSjKp9KrBjIICgJTQRICGgA%3D")
			req.Header.Set("x-youtube-bootstrap-logged-in", "false")
			req.Header.Set("x-youtube-client-name", "1")
			req.Header.Set("x-youtube-client-version", "2.20231208.01.00")
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					a := SearchContinuationStruct{}
					err := json.Unmarshal(body, &a)
					if err == nil {
						return a
					}
					fmt.Println("Error when parsing json.Unmarshal(body, &a): ", err)
				}
			}
		}
	}
}

func search(key string) SearchStruct {
	for {
		data := fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"SA","remoteHost":"5.156.210.49","deviceMake":"","deviceModel":"","visitorData":"CgtoSmd5X1FKUk1YWSjKp9KrBjIICgJTQRICGgA%%3D","userAgent":"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36,gzip(gfe)","clientName":"WEB","clientVersion":"2.20231208.01.00","osName":"Windows","osVersion":"10.0","originalUrl":"https://www.youtube.com/results?search_query=searchKey","screenPixelDensity":1,"platform":"DESKTOP","clientFormFactor":"UNKNOWN_FORM_FACTOR","configInfo":{"appInstallData":"CMqn0qsGEPX5rwUQvZmwBRCVz68FEMyu_hIQg4L_EhDT4a8FEKn3rwUQvPmvBRC2nbAFELfvrwUQ3ej-EhDcmbAFEL2KsAUQrtT-EhC9tq4FELiLrgUQsJ6wBRCmmrAFEKy3rwUQ_IWwBRDonbAFENShrwUQmZSwBRD9n7AFEKKSsAUQ5LP-EhClwv4SEKaBsAUQ7qKvBRDHg7AFEInorgUQyfevBRDnuq8FELfq_hIQzN-uBRDQjbAFEOLUrgUQ65OuBRComrAFENDirwUQlPr-EhDpjLAFEL75rwUQooGwBRDamLAFENnJrwUQ1YiwBRDNlbAFENfprwUQ4fKvBRCI468FEJaDsAUQnouwBRCa8K8FENyCsAUQ6-j-EhC7jrAFEOrDrwUQq4KwBRDbr68FEJj8_hIQiIewBRCZkbAFEL-dsAUQyfy3IhCKgP8SEI6F_xIQjpywBQ%%3D%%3D"},"mainAppWebInfo":{"graftUrl":"/results?search_query=searchKey","pwaInstallabilityStatus":"PWA_INSTALLABILITY_STATUS_UNKNOWN","webDisplayMode":"WEB_DISPLAY_MODE_BROWSER","isWebNativeShareAvailable":true},"timeZone":"Asia/Riyadh"},"user":{"lockedSafetyMode":false},"request":{"useSsl":true,"internalExperimentFlags":[],"consistencyTokenJars":[]},"clickTracking":{"clickTrackingParams":"CA0Q7VAiEwi7_Neg4YKDAxU-x0kHHb6TCIM="}},"query":"%s"}`, strings.ReplaceAll(key, "\"", "\\\""))
		req, err := http.NewRequest("POST", "https://www.youtube.com/youtubei/v1/search?key=AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8&prettyPrint=false", strings.NewReader(data))
		if err == nil {
			req.Header.Set("Host", "www.youtube.com")
			req.Header.Set("accept", "*/*")
			req.Header.Set("accept-language", "en-US,en;q=0.9")
			req.Header.Set("content-type", "application/json")
			req.Header.Set("origin", "https://www.youtube.com")
			req.Header.Set("referer", "https://www.youtube.com/results?search_query=searchKey")
			req.Header.Set("sec-ch-ua", "\"Not_A Brand\";v=\"8\", \"Chromium\";v=\"120\", \"Google Chrome\";v=\"120\"")
			req.Header.Set("sec-ch-ua-arch", "\"x86\"")
			req.Header.Set("sec-ch-ua-bitness", "\"64\"")
			req.Header.Set("sec-ch-ua-full-version", "\"120.0.6099.71\"")
			req.Header.Set("sec-ch-ua-full-version-list", "\"Not_A Brand\";v=\"8.0.0.0\", \"Chromium\";v=\"120.0.6099.71\", \"Google Chrome\";v=\"120.0.6099.71\"")
			req.Header.Set("sec-ch-ua-mobile", "?0")
			req.Header.Set("sec-ch-ua-model", "\"\"")
			req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
			req.Header.Set("sec-ch-ua-platform-version", "\"10.0.0\"")
			req.Header.Set("sec-ch-ua-wow64", "?0")
			req.Header.Set("sec-fetch-dest", "empty")
			req.Header.Set("sec-fetch-mode", "same-origin")
			req.Header.Set("sec-fetch-site", "same-origin")
			req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
			req.Header.Set("x-goog-visitor-id", "CgtoSmd5X1FKUk1YWSjKp9KrBjIICgJTQRICGgA%3D")
			req.Header.Set("x-youtube-bootstrap-logged-in", "false")
			req.Header.Set("x-youtube-client-name", "1")
			req.Header.Set("x-youtube-client-version", "2.20231208.01.00")
			resp, err := http.DefaultClient.Do(req)
			if err == nil {
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil {
					a := SearchStruct{}
					err := json.Unmarshal(body, &a)
					if err == nil {
						return a
					}
					fmt.Println("Error when parsing json.Unmarshal(body, &a): ", err)
				}
			}
		}
	}
}
