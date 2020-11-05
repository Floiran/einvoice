import React from 'react'
import {connect} from 'react-redux'
import {compose, withHandlers} from 'recompose'
import {Button, Card, Col, Form} from 'react-bootstrap'
import {withTranslation} from 'react-i18next'
import {DateField, TextField} from './Fields'
import {setCreateInvoiceValue} from '../../../actions/invoices'
import {ublCreator} from '../../../data/defaultUbl'

const InvoiceForm = ({generateInvoice, t}) => (
  <Card>
    <Card.Header className="bg-primary text-white" as="h5">{t('invoiceForm')}</Card.Header>
    <Card.Body>
      <TextField fieldName="invoiceId" />
      <Form.Row>
        <Col>
          <TextField fieldName="sender" />
        </Col>
        <Col>
          <TextField fieldName="receiver" />
        </Col>
      </Form.Row>
      <Form.Row>
        <Col>
          <TextField fieldName="price" />
        </Col>
        <Col>
          <TextField fieldName="vat" />
        </Col>
      </Form.Row>
      <Form.Row>
        <Col>
          <DateField fieldName="issueDate" />
        </Col>
        <Col>
          <DateField fieldName="dueDate" />
        </Col>
        <Col>
          <DateField fieldName="invoicePeriodStartDate" />
        </Col>
        <Col>
          <DateField fieldName="invoicePeriodEndDate" />
        </Col>
      </Form.Row>
      <div style={{display: 'flex'}}>
        <Button variant="primary" style={{marginLeft: 'auto'}} onClick={generateInvoice}>
          {t('generateInvoice')}
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
    {setInvoiceValue: setCreateInvoiceValue('generated')}
  ),
  withHandlers({
    generateInvoice: ({fields, history, setInvoiceValue}) => () => {
      setInvoiceValue(ublCreator(fields))
      history.push('/create-invoice/generated')
    },
  }),
  withTranslation('invoices'),
)(InvoiceForm)
