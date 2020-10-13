import React from 'react'
import {connect} from 'react-redux'
import {NavLink} from 'react-router-dom'
import {branch, compose, lifecycle, renderNothing, withHandlers} from 'recompose'
import {getInvoiceDetail, setCurrentInvoice} from '../actions/invoices'

const InvoiceView = ({resetCurrentInvoice, invoice, match: {params: {id}}}) => (
  <div>
    <h2>Invoice {id}</h2>
    <div className='row justify-content-center'>
      <NavLink to="/invoices">
        <button className='btn btn-primary' onClick={resetCurrentInvoice}>Close</button>
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
  branch(
    ({user}) => !user,
    // TODO: this should be visible for everyone
    renderNothing,
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
  })
)(InvoiceView)
