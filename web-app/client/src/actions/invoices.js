import swal from 'sweetalert'
import {setData, loadingWrapper} from './common'

export const setCreateInvoiceValue = (tab) => setData(['createInvoiceScreen', tab, 'invoice'])

const setInvoiceIds = setData(['invoicesScreen', 'ids'])

const setInvoice = (id, data) => ({
  type: 'SET INVOICES',
  path: ['invoices', id],
  payload: data,
  reducer: (state, data) => ({
    ...state,
    ...data,
  }),
})

const setInvoices = (data) => ({
  type: 'SET INVOICES',
  path: ['invoices'],
  payload: data,
  reducer: (state, data) => ({
    ...state,
    ...data,
  }),
})

const setInvoiceNotFound = (id) => ({
  type: 'SET INVOICE NOT FOUND',
  path: ['invoices', id],
  payload: null,
  reducer: () => ({
    notFound: true,
    data: '',
    attachments: [],
  }),
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

    try {
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
    } catch (error) {
      await swal({
        title: 'Invoices could not be fetched',
        text: error.message,
        icon: 'error',
      })
    }
  }
)

export const getInvoiceDetail = (id) => loadingWrapper(
  async (dispatch, getState, {api}) => {
    try {
      const invoiceDetail = await api.getInvoiceDetail(id)
      dispatch(setInvoice(id, {data: invoiceDetail}))
    } catch (error) {
      if (error.statusCode === 404) {
        dispatch(setInvoiceNotFound(id))
      } else {
        await swal({
          title: `Invoice ${id} could not be fetched`,
          text: error.message,
          icon: 'error',
        })
      }
    }
  }
)

export const getInvoiceMeta = (id) => loadingWrapper(
  async (dispatch, getState, {api}) => {
    try {
      const meta = await api.getInvoiceMeta(id)
      dispatch(setInvoice(id, meta))
    } catch (error) {
      if (error.statusCode === 404) {
        dispatch(setInvoiceNotFound(id))
      } else {
        await swal({
          title: `Invoice ${id} could not be fetched`,
          text: error.message,
          icon: 'error',
        })
      }
    }
  }
)

export const createInvoice = (data) => loadingWrapper(
  async (dispatch, getState, {api}) => {
    try {
      const invoice = await api.createInvoice(data)
      dispatch(setInvoices({
        [invoice.id]: invoice,
      }))

      await swal({
        title: 'Invoice was created',
        icon: 'success',
      })
      return invoice.id
    } catch (error) {
      await swal({
        title: 'Invoice could not be created',
        text: error.message,
        icon: 'error',
      })
      return null
    }
  }
)

export const addAttachment = (tab, attachment) => ({
  type: 'ADD ATTACHMENT',
  path: ['createInvoiceScreen', tab, 'attachments'],
  payload: attachment,
  reducer: (state, attachment) => [...state, attachment],
})

export const removeAttachment = (tab, attachment) => ({
  type: 'REMOVE ATTACHMENT',
  path: ['createInvoiceScreen', tab, 'attachments'],
  payload: attachment,
  reducer: (state, attachment) => state.filter((a) => a !== attachment),
})

export const clearAttachments = (tab) => ({
  type: 'CLEAR ATTACHMENTS',
  path: ['createInvoiceScreen', tab, 'attachments'],
  reducer: (state) => [],
})
