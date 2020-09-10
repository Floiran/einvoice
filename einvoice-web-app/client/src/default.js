

export let defaultUbl =
`<Invoice xmlns="urn:oasis:names:specification:ubl:schema:xsd:Invoice-2"
          xmlns:cbc="urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2"
          xmlns:cac="urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2">
    <cbc:ID>123</cbc:ID>
    <cbc:IssueDate>2011-09-22</cbc:IssueDate>
    <cac:InvoicePeriod>
        <cbc:StartDate>2011-08-01</cbc:StartDate>
        <cbc:EndDate>2011-08-31</cbc:EndDate>
    </cac:InvoicePeriod>
    <cac:AccountingSupplierParty>
        <cac:Party>
            <cac:PartyName>
                <cbc:Name>Custom Cotter Pins</cbc:Name>
            </cac:PartyName>
        </cac:Party>
    </cac:AccountingSupplierParty>
    <cac:AccountingCustomerParty>
        <cac:Party>
            <cac:PartyName>
                <cbc:Name>North American Veeblefetzer</cbc:Name>
            </cac:PartyName>
        </cac:Party>
    </cac:AccountingCustomerParty>
    <cac:LegalMonetaryTotal>
        <cbc:PayableAmount currencyID="CAD">100.00</cbc:PayableAmount>
    </cac:LegalMonetaryTotal>
    <cac:InvoiceLine>
        <cbc:ID>1</cbc:ID>
        <cbc:LineExtensionAmount currencyID="CAD">100.00</cbc:LineExtensionAmount>
        <cac:Item>
            <cbc:Description>Cotter pin, MIL-SPEC</cbc:Description>
        </cac:Item>
    </cac:InvoiceLine>
</Invoice>`

export let defaultD16b =
`<rsm:CrossIndustryInvoice xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                           xsi:schemaLocation="urn:un:unece:uncefact:data:standard:CrossIndustryInvoice:100 ../xsd/data/standard/CrossIndustryInvoice_100pD16B.xsd"
                           xmlns:udt="urn:un:unece:uncefact:data:standard:UnqualifiedDataType:100"
                           xmlns:rsm="urn:un:unece:uncefact:data:standard:CrossIndustryInvoice:100"
                           xmlns:ram="urn:un:unece:uncefact:data:standard:ReusableAggregateBusinessInformationEntity:100">
    <rsm:ExchangedDocumentContext>
        <ram:GuidelineSpecifiedDocumentContextParameter>
            <ram:ID>urn:cen.eu:en16931:2017</ram:ID>
        </ram:GuidelineSpecifiedDocumentContextParameter>
    </rsm:ExchangedDocumentContext>
    <rsm:ExchangedDocument>
        <ram:ID>TOSL110</ram:ID>
        <ram:TypeCode>380</ram:TypeCode>
        <ram:IssueDateTime>
            <udt:DateTimeString format="102">20130410</udt:DateTimeString>
        </ram:IssueDateTime>
    </rsm:ExchangedDocument>
    <rsm:SupplyChainTradeTransaction>
        <ram:IncludedSupplyChainTradeLineItem>
            <ram:AssociatedDocumentLineDocument>
                <ram:LineID>3</ram:LineID>
            </ram:AssociatedDocumentLineDocument>
            <ram:SpecifiedTradeProduct>
                <ram:SellerAssignedID>JB009</ram:SellerAssignedID>
                <ram:Name>American Cookies</ram:Name>
            </ram:SpecifiedTradeProduct>
            <ram:SpecifiedLineTradeAgreement>
                <ram:NetPriceProductTradePrice>
                    <ram:ChargeAmount>5</ram:ChargeAmount>
                </ram:NetPriceProductTradePrice>
            </ram:SpecifiedLineTradeAgreement>
            <ram:SpecifiedLineTradeDelivery>
                <ram:BilledQuantity unitCode="C62">500</ram:BilledQuantity>
            </ram:SpecifiedLineTradeDelivery>
            <ram:SpecifiedLineTradeSettlement>
                <ram:ApplicableTradeTax>
                    <ram:TypeCode>VAT</ram:TypeCode>
                    <ram:CategoryCode>S</ram:CategoryCode>
                    <ram:RateApplicablePercent>12</ram:RateApplicablePercent>
                </ram:ApplicableTradeTax>
                <ram:SpecifiedTradeSettlementLineMonetarySummation>
                    <ram:LineTotalAmount>2500</ram:LineTotalAmount>
                </ram:SpecifiedTradeSettlementLineMonetarySummation>
            </ram:SpecifiedLineTradeSettlement>
        </ram:IncludedSupplyChainTradeLineItem>
        <ram:ApplicableHeaderTradeAgreement>
            <ram:SellerTradeParty>
                <ram:GlobalID schemeID="0088">5790000436101</ram:GlobalID>
                <ram:Name>SellerCompany</ram:Name>
                <ram:SpecifiedLegalOrganization>
                    <ram:ID>DK16356706</ram:ID>
                </ram:SpecifiedLegalOrganization>
                <ram:DefinedTradeContact>
                    <ram:PersonName>Anthon Larsen</ram:PersonName>
                    <ram:TelephoneUniversalCommunication>
                        <ram:CompleteNumber>+4598989898</ram:CompleteNumber>
                    </ram:TelephoneUniversalCommunication>
                    <ram:EmailURIUniversalCommunication>
                        <ram:URIID>Anthon@SellerCompany.dk</ram:URIID>
                    </ram:EmailURIUniversalCommunication>
                </ram:DefinedTradeContact>
                <ram:PostalTradeAddress>
                    <ram:PostcodeCode>54321</ram:PostcodeCode>
                    <ram:LineOne>Main street 2, Building 4</ram:LineOne>
                    <ram:CityName>Big city</ram:CityName>
                    <ram:CountryID>DK</ram:CountryID>
                </ram:PostalTradeAddress>
                <ram:SpecifiedTaxRegistration>
                    <ram:ID schemeID="VA">DK16356706</ram:ID>
                </ram:SpecifiedTaxRegistration>
            </ram:SellerTradeParty>
            <ram:BuyerTradeParty>
                <ram:GlobalID schemeID="0088">5790000436057</ram:GlobalID>
                <ram:Name>Buyercompany ltd</ram:Name>
                <ram:DefinedTradeContact>
                    <ram:PersonName>John Hansen</ram:PersonName>
                </ram:DefinedTradeContact>
                <ram:PostalTradeAddress>
                    <ram:PostcodeCode>101</ram:PostcodeCode>
                    <ram:LineOne>Anystreet, Building 1</ram:LineOne>
                    <ram:CityName>Anytown</ram:CityName>
                    <ram:CountryID>DK</ram:CountryID>
                </ram:PostalTradeAddress>
            </ram:BuyerTradeParty>
            <ram:BuyerOrderReferencedDocument>
                <ram:IssuerAssignedID>123</ram:IssuerAssignedID>
            </ram:BuyerOrderReferencedDocument>
        </ram:ApplicableHeaderTradeAgreement>
        <ram:ApplicableHeaderTradeDelivery>
            <ram:ShipToTradeParty>
                <ram:PostalTradeAddress>
                    <ram:PostcodeCode>9000</ram:PostcodeCode>
                    <ram:LineOne>Deliverystreet</ram:LineOne>
                    <ram:CityName>Deliverycity</ram:CityName>
                    <ram:CountryID>DK</ram:CountryID>
                </ram:PostalTradeAddress>
            </ram:ShipToTradeParty>
            <ram:ActualDeliverySupplyChainEvent>
                <ram:OccurrenceDateTime>
                    <udt:DateTimeString format="102">20130415</udt:DateTimeString>
                </ram:OccurrenceDateTime>
            </ram:ActualDeliverySupplyChainEvent>
        </ram:ApplicableHeaderTradeDelivery>
        <ram:ApplicableHeaderTradeSettlement>
            <ram:PaymentReference>Payref1</ram:PaymentReference>
            <ram:InvoiceCurrencyCode>DKK</ram:InvoiceCurrencyCode>
            <ram:SpecifiedTradeSettlementPaymentMeans>
                <ram:TypeCode>30</ram:TypeCode>
                <ram:PayeePartyCreditorFinancialAccount>
                    <ram:IBANID>DK1212341234123412</ram:IBANID>
                </ram:PayeePartyCreditorFinancialAccount>
            </ram:SpecifiedTradeSettlementPaymentMeans>
            <ram:SpecifiedTradeSettlementPaymentMeans>
                <ram:TypeCode>30</ram:TypeCode>
                <ram:PayeePartyCreditorFinancialAccount>
                    <ram:IBANID>DK1212341234123412</ram:IBANID>
                </ram:PayeePartyCreditorFinancialAccount>
            </ram:SpecifiedTradeSettlementPaymentMeans>
            <ram:ApplicableTradeTax>
                <ram:CalculatedAmount>375</ram:CalculatedAmount>
                <ram:TypeCode>VAT</ram:TypeCode>
                <ram:BasisAmount>1500</ram:BasisAmount>
                <ram:CategoryCode>S</ram:CategoryCode>
                <ram:RateApplicablePercent>25</ram:RateApplicablePercent>
            </ram:ApplicableTradeTax>
            <ram:ApplicableTradeTax>
                <ram:CalculatedAmount>300</ram:CalculatedAmount>
                <ram:TypeCode>VAT</ram:TypeCode>
                <ram:BasisAmount>2500</ram:BasisAmount>
                <ram:CategoryCode>S</ram:CategoryCode>
                <ram:RateApplicablePercent>12</ram:RateApplicablePercent>
            </ram:ApplicableTradeTax>
            <ram:SpecifiedTradePaymentTerms>
                <ram:DueDateDateTime><udt:DateTimeString format="102">20130510</udt:DateTimeString></ram:DueDateDateTime>
            </ram:SpecifiedTradePaymentTerms>
            <ram:SpecifiedTradeSettlementHeaderMonetarySummation>
                <ram:LineTotalAmount>4000</ram:LineTotalAmount>
                <ram:TaxBasisTotalAmount>4000</ram:TaxBasisTotalAmount>
                <ram:TaxTotalAmount currencyID="DKK">675</ram:TaxTotalAmount>
                <ram:GrandTotalAmount>4675</ram:GrandTotalAmount>
                <ram:DuePayableAmount>4675</ram:DuePayableAmount>
            </ram:SpecifiedTradeSettlementHeaderMonetarySummation>
        </ram:ApplicableHeaderTradeSettlement>
    </rsm:SupplyChainTradeTransaction>
</rsm:CrossIndustryInvoice>`