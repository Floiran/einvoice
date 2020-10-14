import 'bootstrap/dist/css/bootstrap.css'
import './App.css'
import React from 'react'
import {connect} from 'react-redux'
import {branch, compose, lifecycle, renderNothing} from 'recompose'
import {Route, withRouter} from 'react-router-dom'
import InvoiceList from './InvoiceList'
import LandingPage from './LandingPage'
import TopBar from './TopBar'
import CreateInvoice from './CreateInvoice'
import InvoiceView from './InvoiceView'
import {initializeApi} from '../actions/api'
import {getMyInfo} from '../actions/users'
import AccountSettings from "./AccountSettings";

const App = () => (
  <div>
    <TopBar />
    <div className="App container">
      <Route exact path="/" component={LandingPage} />
      <Route path="/account" component={AccountSettings} />
      <Route path="/create-invoice" component={CreateInvoice} />
      <Route exact path="/invoices" component={InvoiceList} />
      <Route exact path="/invoices/:id" component={InvoiceView} />
    </div>
  </div>
)

export default withRouter(
  compose(
    connect(
      (state) => ({
        apiInitialized: state.apiInitialized,
      }),
      {initializeApi, getMyInfo}
    ),
    lifecycle({
      async componentDidMount() {
        await this.props.initializeApi()
        await this.props.getMyInfo(this.props.location, this.props.history)
      },
    }),
    branch(
      ({apiInitialized}) => !apiInitialized,
      // TODO: loading component
      renderNothing,
    ),
  )(App)
)
