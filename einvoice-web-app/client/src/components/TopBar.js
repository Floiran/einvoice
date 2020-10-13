import React from 'react'
import {connect} from 'react-redux'
import {compose, withHandlers} from 'recompose'
import {Navbar} from 'react-bootstrap'
import {NavLink, withRouter} from 'react-router-dom'
import lodash from 'lodash'
import {login, logout} from '../actions/users'

const TopBar = ({isLogged, login, logout, userId}) => (
  <Navbar bg="primary" variant="dark">
    <NavLink to="/">
      <Navbar.Brand>E-invoice</Navbar.Brand>
    </NavLink>
    <NavLink to="/invoices">
      <Navbar.Text>All invoices</Navbar.Text>
    </NavLink>
    <Navbar.Collapse className="justify-content-end">
      {isLogged ?
        <React.Fragment>
          <NavLink to="/account">
            <Navbar.Text>
              User id: {userId}
            </Navbar.Text>
          </NavLink>
          <button className="btn btn-danger" onClick={logout}>Logout</button>
        </React.Fragment>
        :
        <button className="btn btn-success" onClick={login}>Login</button>
      }
    </Navbar.Collapse>
  </Navbar>
)

export default withRouter(
  compose(
    connect(
      (state) => ({
        isLogged: !!state.user,
        userId: lodash.get(state, ['user', 'id'])
      }),
      {login, logout}
    ),
    withHandlers({
      login: ({history, login}) => async () => {
        await login()
        history.push('/account')
      },
      logout: ({history, logout}) => async () => {
        await logout()
        history.push('/')
      },
    })
  )(TopBar)
)
