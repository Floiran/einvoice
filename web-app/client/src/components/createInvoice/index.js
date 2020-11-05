import React from 'react'
import {NavLink, Route, Switch} from 'react-router-dom'
import {Button} from 'react-bootstrap'
import InvoiceForm from './invoiceForm'
import CreateInvoice from './CreateInvoice'
import {useTranslation} from 'react-i18next'

export default () => {
  const {t} = useTranslation(['invoices', 'TopBar'])
  return (
    <div className="container">
      <h2 style={{textAlign: 'center'}}>{t('TopBar:tabs.createInvoice')}</h2>
      <div className="row justify-content-center">
        <NavLink to="/create-invoice/form" activeClassName="selected">
          <Button variant="primary" size="lg">{t('form')}</Button>
        </NavLink>
        <NavLink to="/create-invoice/generated" activeClassName="selected">
          <Button variant="primary" size="lg">{t('generated')}</Button>
        </NavLink>
        <NavLink to="/create-invoice/ubl" activeClassName="selected">
          <Button variant="primary" size="lg">UBL2.1</Button>
        </NavLink>
        <NavLink to="/create-invoice/d16b" activeClassName="selected">
          <Button variant="primary" size="lg">D16B</Button>
        </NavLink>
      </div>
      <Switch>
        <Route exact path="/create-invoice/form" component={InvoiceForm} />
        <Route path="/create-invoice" component={CreateInvoice} />
      </Switch>
    </div>
  )
}
