import {loadingWrapper, setData} from './common'

export const LOGGING = 'logging'
export const LOGGING_FAILED = 'failed'
export const LOGGED_IN = 'loggedIn'
export const LOGGED_OUT = 'loggedOut'

export const setLoggingStatus = setData(['loggingStatus'])
const setUser = setData(['user'])

export const getMyInfo = () => (
  async (dispatch, getState, {api}) => {
    if (localStorage.getItem('token')) {
      try {
        const userData = await api.getUserInfo()
        dispatch(setUser(userData))
      } catch (error) {
        dispatch(setUser({unauthorized: true}))
      }
    } else {
      dispatch(setUser({unauthorized: true}))
    }
  }
)

export const updateUser = (data) => loadingWrapper(
  async (dispatch, getState, {api}) => {
    const userData = await api.updateUser(data)
    dispatch(setUser(userData))
  }
)

export const loginWithSlovenskoSkToken = (token) => (
  async (dispatch, getState, {api}) => {
    try {
      dispatch(setLoggingStatus(LOGGING))
      const userData = await api.loginWithSlovenskoSkToken(token)
      dispatch(setUser(userData))
      localStorage.setItem('token', userData.token)
      dispatch(setLoggingStatus(LOGGED_IN))
      return true
    } catch (error) {
      dispatch(setLoggingStatus(LOGGING_FAILED))
      if (error.statusCode === 401) {
        dispatch(setUser(null))
      }
      localStorage.removeItem('token')
      return false
    }
  }
)

export const logout = () => (
  async (dispatch, getState, {api}) => {
    await api.logout()
    dispatch(setUser(null))
    localStorage.removeItem('token')
    dispatch(setLoggingStatus(LOGGED_OUT))
  }
)
