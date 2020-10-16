import React from 'react'
import {connect} from 'react-redux'
import {compose, withHandlers} from 'recompose'
import {Button, Card, Form} from 'react-bootstrap'
import {withTranslation} from 'react-i18next'
import {updateInvoiceFormProperty, setGeneratedXmlInputValue} from '../../actions/invoices'
import {ublCreator} from '../../data/defaultUbl'

const InvoiceForm = ({fields, generateXml, updateTextField, t}) => (
  <Card>
    <Card.Header className="bg-primary text-white">{t('invoices:invoiceForm')}</Card.Header>
    <Card.Body>
      <Form.Group>
        <Form.Label>{t('sender')}</Form.Label>
        <Form.Control
          value={fields.sender}
          onChange={updateTextField('sender')}
        />
      </Form.Group>
      <Form.Group>
        <Form.Label>{t('receiver')}</Form.Label>
        <Form.Control
          value={fields.receiver}
          onChange={updateTextField('receiver')}
        />
      </Form.Group>
      <Form.Group>
        <Form.Label>{t('price')}</Form.Label>
        <Form.Control
          value={fields.price}
          onChange={updateTextField('price')}
        />
      </Form.Group>
      <div style={{display: 'flex'}}>
        <Button variant="primary" style={{marginLeft: 'auto'}} onClick={generateXml}>
          {t('invoices:generateInvoice')}
        </Button>
      </div>
    </Card.Body>
  </Card>
)

export default compose(
  connect(
    (state) => ({
      fields: state.createInvoiceScreen.form,
    }),
    {setGeneratedXmlInputValue, updateInvoiceFormProperty}
  ),
  withHandlers({
    updateTextField: ({updateInvoiceFormProperty}) => (property) => (e) =>
      updateInvoiceFormProperty(property, e.target.value),
    generateXml: ({fields, history, setGeneratedXmlInputValue}) => () => {
      setGeneratedXmlInputValue(ublCreator(fields))
      history.push('/create-invoice/generated')
    }
  }),
  withTranslation(['common', 'invoices']),
)(InvoiceForm)
