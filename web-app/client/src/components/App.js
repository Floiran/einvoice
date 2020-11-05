import './App.css'
import React from 'react'
import {connect} from 'react-redux'
import {branch, compose, lifecycle, renderComponent} from 'recompose'
import {Redirect, Route, Switch} from 'react-router-dom'
import InvoiceList from './invoiceList'
import LandingPage from './LandingPage'
import TopBar from './TopBar'
import CreateInvoice from './createInvoice'
import InvoiceView from './InvoiceView'
import {getMyInfo} from '../actions/users'
import AccountSettings from './AccountSettings'
import Auth from './helpers/Auth'
import LoadingModal from './helpers/LoadingModal'
import NotFound from './helpers/NotFound'

const App = ({isLoading}) => (
  <div>
    <TopBar />
    <div className="container">
      <Switch>
        <Route exact path="/" component={LandingPage} />
        <Route path="/account" component={Auth(AccountSettings)} />
        <Redirect exact from="/create-invoice" to="/create-invoice/form" />
        <Route path="/create-invoice" component={Auth(CreateInvoice)} />
        <Route exact path="/invoices" component={InvoiceList} />
        <Route exact path="/invoices/:id([0-9]+)" component={InvoiceView} />
        <Route component={NotFound} />
      </Switch>
    </div>
    {isLoading && <LoadingModal />}
  </div>
)

export default compose(
  connect(
    (state) => ({
      isLoading: state.loadingRequests > 0,
      loggedUser: state.loggedUser,
    }),
    {getMyInfo}
  ),
  lifecycle({
    async componentDidMount() {
      // try to get user only if not already logged
      if (this.props.loggedUser.id == null) {
        await this.props.getMyInfo()
      }
    },
  }),
  branch(
    ({loggedUser}) => loggedUser.loading,
    renderComponent(LoadingModal),
  ),
)(App)
