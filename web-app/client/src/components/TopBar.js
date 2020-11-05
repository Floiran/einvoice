import React from 'react'
import {connect} from 'react-redux'
import {compose, withHandlers} from 'recompose'
import {Navbar} from 'react-bootstrap'
import {NavLink} from 'react-router-dom'
import {withTranslation} from 'react-i18next'
import {CONFIG} from '../appSettings'
import {logout} from '../actions/users'
import {setLoadingState} from '../actions/common'

const TopBar = ({i18n, isLogged, loggedUser, logout, startLoading, t}) => (
  <Navbar bg="primary" variant="dark">
    <NavLink to="/">
      <Navbar.Brand>{t('title')}</Navbar.Brand>
    </NavLink>
    <button onClick={() => i18n.changeLanguage('sk')}>sk</button>
    <button onClick={() => i18n.changeLanguage('en')}>en</button>
    <NavLink className="nav-link" to="/invoices">
      <Navbar.Text>{t('tabs.allInvoices')}</Navbar.Text>
    </NavLink>
    <Navbar.Collapse className="justify-content-end">
      {isLogged ?
        <React.Fragment>
          <NavLink className="nav-link" to="/create-invoice">
            <Navbar.Text>{t('tabs.createInvoice')}</Navbar.Text>
          </NavLink>
          <NavLink className="nav-link" to="/account">
            <Navbar.Text>{t('tabs.accountSettings')}</Navbar.Text>
          </NavLink>
          <NavLink to="/account">
            <Navbar.Text>{loggedUser.name}</Navbar.Text>
          </NavLink>
          <button className="btn btn-danger" onClick={logout}>{t('logout')}</button>
        </React.Fragment>
        :
        <a href={CONFIG.slovenskoSkLoginUrl}>
          <button className="btn btn-success" onClick={startLoading}>
            {t('login')}
          </button>
        </a>
      }
    </Navbar.Collapse>
  </Navbar>
)

export default compose(
  connect(
    (state) => ({
      isLogged: state.loggedUser.id != null,
      loggedUser: state.loggedUser,
    }),
    {logout, setLoadingState}
  ),
  withHandlers({
    logout: ({history, logout}) => async () => {
      await logout()
      history.push('/')
    },
    startLoading: ({setLoadingState}) => () => setLoadingState(true)
  }),
  withTranslation('TopBar'),
)(TopBar)
