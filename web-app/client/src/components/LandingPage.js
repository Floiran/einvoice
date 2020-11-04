import React from 'react'
import {withTranslation} from 'react-i18next'

const LandingPage = ({t}) => (
  <div style={{textAlign: 'center'}}>
    <h1>{t('title')}</h1>
    <h2>{t('description')}</h2>
  </div>
)

export default withTranslation('LandingPage')(LandingPage)
