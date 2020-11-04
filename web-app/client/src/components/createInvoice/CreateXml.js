import React from 'react'
import {connect} from 'react-redux'
import {withTranslation} from 'react-i18next'
import ConfirmationButton from '../helpers/ConfirmationButton'
import {
  clearAttachments,
  createInvoice,
  setD16bInputValue,
  setGeneratedXmlInputValue,
  setUblInputValue,
} from '../../actions/invoices'
import {invoiceFormats} from '../../utils/constants'

class CreateXml extends React.Component {
  updateXmlInputValue = (event) => {
    this.props.setXmlInputValue(event.target.value)
  }

  submitXmlInvoice = async () => {
    const formData = new FormData()
    formData.append('format', this.props.format)
    formData.append('data', this.props.xmlInputValue)
    this.props.attachments.forEach((a, i) => formData.append(`attachment${i}`, a, a.name))

    const newInvoiceId = await this.props.createInvoice(formData)
    await this.props.clearAttachments()
    if (newInvoiceId) {
      this.props.history.push(`/invoices/${newInvoiceId}`)
    }
  }

  render = () => (
    <div>
      <p className="row justify-content-center">{this.props.t('format')}: {this.props.format}</p>
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
            confirmationTitle={this.props.t('TopBar:tabs.createInvoice')}
            confirmationText={this.props.t('invoices:confirmationQuestion')}
          >
            {this.props.t('submit')}
          </ConfirmationButton>
        </div>
      </div>
    </div>
  )
}

const TranslatedCreateXml = withTranslation(['common', 'TopBar', 'invoices'])(CreateXml)

export const CreateUbl = connect(
  (state) => ({
    xmlInputValue: state.createInvoiceScreen.ublInput,
    format: invoiceFormats.UBL,
    attachments: state.createInvoiceScreen.attachments,
  }),
  {createInvoice, setXmlInputValue: setUblInputValue, clearAttachments}
)(TranslatedCreateXml)

export const CreateD16b = connect(
  (state) => ({
    xmlInputValue: state.createInvoiceScreen.d16bInput,
    format: invoiceFormats.D16B,
    attachments: state.createInvoiceScreen.attachments,
  }),
  {createInvoice, setXmlInputValue: setD16bInputValue, clearAttachments}
)(TranslatedCreateXml)

export const CreateGenerated = connect(
  (state) => ({
    xmlInputValue: state.createInvoiceScreen.formGeneratedInput,
    format: invoiceFormats.UBL,
    attachments: state.createInvoiceScreen.attachments,
  }),
  {createInvoice, setXmlInputValue: setGeneratedXmlInputValue, clearAttachments}
)(TranslatedCreateXml)
