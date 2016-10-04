package gobay
import "encoding/xml"
type GetMyeBaySellingResult struct {
	Ack string
	Build string
	CorrelationID string
	HardExpirationWarning string
	Timestamp string
	Version string
	ActiveList []struct {
		ItemArray []struct {
			Item []struct {
				BuyItNowPrice string
				ClassifiedAdPayPerLeadFee string
				eBayNotes string
				HideFromSearch string
				ID string
				LeadCount string
				ListingDuration string
				ListingType string
				NewLeadCount string
				PrivateNotes string
				Quantity string
				QuantityAvailable string
				QuestionCount string
				ReasonHideFromSearch string
				ReservePrice string
				SKU string
				StartPrice string
				TimeLeft string
				Title string
				WatchCount string
				BestOfferDetails []struct {
					BestOfferCount string
				}

				ListingDetails []struct {
					ConvertedBuyItNowPrice string
					ConvertedReservePrice string
					ConvertedStartPrice string
					StartTime string
				}

				PictureDetails []struct {
					GalleryURL string
				}

				SellingStatus []struct {
					BidCount string
					BidderCount string
					ConvertedCurrentPrice string
					CurrentPrice string
					QuantitySold string
					ReserveMet string
					HighBidder []struct {
						FeedbackRatingStar string
						FeedbackScore string
						UserID string
					}

					PromotionalSaleDetails []struct {
						EndTime string
						OriginalPrice string
						StartTime string
					}

				}

				ShippingDetails []struct {
					GlobalShipping string
					ShippingType string
					ShippingServiceOptions []struct {
						LocalPickup string
						LogisticPlanType string
						ShippingServiceCost string
					}

				}

				Variations []struct {
					Variation []struct {
						PrivateNotes string
						Quantity string
						SKU string
						StartPrice string
						Title string
						WatchCount string
						SellingStatus []struct {
							BidCount string
							BidderCount string
							ConvertedCurrentPrice string
							CurrentPrice string
							QuantitySold string
							ReserveMet string
							HighBidder []struct {
								FeedbackRatingStar string
								FeedbackScore string
								UserID string
							}

							PromotionalSaleDetails []struct {
								EndTime string
								OriginalPrice string
								StartTime string
							}

						}

						Specifics []struct {
							NameValueList []struct {
								Name string
								Value string
							}

						}

					}

				}

			}

		}

		PaginationResult []struct {
			TotalNumberOfEntries string
			TotalNumberOfPages string
		}

	}

	DeletedFromSoldList []struct {
		OrderTransactionArray []struct {
			OrderTransaction []struct {
				Order []struct {
					ID string
					Subtotal string
					TransactionArray []struct {
						Transaction []struct {
							ConvertedPrice string
							CreatedDate string
							IsMultiLegShipping string
							LineItemID string
							PaidTime string
							PaisaPayID string
							Platform string
							QuantityPurchased string
							SellerPaidStatus string
							ShippedTime string
							ID string
							Price string
							Buyer []struct {
								Email string
								StaticAlias string
								UserID string
								Info []struct {
									ShippingAddress []struct {
										PostalCode string
									}

								}

							}

							FeedbackLeft []struct {
								CommentType string
							}

							FeedbackReceived []struct {
								CommentType string
							}

							Item []struct {
								BuyItNowPrice string
								ClassifiedAdPayPerLeadFee string
								HideFromSearch string
								ID string
								ListingType string
								PrivateNotes string
								Quantity string
								QuantityAvailable string
								QuestionCount string
								ReasonHideFromSearch string
								ReservePrice string
								SKU string
								StartPrice string
								TimeLeft string
								Title string
								WatchCount string
								ListingDetails []struct {
									ConvertedBuyItNowPrice string
									ConvertedReservePrice string
									ConvertedStartPrice string
									EndTime string
									StartTime string
								}

								PictureDetails []struct {
									GalleryURL string
								}

								SellingStatus []struct {
									ConvertedCurrentPrice string
									CurrentPrice string
									QuantitySold string
								}

								ShippingDetails []struct {
									GlobalShipping string
									ShippingType string
									ShippingServiceOptions []struct {
										LocalPickup string
										LogisticPlanType string
										ShippingServiceCost string
									}

								}

								Variations []struct {
									Variation []struct {
										Quantity string
										SKU string
										StartPrice string
										Title string
										WatchCount string
										SellingStatus []struct {
											ConvertedCurrentPrice string
											CurrentPrice string
											QuantitySold string
										}

										Specifics []struct {
											NameValueList []struct {
												Name string
												Value string
											}

										}

									}

								}

							}

							Status []struct {
								PaymentHold string
							}

						}

					}

				}

				Transaction []struct {
					ConvertedPrice string
					CreatedDate string
					IsMultiLegShipping string
					OrderLineItemID string
					PaidTime string
					PaisaPayID string
					Platform string
					QuantityPurchased string
					SellerPaidStatus string
					ShippedTime string
					ID string
					Price string
					Buyer []struct {
						Email string
						StaticAlias string
						UserID string
						Info []struct {
							ShippingAddress []struct {
								PostalCode string
							}

						}

					}

					FeedbackLeft []struct {
						CommentType string
					}

					FeedbackReceived []struct {
						CommentType string
					}

					Item []struct {
						BuyItNowPrice string
						ClassifiedAdPayPerLeadFee string
						HideFromSearch string
						ID string
						ListingType string
						PrivateNotes string
						Quantity string
						QuantityAvailable string
						QuestionCount string
						ReasonHideFromSearch string
						ReservePrice string
						SKU string
						StartPrice string
						TimeLeft string
						Title string
						WatchCount string
						ListingDetails []struct {
							ConvertedBuyItNowPrice string
							ConvertedReservePrice string
							ConvertedStartPrice string
							EndTime string
							StartTime string
						}

						PictureDetails []struct {
							GalleryURL string
						}

						SellingStatus []struct {
							ConvertedCurrentPrice string
							CurrentPrice string
							QuantitySold string
						}

						ShippingDetails []struct {
							GlobalShipping string
							ShippingType string
							ShippingServiceOptions []struct {
								LocalPickup string
								LogisticPlanType string
								ShippingServiceCost string
							}

						}

						Variations []struct {
							Variation []struct {
								Quantity string
								SKU string
								StartPrice string
								Title string
								WatchCount string
								SellingStatus []struct {
									ConvertedCurrentPrice string
									CurrentPrice string
									QuantitySold string
								}

								Specifics []struct {
									NameValueList []struct {
										Name string
										Value string
									}

								}

							}

						}

					}

					Status []struct {
						PaymentHold string
					}

				}

			}

		}

		PaginationResult []struct {
			TotalNumberOfEntries string
			TotalNumberOfPages string
		}

	}

	DeletedFromUnsoldList []struct {
		ItemArray []struct {
			Item []struct {
				BuyItNowPrice string
				ClassifiedAdPayPerLeadFee string
				HideFromSearch string
				ID string
				ListingType string
				PrivateNotes string
				Quantity string
				QuantityAvailable string
				QuestionCount string
				ReasonHideFromSearch string
				ReservePrice string
				SKU string
				StartPrice string
				TimeLeft string
				Title string
				WatchCount string
				ListingDetails []struct {
					ConvertedBuyItNowPrice string
					ConvertedReservePrice string
					ConvertedStartPrice string
					EndTime string
					StartTime string
				}

				PictureDetails []struct {
					GalleryURL string
				}

				SellingStatus []struct {
					ConvertedCurrentPrice string
					CurrentPrice string
					QuantitySold string
				}

				ShippingDetails []struct {
					GlobalShipping string
					ShippingType string
					ShippingServiceOptions []struct {
						LocalPickup string
						LogisticPlanType string
						ShippingServiceCost string
					}

				}

				Variations []struct {
					Variation []struct {
						PrivateNotes string
						Quantity string
						SKU string
						StartPrice string
						Title string
						WatchCount string
						SellingStatus []struct {
							ConvertedCurrentPrice string
							CurrentPrice string
							QuantitySold string
						}

						Specifics []struct {
							NameValueList []struct {
								Name string
								Value string
							}

						}

					}

				}

			}

		}

		PaginationResult []struct {
			TotalNumberOfEntries string
			TotalNumberOfPages string
		}

	}

	ScheduledList []struct {
		ItemArray []struct {
			Item []struct {
				BuyItNowPrice string
				ClassifiedAdPayPerLeadFee string
				eBayNotes string
				HideFromSearch string
				ID string
				ListingDuration string
				ListingType string
				PrivateNotes string
				Quantity string
				QuantityAvailable string
				QuestionCount string
				ReasonHideFromSearch string
				ReservePrice string
				SKU string
				StartPrice string
				TimeLeft string
				Title string
				WatchCount string
				ListingDetails []struct {
					ConvertedBuyItNowPrice string
					ConvertedReservePrice string
					ConvertedStartPrice string
					StartTime string
				}

				PictureDetails []struct {
					GalleryURL string
				}

				SellingStatus []struct {
					ConvertedCurrentPrice string
					CurrentPrice string
					QuantitySold string
					ReserveMet string
				}

				ShippingDetails []struct {
					GlobalShipping string
					ShippingType string
					ShippingServiceOptions []struct {
						LocalPickup string
						LogisticPlanType string
						ShippingServiceCost string
						ShippingSurcharge string
					}

				}

				Variations []struct {
					Variation []struct {
						PrivateNotes string
						Quantity string
						SKU string
						StartPrice string
						Title string
						WatchCount string
						SellingStatus []struct {
							ConvertedCurrentPrice string
							CurrentPrice string
							QuantitySold string
							ReserveMet string
						}

						Specifics []struct {
							NameValueList []struct {
								Name string
								Value string
							}

						}

					}

				}

			}

		}

		PaginationResult []struct {
			TotalNumberOfEntries string
			TotalNumberOfPages string
		}

	}

	SellingSummary []struct {
		ActiveAuctionCount string
		AuctionBidCount string
		AuctionSellingCount string
		SoldDurationInDays string
		TotalAuctionSellingValue string
		TotalSoldCount string
		TotalSoldValue string
	}

	SoldList []struct {
		OrderTransactionArray []struct {
			OrderTransaction []struct {
				Order []struct {
					ID string
					RefundAmount string
					RefundStatus string
					Subtotal string
					TransactionArray []struct {
						Transaction []struct {
							ConvertedPrice string
							CreatedDate string
							IsMultiLegShipping string
							LineItemID string
							PaidTime string
							PaisaPayID string
							Platform string
							QuantityPurchased string
							SellerPaidStatus string
							ShippedTime string
							TotalPrice string
							ID string
							Price string
							Buyer []struct {
								Email string
								StaticAlias string
								UserID string
								Info []struct {
									ShippingAddress []struct {
										PostalCode string
									}

								}

							}

							FeedbackLeft []struct {
								CommentType string
							}

							FeedbackReceived []struct {
								CommentType string
							}

							Item []struct {
								BuyItNowPrice string
								ClassifiedAdPayPerLeadFee string
								HideFromSearch string
								ID string
								ListingType string
								PrivateNotes string
								Quantity string
								QuantityAvailable string
								QuestionCount string
								ReasonHideFromSearch string
								ReservePrice string
								SKU string
								StartPrice string
								TimeLeft string
								Title string
								WatchCount string
								ListingDetails []struct {
									ConvertedBuyItNowPrice string
									ConvertedReservePrice string
									ConvertedStartPrice string
									EndTime string
									StartTime string
								}

								PictureDetails []struct {
									GalleryURL string
								}

								SellingStatus []struct {
									ConvertedCurrentPrice string
									CurrentPrice string
									QuantitySold string
									ReserveMet string
								}

								ShippingDetails []struct {
									GlobalShipping string
									ShippingType string
									ShippingServiceOptions []struct {
										LocalPickup string
										LogisticPlanType string
										ShippingServiceCost string
									}

								}

								Variations []struct {
									Variation []struct {
										Quantity string
										SKU string
										StartPrice string
										Title string
										WatchCount string
										SellingStatus []struct {
											ConvertedCurrentPrice string
											CurrentPrice string
											QuantitySold string
											ReserveMet string
										}

										Specifics []struct {
											NameValueList []struct {
												Name string
												Value string
											}

										}

									}

								}

							}

							PaymentHoldDetails []struct {
								ExpectedReleaseDate string
								PaymentHoldReason string
								RequiredSellerActionArray []struct {
									RequiredSellerAction string
								}

							}

							Status []struct {
								PaymentHold string
							}

						}

					}

				}

				Transaction []struct {
					ConvertedPrice string
					CreatedDate string
					IsMultiLegShipping string
					OrderLineItemID string
					PaidTime string
					PaisaPayID string
					Platform string
					QuantityPurchased string
					SellerPaidStatus string
					ShippedTime string
					TotalPrice string
					ID string
					Price string
					Buyer []struct {
						Email string
						StaticAlias string
						UserID string
						Info []struct {
							ShippingAddress []struct {
								PostalCode string
							}

						}

					}

					FeedbackLeft []struct {
						CommentType string
					}

					FeedbackReceived []struct {
						CommentType string
					}

					Item []struct {
						BuyItNowPrice string
						ClassifiedAdPayPerLeadFee string
						HideFromSearch string
						ID string
						ListingType string
						PrivateNotes string
						Quantity string
						QuantityAvailable string
						QuestionCount string
						ReasonHideFromSearch string
						ReservePrice string
						SKU string
						StartPrice string
						TimeLeft string
						Title string
						WatchCount string
						ListingDetails []struct {
							ConvertedBuyItNowPrice string
							ConvertedReservePrice string
							ConvertedStartPrice string
							EndTime string
							StartTime string
						}

						PictureDetails []struct {
							GalleryURL string
						}

						SellingStatus []struct {
							ConvertedCurrentPrice string
							CurrentPrice string
							QuantitySold string
							ReserveMet string
						}

						ShippingDetails []struct {
							GlobalShipping string
							ShippingType string
							ShippingServiceOptions []struct {
								LocalPickup string
								LogisticPlanType string
								ShippingServiceCost string
							}

						}

						Variations []struct {
							Variation []struct {
								Quantity string
								SKU string
								StartPrice string
								Title string
								WatchCount string
								SellingStatus []struct {
									ConvertedCurrentPrice string
									CurrentPrice string
									QuantitySold string
									ReserveMet string
								}

								Specifics []struct {
									NameValueList []struct {
										Name string
										Value string
									}

								}

							}

						}

					}

					PaymentHoldDetails []struct {
						ExpectedReleaseDate string
						PaymentHoldReason string
						RequiredSellerActionArray []struct {
							RequiredSellerAction string
						}

					}

					Status []struct {
						PaymentHold string
					}

				}

			}

		}

		PaginationResult []struct {
			TotalNumberOfEntries string
			TotalNumberOfPages string
		}

	}

	Summary []struct {
		ActiveAuctionCount string
		AmountLimitRemaining string
		AuctionBidCount string
		AuctionSellingCount string
		ClassifiedAdCount string
		ClassifiedAdOfferCount string
		QuantityLimitRemaining string
		SoldDurationInDays string
		TotalAuctionSellingValue string
		TotalLeadCount string
		TotalListingsWithLeads string
		TotalSoldCount string
		TotalSoldValue string
	}

	UnsoldList []struct {
		ItemArray []struct {
			Item []struct {
				BuyItNowPrice string
				ClassifiedAdPayPerLeadFee string
				eBayNotes string
				HideFromSearch string
				ID string
				LeadCount string
				ListingDuration string
				ListingType string
				PrivateNotes string
				Quantity string
				QuantityAvailable string
				QuestionCount string
				ReasonHideFromSearch string
				Relisted string
				ReservePrice string
				SKU string
				StartPrice string
				TimeLeft string
				Title string
				WatchCount string
				ListingDetails []struct {
					ConvertedBuyItNowPrice string
					ConvertedReservePrice string
					ConvertedStartPrice string
					EndTime string
					StartTime string
				}

				PictureDetails []struct {
					GalleryURL string
				}

				SellingStatus []struct {
					BidCount string
					BidderCount string
					ConvertedCurrentPrice string
					CurrentPrice string
					QuantitySold string
					ReserveMet string
				}

				ShippingDetails []struct {
					GlobalShipping string
					ShippingType string
					ShippingServiceOptions []struct {
						LocalPickup string
						LogisticPlanType string
						ShippingServiceCost string
					}

				}

				Variations []struct {
					Variation []struct {
						PrivateNotes string
						Quantity string
						SKU string
						StartPrice string
						Title string
						WatchCount string
						SellingStatus []struct {
							BidCount string
							BidderCount string
							ConvertedCurrentPrice string
							CurrentPrice string
							QuantitySold string
							ReserveMet string
						}

						Specifics []struct {
							NameValueList []struct {
								Name string
								Value string
							}

						}

					}

				}

			}

		}

		PaginationResult []struct {
			TotalNumberOfEntries string
			TotalNumberOfPages string
		}

	}

	Errors []struct {
		ErrorClassification string
		ErrorCode string
		LongMessage string
		SeverityCode string
		ShortMessage string
		ErrorParameters []struct {
			Value string
		}

	}

}
func NewGetMyeBaySellingResult(data []byte) *GetMyeBaySellingResult {
    var o GetMyeBaySellingResult
    xml.Unmarshal(data, &o)
    return &o
}
