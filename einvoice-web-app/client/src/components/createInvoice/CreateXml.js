import React from 'react'
import {connect} from 'react-redux'
import ConfirmationButton from '../helpers/ConfirmationButton'
import {createInvoice, setUblInputValue, setD16bInputValue, setGeneratedXmlInputValue} from '../../actions/invoices'

class CreateXml extends React.Component {
  updateXmlInputValue = (event) => {
    this.props.setXmlInputValue(event.target.value)
  }

  submitXmlInvoice = async () => {
    const newInvoiceId = await this.props.createInvoice(this.props.format, this.props.xmlInputValue)
    if (newInvoiceId) {
      this.props.history.push(`/invoices/${newInvoiceId}`)
    }
  }

  render = () => (
    <div>
      <p className="row justify-content-center">Format: {this.props.format}</p>
      <div>
        <div className="row justify-content-center">
          <textarea
            className="col"
            name="xml"
            cols="50"
            rows="15"
            value={this.props.xmlInputValue}
            onChange={this.updateXmlInputValue}
          />
        </div>
        <div className="row justify-content-center">
          <ConfirmationButton
            className="btn btn-primary"
            onClick={this.submitXmlInvoice}
            confirmationTitle="Create invoice"
            confirmationText="Do you really want to create this invoice?"
          >
            Submit
          </ConfirmationButton>
        </div>
      </div>
    </div>
  )
}

export const CreateUbl = connect(
  (state) => ({
    xmlInputValue: state.createInvoiceScreen.ublInput,
    format: 'ubl',
  }),
  {createInvoice, setXmlInputValue: setUblInputValue}
)(CreateXml)

export const CreateD16b = connect(
  (state) => ({
    xmlInputValue: state.createInvoiceScreen.d16bInput,
    format: 'd16b',
  }),
  {createInvoice, setXmlInputValue: setD16bInputValue}
)(CreateXml)

export const CreateGenerated = connect(
  (state) => ({
    xmlInputValue: state.createInvoiceScreen.formGeneratedInput,
    format: 'ubl',
  }),
  {createInvoice, setXmlInputValue: setGeneratedXmlInputValue}
)(CreateXml)
