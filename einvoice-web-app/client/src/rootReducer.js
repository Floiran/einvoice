import {forwardReducerTo} from './utils/helpers'
import {invoiceFormats} from './utils/constants'
import {defaultUbl} from './data/defaultUbl'
import defaultD16b from './data/defaultD16b'

const getInitialState = () => ({
  user: null,
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
      }
    }
  },
  // TODO: get rid of this
  apiInitialized: false,
})

const rootReducer = (state = getInitialState(), action) => {
  if (action.reducer) {
    return forwardReducerTo(action.reducer, action.path)(state, action.payload)
  }

  return state
}

export default rootReducer
