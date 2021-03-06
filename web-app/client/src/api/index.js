import {CONFIG} from '../appSettings'
import ApiError from './ApiError'

export default class Api {

  validateResponse = ({status, body}) => {
    if (status < 200 || status >= 300) {
      throw new ApiError({statusCode: status, message: body && body.error})
    }
  }

  getUserInfo = async () =>
    await this.apiRequest({
      route: '/users/me',
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
    })

  updateUser = async (data) =>
    await this.apiRequest({
      method: 'PATCH',
      route: '/users/me',
      data,
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
    })

  loginWithSlovenskoSkToken = async (token) =>
    await this.apiRequest({
      route: '/login',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })

  logout = async () =>
    await this.apiRequest({
      route: '/logout',
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
      jsonResponse: false,
    })

  getInvoices = async (formats) => {
    const queryParams = formats.map((f) => ['format', f])
    return await this.apiRequest({
      route: `/invoices?${new URLSearchParams(queryParams)}`,
    })
  }

  getInvoiceDetail = async (id) =>
    await this.apiRequest({
      route: `/invoices/${id}/detail`,
      jsonResponse: false,
    })

  getInvoiceMeta = async (id) => {
    return await this.apiRequest({
      route: `/invoices/${id}`,
    })
  }

  createInvoice = async (formData) =>
    await this.apiRequest({
      method: 'POST',
      route: '/invoices',
      data: formData,
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
      jsonBody: false,
    })

  apiRequest = (params) =>
    this.standardRequest({
      ...params,
      route: `${CONFIG.authServerUrl}${params.route}`,
    })

  //TODO: ideally all request & response bodies are jsons
  async standardRequest({method, data, route, jsonResponse = true, jsonBody = true, ...opts}) {
    let contentType = {}
    if (jsonBody) {
      contentType = {'Content-Type': 'application/json'}
    }
    const response = await fetch(route, {
      method,
      body: jsonBody ? JSON.stringify(data) : data,
      ...opts,
      headers: {
        ...contentType,
        ...opts.headers,
      },
    })

    // TODO: API should be unified and return only one type of responses
    const body = jsonResponse ? await response.json() : await response.text()
    this.validateResponse({status: response.status, body})
    return body
  }
}
