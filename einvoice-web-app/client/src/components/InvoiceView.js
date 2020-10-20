import React from 'react'
import {connect} from 'react-redux'
import {NavLink} from 'react-router-dom'
import {compose, lifecycle, withHandlers} from 'recompose'
import {getInvoiceDetail, getInvoiceMeta, setCurrentInvoice, setCurrentInvoiceMeta} from '../actions/invoices'
import {withTranslation} from 'react-i18next'

const InvoiceView = ({invoice, match: {params: {id}}, resetCurrentInvoice, meta, apiServerUrl, t}) => (
  <div>
    <h2 style={{textAlign: 'center'}}>{t('invoice')} {id}</h2>
    <div className='row justify-content-center'>
      <NavLink to="/invoices">
        <button className='btn btn-primary' onClick={resetCurrentInvoice}>{t('close')}</button>
      </NavLink>
    </div>
    <div style={{borderStyle: 'solid'}}>
      {invoice}
    </div>
    {
      meta && meta.attachments.map(a =>
        <p><a className={'row'} href={`${apiServerUrl}/api/attachments/${a.id}`}>{a.name}</a></p>
      )
    }
  </div>
)

export default compose(
  connect(
    (state) => ({
      invoice: state.currentInvoice,
      meta: state.currentInvoiceMeta,
      user: state.user,
      apiServerUrl: state.urls.apiServerUrl,
    }),
    {getInvoiceDetail, setCurrentInvoice, getInvoiceMeta, setCurrentInvoiceMeta}
  ),
  lifecycle({
    componentDidMount() {
      this.props.getInvoiceDetail(this.props.match.params.id)
      this.props.getInvoiceMeta(this.props.match.params.id)
    },
  }),
  withHandlers({
    resetCurrentInvoice: ({setCurrentInvoice, setCurrentInvoiceMeta}) => () => {
      setCurrentInvoice(null)
      setCurrentInvoiceMeta(null)
    },
  }),
  withTranslation(['common']),
)(InvoiceView)
