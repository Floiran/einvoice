import React from 'react'
import {connect} from 'react-redux'
import {compose, withHandlers} from 'recompose'
import {Navbar} from 'react-bootstrap'
import {NavLink, withRouter} from 'react-router-dom'
import {LOGGING, LOGGING_FAILED, login, logout} from '../actions/users'

const TopBar = ({isLogged, login, logout, user, loggingStatus}) => (
  <Navbar bg="primary" variant="dark">
    <NavLink to="/">
      <Navbar.Brand>E-invoice</Navbar.Brand>
    </NavLink>
    <NavLink className="nav-link" to="/invoices">
      <Navbar.Text>All invoices</Navbar.Text>
    </NavLink>
    {
      isLogged &&
      <NavLink className="nav-link" to="/create-invoice">
        <Navbar.Text>Create invoice</Navbar.Text>
      </NavLink>
    }
    {
      isLogged &&
      <NavLink className="nav-link" to="/account">
        <Navbar.Text>Account settings</Navbar.Text>
      </NavLink>
    }
    <Navbar.Collapse className="justify-content-end">
      { loggingStatus === LOGGING_FAILED &&
      <p>Logging failed</p>
      }
      {isLogged ?
        <React.Fragment>
          <NavLink to="/account">
            <Navbar.Text>
              User name: {user.name}
            </Navbar.Text>
          </NavLink>
          <button className="btn btn-danger" onClick={logout}>Logout</button>
        </React.Fragment>
        :
        <button className="btn btn-success" onClick={login}>
          { loggingStatus === LOGGING ?
            "Logging in..."
            :
            "Login"
          }
        </button>
      }
    </Navbar.Collapse>
  </Navbar>
)

export default withRouter(
  compose(
    connect(
      (state) => ({
        loggingStatus: state.loggingStatus,
        isLogged: !!state.user,
        user: state.user
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
