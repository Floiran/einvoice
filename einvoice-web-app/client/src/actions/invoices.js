const setInvoices = (invoices) => ({
  type: 'SET INVOICES',
  path: ['invoices'],
  payload: invoices,
  reducer: (state, invoices) => invoices,
})

export const setCurrentInvoice = (invoice) => ({
  type: 'SET CURRENT INVOICE',
  path: ['currentInvoice'],
  payload: invoice,
  reducer: (state, invoice) => invoice,
})

export const addInvoice = (invoice) => ({
  type: 'ADD INVOICE',
  path: ['invoices'],
  payload: invoice,
  reducer: (state, invoice) => [...state, invoice],
})

export const toggleCreatingInvoice = () => ({
  type: 'TOGGLE INVOICE CREATION',
  path: ['isCreatingInvoice'],
  payload: null,
  reducer: (state) => !state,
})

export const getInvoices = () => (
  async (dispatch, getState, {api}) => {
    const invoices = await api.getInvoices(getState().user)
    dispatch(setInvoices(invoices))
  }
)

export const getInvoiceDetail = (id) => (
  async (dispatch, getState, {api}) => {
    const invoiceDetail = await api.getInvoiceDetail(getState().user, id)
    dispatch(setCurrentInvoice(invoiceDetail))
  }
)

export const createInvoice = (format, data) => (
  async (dispatch, getState, {api}) => {
    const invoice = await api.createInvoice(getState().user, format, data)
    dispatch(addInvoice(invoice))
  }
)
