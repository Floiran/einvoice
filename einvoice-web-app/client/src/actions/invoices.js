import swal from 'sweetalert'
import {setData, loadingWrapper} from './common'

const setInvoices = setData(['invoices'])
export const setCurrentInvoice = setData(['currentInvoice'])
export const setUblInputValue = setData(['createInvoiceScreen', 'ublInput'])
export const setD16bInputValue = setData(['createInvoiceScreen', 'd16bInput'])
export const setGeneratedXmlInputValue = setData(['createInvoiceScreen', 'formGeneratedInput'])

const addInvoice = (invoice) => ({
  type: 'ADD INVOICE',
  path: ['invoices'],
  payload: invoice,
  reducer: (state, invoice) => [...state, invoice],
})

export const updateInvoiceFormProperty = (property, data) => ({
  type: `UPDATE INVOICE FORM PROPERTY ${property}`,
  path: ['createInvoiceScreen', 'form', property],
  payload: data,
  reducer: (state, value) => value,
})

export const getInvoices = () => loadingWrapper(
  async (dispatch, getState, {api}) => {
    const invoices = await api.getInvoices()
    dispatch(setInvoices(invoices))
  }
)

export const getInvoiceDetail = (id) => loadingWrapper(
  async (dispatch, getState, {api}) => {
    const invoiceDetail = await api.getInvoiceDetail(id)
    dispatch(setCurrentInvoice(invoiceDetail))
  }
)

export const createInvoice = (format, data) => loadingWrapper(
  async (dispatch, getState, {api}) => {
    try {
      const invoice = await api.createInvoice(getState().user, format, data)
      dispatch(addInvoice(invoice))
      await swal({
        title: 'Invoice was created',
        icon: 'success',
      })
      return invoice.id
    } catch(error) {
      await swal({
        title: 'Invoice could not be created',
        text: error.message,
        icon: 'error',
      })
    }
  }
)
