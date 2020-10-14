const setUser = (user) => ({
  type: 'SET USER',
  path: ['user'],
  payload: user,
  reducer: (state, user) => user,
})

export const LOGGING = "logging"
export const LOGGING_FAILED = "failed"
export const LOGGED_IN = "loggedIn"
export const LOGGED_OUT = "loggedOut"

export const setLoggingStatus = (status) => ({
  type: 'SET LOGGING STATUS',
  path: ['loggingStatus'],
  payload: status,
  reducer: (state, status) => status,
})

export const getMyInfo = (location, history) => (
  async (dispatch, getState, {api}) => {
    let userString = localStorage.getItem('user')
    let user

    try {
      user = userString && JSON.parse(userString)
    } catch (error) {
      localStorage.removeItem('user')
    }

    if (user) {
      try {
        let userData = await api.getUserInfo(user)
        dispatch(setUser(userData))
        localStorage.setItem('user', JSON.stringify(userData))
      } catch (error) {
        if (error.statusCode === 401) {
          dispatch(setUser(null))
        }
      }
    }

    let urlParams = new URLSearchParams(location.search)

    if(urlParams.has('token')) {
      await dispatch(loginWithSlovenskoSkToken(history, urlParams.get('token')))
    }
  }
)

export const login = (history) => (
  async (dispatch, getState, {api}) => {
    try {
      dispatch(setLoggingStatus(LOGGING))
      await api.login(history)
    } catch (error) {
      dispatch(setLoggingStatus(LOGGING_FAILED))
    }
  }
)

export const updateUser = (data) => (
  async (dispatch, getState, {api}) => {
    let userData = await api.updateUser(getState().user, data)
    dispatch(setUser(userData))
    localStorage.setItem('user', JSON.stringify(userData))
  }
)

export const loginWithSlovenskoSkToken = (history, token) => (
  async (dispatch, getState, {api}) => {
    try {
      dispatch(setLoggingStatus(LOGGING))
      let userData = await api.loginWithSlovenskoSkToken(token)
      history.push("/account")
      dispatch(setUser(userData))
      localStorage.setItem('user', JSON.stringify(userData))
      dispatch(setLoggingStatus(LOGGED_IN))
    } catch (error) {
      dispatch(setLoggingStatus(LOGGING_FAILED))
      if (error.statusCode === 401) {
        dispatch(setUser(null))
      }
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