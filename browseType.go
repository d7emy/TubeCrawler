package main

type browseContent struct {
	RichItemRenderer struct {
		Content struct {
			VideoRenderer struct {
				VideoID   string `json:"videoId"`
				Thumbnail struct {
					Thumbnails []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"thumbnails"`
				} `json:"thumbnail"`
				Title struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
					Accessibility struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
				} `json:"title"`
				DescriptionSnippet struct {
					Runs []struct {
						Text string `json:"text"`
					} `json:"runs"`
				} `json:"descriptionSnippet"`
				PublishedTimeText struct {
					SimpleText string `json:"simpleText"`
				} `json:"publishedTimeText"`
				LengthText struct {
					Accessibility struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
					SimpleText string `json:"simpleText"`
				} `json:"lengthText"`
				ViewCountText struct {
					SimpleText string `json:"simpleText"`
				} `json:"viewCountText"`
				NavigationEndpoint struct {
					ClickTrackingParams string `json:"clickTrackingParams"`
					CommandMetadata     struct {
						WebCommandMetadata struct {
							URL         string `json:"url"`
							WebPageType string `json:"webPageType"`
							RootVe      int    `json:"rootVe"`
						} `json:"webCommandMetadata"`
					} `json:"commandMetadata"`
					WatchEndpoint struct {
						VideoID                            string `json:"videoId"`
						WatchEndpointSupportedOnesieConfig struct {
							HTML5PlaybackOnesieConfig struct {
								CommonConfig struct {
									URL string `json:"url"`
								} `json:"commonConfig"`
							} `json:"html5PlaybackOnesieConfig"`
						} `json:"watchEndpointSupportedOnesieConfig"`
					} `json:"watchEndpoint"`
				} `json:"navigationEndpoint"`
				OwnerBadges []struct {
					MetadataBadgeRenderer struct {
						Icon struct {
							IconType string `json:"iconType"`
						} `json:"icon"`
						Style             string `json:"style"`
						Tooltip           string `json:"tooltip"`
						TrackingParams    string `json:"trackingParams"`
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"metadataBadgeRenderer"`
				} `json:"ownerBadges"`
				TrackingParams     string `json:"trackingParams"`
				ShowActionMenu     bool   `json:"showActionMenu"`
				ShortViewCountText struct {
					Accessibility struct {
						AccessibilityData struct {
							Label string `json:"label"`
						} `json:"accessibilityData"`
					} `json:"accessibility"`
					SimpleText string `json:"simpleText"`
				} `json:"shortViewCountText"`
				Menu struct {
					MenuRenderer struct {
						Items []struct {
							MenuServiceItemRenderer struct {
								Text struct {
									Runs []struct {
										Text string `json:"text"`
									} `json:"runs"`
								} `json:"text"`
								Icon struct {
									IconType string `json:"iconType"`
								} `json:"icon"`
								ServiceEndpoint struct {
									ClickTrackingParams string `json:"clickTrackingParams"`
									CommandMetadata     struct {
										WebCommandMetadata struct {
											SendPost bool `json:"sendPost"`
										} `json:"webCommandMetadata"`
									} `json:"commandMetadata"`
									SignalServiceEndpoint struct {
										Signal  string `json:"signal"`
										Actions []struct {
											ClickTrackingParams  string `json:"clickTrackingParams"`
											AddToPlaylistCommand struct {
												OpenMiniplayer      bool   `json:"openMiniplayer"`
												VideoID             string `json:"videoId"`
												ListType            string `json:"listType"`
												OnCreateListCommand struct {
													ClickTrackingParams string `json:"clickTrackingParams"`
													CommandMetadata     struct {
														WebCommandMetadata struct {
															SendPost bool   `json:"sendPost"`
															APIURL   string `json:"apiUrl"`
														} `json:"webCommandMetadata"`
													} `json:"commandMetadata"`
													CreatePlaylistServiceEndpoint struct {
														VideoIds []string `json:"videoIds"`
														Params   string   `json:"params"`
													} `json:"createPlaylistServiceEndpoint"`
												} `json:"onCreateListCommand"`
												VideoIds []string `json:"videoIds"`
											} `json:"addToPlaylistCommand"`
										} `json:"actions"`
									} `json:"signalServiceEndpoint"`
								} `json:"serviceEndpoint"`
								TrackingParams string `json:"trackingParams"`
							} `json:"menuServiceItemRenderer,omitempty"`
							MenuServiceItemDownloadRenderer struct {
								ServiceEndpoint struct {
									ClickTrackingParams  string `json:"clickTrackingParams"`
									OfflineVideoEndpoint struct {
										VideoID      string `json:"videoId"`
										OnAddCommand struct {
											ClickTrackingParams      string `json:"clickTrackingParams"`
											GetDownloadActionCommand struct {
												VideoID string `json:"videoId"`
												Params  string `json:"params"`
											} `json:"getDownloadActionCommand"`
										} `json:"onAddCommand"`
									} `json:"offlineVideoEndpoint"`
								} `json:"serviceEndpoint"`
								TrackingParams string `json:"trackingParams"`
							} `json:"menuServiceItemDownloadRenderer,omitempty"`
						} `json:"items"`
						TrackingParams string `json:"trackingParams"`
						Accessibility  struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"accessibility"`
					} `json:"menuRenderer"`
				} `json:"menu"`
				ThumbnailOverlays []struct {
					ThumbnailOverlayTimeStatusRenderer struct {
						Text struct {
							Accessibility struct {
								AccessibilityData struct {
									Label string `json:"label"`
								} `json:"accessibilityData"`
							} `json:"accessibility"`
							SimpleText string `json:"simpleText"`
						} `json:"text"`
						Style string `json:"style"`
					} `json:"thumbnailOverlayTimeStatusRenderer,omitempty"`

					ThumbnailOverlayToggleButtonRenderer struct {
						UntoggledIcon struct {
							IconType string `json:"iconType"`
						} `json:"untoggledIcon"`
						ToggledIcon struct {
							IconType string `json:"iconType"`
						} `json:"toggledIcon"`
						UntoggledTooltip         string `json:"untoggledTooltip"`
						ToggledTooltip           string `json:"toggledTooltip"`
						UntoggledServiceEndpoint struct {
							ClickTrackingParams string `json:"clickTrackingParams"`
							CommandMetadata     struct {
								WebCommandMetadata struct {
									SendPost bool `json:"sendPost"`
								} `json:"webCommandMetadata"`
							} `json:"commandMetadata"`
							SignalServiceEndpoint struct {
								Signal  string `json:"signal"`
								Actions []struct {
									ClickTrackingParams  string `json:"clickTrackingParams"`
									AddToPlaylistCommand struct {
										OpenMiniplayer      bool   `json:"openMiniplayer"`
										VideoID             string `json:"videoId"`
										ListType            string `json:"listType"`
										OnCreateListCommand struct {
											ClickTrackingParams string `json:"clickTrackingParams"`
											CommandMetadata     struct {
												WebCommandMetadata struct {
													SendPost bool   `json:"sendPost"`
													APIURL   string `json:"apiUrl"`
												} `json:"webCommandMetadata"`
											} `json:"commandMetadata"`
											CreatePlaylistServiceEndpoint struct {
												VideoIds []string `json:"videoIds"`
												Params   string   `json:"params"`
											} `json:"createPlaylistServiceEndpoint"`
										} `json:"onCreateListCommand"`
										VideoIds []string `json:"videoIds"`
									} `json:"addToPlaylistCommand"`
								} `json:"actions"`
							} `json:"signalServiceEndpoint"`
						} `json:"untoggledServiceEndpoint"`
						UntoggledAccessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"untoggledAccessibility"`
						ToggledAccessibility struct {
							AccessibilityData struct {
								Label string `json:"label"`
							} `json:"accessibilityData"`
						} `json:"toggledAccessibility"`
						TrackingParams string `json:"trackingParams"`
					} `json:"thumbnailOverlayToggleButtonRenderer,omitempty"`
					ThumbnailOverlayNowPlayingRenderer struct {
						Text struct {
							Runs []struct {
								Text string `json:"text"`
							} `json:"runs"`
						} `json:"text"`
					} `json:"thumbnailOverlayNowPlayingRenderer,omitempty"`
				} `json:"thumbnailOverlays"`
				RichThumbnail struct {
					MovingThumbnailRenderer struct {
						MovingThumbnailDetails struct {
							Thumbnails []struct {
								URL    string `json:"url"`
								Width  int    `json:"width"`
								Height int    `json:"height"`
							} `json:"thumbnails"`
							LogAsMovingThumbnail bool `json:"logAsMovingThumbnail"`
						} `json:"movingThumbnailDetails"`
						EnableHoveredLogging bool `json:"enableHoveredLogging"`
						EnableOverlay        bool `json:"enableOverlay"`
					} `json:"movingThumbnailRenderer"`
				} `json:"richThumbnail"`
			} `json:"videoRenderer"`
		} `json:"content"`
		TrackingParams string `json:"trackingParams"`
	} `json:"richItemRenderer"`
	ContinuationItemRenderer struct {
		Trigger              string `json:"trigger"`
		ContinuationEndpoint struct {
			ClickTrackingParams string `json:"clickTrackingParams"`
			CommandMetadata     struct {
				WebCommandMetadata struct {
					SendPost bool   `json:"sendPost"`
					APIURL   string `json:"apiUrl"`
				} `json:"webCommandMetadata"`
			} `json:"commandMetadata"`
			ContinuationCommand struct {
				Token   string `json:"token"`
				Request string `json:"request"`
			} `json:"continuationCommand"`
		} `json:"continuationEndpoint"`
	} `json:"continuationItemRenderer"`
}
type browseStruct struct {
	Contents struct {
		TwoColumnBrowseResultsRenderer struct {
			Tabs []struct {
				TabRenderer struct {
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								APIURL      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseID         string `json:"browseId"`
							Params           string `json:"params"`
							CanonicalBaseURL string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					Title    string `json:"title"`
					Selected bool   `json:"selected"`
					Content  struct {
						RichGridRenderer struct {
							Contents       []browseContent `json:"contents"`
							TrackingParams string          `json:"trackingParams"`
							Header         struct {
								FeedFilterChipBarRenderer struct {
									Contents []struct {
										ChipCloudChipRenderer struct {
											Text struct {
												SimpleText string `json:"simpleText"`
											} `json:"text"`
											NavigationEndpoint struct {
												ClickTrackingParams string `json:"clickTrackingParams"`
												CommandMetadata     struct {
													WebCommandMetadata struct {
														SendPost bool   `json:"sendPost"`
														APIURL   string `json:"apiUrl"`
													} `json:"webCommandMetadata"`
												} `json:"commandMetadata"`
												ContinuationCommand struct {
													Token   string `json:"token"`
													Request string `json:"request"`
													Command struct {
														ClickTrackingParams string `json:"clickTrackingParams"`
														ShowReloadUICommand struct {
															TargetID string `json:"targetId"`
														} `json:"showReloadUiCommand"`
													} `json:"command"`
												} `json:"continuationCommand"`
											} `json:"navigationEndpoint"`
											TrackingParams string `json:"trackingParams"`
											IsSelected     bool   `json:"isSelected"`
										} `json:"chipCloudChipRenderer"`
									} `json:"contents"`
									TrackingParams string `json:"trackingParams"`
									NextButton     struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											Size       string `json:"size"`
											IsDisabled bool   `json:"isDisabled"`
											Icon       struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Tooltip           string `json:"tooltip"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibilityData"`
										} `json:"buttonRenderer"`
									} `json:"nextButton"`
									PreviousButton struct {
										ButtonRenderer struct {
											Style      string `json:"style"`
											Size       string `json:"size"`
											IsDisabled bool   `json:"isDisabled"`
											Icon       struct {
												IconType string `json:"iconType"`
											} `json:"icon"`
											Tooltip           string `json:"tooltip"`
											TrackingParams    string `json:"trackingParams"`
											AccessibilityData struct {
												AccessibilityData struct {
													Label string `json:"label"`
												} `json:"accessibilityData"`
											} `json:"accessibilityData"`
										} `json:"buttonRenderer"`
									} `json:"previousButton"`
									StyleType string `json:"styleType"`
								} `json:"feedFilterChipBarRenderer"`
							} `json:"header"`
							TargetID string `json:"targetId"`
							Style    string `json:"style"`
						} `json:"richGridRenderer"`
					} `json:"content"`
					TrackingParams string `json:"trackingParams"`
				} `json:"tabRenderer,omitempty"`
				ExpandableTabRenderer struct {
					Endpoint struct {
						ClickTrackingParams string `json:"clickTrackingParams"`
						CommandMetadata     struct {
							WebCommandMetadata struct {
								URL         string `json:"url"`
								WebPageType string `json:"webPageType"`
								RootVe      int    `json:"rootVe"`
								APIURL      string `json:"apiUrl"`
							} `json:"webCommandMetadata"`
						} `json:"commandMetadata"`
						BrowseEndpoint struct {
							BrowseID         string `json:"browseId"`
							Params           string `json:"params"`
							CanonicalBaseURL string `json:"canonicalBaseUrl"`
						} `json:"browseEndpoint"`
					} `json:"endpoint"`
					Title    string `json:"title"`
					Selected bool   `json:"selected"`
				} `json:"expandableTabRenderer,omitempty"`
			} `json:"tabs"`
		} `json:"twoColumnBrowseResultsRenderer"`
	} `json:"contents"`
}
type browseContinuationStruct struct {
		ResponseContext struct {
			VisitorData           string `json:"visitorData"`
			ServiceTrackingParams []struct {
				Service string `json:"service"`
				Params  []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"params"`
			} `json:"serviceTrackingParams"`
			MaxAgeSeconds             int `json:"maxAgeSeconds"`
			MainAppWebResponseContext struct {
				LoggedOut     bool   `json:"loggedOut"`
				TrackingParam string `json:"trackingParam"`
			} `json:"mainAppWebResponseContext"`
			WebResponseContextExtensionData struct {
				HasDecorated bool `json:"hasDecorated"`
			} `json:"webResponseContextExtensionData"`
		} `json:"responseContext"`
		TrackingParams            string `json:"trackingParams"`
		OnResponseReceivedActions []struct {
			ClickTrackingParams           string `json:"clickTrackingParams"`
			AppendContinuationItemsAction struct {
				ContinuationItems []browseContent`json:"continuationItems"`
				TargetID string `json:"targetId"`
			} `json:"appendContinuationItemsAction"`
		} `json:"onResponseReceivedActions"`
	}