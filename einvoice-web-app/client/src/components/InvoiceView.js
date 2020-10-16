import React from 'react'
import {connect} from 'react-redux'
import {NavLink} from 'react-router-dom'
import {compose, lifecycle, withHandlers} from 'recompose'
import {getInvoiceDetail, setCurrentInvoice} from '../actions/invoices'
import {withTranslation} from 'react-i18next'

const InvoiceView = ({invoice, match: {params: {id}}, resetCurrentInvoice, t}) => (
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
  </div>
)

export default compose(
  connect(
  (state) => ({
    invoice: state.currentInvoice,
    user: state.user,
  }),
  {getInvoiceDetail, setCurrentInvoice}
  ),
  lifecycle({
    componentDidMount() {
      this.props.getInvoiceDetail(this.props.match.params.id)
    },
  }),
  withHandlers({
    resetCurrentInvoice: ({setCurrentInvoice}) => () => {
      setCurrentInvoice(null)
    },
  }),
  withTranslation('common'),
)(InvoiceView)
