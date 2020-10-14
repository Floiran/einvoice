import React from 'react'
import {connect} from 'react-redux'
import {branch, compose, lifecycle, renderNothing} from 'recompose'
import {NavLink} from 'react-router-dom'
import {getInvoices} from '../actions/invoices'

const InvoiceList = ({invoices})  => {
  let rows = []
  if(invoices) {
    rows = invoices.map((invoice, i) => {
      return <tr key={i}>
        <th scope="row"><p>{i+1}</p></th>
        <td><p>{invoice.id}</p></td>
        <td><p>{invoice.sender}</p></td>
        <td><p>{invoice.receiver}</p></td>
        <td><p>{invoice.price}</p></td>
        <td><p>{invoice.format}</p></td>
        <td><p>
          <NavLink to={`/invoices/${invoice.id}`}>view</NavLink>
        </p></td>
      </tr>
    })
  }

  return (
    <div className="container">
      <h2>All invoices</h2>
      <table className="table table-striped table-borderless table-hover">
        <thead>
        <tr>
          <th>#</th>
          <th>ID</th>
          <th>Sender</th>
          <th>Receiver</th>
          <th>Price</th>
          <th>Format</th>
          <th />
        </tr>
        </thead>
        <tbody>
          {rows}
        </tbody>
      </table>
    </div>
  )
}

export default compose(
  connect(
    (state) => ({
      invoices: state.invoices,
      user: state.user,
    }),
    {getInvoices}
  ),
  lifecycle({
    componentDidMount() {
      this.props.getInvoices()
    },
  }),
)(InvoiceList)
