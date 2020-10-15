import './LoadingModal.css'
import React from 'react'
import {Modal} from 'react-bootstrap'

export default () => (
  <div className="static-modal Modal">
    <Modal.Dialog>
      <Modal.Header style={{display: 'flex', backgroundColor: '#f3f3f3'}}>
        <Modal.Title style={{margin: 'auto'}}>PROCESSING REQUEST</Modal.Title>
      </Modal.Header>

      <Modal.Body style={{display: 'flex', backgroundColor: '#f3f3f3'}}>
        <div style={{margin: 'auto'}} className="loader" />
      </Modal.Body>
    </Modal.Dialog>
  </div>
)
