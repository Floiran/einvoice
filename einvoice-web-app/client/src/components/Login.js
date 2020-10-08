import React from 'react'
import {connect} from 'react-redux'
import {login} from '../actions/users'

const Login = ({login}) => (
  <div className="App container">
    <header className="App-header row">
      <h1 className='col'>E-invoice</h1>
    </header>
    <div className='row justify-content-center'>
      <button className='btn btn-primary col-sm-2' onClick={login}>Login</button>
    </div>
  </div>
)

export default connect(
  null,
  {login}
)(Login)
