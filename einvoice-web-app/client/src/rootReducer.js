import {forwardReducerTo} from "./utils/helpers"

export const getInitialState = () => ({
  user: null,
  invoices: [],
  serviceAccounts: [],
  // TODO: be smarter and create cache
  currentInvoice: null,
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
