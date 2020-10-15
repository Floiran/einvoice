import 'bootstrap/dist/css/bootstrap.css'
import './App.css'
import React from 'react'
import {connect} from 'react-redux'
import {branch, compose, lifecycle, renderComponent} from 'recompose'
import {Redirect, Route, Switch, withRouter} from 'react-router-dom'
import InvoiceList from './InvoiceList'
import LandingPage from './LandingPage'
import TopBar from './TopBar'
import CreateInvoice from './createInvoice'
import InvoiceView from './InvoiceView'
import {initializeApi} from '../actions/api'
import {getMyInfo, loginWithSlovenskoSkToken} from '../actions/users'
import AccountSettings from './AccountSettings'
import Auth from './helpers/Auth'
import LoadingModal from './helpers/LoadingModal'

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
        <Route exact path="/invoices/:id" component={InvoiceView} />
      </Switch>
    </div>
    {isLoading && <LoadingModal />}
  </div>
)

export default withRouter(
  compose(
    connect(
      (state) => ({
        apiInitialized: state.apiInitialized,
        isLoading: state.isLoading,
      }),
      {initializeApi, loginWithSlovenskoSkToken, getMyInfo}
    ),
    lifecycle({
      async componentDidMount() {
        await this.props.initializeApi()
        await this.props.getMyInfo()

        const urlParams = new URLSearchParams(this.props.location.search)

        if(urlParams.has('token')) {
          if (await this.props.loginWithSlovenskoSkToken(urlParams.get('token'))) {
            this.props.history.push('/account')
          }
        }
      },
    }),
    branch(
      ({apiInitialized}) => !apiInitialized,
      renderComponent(LoadingModal),
    ),
  )(App)
)
