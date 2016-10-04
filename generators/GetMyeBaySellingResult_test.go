package gobay
import "testing"
func TestGetMyeBaySellingResult(t *testing.T) {
	data := `<?xml version="1.0" encoding="utf-8"?>
<GetMyeBaySellingResponse xmlns="urn:ebay:apis:eBLBaseComponents">
  <!-- Call-specific Output Fields -->
  <ActiveList> <!-- PaginatedItemArrayType -->
    <ItemArray> <!-- ItemArrayType -->
      <Item> <!-- ItemType -->
        <BestOfferDetails> <!-- BestOfferDetailsType -->
          <BestOfferCount> int </BestOfferCount>
        </BestOfferDetails>
        <BuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </BuyItNowPrice>
        <ClassifiedAdPayPerLeadFee currencyID="CurrencyCodeType"> AmountType (double) </ClassifiedAdPayPerLeadFee>
        <eBayNotes> string </eBayNotes>
        <HideFromSearch> boolean </HideFromSearch>
        <ItemID> ItemIDType (string) </ItemID>
        <LeadCount> int </LeadCount>
        <ListingDetails> <!-- ListingDetailsType -->
          <ConvertedBuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedBuyItNowPrice>
          <ConvertedReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedReservePrice>
          <ConvertedStartPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedStartPrice>
          <StartTime> dateTime </StartTime>
        </ListingDetails>
        <ListingDuration> token </ListingDuration>
        <ListingType> ListingTypeCodeType </ListingType>
        <NewLeadCount> int </NewLeadCount>
        <PictureDetails> <!-- PictureDetailsType -->
          <GalleryURL> anyURI </GalleryURL>
        </PictureDetails>
        <PrivateNotes> string </PrivateNotes>
        <Quantity> int </Quantity>
        <QuantityAvailable> int </QuantityAvailable>
        <QuestionCount> long </QuestionCount>
        <ReasonHideFromSearch> ReasonHideFromSearchCodeType </ReasonHideFromSearch>
        <ReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ReservePrice>
        <SellingStatus> <!-- SellingStatusType -->
          <BidCount> int </BidCount>
          <BidderCount> long </BidderCount>
          <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
          <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
          <HighBidder> <!-- UserType -->
            <FeedbackRatingStar> FeedbackRatingStarCodeType </FeedbackRatingStar>
            <FeedbackScore> int </FeedbackScore>
            <UserID> UserIDType (string) </UserID>
          </HighBidder>
          <PromotionalSaleDetails> <!-- PromotionalSaleDetailsType -->
            <EndTime> dateTime </EndTime>
            <OriginalPrice currencyID="CurrencyCodeType"> AmountType (double) </OriginalPrice>
            <StartTime> dateTime </StartTime>
          </PromotionalSaleDetails>
          <QuantitySold> int </QuantitySold>
          <ReserveMet> boolean </ReserveMet>
        </SellingStatus>
        <ShippingDetails> <!-- ShippingDetailsType -->
          <GlobalShipping> boolean </GlobalShipping>
          <ShippingServiceOptions> <!-- ShippingServiceOptionsType -->
            <LocalPickup> boolean </LocalPickup>
            <LogisticPlanType> string </LogisticPlanType>
            <ShippingServiceCost currencyID="CurrencyCodeType"> AmountType (double) </ShippingServiceCost>
          </ShippingServiceOptions>
          <!-- ... more ShippingServiceOptions nodes allowed here ... -->
          <ShippingType> ShippingTypeCodeType </ShippingType>
        </ShippingDetails>
        <SKU> SKUType (string) </SKU>
        <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
        <TimeLeft> duration </TimeLeft>
        <Title> string </Title>
        <Variations> <!-- VariationsType -->
          <Variation> <!-- VariationType -->
            <PrivateNotes> string </PrivateNotes>
            <Quantity> int </Quantity>
            <SellingStatus> <!-- SellingStatusType -->
              <BidCount> int </BidCount>
              <BidderCount> long </BidderCount>
              <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
              <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
              <HighBidder> <!-- UserType -->
                <FeedbackRatingStar> FeedbackRatingStarCodeType </FeedbackRatingStar>
                <FeedbackScore> int </FeedbackScore>
                <UserID> UserIDType (string) </UserID>
              </HighBidder>
              <PromotionalSaleDetails> <!-- PromotionalSaleDetailsType -->
                <EndTime> dateTime </EndTime>
                <OriginalPrice currencyID="CurrencyCodeType"> AmountType (double) </OriginalPrice>
                <StartTime> dateTime </StartTime>
              </PromotionalSaleDetails>
              <QuantitySold> int </QuantitySold>
              <ReserveMet> boolean </ReserveMet>
            </SellingStatus>
            <SKU> SKUType (string) </SKU>
            <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
            <VariationSpecifics> <!-- NameValueListArrayType -->
              <NameValueList> <!-- NameValueListType -->
                <Name> string </Name>
                <Value> string </Value>
                <!-- ... more Value values allowed here ... -->
              </NameValueList>
              <!-- ... more NameValueList nodes allowed here ... -->
            </VariationSpecifics>
            <!-- ... more VariationSpecifics nodes allowed here ... -->
            <VariationTitle> string </VariationTitle>
            <WatchCount> long </WatchCount>
          </Variation>
          <!-- ... more Variation nodes allowed here ... -->
        </Variations>
        <WatchCount> long </WatchCount>
      </Item>
      <!-- ... more Item nodes allowed here ... -->
    </ItemArray>
    <PaginationResult> <!-- PaginationResultType -->
      <TotalNumberOfEntries> int </TotalNumberOfEntries>
      <TotalNumberOfPages> int </TotalNumberOfPages>
    </PaginationResult>
  </ActiveList>
  <DeletedFromSoldList> <!-- PaginatedOrderTransactionArrayType -->
    <OrderTransactionArray> <!-- OrderTransactionArrayType -->
      <OrderTransaction> <!-- OrderTransactionType -->
        <Order> <!-- OrderType -->
          <OrderID> OrderIDType (string) </OrderID>
          <Subtotal currencyID="CurrencyCodeType"> AmountType (double) </Subtotal>
          <TransactionArray> <!-- TransactionArrayType -->
            <Transaction> <!-- TransactionType -->
              <Buyer> <!-- UserType -->
                <BuyerInfo> <!-- BuyerType -->
                  <ShippingAddress> <!-- AddressType -->
                    <PostalCode> string </PostalCode>
                  </ShippingAddress>
                </BuyerInfo>
                <Email> string </Email>
                <StaticAlias> string </StaticAlias>
                <UserID> UserIDType (string) </UserID>
              </Buyer>
              <ConvertedTransactionPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedTransactionPrice>
              <CreatedDate> dateTime </CreatedDate>
              <FeedbackLeft> <!-- FeedbackInfoType -->
                <CommentType> CommentTypeCodeType </CommentType>
              </FeedbackLeft>
              <FeedbackReceived> <!-- FeedbackInfoType -->
                <CommentType> CommentTypeCodeType </CommentType>
              </FeedbackReceived>
              <IsMultiLegShipping> boolean </IsMultiLegShipping>
              <Item> <!-- ItemType -->
                <BuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </BuyItNowPrice>
                <ClassifiedAdPayPerLeadFee currencyID="CurrencyCodeType"> AmountType (double) </ClassifiedAdPayPerLeadFee>
                <HideFromSearch> boolean </HideFromSearch>
                <ItemID> ItemIDType (string) </ItemID>
                <ListingDetails> <!-- ListingDetailsType -->
                  <ConvertedBuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedBuyItNowPrice>
                  <ConvertedReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedReservePrice>
                  <ConvertedStartPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedStartPrice>
                  <EndTime> dateTime </EndTime>
                  <StartTime> dateTime </StartTime>
                </ListingDetails>
                <ListingType> ListingTypeCodeType </ListingType>
                <PictureDetails> <!-- PictureDetailsType -->
                  <GalleryURL> anyURI </GalleryURL>
                </PictureDetails>
                <PrivateNotes> string </PrivateNotes>
                <Quantity> int </Quantity>
                <QuantityAvailable> int </QuantityAvailable>
                <QuestionCount> long </QuestionCount>
                <ReasonHideFromSearch> ReasonHideFromSearchCodeType </ReasonHideFromSearch>
                <ReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ReservePrice>
                <SellingStatus> <!-- SellingStatusType -->
                  <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
                  <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
                  <QuantitySold> int </QuantitySold>
                </SellingStatus>
                <ShippingDetails> <!-- ShippingDetailsType -->
                  <GlobalShipping> boolean </GlobalShipping>
                  <ShippingServiceOptions> <!-- ShippingServiceOptionsType -->
                    <LocalPickup> boolean </LocalPickup>
                    <LogisticPlanType> string </LogisticPlanType>
                    <ShippingServiceCost currencyID="CurrencyCodeType"> AmountType (double) </ShippingServiceCost>
                  </ShippingServiceOptions>
                  <!-- ... more ShippingServiceOptions nodes allowed here ... -->
                  <ShippingType> ShippingTypeCodeType </ShippingType>
                </ShippingDetails>
                <SKU> SKUType (string) </SKU>
                <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
                <TimeLeft> duration </TimeLeft>
                <Title> string </Title>
                <Variations> <!-- VariationsType -->
                  <Variation> <!-- VariationType -->
                    <Quantity> int </Quantity>
                    <SellingStatus> <!-- SellingStatusType -->
                      <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
                      <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
                      <QuantitySold> int </QuantitySold>
                    </SellingStatus>
                    <SKU> SKUType (string) </SKU>
                    <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
                    <VariationSpecifics> <!-- NameValueListArrayType -->
                      <NameValueList> <!-- NameValueListType -->
                        <Name> string </Name>
                        <Value> string </Value>
                        <!-- ... more Value values allowed here ... -->
                      </NameValueList>
                      <!-- ... more NameValueList nodes allowed here ... -->
                    </VariationSpecifics>
                    <!-- ... more VariationSpecifics nodes allowed here ... -->
                    <VariationTitle> string </VariationTitle>
                    <WatchCount> long </WatchCount>
                  </Variation>
                  <!-- ... more Variation nodes allowed here ... -->
                </Variations>
                <WatchCount> long </WatchCount>
              </Item>
              <OrderLineItemID> string </OrderLineItemID>
              <PaidTime> dateTime </PaidTime>
              <PaisaPayID> string </PaisaPayID>
              <Platform> TransactionPlatformCodeType </Platform>
              <QuantityPurchased> int </QuantityPurchased>
              <SellerPaidStatus> PaidStatusCodeType </SellerPaidStatus>
              <ShippedTime> dateTime </ShippedTime>
              <Status> <!-- TransactionStatusType -->
                <PaymentHoldStatus> PaymentHoldStatusCodeType </PaymentHoldStatus>
              </Status>
              <TransactionID> string </TransactionID>
              <TransactionPrice currencyID="CurrencyCodeType"> AmountType (double) </TransactionPrice>
            </Transaction>
            <!-- ... more Transaction nodes allowed here ... -->
          </TransactionArray>
        </Order>
        <Transaction> <!-- TransactionType -->
          <Buyer> <!-- UserType -->
            <BuyerInfo> <!-- BuyerType -->
              <ShippingAddress> <!-- AddressType -->
                <PostalCode> string </PostalCode>
              </ShippingAddress>
            </BuyerInfo>
            <Email> string </Email>
            <StaticAlias> string </StaticAlias>
            <UserID> UserIDType (string) </UserID>
          </Buyer>
          <ConvertedTransactionPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedTransactionPrice>
          <CreatedDate> dateTime </CreatedDate>
          <FeedbackLeft> <!-- FeedbackInfoType -->
            <CommentType> CommentTypeCodeType </CommentType>
          </FeedbackLeft>
          <FeedbackReceived> <!-- FeedbackInfoType -->
            <CommentType> CommentTypeCodeType </CommentType>
          </FeedbackReceived>
          <IsMultiLegShipping> boolean </IsMultiLegShipping>
          <Item> <!-- ItemType -->
            <BuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </BuyItNowPrice>
            <ClassifiedAdPayPerLeadFee currencyID="CurrencyCodeType"> AmountType (double) </ClassifiedAdPayPerLeadFee>
            <HideFromSearch> boolean </HideFromSearch>
            <ItemID> ItemIDType (string) </ItemID>
            <ListingDetails> <!-- ListingDetailsType -->
              <ConvertedBuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedBuyItNowPrice>
              <ConvertedReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedReservePrice>
              <ConvertedStartPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedStartPrice>
              <EndTime> dateTime </EndTime>
              <StartTime> dateTime </StartTime>
            </ListingDetails>
            <ListingType> ListingTypeCodeType </ListingType>
            <PictureDetails> <!-- PictureDetailsType -->
              <GalleryURL> anyURI </GalleryURL>
            </PictureDetails>
            <PrivateNotes> string </PrivateNotes>
            <Quantity> int </Quantity>
            <QuantityAvailable> int </QuantityAvailable>
            <QuestionCount> long </QuestionCount>
            <ReasonHideFromSearch> ReasonHideFromSearchCodeType </ReasonHideFromSearch>
            <ReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ReservePrice>
            <SellingStatus> <!-- SellingStatusType -->
              <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
              <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
              <QuantitySold> int </QuantitySold>
            </SellingStatus>
            <ShippingDetails> <!-- ShippingDetailsType -->
              <GlobalShipping> boolean </GlobalShipping>
              <ShippingServiceOptions> <!-- ShippingServiceOptionsType -->
                <LocalPickup> boolean </LocalPickup>
                <LogisticPlanType> string </LogisticPlanType>
                <ShippingServiceCost currencyID="CurrencyCodeType"> AmountType (double) </ShippingServiceCost>
              </ShippingServiceOptions>
              <!-- ... more ShippingServiceOptions nodes allowed here ... -->
              <ShippingType> ShippingTypeCodeType </ShippingType>
            </ShippingDetails>
            <SKU> SKUType (string) </SKU>
            <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
            <TimeLeft> duration </TimeLeft>
            <Title> string </Title>
            <Variations> <!-- VariationsType -->
              <Variation> <!-- VariationType -->
                <Quantity> int </Quantity>
                <SellingStatus> <!-- SellingStatusType -->
                  <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
                  <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
                  <QuantitySold> int </QuantitySold>
                </SellingStatus>
                <SKU> SKUType (string) </SKU>
                <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
                <VariationSpecifics> <!-- NameValueListArrayType -->
                  <NameValueList> <!-- NameValueListType -->
                    <Name> string </Name>
                    <Value> string </Value>
                    <!-- ... more Value values allowed here ... -->
                  </NameValueList>
                  <!-- ... more NameValueList nodes allowed here ... -->
                </VariationSpecifics>
                <!-- ... more VariationSpecifics nodes allowed here ... -->
                <VariationTitle> string </VariationTitle>
                <WatchCount> long </WatchCount>
              </Variation>
              <!-- ... more Variation nodes allowed here ... -->
            </Variations>
            <WatchCount> long </WatchCount>
          </Item>
          <OrderLineItemID> string </OrderLineItemID>
          <PaidTime> dateTime </PaidTime>
          <PaisaPayID> string </PaisaPayID>
          <Platform> TransactionPlatformCodeType </Platform>
          <QuantityPurchased> int </QuantityPurchased>
          <SellerPaidStatus> PaidStatusCodeType </SellerPaidStatus>
          <ShippedTime> dateTime </ShippedTime>
          <Status> <!-- TransactionStatusType -->
            <PaymentHoldStatus> PaymentHoldStatusCodeType </PaymentHoldStatus>
          </Status>
          <TransactionID> string </TransactionID>
          <TransactionPrice currencyID="CurrencyCodeType"> AmountType (double) </TransactionPrice>
        </Transaction>
      </OrderTransaction>
      <!-- ... more OrderTransaction nodes allowed here ... -->
    </OrderTransactionArray>
    <PaginationResult> <!-- PaginationResultType -->
      <TotalNumberOfEntries> int </TotalNumberOfEntries>
      <TotalNumberOfPages> int </TotalNumberOfPages>
    </PaginationResult>
  </DeletedFromSoldList>
  <DeletedFromUnsoldList> <!-- PaginatedItemArrayType -->
    <ItemArray> <!-- ItemArrayType -->
      <Item> <!-- ItemType -->
        <BuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </BuyItNowPrice>
        <ClassifiedAdPayPerLeadFee currencyID="CurrencyCodeType"> AmountType (double) </ClassifiedAdPayPerLeadFee>
        <HideFromSearch> boolean </HideFromSearch>
        <ItemID> ItemIDType (string) </ItemID>
        <ListingDetails> <!-- ListingDetailsType -->
          <ConvertedBuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedBuyItNowPrice>
          <ConvertedReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedReservePrice>
          <ConvertedStartPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedStartPrice>
          <EndTime> dateTime </EndTime>
          <StartTime> dateTime </StartTime>
        </ListingDetails>
        <ListingType> ListingTypeCodeType </ListingType>
        <PictureDetails> <!-- PictureDetailsType -->
          <GalleryURL> anyURI </GalleryURL>
        </PictureDetails>
        <PrivateNotes> string </PrivateNotes>
        <Quantity> int </Quantity>
        <QuantityAvailable> int </QuantityAvailable>
        <QuestionCount> long </QuestionCount>
        <ReasonHideFromSearch> ReasonHideFromSearchCodeType </ReasonHideFromSearch>
        <ReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ReservePrice>
        <SellingStatus> <!-- SellingStatusType -->
          <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
          <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
          <QuantitySold> int </QuantitySold>
        </SellingStatus>
        <ShippingDetails> <!-- ShippingDetailsType -->
          <GlobalShipping> boolean </GlobalShipping>
          <ShippingServiceOptions> <!-- ShippingServiceOptionsType -->
            <LocalPickup> boolean </LocalPickup>
            <LogisticPlanType> string </LogisticPlanType>
            <ShippingServiceCost currencyID="CurrencyCodeType"> AmountType (double) </ShippingServiceCost>
          </ShippingServiceOptions>
          <!-- ... more ShippingServiceOptions nodes allowed here ... -->
          <ShippingType> ShippingTypeCodeType </ShippingType>
        </ShippingDetails>
        <SKU> SKUType (string) </SKU>
        <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
        <TimeLeft> duration </TimeLeft>
        <Title> string </Title>
        <Variations> <!-- VariationsType -->
          <Variation> <!-- VariationType -->
            <PrivateNotes> string </PrivateNotes>
            <Quantity> int </Quantity>
            <SellingStatus> <!-- SellingStatusType -->
              <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
              <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
              <QuantitySold> int </QuantitySold>
            </SellingStatus>
            <SKU> SKUType (string) </SKU>
            <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
            <VariationSpecifics> <!-- NameValueListArrayType -->
              <NameValueList> <!-- NameValueListType -->
                <Name> string </Name>
                <Value> string </Value>
                <!-- ... more Value values allowed here ... -->
              </NameValueList>
              <!-- ... more NameValueList nodes allowed here ... -->
            </VariationSpecifics>
            <!-- ... more VariationSpecifics nodes allowed here ... -->
            <VariationTitle> string </VariationTitle>
            <WatchCount> long </WatchCount>
          </Variation>
          <!-- ... more Variation nodes allowed here ... -->
        </Variations>
        <WatchCount> long </WatchCount>
      </Item>
      <!-- ... more Item nodes allowed here ... -->
    </ItemArray>
    <PaginationResult> <!-- PaginationResultType -->
      <TotalNumberOfEntries> int </TotalNumberOfEntries>
      <TotalNumberOfPages> int </TotalNumberOfPages>
    </PaginationResult>
  </DeletedFromUnsoldList>
  <ScheduledList> <!-- PaginatedItemArrayType -->
    <ItemArray> <!-- ItemArrayType -->
      <Item> <!-- ItemType -->
        <BuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </BuyItNowPrice>
        <ClassifiedAdPayPerLeadFee currencyID="CurrencyCodeType"> AmountType (double) </ClassifiedAdPayPerLeadFee>
        <eBayNotes> string </eBayNotes>
        <HideFromSearch> boolean </HideFromSearch>
        <ItemID> ItemIDType (string) </ItemID>
        <ListingDetails> <!-- ListingDetailsType -->
          <ConvertedBuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedBuyItNowPrice>
          <ConvertedReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedReservePrice>
          <ConvertedStartPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedStartPrice>
          <StartTime> dateTime </StartTime>
        </ListingDetails>
        <ListingDuration> token </ListingDuration>
        <ListingType> ListingTypeCodeType </ListingType>
        <PictureDetails> <!-- PictureDetailsType -->
          <GalleryURL> anyURI </GalleryURL>
        </PictureDetails>
        <PrivateNotes> string </PrivateNotes>
        <Quantity> int </Quantity>
        <QuantityAvailable> int </QuantityAvailable>
        <QuestionCount> long </QuestionCount>
        <ReasonHideFromSearch> ReasonHideFromSearchCodeType </ReasonHideFromSearch>
        <ReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ReservePrice>
        <SellingStatus> <!-- SellingStatusType -->
          <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
          <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
          <QuantitySold> int </QuantitySold>
          <ReserveMet> boolean </ReserveMet>
        </SellingStatus>
        <ShippingDetails> <!-- ShippingDetailsType -->
          <GlobalShipping> boolean </GlobalShipping>
          <ShippingServiceOptions> <!-- ShippingServiceOptionsType -->
            <LocalPickup> boolean </LocalPickup>
            <LogisticPlanType> string </LogisticPlanType>
            <ShippingServiceCost currencyID="CurrencyCodeType"> AmountType (double) </ShippingServiceCost>
            <ShippingSurcharge currencyID="CurrencyCodeType"> AmountType (double) </ShippingSurcharge>
          </ShippingServiceOptions>
          <!-- ... more ShippingServiceOptions nodes allowed here ... -->
          <ShippingType> ShippingTypeCodeType </ShippingType>
        </ShippingDetails>
        <SKU> SKUType (string) </SKU>
        <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
        <TimeLeft> duration </TimeLeft>
        <Title> string </Title>
        <Variations> <!-- VariationsType -->
          <Variation> <!-- VariationType -->
            <PrivateNotes> string </PrivateNotes>
            <Quantity> int </Quantity>
            <SellingStatus> <!-- SellingStatusType -->
              <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
              <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
              <QuantitySold> int </QuantitySold>
              <ReserveMet> boolean </ReserveMet>
            </SellingStatus>
            <SKU> SKUType (string) </SKU>
            <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
            <VariationSpecifics> <!-- NameValueListArrayType -->
              <NameValueList> <!-- NameValueListType -->
                <Name> string </Name>
                <Value> string </Value>
                <!-- ... more Value values allowed here ... -->
              </NameValueList>
              <!-- ... more NameValueList nodes allowed here ... -->
            </VariationSpecifics>
            <!-- ... more VariationSpecifics nodes allowed here ... -->
            <VariationTitle> string </VariationTitle>
            <WatchCount> long </WatchCount>
          </Variation>
          <!-- ... more Variation nodes allowed here ... -->
        </Variations>
        <WatchCount> long </WatchCount>
      </Item>
      <!-- ... more Item nodes allowed here ... -->
    </ItemArray>
    <PaginationResult> <!-- PaginationResultType -->
      <TotalNumberOfEntries> int </TotalNumberOfEntries>
      <TotalNumberOfPages> int </TotalNumberOfPages>
    </PaginationResult>
  </ScheduledList>
  <SellingSummary> <!-- SellingSummaryType -->
    <ActiveAuctionCount> int </ActiveAuctionCount>
    <AuctionBidCount> int </AuctionBidCount>
    <AuctionSellingCount> int </AuctionSellingCount>
    <SoldDurationInDays> int </SoldDurationInDays>
    <TotalAuctionSellingValue currencyID="CurrencyCodeType"> AmountType (double) </TotalAuctionSellingValue>
    <TotalSoldCount> int </TotalSoldCount>
    <TotalSoldValue currencyID="CurrencyCodeType"> AmountType (double) </TotalSoldValue>
  </SellingSummary>
  <SoldList> <!-- PaginatedOrderTransactionArrayType -->
    <OrderTransactionArray> <!-- OrderTransactionArrayType -->
      <OrderTransaction> <!-- OrderTransactionType -->
        <Order> <!-- OrderType -->
          <OrderID> OrderIDType (string) </OrderID>
          <RefundAmount currencyID="CurrencyCodeType"> AmountType (double) </RefundAmount>
          <RefundStatus> string </RefundStatus>
          <Subtotal currencyID="CurrencyCodeType"> AmountType (double) </Subtotal>
          <TransactionArray> <!-- TransactionArrayType -->
            <Transaction> <!-- TransactionType -->
              <Buyer> <!-- UserType -->
                <BuyerInfo> <!-- BuyerType -->
                  <ShippingAddress> <!-- AddressType -->
                    <PostalCode> string </PostalCode>
                  </ShippingAddress>
                </BuyerInfo>
                <Email> string </Email>
                <StaticAlias> string </StaticAlias>
                <UserID> UserIDType (string) </UserID>
              </Buyer>
              <ConvertedTransactionPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedTransactionPrice>
              <CreatedDate> dateTime </CreatedDate>
              <FeedbackLeft> <!-- FeedbackInfoType -->
                <CommentType> CommentTypeCodeType </CommentType>
              </FeedbackLeft>
              <FeedbackReceived> <!-- FeedbackInfoType -->
                <CommentType> CommentTypeCodeType </CommentType>
              </FeedbackReceived>
              <IsMultiLegShipping> boolean </IsMultiLegShipping>
              <Item> <!-- ItemType -->
                <BuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </BuyItNowPrice>
                <ClassifiedAdPayPerLeadFee currencyID="CurrencyCodeType"> AmountType (double) </ClassifiedAdPayPerLeadFee>
                <HideFromSearch> boolean </HideFromSearch>
                <ItemID> ItemIDType (string) </ItemID>
                <ListingDetails> <!-- ListingDetailsType -->
                  <ConvertedBuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedBuyItNowPrice>
                  <ConvertedReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedReservePrice>
                  <ConvertedStartPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedStartPrice>
                  <EndTime> dateTime </EndTime>
                  <StartTime> dateTime </StartTime>
                </ListingDetails>
                <ListingType> ListingTypeCodeType </ListingType>
                <PictureDetails> <!-- PictureDetailsType -->
                  <GalleryURL> anyURI </GalleryURL>
                </PictureDetails>
                <PrivateNotes> string </PrivateNotes>
                <Quantity> int </Quantity>
                <QuantityAvailable> int </QuantityAvailable>
                <QuestionCount> long </QuestionCount>
                <ReasonHideFromSearch> ReasonHideFromSearchCodeType </ReasonHideFromSearch>
                <ReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ReservePrice>
                <SellingStatus> <!-- SellingStatusType -->
                  <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
                  <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
                  <QuantitySold> int </QuantitySold>
                  <ReserveMet> boolean </ReserveMet>
                </SellingStatus>
                <ShippingDetails> <!-- ShippingDetailsType -->
                  <GlobalShipping> boolean </GlobalShipping>
                  <ShippingServiceOptions> <!-- ShippingServiceOptionsType -->
                    <LocalPickup> boolean </LocalPickup>
                    <LogisticPlanType> string </LogisticPlanType>
                    <ShippingServiceCost currencyID="CurrencyCodeType"> AmountType (double) </ShippingServiceCost>
                  </ShippingServiceOptions>
                  <!-- ... more ShippingServiceOptions nodes allowed here ... -->
                  <ShippingType> ShippingTypeCodeType </ShippingType>
                </ShippingDetails>
                <SKU> SKUType (string) </SKU>
                <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
                <TimeLeft> duration </TimeLeft>
                <Title> string </Title>
                <Variations> <!-- VariationsType -->
                  <Variation> <!-- VariationType -->
                    <Quantity> int </Quantity>
                    <SellingStatus> <!-- SellingStatusType -->
                      <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
                      <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
                      <QuantitySold> int </QuantitySold>
                      <ReserveMet> boolean </ReserveMet>
                    </SellingStatus>
                    <SKU> SKUType (string) </SKU>
                    <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
                    <VariationSpecifics> <!-- NameValueListArrayType -->
                      <NameValueList> <!-- NameValueListType -->
                        <Name> string </Name>
                        <Value> string </Value>
                        <!-- ... more Value values allowed here ... -->
                      </NameValueList>
                      <!-- ... more NameValueList nodes allowed here ... -->
                    </VariationSpecifics>
                    <!-- ... more VariationSpecifics nodes allowed here ... -->
                    <VariationTitle> string </VariationTitle>
                    <WatchCount> long </WatchCount>
                  </Variation>
                  <!-- ... more Variation nodes allowed here ... -->
                </Variations>
                <WatchCount> long </WatchCount>
              </Item>
              <OrderLineItemID> string </OrderLineItemID>
              <PaidTime> dateTime </PaidTime>
              <PaisaPayID> string </PaisaPayID>
              <PaymentHoldDetails> <!-- PaymentHoldDetailType -->
                <ExpectedReleaseDate> dateTime </ExpectedReleaseDate>
                <PaymentHoldReason> PaymentHoldReasonCodeType </PaymentHoldReason>
                <RequiredSellerActionArray> <!-- RequiredSellerActionArrayType -->
                  <RequiredSellerAction> RequiredSellerActionCodeType </RequiredSellerAction>
                  <!-- ... more RequiredSellerAction values allowed here ... -->
                </RequiredSellerActionArray>
              </PaymentHoldDetails>
              <Platform> TransactionPlatformCodeType </Platform>
              <QuantityPurchased> int </QuantityPurchased>
              <SellerPaidStatus> PaidStatusCodeType </SellerPaidStatus>
              <ShippedTime> dateTime </ShippedTime>
              <Status> <!-- TransactionStatusType -->
                <PaymentHoldStatus> PaymentHoldStatusCodeType </PaymentHoldStatus>
              </Status>
              <TotalPrice currencyID="CurrencyCodeType"> AmountType (double) </TotalPrice>
              <TransactionID> string </TransactionID>
              <TransactionPrice currencyID="CurrencyCodeType"> AmountType (double) </TransactionPrice>
            </Transaction>
            <!-- ... more Transaction nodes allowed here ... -->
          </TransactionArray>
        </Order>
        <Transaction> <!-- TransactionType -->
          <Buyer> <!-- UserType -->
            <BuyerInfo> <!-- BuyerType -->
              <ShippingAddress> <!-- AddressType -->
                <PostalCode> string </PostalCode>
              </ShippingAddress>
            </BuyerInfo>
            <Email> string </Email>
            <StaticAlias> string </StaticAlias>
            <UserID> UserIDType (string) </UserID>
          </Buyer>
          <ConvertedTransactionPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedTransactionPrice>
          <CreatedDate> dateTime </CreatedDate>
          <FeedbackLeft> <!-- FeedbackInfoType -->
            <CommentType> CommentTypeCodeType </CommentType>
          </FeedbackLeft>
          <FeedbackReceived> <!-- FeedbackInfoType -->
            <CommentType> CommentTypeCodeType </CommentType>
          </FeedbackReceived>
          <IsMultiLegShipping> boolean </IsMultiLegShipping>
          <Item> <!-- ItemType -->
            <BuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </BuyItNowPrice>
            <ClassifiedAdPayPerLeadFee currencyID="CurrencyCodeType"> AmountType (double) </ClassifiedAdPayPerLeadFee>
            <HideFromSearch> boolean </HideFromSearch>
            <ItemID> ItemIDType (string) </ItemID>
            <ListingDetails> <!-- ListingDetailsType -->
              <ConvertedBuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedBuyItNowPrice>
              <ConvertedReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedReservePrice>
              <ConvertedStartPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedStartPrice>
              <EndTime> dateTime </EndTime>
              <StartTime> dateTime </StartTime>
            </ListingDetails>
            <ListingType> ListingTypeCodeType </ListingType>
            <PictureDetails> <!-- PictureDetailsType -->
              <GalleryURL> anyURI </GalleryURL>
            </PictureDetails>
            <PrivateNotes> string </PrivateNotes>
            <Quantity> int </Quantity>
            <QuantityAvailable> int </QuantityAvailable>
            <QuestionCount> long </QuestionCount>
            <ReasonHideFromSearch> ReasonHideFromSearchCodeType </ReasonHideFromSearch>
            <ReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ReservePrice>
            <SellingStatus> <!-- SellingStatusType -->
              <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
              <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
              <QuantitySold> int </QuantitySold>
              <ReserveMet> boolean </ReserveMet>
            </SellingStatus>
            <ShippingDetails> <!-- ShippingDetailsType -->
              <GlobalShipping> boolean </GlobalShipping>
              <ShippingServiceOptions> <!-- ShippingServiceOptionsType -->
                <LocalPickup> boolean </LocalPickup>
                <LogisticPlanType> string </LogisticPlanType>
                <ShippingServiceCost currencyID="CurrencyCodeType"> AmountType (double) </ShippingServiceCost>
              </ShippingServiceOptions>
              <!-- ... more ShippingServiceOptions nodes allowed here ... -->
              <ShippingType> ShippingTypeCodeType </ShippingType>
            </ShippingDetails>
            <SKU> SKUType (string) </SKU>
            <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
            <TimeLeft> duration </TimeLeft>
            <Title> string </Title>
            <Variations> <!-- VariationsType -->
              <Variation> <!-- VariationType -->
                <Quantity> int </Quantity>
                <SellingStatus> <!-- SellingStatusType -->
                  <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
                  <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
                  <QuantitySold> int </QuantitySold>
                  <ReserveMet> boolean </ReserveMet>
                </SellingStatus>
                <SKU> SKUType (string) </SKU>
                <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
                <VariationSpecifics> <!-- NameValueListArrayType -->
                  <NameValueList> <!-- NameValueListType -->
                    <Name> string </Name>
                    <Value> string </Value>
                    <!-- ... more Value values allowed here ... -->
                  </NameValueList>
                  <!-- ... more NameValueList nodes allowed here ... -->
                </VariationSpecifics>
                <!-- ... more VariationSpecifics nodes allowed here ... -->
                <VariationTitle> string </VariationTitle>
                <WatchCount> long </WatchCount>
              </Variation>
              <!-- ... more Variation nodes allowed here ... -->
            </Variations>
            <WatchCount> long </WatchCount>
          </Item>
          <OrderLineItemID> string </OrderLineItemID>
          <PaidTime> dateTime </PaidTime>
          <PaisaPayID> string </PaisaPayID>
          <PaymentHoldDetails> <!-- PaymentHoldDetailType -->
            <ExpectedReleaseDate> dateTime </ExpectedReleaseDate>
            <PaymentHoldReason> PaymentHoldReasonCodeType </PaymentHoldReason>
            <RequiredSellerActionArray> <!-- RequiredSellerActionArrayType -->
              <RequiredSellerAction> RequiredSellerActionCodeType </RequiredSellerAction>
              <!-- ... more RequiredSellerAction values allowed here ... -->
            </RequiredSellerActionArray>
          </PaymentHoldDetails>
          <Platform> TransactionPlatformCodeType </Platform>
          <QuantityPurchased> int </QuantityPurchased>
          <SellerPaidStatus> PaidStatusCodeType </SellerPaidStatus>
          <ShippedTime> dateTime </ShippedTime>
          <Status> <!-- TransactionStatusType -->
            <PaymentHoldStatus> PaymentHoldStatusCodeType </PaymentHoldStatus>
          </Status>
          <TotalPrice currencyID="CurrencyCodeType"> AmountType (double) </TotalPrice>
          <TransactionID> string </TransactionID>
          <TransactionPrice currencyID="CurrencyCodeType"> AmountType (double) </TransactionPrice>
        </Transaction>
      </OrderTransaction>
      <!-- ... more OrderTransaction nodes allowed here ... -->
    </OrderTransactionArray>
    <PaginationResult> <!-- PaginationResultType -->
      <TotalNumberOfEntries> int </TotalNumberOfEntries>
      <TotalNumberOfPages> int </TotalNumberOfPages>
    </PaginationResult>
  </SoldList>
  <Summary> <!-- MyeBaySellingSummaryType -->
    <ActiveAuctionCount> int </ActiveAuctionCount>
    <AmountLimitRemaining currencyID="CurrencyCodeType"> AmountType (double) </AmountLimitRemaining>
    <AuctionBidCount> int </AuctionBidCount>
    <AuctionSellingCount> int </AuctionSellingCount>
    <ClassifiedAdCount> int </ClassifiedAdCount>
    <ClassifiedAdOfferCount> int </ClassifiedAdOfferCount>
    <QuantityLimitRemaining> long </QuantityLimitRemaining>
    <SoldDurationInDays> int </SoldDurationInDays>
    <TotalAuctionSellingValue currencyID="CurrencyCodeType"> AmountType (double) </TotalAuctionSellingValue>
    <TotalLeadCount> int </TotalLeadCount>
    <TotalListingsWithLeads> int </TotalListingsWithLeads>
    <TotalSoldCount> int </TotalSoldCount>
    <TotalSoldValue currencyID="CurrencyCodeType"> AmountType (double) </TotalSoldValue>
  </Summary>
  <UnsoldList> <!-- PaginatedItemArrayType -->
    <ItemArray> <!-- ItemArrayType -->
      <Item> <!-- ItemType -->
        <BuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </BuyItNowPrice>
        <ClassifiedAdPayPerLeadFee currencyID="CurrencyCodeType"> AmountType (double) </ClassifiedAdPayPerLeadFee>
        <eBayNotes> string </eBayNotes>
        <HideFromSearch> boolean </HideFromSearch>
        <ItemID> ItemIDType (string) </ItemID>
        <LeadCount> int </LeadCount>
        <ListingDetails> <!-- ListingDetailsType -->
          <ConvertedBuyItNowPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedBuyItNowPrice>
          <ConvertedReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedReservePrice>
          <ConvertedStartPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedStartPrice>
          <EndTime> dateTime </EndTime>
          <StartTime> dateTime </StartTime>
        </ListingDetails>
        <ListingDuration> token </ListingDuration>
        <ListingType> ListingTypeCodeType </ListingType>
        <PictureDetails> <!-- PictureDetailsType -->
          <GalleryURL> anyURI </GalleryURL>
        </PictureDetails>
        <PrivateNotes> string </PrivateNotes>
        <Quantity> int </Quantity>
        <QuantityAvailable> int </QuantityAvailable>
        <QuestionCount> long </QuestionCount>
        <ReasonHideFromSearch> ReasonHideFromSearchCodeType </ReasonHideFromSearch>
        <Relisted> boolean </Relisted>
        <ReservePrice currencyID="CurrencyCodeType"> AmountType (double) </ReservePrice>
        <SellingStatus> <!-- SellingStatusType -->
          <BidCount> int </BidCount>
          <BidderCount> long </BidderCount>
          <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
          <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
          <QuantitySold> int </QuantitySold>
          <ReserveMet> boolean </ReserveMet>
        </SellingStatus>
        <ShippingDetails> <!-- ShippingDetailsType -->
          <GlobalShipping> boolean </GlobalShipping>
          <ShippingServiceOptions> <!-- ShippingServiceOptionsType -->
            <LocalPickup> boolean </LocalPickup>
            <LogisticPlanType> string </LogisticPlanType>
            <ShippingServiceCost currencyID="CurrencyCodeType"> AmountType (double) </ShippingServiceCost>
          </ShippingServiceOptions>
          <!-- ... more ShippingServiceOptions nodes allowed here ... -->
          <ShippingType> ShippingTypeCodeType </ShippingType>
        </ShippingDetails>
        <SKU> SKUType (string) </SKU>
        <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
        <TimeLeft> duration </TimeLeft>
        <Title> string </Title>
        <Variations> <!-- VariationsType -->
          <Variation> <!-- VariationType -->
            <PrivateNotes> string </PrivateNotes>
            <Quantity> int </Quantity>
            <SellingStatus> <!-- SellingStatusType -->
              <BidCount> int </BidCount>
              <BidderCount> long </BidderCount>
              <ConvertedCurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </ConvertedCurrentPrice>
              <CurrentPrice currencyID="CurrencyCodeType"> AmountType (double) </CurrentPrice>
              <QuantitySold> int </QuantitySold>
              <ReserveMet> boolean </ReserveMet>
            </SellingStatus>
            <SKU> SKUType (string) </SKU>
            <StartPrice currencyID="CurrencyCodeType"> AmountType (double) </StartPrice>
            <VariationSpecifics> <!-- NameValueListArrayType -->
              <NameValueList> <!-- NameValueListType -->
                <Name> string </Name>
                <Value> string </Value>
                <!-- ... more Value values allowed here ... -->
              </NameValueList>
              <!-- ... more NameValueList nodes allowed here ... -->
            </VariationSpecifics>
            <!-- ... more VariationSpecifics nodes allowed here ... -->
            <VariationTitle> string </VariationTitle>
            <WatchCount> long </WatchCount>
          </Variation>
          <!-- ... more Variation nodes allowed here ... -->
        </Variations>
        <WatchCount> long </WatchCount>
      </Item>
      <!-- ... more Item nodes allowed here ... -->
    </ItemArray>
    <PaginationResult> <!-- PaginationResultType -->
      <TotalNumberOfEntries> int </TotalNumberOfEntries>
      <TotalNumberOfPages> int </TotalNumberOfPages>
    </PaginationResult>
  </UnsoldList>
  <!-- Standard Output Fields -->
  <Ack> AckCodeType </Ack>
  <Build> string </Build>
  <CorrelationID> string </CorrelationID>
  <Errors> <!-- ErrorType -->
    <ErrorClassification> ErrorClassificationCodeType </ErrorClassification>
    <ErrorCode> token </ErrorCode>
    <ErrorParameters ParamID="string"> <!-- ErrorParameterType -->
      <Value> string </Value>
    </ErrorParameters>
    <!-- ... more ErrorParameters nodes allowed here ... -->
    <LongMessage> string </LongMessage>
    <SeverityCode> SeverityCodeType </SeverityCode>
    <ShortMessage> string </ShortMessage>
  </Errors>
  <!-- ... more Errors nodes allowed here ... -->
  <HardExpirationWarning> string </HardExpirationWarning>
  <Timestamp> dateTime </Timestamp>
  <Version> string </Version>
</GetMyeBaySellingResponse>`
	o := NewGetMyeBaySellingResult([]byte(data))
	if o.SellingSummary[4].ActiveAuctionCount != " int " {
		t.Errorf("o.SellingSummary[4].ActiveAuctionCount has not been filled out!! %+v\n",o.SellingSummary[4])
	}
	if o.SellingSummary[4].AuctionBidCount != " int " {
		t.Errorf("o.SellingSummary[4].AuctionBidCount has not been filled out!! %+v\n",o.SellingSummary[4])
	}
	if o.SellingSummary[4].AuctionSellingCount != " int " {
		t.Errorf("o.SellingSummary[4].AuctionSellingCount has not been filled out!! %+v\n",o.SellingSummary[4])
	}
	if o.SellingSummary[4].SoldDurationInDays != " int " {
		t.Errorf("o.SellingSummary[4].SoldDurationInDays has not been filled out!! %+v\n",o.SellingSummary[4])
	}
	if o.SellingSummary[4].TotalAuctionSellingValue != " AmountType (double) " {
		t.Errorf("o.SellingSummary[4].TotalAuctionSellingValue has not been filled out!! %+v\n",o.SellingSummary[4])
	}
	if o.SellingSummary[4].TotalSoldCount != " int " {
		t.Errorf("o.SellingSummary[4].TotalSoldCount has not been filled out!! %+v\n",o.SellingSummary[4])
	}
	if o.SellingSummary[4].TotalSoldValue != " AmountType (double) " {
		t.Errorf("o.SellingSummary[4].TotalSoldValue has not been filled out!! %+v\n",o.SellingSummary[4])
	}
	if o.Summary[6].ActiveAuctionCount != " int " {
		t.Errorf("o.Summary[6].ActiveAuctionCount has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].AmountLimitRemaining != " AmountType (double) " {
		t.Errorf("o.Summary[6].AmountLimitRemaining has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].AuctionBidCount != " int " {
		t.Errorf("o.Summary[6].AuctionBidCount has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].AuctionSellingCount != " int " {
		t.Errorf("o.Summary[6].AuctionSellingCount has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].ClassifiedAdCount != " int " {
		t.Errorf("o.Summary[6].ClassifiedAdCount has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].ClassifiedAdOfferCount != " int " {
		t.Errorf("o.Summary[6].ClassifiedAdOfferCount has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].QuantityLimitRemaining != " long " {
		t.Errorf("o.Summary[6].QuantityLimitRemaining has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].SoldDurationInDays != " int " {
		t.Errorf("o.Summary[6].SoldDurationInDays has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].TotalAuctionSellingValue != " AmountType (double) " {
		t.Errorf("o.Summary[6].TotalAuctionSellingValue has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].TotalLeadCount != " int " {
		t.Errorf("o.Summary[6].TotalLeadCount has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].TotalListingsWithLeads != " int " {
		t.Errorf("o.Summary[6].TotalListingsWithLeads has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].TotalSoldCount != " int " {
		t.Errorf("o.Summary[6].TotalSoldCount has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Summary[6].TotalSoldValue != " AmountType (double) " {
		t.Errorf("o.Summary[6].TotalSoldValue has not been filled out!! %+v\n",o.Summary[6])
	}
	if o.Ack != " AckCodeType " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n",o.Ack)
	}
	if o.Build != " string " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n",o.Build)
	}
	if o.CorrelationID != " string " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n",o.CorrelationID)
	}
	if o.Errors[8].ErrorClassification != " ErrorClassificationCodeType " {
		t.Errorf("o.Errors[8].ErrorClassification has not been filled out!! %+v\n",o.Errors[8])
	}
	if o.Errors[8].ErrorCode != " token " {
		t.Errorf("o.Errors[8].ErrorCode has not been filled out!! %+v\n",o.Errors[8])
	}
	if o.Errors[8].LongMessage != " string " {
		t.Errorf("o.Errors[8].LongMessage has not been filled out!! %+v\n",o.Errors[8])
	}
	if o.Errors[8].SeverityCode != " SeverityCodeType " {
		t.Errorf("o.Errors[8].SeverityCode has not been filled out!! %+v\n",o.Errors[8])
	}
	if o.Errors[8].ShortMessage != " string " {
		t.Errorf("o.Errors[8].ShortMessage has not been filled out!! %+v\n",o.Errors[8])
	}
	if o.HardExpirationWarning != " string " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n",o.HardExpirationWarning)
	}
	if o.Timestamp != " dateTime " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n",o.Timestamp)
	}
	if o.Version != " string " {
		t.Errorf("GetMyeBaySellingResult has not been filled out %+v!!\n",o.Version)
	}
}
