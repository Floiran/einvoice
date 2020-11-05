import {forwardReducerTo} from './utils/helpers'
import {invoiceFormats} from './utils/constants'
import {defaultUbl} from './data/defaultUbl'
import defaultD16b from './data/defaultD16b'

const getInitialState = () => ({
  invoices: {},
  // Count of running requests
  // If there is at least one running request show Loading Modal
  loadingRequests: 0,
  createInvoiceScreen: {
    ubl: {
      invoice: defaultUbl,
      attachments: [],
    },
    d16b: {
      invoice: defaultD16b,
      attachments: [],
    },
    generated: {
      invoice: '',
      attachments: [],
    },
    form: {
      sender: '',
      receiver: '',
      price: '',
      issueDate: new Date(),
      dueDate: new Date(),
      invoicePeriodStartDate: new Date(),
      invoicePeriodEndDate: new Date(),
      vat: '',
      invoiceId: '',
    },
  },
  invoicesScreen: {
    filters: {
      formats: {
        [invoiceFormats.UBL]: true,
        [invoiceFormats.D16B]: true,
      },
    },
  },
  // loggedUser can be in 3 possible states
  // 1. {unknown: true} - no user is logged
  // 2. {loading: true} - we are trying to request user data
  // 3. {id: 1, ...} - user is logged
  loggedUser: {
    // We start with loading user data
    loading: true,
  }
})

const rootReducer = (state = getInitialState(), action) => {
  if (action.reducer) {
    return forwardReducerTo(action.reducer, action.path)(state, action.payload)
  }

  return state
}

export default rootReducer
