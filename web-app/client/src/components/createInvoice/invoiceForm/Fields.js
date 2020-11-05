import React from 'react'
import {connect} from 'react-redux'
import {Form} from 'react-bootstrap'
import {useTranslation} from 'react-i18next'
import DatePicker from '../../helpers/DatePicker'
import {updateInvoiceFormProperty} from '../../../actions/invoices'

const _TextField = ({fieldName, updateField, value}) => {
  const {t} = useTranslation('invoices')
  return (
    <Form.Group>
      <Form.Label>{t(fieldName)}</Form.Label>
      <Form.Control
        value={value}
        onChange={updateField}
      />
    </Form.Group>
  )
}

export const TextField = connect(
  (state, {fieldName}) => ({
    value: state.createInvoiceScreen.form[fieldName],
  }),
  (dispatch, {fieldName}) => ({
    updateField: (e) => dispatch(updateInvoiceFormProperty(fieldName, e.target.value)),
  })
)(_TextField)

const _DateField = ({fieldName, updateField, value}) => {
  const {t} = useTranslation('invoices')
  return (
    <Form.Group>
      <Form.Label>{t(fieldName)}</Form.Label>
      <div>
        <DatePicker
          selected={value}
          onChange={updateField}
          className="form-control"
          dateFormat="yyyy-MM-dd"
        />
      </div>
    </Form.Group>
  )
}

export const DateField = connect(
  (state, {fieldName}) => ({
    value: state.createInvoiceScreen.form[fieldName],
  }),
  (dispatch, {fieldName}) => ({
    updateField: (time) => dispatch(updateInvoiceFormProperty(fieldName, time)),
  })
)(_DateField)
