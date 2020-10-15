import {connect} from 'react-redux'
import {branch, compose,  renderComponent} from 'recompose'
import Unauthorized from './Unauthorized'
import LoadingModal from './LoadingModal'

export default compose(
  connect(
    (state) => ({
      user: state.user
    })
  ),
  branch(
    ({user}) => user == null,
    renderComponent(LoadingModal),
  ),
  branch(
    ({user}) => user.unauthorized,
    renderComponent(Unauthorized),
  ),
)
