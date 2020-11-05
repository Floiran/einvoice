import {format} from 'date-fns'

const dueDateCreator = (dueDate) =>
  `    <cbc:DueDate>${format(dueDate, 'yyyy-MM-dd')}</cbc:DueDate>`

const invoicePeriodCreator = (invoicePeriodStartDate, invoicePeriodEndDate) =>
  `    <cac:InvoicePeriod>
        <cbc:StartDate>${format(invoicePeriodStartDate, 'yyyy-MM-dd')}</cbc:StartDate>
        <cbc:EndDate>${format(invoicePeriodEndDate, 'yyyy-MM-dd')}</cbc:EndDate>
    </cac:InvoicePeriod>`

const accountingSupplierPartyCreator = (sender) =>
  `    <cac:AccountingSupplierParty>
        <cac:Party>
            <cac:PartyName>
                <cbc:Name>${sender}</cbc:Name>
            </cac:PartyName>
        </cac:Party>
    </cac:AccountingSupplierParty>`

const accountingCustomerPartyCreator = (receiver) =>
  `    <cac:AccountingCustomerParty>
        <cac:Party>
            <cac:PartyName>
                <cbc:Name>${receiver}</cbc:Name>
            </cac:PartyName>
        </cac:Party>
    </cac:AccountingCustomerParty>`

const legalMonetaryTotalCreator = (price) =>
  `    <cac:LegalMonetaryTotal>
        <cbc:PayableAmount currencyID="CAD">${price}</cbc:PayableAmount>
    </cac:LegalMonetaryTotal>`

const invoiceLineCreator = () =>
  `    <cac:InvoiceLine>
        <cbc:ID>1</cbc:ID>
        <cbc:LineExtensionAmount currencyID="CAD">100.00</cbc:LineExtensionAmount>
        <cac:Item>
            <cbc:Description>something</cbc:Description>
        </cac:Item>
    </cac:InvoiceLine>`

export const ublCreator = ({
  sender, receiver, price, issueDate, dueDate, invoiceId,
  invoicePeriodStartDate, invoicePeriodEndDate,
}) => {
  const invoiceParts = [`<Invoice xmlns="urn:oasis:names:specification:ubl:schema:xsd:Invoice-2"
          xmlns:cbc="urn:oasis:names:specification:ubl:schema:xsd:CommonBasicComponents-2"
          xmlns:cac="urn:oasis:names:specification:ubl:schema:xsd:CommonAggregateComponents-2">
    <cbc:ID>${invoiceId}</cbc:ID>
    <cbc:IssueDate>${format(issueDate, 'yyyy-MM-dd')}</cbc:IssueDate>`]
  if (dueDate) invoiceParts.push(dueDateCreator(dueDate))
  if (invoicePeriodStartDate && invoicePeriodEndDate) {
    invoiceParts.push(invoicePeriodCreator(invoicePeriodStartDate, invoicePeriodEndDate))
  }
  invoiceParts.push(
    accountingSupplierPartyCreator(sender),
    accountingCustomerPartyCreator(receiver),
    legalMonetaryTotalCreator(price),
    invoiceLineCreator(),
    '</Invoice>'
  )
  return invoiceParts.join('\n')
}

export const defaultUbl = ublCreator({
  sender: 'SubjectA',
  receiver: 'subjectB',
  price: 100,
  vat: 5,
  invoiceId: 123,
  issueDate: Date.now(),
  dueDate: Date.now(),
  invoicePeriodStartDate: Date.now(),
  invoicePeriodEndDate: Date.now(),
})
