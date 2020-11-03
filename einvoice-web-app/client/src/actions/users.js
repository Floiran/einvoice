import {loadingWrapper, setData} from './common'

export const LOGGING = 'logging'
export const LOGGING_FAILED = 'failed'
export const LOGGED_IN = 'loggedIn'
export const LOGGED_OUT = 'loggedOut'

export const setLoggingStatus = setData(['loggingStatus'])
const setUser = setData(['user'])

export const getMyInfo = () => (
  async (dispatch, getState, {api}) => {
    const userString = localStorage.getItem('user')
    let user

    try {
      user = userString && JSON.parse(userString)
    } catch (error) {
      localStorage.removeItem('user')
    }

    if (user) {
      try {
        const userData = await api.getUserInfo(user)
        dispatch(setUser(userData))
        localStorage.setItem('user', JSON.stringify(userData))
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
    const updateData = await api.updateUser(getState().user, data)
    // TODO: fixme
    const userData = {...getState().user, ...updateData}
    dispatch(setUser(userData))
    localStorage.setItem('user', JSON.stringify(userData))
  }
)

export const loginWithSlovenskoSkToken = (token) => (
  async (dispatch, getState, {api}) => {
    try {
      dispatch(setLoggingStatus(LOGGING))
      const userData = await api.loginWithSlovenskoSkToken(token)
      dispatch(setUser(userData))
      localStorage.setItem('user', JSON.stringify(userData))
      dispatch(setLoggingStatus(LOGGED_IN))
      return true
    } catch (error) {
      dispatch(setLoggingStatus(LOGGING_FAILED))
      if (error.statusCode === 401) {
        dispatch(setUser(null))
      }
      return false
    }
  }
)

export const logout = () => (
  async (dispatch, getState, {api}) => {
    await api.logout(getState().user)
    dispatch(setUser(null))
    localStorage.removeItem('user')
    dispatch(setLoggingStatus(LOGGED_OUT))
  }
)
