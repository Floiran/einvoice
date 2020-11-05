import {connect} from 'react-redux'
import {branch, compose, renderComponent} from 'recompose'
import Unauthorized from './Unauthorized'

export default compose(
  connect(
    (state) => ({
      isLogged: state.loggedUser.id != null,
    })
  ),
  branch(
    ({isLogged}) => !isLogged,
    renderComponent(Unauthorized),
  ),
)
