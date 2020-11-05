import 'bootstrap/dist/css/bootstrap.css'
import React from 'react'
import {connect} from 'react-redux'
import {compose, lifecycle} from 'recompose'
import {Route, Switch} from 'react-router-dom'
import App from './App'
import LoadingModal from './helpers/LoadingModal'
import {loginWithSlovenskoSkToken} from '../actions/users'
import {registerLocale} from 'react-datepicker'
import sk from 'date-fns/locale/sk'
// Load slovak translations for time
registerLocale('sk', sk)

const LoginCallback = compose(
  connect(
    null,
    {loginWithSlovenskoSkToken}
  ),
  lifecycle({
    async componentDidMount() {
      const urlParams = new URLSearchParams(this.props.location.search)
      if (await this.props.loginWithSlovenskoSkToken(urlParams.get('token'))) {
        this.props.history.push('/account')
      } else {
        this.props.history.push('/')
      }
    },
  })
)(LoadingModal)

export default () => (
  <Switch>
    <Route path="/login-callback" component={LoginCallback} />
    <Route component={App} />
  </Switch>
)
