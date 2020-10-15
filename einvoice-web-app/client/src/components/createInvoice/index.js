import React from 'react'
import {Route, NavLink} from 'react-router-dom'
import {Button} from 'react-bootstrap'
import InvoiceForm from './InvoiceForm'
import {CreateUbl, CreateD16b, CreateGenerated} from './CreateXml'

export default () => (
  <div className="container">
    <h2 style={{textAlign: 'center'}}>Create invoice</h2>
    <div className="row justify-content-center">
      <NavLink to="/create-invoice/form" activeClassName="selected">
        <Button variant="primary" size="lg">Form</Button>
      </NavLink>
      <NavLink to="/create-invoice/generated" activeClassName="selected">
        <Button variant="primary" size="lg">Generated</Button>
      </NavLink>
      <NavLink to="/create-invoice/ubl" activeClassName="selected">
        <Button variant="primary" size="lg">UBL2.1</Button>
      </NavLink>
      <NavLink to="/create-invoice/d16b" activeClassName="selected">
        <Button variant="primary" size="lg">D16B</Button>
      </NavLink>
    </div>
    <Route exact path="/create-invoice/form" component={InvoiceForm} />
    <Route exact path="/create-invoice/generated" component={CreateGenerated} />
    <Route exact path="/create-invoice/ubl" component={CreateUbl} />
    <Route exact path="/create-invoice/d16b" component={CreateD16b} />
  </div>
)
