import React from 'react'
import {connect} from 'react-redux'
import {compose, withHandlers} from 'recompose'
import {Navbar} from 'react-bootstrap'
import {NavLink, withRouter} from 'react-router-dom'
import {withTranslation} from 'react-i18next'
import {LOGGING, LOGGING_FAILED, logout} from '../actions/users'

const TopBar = ({i18n, isLogged, loggingStatus, logout, slovenskoSkLoginUrl, t, user}) => (
  <Navbar bg="primary" variant="dark">
    <NavLink to="/">
      <Navbar.Brand>{t('title')}</Navbar.Brand>
    </NavLink>
    <button onClick={() => i18n.changeLanguage('sk')}>sk</button>
    <button onClick={() => i18n.changeLanguage('en')}>en</button>
    <NavLink className="nav-link" to="/invoices">
      <Navbar.Text>{t('tabs.allInvoices')}</Navbar.Text>
    </NavLink>
    {
      isLogged && <React.Fragment>
        <NavLink className="nav-link" to="/create-invoice">
          <Navbar.Text>{t('tabs.createInvoice')}</Navbar.Text>
        </NavLink>
        <NavLink className="nav-link" to="/account">
          <Navbar.Text>{t('tabs.accountSettings')}</Navbar.Text>
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
            <Navbar.Text>{user.name}</Navbar.Text>
          </NavLink>
          <button className="btn btn-danger" onClick={logout}>{t('logout')}</button>
        </React.Fragment>
        :
        <a href={slovenskoSkLoginUrl}>
          <button className="btn btn-success">
            { loggingStatus === LOGGING ?
              t('logging')
              :
              t('login')
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
        slovenskoSkLoginUrl: state.urls.slovenskoSkLoginUrl,
      }),
      {logout}
    ),
    withHandlers({
      logout: ({history, logout}) => async () => {
        await logout()
        history.push('/')
      },
    }),
    withTranslation('TopBar'),
  )(TopBar)
)
