import React from 'react'
import {connect} from 'react-redux'
import {compose} from 'recompose'
import {Button, Table} from 'react-bootstrap'
import {withTranslation} from 'react-i18next'
import {addAttachment, removeAttachment} from "../../actions/invoices";

const Attachments = ({attachments, addAttachment, removeAttachment, t}) => {
  const hiddenFileInput = React.useRef(null);

  const handleClick = () => {
    hiddenFileInput.current.click();
  };

  const handleChange = event => {
    const fileUploaded = event.target.files[0];
    if(fileUploaded) {
      addAttachment(fileUploaded)
    }
  };

  return (
    <div>
      {attachments.length > 0 && <Table>
        <thead>
        <tr>
          <th><p>{t('name')}</p></th>
          <th></th>
        </tr>
        </thead>
        <tbody>
        {
          attachments.map((a, i) =>
            <tr key={i}>
              <td><p>{a.name}</p></td>
              <td><p className={'link'} onClick={() => removeAttachment(a)}>{t('remove')}</p></td>
            </tr>
          )
        }
        </tbody>
      </Table>
      }
      <div className="row justify-content-center">
        <Button className="col-2" onClick={handleClick} >
          {t('add')}
        </Button>
        <input type="file"
               ref={hiddenFileInput}
               onChange={handleChange}
               style={{display:'none'}} />
      </div>
    </div>
  )
}

export default compose(
  connect(
    (state) => ({
      attachments: state.createInvoiceScreen.attachments,
    }),
    {addAttachment, removeAttachment}
  ),
  withTranslation(['common']),
)(Attachments)
