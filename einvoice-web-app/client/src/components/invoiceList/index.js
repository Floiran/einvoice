import React from 'react'
import {connect} from 'react-redux'
import {branch, compose, lifecycle, renderNothing} from 'recompose'
import {NavLink} from 'react-router-dom'
import {withTranslation} from 'react-i18next'
import Filters from './Filters'
import {getInvoices} from '../../actions/invoices'

const Index = ({invoices, invoiceIds, t})  => (
  <div className="container" style={{textAlign: 'center'}}>
    <h2>{t('TopBar:tabs.allInvoices')}</h2>
    <Filters />
    <table className="table table-striped table-borderless table-hover">
      <thead>
      <tr>
        <th>#</th>
        <th>ID</th>
        <th>{t('sender')}</th>
        <th>{t('receiver')}</th>
        <th>{t('price')}</th>
        <th>{t('format')}</th>
        <th />
      </tr>
      </thead>
      <tbody>
        {invoiceIds.map((invoiceId, i) => (
          <tr key={i}>
            <th scope="row"><p>{i+1}</p></th>
            <td><p>{invoiceId}</p></td>
            <td><p>{invoices[invoiceId].sender}</p></td>
            <td><p>{invoices[invoiceId].receiver}</p></td>
            <td><p>{invoices[invoiceId].price}</p></td>
            <td><p>{invoices[invoiceId].format}</p></td>
            <td><p>
              <NavLink to={`/invoices/${invoiceId}`}>detail</NavLink>
            </p></td>
          </tr>
        ))}
      </tbody>
    </table>
  </div>
)

export default compose(
  connect(
    (state) => ({
      invoices: state.invoices,
      invoiceIds: state.invoicesScreen.ids,
    }),
    {getInvoices}
  ),
  lifecycle({
    componentDidMount() {
      this.props.getInvoices()
    },
  }),
  branch(
    ({invoiceIds}) => invoiceIds == null,
    renderNothing,
  ),
  withTranslation(['common', 'TopBar']),
)(Index)
