import React from 'react'
import {connect} from 'react-redux'
import {compose, withHandlers} from 'recompose'
import {Navbar} from 'react-bootstrap'
import {NavLink, withRouter} from 'react-router-dom'
import {LOGGING, LOGGING_FAILED, logout} from '../actions/users'

const TopBar = ({isLogged, slovenskoSkUrl, logout, user, loggingStatus}) => (
  <Navbar bg="primary" variant="dark">
    <NavLink to="/">
      <Navbar.Brand>E-invoice</Navbar.Brand>
    </NavLink>
    <NavLink className="nav-link" to="/invoices">
      <Navbar.Text>All invoices</Navbar.Text>
    </NavLink>
    {
      isLogged && <React.Fragment>
        <NavLink className="nav-link" to="/create-invoice">
          <Navbar.Text>Create invoice</Navbar.Text>
        </NavLink>
        <NavLink className="nav-link" to="/account">
          <Navbar.Text>Account settings</Navbar.Text>
        </NavLink>
      </React.Fragment>
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
        <a href={slovenskoSkUrl}>
          <button className="btn btn-success">
            { loggingStatus === LOGGING ?
              "Logging in..."
              :
              "Login"
            }
          </button>
        </a>
      }
    </Navbar.Collapse>
  </Navbar>
)

export default withRouter(
  compose(
    connect(
      (state) => ({
        loggingStatus: state.loggingStatus,
        isLogged: state.user && !state.user.unauthorized,
        user: state.user,
        slovenskoSkUrl: state.slovenskoSkUrl,
      }),
      {logout}
    ),
    withHandlers({
      logout: ({history, logout}) => async () => {
        await logout()
        history.push('/')
      },
    })
  )(TopBar)
)
