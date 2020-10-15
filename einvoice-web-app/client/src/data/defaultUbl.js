// TODO: make template more configurable
export const ublCreator = ({sender, receiver, price}) =>
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
                <cbc:Name>${sender}</cbc:Name>
            </cac:PartyName>
        </cac:Party>
    </cac:AccountingSupplierParty>
    <cac:AccountingCustomerParty>
        <cac:Party>
            <cac:PartyName>
                <cbc:Name>${receiver}</cbc:Name>
            </cac:PartyName>
        </cac:Party>
    </cac:AccountingCustomerParty>
    <cac:LegalMonetaryTotal>
        <cbc:PayableAmount currencyID="CAD">${price}</cbc:PayableAmount>
    </cac:LegalMonetaryTotal>
    <cac:InvoiceLine>
        <cbc:ID>1</cbc:ID>
        <cbc:LineExtensionAmount currencyID="CAD">100.00</cbc:LineExtensionAmount>
        <cac:Item>
            <cbc:Description>something</cbc:Description>
        </cac:Item>
    </cac:InvoiceLine>
</Invoice>`

export const defaultUbl = ublCreator({
  sender: 'SubjectA',
  receiver: 'subjectB',
  price: 100,
})
