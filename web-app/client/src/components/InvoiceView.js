import React from 'react'
import {connect} from 'react-redux'
import {NavLink} from 'react-router-dom'
import {branch, compose, lifecycle, renderComponent, renderNothing} from 'recompose'
import {Button} from 'react-bootstrap'
import {withTranslation} from 'react-i18next'
import {get} from 'lodash'
import NotFound from './helpers/NotFound'
import {getInvoiceDetail, getInvoiceMeta} from '../actions/invoices'
import {CONFIG} from '../appSettings'

const InvoiceView = ({attachments, invoice, match: {params: {id}}, t}) => (
  <div>
    <h2 style={{textAlign: 'center'}}>{t('invoice')} {id}</h2>
    <div className="row justify-content-center">
      <NavLink to="/invoices">
        <Button variant="primary">{t('close')}</Button>
      </NavLink>
    </div>
    <div style={{borderStyle: 'solid'}}>
      {invoice}
    </div>
    {attachments.map((a, index) => (
      <p key={index}>
        <a className="row" href={`${CONFIG.authServerUrl}/attachments/${a.id}`}>{a.name}</a>
      </p>
    ))}
  </div>
)

export default compose(
  connect(
    (state, {match: {params: {id}}}) => ({
      invoice: get(state, ['invoices', id, 'data']),
      attachments: get(state, ['invoices', id, 'attachments']),
      invoiceDoesNotExist: get(state, ['invoices', id, 'notFound']),
    }),
    {getInvoiceDetail, getInvoiceMeta}
  ),
  lifecycle({
    componentDidMount() {
      this.props.getInvoiceDetail(this.props.match.params.id)
      this.props.getInvoiceMeta(this.props.match.params.id)
    },
  }),
  branch(
    ({invoice, attachments}) => invoice == null || !attachments,
    renderNothing,
  ),
  branch(
    ({invoiceDoesNotExist}) => invoiceDoesNotExist,
    renderComponent(NotFound),
  ),
  withTranslation('common'),
)(InvoiceView)
