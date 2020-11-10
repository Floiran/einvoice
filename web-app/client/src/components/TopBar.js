import React from 'react'
import {connect} from 'react-redux'
import {compose, withHandlers} from 'recompose'
import {Dropdown, DropdownButton, Navbar} from 'react-bootstrap'
import {NavLink, withRouter} from 'react-router-dom'
import {withTranslation} from 'react-i18next'
import {CONFIG} from '../appSettings'
import {logout} from '../actions/users'
import {updateRunningRequests} from '../actions/common'

const TopBar = ({i18n, isLogged, loggedUser, logout, startLoading, t}) => (
  <Navbar bg="primary" variant="dark">
    <NavLink to="/">
      <Navbar.Brand>{t('title')}</Navbar.Brand>
    </NavLink>
    <DropdownButton size="sm" title={i18n.language.toUpperCase()} variant="secondary">
      <Dropdown.Item active={i18n.language === 'sk'} onClick={() => i18n.changeLanguage('sk')}>
        SK
      </Dropdown.Item>
      <Dropdown.Item active={i18n.language === 'en'} onClick={() => i18n.changeLanguage('en')}>
        EN
      </Dropdown.Item>
    </DropdownButton>
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

export default withRouter(
  compose(
    connect(
      (state) => ({
        isLogged: state.loggedUser.id != null,
        loggedUser: state.loggedUser,
      }),
      {logout, updateRunningRequests}
    ),
    withHandlers({
      logout: ({logout, history}) => async () => {
        await logout()
        history.push('/')
      },
      startLoading: ({updateRunningRequests}) => () => updateRunningRequests(1)
    }),
    withTranslation('TopBar'),
  )(TopBar)
)
