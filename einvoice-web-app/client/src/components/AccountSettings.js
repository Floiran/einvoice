import React from 'react'
import {connect} from 'react-redux'
import {updateUser} from "../actions/users";
import {set} from 'object-path-immutable'

const EditableField = ({label, valueField, inputField, change, toggleChange, name, save}) => (
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
      {!change && <button className='btn btn-primary col-1' onClick={() => toggleChange(name)}>Edit</button>}
      {change && <button className='btn btn-primary col-1' onClick={() => save(name)}>Save</button>}
      {change && <button className='btn btn-primary col-1' onClick={() => toggleChange(name)}>Cancel</button>}
    </div>
  </div>
)

class AccountSettings extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      email: {
        input: this.props.user.email,
        change: false
      },
      serviceAccountKey: {
        input: this.props.user.serviceAccountKey,
        change: false
      },
    }
  }

  toggleChange = (field) => {
    this.setState( oldState => {
      let state = {...oldState}
      state[field].change = !state[field].change
      if(!state[field].change) state[field].input = this.props.user[field]
      return state
    })
  }

  handleInputChange = (event) => {
    const target = event.target
    const name = target.name
    this.setState( oldState => set(oldState, [name, 'input'], target.value))
  }

  submit = async (field) => {
    await this.props.updateUser( {
      [field]: this.state[field].input
    })
    this.setState( oldState => set(oldState, [field, 'change'], false))
  }

  render() {
    return (
      <div className="container">
        <h2 style={{textAlign: 'center'}}>Settings</h2>
        <EditableField
          name={"email"}
          label={<p className="col-1">Email:</p>}
          valueField={<p className="col-3">{this.props.user.email}</p>}
          inputField={<input className="col-5" type="text" name="email" value={this.state.email.input}
                             onChange={this.handleInputChange}/>}
          toggleChange={this.toggleChange}
          save={this.submit}
          {...this.state.email}
        />
        <EditableField
          name={"serviceAccountKey"}
          label={<p className="col-2">Service account key:</p>}
          valueField={<textarea readOnly className="col-7" name="serviceAccountKey" value={this.props.user.serviceAccountKey}
                                rows="15"/>}
          inputField={<textarea className="col-7" name="serviceAccountKey" value={this.state.serviceAccountKey.input}
                                onChange={this.handleInputChange}  rows="15"/>}
          toggleChange={this.toggleChange}
          save={this.submit}
          {...this.state.serviceAccountKey}
        />
      </div>
    )
  }
}

export default connect(
  (state) => ({
    user: state.user
  }),
  {updateUser}
)(AccountSettings)
