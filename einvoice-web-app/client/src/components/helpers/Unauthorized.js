import React from 'react'
import {withTranslation} from 'react-i18next'

export default withTranslation('helpers')(
  ({t}) => (
    <div>
      <h1>401</h1>
      <h2>{t('auth.title')}</h2>
      <div>{t('auth.description')}</div>
    </div>
  )
)
