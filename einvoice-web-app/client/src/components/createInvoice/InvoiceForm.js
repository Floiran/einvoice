import React from 'react'
import {connect} from 'react-redux'
import {compose, withHandlers} from 'recompose'
import {Button, Card, Form} from 'react-bootstrap'
import {updateInvoiceFormProperty, setGeneratedXmlInputValue} from '../../actions/invoices'
import {ublCreator} from '../../data/defaultUbl'

const InvoiceForm = ({fields, updateTextField, generateXml}) => (
  <Card>
    <Card.Header className="bg-primary text-white">Invoice Form</Card.Header>
    <Card.Body>
      <Form.Group>
        <Form.Label>Sender</Form.Label>
        <Form.Control
          value={fields.sender}
          onChange={updateTextField('sender')}
        />
      </Form.Group>
      <Form.Group>
        <Form.Label>Receiver</Form.Label>
        <Form.Control
          value={fields.receiver}
          onChange={updateTextField('receiver')}
        />
      </Form.Group>
      <Form.Group>
        <Form.Label>Price</Form.Label>
        <Form.Control
          value={fields.price}
          onChange={updateTextField('price')}
        />
      </Form.Group>
      <div style={{display: 'flex'}}>
        <Button variant="primary" style={{marginLeft: 'auto'}} onClick={generateXml}>
          Generate invoice
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
)(InvoiceForm)
