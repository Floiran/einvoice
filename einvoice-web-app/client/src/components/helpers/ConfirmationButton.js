import React from 'react'
import ConfirmationModal from './ConfirmationModal'

export default ({confirmationTitle, confirmationText, onClick, children, ...props}) => (
  <ConfirmationModal
    title={confirmationTitle}
    text={confirmationText}
  >
    {(confirm) => (
      <button onClick={confirm(onClick)} {...props}>
        {children}
      </button>
    )}
  </ConfirmationModal>
)
