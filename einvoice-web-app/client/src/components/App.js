import './App.css'
import React from 'react'
import {connect} from 'react-redux'
import {branch, compose, lifecycle, renderComponent, renderNothing} from 'recompose'
import CreateInvoice from './CreateInvoice'
import InvoiceList from './InvoiceList'
import Login from './Login'
import InvoiceView from './InvoiceView'
import {initializeApi} from '../actions/api'
import {getMyInfo, setUser, logout} from '../actions/users'
import {getInvoices, toggleCreatingInvoice} from '../actions/invoices'

const App = ({currentInvoice, isCreatingInvoice, logout, user, toggleCreatingInvoice}) => (
  <div className="App container">
    <header className="App-header row">
      <h1 className='col'>E-invoice</h1>
    </header>
    <div className='row justify-content-center'>
      <p className='col-sm-1'>User id:</p>
      <p className='col-sm-1'>{user.id}</p>
      <button className='btn btn-primary' onClick={logout}>Logout</button>
    </div>
    { !currentInvoice &&
    <div>
      <div className='row justify-content-center'>
        <button className='btn btn-primary' onClick={toggleCreatingInvoice}>Create invoice</button>
      </div>
      {isCreatingInvoice && <CreateInvoice />}
      <InvoiceList />
    </div>
    }
    { currentInvoice && <InvoiceView /> }
  </div>
)

export default compose(
  connect(
    (state) => ({
      user: state.user,
      currentInvoice: state.currentInvoice,
      apiInitialized: state.apiInitialized,
      isCreatingInvoice: state.isCreatingInvoice,
    }),
    {
      initializeApi, getMyInfo, getInvoices, setUser, logout, toggleCreatingInvoice
    }
  ),
  lifecycle({
    async componentDidMount() {
      await this.props.initializeApi()
      await this.props.getMyInfo()
    },
  }),
  branch(
    ({apiInitialized}) => !apiInitialized,
    renderNothing,
  ),
  branch(
    ({user}) => !user,
    renderComponent(Login),
  ),
)(App)
