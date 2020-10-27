import ApiError from './ApiError'

// TODO: user token should be cookie and sent automatically with all requests in future
export default class Api {

  validateResponse = ({status, body}) => {
    if (status < 200 || status >= 300) {
      throw new ApiError({statusCode: status, message: body && body.error})
    }
  }

  getApiUrl = async () => {
    const urls = await this.standardRequest({route: '/api/urls'})
    this.baseUrl = urls.apiServerUrl
    return urls
  }

  getUserInfo = async (user) =>
    await this.apiRequest({
      route: '/users/me',
      headers: {
        Authorization: user.token
      }
    })

  updateUser = async (user, data) =>
    await this.apiRequest({
      method: "PUT",
      route: '/users/me',
      data,
      headers: {
        Authorization: user.token
      }
    })

  loginWithSlovenskoSkToken = async (token) =>
    await this.apiRequest({
      route: '/login',
      headers: {
        Authorization: token
      }
    })

  logout = async (user) =>
    await this.apiRequest({
      route: '/logout',
      headers: {
        Authorization: user.token
      },
      jsonResponse: false
    })

  getInvoices = async (formats) => {
    const queryParams = formats.map((f) => ['format', f])
    return await this.apiRequest({
      route: `/invoices?${new URLSearchParams(queryParams)}`
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

  createInvoice = async (user, formData) =>
    await this.apiRequest({
      method: 'POST',
      route: `/invoices`,
      data: formData,
      headers: {
        'Authorization': user.token,
      },
      jsonBody: false,
    })

  apiRequest = (params) =>
    this.standardRequest({
      ...params,
      route: `${this.baseUrl}${params.route}`,
    })

  //TODO: ideally all request & response bodies are jsons
  async standardRequest({method, data, route, jsonResponse = true, jsonBody = true, ...opts}) {
    let contentType = {}
    if(jsonBody) {
      contentType = {'Content-Type':  'application/json'}
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
