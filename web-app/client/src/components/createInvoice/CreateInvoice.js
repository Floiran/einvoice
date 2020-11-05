import React from 'react'
import {connect} from 'react-redux'
import {compose, withHandlers} from 'recompose'
import {withTranslation} from 'react-i18next'
import Attachments from './Attachments'
import ConfirmationButton from '../helpers/ConfirmationButton'
import {
  addAttachment, clearAttachments, createInvoice, removeAttachment, setCreateInvoiceValue
} from '../../actions/invoices'
import {tabToInvoiceFormat} from '../../utils/constants'

const CreateInvoice = ({
  addAttachment, attachments, format, invoice, removeAttachment, submitInvoice, t, updateInvoiceInputValue,
}) => (
  <div>
    <p className="row justify-content-center">{t('format')}: {format}</p>
    <div className="row justify-content-center">
      <textarea
        className="col"
        cols="50"
        rows="13"
        value={invoice}
        onChange={updateInvoiceInputValue}
      />
    </div>
    <Attachments
      attachments={attachments}
      addAttachment={addAttachment}
      removeAttachment={removeAttachment}
    />
    <div className="row justify-content-center">
      <ConfirmationButton
        className="btn btn-primary"
        onClick={submitInvoice}
        confirmationTitle={t('TopBar:tabs.createInvoice')}
        confirmationText={t('invoices:confirmationQuestion')}
      >
        {t('submit')}
      </ConfirmationButton>
    </div>
  </div>
)

export default compose(
  connect(
    (state, {location: {pathname}}) => {
      const tab = pathname.split('/').slice(-1)[0]
      return {
        format: tabToInvoiceFormat[tab],
        invoice: state.createInvoiceScreen[tab].invoice,
        attachments: state.createInvoiceScreen[tab].attachments,
      }
    },
    (dispatch, {location: {pathname}}) => {
      const tab = pathname.split('/').slice(-1)[0]
      return {
        createInvoice: (data) => dispatch(createInvoice(data)),
        setInvoiceInputValue: (invoice) => dispatch(setCreateInvoiceValue(tab)(invoice)),
        clearAttachments: () => dispatch(clearAttachments(tab)),
        addAttachment: (attachment) => dispatch(addAttachment(tab, attachment)),
        removeAttachment: (attachment) => dispatch(removeAttachment(tab, attachment)),
      }
    }
  ),
  withHandlers({
    updateInvoiceInputValue: ({setInvoiceInputValue}) => (e) => setInvoiceInputValue(e.target.value),
    submitInvoice: ({attachments, clearAttachments, createInvoice, format, history, invoice}) => async () => {
      const formData = new FormData()
      formData.append('format', format)
      formData.append('data', invoice)
      attachments.forEach((a, i) => formData.append(`attachment${i}`, a, a.name))

      const newInvoiceId = await createInvoice(formData)
      if (newInvoiceId) {
        history.push(`/invoices/${newInvoiceId}`)
        clearAttachments()
      }
    }
  }),
  withTranslation(['common', 'TopBar', 'invoices']),
)(CreateInvoice)
