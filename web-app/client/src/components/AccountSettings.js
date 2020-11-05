import React, {useState} from 'react'
import {connect} from 'react-redux'
import {compose} from 'recompose'
import {Button, Card, Form} from 'react-bootstrap'
import {useTranslation, withTranslation} from 'react-i18next'
import {updateUser} from '../actions/users'

const EditableField = ({actualValue, label, save, ...props}) => {
  const {t} = useTranslation('common')
  const [isEditing, setEditing] = useState(false)
  const [value, setValue] = useState(actualValue)
  return (
    <Form.Group>
      <div style={{marginBottom: '5px'}}>
        <Form.Label>{label}</Form.Label>
        {!isEditing &&
        <Button variant="primary" size="sm" onClick={() => setEditing(true)}>
          {t('edit')}
        </Button>
        }
      </div>
      <Form.Control
        value={value}
        readOnly={!isEditing}
        onChange={(e) => setValue(e.target.value)}
        {...props}
      />
      {isEditing && <div style={{marginTop: '5px'}}>
        <Button
          variant="danger"
          size="sm"
          onClick={() => {setValue(actualValue); setEditing(false)}}
        >
          {t('cancel')}
        </Button>
        <Button
          variant="success"
          size="sm"
          onClick={() => {save(value); setEditing(false)}}
        >
          {t('save')}
        </Button>
      </div>}
    </Form.Group>
  )
}

const AccountSettings = ({loggedUser, t, updateUser}) => {
  return (
    <Card style={{margin: '15px'}}>
      <Card.Header className="bg-primary text-white text-center" as="h2">{t('tabs.accountSettings')}</Card.Header>
      <Card.Body>
        <EditableField
          actualValue={loggedUser.email}
          label={t('common:email')}
          save={(email) => updateUser({email})}
        />
        <EditableField
          actualValue={loggedUser.serviceAccountKey}
          label={t('common:serviceAccountKey')}
          save={(serviceAccountKey) => updateUser({serviceAccountKey})}
          as="textarea"
          rows={10}
        />
      </Card.Body>
    </Card>
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
