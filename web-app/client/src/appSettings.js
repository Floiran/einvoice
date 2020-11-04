export const CONFIG = process.env.NODE_ENV === 'development' ?
  {
    authServerUrl: process.env.REACT_APP_AUTHSERVER_URL,
    slovenskoSkLoginUrl: process.env.REACT_APP_SLOVENSKOSK_LOGIN_URL
  } :
  JSON.parse(document.getElementsByTagName('body')[0].dataset.einvoiceconfig)
