import React from 'react'
import {connect} from 'react-redux'
import {compose} from 'recompose'
import {set} from 'object-path-immutable'
import {withTranslation} from 'react-i18next'
import {updateUser} from '../actions/users'

const EditableField = withTranslation('common')(
  ({label, valueField, inputField, change, toggleChange, name, save, t}) => (
    <div>
      <div className="row justify-content-center">
        {label}
        {change ?
          inputField
          :
          valueField
        }
      </div>
      <div className="row justify-content-center">
        {!change && <button className="btn btn-primary col-1" onClick={() => toggleChange(name)}>{t('edit')}</button>}
        {change && <button className="btn btn-primary col-1" onClick={() => save(name)}>{t('save')}</button>}
        {change && <button className="btn btn-primary col-1" onClick={() => toggleChange(name)}>{t('cancel')}</button>}
      </div>
    </div>
  )
)

class AccountSettings extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      email: {
        input: this.props.loggedUser.email,
        change: false,
      },
      serviceAccountKey: {
        input: this.props.loggedUser.serviceAccountKey,
        change: false,
      },
    }
  }

  toggleChange = (field) => {
    this.setState((oldState) => {
      const state = {...oldState}
      state[field].change = !state[field].change
      if (!state[field].change) state[field].input = this.props.loggedUser[field]
      return state
    })
  }

  handleInputChange = (event) => {
    const target = event.target
    const name = target.name
    this.setState((state) => set(state, [name, 'input'], target.value))
  }

  submit = async (field) => {
    await this.props.updateUser({
      [field]: this.state[field].input,
    })
    this.setState((state) => set(state, [field, 'change'], false))
  }

  render = () => (
    <div className="container">
      <h2 style={{textAlign: 'center'}}>{this.props.t('tabs.accountSettings')}</h2>
      <EditableField
        name="email"
        label={<p className="col-1">{this.props.t('common:email')}:</p>}
        valueField={<p className="col-3">{this.props.loggedUser.email}</p>}
        inputField={
          <input className="col-5" type="text" name="email" value={this.state.email.input} onChange={this.handleInputChange} />
        }
        toggleChange={this.toggleChange}
        save={this.submit}
        {...this.state.email}
      />
      <EditableField
        name="serviceAccountKey"
        label={<p className="col-2">{this.props.t('common:serviceAccountKey')}:</p>}
        valueField={
          <textarea
            readOnly
            className="col-7"
            name="serviceAccountKey"
            value={this.props.loggedUser.serviceAccountKey}
            rows="15"
          />
        }
        inputField={
          <textarea
            className="col-7"
            name="serviceAccountKey"
            value={this.state.serviceAccountKey.input}
            onChange={this.handleInputChange}
            rows="15"
          />
        }
        toggleChange={this.toggleChange}
        save={this.submit}
        {...this.state.serviceAccountKey}
      />
    </div>
  )
}

export default compose(
  connect(
    (state) => ({
      loggedUser: state.loggedUser,
    }),
    {updateUser},
  ),
  withTranslation(['TopBar', 'common']),
)(AccountSettings)
