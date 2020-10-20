import swal from 'sweetalert'
import {setData, loadingWrapper} from './common'

export const setUblInputValue = setData(['createInvoiceScreen', 'ublInput'])
export const setD16bInputValue = setData(['createInvoiceScreen', 'd16bInput'])
export const setGeneratedXmlInputValue = setData(['createInvoiceScreen', 'formGeneratedInput'])

const setInvoiceIds = setData(['invoicesScreen', 'ids'])

const setInvoice = (id, data) => ({
  type: 'SET INVOICES',
  path: ['invoices', id],
  payload: data,
  reducer: (state, data) => ({
    ...state,
    ...data,
  })
})

const setInvoices = (data) => ({
  type: 'SET INVOICES',
  path: ['invoices'],
  payload: data,
  reducer: (state, data) => ({
    ...state,
    ...data,
  })
})

export const updateInvoiceFormProperty = (property, data) => ({
  type: `UPDATE INVOICE FORM PROPERTY ${property}`,
  path: ['createInvoiceScreen', 'form', property],
  payload: data,
  reducer: (state, value) => value,
})

export const toggleFormatFilter = (format) => ({
  type: 'TOGGLE FORMAT FILTER',
  path: ['invoicesScreen', 'filters', 'formats', format],
  payload: null,
  reducer: (state) => !state,
})

export const getInvoices = () => loadingWrapper(
  async (dispatch, getState, {api}) => {
    const filters = getState().invoicesScreen.filters
    const formats = Object.keys(filters.formats).filter((k) => filters.formats[k])

    const invoices = await api.getInvoices(formats)
    dispatch(setInvoices(
      invoices.reduce((acc, val) => ({
        ...acc,
        [val.id]: val,
      }), {})
    ))

    dispatch(setInvoiceIds(
      invoices.reduce((acc, val) => ([
        ...acc,
        val.id,
      ]), []))
    )
  }
)

export const getInvoiceDetail = (id) => loadingWrapper(
  async (dispatch, getState, {api}) => {
    const invoiceDetail = await api.getInvoiceDetail(id)
    dispatch(setInvoice(id, {data: invoiceDetail}))
  }
)

export const getInvoiceMeta = (id) => loadingWrapper(
  async (dispatch, getState, {api}) => {
    const meta = await api.getInvoiceMeta(id)
    dispatch(setInvoice(id, meta))
  }
)

export const createInvoice = (data) => loadingWrapper(
  async (dispatch, getState, {api}) => {
    try {
      const invoice = await api.createInvoice(getState().user, data)
      dispatch(setInvoices({
        [invoice.id]: invoice
      }))

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

export const addAttachment = (attachment) => ({
  type: 'ADD ATTACHMENT',
  path: ['createInvoiceScreen', 'attachments'],
  payload: attachment,
  reducer: (state, attachment) => [...state, attachment],
})

export const removeAttachment = (attachment) => ({
  type: 'REMOVE ATTACHMENT',
  path: ['createInvoiceScreen', 'attachments'],
  payload: attachment,
  reducer: (state, attachment) => state.filter(a => a !== attachment),
})

export const clearAttachments = () => ({
  type: 'CLEAR ATTACHMENTS',
  path: ['createInvoiceScreen', 'attachments'],
  reducer: (state, attachment) => [],
})
