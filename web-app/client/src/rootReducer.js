import {forwardReducerTo} from './utils/helpers'
import {invoiceFormats} from './utils/constants'
import {defaultUbl} from './data/defaultUbl'
import defaultD16b from './data/defaultD16b'

const getInitialState = () => ({
  invoices: {},
  serviceAccounts: [],
  isLoading: false,
  createInvoiceScreen: {
    ublInput: defaultUbl,
    d16bInput: defaultD16b,
    formGeneratedInput: '',
    form: {
      sender: '',
      receiver: '',
      price: '',
    },
    attachments: [],
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
