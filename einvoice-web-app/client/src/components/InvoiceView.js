import React from 'react'
import {connect} from 'react-redux'
import {setCurrentInvoice} from '../actions/invoices'

const InvoiceView = ({invoice, setCurrentInvoice}) => (
  <div className="container">
    <div className='row justify-content-center'>
      <button className='btn btn-primary col-sm-2' onClick={() => setCurrentInvoice(null)}>Close</button>
    </div>
    <div className='row justify-content-center'>
      <textarea rows="40" cols="100">{invoice}</textarea>
    </div>
  </div>
)

export default connect(
  (state) => ({
    invoice: state.currentInvoice,
  }),
  {setCurrentInvoice}
)(InvoiceView)
