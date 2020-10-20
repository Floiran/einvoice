import React from 'react'
import {connect} from 'react-redux'
import {NavLink} from 'react-router-dom'
import {branch, compose, lifecycle, renderNothing} from 'recompose'
import {Button} from 'react-bootstrap'
import {withTranslation} from 'react-i18next'
import {get} from 'lodash'
import {getInvoiceDetail, getInvoiceMeta} from '../actions/invoices'

const InvoiceView = ({attachments, invoice, match: {params: {id}}, apiServerUrl, t}) => (
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
    {attachments.map((a, index) =>
      <p key={index} >
        <a className="row" href={`${apiServerUrl}/api/attachments/${a.id}`}>{a.name}</a>
      </p>
    )}
  </div>
)

export default compose(
  connect(
    (state, {match: {params: {id}}}) => ({
      invoice: get(state, ['invoices', id, 'data']),
      attachments: get(state, ['invoices', id, 'attachments']),
      apiServerUrl: state.urls.apiServerUrl,
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
    ({invoice, attachments}) => !invoice || !attachments,
    renderNothing,
  ),
  withTranslation('common'),
)(InvoiceView)
