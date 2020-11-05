import React from 'react'
import {connect} from 'react-redux'
import {branch, compose, lifecycle, renderNothing} from 'recompose'
import {NavLink} from 'react-router-dom'
import {Table} from 'react-bootstrap'
import {withTranslation} from 'react-i18next'
import Filters from './Filters'
import {getInvoices} from '../../actions/invoices'

const Index = ({invoices, invoiceIds, t}) => (
  <div className="container" style={{textAlign: 'center'}}>
    <h2>{t('TopBar:tabs.allInvoices')}</h2>
    <Filters />
    <Table striped hover responsive size="sm">
      <thead>
        <tr>
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
    </Table>
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
