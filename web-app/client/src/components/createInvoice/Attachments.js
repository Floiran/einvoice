import React from 'react'
import {Button, Table} from 'react-bootstrap'
import {useTranslation} from 'react-i18next'

export default ({attachments, addAttachment, removeAttachment}) => {
  const {t} = useTranslation(['common', 'invoices'])
  const hiddenFileInput = React.useRef(null)

  const handleClick = () => {
    hiddenFileInput.current.click()
  }

  const handleChange = (event) => {
    const fileUploaded = event.target.files[0]
    if (fileUploaded) {
      addAttachment(fileUploaded)
    }
  }

  return (
    <div>
      <h4>{t('invoices:attachments')}</h4>
      {attachments.length > 0 &&
        <Table bordered>
          <thead>
            <tr>
              <th><p>{t('name')}</p></th>
              <th />
            </tr>
          </thead>
          <tbody>
            {attachments.map((a, i) => (
              <tr key={i}>
                <td><p>{a.name}</p></td>
                <td>
                  <Button variant="danger" onClick={() => removeAttachment(a)}>{t('remove')}</Button>
                </td>
              </tr>
            ))}
          </tbody>
        </Table>
      }
      <div>
        <Button onClick={handleClick}>
          {t('add')}
        </Button>
        <input
          type="file"
          ref={hiddenFileInput}
          onChange={handleChange}
          style={{display: 'none'}}
        />
      </div>
    </div>
  )
}
