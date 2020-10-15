import React from 'react'
import {Modal, Button} from 'react-bootstrap'

const _ConfirmationModal = ({title, text, confirm, cancel}) => (
  <div className="static-modal Modal" style={{cursor: 'default'}}>
    <Modal.Dialog>
      <Modal.Header style={{display: 'flex'}}>
        <Modal.Title style={{margin: 'auto'}}>{title}</Modal.Title>
      </Modal.Header>

      <Modal.Body>{text}</Modal.Body>

      <Modal.Footer>
        <Button onClick={cancel}>Cancel</Button>
        <Button variant="primary" onClick={confirm}>Confirm</Button>
      </Modal.Footer>
    </Modal.Dialog>
  </div>
)


export default class ConfirmationModal extends React.Component {
  state = {
    open: false,
    callback: null,
  }

  show = (callback) => () => {
    this.setState({
      open: true,
      callback,
    })
  }

  hide = () => this.setState({open: false, callback: null})

  confirm = (arg) => {
    this.state.callback(arg)
    this.hide()
  }

  render = () => (
    <React.Fragment>
      {this.props.children(this.show)}
      {this.state.open && (
        <_ConfirmationModal
          title={this.props.title}
          text={this.props.text}
          confirm={this.confirm}
          cancel={this.hide}
        />
      )}
    </React.Fragment>
  )
}
