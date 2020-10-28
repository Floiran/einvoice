import React from 'react'
import {connect} from 'react-redux'
import {compose} from 'recompose'
import {Accordion, Button, Card, FormCheck} from 'react-bootstrap'
import {withTranslation} from 'react-i18next'
import {invoiceFormats} from '../../utils/constants'
import {getInvoices, toggleFormatFilter} from '../../actions/invoices'

const Filters = ({formats, getInvoices, searchEnabled, t, toggleFormatFilter}) => (
  <Accordion>
    <Card style={{textAlign: 'left'}}>
      <Accordion.Toggle
        as={Card.Header}
        eventKey="0"
        className="bg-primary text-white"
        style={{cursor: 'pointer'}}
      >
        {t('filters')}
      </Accordion.Toggle>
      <Accordion.Collapse eventKey="0">
        <Card.Body>
          <div>
            <strong style={{textDecoration: 'underline', fontSize: '20px'}}>{t('format')}</strong>
            <div style={{display: 'flex'}}>
              {Object.values(invoiceFormats).map((format) => (
                <FormCheck
                  type="checkbox"
                  key={format}
                  checked={formats[format]}
                  onChange={() => toggleFormatFilter(format)}
                  label={format}
                  style={{marginRight: '5px'}}
                />
              ))}
            </div>
          </div>
          <div style={{display: 'flex'}}>
            <Button
              variant="primary"
              style={{marginLeft: 'auto'}}
              onClick={getInvoices}
              disabled={!searchEnabled}
            >
              {t('search')}
            </Button>
          </div>
        </Card.Body>
      </Accordion.Collapse>
    </Card>
  </Accordion>
)

export default compose(
  connect(
    (state) => {
      const filters = state.invoicesScreen.filters
      return {
        formats: filters.formats,
        searchEnabled: Object.values(filters.formats).reduce((acc, v) => acc || v, false),
      }
    },
    {getInvoices, toggleFormatFilter}
  ),
  withTranslation('common')
)(Filters)
